package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CommitTsPagePrecedes(m *base.Module, l0 int64, l1 int64) int32 {
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v6 = int32(819)
	v8 = int32(4)
	v9 = base.I32_wrap_i64(l0)*v6 + v8
	v12 = base.I32_wrap_i64(l1) * v6
	v14 = v12 + v8
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v14))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v9)) == int32(0) {
		v26 = base.B2i32(base.Ui32(v9) < base.Ui32(v14))
	} else {
		v26 = int32(base.Ui32(v9-v14) >> (uint(int32(31)) % 32))
	}
	if v26 != 0 {
		v28 = v12 + int32(822)
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v28))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v9)) == int32(0) {
			v40 = base.B2i32(base.Ui32(v9) < base.Ui32(v28))
		} else {
			v40 = int32(base.Ui32(v9-v28) >> (uint(int32(31)) % 32))
		}
		v42 = v40
	} else {
		v42 = int32(0)
	}
	return v42
}
func F_commit_ts_identify(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	if l0 == int32(16) {
		v6 = int32(_a_F_commit_ts_identify_0)
	} else {
		v6 = int32(0)
	}
	if l0 != 0 {
		v8 = v6
	} else {
		v8 = int32(_a_F_commit_ts_identify_1)
	}
	return v8
}
func F_commit_ts_redo(m *base.Module, l0 int32) {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v81 int64
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	v13 = v11 & int32(240)
	if v13 != 0 {
		if v13 == int32(16) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_commit_ts_redo[0]))
			v23 = F_LWLockAcquire(m, v19+int32(_a_F_commit_ts_redo_0), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_commit_ts_redo[1]))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+40))
				if v27 == int32(0) {
				} else {
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v17))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v27)) == int32(0) {
						v41 = base.B2i32(base.Ui32(v27) < base.Ui32(v17))
					} else {
						v41 = int32(base.Ui32(v27-v17) >> (uint(int32(31)) % 32))
					}
					if v41 == int32(0) {
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, _c_F_commit_ts_redo[1]))
						*(*int32)(unsafe.Add(mBase, uint32(v45)+40)) = v17
					}
				}
				v48 = *(*int32)(unsafe.Add(mBase, _c_F_commit_ts_redo[0]))
				F_LWLockRelease(m, v48+int32(_a_F_commit_ts_redo_0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					v53 = int32(_a_F_commit_ts_redo_1)
					v54 = *(*int32)(unsafe.Add(mBase, _c_F_commit_ts_redo[2]))
					v55 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
					*(*int64)(unsafe.Add(mBase, uint32(v54)+48)) = v55
					v58 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
					F_SimpleLruTruncate(m, v53, v58)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(23), int32(0))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v13
				F_errmsg_internal(m, int32(_a_F_commit_ts_redo_2), v8)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_commit_ts_redo_3), int32(1063), int32(_a_F_commit_ts_redo_4))
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
		}
	} else {
		v75 = *(*int32)(unsafe.Add(mBase, _c_F_commit_ts_redo[2]))
		v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+28))
		v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
		v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
		v80 = int64(*(*uint16)(unsafe.Add(mBase, _c_F_commit_ts_redo[3])))
		v81 = base.I64_rem_s(v78, v80)
		v85 = v76 + base.I32_wrap_i64(v81)<<(uint(int32(7))%32)
		v87 = F_LWLockAcquire(m, v85, int32(0))
		mBase = m.M
		v88 = m.ExcPending
		if v88 != 0 {
			return
		} else {
			v89 = int32(_a_F_commit_ts_redo_1)
			v91 = F_SimpleLruZeroPage(m, v89, v78)
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return
			} else {
				F_SimpleLruWritePage(m, v89, v91)
				mBase = m.M
				v94 = m.ExcPending
				if v94 != 0 {
					return
				} else {
					F_LWLockRelease(m, v85)
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_get_ts_config_oid(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13903(m, l0, l1, int32(_a_F_get_ts_config_oid_0), int32(3201), int32(_a_F_get_ts_config_oid_1), int32(73))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_get_ts_template_oid(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13903(m, l0, l1, int32(_a_F_get_ts_template_oid_0), int32(3056), int32(_a_F_get_ts_template_oid_1), int32(79))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_lookup_ts_config_cache(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v135 int64
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v310 int32
	_ = v310
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	v10 = m.G0
	v12 = v10 - int32(2560)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+2556)) = l0
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[0]))
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[1]))
	if v46 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+464)) = int64(85899345924)
	v25 = F_hash_create(m, int32(_a_F_lookup_ts_config_cache_0), int32(16), v12+int32(448), int32(40))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[0])) = v25
	F_CacheRegisterSyscacheCallback(m, int32(74), int32(1597), v25)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[0]))
	F_CacheRegisterSyscacheCallback(m, int32(72), int32(1597), v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[2]))
	if v41 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L1
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L3
	} else {
		goto L93
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L3
	} else {
		goto L90
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L3
	} else {
		goto L87
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L3
	} else {
		goto L84
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L3
	} else {
		goto L81
	}
