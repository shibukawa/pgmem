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
	F_errmsg(m, int32(142890), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(470610), int32(2098), int32(70401))
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
	F_errmsg(m, int32(232906), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	F_errhint(m, int32(534683), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(476708), int32(1648), int32(98701))
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
	v5 = F_DirectFunctionCall2Coll(m, int32(1559), int32(100), l0, l1)
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
	v6 = F_DirectFunctionCall2Coll(m, int32(1558), v3, v4, v5)
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
	v4 = F_DirectFunctionCall1Coll(m, int32(1585), int32(100), l0)
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
						F_errmsg(m, int32(232906), int32(0))
						mBase = m.M
						v151 = m.ExcPending
						if v151 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(534683), int32(0))
							mBase = m.M
							v155 = m.ExcPending
							if v155 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(476708), int32(1648), int32(98701))
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
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v761 int32
	_ = v761
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v788 int32
	_ = v788
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v813 int32
	_ = v813
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var __phi940 int32
	_ = __phi940
	var v941 int32
	_ = v941
	var __phi941 int32
	_ = __phi941
	var v942 int32
	_ = v942
	var __phi942 int32
	_ = __phi942
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1019 int64
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1076 int32
	_ = v1076
	var v1093 int32
	_ = v1093
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1115 int32
	_ = v1115
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1134 int64
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1154 int32
	_ = v1154
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1261 int64
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1363 int32
	_ = v1363
	var v1368 int32
	_ = v1368
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1385 int32
	_ = v1385
	var v1390 int32
	_ = v1390
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1401 int32
	_ = v1401
	var v1406 int32
	_ = v1406
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1429 int32
	_ = v1429
	var v1434 int32
	_ = v1434
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1457 int32
	_ = v1457
	var v1461 int32
	_ = v1461
	var v1466 int32
	_ = v1466
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1479 int32
	_ = v1479
	var v1484 int32
	_ = v1484
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
	v52 = int32(367154)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, _consts[911])))
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
	v420 = int32(16346)
	v423 = int32(*(*uint8)(unsafe.Add(mBase, _consts[912])))
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
	v89 = F_get_tsearch_config_filename(m, v86, int32(145331))
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
	F_errmsg(m, int32(205441), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(470495), int32(212), int32(442778))
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
	F_errmsg(m, int32(358240), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(470495), int32(262), int32(442778))
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
	F_errmsg(m, int32(358240), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(470495), int32(278), int32(442778))
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
	F_errmsg(m, int32(355819), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(470495), int32(287), int32(442778))
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
	F_errmsg(m, int32(11170), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(470495), int32(292), int32(442778))
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
	F_errmsg(m, int32(283084), v19+int32(80))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(470495), int32(180), int32(442778))
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
	F_errmsg(m, int32(123235), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(470495), int32(616), int32(93710))
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
	F_errmsg(m, int32(123046), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(470495), int32(625), int32(93710))
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
	F_errmsg(m, int32(683504), v19+int32(96))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(470495), int32(633), int32(93710))
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
	F_errmsg(m, int32(206051), int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(470495), int32(640), int32(93710))
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
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L1
	} else {
		goto L357
	}
L160:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L1
	} else {
		goto L352
	}
L161:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L1
	} else {
		goto L348
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
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L1
	} else {
		goto L344
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
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v922 != 0 {
		goto L233
	} else {
		goto L234
	}
L170:
	;
	v908 = v581
	v913 = v593
	v914 = int32(16)
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
	v908 = v878
	v913 = v883
	v914 = v884
	goto L169
L175:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v892+v617)))
	F_pfree(m, v894)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L1
	} else {
		goto L230
	}
L176:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v662 = int32(0)
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v659)+44))
	if v620&int32(3) == v662 {
		v687 = v620
		goto L186
	} else {
		goto L187
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
	v878 = v602 + v646
	v883 = v633
	v884 = v634
	goto L175
L184:
	;
	v722 = F_FunctionCall4Coll(m, v659+int32(12), v662, v663, v620, v720, int32(0))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L201
	}
L185:
	;
	v720 = v712 - v620
	goto L184
L186:
	;
	v691 = v687
	goto L195
L187:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620))))
	if v671 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v720 = int32(0)
	goto L184
L189:
	;
	goto L190
L190:
	;
	v676 = v620
	goto L191
L191:
	;
	v680 = v676 + int32(1)
	if v680&int32(3) == int32(0) {
		v687 = v680
		goto L186
	} else {
		goto L193
	}
L192:
	;
	v712 = v680
	goto L185
L193:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680))))
	if v685 != 0 {
		v676 = v680
		goto L191
	} else {
		goto L194
	}
L194:
	;
	goto L192
L195:
	;
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v691)))
	v700 = int32(-2139062144)
	if (int32(16843008)-v697|v697)&v700 == v700 {
		v691 = v691 + int32(4)
		goto L195
	} else {
		goto L197
	}
L196:
	;
	v706 = v691
	goto L198
L197:
	;
	goto L196
L198:
	;
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706))))
	if v710 != 0 {
		v706 = v706 + int32(1)
		goto L198
	} else {
		goto L200
	}
L199:
	;
	v712 = v706
	goto L185
L200:
	;
	goto L199
L201:
	;
	if v722 == int32(0) {
		goto L161
	} else {
		goto L202
	}
L202:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v722)+4))
	if v726 == int32(0) {
		goto L160
	} else {
		goto L203
	}
L203:
	;
	v731 = v602
	v734 = v722
	v736 = v607
	v737 = v608
	goto L204
L204:
	;
	v745 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v734))))
	v746 = int32(1)
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v734)+12))
	if v747 == int32(0) {
		v788 = v746
		goto L206
	} else {
		goto L207
	}
L205:
	;
	v878 = v867
	v883 = v870
	v884 = v871
	goto L175
L206:
	;
	v798 = v734
	v799 = v731
	v800 = v734 + int32(4)
	v804 = v736
	v805 = v737
	goto L212
L207:
	;
	v750 = v734
	v761 = v746
	goto L208
L208:
	;
	v767 = v750 + int32(8)
	v768 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v767))))
	if v768 != v745&int32(65535) {
		v788 = v761
		goto L206
	} else {
		goto L210
	}
L209:
	;
	v788 = v773
	goto L206
L210:
	;
	v773 = v761 + int32(1)
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v750+int32(20))))
	if v776 != 0 {
		v750 = v767
		v761 = v773
		goto L208
	} else {
		goto L211
	}
L211:
	;
	goto L209
L212:
	;
	v813 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v798))))
	if v745&int32(65535) != v813 {
		goto L215
	} else {
		goto L216
	}
L213:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v869)+4))
	if v875 != 0 {
		v731 = v867
		v734 = v869
		v736 = v870
		v737 = v871
		goto L204
	} else {
		goto L229
	}
L214:
	;
	goto L213
L215:
	;
	v867 = v799
	v869 = v798
	v870 = v804
	v871 = v805
	goto L214
