package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_pg_stat_get_backend_idset(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
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
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	if v6 == int32(0) {
		v9 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
			v15 = F_MemoryContextAlloc(m, v13, int32(4))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v15
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(0)
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v26 = v24 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = v26
				v28 = F_pgstat_fetch_stat_numbackends(m)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					if v26 <= v28 {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
						v32 = F_pgstat_get_local_beentry_by_index(m, v31)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							v34 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
							*(*int64)(unsafe.Add(mBase, uint32(v22))) = v34 + int64(1)
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = int32(1)
							v41 = int64(*(*int32)(unsafe.Add(mBase, uint32(v32)+408)))
							return v41
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = int32(2)
							v48 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v48)
							return int64(0)
						}
					}
				}
			}
		}
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		v26 = v24 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v23))) = v26
		v28 = F_pgstat_fetch_stat_numbackends(m)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int64(0)
		} else {
			if v26 <= v28 {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v32 = F_pgstat_get_local_beentry_by_index(m, v31)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int64(0)
				} else {
					v34 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
					*(*int64)(unsafe.Add(mBase, uint32(v22))) = v34 + int64(1)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = int32(1)
					v41 = int64(*(*int32)(unsafe.Add(mBase, uint32(v32)+408)))
					return v41
				}
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = int32(2)
					v48 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v48)
					return int64(0)
				}
			}
		}
	}
}
func F_pg_stat_get_backend_xact_start(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pgstat_get_beentry_by_proc_number(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 == int32(0) {
			v29 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
			v32 = int64(0)
			return v32
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_xact_start[0]))
			v14 = F_has_privs_of_role(m, v12, int32(3375))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				if v14 == int32(0) {
					v19 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_xact_start[0]))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+52))
					v21 = F_has_privs_of_role(m, v19, v20)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						if v21 == int32(0) {
							v29 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
							v32 = int64(0)
						} else {
							v25 = *(*int64)(unsafe.Add(mBase, uint32(v5)+24))
							if v25 != int64(0) {
								v32 = v25
							} else {
								v29 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
								v32 = int64(0)
							}
						}
						return v32
					}
				} else {
					v25 = *(*int64)(unsafe.Add(mBase, uint32(v5)+24))
					if v25 != int64(0) {
						v32 = v25
					} else {
						v29 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
						v32 = int64(0)
					}
					return v32
				}
			}
		}
	}
}
func F_pg_stat_get_bgwriter_buf_written_clean(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	v2 = F_pgstat_fetch_stat_bgwriter(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
		return v6
	}
}
func F_pg_stat_get_bgwriter_maxwritten_clean(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	v2 = F_pgstat_fetch_stat_bgwriter(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+8))
		return v6
	}
}
func F_pg_stat_get_db_blk_write_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+176))
			return base.I64_reinterpret_f64(base.F64_div(base.F64_convert_i64_s(v11), float64(1000)))
		}
	}
}
func F_pg_stat_get_db_checksum_failures(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int64
	_ = v24
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_db_checksum_failures[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+268))
	if base.B2i32(v6 != int32(0)) == int32(0) {
		v11 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
		return int64(0)
	} else {
		v16 = F_pgstat_fetch_stat_dbentry(m, base.I32_wrap_i64(v3))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			if v16 == int32(0) {
				return int64(0)
			} else {
				v24 = *(*int64)(unsafe.Add(mBase, uint32(v16)+152))
				return v24
			}
		}
	}
}
func F_pg_stat_get_db_conflict_snapshot(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+96))
			return v11
		}
	}
}
func F_pg_stat_get_db_idle_in_transaction_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+208))
			return base.I64_reinterpret_f64(base.F64_div(base.F64_convert_i64_s(v11), float64(1000)))
		}
	}
}
func F_pg_stat_get_db_parallel_workers_to_launch(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+240))
			return v11
		}
	}
}
func F_pg_stat_get_db_sessions(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+184))
			return v11
		}
	}
}
func F_pg_stat_get_db_tuples_deleted(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+64))
			return v11
		}
	}
}
func F_pg_stat_get_dead_tuples(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+80))
			return v11
		}
	}
}
func F_pg_stat_get_progress_info(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v313 int64
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int64
	_ = v323
	var v325 int64
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int64
	_ = v341
	var v343 int64
	_ = v343
	var v345 int64
	_ = v345
	var v347 int64
	_ = v347
	var v349 int64
	_ = v349
	var v351 int64
	_ = v351
	var v353 int64
	_ = v353
	var v355 int64
	_ = v355
	var v357 int64
	_ = v357
	var v359 int64
	_ = v359
	var v361 int64
	_ = v361
	var v363 int64
	_ = v363
	var v365 int64
	_ = v365
	var v367 int64
	_ = v367
	var v369 int64
	_ = v369
	var v371 int64
	_ = v371
	var v373 int64
	_ = v373
	var v375 int64
	_ = v375
	var v377 int64
	_ = v377
	var v379 int64
	_ = v379
	var v381 int64
	_ = v381
	var v383 int64
	_ = v383
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	v8 = m.G0
	v10 = v8 - int32(240)
	m.G0 = v10
	v12 = F_pgstat_fetch_stat_numbackends(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = F_text_to_cstring(m, v17)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v26 = v19
	v27 = int32(_a_F_pg_stat_get_progress_info_0)
	goto L8
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L119
	}
L6:
	;
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L99
	}
