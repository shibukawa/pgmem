package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_systable_beginscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	v7 = int32(0)
	if l2 == v7 {
		v33 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v35 = F_palloc(m, int32(24))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L11
	}
L2:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_systable_beginscan[0])))
	if v17&int32(1) != 0 {
		v33 = v7
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_systable_beginscan[1]))
	if v21 != l1 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v27 != 0 {
		v33 = v7
		goto L1
	} else {
		goto L8
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_systable_beginscan[2]))
	v25 = F_list_member_ptr(m, v24, l1)
	mBase = m.M
	v27 = v25
	goto L7
L6:
	;
	v27 = int32(1)
	goto L7
L7:
	;
	goto L4
L8:
	;
	v29 = F_index_open(m, l1, int32(1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v33 = v29
	goto L1
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = l0
	v40 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v40
	if l3 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v46 = F_GetCatalogSnapshot(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	v50 = l3
	v51 = v7
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v51
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_systable_beginscan[3]))
	if v54 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v48 = F_RegisterSnapshot(m, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v50 = v48
	v51 = v48
	goto L15
L18:
	;
	v56 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_systable_beginscan[4])) = uint8(v56)
	goto L20
L19:
	;
	goto L20
L20:
	;
	if v33 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v59 = F_palloc_mul(m, int32(56), l4)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L9
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v54 != 0 {
		goto L49
	} else {
		goto L50
	}
L24:
	;
	if l4 <= int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v185 = int32(0)
	v188 = F_index_beginscan(m, l0, v33, v50, v185, l4, v185, v185)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L9
	} else {
		goto L45
	}
L26:
	;
	v74 = v7
	goto L27
L27:
	;
	v77 = v74 * int32(56)
	v78 = v59 + v77
	v79 = v77 + l5
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v78)+48)) = v80
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v79)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v78)+40)) = v82
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v79)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v78)+32)) = v84
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v79)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v78)+24)) = v86
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v79)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v78)+16)) = v88
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v79)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v78)+8)) = v90
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v33)+192))
	v95 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94)+8)))
	if v95 <= int32(0) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L9
	} else {
		goto L42
	}
L29:
	;
	goto L28
L30:
	;
	if v131 == v135 {
		goto L29
	} else {
		goto L40
	}
L31:
	;
	v131 = v95
	v135 = int32(0)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+4)))
	v109 = int32(0)
	goto L34
L34:
	;
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94+int32(48)+v109<<(uint(int32(1))%32)))))
	if v119 == v102 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L29
L36:
	;
	v122 = v109 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v78)+4)) = uint16(v122)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v33)+192))
	v125 = int32(*(*int16)(unsafe.Add(mBase, uint32(v124)+8)))
	v131 = v125
	v135 = v109
	goto L30
L37:
	;
	goto L38
L38:
	;
	v127 = v109 + int32(1)
	if v127 != v95 {
		v109 = v127
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	v144 = v74 + int32(1)
	if l4 != v144 {
		v74 = v144
		goto L27
	} else {
		goto L41
	}
L41:
	;
	goto L25
L42:
	;
	F_errmsg_internal(m, int32(_a_F_systable_beginscan_0), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_systable_beginscan_1), int32(454), int32(_a_F_systable_beginscan_2))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v188
	v191 = int32(0)
	F_index_rescan(m, v188, v59, l4, v191, v191)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = int32(0)
	F_pfree(m, v59)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	return v35
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L9
	} else {
		goto L54
	}
L49:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_systable_beginscan[4])))
	if v201&int32(1) == int32(0) {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+8))
	v210 = m.T0[v209].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v50, l4, l5, int32(0), int32(321))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L9
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v210
	return v35