L216:
	;
	goto L217
L217:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v815+v617)+4))
	if v805 <= v799 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v821 = F_repalloc(m, v804, v805<<(uint(int32(4))%32))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L221
	}
L219:
	;
	v825 = v804
	v826 = v805
	goto L220
L220:
	;
	v829 = v825 + v799<<(uint(int32(3))%32)
	v831 = v829 + int32(4)
	v833 = F_palloc(m, int32(16))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L1
	} else {
		goto L222
	}
L221:
	;
	v825 = v821
	v826 = v805 << (uint(int32(1)) % 32)
	goto L220
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v831))) = v833
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v800)))
	if v836 == int32(0) {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v829))) = v845
	*(*uint16)(unsafe.Add(mBase, uint32(v844)+6)) = uint16(v846)
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v831)))
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v817)))
	*(*int32)(unsafe.Add(mBase, uint32(v849))) = v850
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v831)))
	v853 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v817)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v852)+4)) = uint16(v853)
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v831)))
	*(*int32)(unsafe.Add(mBase, uint32(v855)+8)) = int32(0)
	v859 = v799 + int32(1)
	v861 = v798 + int32(12)
	v863 = v798 + int32(8)
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v861)))
	if v864 != 0 {
		v798 = v863
		v799 = v859
		v800 = v861
		v804 = v825
		v805 = v826
		goto L212
	} else {
		goto L228
	}
L224:
	;
	v844 = v833
	v845 = int32(0)
	v846 = int32(1)
	goto L223
L225:
	;
	goto L226
L226:
	;
	v841 = F_pstrdup(m, v836)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v831)))
	v844 = v843
	v845 = v841
	v846 = v788
	goto L223
L228:
	;
	v867 = v859
	v869 = v863
	v870 = v825
	v871 = v826
	goto L214
L229:
	;
	goto L205
L230:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v897+v617)+4))
	F_pfree(m, v899)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	v903 = v615 + int32(1)
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v903 < v904 {
		v602 = v878
		v607 = v883
		v608 = v884
		v615 = v903
		goto L173
	} else {
		goto L232
	}
L232:
	;
	goto L174
L233:
	;
	F_pfree(m, v922)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L1
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v914
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v908
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v913
	if int32(2) <= v908 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	goto L235
L237:
	;
	F_pg_qsort(m, v913, v908, int32(8), int32(1164))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L1
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	if int32(0) < v1076 {
		goto L277
	} else {
		goto L278
	}
L240:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v935 < int32(2) {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v1050 = int32(3)
	v1053 = (v1034-v1036)>>(uint(v1050)%32) + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v1053
	v1057 = F_repalloc(m, v1036, v1053<<(uint(v1050)%32))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L1
	} else {
		goto L274
	}
L242:
	;
	v1034 = v934
	v1036 = v934
	goto L241
L243:
	;
	goto L244
L244:
	;
	__phi940 = v934
	__phi941 = v934
	__phi942 = v934 + int32(8)
	v940 = __phi940
	v941 = __phi941
	v942 = __phi942
	goto L245
L245:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v941)))
	v958 = v940 + int32(8)
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v958)))
	if v959 != 0 {
		goto L249
	} else {
		goto L250
	}
L246:
	;
	v1034 = v1021
	v1036 = v1028
	goto L241
L247:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v1027 = v942 + int32(8)
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if (v1027-v1028)>>(uint(int32(3))%32) < v1025 {
		__phi940 = v942
		__phi941 = v1021
		__phi942 = v1027
		v940 = __phi940
		v941 = __phi941
		v942 = __phi942
		goto L245
	} else {
		goto L273
	}
L248:
	;
	v1018 = v941 + int32(8)
	v1019 = *(*int64)(unsafe.Add(mBase, uint32(v958)))
	*(*int64)(unsafe.Add(mBase, uint32(v1018))) = v1019
	v1021 = v1018
	goto L247
L249:
	;
	if v956 == int32(0) {
		goto L248
	} else {
		goto L252
	}
L250:
	;
	v987 = v956
	goto L251
L251:
	;
	if v987 != 0 {
		goto L248
	} else {
		goto L261
	}
L252:
	;
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v956))))
	v965 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v959))))
	if v965 == int32(0) {
		v984 = v964
		v985 = v965
		goto L254
	} else {
		goto L255
	}
L253:
	;
	v987 = v985 - v984
	goto L251
L254:
	;
	goto L253
L255:
	;
	if v964 != v965 {
		v984 = v964
		v985 = v965
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v969 = v959
	v970 = v956
	goto L257
L257:
	;
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970)+1)))
	v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v969)+1)))
	if v974 == int32(0) {
		v984 = v973
		v985 = v974
		goto L254
	} else {
		goto L259
	}
L258:
	;
	v984 = v973
	v985 = v974
	goto L254
L259:
	;
	v977 = int32(1)
	if v973 == v974 {
		v969 = v969 + v977
		v970 = v970 + v977
		goto L257
	} else {
		goto L260
	}
L260:
	;
	goto L258
L261:
	;
	v989 = v940 + int32(12)
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v989)))
	if v990 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v958)))
	if v1012 == int32(0) {
		v1021 = v941
		goto L247
	} else {
		goto L271
	}
L263:
	;
	F_pfree(m, v990)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L1
	} else {
		goto L270
	}
L264:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v941)+4))
	if v993 == int32(0) {
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v990)))
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v993)))
	if v996 != v997 {
		goto L266
	} else {
		goto L267
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v990)+8)) = v993
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v989)))
	*(*int32)(unsafe.Add(mBase, uint32(v941)+4)) = v1006
	goto L262
L267:
	;
	v999 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v990)+4)))
	v1000 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v993)+4)))
	if v999 != v1000 {
		goto L266
	} else {
		goto L268
	}
L268:
	;
	v1002 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v990)+6)))
	v1003 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v993)+6)))
	if v1002 == v1003 {
		goto L263
	} else {
		goto L269
	}
L269:
	;
	goto L266
L270:
	;
	goto L262
L271:
	;
	F_pfree(m, v1012)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	v1021 = v941
	goto L247
L273:
	;
	goto L246
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v1057
	goto L239
L275:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1372 = m.ExcPending
	if v1372 != 0 {
		goto L1
	} else {
		goto L340
	}
L276:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L1
	} else {
		goto L336
	}
L277:
	;
	v1093 = int32(0)
	goto L280
L278:
	;
	goto L279
L279:
	;
	m.G0 = v19 + int32(160)
	return v23
L280:
	;
	v1097 = v1093 << (uint(int32(3)) % 32)
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1097+v1098)+4))
	v1102 = F_palloc(m, int32(16))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L1
	} else {
		goto L282
	}
L281:
	;
	goto L279
L282:
	;
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1104+v1097)+4)) = v1102
	v1107 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1102)+4)) = v1107
	if v1100 == v1107 {
		v1297 = v1102
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v1314 = v1313 + v1097
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1314)+4))
	if v1297 == v1315 {
		goto L159
	} else {
		goto L333
	}