L7:
	;
	if v64 == int32(0) {
		v290 = int32(1)
		goto L6
	} else {
		goto L20
	}
L8:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v30 == v31 {
		v53 = v30
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v64 = int32(0)
	goto L7
L10:
	;
	v55 = int32(1)
	if v53 != 0 {
		v26 = v26 + v55
		v27 = v27 + v55
		goto L8
	} else {
		goto L19
	}
L11:
	;
	if base.Ui32((v30-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v41 = v30 | int32(32)
	goto L14
L13:
	;
	v41 = v30
	goto L14
L14:
	;
	if base.Ui32((v31-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v50 = v31 | int32(32)
	goto L17
L16:
	;
	v50 = v31
	goto L17
L17:
	;
	if v41 == v50 {
		v53 = v41
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v64 = v41 - v50
	goto L7
L19:
	;
	goto L9
L20:
	;
	v70 = v19
	v71 = int32(_a_F_pg_stat_get_progress_info_1)
	goto L22
L21:
	;
	if v108 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v74 == v75 {
		v97 = v74
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v108 = int32(0)
	goto L21
L24:
	;
	v99 = int32(1)
	if v97 != 0 {
		v70 = v70 + v99
		v71 = v71 + v99
		goto L22
	} else {
		goto L33
	}
L25:
	;
	if base.Ui32((v74-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v85 = v74 | int32(32)
	goto L28
L27:
	;
	v85 = v74
	goto L28
L28:
	;
	if base.Ui32((v75-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v94 = v75 | int32(32)
	goto L31
L30:
	;
	v94 = v75
	goto L31
L31:
	;
	if v85 == v94 {
		v97 = v85
		goto L24
	} else {
		goto L32
	}
L32:
	;
	v108 = v85 - v94
	goto L21
L33:
	;
	goto L23
L34:
	;
	v290 = int32(2)
	goto L6
L35:
	;
	goto L36
L36:
	;
	v115 = v19
	v116 = int32(_a_F_pg_stat_get_progress_info_2)
	goto L38
L37:
	;
	if v153 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v119 == v120 {
		v142 = v119
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v153 = int32(0)
	goto L37
L40:
	;
	v144 = int32(1)
	if v142 != 0 {
		v115 = v115 + v144
		v116 = v116 + v144
		goto L38
	} else {
		goto L49
	}
L41:
	;
	if base.Ui32((v119-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v130 = v119 | int32(32)
	goto L44
L43:
	;
	v130 = v119
	goto L44
L44:
	;
	if base.Ui32((v120-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v139 = v120 | int32(32)
	goto L47
L46:
	;
	v139 = v120
	goto L47
L47:
	;
	if v130 == v139 {
		v142 = v130
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v153 = v130 - v139
	goto L37
L49:
	;
	goto L39
L50:
	;
	v290 = int32(6)
	goto L6
L51:
	;
	goto L52
L52:
	;
	v160 = v19
	v161 = int32(_a_F_pg_stat_get_progress_info_3)
	goto L54
L53:
	;
	if v198 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L54:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v164 == v165 {
		v187 = v164
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v198 = int32(0)
	goto L53
L56:
	;
	v189 = int32(1)
	if v187 != 0 {
		v160 = v160 + v189
		v161 = v161 + v189
		goto L54
	} else {
		goto L65
	}
L57:
	;
	if base.Ui32((v164-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v175 = v164 | int32(32)
	goto L60
L59:
	;
	v175 = v164
	goto L60
L60:
	;
	if base.Ui32((v165-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v184 = v165 | int32(32)
	goto L63
L62:
	;
	v184 = v165
	goto L63
L63:
	;
	if v175 == v184 {
		v187 = v175
		goto L56
	} else {
		goto L64
	}
L64:
	;
	v198 = v175 - v184
	goto L53
L65:
	;
	goto L55
L66:
	;
	v290 = int32(3)
	goto L6
L67:
	;
	goto L68
L68:
	;
	v205 = v19
	v206 = int32(_a_F_pg_stat_get_progress_info_4)
	goto L70
L69:
	;
	if v243 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L70:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if v209 == v210 {
		v232 = v209
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v243 = int32(0)
	goto L69
L72:
	;
	v234 = int32(1)
	if v232 != 0 {
		v205 = v205 + v234
		v206 = v206 + v234
		goto L70
	} else {
		goto L81
	}
L73:
	;
	if base.Ui32((v209-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v220 = v209 | int32(32)
	goto L76
L75:
	;
	v220 = v209
	goto L76
L76:
	;
	if base.Ui32((v210-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v229 = v210 | int32(32)
	goto L79
L78:
	;
	v229 = v210
	goto L79
L79:
	;
	if v220 == v229 {
		v232 = v220
		goto L72
	} else {
		goto L80
	}
L80:
	;
	v243 = v220 - v229
	goto L69
L81:
	;
	goto L71
L82:
	;
	v290 = int32(4)
	goto L6
L83:
	;
	goto L84
L84:
	;
	v250 = v19
	v251 = int32(_a_F_pg_stat_get_progress_info_5)
	goto L86
L85:
	;
	if v288 != 0 {
		goto L5
	} else {
		goto L98
	}
L86:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if v254 == v255 {
		v277 = v254
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v288 = int32(0)
	goto L85
L88:
	;
	v279 = int32(1)
	if v277 != 0 {
		v250 = v250 + v279
		v251 = v251 + v279
		goto L86
	} else {
		goto L97
	}
L89:
	;
	if base.Ui32((v254-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v265 = v254 | int32(32)
	goto L92
L91:
	;
	v265 = v254
	goto L92
L92:
	;
	if base.Ui32((v255-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v274 = v255 | int32(32)
	goto L95
L94:
	;
	v274 = v255
	goto L95
L95:
	;
	if v265 == v274 {
		v277 = v265
		goto L88
	} else {
		goto L96
	}
L96:
	;
	v288 = v265 - v274
	goto L85
L97:
	;
	goto L87
L98:
	;
	v290 = int32(5)
	goto L6
L99:
	;
	if int32(0) < v12 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v299 = v10 + int32(16) | int32(2)
	v305 = int32(1)
	goto L103
L101:
	;
	goto L102
L102:
	;
	m.G0 = v10 + int32(240)
	return int64(0)
L103:
	;
	base.MemoryFill(m, v10+int32(48), int32(0), int32(184))
	v313 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+31)) = v313
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v313
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v313
	v319 = F_pgstat_get_local_beentry_by_index(m, v305)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L105
	}
L104:
	;
	goto L102
L105:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v319)+220))
	if v290 == v321 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v323 = int64(*(*int32)(unsafe.Add(mBase, uint32(v319)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v323
	v325 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v319)+48)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v325
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_progress_info[0]))
	v330 = F_has_privs_of_role(m, v328, int32(3375))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L111
	}
L107:
	;
	goto L108
L108:
	;
	v398 = v305 + int32(1)
	if v398 <= v12 {
		v305 = v398
		goto L103
	} else {
		goto L118
	}
L109:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	F_tuplestore_putvalues(m, v389, v390, v10+int32(48), v10+int32(16))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L117
	}
L110:
	;
	v383 = int64(72340172838076673)
	*(*int64)(unsafe.Add(mBase, uint32(v299)+13)) = v383
	*(*int64)(unsafe.Add(mBase, uint32(v299)+8)) = v383
	*(*int64)(unsafe.Add(mBase, uint32(v299))) = v383
	goto L109
L111:
	;
	if v330 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_progress_info[0]))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v319)+52))
	v337 = F_has_privs_of_role(m, v335, v336)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v341 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v319)+224)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = v341
	v343 = *(*int64)(unsafe.Add(mBase, uint32(v319)+232))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v343
	v345 = *(*int64)(unsafe.Add(mBase, uint32(v319)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = v345
	v347 = *(*int64)(unsafe.Add(mBase, uint32(v319)+248))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+88)) = v347
	v349 = *(*int64)(unsafe.Add(mBase, uint32(v319)+256))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+96)) = v349
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v319)+264))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = v351
	v353 = *(*int64)(unsafe.Add(mBase, uint32(v319)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+112)) = v353
	v355 = *(*int64)(unsafe.Add(mBase, uint32(v319)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+120)) = v355
	v357 = *(*int64)(unsafe.Add(mBase, uint32(v319)+288))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+128)) = v357
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v319)+296))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+136)) = v359
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v319)+304))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+144)) = v361
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v319)+312))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+152)) = v363
	v365 = *(*int64)(unsafe.Add(mBase, uint32(v319)+320))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+160)) = v365
	v367 = *(*int64)(unsafe.Add(mBase, uint32(v319)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+168)) = v367
	v369 = *(*int64)(unsafe.Add(mBase, uint32(v319)+336))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+176)) = v369
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v319)+344))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+184)) = v371
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v319)+352))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+192)) = v373
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v319)+360))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+200)) = v375
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v319)+368))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+208)) = v377
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v319)+376))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+216)) = v379
	v381 = *(*int64)(unsafe.Add(mBase, uint32(v319)+384))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+224)) = v381
	goto L109
