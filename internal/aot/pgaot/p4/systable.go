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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v138 int32
	_ = v138
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
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
	if v33 != 0 {
		goto L19
	} else {
		goto L20
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
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_systable_beginscan[3]))
	if v215 != 0 {
		goto L47
	} else {
		goto L48
	}
L19:
	;
	v55 = F_palloc(m, l4*int32(48))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L9
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+8))
	v196 = m.T0[v195].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v50, l4, l5, int32(0), int32(321))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L9
	} else {
		goto L46
	}
L22:
	;
	if l4 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v179 = int32(0)
	v181 = F_index_beginscan(m, l0, v33, v50, v179, l4, v179)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L9
	} else {
		goto L43
	}
L24:
	;
	v70 = v7
	goto L25
L25:
	;
	v73 = v70 * int32(48)
	v74 = v55 + v73
	v75 = v73 + l5
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v75)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v74)+40)) = v76
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v75)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v74)+32)) = v78
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v75)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v74)+24)) = v80
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v75)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v74)+16)) = v82
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v75)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v74)+8)) = v84
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
	*(*int64)(unsafe.Add(mBase, uint32(v74))) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v33)+192))
	v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88)+8)))
	if v89 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L9
	} else {
		goto L40
	}
L27:
	;
	goto L26
L28:
	;
	if v125 == v129 {
		goto L27
	} else {
		goto L38
	}
L29:
	;
	v125 = v89
	v129 = int32(0)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+4)))
	v103 = int32(0)
	goto L32
L32:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v88+int32(48)+v103<<(uint(int32(1))%32)))))
	if v113 == v96 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L27
L34:
	;
	v116 = v103 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)) = uint16(v116)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v33)+192))
	v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(v118)+8)))
	v125 = v119
	v129 = v103
	goto L28
L35:
	;
	goto L36
L36:
	;
	v121 = v103 + int32(1)
	if v121 != v89 {
		v103 = v121
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v138 = v70 + int32(1)
	if l4 != v138 {
		v70 = v138
		goto L25
	} else {
		goto L39
	}
L39:
	;
	goto L23
L40:
	;
	F_errmsg_internal(m, int32(_a_F_systable_beginscan_0), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_systable_beginscan_1), int32(446), int32(_a_F_systable_beginscan_2))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L9
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v181
	v184 = int32(0)
	F_index_rescan(m, v181, v55, l4, v184, v184)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = int32(0)
	F_pfree(m, v55)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L9
	} else {
		goto L45
	}
L45:
	;
	goto L18
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v196
	goto L18
L47:
	;
	v217 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_systable_beginscan[4])) = uint8(v217)
	goto L49
L48:
	;
	goto L49
L49:
	;
	return v35
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
	var v142 int32
	_ = v142
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
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
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
	v333 = m.ExcPending
	if v333 != 0 {
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
	v317 = m.ExcPending
	if v317 != 0 {
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
	v310 = F_heap_copytuple(m, v53)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
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
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+40))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+68))
	v67 = m.G0
	v69 = v67 - int32(32)
	m.G0 = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = v71
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v69)+16)) = v73
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	*(*int64)(unsafe.Add(mBase, uint32(v69)+8)) = v75
	F_CacheInvalidateHeapTupleCommon(m, v63, v65, int32(0), int32(1584))
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
	F_LockBuffer(m, v66, int32(2))
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
	if v304 == int32(0) {
		v39 = v39 + int32(1)
		goto L9
	} else {
		goto L99
	}
L27:
	;
	F_UnlockTuple(m, v63, v84, int32(7))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L14
	} else {
		goto L96
	}
L28:
	;
	F_LockBuffer(m, v66, int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
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
		v304 = v91
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
		v304 = v91
		goto L26
	} else {
		goto L46
	}
L46:
	;
	F_LockBuffer(m, v66, int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L14
	} else {
		goto L47
	}
L47:
	;
	F_systable_endscan(m, v51)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L14
	} else {
		goto L48
	}
L48:
	;
	v145 = int32(4)
	v146 = int32(0)
	v151 = F_Do_MultiXactIdWait(m, v130, v145, v131, v146, v63, v84, int32(1), v69+v145, v146)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L14
	} else {
		goto L49
	}
L49:
	;
	goto L27
L50:
	;
	if v272|base.B2i32(v131&int32(80) == int32(16)) != 0 {
		v304 = v91
		goto L26
	} else {
		goto L90
	}
L51:
	;
	v272 = int32(0)
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
	v272 = int32(1)
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
	v272 = v264
	goto L50