L284:
	;
	v1115 = int32(2)
	v1119 = v1100
	v1122 = v1102
	goto L285
L285:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+4))
	if v1128 == int32(0) {
		goto L287
	} else {
		goto L288
	}
L286:
	;
	v1297 = v1268
	goto L283
L287:
	;
	v1297 = v1122
	goto L283
L288:
	;
	goto L289
L289:
	;
	v1131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1119)+3)))
	if v1131&int32(16) != 0 {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1209)+4))
	if v1210 == int32(0) {
		goto L276
	} else {
		goto L313
	}
L291:
	;
	v1134 = *(*int64)(unsafe.Add(mBase, uint32(v1119)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+112)) = v1134
	v1136 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+124)) = v1136
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+114)) = uint16(v1136)
	v1209 = v19 + int32(112)
	goto L290
L292:
	;
	goto L293
L293:
	;
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v1145 = int32(0)
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1142)+44))
	if v1128&int32(3) == v1145 {
		v1170 = v1128
		goto L296
	} else {
		goto L297
	}
L294:
	;
	v1205 = F_FunctionCall4Coll(m, v1142+int32(12), v1145, v1146, v1128, v1203, int32(0))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L1
	} else {
		goto L311
	}
L295:
	;
	v1203 = v1195 - v1128
	goto L294
L296:
	;
	v1174 = v1170
	goto L305
L297:
	;
	v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1128))))
	if v1154 == int32(0) {
		goto L298
	} else {
		goto L299
	}
L298:
	;
	v1203 = int32(0)
	goto L294
L299:
	;
	goto L300
L300:
	;
	v1159 = v1128
	goto L301
L301:
	;
	v1163 = v1159 + int32(1)
	if v1163&int32(3) == int32(0) {
		v1170 = v1163
		goto L296
	} else {
		goto L303
	}
L302:
	;
	v1195 = v1163
	goto L295
L303:
	;
	v1168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1163))))
	if v1168 != 0 {
		v1159 = v1163
		goto L301
	} else {
		goto L304
	}
L304:
	;
	goto L302
L305:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1174)))
	v1183 = int32(-2139062144)
	if (int32(16843008)-v1180|v1180)&v1183 == v1183 {
		v1174 = v1174 + int32(4)
		goto L305
	} else {
		goto L307
	}
L306:
	;
	v1189 = v1174
	goto L308
L307:
	;
	goto L306
L308:
	;
	v1193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1189))))
	if v1193 != 0 {
		v1189 = v1189 + int32(1)
		goto L308
	} else {
		goto L310
	}
L309:
	;
	v1195 = v1189
	goto L295
L310:
	;
	goto L309
L311:
	;
	if v1205 == int32(0) {
		goto L275
	} else {
		goto L312
	}
L312:
	;
	v1209 = v1205
	goto L290
L313:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1215+v1097)+4))
	v1221 = v1122
	v1222 = v1209
	v1223 = v1209 + int32(4)
	v1224 = v1115
	goto L314
L314:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1237+v1097)+4))
	v1240 = v1221 - v1239
	if v1224 <= v1240>>(uint(int32(3))%32)+int32(1) {
		goto L316
	} else {
		goto L317
	}
L315:
	;
	if v1122 == v1217 {
		goto L322
	} else {
		goto L323
	}
L316:
	;
	v1248 = F_repalloc(m, v1239, v1224<<(uint(int32(4))%32))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L1
	} else {
		goto L319
	}
L317:
	;
	v1259 = v1221
	v1260 = v1224
	goto L318
L318:
	;
	v1261 = *(*int64)(unsafe.Add(mBase, uint32(v1222)))
	*(*int64)(unsafe.Add(mBase, uint32(v1259))) = v1261
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1223)))
	v1264 = F_pstrdup(m, v1263)
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L1
	} else {
		goto L320
	}
L319:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1250+v1097)+4)) = v1248
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(v1255+v1097)+4))
	v1259 = v1257 + v1240
	v1260 = v1224 << (uint(int32(1)) % 32)
	goto L318
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1259)+4)) = v1264
	v1267 = int32(8)
	v1268 = v1259 + v1267
	v1270 = v1222 + int32(12)
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1270)))
	if v1273 != 0 {
		v1221 = v1268
		v1222 = v1222 + v1267
		v1223 = v1270
		v1224 = v1260
		goto L314
	} else {
		goto L321
	}
L321:
	;
	goto L315
L322:
	;
	v1276 = int32(-1)
	goto L324
L323:
	;
	v1276 = (v1122 - v1217) >> (uint(int32(3)) % 32)
	goto L324
L324:
	;
	if int32(0) < v1276 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1279+v1097)+4))
	v1286 = v1281 + v1276<<(uint(int32(3))%32) + int32(2)
	v1287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1286))))
	v1289 = v1287 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v1286))) = uint16(v1289)
	goto L327
L326:
	;
	goto L327
L327:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+4))
	if v1292 != 0 {
		goto L328
	} else {
		goto L329
	}
L328:
	;
	F_pfree(m, v1292)
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L1
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	v1296 = v1119 + int32(8)
	if v1296 != 0 {
		v1115 = v1260
		v1119 = v1296
		v1122 = v1268
		goto L285
	} else {
		goto L332
	}
L331:
	;
	goto L330
L332:
	;
	goto L286
L333:
	;
	v1319 = int32(base.Ui32(v1297-v1315) >> (uint(int32(3)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v1314)+2)) = uint16(v1319)
	F_pfree(m, v1100)
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L1
	} else {
		goto L334
	}
L334:
	;
	v1324 = v1093 + int32(1)
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	if v1324 < v1325 {
		v1093 = v1324
		goto L280
	} else {
		goto L335
	}
L335:
	;
	goto L281
L336:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L1
	} else {
		goto L337
	}
L337:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v1093 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v1354
	F_errmsg(m, int32(639279), v19+int32(32))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	F_errfinish(m, int32(470495), int32(568), int32(331550))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L1
	} else {
		goto L339
	}
L339:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L340:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L1
	} else {
		goto L341
	}
L341:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v1093 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v1376
	F_errmsg(m, int32(639133), v19+int32(16))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	F_errfinish(m, int32(470495), int32(575), int32(331550))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L1
	} else {
		goto L343
	}
L343:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L344:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L1
	} else {
		goto L345
	}
L345:
	;
	F_errmsg(m, int32(205790), int32(0))
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L1
	} else {
		goto L346
	}
L346:
	;
	F_errfinish(m, int32(470495), int32(644), int32(93710))
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L348:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L1
	} else {
		goto L349
	}
L349:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v1417 = v1414 + v615<<(uint(int32(3))%32)
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v1417)+4))
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v1418)))
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v1417)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v1420
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v1419 + int32(1)
	F_errmsg(m, int32(639208), v19+int32(48))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L1
	} else {
		goto L350
	}
