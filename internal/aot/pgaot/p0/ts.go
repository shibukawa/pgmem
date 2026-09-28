package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CommitTsPagePrecedes(m *base.Module, l0 int64, l1 int64) int32 {
	var v3 int32
	_ = v3
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	v3 = int32(0)
	v9 = int32(819)
	v10 = base.I32_wrap_i64(l0) * v9
	v11 = int32(4)
	v12 = v10 + v11
	v13 = int32(3)
	v17 = base.I32_wrap_i64(l1) * v9
	v19 = v17 + v11
	if base.B2i32(base.Ui32(v12) < base.Ui32(v13))|base.B2i32(base.Ui32(v19) < base.Ui32(v13)) == v3 {
		if v10-v17 < int32(0) {
			v30 = v17 + int32(822)
			if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v30))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v12)) == int32(0) {
				return base.B2i32(base.Ui32(v12) < base.Ui32(v30))
			} else {
				v44 = int32(base.Ui32(v12-v30) >> (uint(int32(31)) % 32))
				return v44
			}
		} else {
			v44 = v3
			return v44
		}
	} else {
		if base.Ui32(v19) <= base.Ui32(v12) {
			v44 = v3
			return v44
		} else {
			v30 = v17 + int32(822)
			if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v30))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v12)) == int32(0) {
				return base.B2i32(base.Ui32(v12) < base.Ui32(v30))
			} else {
				v44 = int32(base.Ui32(v12-v30) >> (uint(int32(31)) % 32))
				return v44
			}
		}
	}
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
	var v30 int32
	_ = v30
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
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
					v30 = int32(3)
					if base.B2i32(base.Ui32(v17) < base.Ui32(v30))|base.B2i32(base.Ui32(v27) < base.Ui32(v30)) == int32(0) {
						if v27-v17 < int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v17
						} else {
						}
					} else {
						if base.Ui32(v17) <= base.Ui32(v27) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v17
						}
					}
				}
				v43 = *(*int32)(unsafe.Add(mBase, _c_F_commit_ts_redo[0]))
				F_LWLockRelease(m, v43+int32(_a_F_commit_ts_redo_0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, _c_F_commit_ts_redo[2]))
					v50 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
					v52 = base.AtomicRmwXchg64(m, v49, int32(48), v50)
					v54 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
					F_SimpleLruTruncate(m, int32(_a_F_commit_ts_redo_1), v54)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(24), int32(0))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v13
				F_errmsg_internal(m, int32(_a_F_commit_ts_redo_2), v8)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_commit_ts_redo_3), int32(1025), int32(_a_F_commit_ts_redo_4))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
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
		v71 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
		v72 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
		F_SimpleLruZeroAndWritePage(m, int32(_a_F_commit_ts_redo_1), v72)
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return
		} else {
			m.G0 = v8 + int32(16)
			return
		}
	}
}
func F_get_ts_config_oid(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14294(m, l0, l1, int32(_a_F_get_ts_config_oid_0), int32(3270), int32(_a_F_get_ts_config_oid_1), int32(73))
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
	v7 = Fn14294(m, l0, l1, int32(_a_F_get_ts_template_oid_0), int32(3125), int32(_a_F_get_ts_template_oid_1), int32(79))
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
	var v34 int32
	_ = v34
	var v38 int64
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
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int64
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v157 int64
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v311 int32
	_ = v311
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	v10 = m.G0
	v12 = v10 - int32(2576)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+2572)) = l0
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[0]))
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[1]))
	if v47 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+456)) = int64(85899345924)
	v25 = F_hash_create(m, int32(_a_F_lookup_ts_config_cache_0), int64(16), v12+int32(448), int32(40))
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
	F_CacheRegisterSyscacheCallback(m, int32(74), int32(1811), base.I64_extend_i32_u(v25))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v38 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[0])))
	F_CacheRegisterSyscacheCallback(m, int32(72), int32(1811), v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[2]))
	if v42 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
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
	v383 = m.ExcPending
	if v383 != 0 {
		goto L3
	} else {
		goto L93
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L3
	} else {
		goto L90
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L3
	} else {
		goto L87
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L3
	} else {
		goto L84
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L3
	} else {
		goto L81
	}