L14:
	;
	m.G0 = v12 + int32(2560)
	return v310
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[0]))
	v57 = int32(0)
	v59 = F_hash_search(m, v54, v12+int32(2556), v57, v57)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L3
	} else {
		goto L20
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2556))
	if v49 != v50 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+4)))
	if v52 != 0 {
		v310 = v46
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[1])) = v299
	v310 = v299
	goto L14
L20:
	;
	if v59 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+4)))
	if v61 != 0 {
		v299 = v59
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2556))
	v64 = F_SearchSysCache1(m, int32(74), v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L3
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	if v64 == int32(0) {
		goto L13
	} else {
		goto L26
	}
L26:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+22)))
	v70 = v68 + v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+76))
	if v71 == int32(0) {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	if v59 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v135 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v128)+12)) = v135
	*(*int64)(unsafe.Add(mBase, uint32(v128)+4)) = v135
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2556))
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v139
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v70)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v128)+8)) = v141
	F_ReleaseCatCache(m, v64)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L3
	} else {
		goto L45
	}
L29:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[0]))
	v83 = F_hash_search(m, v77, v12+int32(2556), int32(1), v12+int32(448))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	if v85 == int32(0) {
		v128 = v59
		goto L28
	} else {
		goto L33
	}
L32:
	;
	v128 = v83
	goto L28
L33:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	if int32(0) < v88 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v92 = int32(0)
	v95 = v88
	goto L37
L35:
	;
	v123 = v85
	goto L36
L36:
	;
	F_pfree(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L3
	} else {
		goto L44
	}
L37:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101+v92<<(uint(int32(3))%32))+4))
	if v105 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v123 = v113
	goto L36
L39:
	;
	F_pfree(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L3
	} else {
		goto L42
	}
L40:
	;
	v109 = v95
	goto L41
L41:
	;
	v111 = v92 + int32(1)
	if v111 < v109 {
		v92 = v111
		v95 = v109
		goto L37
	} else {
		goto L43
	}
L42:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v109 = v108
	goto L41
L43:
	;
	goto L38
L44:
	;
	v128 = v59
	goto L28
L45:
	;
	v145 = int32(0)
	base.MemoryFill(m, v12+int32(448), v145, int32(2056))
	v152 = v12 + int32(2508)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2556))
	F_ScanKeyInit(m, v152, int32(1), int32(3), int32(184), v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	v161 = F_table_open(m, int32(3603), int32(1))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L3
	} else {
		goto L48
	}
L47:
	;
	F_systable_endscan_ordered(m, v169)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L3
	} else {
		goto L70
	}
L48:
	;
	v165 = F_index_open(m, int32(3609), int32(1))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v169 = F_systable_beginscan_ordered(m, v161, v165, int32(0), int32(1), v152)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	v172 = F_systable_getnext_ordered(m, v169, int32(1))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	if v172 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v239 = int32(0)
	v243 = v145
	goto L47
L53:
	;
	goto L54
L54:
	;
	v178 = int32(0)
	v181 = v172
	v182 = v145
	goto L55
L55:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v181)+16))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+22)))
	v189 = v187 + v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	if base.Ui32(v190-int32(257)) <= base.Ui32(int32(-257)) {
		goto L11
	} else {
		goto L57
	}
L56:
	;
	v239 = v235
	v243 = v234
	goto L47
L57:
	;
	if v190 < v182 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	if v182 < v190 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v237 = F_systable_getnext_ordered(m, v169, int32(1))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L3
	} else {
		goto L68
	}
L60:
	;
	if v178 <= int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	if int32(100) <= v178 {
		goto L9
	} else {
		goto L67
	}
L63:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v219
	v234 = v190
	v235 = int32(1)
	goto L59
L64:
	;
	v203 = v12 + int32(448) + v182<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v203))) = v178
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[2]))
	v208 = v178 << (uint(int32(2)) % 32)
	v209 = F_MemoryContextAlloc(m, v206, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+4)) = v209
	if v208 == int32(0) {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	base.MemoryCopy(m, v209, v12+int32(48), v208)
	goto L63
L67:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v189)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(48)+v178<<(uint(int32(2))%32)))) = v229
	v234 = v182
	v235 = v178 + int32(1)
	goto L59
L68:
	;
	if v237 != 0 {
		v178 = v235
		v181 = v237
		v182 = v234
		goto L55
	} else {
		goto L69
	}
L69:
	;
	goto L56
L70:
	;
	F_relation_close(m, v165, int32(1))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	F_relation_close(m, v161, int32(1))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L3
	} else {
		goto L72
	}
L72:
	;
	if v239 <= int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v295 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v128)+4)) = uint8(v295)
	v299 = v128
	goto L19
L74:
	;
	v262 = v12 + int32(448) + v243<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v262))) = v239
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[2]))
	v267 = v239 << (uint(int32(2)) % 32)
	v268 = F_MemoryContextAlloc(m, v265, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262)+4)) = v268
	if v267 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	base.MemoryCopy(m, v268, v12+int32(48), v267)
	goto L78