L115:
	;
	if v337 == int32(0) {
		goto L110
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	goto L108
L118:
	;
	goto L104
L119:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v19
	F_errmsg(m, int32(_a_F_pg_stat_get_progress_info_6), v10)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_pg_stat_get_progress_info_7), int32(303), int32(_a_F_pg_stat_get_progress_info_8))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_stat_get_slru(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(96)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	F_pgstat_snapshot_fixed(m, int32(12))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = int32(0)
	base.MemoryFill(m, v16+int32(16), v29, int32(72))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)) = uint8(v29)
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(0)
	goto L5
L4:
	;
	if v41 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_slru[0]))
	goto L7
L7:
	;
	goto L4
L8:
	;
	v47 = v41
	v48 = v2
	goto L11
L9:
	;
	goto L10
L10:
	;
	m.G0 = v16 + int32(96)
	return int64(0)
L11:
	;
	v58 = v48 << (uint(int32(6)) % 32)
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v58)+uint32(_c_F_pg_stat_get_slru[1])))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v58)+uint32(_c_F_pg_stat_get_slru[2])))
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v58)+uint32(_c_F_pg_stat_get_slru[3])))
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v58)+uint32(_c_F_pg_stat_get_slru[4])))
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v58)+uint32(_c_F_pg_stat_get_slru[5])))
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v58)+uint32(_c_F_pg_stat_get_slru[6])))
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v58)+uint32(_c_F_pg_stat_get_slru[7])))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v58)+uint32(_c_F_pg_stat_get_slru[8])))
	v69 = F_cstring_to_text(m, v47)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v68
	*(*int64)(unsafe.Add(mBase, uint32(v16)+72)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = v66
	*(*int64)(unsafe.Add(mBase, uint32(v16)+56)) = v65
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v62
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = base.I64_extend_i32_u(v69)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v84 = v16 + int32(16)
	F_tuplestore_putvalues(m, v81, v82, v84, v16)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v87 = int32(0)
	base.MemoryFill(m, v84, v87, int32(72))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)) = uint8(v87)
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(0)
	v95 = v48 + int32(1)
	if base.Ui32(v95) <= base.Ui32(int32(7)) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v102 != 0 {
		v47 = v102
		v48 = v95
		goto L11
	} else {
		goto L19
	}