L54:
	;
	F_errmsg_internal(m, int32(_a_F_systable_beginscan_3), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_systable_beginscan_4), int32(931), int32(_a_F_systable_beginscan_5))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L9
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_systable_endscan_ordered(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v3 != 0 {
		F_ExecDropSingleTupleTableSlot(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_index_endscan(m, v8)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v11 != 0 {
					F_UnregisterSnapshot(m, v11)
					mBase = m.M
					v13 = m.ExcPending
					if v13 != 0 {
						return
					} else {
						v15 = *(*int32)(unsafe.Add(mBase, _c_F_systable_endscan_ordered[0]))
						if v15 != 0 {
							v17 = int32(0)
							*(*uint8)(unsafe.Add(mBase, _c_F_systable_endscan_ordered[1])) = uint8(v17)
						} else {
						}
						F_pfree(m, l0)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, _c_F_systable_endscan_ordered[0]))
					if v15 != 0 {
						v17 = int32(0)
						*(*uint8)(unsafe.Add(mBase, _c_F_systable_endscan_ordered[1])) = uint8(v17)
					} else {
					}
					F_pfree(m, l0)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		F_index_endscan(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v11 != 0 {
				F_UnregisterSnapshot(m, v11)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, _c_F_systable_endscan_ordered[0]))
					if v15 != 0 {
						v17 = int32(0)
						*(*uint8)(unsafe.Add(mBase, _c_F_systable_endscan_ordered[1])) = uint8(v17)
					} else {
					}
					F_pfree(m, l0)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, _c_F_systable_endscan_ordered[0]))
				if v15 != 0 {
					v17 = int32(0)
					*(*uint8)(unsafe.Add(mBase, _c_F_systable_endscan_ordered[1])) = uint8(v17)
				} else {
				}
				F_pfree(m, l0)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_systable_inplace_update_begin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_begin[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L14
	} else {
		goto L105
	}
L2:
	;
	if v22&int32(1) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v22 = int32(1)
	goto L5
L4:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+76)))
	v22 = v21
	goto L5
L5:
	;
	goto L2
L6:
	;
	v39 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L14
	} else {
		goto L101
	}
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_begin[1]))
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v319 = F_heap_copytuple(m, v53)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L14
	} else {
		goto L100
	}
L11:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if v39 == int32(_a_F_systable_inplace_update_begin_0) {
		goto L1
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
	v48 = int32(1)
	v51 = F_systable_beginscan(m, l0, l1, v48, int32(0), v48, l2)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v53 = F_systable_getnext(m, v51)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	if v53 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_systable_endscan(m, v51)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L14
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+44))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+72))
	v67 = m.G0
	v69 = v67 - int32(32)
	m.G0 = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = v71
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v69)+16)) = v73
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	*(*int64)(unsafe.Add(mBase, uint32(v69)+8)) = v75
	F_CacheInvalidateHeapTupleCommon(m, v63, v65, int32(0), int32(1797))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L14
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	return
L23:
	;
	v82 = v69 + int32(8)
	v84 = v82 | int32(4)
	F_LockTuple(m, v63, v84, int32(7))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	F_LockBufferInternal(m, v66, int32(3))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	v91 = int32(1)
	v93 = F_GetCurrentCommandId(m, int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L14
	} else {
		goto L32
	}
L26:
	;
	m.G0 = v69 + int32(32)
	if v313 == int32(0) {
		v39 = v39 + int32(1)
		goto L9
	} else {
		goto L99
	}
L27:
	;
	F_UnlockTuple(m, v63, v84, int32(7))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L14
	} else {
		goto L96
	}
L28:
	;
	F_UnlockBuffer(m, v66)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L14
	} else {
		goto L94
	}
L29:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v69)+24))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129)+20)))
	if v131&int32(_a_F_systable_inplace_update_begin_1) != 0 {
		goto L42
	} else {
		goto L43
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L14
	} else {
		goto L38
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L14
	} else {
		goto L34
	}
L32:
	;
	v95 = F_HeapTupleSatisfiesUpdate(m, v82, v93, v66)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L14
	} else {
		goto L33
	}
L33:
	;
	switch v95 {
	case 0:
		v313 = v91
		goto L26
	case 1:
		goto L31
	case 2:
		goto L30
	default:
		goto L28
	case 5:
		goto L29
	}
L34:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L14
	} else {
		goto L35
	}
L35:
	;
	F_errmsg_internal(m, int32(_a_F_systable_inplace_update_begin_2), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L14
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_systable_inplace_update_begin_3), int32(_a_F_systable_inplace_update_begin_4), int32(_a_F_systable_inplace_update_begin_5))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	F_errmsg(m, int32(_a_F_systable_inplace_update_begin_6), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_systable_inplace_update_begin_3), int32(_a_F_systable_inplace_update_begin_7), int32(_a_F_systable_inplace_update_begin_5))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v136 = F_DoesMultiXactIdConflict(m, v130, v131, int32(2), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L14
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if base.Ui32(v130) < base.Ui32(int32(3)) {
		goto L51
	} else {
		goto L52
	}
L45:
	;
	if v136 == int32(0) {
		v313 = v91
		goto L26
	} else {
		goto L46
	}
L46:
	;
	F_UnlockBuffer(m, v66)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L14
	} else {
		goto L47
	}