L58:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_begin[0]))
	if v171 == int32(0) {
		v264 = int32(0)
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_begin[4]))
	v235 = int32(0)
	v237 = v167 - int32(1)
	goto L80
L61:
	;
	v176 = v171
	goto L62
L62:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v176)+20))
	if v181 == int32(4) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v264 = int32(0)
	goto L57
L64:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v176)+80))
	if v228 != 0 {
		v176 = v228
		goto L62
	} else {
		goto L79
	}
L65:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	if v184 == int32(0) {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v187 = int32(1)
	if v130 == v184 {
		v264 = v187
		goto L57
	} else {
		goto L67
	}
L67:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v176)+52))
	v191 = v189 - int32(1)
	if v191 < int32(0) {
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v196 = int32(0)
	v198 = v191
	goto L69
L69:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v176)+48))
	v204 = int32(2)
	v205 = base.I32_div_s(v198-v196, v204)
	v206 = v205 + v196
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v202+v206<<(uint(v204)%32))))
	if v210 == v130 {
		v264 = v187
		goto L57
	} else {
		goto L71
	}
L70:
	;
	goto L64
L71:
	;
	v214 = F_TransactionIdPrecedes(m, v210, v130)
	mBase = m.M
	if v214 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v215 = v206 + int32(1)
	goto L74
L73:
	;
	v215 = v196
	goto L74
L74:
	;
	if v214 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v218 = v198
	goto L77
L76:
	;
	v218 = v206 - int32(1)
	goto L77
L77:
	;
	if v215 <= v218 {
		v196 = v215
		v198 = v218
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
	v242 = int32(2)
	v243 = base.I32_div_s(v237-v235, v242)
	v244 = v243 + v235
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(v242)%32))))
	v249 = base.B2i32(v248 == v130)
	if v248 == v130 {
		v264 = v249
		goto L57
	} else {
		goto L82
	}
L81:
	;
	v264 = v249
	goto L57
L82:
	;
	v252 = base.B2i32(base.Ui32(v248) < base.Ui32(v130))
	if base.Ui32(v248) < base.Ui32(v130) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v253 = v244 + int32(1)
	goto L85
L84:
	;
	v253 = v235
	goto L85
L85:
	;
	if base.Ui32(v248) < base.Ui32(v130) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v256 = v237
	goto L88
L87:
	;
	v256 = v244 - int32(1)
	goto L88
L88:
	;
	if v253 <= v256 {
		v235 = v253
		v237 = v256
		goto L80
	} else {
		goto L89
	}
L89:
	;
	goto L81
L90:
	;
	F_LockBuffer(m, v66, int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L14
	} else {
		goto L91
	}
L91:
	;
	F_systable_endscan(m, v51)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L14
	} else {
		goto L92
	}
L92:
	;
	F_XactLockTableWait(m, v130, v63, v84, int32(1))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
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
	v290 = m.ExcPending
	if v290 != 0 {
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
	v300 = m.ExcPending
	if v300 != 0 {
		goto L14
	} else {
		goto L98
	}
L98:
	;
	v304 = int32(0)
	goto L26
L99:
	;
	goto L10
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v310
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v51
	return
L101:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L14
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(_a_F_systable_inplace_update_begin_8), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L14
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_systable_inplace_update_begin_9), int32(829), int32(_a_F_systable_inplace_update_begin_10))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
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
	v337 = m.ExcPending
	if v337 != 0 {
		goto L14
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_systable_inplace_update_begin_9), int32(852), int32(_a_F_systable_inplace_update_begin_10))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
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
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
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
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v230 int64
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
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
	var v270 int64
	_ = v270
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v291 int64
	_ = v291
	var v292 int32
	_ = v292
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	v3 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
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
	v388 = m.ExcPending
	if v388 != 0 {
		goto L12
	} else {
		goto L101
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
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
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[6]))
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
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[7]))
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
	v129 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[4])) = v128 + v129
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[5]))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+120)) = v134 | v129
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v16)+48))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+118)))
	if v139 != int32(112) {
		goto L30
	} else {
		goto L31
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
	if v34 != 0 {
		goto L72
	} else {
		goto L73
	}
L31:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[2]))
	if v143 <= int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	if v146 != 0 {
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v19 < int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	if v147 != 0 {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+14)))
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+12)))
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[11]))) = uint16(v168)
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[13]))) = v171
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[15]))) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[16]))) = v110
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[1]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[17]))) = uint8(v177)
	F_XLogBeginInsert(m)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L12
	} else {
		goto L41
	}
L38:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[9]))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v151+(v19^int32(-1))<<(uint(int32(2))%32))))
	v165 = v157
	goto L37
L39:
	;
	goto L40
