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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	v7 = int32(0)
	if l2 == v7 {
		v32 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v34 = F_palloc(m, int32(24))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L9
	} else {
		goto L11
	}
L2:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, _consts[112])))
	if v17 != 0 {
		v32 = v7
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	if v21 != l1 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v26 != 0 {
		v32 = v7
		goto L1
	} else {
		goto L8
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	v25 = F_list_member_ptr(m, v24, l1)
	mBase = m.M
	v26 = v25
	goto L7
L6:
	;
	v26 = int32(1)
	goto L7
L7:
	;
	goto L4
L8:
	;
	v28 = F_index_open(m, l1, int32(1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v32 = v28
	goto L1
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = l0
	v39 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v39
	if l3 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v45 = F_GetCatalogSnapshot(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	v49 = l3
	v50 = v7
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v50
	if v32 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v47 = F_RegisterSnapshot(m, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v49 = v47
	v50 = v47
	goto L15
L18:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	if v214 != 0 {
		goto L47
	} else {
		goto L48
	}
L19:
	;
	v54 = F_palloc(m, l4*int32(48))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
	v195 = m.T0[v194].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v49, l4, l5, int32(0), int32(321))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
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
	v178 = int32(0)
	v180 = F_index_beginscan(m, l0, v32, v49, v178, l4, v178)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
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
	v72 = v70 * int32(48)
	v73 = v54 + v72
	v74 = l5 + v72
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v74)))
	*(*int64)(unsafe.Add(mBase, uint32(v73))) = v75
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v74)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+40)) = v77
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v74)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+32)) = v79
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v74)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+24)) = v81
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v74)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+16)) = v83
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v74)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+8)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v32)+192))
	v88 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87)+8)))
	if v88 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L9
	} else {
		goto L40
	}
L27:
	;
	goto L26
L28:
	;
	if v124 == v128 {
		goto L27
	} else {
		goto L38
	}
L29:
	;
	v124 = v88
	v128 = int32(0)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)))
	v102 = int32(0)
	goto L32
L32:
	;
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87+int32(48)+v102<<(uint(int32(1))%32)))))
	if v112 == v95 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L27
L34:
	;
	v115 = v102 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v73)+4)) = uint16(v115)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v32)+192))
	v118 = int32(*(*int16)(unsafe.Add(mBase, uint32(v117)+8)))
	v124 = v118
	v128 = v102
	goto L28
L35:
	;
	goto L36
L36:
	;
	v120 = v102 + int32(1)
	if v120 != v88 {
		v102 = v120
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v137 = v70 + int32(1)
	if l4 != v137 {
		v70 = v137
		goto L25
	} else {
		goto L39
	}
L39:
	;
	goto L23
L40:
	;
	F_errmsg_internal(m, int32(28345), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(494121), int32(446), int32(282736))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v180
	v183 = int32(0)
	F_index_rescan(m, v180, v54, l4, v183, v183)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = int32(0)
	F_pfree(m, v54)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L9
	} else {
		goto L45
	}
L45:
	;
	goto L18
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v195
	goto L18
L47:
	;
	v216 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[116])) = uint8(v216)
	goto L49
L48:
	;
	goto L49
L49:
	;
	return v34
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
						v15 = *(*int32)(unsafe.Add(mBase, _consts[115]))
						if v15 != 0 {
							v17 = int32(0)
							*(*uint8)(unsafe.Add(mBase, _consts[116])) = uint8(v17)
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
					v15 = *(*int32)(unsafe.Add(mBase, _consts[115]))
					if v15 != 0 {
						v17 = int32(0)
						*(*uint8)(unsafe.Add(mBase, _consts[116])) = uint8(v17)
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
					v15 = *(*int32)(unsafe.Add(mBase, _consts[115]))
					if v15 != 0 {
						v17 = int32(0)
						*(*uint8)(unsafe.Add(mBase, _consts[116])) = uint8(v17)
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
				v15 = *(*int32)(unsafe.Add(mBase, _consts[115]))
				if v15 != 0 {
					v17 = int32(0)
					*(*uint8)(unsafe.Add(mBase, _consts[116])) = uint8(v17)
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	v21 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
	if v22 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L14
	} else {
		goto L106
	}
L2:
	;
	if v24&int32(1) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v24 = int32(1)
	goto L5
L4:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+76)))
	v24 = v23
	goto L5