L47:
	;
	F_systable_endscan(m, v51)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L14
	} else {
		goto L48
	}
L48:
	;
	v144 = int32(4)
	v145 = int32(0)
	v150 = F_Do_MultiXactIdWait(m, v130, v144, v131, v145, v63, v84, int32(1), v69+v144, v145)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L14
	} else {
		goto L49
	}
L49:
	;
	goto L27
L50:
	;
	if v283|base.B2i32(v131&int32(80) == int32(16)) != 0 {
		v313 = v91
		goto L26
	} else {
		goto L90
	}
L51:
	;
	v283 = int32(0)
	goto L50
L52:
	;
	goto L53
L53:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_begin[2]))
	if v163 == v130 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v283 = int32(1)
	goto L50
L55:
	;
	goto L56
L56:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_begin[3]))
	if v167 <= int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v283 = v273
	goto L50
L58:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_begin[0]))
	if v171 == int32(0) {
		v273 = int32(0)
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_begin[4]))
	v243 = int32(0)
	v246 = v167 - int32(1)
	goto L80
L61:
	;
	v176 = v171
	goto L62
L62:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v176)+20))
	if v182 == int32(4) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v273 = int32(0)
	goto L57
L64:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v176)+80))
	if v236 != 0 {
		v176 = v236
		goto L62
	} else {
		goto L79
	}
L65:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	if v185 == int32(0) {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v188 = int32(1)
	if v130 == v185 {
		v273 = v188
		goto L57
	} else {
		goto L67
	}
L67:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v176)+52))
	v192 = v190 - int32(1)
	if v192 < int32(0) {
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v176)+48))
	v198 = int32(0)
	v201 = v192
	goto L69
L69:
	;
	v206 = int32(2)
	v207 = base.I32_div_s(v201-v198, v206)
	v208 = v207 + v198
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v195+v208<<(uint(v206)%32))))
	if v212 == v130 {
		v273 = v188
		goto L57
	} else {
		goto L71
	}
L70:
	;
	goto L64
L71:
	;
	v221 = base.B2i32(v212-v130 < int32(0)) | base.B2i32(base.Ui32(v212) < base.Ui32(int32(3)))
	if v221 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v222 = v208 + int32(1)
	goto L74
L73:
	;
	v222 = v198
	goto L74
L74:
	;
	if v221 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v225 = v201
	goto L77
L76:
	;
	v225 = v208 - int32(1)
	goto L77
L77:
	;
	if v222 <= v225 {
		v198 = v222
		v201 = v225
		goto L69
	} else {
		goto L78
	}
L78:
	;
	goto L70
L79:
	;
	goto L63
L80:
	;
	v251 = int32(2)
	v252 = base.I32_div_s(v246-v243, v251)
	v253 = v252 + v243
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v241+v253<<(uint(v251)%32))))
	v258 = base.B2i32(v257 == v130)
	if v257 == v130 {
		v273 = v258
		goto L57
	} else {
		goto L82
	}
L81:
	;
	v273 = v258
	goto L57
L82:
	;
	v261 = base.B2i32(base.Ui32(v257) < base.Ui32(v130))
	if base.Ui32(v257) < base.Ui32(v130) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v262 = v253 + int32(1)
	goto L85
L84:
	;
	v262 = v243
	goto L85
L85:
	;
	if base.Ui32(v257) < base.Ui32(v130) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v265 = v246
	goto L88
L87:
	;
	v265 = v253 - int32(1)
	goto L88
L88:
	;
	if v262 <= v265 {
		v243 = v262
		v246 = v265
		goto L80
	} else {
		goto L89
	}
L89:
	;
	goto L81
L90:
	;
	F_UnlockBuffer(m, v66)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L14
	} else {
		goto L91
	}
L91:
	;
	F_systable_endscan(m, v51)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L14
	} else {
		goto L92
	}
L92:
	;
	F_XactLockTableWait(m, v130, v63, v84, int32(1))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L14
	} else {
		goto L93
	}
L93:
	;
	goto L27
L94:
	;
	F_systable_endscan(m, v51)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L14
	} else {
		goto L95
	}
L95:
	;
	goto L27
L96:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_begin[5])) = int32(0)
	goto L97