L350:
	;
	F_errfinish(m, int32(470495), int32(418), int32(358273))
	mBase = m.M
	v1434 = m.ExcPending
	if v1434 != 0 {
		goto L1
	} else {
		goto L351
	}
L351:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L352:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L1
	} else {
		goto L353
	}
L353:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v1445 = v1442 + v615<<(uint(int32(3))%32)
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1445)+4))
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1446)))
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1445)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v1448
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v1447 + int32(1)
	F_errmsg(m, int32(639335), v19-int32(-64))
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L1
	} else {
		goto L354
	}
L354:
	;
	F_errhint(m, int32(591656), int32(0))
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L1
	} else {
		goto L355
	}
L355:
	;
	F_errfinish(m, int32(470495), int32(425), int32(358273))
	mBase = m.M
	v1466 = m.ExcPending
	if v1466 != 0 {
		goto L1
	} else {
		goto L356
	}
L356:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L357:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L1
	} else {
		goto L358
	}
L358:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v1093 + int32(1)
	F_errmsg(m, int32(639086), v19)
	mBase = m.M
	v1479 = m.ExcPending
	if v1479 != 0 {
		goto L1
	} else {
		goto L359
	}
L359:
	;
	F_errfinish(m, int32(470495), int32(587), int32(331550))
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L1
	} else {
		goto L360
	}
L360:
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
	v22 = F_pg_snprintf(m, v7+v14, int32(32), int32(629736), v7)
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
	v94 = int64(*(*int32)(unsafe.Add(mBase, uint32(v89<<(uint(int32(2))%32))+uint32(_consts[1232]))))
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
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v264<<(uint(int32(2))%32))+uint32(_consts[1232])))
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
	*(*int32)(unsafe.Add(mBase, _consts[1233])) = v288
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
	v313 = *(*int32)(unsafe.Add(mBase, uint32(base.B2i32(v306 == int32(0))<<(uint(int32(2))%32))+uint32(_consts[1232])))
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
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v325<<(uint(int32(2))%32))+uint32(_consts[1232])))
	v285 = v285 - v334
	v288 = v288 + int32(1)
	goto L66
L80:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1234])) = v285
	*(*int32)(unsafe.Add(mBase, _consts[1233])) = v288 - int32(1900)
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
	*(*uint32)(unsafe.Add(mBase, _consts[1235])) = uint32(v379)
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
	*(*int32)(unsafe.Add(mBase, _consts[1236])) = v392
	v398 = base.I32_wrap_i64(v234 - v379*int64(3600))
	v399 = int32(65535)
	v401 = int32(60)
	v402 = base.I32_div_u_s(v398&v399, v401)
	*(*int32)(unsafe.Add(mBase, _consts[1237])) = v402
	*(*int32)(unsafe.Add(mBase, _consts[1238])) = v56 + (v398-v402*v401)&v399
	if v298 != 0 {
		v420 = int32(0)
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v422 = v420 * int32(48)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v422)+uint32(_consts[1239])))
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
	*(*int32)(unsafe.Add(mBase, _consts[1240])) = l1
	*(*int32)(unsafe.Add(mBase, _consts[1241])) = v451
	*(*int32)(unsafe.Add(mBase, _consts[1242])) = v447 + int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[1243])) = int32(0)
	return int32(4442920)
L94:
	;
	v439 = v427 - v430
	v441 = v431 + int32(1)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v422+int32(1790880)+v441<<(uint(int32(2))%32))))
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
	F_errmsg(m, int32(11905), v11)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(475474), int32(590), int32(84664))
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
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
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
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v432 int32
	_ = v432
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v485 int32
	_ = v485
	var v505 int32
	_ = v505
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v578 int32
	_ = v578
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v723 int32
	_ = v723
	var v731 int32
	_ = v731
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v838 int32
	_ = v838
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v865 int32
	_ = v865
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v892 int32
	_ = v892
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v924 int32
	_ = v924
	var v940 int32
	_ = v940
	var v951 int32
	_ = v951
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1060 int32
	_ = v1060
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1116 int32
	_ = v1116
	var v1138 int32
	_ = v1138
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1205 int32
	_ = v1205
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1291 int32
	_ = v1291
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1361 int32
	_ = v1361
	var v1375 int32
	_ = v1375
	var v1378 int32
	_ = v1378
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	v24 = m.G0
	v26 = v24 - int32(80)
	m.G0 = v26
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = int32(798)
	v30 = int32(4435480)
	v31 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v26 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v26 + int32(24)
	v44 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v49 = F_AllocSetContextCreateInternal(m, v44, int32(367563), int32(0), int32(1024), int32(8192))
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
	v51 = int32(4442576)
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
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v1378
	F_MemoryContextDelete(m, v1375)
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L1
	} else {
		goto L271
	}
L8:
	;
	if int32(base.Ui32(v68)>>(uint(int32(4))%32))&int32(1) != 0 {
		v1361 = v26
		v1375 = v49
		v1378 = v52
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
	v1361 = v995
	v1375 = v1009
	v1378 = v1012
	goto L7
L16:
	;
	if int32(base.Ui32(v104)>>(uint(int32(5))%32))&int32(1) != 0 {
		v1361 = v81
		v1375 = v95
		v1378 = v98
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
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v77)+76))
	if v321 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L24:
	;
	if v124 == int32(0) {
		v315 = v121
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
	if v151&int32(3) == int32(0) {
		v175 = v151
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v315 = v292
	goto L23
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+48)) = v275
	if v275 <= v134 {
		v315 = v145
		goto L23
	} else {
		goto L54
	}
L29:
	;
	if v208 <= int32(0) {
		goto L46
	} else {
		goto L47
	}
L30:
	;
	v208 = v200 - v151
	goto L29
L31:
	;
	v179 = v175
	goto L40
L32:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v159 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v208 = int32(0)
	goto L29
L34:
	;
	goto L35
L35:
	;
	v164 = v151
	goto L36
L36:
	;
	v168 = v164 + int32(1)
	if v168&int32(3) == int32(0) {
		v175 = v168
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v200 = v168
	goto L30
L38:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v173 != 0 {
		v164 = v168
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v188 = int32(-2139062144)
	if (int32(16843008)-v185|v185)&v188 == v188 {
		v179 = v179 + int32(4)
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v194 = v179
	goto L43
L42:
	;
	goto L41
L43:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v198 != 0 {
		v194 = v194 + int32(1)
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v200 = v194
	goto L30
L45:
	;
	goto L44
L46:
	;
	v275 = v208
	goto L28
L47:
	;
	goto L48
L48:
	;
	v217 = v208
	goto L49
L49:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217+(v151-int32(1))))))
	switch v237 - int32(10) {
	case 0, 3:
		goto L52
	default:
		v249 = v217
		goto L51
	}
L50:
	;
	v275 = v249
	goto L28
L51:
	;
	goto L50