L14:
	;
	m.G0 = v12 + int32(2576)
	return v311
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[0]))
	v58 = int32(0)
	v60 = F_hash_search(m, v55, v12+int32(2572), v58, v58)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
	} else {
		goto L20
	}
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2572))
	if v50 != v51 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+4)))
	if v53 != 0 {
		v311 = v47
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[1])) = v300
	v311 = v300
	goto L14
L20:
	;
	if v60 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)))
	if v62 != 0 {
		v300 = v60
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v64 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+2572)))
	v65 = F_SearchSysCache1(m, int32(74), v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L3
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	if v65 == int32(0) {
		goto L13
	} else {
		goto L26
	}
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+22)))
	v71 = v69 + v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+76))
	if v72 == int32(0) {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	if v60 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v136 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v129)+12)) = v136
	*(*int64)(unsafe.Add(mBase, uint32(v129)+4)) = v136
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2572))
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v140
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v71)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v129)+8)) = v142
	F_ReleaseCatCache(m, v65)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L3
	} else {
		goto L45
	}
L29:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[0]))
	v84 = F_hash_search(m, v78, v12+int32(2572), int32(1), v12+int32(448))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	if v86 == int32(0) {
		v129 = v60
		goto L28
	} else {
		goto L33
	}
L32:
	;
	v129 = v84
	goto L28
L33:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	if int32(0) < v89 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v93 = int32(0)
	v96 = v89
	goto L37
L35:
	;
	v124 = v86
	goto L36
L36:
	;
	F_pfree(m, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L3
	} else {
		goto L44
	}
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+v93<<(uint(int32(3))%32))+4))
	if v106 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v124 = v114
	goto L36
L39:
	;
	F_pfree(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L42
	}
L40:
	;
	v110 = v96
	goto L41
L41:
	;
	v112 = v93 + int32(1)
	if v112 < v110 {
		v93 = v112
		v96 = v110
		goto L37
	} else {
		goto L43
	}
L42:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v110 = v109
	goto L41
L43:
	;
	goto L38
L44:
	;
	v129 = v60
	goto L28
L45:
	;
	v146 = int32(0)
	base.MemoryFill(m, v12+int32(448), v146, int32(2056))
	v153 = v12 + int32(2512)
	v157 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+2572)))
	F_ScanKeyInit(m, v153, int32(1), int32(3), int32(184), v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	v162 = F_table_open(m, int32(3603), int32(1))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L3
	} else {
		goto L48
	}
L47:
	;
	F_systable_endscan_ordered(m, v170)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L3
	} else {
		goto L70
	}
L48:
	;
	v166 = F_index_open(m, int32(3609), int32(1))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v170 = F_systable_beginscan_ordered(m, v162, v166, int32(0), int32(1), v153)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L3
	} else {
		goto L50
	}
L50:
	;
	v173 = F_systable_getnext_ordered(m, v170, int32(1))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L3
	} else {
		goto L51
	}
L51:
	;
	if v173 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v240 = int32(0)
	v244 = v146
	goto L47
L53:
	;
	goto L54
L54:
	;
	v179 = int32(0)
	v182 = v173
	v183 = v146
	goto L55
L55:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v182)+16))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+22)))
	v190 = v188 + v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	if base.Ui32(v191-int32(257)) <= base.Ui32(int32(-257)) {
		goto L11
	} else {
		goto L57
	}
L56:
	;
	v240 = v236
	v244 = v235
	goto L47
L57:
	;
	if v191 < v183 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	if v183 < v191 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v238 = F_systable_getnext_ordered(m, v170, int32(1))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L3
	} else {
		goto L68
	}
L60:
	;
	if v179 <= int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	if int32(100) <= v179 {
		goto L9
	} else {
		goto L67
	}
L63:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v220
	v235 = v191
	v236 = int32(1)
	goto L59
L64:
	;
	v204 = v12 + int32(448) + v183<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v179
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[2]))
	v209 = v179 << (uint(int32(2)) % 32)
	v210 = F_MemoryContextAlloc(m, v207, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+4)) = v210
	if v209 == int32(0) {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	base.MemoryCopy(m, v210, v12+int32(48), v209)
	goto L63
L67:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v190)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(48)+v179<<(uint(int32(2))%32)))) = v230
	v235 = v183
	v236 = v179 + int32(1)
	goto L59