L5:
	;
	goto L2
L6:
	;
	v42 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L14
	} else {
		goto L102
	}
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v46 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v316 = F_heap_copytuple(m, v56)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L14
	} else {
		goto L101
	}
L11:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if v42 == int32(10001) {
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
	v51 = int32(1)
	v54 = F_systable_beginscan(m, l0, l1, v51, int32(0), v51, l2)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v56 = F_systable_getnext(m, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	if v56 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_systable_endscan(m, v54)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L14
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+40))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+68))
	v70 = m.G0
	v72 = v70 - int32(32)
	m.G0 = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v72)+24)) = v74
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v68)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v72)+16)) = v76
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v68)))
	*(*int64)(unsafe.Add(mBase, uint32(v72)+8)) = v78
	F_CacheInvalidateHeapTupleCommon(m, v66, v68, int32(0), int32(1600))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
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
	v87 = v72 + int32(8) | int32(4)
	F_LockTuple(m, v66, v87, int32(7))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L14
	} else {
		goto L24
	}
L24:
	;
	F_LockBuffer(m, v69, int32(2))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	v94 = int32(1)
	v98 = F_GetCurrentCommandId(m, int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L14
	} else {
		goto L32
	}
L26:
	;
	m.G0 = v72 + int32(32)
	if v307 == int32(0) {
		v42 = v42 + int32(1)
		goto L9
	} else {
		goto L100
	}
L27:
	;
	F_UnlockTuple(m, v66, v87, int32(7))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L14
	} else {
		goto L97
	}
L28:
	;
	F_LockBuffer(m, v69, int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L14
	} else {
		goto L95
	}
L29:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v72)+24))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v134)+20)))
	if v136&int32(4096) != 0 {
		goto L42
	} else {
		goto L43
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L14
	} else {
		goto L38
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L14
	} else {
		goto L34
	}
L32:
	;
	v100 = F_HeapTupleSatisfiesUpdate(m, v72+int32(8), v98, v69)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L14
	} else {
		goto L33
	}
L33:
	;
	switch v100 {
	case 0:
		v307 = v94
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
	v108 = m.ExcPending
	if v108 != 0 {
		goto L14
	} else {
		goto L35
	}
L35:
	;
	F_errmsg_internal(m, int32(382056), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L14
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(494112), int32(6413), int32(315084))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
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
	v124 = m.ExcPending
	if v124 != 0 {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	F_errmsg(m, int32(425450), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(494112), int32(6425), int32(315084))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
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
	v141 = F_DoesMultiXactIdConflict(m, v135, v136, int32(2), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L14
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if base.Ui32(v135) < base.Ui32(int32(3)) {
		goto L51
	} else {
		goto L52
	}
L45:
	;
	if v141 == int32(0) {
		v307 = v94
		goto L26
	} else {
		goto L46
	}
L46:
	;
	F_LockBuffer(m, v69, int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L14
	} else {
		goto L47
	}
L47:
	;
	F_systable_endscan(m, v54)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L14
	} else {
		goto L48
	}
L48:
	;
	v150 = int32(4)
	v151 = int32(0)
	v156 = F_Do_MultiXactIdWait(m, v135, v150, v136, v151, v66, v87, int32(1), v72+v150, v151)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L14
	} else {
		goto L49
	}
L49:
	;
	goto L27
L50:
	;
	if v136&int32(80) == int32(16) {
		v307 = v94
		goto L26
	} else {
		goto L90
	}
L51:
	;
	v277 = int32(0)
	goto L50
L52:
	;
	goto L53
L53:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	if v168 == v135 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v277 = int32(1)
	goto L50
L55:
	;
	goto L56
L56:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	if v172 <= int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v277 = v269
	goto L50
L58:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	if v176 == int32(0) {
		v269 = int32(0)
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _consts[92]))
	v240 = int32(0)
	v242 = v172 - int32(1)
	goto L80