L16:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v95<<(uint(int32(2))%32))+uint32(_c_F_pg_stat_get_slru[0])))
	v102 = v100
	goto L18
L17:
	;
	v102 = int32(0)
	goto L18
L18:
	;
	goto L15
L19:
	;
	goto L12
}
func F_pg_stat_get_stat_reset_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pgstat_fetch_stat_tabentry(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 != 0 {
			v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)+216))
			if v9 != int64(0) {
				v16 = v9
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				v16 = int64(0)
			}
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			v16 = int64(0)
		}
		return v16
	}
}
func F_pg_stat_get_wal_receiver(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
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
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
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
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v432 int64
	_ = v432
	var v435 int64
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int64
	_ = v563
	var v564 int32
	_ = v564
	var v570 int64
	_ = v570
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	v17 = m.G0
	v19 = v17 - int32(1360)
	m.G0 = v19
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_receiver[0]))
	v25 = base.AtomicRmwXchg32(m, v22, int32(1456), int32(1))
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, v22+int32(1456), int32(_a_F_pg_stat_get_wal_receiver_0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_receiver[0]))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v34)+96))
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v34)+88))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v34)+80))
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v34)+72))
	v39 = int64(*(*int32)(unsafe.Add(mBase, uint32(v34)+56)))
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v34)+48))
	v41 = int64(*(*int32)(unsafe.Add(mBase, uint32(v34)+40)))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v34)+32))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1453)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v47 = v19 + int32(1024)
	v49 = v34 + int32(1388)
	goto L9
L4:
	;
	return int64(0)
L5:
	;
	goto L3
L6:
	;
	v170 = v19 + int32(1088)
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_receiver[0]))
	v174 = v172 + int32(1128)
	goto L40