L52:
	;
	v240 = int32(0)
	v241 = int32(1)
	v242 = v217 - v241
	*(*uint8)(unsafe.Add(mBase, uint32(v151+v242))) = uint8(v240)
	if v241 < v217 {
		v217 = v242
		goto L49
	} else {
		goto L53
	}
L53:
	;
	v249 = v240
	goto L51
L54:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v81)+44))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278+v275-int32(1)))))
	if v282 != int32(92) {
		v315 = v145
		goto L23
	} else {
		goto L55
	}
L55:
	;
	v285 = int32(1)
	v286 = v275 - v285
	*(*int32)(unsafe.Add(mBase, uint32(v81)+48)) = v286
	v289 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v286+v278))) = uint8(v289)
	v292 = v145 + v285
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v81)+48))
	v296 = F_pg_get_line_append(m, v77, v81+int32(44))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	if v296 != 0 {
		v134 = v293
		v145 = v292
		goto L26
	} else {
		goto L57
	}
L57:
	;
	goto L27
L58:
	;
	if int32(base.Ui32(v326)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	goto L58
L60:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v326 = v324
	goto L59
L61:
	;
	goto L62
L62:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v326 = v325
	goto L59
L63:
	;
	v332 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v334 = F_errstart(m, v79, int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v81)+20))
	v356 = int32(0)
	v357 = base.B2i32(v355 == v356)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v81)+44))
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359))))
	if v360 == v356 {
		v990 = v76
		v991 = v77
		v992 = v78
		v993 = v79
		v995 = v81
		v996 = v357
		v1003 = v356
		v1004 = v90
		v1006 = v92
		v1007 = v315
		v1009 = v95
		v1012 = v98
		goto L74
	} else {
		goto L75
	}
L66:
	;
	if v334 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v332
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v76
	v353 = F_psprintf(m, int32(284781), v81)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L73
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = v76
	F_errmsg(m, int32(284781), v81+int32(16))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(476740), int32(769), int32(367563))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	v1361 = v81
	v1375 = v95
	v1378 = v98
	goto L7
L74:
	;
	if v996 != 0 {
		goto L189
	} else {
		goto L190
	}
L75:
	;
	if v355 != 0 {
		v990 = v76
		v991 = v77
		v992 = v78
		v993 = v79
		v995 = v81
		v996 = v357
		v1003 = v356
		v1004 = v90
		v1006 = v92
		v1007 = v315
		v1009 = v95
		v1012 = v98
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v363 = v76
	v364 = v77
	v365 = v78
	v366 = v79
	v368 = v81
	v372 = v359
	v376 = v356
	v377 = v90
	v379 = v92
	v380 = v315
	v382 = v95
	v385 = v98
	goto L77
L77:
	;
	F_initStringInfo(m, v368+int32(60))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L79
	}
L78:
	;
	v990 = v363
	v991 = v364
	v992 = v365
	v993 = v366
	v995 = v368
	v996 = v984
	v1003 = v981
	v1004 = v377
	v1006 = v379
	v1007 = v380
	v1009 = v382
	v1012 = v385
	goto L74
L79:
	;
	v398 = int32(0)
	v400 = v372
	goto L80
L80:
	;
	v415 = v368 + int32(60)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	v417 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v416))) = uint8(v417)
	*(*int32)(unsafe.Add(mBase, uint32(v415)+12)) = v417
	*(*int32)(unsafe.Add(mBase, uint32(v415)+4)) = v417
	goto L82
L81:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v368)+60))
	F_pfree(m, v967)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L181
	}
L82:
	;
	v432 = v400
	goto L84
L83:
	;
	v457 = int32(0)
	v465 = v446
	v467 = v457
	v470 = v449
	v471 = v456
	v473 = v457
	v476 = v457
	goto L90
L84:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432))))
	v447 = int32(1)
	v449 = v432 + v447
	switch v446 {
	case 0:
		v456 = v446
		goto L83
	default:
		goto L87
	case 9, 13, 32:
		v451 = v447
		goto L86
	}
L85:
	;
	v456 = int32(0)
	goto L83
L86:
	;
	if v446 == int32(44) {
		v432 = v449
		goto L84
	} else {
		goto L88
	}
L87:
	;
	v451 = int32(0)
	goto L86
L88:
	;
	if v451 != 0 {
		v432 = v449
		goto L84
	} else {
		goto L89
	}
L89:
	;
	goto L85
L90:
	;
	v485 = v465 & int32(255)
	switch v485 {
	case 0:
		v578 = v470
		v589 = v457
		goto L92
	default:
		goto L96
	case 9, 13, 32:
		goto L97
	}
L91:
	;
	v592 = int32(1)
	v593 = v578 - v592
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v368)+64))
	if (v476|base.B2i32(int32(0) < v594))&v592 != 0 {
		goto L114
	} else {
		goto L115
	}
L92:
	;
	goto L91
L93:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v470))))
	v465 = v566
	v467 = v561
	v470 = v470 + int32(1)
	v471 = v563
	v473 = v562
	v476 = v565
	goto L90
L94:
	;
	v555 = int32(1)
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v368)+64))
	if v557 != 0 {
		goto L111
	} else {
		goto L112
	}
L95:
	;
	F_appendStringInfoChar(m, v368+int32(60), base.I32_extend8_s(v465))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L109
	}
L96:
	;
	if (base.B2i32(v485 != int32(35))|v467)&int32(1) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	if v467&int32(1) != 0 {
		v539 = int32(0)
		goto L95
	} else {
		goto L98
	}
L98:
	;
	v578 = v470
	v589 = v457
	goto L92
L99:
	;
	v505 = v470
	goto L102
L100:
	;
	goto L101
L101:
	;
	if (base.B2i32(v485 != int32(44))|v467)&int32(1) == int32(0) {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505))))
	v521 = v505 + int32(1)
	if v519 != 0 {
		v505 = v521
		goto L102
	} else {
		goto L104
	}
L103:
	;
	v578 = v521
	v589 = v457
	goto L92
L104:
	;
	goto L103
L105:
	;
	v578 = v470
	v589 = int32(1)
	goto L92
L106:
	;
	goto L107
L107:
	;
	if (base.B2i32(v485 != int32(34))|v473)&int32(1) == int32(0) {
		v554 = v467
		goto L94
	} else {
		goto L108
	}
L108:
	;
	v539 = base.B2i32(v485 == int32(34))
	goto L95
L109:
	;
	v545 = int32(0)
	if v539 == v545 {
		v561 = v467
		v562 = v545
		v563 = v471
		v565 = v476
		goto L93
	} else {
		goto L110
	}
L110:
	;
	v554 = v467 & v539 & (v473 ^ int32(1))
	goto L94
L111:
	;
	v558 = v471
	goto L113
L112:
	;
	v558 = v555
	goto L113
L113:
	;
	v561 = v467 ^ int32(1)
	v562 = v554
	v563 = v558
	v565 = v555
	goto L93