L61:
	;
	v181 = v176
	goto L62
L62:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
	if v186 == int32(4) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v269 = int32(0)
	goto L57
L64:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v181)+80))
	if v233 != 0 {
		v181 = v233
		goto L62
	} else {
		goto L79
	}
L65:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	if v189 == int32(0) {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v192 = int32(1)
	if v135 == v189 {
		v269 = v192
		goto L57
	} else {
		goto L67
	}
L67:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v181)+52))
	v196 = v194 - int32(1)
	if v196 < int32(0) {
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v201 = int32(0)
	v203 = v196
	goto L69
L69:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v181)+48))
	v209 = int32(2)
	v210 = base.I32_div_s(v203-v201, v209)
	v211 = v210 + v201
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v207+v211<<(uint(v209)%32))))
	if v215 == v135 {
		v269 = v192
		goto L57
	} else {
		goto L71
	}
L70:
	;
	goto L64
L71:
	;
	v219 = F_TransactionIdPrecedes(m, v215, v135)
	mBase = m.M
	if v219 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v220 = v211 + int32(1)
	goto L74
L73:
	;
	v220 = v201
	goto L74
L74:
	;
	if v219 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v223 = v203
	goto L77
L76:
	;
	v223 = v211 - int32(1)
	goto L77
L77:
	;
	if v220 <= v223 {
		v201 = v220
		v203 = v223
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
	v247 = int32(2)
	v248 = base.I32_div_s(v242-v240, v247)
	v249 = v248 + v240
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v238+v249<<(uint(v247)%32))))
	v254 = base.B2i32(v253 == v135)
	if v253 == v135 {
		v269 = v254
		goto L57
	} else {
		goto L82
	}
L81:
	;
	v269 = v254
	goto L57
L82:
	;
	v257 = base.B2i32(base.Ui32(v253) < base.Ui32(v135))
	if base.Ui32(v253) < base.Ui32(v135) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v258 = v249 + int32(1)
	goto L85
L84:
	;
	v258 = v240
	goto L85
L85:
	;
	if base.Ui32(v253) < base.Ui32(v135) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v261 = v242
	goto L88
L87:
	;
	v261 = v249 - int32(1)
	goto L88
L88:
	;
	if v258 <= v261 {
		v240 = v258
		v242 = v261
		goto L80
	} else {
		goto L89
	}
L89:
	;
	goto L81
L90:
	;
	if v277 != 0 {
		v307 = v94
		goto L26
	} else {
		goto L91
	}
L91:
	;
	F_LockBuffer(m, v69, int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L14
	} else {
		goto L92
	}
L92:
	;
	F_systable_endscan(m, v54)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L14
	} else {
		goto L93
	}
L93:
	;
	F_XactLockTableWait(m, v135, v66, v87, int32(1))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L14
	} else {
		goto L94
	}
L94:
	;
	goto L27
L95:
	;
	F_systable_endscan(m, v54)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L14
	} else {
		goto L96
	}
L96:
	;
	goto L27
L97:
	;
	*(*int32)(unsafe.Add(mBase, _consts[117])) = int32(0)
	goto L98
L98:
	;
	F_InvalidateCatalogSnapshot(m)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L14
	} else {
		goto L99
	}
L99:
	;
	v307 = int32(0)
	goto L26
L100:
	;
	goto L10
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v316
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v54
	return