L77:
	;
	goto L78
L78:
	;
	v275 = v243 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v128)+12)) = v275
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[2]))
	v281 = F_MemoryContextAlloc(m, v278, v275<<(uint(int32(3))%32))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128)+16)) = v281
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v286 = v284 << (uint(int32(3)) % 32)
	if v286 == int32(0) {
		goto L73
	} else {
		goto L80
	}
L80:
	;
	base.MemoryCopy(m, v281, v12+int32(448), v286)
	goto L73
L81:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2556))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v325
	F_errmsg_internal(m, int32(_a_F_lookup_ts_config_cache_1), v12)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L3
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_config_cache_2), int32(426), int32(_a_F_lookup_ts_config_cache_3))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L3
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2556))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v339
	F_errmsg_internal(m, int32(_a_F_lookup_ts_config_cache_4), v12+int32(16))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L3
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_config_cache_2), int32(433), int32(_a_F_lookup_ts_config_cache_3))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L3
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v190
	F_errmsg_internal(m, int32(_a_F_lookup_ts_config_cache_5), v12+int32(32))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L3
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_config_cache_2), int32(491), int32(_a_F_lookup_ts_config_cache_3))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L3
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	F_errmsg_internal(m, int32(_a_F_lookup_ts_config_cache_6), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L3
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_config_cache_2), int32(493), int32(_a_F_lookup_ts_config_cache_3))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L3
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errmsg_internal(m, int32(_a_F_lookup_ts_config_cache_7), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L3
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_config_cache_2), int32(514), int32(_a_F_lookup_ts_config_cache_3))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L3
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ts_process_call(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
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
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16+v17<<(uint(int32(2))%32))))
	if v21 == v2 {
		v153 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v13 + int32(80)
	return v153
L2:
	;
	v25 = v21
	v26 = v17
	goto L3
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v34 != 0 {
		v77 = v25
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v153 = v2
	goto L1
L5:
	;
	if v26 == int32(0) {
		v153 = v2
		goto L1
	} else {
		goto L24
	}
L6:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v89 = F_palloc(m, v86+int32(1))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v35 == int32(0) {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v39 = v26 + int32(1)
	v42 = v16 + v39<<(uint(int32(2))%32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v35 == v43 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v39
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	if v48 == int32(0) {
		v77 = v46
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v55 = v46 + int32(8)
	goto L11
L11:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v65 = v63 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	*(*int32)(unsafe.Add(mBase, uint32(v67+v65<<(uint(int32(2))%32)))) = v71
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	if v75 != 0 {
		v55 = v71 + int32(8)
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v77 = v71
	goto L6
L13:
	;
	goto L12
L14:
	;
	return int32(0)
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v89
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	if v94 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	base.MemoryCopy(m, v89, v77+int32(20), v94)
	goto L18
L17:
	;
	goto L18
L18:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
	v100 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v89+v98))) = uint8(v100)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v102
	v105 = v13 + int32(48)
	v109 = F_pg_sprintf(m, v105, int32(_a_F_ts_process_call_0), v13+int32(16))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v105
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v112
	v115 = v13 + int32(32)
	v117 = F_pg_sprintf(m, v115, int32(_a_F_ts_process_call_0), v13)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v115
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v123 = F_BuildTupleFromCStrings(m, v120, v13+int32(68))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	v126 = F_HeapTupleHeaderGetDatum(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	F_pfree(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L14
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(0)
	v153 = v126
	goto L1
L24:
	;
	v138 = v26 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v138
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v16+v138<<(uint(int32(2))%32))))
	if v143 != 0 {
		v25 = v143
		v26 = v138
		goto L3
	} else {
		goto L25
	}
L25:
	;
	goto L4
}
func F_ts_setup_firstcall(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = l2
	v13 = int32(_a_F_ts_setup_firstcall_0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_ts_setup_firstcall[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_ts_setup_firstcall[0])) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v23 = F_palloc0(m, v18<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v23
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
		if v28 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = v28
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
			if v30 == int32(0) {
			} else {
				v39 = v28 + int32(8)
				for {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
					v44 = v42 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v44
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					*(*int32)(unsafe.Add(mBase, uint32(v46+v44<<(uint(int32(2))%32)))) = v50
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
					if v54 != 0 {
						v39 = v50 + int32(8)
						continue
					} else {
						break
					}
					break
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(0)
		}
		v67 = F_get_call_result_type(m, l0, int32(0), v10+int32(12))
		mBase = m.M
		v68 = m.ExcPending
		if v68 != 0 {
			return
		} else {
			if v67 == int32(1) {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v71
				v73 = F_TupleDescGetAttInMetadata(m, v71)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v73
					*(*int32)(unsafe.Add(mBase, _c_F_ts_setup_firstcall[0])) = v14
					m.G0 = v10 + int32(16)
					return
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_ts_setup_firstcall_1), int32(0))
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ts_setup_firstcall_2), int32(2481), int32(_a_F_ts_setup_firstcall_3))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
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
}