L114:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v368)+60))
	v602 = v471 & int32(1)
	if v602 != 0 {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	v951 = v398
	goto L116
L116:
	;
	goto L81
L117:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v368)+20))
	if v589&base.B2i32(v940 == int32(0)) != 0 {
		v398 = v924
		v400 = v593
		goto L80
	} else {
		goto L180
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v368)+76)) = int32(0)
	v693 = F_AbsoluteConfigLocation(m, v600+int32(1), v363)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L146
	}
L119:
	;
	v608 = int32(4442576)
	v609 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v612 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v612
	if v600&int32(3) == int32(0) {
		v637 = v600
		goto L125
	} else {
		goto L126
	}
L120:
	;
	if v594 < int32(2) {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600))))
	if v605 == int32(64) {
		goto L118
	} else {
		goto L122
	}
L122:
	;
	goto L119
L123:
	;
	v673 = F_palloc0(m, v670+int32(13))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L140
	}
L124:
	;
	v670 = v662 - v600
	goto L123
L125:
	;
	v641 = v637
	goto L134
L126:
	;
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600))))
	if v621 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v670 = int32(0)
	goto L123
L128:
	;
	goto L129
L129:
	;
	v626 = v600
	goto L130
L130:
	;
	v630 = v626 + int32(1)
	if v630&int32(3) == int32(0) {
		v637 = v630
		goto L125
	} else {
		goto L132
	}
L131:
	;
	v662 = v630
	goto L124
L132:
	;
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630))))
	if v635 != 0 {
		v626 = v630
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
	v650 = int32(-2139062144)
	if (int32(16843008)-v647|v647)&v650 == v650 {
		v641 = v641 + int32(4)
		goto L134
	} else {
		goto L136
	}
L135:
	;
	v656 = v641
	goto L137
L136:
	;
	goto L135
L137:
	;
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656))))
	if v660 != 0 {
		v656 = v656 + int32(1)
		goto L137
	} else {
		goto L139
	}
L138:
	;
	v662 = v656
	goto L124
L139:
	;
	goto L138
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v673)+8)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v673)+4)) = uint8(v602)
	v679 = v673 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v673))) = v679
	v682 = v670 + int32(1)
	if v682 != 0 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v685 = F_lappend(m, v398, v673)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L145
	}
L142:
	;
	v683 = F__emscripten_memcpy_bulkmem(m, v679, v600, v682)
	mBase = m.M
	goto L144
L143:
	;
	goto L144
L144:
	;
	goto L141
L145:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v609
	v924 = v685
	goto L117
L146:
	;
	v697 = F_open_auth_file(m, v693, v366, v377, v368+int32(20))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	if v697 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	F_pfree(m, v693)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L1
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	F_tokenize_auth_file(m, v693, v697, v368+int32(76), v366, v377)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L152
	}
L151:
	;
	v924 = v398
	goto L117
L152:
	;
	F_pfree(m, v693)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v368)+76))
	if v709 == int32(0) {
		v892 = v398
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v908 = F_FreeFile(m, v697)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L1
	} else {
		goto L177
	}
L155:
	;
	v712 = int32(0)
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v709)+4))
	if v713 <= v712 {
		v892 = v398
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v723 = v398
	v731 = v712
	goto L157
L157:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v709)+12))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v739+v731<<(uint(int32(2))%32))))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v743)+16))
	if v744 != 0 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v892 = v865
	goto L154
L159:
	;
	v745 = F_pstrdup(m, v744)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v743)))
	if v748 == int32(0) {
		v865 = v723
		goto L163
	} else {
		goto L164
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v368)+20)) = v745
	v892 = v723
	goto L154
L163:
	;
	v882 = v731 + int32(1)
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v709)+4))
	if v882 < v883 {
		v723 = v865
		v731 = v882
		goto L157
	} else {
		goto L176
	}
L164:
	;
	v751 = int32(0)
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v748)+4))
	if v752 <= v751 {
		v865 = v723
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v762 = v723
	v765 = v751
	goto L166
L166:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v748)+12))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v778+v765<<(uint(int32(2))%32))))
	if v782 == int32(0) {
		v838 = v762
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v865 = v838
	goto L163
L168:
	;
	v855 = v765 + int32(1)
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v748)+4))
	if v855 < v856 {
		v762 = v838
		v765 = v855
		goto L166
	} else {
		goto L175
	}
L169:
	;
	v785 = int32(0)
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v782)+4))
	if v786 <= v785 {
		v838 = v762
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v793 = v785
	v796 = v762
	goto L171
L171:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v782)+12))
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v812+v793<<(uint(int32(2))%32))))
	v817 = int32(4442576)
	v818 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v821 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v821
	v823 = F_lappend(m, v796, v816)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L1
	} else {
		goto L173
	}
L172:
	;
	v838 = v823
	goto L168
L173:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v818
	v828 = v793 + int32(1)
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v782)+4))
	if v828 < v829 {
		v793 = v828
		v796 = v823
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	goto L167
L176:
	;
	goto L158
L177:
	;
	if v377 != 0 {
		v924 = v892
		goto L117
	} else {
		goto L178
	}
L178:
	;
	v911 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	F_MemoryContextDelete(m, v911)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, _consts[444])) = int32(0)
	v924 = v892
	goto L117
L180:
	;
	v951 = v924
	goto L116
L181:
	;
	if v951 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v970 = int32(4442576)
	v971 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v974 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v974
	v976 = F_lappend(m, v376, v951)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L1
	} else {
		goto L185
	}
L183:
	;
	v981 = v376
	goto L184
L184:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v368)+20))
	v983 = int32(0)
	v984 = base.B2i32(v982 == v983)
	v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593))))
	if v985 == v983 {
		v990 = v363
		v991 = v364
		v992 = v365
		v993 = v366
		v995 = v368
		v996 = v984
		v1003 = v981
		v1004 = v377
		v1006 = v379
		v1007 = v380
		v1009 = v382
		v1012 = v385
		goto L74
	} else {
		goto L186
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v971
	v981 = v976
	goto L184
L186:
	;
	if v982 == int32(0) {
		v372 = v593
		v376 = v981
		goto L77
	} else {
		goto L187
	}
L187:
	;
	goto L78
L188:
	;
	v1342 = v1006 + v1007 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v995)+28)) = v1342
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v991)+76))
	if v1344 < int32(0) {
		goto L267
	} else {
		goto L268
	}
L189:
	;
	v1014 = v1003
	goto L191
L190:
	;
	v1014 = int32(1)
	goto L191
L191:
	;
	if v1014 == int32(0) {
		goto L188
	} else {
		goto L192
	}
L192:
	;
	if v1003 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v1287 = int32(4442576)
	v1288 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v1291 = *(*int32)(unsafe.Add(mBase, _consts[444]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v1291
	v1294 = F_palloc0(m, int32(20))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L1
	} else {
		goto L257
	}
L194:
	;
	if v996^int32(1) != 0 {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v1003)+4))
	if v1021 != int32(2) {
		goto L193
	} else {
		goto L196
	}