L102:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L14
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(258889), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L14
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(494121), int32(829), int32(275341))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L14
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	F_errmsg_internal(m, int32(30253), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L14
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(494121), int32(852), int32(275341))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L14
	} else {
		goto L108
	}
L108:
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
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
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v235 int64
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int64
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v296 int64
	_ = v296
	var v297 int32
	_ = v297
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	v3 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+68))
	v21 = m.G0
	v23 = v21 - int32(8256)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[118]))) = v3
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[119]))) = uint8(v3)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
	if v30 != v32 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_systable_endscan(m, l0)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L12
	} else {
		goto L107
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L12
	} else {
		goto L104
	}
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v35 = v34 - v30
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v35 != v36-v32 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if int32(0) < v40 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	if v48 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v112 = v3
	goto L7
L7:
	;
	v117 = v31 + v32
	v118 = v25 + v30
	v120 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	if v120 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L8:
	;
	v112 = v111
	goto L7
L9:
	;
	v51 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[119]))) = uint8(v51)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[118]))) = v51
	v111 = v51
	goto L8
L10:
	;
	goto L11
L11:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[119]))) = uint8(v56)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v67 = F_palloc(m, (v58+v59-(v61+v62))<<(uint(int32(4))%32))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[118]))) = v67
	v71 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v74 = v72 - v73
	if int32(0) < v74 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[121]))
	v79 = int32(4)
	v83 = v74 << (uint(v79) % 32)
	if v83 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v86 = v3
	goto L16
L16:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v89 = v87 - v88
	if int32(0) < v89 {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v86 = v74
	goto L16
L18:
	;
	v84 = F__emscripten_memcpy_bulkmem(m, v67, v78+v73<<(uint(v79)%32), v83)
	mBase = m.M
	goto L20
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	v92 = int32(4)
	v96 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	v101 = v89 << (uint(v92) % 32)
	if v101 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v105 = v86
	goto L23
L23:
	;
	v111 = v105
	goto L8
L24:
	;
	v105 = v86 + v89
	goto L23
L25:
	;
	v102 = F__emscripten_memcpy_bulkmem(m, v67+v86<<(uint(v92)%32), v96+v88<<(uint(v92)%32), v101)
	mBase = m.M
	goto L27
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	v128 = int32(4481700)
	v130 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v131 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v130 + v131
	v135 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+120)) = v136 | v131
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v16)+48))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+118)))
	if v141 != int32(112) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+16)))
	if v123 != int32(1) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	F_RelationCacheInitFilePreInvalidate(m)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	if v35 != 0 {
		goto L78
	} else {
		goto L79
	}
L33:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v145 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	if v148 != 0 {
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v19 < int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	if v149 != 0 {
		goto L32
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+14)))
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+12)))
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[124]))) = uint16(v170)
	v173 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[125]))) = v173
	v176 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[127]))) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[128]))) = v112
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[119]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[129]))) = uint8(v179)
	F_XLogBeginInsert(m)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L12
	} else {
		goto L43
	}
L40:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v153+(v19^int32(-1))<<(uint(int32(2))%32))))
	v167 = v159
	goto L39
L41:
	;
	goto L42
L42:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v167 = v161 + v19<<(uint(int32(13))%32) + int32(-8192)
	goto L39
L43:
	;
	F_XLogRegisterData(m, v23+int32(8228), int32(20))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L12
	} else {
		goto L44
	}