L7:
	;
	v166 = F_strlen(m, v155)
	mBase = m.M
	goto L6
L9:
	;
	goto L10
L10:
	;
	v56 = int32(63)
	if (v47^v49)&int32(3) != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v159 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v159)
	goto L7
L12:
	;
	v140 = v135
	v141 = v136
	v142 = v137
	goto L33
L13:
	;
	if v130 == int32(0) {
		v155 = v128
		v156 = v129
		goto L11
	} else {
		goto L32
	}
L14:
	;
	v128 = v49
	v129 = v47
	v130 = v56
	goto L13
L15:
	;
	goto L16
L16:
	;
	v60 = int32(0)
	if base.B2i32(v49&int32(3) == v60)|int32(0) == v60 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v96 == int32(0) {
		v155 = v93
		v156 = v94
		goto L11
	} else {
		goto L26
	}
L18:
	;
	v72 = v49
	v73 = v47
	v74 = v56
	goto L21
L19:
	;
	goto L20
L20:
	;
	v93 = v49
	v94 = v47
	v95 = v56
	v96 = int32(1)
	goto L17
L21:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v76)
	if v76 == int32(0) {
		v135 = v72
		v136 = v73
		v137 = v74
		goto L12
	} else {
		goto L23
	}
L22:
	;
	v93 = v87
	v94 = v81
	v95 = v83
	v96 = v85
	goto L17
L23:
	;
	v80 = int32(1)
	v81 = v73 + v80
	v83 = v74 - v80
	v84 = int32(0)
	v85 = base.B2i32(v83 != v84)
	v87 = v72 + v80
	if v87&int32(3) == v84 {
		v93 = v87
		v94 = v81
		v95 = v83
		v96 = v85
		goto L17
	} else {
		goto L24
	}
L24:
	;
	if v83 != 0 {
		v72 = v87
		v73 = v81
		v74 = v83
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if base.B2i32(v99 == int32(0))|base.B2i32(base.Ui32(v95) < base.Ui32(int32(4))) != 0 {
		v128 = v93
		v129 = v94
		v130 = v95
		goto L13
	} else {
		goto L27
	}
L27:
	;
	v106 = v93
	v107 = v94
	v108 = v95
	goto L28
L28:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v114 = int32(-2139062144)
	if (int32(16843008)-v111|v111)&v114 != v114 {
		v135 = v106
		v136 = v107
		v137 = v108
		goto L12
	} else {
		goto L30
	}
L29:
	;
	v128 = v122
	v129 = v120
	v130 = v124
	goto L13
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v111
	v119 = int32(4)
	v120 = v107 + v119
	v122 = v106 + v119
	v124 = v108 - v119
	if base.Ui32(int32(3)) < base.Ui32(v124) {
		v106 = v122
		v107 = v120
		v108 = v124
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v135 = v128
	v136 = v129
	v137 = v130
	goto L12
L33:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v144)
	if v144 == int32(0) {
		v155 = v140
		v156 = v141
		goto L11
	} else {
		goto L35
	}
L34:
	;
	v155 = v151
	v156 = v149
	goto L11
L35:
	;
	v148 = int32(1)
	v149 = v141 + v148
	v151 = v140 + v148
	v153 = v142 - v148
	if v153 != 0 {
		v140 = v151
		v141 = v149
		v142 = v153
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_receiver[0]))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)+1384))
	v298 = v295 + int32(104)
	goto L71
L38:
	;
	v291 = F_strlen(m, v280)
	mBase = m.M
	goto L37
L40:
	;
	goto L41
L41:
	;
	v181 = int32(254)
	if (v170^v174)&int32(3) != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v284 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v281))) = uint8(v284)
	goto L38
L43:
	;
	v265 = v260
	v266 = v261
	v267 = v262
	goto L64
L44:
	;
	if v255 == int32(0) {
		v280 = v253
		v281 = v254
		goto L42
	} else {
		goto L63
	}
L45:
	;
	v253 = v174
	v254 = v170
	v255 = v181
	goto L44
L46:
	;
	goto L47
L47:
	;
	v185 = int32(0)
	if base.B2i32(v174&int32(3) == v185)|int32(0) == v185 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v221 == int32(0) {
		v280 = v218
		v281 = v219
		goto L42
	} else {
		goto L57
	}
L49:
	;
	v197 = v174
	v198 = v170
	v199 = v181
	goto L52
L50:
	;
	goto L51
L51:
	;
	v218 = v174
	v219 = v170
	v220 = v181
	v221 = int32(1)
	goto L48