L196:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1003)+12))
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+4))
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v1025)+12))
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1026)))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1024)))
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+12))
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1029)))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1030)))
	v1032 = int32(391157)
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, _consts[445])))
	v1036 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1031))))
	if v1036 == int32(0) {
		v1055 = v1035
		v1056 = v1036
		goto L198
	} else {
		goto L199
	}
L197:
	;
	if v1056-v1055 == int32(0) {
		goto L205
	} else {
		goto L206
	}
L198:
	;
	goto L197
L199:
	;
	if v1035 != v1036 {
		v1055 = v1035
		v1056 = v1036
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v1040 = v1031
	v1041 = v1032
	goto L201
L201:
	;
	v1044 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1041)+1)))
	v1045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1040)+1)))
	if v1045 == int32(0) {
		v1055 = v1044
		v1056 = v1045
		goto L198
	} else {
		goto L203
	}
L202:
	;
	v1055 = v1044
	v1056 = v1045
	goto L198
L203:
	;
	v1048 = int32(1)
	if v1044 == v1045 {
		v1040 = v1040 + v1048
		v1041 = v1041 + v1048
		goto L201
	} else {
		goto L204
	}
L204:
	;
	goto L202
L205:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1027)))
	F_tokenize_include_file(m, v990, v1060, v992, v993, v1004, int32(0), v995+int32(20))
	mBase = m.M
	v1065 = m.ExcPending
	if v1065 != 0 {
		goto L1
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	v1067 = int32(202797)
	v1070 = int32(*(*uint8)(unsafe.Add(mBase, _consts[446])))
	v1071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1031))))
	if v1071 == int32(0) {
		v1090 = v1070
		v1091 = v1071
		goto L213
	} else {
		goto L214
	}
L208:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v995)+20))
	if v1066 != 0 {
		goto L193
	} else {
		goto L209
	}
L209:
	;
	goto L188
L210:
	;
	F_pfree(m, v1100)
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L1
	} else {
		goto L255
	}
L211:
	;
	v1205 = v1163
	goto L251
L212:
	;
	if v1091-v1090 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L213:
	;
	goto L212
L214:
	;
	if v1070 != v1071 {
		v1090 = v1070
		v1091 = v1071
		goto L213
	} else {
		goto L215
	}
L215:
	;
	v1075 = v1031
	v1076 = v1067
	goto L216
L216:
	;
	v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1076)+1)))
	v1080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075)+1)))
	if v1080 == int32(0) {
		v1090 = v1079
		v1091 = v1080
		goto L213
	} else {
		goto L218
	}
L217:
	;
	v1090 = v1079
	v1091 = v1080
	goto L213
L218:
	;
	v1083 = int32(1)
	if v1079 == v1080 {
		v1075 = v1075 + v1083
		v1076 = v1076 + v1083
		goto L216
	} else {
		goto L219
	}
L219:
	;
	goto L217
L220:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1027)))
	v1100 = F_GetConfFilesInDir(m, v1095, v990, v993, v995+int32(76), v995+int32(20))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L1
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v1166 = int32(107691)
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, _consts[447])))
	v1170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1031))))
	if v1170 == int32(0) {
		v1189 = v1169
		v1190 = v1170
		goto L241
	} else {
		goto L242
	}
L223:
	;
	if v1100 == int32(0) {
		goto L193
	} else {
		goto L224
	}
L224:
	;
	F_initStringInfo(m, v995+int32(60))
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v1108 = int32(0)
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v995)+76))
	if v1109 <= v1108 {
		goto L210
	} else {
		goto L226
	}
L226:
	;
	v1116 = v1108
	goto L227
L227:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1100+v1116<<(uint(int32(2))%32))))
	F_tokenize_include_file(m, v990, v1138, v992, v993, v1004, int32(0), v995+int32(20))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L1
	} else {
		goto L229
	}
L228:
	;
	v1163 = int32(0)
	if v1163 < v1161 {
		goto L211
	} else {
		goto L239
	}
L229:
	;
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v995)+20))
	if v1144 != 0 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v995)+64))
	if int32(0) < v1147 {
		goto L233
	} else {
		goto L234
	}
L231:
	;
	goto L232
L232:
	;
	v1160 = v1116 + int32(1)
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v995)+76))
	if v1160 < v1161 {
		v1116 = v1160
		goto L227
	} else {
		goto L238
	}
L233:
	;
	F_appendStringInfoChar(m, v995+int32(60), int32(10))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L1
	} else {
		goto L236
	}
L234:
	;
	v1156 = v1144
	goto L235
L235:
	;
	F_appendStringInfoString(m, v995+int32(60), v1156)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L1
	} else {
		goto L237
	}
L236:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v995)+20))
	v1156 = v1155
	goto L235
L237:
	;
	goto L232
L238:
	;
	goto L228
L239:
	;
	goto L210
L240:
	;
	if v1190-v1189 != 0 {
		goto L193
	} else {
		goto L248
	}
L241:
	;
	goto L240
L242:
	;
	if v1169 != v1170 {
		v1189 = v1169
		v1190 = v1170
		goto L241
	} else {
		goto L243
	}
L243:
	;
	v1174 = v1031
	v1175 = v1166
	goto L244
L244:
	;
	v1178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1175)+1)))
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1174)+1)))
	if v1179 == int32(0) {
		v1189 = v1178
		v1190 = v1179
		goto L241
	} else {
		goto L246
	}
L245:
	;
	v1189 = v1178
	v1190 = v1179
	goto L241
L246:
	;
	v1182 = int32(1)
	if v1178 == v1179 {
		v1174 = v1174 + v1182
		v1175 = v1175 + v1182
		goto L244
	} else {
		goto L247
	}
L247:
	;
	goto L245
L248:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1027)))
	F_tokenize_include_file(m, v990, v1192, v992, v993, v1004, int32(1), v995+int32(20))
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v995)+20))
	if v1198 == int32(0) {
		goto L188
	} else {
		goto L250
	}
L250:
	;
	goto L193
L251:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1100+v1205<<(uint(int32(2))%32))))
	F_pfree(m, v1227)
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L1
	} else {
		goto L253
	}
L252:
	;
	goto L210
L253:
	;
	v1231 = v1205 + int32(1)
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v995)+76))
	if v1231 < v1232 {
		v1205 = v1231
		goto L251
	} else {
		goto L254
	}
L254:
	;
	goto L252
L255:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v995)+64))
	if v1259 == int32(0) {
		goto L188
	} else {
		goto L256
	}
L256:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v995)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v995)+20)) = v1262
	goto L193
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1294))) = v1003
	v1297 = F_pstrdup(m, v990)
	mBase = m.M
	v1298 = m.ExcPending
	if v1298 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1294)+8)) = v1006
	*(*int32)(unsafe.Add(mBase, uint32(v1294)+4)) = v1297
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v995)+44))
	v1302 = F_pstrdup(m, v1301)
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1294)+12)) = v1302
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v995)+20))
	if v1305 != 0 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1306 = F_pstrdup(m, v1305)
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L1
	} else {
		goto L263
	}