L97:
	;
	F_InvalidateCatalogSnapshot(m)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L14
	} else {
		goto L98
	}
L98:
	;
	v313 = int32(0)
	goto L26
L99:
	;
	goto L10
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v319
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v51
	return
L101:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L14
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(_a_F_systable_inplace_update_begin_8), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L14
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_systable_inplace_update_begin_9), int32(832), int32(_a_F_systable_inplace_update_begin_10))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L14
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
	F_errmsg_internal(m, int32(_a_F_systable_inplace_update_begin_11), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L14
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_systable_inplace_update_begin_9), int32(855), int32(_a_F_systable_inplace_update_begin_10))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L14
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_systable_inplace_update_finish(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int64
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int64
	_ = v266
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v287 int64
	_ = v287
	var v288 int32
	_ = v288
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	v3 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	v20 = m.G0
	v22 = v20 - int32(_a_F_systable_inplace_update_finish_0)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[0]))) = v3
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[1]))) = uint8(v3)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
	if v29 != v31 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_systable_endscan(m, l0)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L12
	} else {
		goto L101
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L12
	} else {
		goto L98
	}
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v34 = v33 - v29
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v34 != v35-v31 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[2]))
	if int32(0) < v39 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[3]))
	if v47 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v110 = v3
	goto L7
L7:
	;
	v115 = v31 + v30
	v116 = v24 + v29
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[3]))
	if v118 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L8:
	;
	v110 = v109
	goto L7
L9:
	;
	v50 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[1]))) = uint8(v50)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[0]))) = v50
	v109 = v50
	goto L8
L10:
	;
	goto L11
L11:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[1]))) = uint8(v55)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v66 = F_palloc(m, (v57+v58-(v60+v61))<<(uint(int32(4))%32))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[0]))) = v66
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[3]))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v73 = v71 - v72
	if int32(0) < v73 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v77 = v73 << (uint(int32(4)) % 32)
	if v77 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v84 = v3
	goto L16
L16:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v87 = v85 - v86
	if int32(0) < v87 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[5]))
	base.MemoryCopy(m, v66, v79+v72<<(uint(int32(4))%32), v77)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v84 = v73
	goto L16
L20:
	;
	v91 = v87 << (uint(int32(4)) % 32)
	if v91 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v103 = v84
	goto L22
L22:
	;
	v109 = v103
	goto L8
L23:
	;
	v92 = int32(4)
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[6]))
	base.MemoryCopy(m, v66+v84<<(uint(v92)%32), v96+v86<<(uint(v92)%32), v91)
	goto L25
L24:
	;
	goto L25
L25:
	;
	v103 = v84 + v87
	goto L22
L26:
	;
	v126 = int32(_a_F_systable_inplace_update_finish_5)
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[4])) = v128 + int32(1)
	F_MarkBufferDirty(m, v19)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L12
	} else {
		goto L30
	}
L27:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+16)))
	if v121 != int32(1) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	F_RelationCacheInitFilePreInvalidate(m)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v16)+48))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+118)))
	if v135 != int32(112) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if v34 != 0 {
		goto L73
	} else {
		goto L74
	}
L32:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[2]))
	if v139 <= int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	if v142 != 0 {
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v19 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	if v143 != 0 {
		goto L31
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161)+14)))
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161)+12)))
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[10]))) = uint16(v164)
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[12]))) = v167
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[14]))) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[15]))) = v110
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[1]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[16]))) = uint8(v173)
	F_XLogBeginInsert(m)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L12
	} else {
		goto L42
	}
L39:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[8]))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v147+(v19^int32(-1))<<(uint(int32(2))%32))))
	v161 = v153
	goto L38
L40:
	;
	goto L41
L41:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[9]))
	v161 = v155 + v19<<(uint(int32(13))%32) + int32(-8192)
	goto L38
L42:
	;
	F_XLogRegisterData(m, v22+int32(_a_F_systable_inplace_update_finish_6), int32(20))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L12
	} else {
		goto L43
	}
L43:
	;
	if v110 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[0])))
	F_XLogRegisterData(m, v182, v110<<(uint(int32(4))%32))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L12
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if v163 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	base.MemoryCopy(m, v22+int32(32), v161, v163)
	goto L50
L49:
	;
	goto L50
L50:
	;
	v191 = int32(_a_F_systable_inplace_update_finish_7) - v162
	if v191 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	base.MemoryCopy(m, v22+int32(32)+v162, v161+v162, v191)
	goto L53