L52:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	*(*uint8)(unsafe.Add(mBase, uint32(v198))) = uint8(v201)
	if v201 == int32(0) {
		v260 = v197
		v261 = v198
		v262 = v199
		goto L43
	} else {
		goto L54
	}
L53:
	;
	v218 = v212
	v219 = v206
	v220 = v208
	v221 = v210
	goto L48
L54:
	;
	v205 = int32(1)
	v206 = v198 + v205
	v208 = v199 - v205
	v209 = int32(0)
	v210 = base.B2i32(v208 != v209)
	v212 = v197 + v205
	if v212&int32(3) == v209 {
		v218 = v212
		v219 = v206
		v220 = v208
		v221 = v210
		goto L48
	} else {
		goto L55
	}
L55:
	;
	if v208 != 0 {
		v197 = v212
		v198 = v206
		v199 = v208
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	if base.B2i32(v224 == int32(0))|base.B2i32(base.Ui32(v220) < base.Ui32(int32(4))) != 0 {
		v253 = v218
		v254 = v219
		v255 = v220
		goto L44
	} else {
		goto L58
	}
L58:
	;
	v231 = v218
	v232 = v219
	v233 = v220
	goto L59
L59:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v239 = int32(-2139062144)
	if (int32(16843008)-v236|v236)&v239 != v239 {
		v260 = v231
		v261 = v232
		v262 = v233
		goto L43
	} else {
		goto L61
	}
L60:
	;
	v253 = v247
	v254 = v245
	v255 = v249
	goto L44
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232))) = v236
	v244 = int32(4)
	v245 = v232 + v244
	v247 = v231 + v244
	v249 = v233 - v244
	if base.Ui32(int32(3)) < base.Ui32(v249) {
		v231 = v247
		v232 = v245
		v233 = v249
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v260 = v253
	v261 = v254
	v262 = v255
	goto L43
L64:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265))))
	*(*uint8)(unsafe.Add(mBase, uint32(v266))) = uint8(v269)
	if v269 == int32(0) {
		v280 = v265
		v281 = v266
		goto L42
	} else {
		goto L66
	}
L65:
	;
	v280 = v276
	v281 = v274
	goto L42
L66:
	;
	v273 = int32(1)
	v274 = v266 + v273
	v276 = v265 + v273
	v278 = v267 - v273
	if v278 != 0 {
		v265 = v276
		v266 = v274
		v267 = v278
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_receiver[0]))
	v420 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v419)+1456)), uint32(v420))
	if v44&int32(1) != 0 {
		goto L101
	} else {
		goto L102
	}
L69:
	;
	v415 = F_strlen(m, v404)
	mBase = m.M
	goto L68
L71:
	;
	goto L72
L72:
	;
	v305 = int32(1023)
	if (v19^v298)&int32(3) != 0 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v408 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v405))) = uint8(v408)
	goto L69
L74:
	;
	v389 = v384
	v390 = v385
	v391 = v386
	goto L95
L75:
	;
	if v379 == int32(0) {
		v404 = v377
		v405 = v378
		goto L73
	} else {
		goto L94
	}
L76:
	;
	v377 = v298
	v378 = v19
	v379 = v305
	goto L75
L77:
	;
	goto L78
L78:
	;
	v309 = int32(0)
	if base.B2i32(v298&int32(3) == v309)|int32(0) == v309 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v345 == int32(0) {
		v404 = v342
		v405 = v343
		goto L73
	} else {
		goto L88
	}
L80:
	;
	v321 = v298
	v322 = v19
	v323 = v305
	goto L83
L81:
	;
	goto L82
L82:
	;
	v342 = v298
	v343 = v19
	v344 = v305
	v345 = int32(1)
	goto L79
L83:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	*(*uint8)(unsafe.Add(mBase, uint32(v322))) = uint8(v325)
	if v325 == int32(0) {
		v384 = v321
		v385 = v322
		v386 = v323
		goto L74
	} else {
		goto L85
	}
L84:
	;
	v342 = v336
	v343 = v330
	v344 = v332
	v345 = v334
	goto L79
L85:
	;
	v329 = int32(1)
	v330 = v322 + v329
	v332 = v323 - v329
	v333 = int32(0)
	v334 = base.B2i32(v332 != v333)
	v336 = v321 + v329
	if v336&int32(3) == v333 {
		v342 = v336
		v343 = v330
		v344 = v332
		v345 = v334
		goto L79
	} else {
		goto L86
	}
L86:
	;
	if v332 != 0 {
		v321 = v336
		v322 = v330
		v323 = v332
		goto L83
	} else {
		goto L87
	}
L87:
	;
	goto L84