L40:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[10]))
	v165 = v159 + v19<<(uint(int32(13))%32) + int32(-8192)
	goto L37
L41:
	;
	F_XLogRegisterData(m, v22+int32(_a_F_systable_inplace_update_finish_6), int32(20))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	if v110 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_systable_inplace_update_finish[0])))
	F_XLogRegisterData(m, v186, v110<<(uint(int32(4))%32))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L12
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	if v167 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	base.MemoryCopy(m, v22+int32(32), v165, v167)
	goto L49
L48:
	;
	goto L49
L49:
	;
	v195 = int32(_a_F_systable_inplace_update_finish_7) - v166
	if v195 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	base.MemoryCopy(m, v22+int32(32)+v166, v165+v166, v195)
	goto L52
L51:
	;
	goto L52
L52:
	;
	if v34 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	base.MemoryCopy(m, v22+int32(32)+(v116-v165), v115, v34)
	goto L55
L54:
	;
	goto L55
L55:
	;
	v207 = v22 + int32(20)
	if v19 < int32(0) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[20]))
	if v241 <= int32(0) {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v230 = *(*int64)(unsafe.Add(mBase, uint32(v229)))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v229)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+8)) = v231
	*(*int64)(unsafe.Add(mBase, uint32(v207))) = v230
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(16)))) = v234
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v229)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v22+int32(12)))) = v236
	goto L56
L58:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[18]))
	v229 = v216 + (v19^int32(-1))<<(uint(int32(6))%32)
	goto L57
L59:
	;
	goto L60
L60:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[19]))
	v229 = v223 + v19<<(uint(int32(6))%32) + int32(-64)
	goto L57
L61:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[20])) = int32(1)
	goto L63
L62:
	;
	goto L63
L63:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[21]))
	if v248 <= int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L12
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[22]))
	v267 = v22 + int32(20)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v265)+12)) = v268
	v270 = *(*int64)(unsafe.Add(mBase, uint32(v267)))
	*(*int64)(unsafe.Add(mBase, uint32(v265)+4)) = v270
	v272 = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v265)+24)) = v22 + v272
	*(*int32)(unsafe.Add(mBase, uint32(v265)+20)) = v239
	*(*int32)(unsafe.Add(mBase, uint32(v265)+16)) = v238
	v277 = int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v265)+1)) = uint8(v277)
	v279 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v265)+28)) = v279
	v281 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v265))) = uint8(v281)
	*(*int32)(unsafe.Add(mBase, uint32(v265)+36)) = v265 + v272
	F_XLogRegisterBufData(m, v279, v115, v34)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L12
	} else {
		goto L70
	}
L67:
	;
	F_errmsg_internal(m, int32(_a_F_systable_inplace_update_finish_8), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L12
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_systable_inplace_update_finish_9), int32(320), int32(_a_F_systable_inplace_update_finish_10))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	v291 = F_XLogInsert(m, int32(10), int32(112))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v165))) = base.I64_rotr(v291, int64(32))
	goto L30
L72:
	;
	base.MemoryCopy(m, v116, v115, v34)
	goto L74
L73:
	;
	goto L74
L74:
	;
	F_MarkBufferDirty(m, v19)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	F_LockBuffer(m, v19, int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[3]))
	if v308 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+8))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
	v311 = v309 - v310
	if int32(0) < v311 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[5]))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v347)+120)) = v348 & int32(-2)
	v352 = int32(_a_F_systable_inplace_update_finish_5)
	v354 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[4])) = v354 - int32(1)
	F_UnlockTuple(m, v16, l1+int32(4), int32(7))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L12
	} else {
		goto L92
	}
L80:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[6]))
	F_SIInsertDataEntries(m, v315+v310<<(uint(int32(4))%32), v311)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L12
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v308)+12))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	v323 = v321 - v322
	if int32(0) < v323 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	goto L82
L84:
	;
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[7]))
	F_SIInsertDataEntries(m, v327+v322<<(uint(int32(4))%32), v323)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L12
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v334 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[3]))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+16)))
	if v335 == int32(1) {
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
	v339 = m.ExcPending
	if v339 != 0 {
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
	v364 = m.ExcPending
	if v364 != 0 {
		goto L12
	} else {
		goto L93
	}
L93:
	;
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_systable_inplace_update_finish[8]))
	if v366 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	F_CacheInvalidateHeapTuple(m, v16, l1, int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
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
	v381 = m.ExcPending
	if v381 != 0 {
		goto L12
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_systable_inplace_update_finish_2), int32(_a_F_systable_inplace_update_finish_3), int32(_a_F_systable_inplace_update_finish_4))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
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