L52:
	;
	goto L53
L53:
	;
	if v34 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	base.MemoryCopy(m, v22+int32(32)+(v116-v161), v115, v34)
	goto L56
L55:
	;
	goto L56
L56:
	;
	v203 = v22 + int32(20)
	if v19 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[19]))
	if v237 <= int32(0) {
		goto L62
	} else {
		goto L63
	}
L58:
	;
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v225)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v225)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+8)) = v227
	*(*int64)(unsafe.Add(mBase, uint32(v203))) = v226
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(16)))) = v230
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v225)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(12)))) = v232
	goto L57
L59:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[17]))
	v225 = v212 + (v19^int32(-1))*int32(56)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[18]))
	v220 = int32(56)
	v225 = v219 + v19*v220 - v220
	goto L58
L62:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[19])) = int32(1)
	goto L64
L63:
	;
	goto L64
L64:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[20]))
	if v244 <= int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L12
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[21]))
	v263 = v22 + int32(20)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v261)+12)) = v264
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v263)))
	*(*int64)(unsafe.Add(mBase, uint32(v261)+4)) = v266
	v268 = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v261)+24)) = v22 + v268
	*(*int32)(unsafe.Add(mBase, uint32(v261)+20)) = v235
	*(*int32)(unsafe.Add(mBase, uint32(v261)+16)) = v234
	v273 = int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)) = uint8(v273)
	v275 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v261)+28)) = v275
	v277 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v261))) = uint8(v277)
	*(*int32)(unsafe.Add(mBase, uint32(v261)+36)) = v261 + v268
	F_XLogRegisterBufData(m, v275, v115, v34)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L12
	} else {
		goto L71
	}
L68:
	;
	F_errmsg_internal(m, int32(_a_F_systable_inplace_update_finish_8), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_systable_inplace_update_finish_9), int32(328), int32(_a_F_systable_inplace_update_finish_10))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L12
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	v287 = F_XLogInsert(m, int32(10), int32(112))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L12
	} else {
		goto L72
	}
L72:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v161))) = base.I64_rotl(v287, int64(32))
	goto L31
L73:
	;
	base.MemoryCopy(m, v116, v115, v34)
	goto L75
L74:
	;
	goto L75
L75:
	;
	F_UnlockBuffer(m, v19)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[3]))
	if v301 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)+8))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	v304 = v302 - v303
	if int32(0) < v304 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	v339 = int32(_a_F_systable_inplace_update_finish_5)
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[4])) = v341 - int32(1)
	F_UnlockTuple(m, v16, l1+int32(4), int32(7))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L12
	} else {
		goto L92
	}
L80:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[5]))
	F_SIInsertDataEntries(m, v308+v303<<(uint(int32(4))%32), v304)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L12
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v301)+12))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	v316 = v314 - v315
	if int32(0) < v316 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	goto L82
L84:
	;
	v320 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[6]))
	F_SIInsertDataEntries(m, v320+v315<<(uint(int32(4))%32), v316)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L12
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[3]))
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327)+16)))
	if v328 == int32(1) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L86
L88:
	;
	F_RelationCacheInitFilePostInvalidate(m)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L12
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[3])) = int32(0)
	goto L79
L91:
	;
	goto L90
L92:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L12
	} else {
		goto L93
	}
L93:
	;
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[7]))
	if v353 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	F_CacheInvalidateHeapTuple(m, v16, l1, int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L12
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	m.G0 = v22 + int32(_a_F_systable_inplace_update_finish_0)
	goto L1
L97:
	;
	goto L96
L98:
	;
	F_errmsg_internal(m, int32(_a_F_systable_inplace_update_finish_1), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L12
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_systable_inplace_update_finish_2), int32(_a_F_systable_inplace_update_finish_3), int32(_a_F_systable_inplace_update_finish_4))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L12
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	return
}
func F_systable_recheck_tuple(m *base.Module, l0 int32) int32 {
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
	var v10 int32
	_ = v10
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+56))
	v6 = F_GetCatalogSnapshot(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_RegisterSnapshot(m, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+188))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
			v16 = m.T0[v15].(func(*base.Module, int32, int32, int32) int32)(m, v12, v13, v10)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_UnregisterSnapshot(m, v10)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_HandleConcurrentAbort(m)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						return v16
					}
				}
			}
		}
	}
}