L261:
	;
	v1309 = int32(0)
	goto L262
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1294)+16)) = v1309
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v992)))
	v1312 = F_lappend(m, v1311, v1294)
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L1
	} else {
		goto L264
	}
L263:
	;
	v1309 = v1306
	goto L262
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v992))) = v1312
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v1288
	goto L188
L265:
	;
	if int32(base.Ui32(v1349)>>(uint(int32(4))%32))&int32(1) == int32(0) {
		v76 = v990
		v77 = v991
		v78 = v992
		v79 = v993
		v81 = v995
		v90 = v1004
		v92 = v1342
		v95 = v1009
		v98 = v1012
		goto L14
	} else {
		goto L270
	}
L266:
	;
	goto L265
L267:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v991)))
	v1349 = v1347
	goto L266
L268:
	;
	goto L269
L269:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v991)))
	v1349 = v1348
	goto L266
L270:
	;
	goto L15
L271:
	;
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1361)+32))
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v1384
	m.G0 = v1361 + int32(80)
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
	F_errmsg(m, int32(179332), v15)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L11
	} else {
		goto L63
	}
L63:
	;
	F_errdetail(m, int32(543186), int32(0))
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
	F_errfinish(m, int32(474878), int32(221), int32(290143))
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
												F_errmsg(m, int32(181876), v15+int32(16))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(576960), int32(0))
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
															F_errfinish(m, int32(470280), int32(596), int32(197144))
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
													F_errmsg(m, int32(181876), v15+int32(16))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														F_errhint(m, int32(576960), int32(0))
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
																F_errfinish(m, int32(470280), int32(596), int32(197144))
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
													F_errmsg(m, int32(181876), v15+int32(16))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int32(0)
													} else {
														F_errhint(m, int32(576960), int32(0))
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
																F_errfinish(m, int32(470280), int32(596), int32(197144))
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
									F_errmsg(m, int32(496538), int32(0))
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
											F_errfinish(m, int32(470280), int32(512), int32(197144))
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
									F_errmsg(m, int32(496577), int32(0))
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
											F_errfinish(m, int32(470280), int32(517), int32(197144))
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
				F_errmsg(m, int32(669354), v15)
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
						F_errfinish(m, int32(470280), int32(485), int32(197144))
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
														F_errmsg(m, int32(172784), v19)
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
																F_errfinish(m, int32(470280), int32(1004), int32(110340))
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
													F_errmsg(m, int32(172784), v19)
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
															F_errfinish(m, int32(470280), int32(1004), int32(110340))
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
	v206 = int32(497632)
	goto L52
L51:
	;
	v206 = int32(493984)
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
	v209 = int32(506137)
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
	v371 = int32(4435480)
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
	F_errmsg(m, int32(493864), int32(0))
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
	F_errfinish(m, int32(474990), int32(2066), int32(390912))
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
	F_errmsg(m, int32(493818), v20+int32(48))
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
	v491 = int32(356532)
	goto L132
L132:
	;
	goto L129
L133:
	;
	F_errfinish(m, int32(474990), int32(2076), int32(390912))
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
	F_errmsg(m, int32(138622), v20+int32(32))
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
	F_errfinish(m, int32(474990), int32(2226), int32(390912))
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
	v611 = F_pg_snprintf(m, v20+int32(96), int32(32), int32(465655), v20)
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
	F_errmsg(m, int32(291190), int32(0))
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
	F_errfinish(m, int32(474990), int32(2138), int32(390912))
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
	F_errmsg(m, int32(424495), int32(0))
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
	F_errfinish(m, int32(475120), int32(883), int32(77005))
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
	F_errmsg(m, int32(114839), v14+int32(16))
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
	F_errfinish(m, int32(475120), int32(936), int32(258182))
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
	F_errmsg(m, int32(273741), v14)
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
	F_errfinish(m, int32(475120), int32(944), int32(258182))
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
	F_errmsg(m, int32(424250), int32(0))
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
	F_errfinish(m, int32(475120), int32(979), int32(258182))
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
	F_errmsg(m, int32(424345), int32(0))
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
	F_errfinish(m, int32(475120), int32(985), int32(258182))
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
	F_errmsg(m, int32(424299), int32(0))
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
	F_errfinish(m, int32(475120), int32(991), int32(258182))
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
	F_errmsg(m, int32(424392), int32(0))
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
	F_errfinish(m, int32(475120), int32(997), int32(258182))
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
	F_errmsg_internal(m, int32(397112), int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L38
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(475120), int32(1019), int32(258182))
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
	F_errmsg(m, int32(395832), v14+int32(32))
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
	F_errfinish(m, int32(475120), int32(139), int32(342482))
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
func F_trueTriConsistentFn(m *base.Module, l0 int32) int32 {
	return int32(1)
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
	F_pg_qsort(m, v16, v76, int32(4), int32(1528))
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
	F_pg_qsort(m, v86, v146, int32(4), int32(1528))
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
			F_infix(m, v8+v34, int32(-1), v31)
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
	v78 = int32(4442576)
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
	F_errmsg(m, int32(632344), v20)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errsave_finish(m, v25, int32(471930), int32(215), int32(261755))
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
	F_errmsg(m, int32(632285), v18+int32(-48))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errsave_finish(m, v25, int32(471930), int32(221), int32(261755))
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
	F_qsort_arg(m, v231, v227, int32(12), int32(1534), v233)
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
	F_pg_qsort(m, v346, v343, int32(2), int32(1535))
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
	F_pg_qsort(m, v551, v548, int32(2), int32(1535))
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
	F_errmsg(m, int32(632477), v18+int32(-32))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errsave_finish(m, v25, int32(471930), int32(274), int32(261755))
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
	F_errmsg_internal(m, int32(311220), int32(0))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(471930), int32(292), int32(261755))
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
			v25 = F_pg_sprintf(m, v8+int32(16), int32(465664), v8)
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
			v14 = int32(4442576)
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
								F_errmsg_internal(m, int32(349675), int32(0))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return
								} else {
									F_errfinish(m, int32(472121), int32(69), int32(289779))
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
				F_errmsg_internal(m, int32(41185), v9)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					F_errfinish(m, int32(472121), int32(57), int32(289779))
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
	F_errmsg_internal(m, int32(381393), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L23
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(310305), int32(327), int32(324221))
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
	F_errmsg_internal(m, int32(381393), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L23
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(310305), int32(327), int32(324221))
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
	v320 = int32(4442576)
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
	F_errmsg_internal(m, int32(441036), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L23
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(310305), int32(630), int32(296246))
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
							F_errmsg(m, int32(289374), v8+int32(16))
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
									F_errfinish(m, int32(475439), int32(280), int32(353818))
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
						F_errmsg(m, int32(68911), v8)
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
								F_errfinish(m, int32(475439), int32(274), int32(353818))
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