L88:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342))))
	if base.B2i32(v348 == int32(0))|base.B2i32(base.Ui32(v344) < base.Ui32(int32(4))) != 0 {
		v377 = v342
		v378 = v343
		v379 = v344
		goto L75
	} else {
		goto L89
	}
L89:
	;
	v355 = v342
	v356 = v343
	v357 = v344
	goto L90
L90:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v363 = int32(-2139062144)
	if (int32(16843008)-v360|v360)&v363 != v363 {
		v384 = v355
		v385 = v356
		v386 = v357
		goto L74
	} else {
		goto L92
	}
L91:
	;
	v377 = v371
	v378 = v369
	v379 = v373
	goto L75
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v356))) = v360
	v368 = int32(4)
	v369 = v356 + v368
	v371 = v355 + v368
	v373 = v357 - v368
	if base.Ui32(int32(3)) < base.Ui32(v373) {
		v355 = v371
		v356 = v369
		v357 = v373
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	v384 = v377
	v385 = v378
	v386 = v379
	goto L74
L95:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389))))
	*(*uint8)(unsafe.Add(mBase, uint32(v390))) = uint8(v393)
	if v393 == int32(0) {
		v404 = v389
		v405 = v390
		goto L73
	} else {
		goto L97
	}
L96:
	;
	v404 = v400
	v405 = v398
	goto L73
L97:
	;
	v397 = int32(1)
	v398 = v390 + v397
	v400 = v389 + v397
	v402 = v391 - v397
	if v402 != 0 {
		v389 = v400
		v390 = v398
		v391 = v402
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L4
	} else {
		goto L169
	}
L100:
	;
	m.G0 = v19 + int32(1360)
	return v570
L101:
	;
	v426 = v45
	goto L103
L102:
	;
	v426 = v420
	goto L103
L103:
	;
	if v426 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v429 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v429)
	v570 = int64(0)
	goto L100
L105:
	;
	goto L106
L106:
	;
	v432 = int64(0)
	v435 = base.AtomicRmwCmpxchg64(m, v419, int32(1464), v432, v432)
	v439 = F_get_call_result_type(m, l0, int32(0), v19+int32(1356))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	if v439 != int32(1) {
		goto L99
	} else {
		goto L108
	}
L108:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1356))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v444)))
	v446 = F_palloc0_mul(m, int32(8), v445)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1356))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	v451 = F_palloc0_mul(m, int32(1), v450)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v446))) = base.I64_extend_i32_s(v45)
	v456 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_wal_receiver[1]))
	v458 = F_has_privs_of_role(m, v456, int32(3375))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L112
	}
L111:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1356))
	v560 = F_heap_form_tuple(m, v559, v446, v451)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L4
	} else {
		goto L167
	}
L112:
	;
	if v458 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1356))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	v465 = v463 - int32(1)
	if v465 == int32(0) {
		goto L111
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if base.Ui32(v43) <= base.Ui32(int32(6)) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v468 = int32(1)
	base.MemoryFill(m, v451+v468, v468, v465)
	goto L111
L117:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v43<<(uint(int32(2))%32))+uint32(_c_F_pg_stat_get_wal_receiver[2])))
	v480 = v478
	goto L119
L118:
	;
	v480 = int32(_a_F_pg_stat_get_wal_receiver_1)
	goto L119
L119:
	;
	v481 = F_cstring_to_text(m, v480)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v446)+8)) = base.I64_extend_i32_u(v481)
	if v42 == int64(0) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v446)+24)) = v41
	if v435 == int64(0) {
		goto L126
	} else {
		goto L127
	}
L122:
	;
	v487 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+2)) = uint8(v487)
	goto L121
L123:
	;
	goto L124
L124:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v446)+16)) = v42
	goto L121
L125:
	;
	if v40 == int64(0) {
		goto L130
	} else {
		goto L131
	}
L126:
	;
	v491 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+4)) = uint8(v491)
	goto L125
L127:
	;
	goto L128
L128:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v446)+32)) = v435
	goto L125
L129:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v446)+48)) = v39
	if v38 == int64(0) {
		goto L134
	} else {
		goto L135
	}
L130:
	;
	v496 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+5)) = uint8(v496)
	goto L129
L131:
	;
	goto L132
L132:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v446)+40)) = v40
	goto L129
L133:
	;
	if v37 == int64(0) {
		goto L138
	} else {
		goto L139
	}
L134:
	;
	v502 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+7)) = uint8(v502)
	goto L133
L135:
	;
	goto L136
L136:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v446)+56)) = v38
	goto L133
L137:
	;
	if v36 == int64(0) {
		goto L142
	} else {
		goto L143
	}
L138:
	;
	v507 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+8)) = uint8(v507)
	goto L137
L139:
	;
	goto L140