L44:
	;
	if v112 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[118])))
	F_XLogRegisterData(m, v188, v112<<(uint(int32(4))%32))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L12
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if v169 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L47
L49:
	;
	v202 = int32(8192) - v168
	if v202 != 0 {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v195 = F__emscripten_memcpy_bulkmem(m, v23+int32(32), v167, v169)
	mBase = m.M
	goto L52
L51:
	;
	goto L52
L52:
	;
	goto L49
L53:
	;
	if v35 != 0 {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	v203 = F__emscripten_memcpy_bulkmem(m, v23+int32(32)+v168, v167+v168, v202)
	mBase = m.M
	goto L56
L55:
	;
	goto L56
L56:
	;
	goto L53
L57:
	;
	v212 = v23 + int32(20)
	if v19 < int32(0) {
		goto L63
	} else {
		goto L64
	}
L58:
	;
	v209 = F__emscripten_memcpy_bulkmem(m, v23+int32(32)+(v118-v167), v117, v35)
	mBase = m.M
	goto L60
L59:
	;
	goto L60
L60:
	;
	goto L57
L61:
	;
	v244 = v23 + int32(20)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v250 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	if v250 <= int32(0) {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v234)))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v234)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v212)+8)) = v236
	*(*int64)(unsafe.Add(mBase, uint32(v212))) = v235
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(16)))) = v239
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v234)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(12)))) = v241
	goto L61
L63:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v234 = v221 + (v19^int32(-1))<<(uint(int32(6))%32)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v234 = v228 + v19<<(uint(int32(6))%32) + int32(-64)
	goto L62
L66:
	;
	*(*int32)(unsafe.Add(mBase, _consts[130])) = int32(1)
	goto L68
L67:
	;
	goto L68
L68:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	if v257 <= int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L12
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	v275 = *(*int64)(unsafe.Add(mBase, uint32(v244)))
	*(*int64)(unsafe.Add(mBase, uint32(v274)+4)) = v275
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v274)+24)) = v23 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v274)+20)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v274)+16)) = v245
	v281 = int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v274)+1)) = uint8(v281)
	v283 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v274)+28)) = v283
	v285 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v274))) = uint8(v285)
	*(*int32)(unsafe.Add(mBase, uint32(v274)+12)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v274)+36)) = v274 + int32(32)
	F_XLogRegisterBufData(m, v283, v117, v35)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L12
	} else {
		goto L75
	}
L72:
	;
	F_errmsg_internal(m, int32(134508), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(489916), int32(320), int32(315148))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	v296 = F_XLogInsert(m, int32(10), int32(112))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v167))) = base.I64_rotr(v296, int64(32))
	goto L32
L77:
	;
	F_MarkBufferDirty(m, v19)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L12
	} else {
		goto L81
	}
L78:
	;
	v307 = F__emscripten_memcpy_bulkmem(m, v118, v117, v35)
	mBase = m.M
	goto L80
L79:
	;
	goto L80
L80:
	;
	goto L77
L81:
	;
	F_LockBuffer(m, v19, int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L12
	} else {
		goto L82
	}
L82:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	if v315 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)+8))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
	v318 = v316 - v317
	if int32(0) < v318 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v354)+120)) = v355 & int32(-2)
	v359 = int32(4481700)
	v361 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v361 - int32(1)
	F_UnlockTuple(m, v16, l1+int32(4), int32(7))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L12
	} else {
		goto L98
	}
L86:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _consts[121]))
	F_SendSharedInvalidMessages(m, v322+v317<<(uint(int32(4))%32), v318)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L12
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v315)+12))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v315)+4))
	v330 = v328 - v329
	if int32(0) < v330 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L88
L90:
	;
	v334 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	F_SendSharedInvalidMessages(m, v334+v329<<(uint(int32(4))%32), v330)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L12
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+16)))
	if v342 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	goto L92
L94:
	;
	F_RelationCacheInitFilePostInvalidate(m)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L12
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, _consts[117])) = int32(0)
	goto L85
L97:
	;
	goto L96
L98:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L12
	} else {
		goto L99
	}
L99:
	;
	v373 = *(*int32)(unsafe.Add(mBase, _consts[123]))
	if v373 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	F_CacheInvalidateHeapTuple(m, v16, l1, int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L12
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	m.G0 = v23 + int32(8256)
	goto L1
L103:
	;
	goto L102
L104:
	;
	F_errmsg_internal(m, int32(319172), int32(0))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(494112), int32(6520), int32(314232))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L12
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