L68:
	;
	if v238 != 0 {
		v179 = v236
		v182 = v238
		v183 = v235
		goto L55
	} else {
		goto L69
	}
L69:
	;
	goto L56
L70:
	;
	F_relation_close(m, v166, int32(1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	F_relation_close(m, v162, int32(1))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L3
	} else {
		goto L72
	}
L72:
	;
	if v240 <= int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v296 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v129)+4)) = uint8(v296)
	v300 = v129
	goto L19
L74:
	;
	v263 = v12 + int32(448) + v244<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v263))) = v240
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[2]))
	v268 = v240 << (uint(int32(2)) % 32)
	v269 = F_MemoryContextAlloc(m, v266, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263)+4)) = v269
	if v268 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	base.MemoryCopy(m, v269, v12+int32(48), v268)
	goto L78
L77:
	;
	goto L78
L78:
	;
	v276 = v244 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v129)+12)) = v276
	v279 = *(*int32)(unsafe.Add(mBase, _c_F_lookup_ts_config_cache[2]))
	v282 = F_MemoryContextAlloc(m, v279, v276<<(uint(int32(3))%32))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+16)) = v282
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	v287 = v285 << (uint(int32(3)) % 32)
	if v287 == int32(0) {
		goto L73
	} else {
		goto L80
	}
L80:
	;
	base.MemoryCopy(m, v282, v12+int32(448), v287)
	goto L73
L81:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2572))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v326
	F_errmsg_internal(m, int32(_a_F_lookup_ts_config_cache_1), v12)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L3
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_config_cache_2), int32(442), int32(_a_F_lookup_ts_config_cache_3))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
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
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v12)+2572))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v340
	F_errmsg_internal(m, int32(_a_F_lookup_ts_config_cache_4), v12+int32(16))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L3
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_config_cache_2), int32(449), int32(_a_F_lookup_ts_config_cache_3))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v191
	F_errmsg_internal(m, int32(_a_F_lookup_ts_config_cache_5), v12+int32(32))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L3
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_config_cache_2), int32(507), int32(_a_F_lookup_ts_config_cache_3))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
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
	v374 = m.ExcPending
	if v374 != 0 {
		goto L3
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_config_cache_2), int32(509), int32(_a_F_lookup_ts_config_cache_3))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
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
	v387 = m.ExcPending
	if v387 != 0 {
		goto L3
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_lookup_ts_config_cache_2), int32(530), int32(_a_F_lookup_ts_config_cache_3))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
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
func F_ts_process_call(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
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
	var v126 int64
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
	var v153 int64
	_ = v153
	v10 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16+v17<<(uint(int32(2))%32))))
	if v21 == int32(0) {
		v153 = v10
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
	v153 = v10
	goto L1
L5:
	;
	if v26 == int32(0) {
		v153 = v10
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
	return int64(0)
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = l2
	v13 = int32(_a_F_ts_setup_firstcall_0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_ts_setup_firstcall[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_ts_setup_firstcall[0])) = v16
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v22 = F_palloc0_mul(m, int32(4), v19+int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v22
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
		if v27 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v22))) = v27
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
			if v29 == int32(0) {
			} else {
				v38 = v27 + int32(8)
				for {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
					v43 = v41 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v43
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
					*(*int32)(unsafe.Add(mBase, uint32(v45+v43<<(uint(int32(2))%32)))) = v49
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
					if v53 != 0 {
						v38 = v49 + int32(8)
						continue
					} else {
						break
					}
					break
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(0)
		}
		v66 = F_get_call_result_type(m, l0, int32(0), v10+int32(12))
		mBase = m.M
		v67 = m.ExcPending
		if v67 != 0 {
			return
		} else {
			if v66 == int32(1) {
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v70
				v72 = F_TupleDescGetAttInMetadata(m, v70)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v72
					*(*int32)(unsafe.Add(mBase, _c_F_ts_setup_firstcall[0])) = v14
					m.G0 = v10 + int32(16)
					return
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_ts_setup_firstcall_1), int32(0))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ts_setup_firstcall_2), int32(2485), int32(_a_F_ts_setup_firstcall_3))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
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