L140:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v446)+64)) = v37
	goto L137
L141:
	;
	if v35 == int64(0) {
		goto L146
	} else {
		goto L147
	}
L142:
	;
	v512 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+9)) = uint8(v512)
	goto L141
L143:
	;
	goto L144
L144:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v446)+72)) = v36
	goto L141
L145:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1024)))
	if v520 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L146:
	;
	v517 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+10)) = uint8(v517)
	goto L145
L147:
	;
	goto L148
L148:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v446)+80)) = v35
	goto L145
L149:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1088)))
	if v531 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L150:
	;
	v523 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+11)) = uint8(v523)
	goto L149
L151:
	;
	goto L152
L152:
	;
	v527 = F_cstring_to_text(m, v19+int32(1024))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v446)+88)) = base.I64_extend_i32_u(v527)
	goto L149
L154:
	;
	if v296 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L155:
	;
	v534 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+12)) = uint8(v534)
	goto L154
L156:
	;
	goto L157
L157:
	;
	v538 = F_cstring_to_text(m, v19+int32(1088))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v446)+96)) = base.I64_extend_i32_u(v538)
	goto L154
L159:
	;
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v548 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v544 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+13)) = uint8(v544)
	goto L159
L161:
	;
	goto L162
L162:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v446)+104)) = base.I64_extend_i32_s(v296)
	goto L159
L163:
	;
	v551 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+14)) = uint8(v551)
	goto L111
L164:
	;
	goto L165
L165:
	;
	v553 = F_cstring_to_text(m, v19)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v446)+112)) = base.I64_extend_i32_u(v553)
	goto L111
L167:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v560)+16))
	v563 = F_HeapTupleHeaderGetDatum(m, v562)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	v570 = v563
	goto L100
L169:
	;
	F_errmsg_internal(m, int32(_a_F_pg_stat_get_wal_receiver_2), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L4
	} else {
		goto L170
	}
L170:
	;
	F_errfinish(m, int32(_a_F_pg_stat_get_wal_receiver_3), int32(1581), int32(_a_F_pg_stat_get_wal_receiver_4))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L4
	} else {
		goto L171
	}
L171:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_stat_reset_replication_slot(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v24 int32
	_ = v24
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v4 == int32(1) {
		F_pgstat_reset_of_kind(m, int32(4))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			return int64(0)
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
				v19 = m.G0
				v21 = v19 - int32(16)
				m.G0 = v21
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_replication_slot[0]))
				v28 = F_LWLockAcquire(m, v24+int32(_a_F_pg_stat_reset_replication_slot_0), int32(1))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v31 = F_SearchNamedReplicationSlot(m, v17, int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int64(0)
					} else {
						if v31 != 0 {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+88))
							if v33 != 0 {
								v37 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_replication_slot[1]))
								v40 = base.I32_div_s(v31-v37, int32(296))
								F_pgstat_reset(m, int32(4), int32(0), base.I64_extend_i32_s(v40))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int64(0)
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_replication_slot[0]))
									F_LWLockRelease(m, v45+int32(_a_F_pg_stat_reset_replication_slot_0))
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int64(0)
									} else {
										m.G0 = v21 + int32(16)
										return int64(0)
									}
								}
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_replication_slot[0]))
								F_LWLockRelease(m, v45+int32(_a_F_pg_stat_reset_replication_slot_0))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int64(0)
								} else {
									m.G0 = v21 + int32(16)
									return int64(0)
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v21))) = v17
									F_errmsg(m, int32(_a_F_pg_stat_reset_replication_slot_1), v21)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_pg_stat_reset_replication_slot_2), int32(56), int32(_a_F_pg_stat_reset_replication_slot_3))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
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
	}
}
func F_pg_stat_reset_subscription_stats(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v8 == int32(1) {
		F_pgstat_reset_of_kind(m, int32(5))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			m.G0 = v6 + int32(16)
			return int64(0)
		}
	} else {
		v16 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
		if v16 == int64(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(0)
					F_errmsg(m, int32(_a_F_pg_stat_reset_subscription_stats_0), v6)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_stat_reset_subscription_stats_1), int32(2111), int32(_a_F_pg_stat_reset_subscription_stats_2))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			F_pgstat_reset(m, int32(5), int32(0), v16)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				m.G0 = v6 + int32(16)
				return int64(0)
			}
		}
	}
}
func F_pg_stat_statements_1_2(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_pg_stat_statements_internal(m, l0, int32(2), base.B2i32(v3 != int64(0)))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_pg_stat_statements_reset_1_11(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v8 = F_entry_reset(m, v2, v3, v4, base.B2i32(v5 != int64(0)))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		return v8
	}
}
