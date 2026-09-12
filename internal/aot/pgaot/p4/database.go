package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetDatabasePath(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	switch l1 - int32(1663) {
	case 0:
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l0
		v19 = F_psprintf(m, int32(39137), v6+int32(16))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v30 = v19
			m.G0 = v6 + int32(32)
			return v30
		}
	case 1:
		v11 = F_pstrdup(m, int32(314244))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v30 = v11
			m.G0 = v6 + int32(32)
			return v30
		}
	default:
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(561441)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(488584)
		v28 = F_psprintf(m, int32(39103), v6)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = v28
			m.G0 = v6 + int32(32)
			return v30
		}
	}
}
func F_has_database_privilege_id_id(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v18)
		v22 = F_convert_any_priv_string(m, v14, int32(1656912))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v26 = F_object_aclcheck_ext(m, int32(1262), v11, v12, v22, v9+int32(15))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
				if v28 == int32(1) {
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
					v35 = int32(0)
				} else {
					v35 = base.B2i32(v26 == int32(0))
				}
				m.G0 = v9 + int32(16)
				return v35
			}
		}
	}
}
func F_has_database_privilege_id_name(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v14 = F_text_to_cstring(m, v6)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v17 = F_get_database_oid(m, v14, int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v20 = F_convert_any_priv_string(m, v11, int32(1656912))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = F_object_aclcheck(m, int32(1262), v17, v4, v20)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v22 == int32(0))
						}
					}
				}
			}
		}
	}
}
func F_has_database_privilege_name_name(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_get_role_oid_or_public(m, v4)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v16 = F_text_to_cstring(m, v6)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v19 = F_get_database_oid(m, v16, int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v22 = F_convert_any_priv_string(m, v11, int32(1656912))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							v24 = F_object_aclcheck(m, int32(1262), v19, v13, v22)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v24 == int32(0))
							}
						}
					}
				}
			}
		}
	}
}
func F_rebuild_database_list(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v191 int32
	_ = v191
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v231 int32
	_ = v231
	var v232 int64
	_ = v232
	var v234 int64
	_ = v234
	var v236 int64
	_ = v236
	var v238 int64
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v274 float64
	_ = v274
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int64
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int64
	_ = v301
	var v302 int64
	_ = v302
	var v310 int64
	_ = v310
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v330 int64
	_ = v330
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v355 int64
	_ = v355
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v394 int64
	_ = v394
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = l0
	v23 = *(*int32)(unsafe.Add(mBase, _consts[661]))
	v28 = F_AllocSetContextCreateInternal(m, v23, int32(75483), v2, int32(8192), int32(8388608))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v34 = F_AllocSetContextCreateInternal(m, v28, int32(675480), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v36 = int32(4515248)
	v37 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v19)+44)) = int64(137438953476)
	v48 = F_hash_create(m, int32(323233), int32(20), v19+int32(28), int32(1064))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l0 == int32(0) {
		v65 = v2
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[707]))
	if v67 == int32(0) {
		v111 = v65
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v52 = F_pgstat_fetch_stat_dbentry(m, l0)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v52 == int32(0) {
		v65 = v2
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v56 = int32(1)
	v61 = F_hash_search(m, v48, v19+int32(76), v56, int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = int32(0)
	v65 = v56
	goto L5
L10:
	;
	v125 = F_get_database_list(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L22
	}
L11:
	;
	if v67 == int32(4121424) {
		v111 = v65
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v72 = v67
	v74 = v65
	goto L13
L13:
	;
	v89 = v72 - int32(20)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v91 = F_pgstat_fetch_stat_dbentry(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v111 = v105
	goto L10
L15:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v106 != int32(4121424) {
		v72 = v106
		v74 = v105
		goto L13
	} else {
		goto L20
	}
L16:
	;
	if v91 == int32(0) {
		v105 = v74
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v98 = F_hash_search(m, v48, v89, int32(1), v19+int32(8))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+8)))
	if v100 != 0 {
		v105 = v74
		goto L15
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v98)+16)) = v74
	v105 = v74 + int32(1)
	goto L15
L20:
	;
	goto L14
L21:
	;
	v191 = int32(4121424)
	*(*int32)(unsafe.Add(mBase, _consts[707])) = v191
	*(*int32)(unsafe.Add(mBase, _consts[708])) = v191
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v28
	if v176 <= int32(0) {
		goto L33
	} else {
		goto L34
	}
L22:
	;
	if v125 == int32(0) {
		v176 = v111
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v129 = int32(0)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v130 <= v129 {
		v176 = v111
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v133 = v129
	v135 = v111
	goto L25
L25:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149+v133<<(uint(int32(2))%32))))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v155 = F_pgstat_fetch_stat_dbentry(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v176 = v168
	goto L21
L27:
	;
	v171 = v133 + int32(1)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v171 < v172 {
		v133 = v171
		v135 = v168
		goto L25
	} else {
		goto L32
	}
L28:
	;
	if v155 == int32(0) {
		v168 = v135
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v162 = F_hash_search(m, v48, v153, int32(1), v19+int32(8))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+8)))
	if v164 != 0 {
		v168 = v135
		goto L27
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+16)) = v135
	v168 = v135 + int32(1)
	goto L27
L32:
	;
	goto L26
L33:
	;
	v438 = *(*int32)(unsafe.Add(mBase, _consts[709]))
	if v438 != 0 {
		goto L73
	} else {
		goto L74
	}
L34:
	;
	v202 = F_palloc(m, v176<<(uint(int32(5))%32))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_hash_seq_init(m, v19+int32(8), v48)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v210 = F_hash_seq_search(m, v19+int32(8))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v210 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v213 = v210
	v216 = int32(0)
	goto L41
L39:
	;
	goto L40
L40:
	;
	F_pg_qsort(m, v202, v176, int32(32), int32(920))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L45
	}
L41:
	;
	v231 = v202 + v216<<(uint(int32(5))%32)
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v213)))
	*(*int64)(unsafe.Add(mBase, uint32(v231))) = v232
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v213)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v231)+24)) = v234
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v213)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v231)+16)) = v236
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v213)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v231)+8)) = v238
	v244 = F_hash_seq_search(m, v19+int32(8))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	goto L40
L43:
	;
	if v244 != 0 {
		v213 = v244
		v216 = v216 + int32(1)
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v266 = int32(0)
	v269 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v274 = base.F64_div(base.F64_mul(base.F64_convert_i32_s(v269), float64(1000)), base.F64_convert_i32_u(v176))
	if base.F64_lt(base.F64_abs(v274), float64(2.147483648e+09)) != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v280 < int32(101) {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v278 = base.I32_trunc_f64_s(v274)
	v280 = v278
	goto L46
L48:
	;
	goto L49
L49:
	;
	v280 = int32(-2147483648)
	goto L46
L50:
	;
	v283 = int32(110)
	goto L52
L51:
	;
	v283 = v280
	goto L52
L52:
	;
	v286 = base.I64_extend_i32_u(v283) * int64(1000)
	v287 = int32(1)
	if v176 <= v287 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v290 = v287
	goto L55
L54:
	;
	v290 = v176
	goto L55
L55:
	;
	v296 = m.G0
	v297 = int32(16)
	v298 = v296 - v297
	m.G0 = v298
	F___gettimeofday(m, v298)
	mBase = m.M
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v298)))
	v302 = int64(*(*int32)(unsafe.Add(mBase, uint32(v298)+8)))
	m.G0 = v298 + v297
	v310 = v302 + v301*int64(1000000) - int64(946684800000000)
	goto L56
L56:
	;
	if v176 != int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v320 = v266
	v322 = int32(0)
	v330 = v310
	goto L60
L58:
	;
	v384 = v266
	v394 = v310
	goto L59
L59:
	;
	if v290&int32(1) == int32(0) {
		goto L33
	} else {
		goto L69
	}
L60:
	;
	v335 = v202 + v320<<(uint(int32(5))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v335)+8)) = v286 + v330
	v339 = v335 + int32(20)
	v341 = *(*int32)(unsafe.Add(mBase, _consts[707]))
	if v341 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v384 = v377
	v394 = v355
	goto L59
L62:
	;
	v344 = int32(4121424)
	*(*int32)(unsafe.Add(mBase, _consts[708])) = v344
	v348 = v344
	goto L64
L63:
	;
	v348 = v341
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v335)+20)) = int32(4121424)
	*(*int32)(unsafe.Add(mBase, uint32(v335)+24)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v348))) = v339
	v353 = int32(4121428)
	*(*int32)(unsafe.Add(mBase, _consts[707])) = v339
	v355 = v330 + (v286 + v286)
	*(*int64)(unsafe.Add(mBase, uint32(v335)+40)) = v355
	v358 = v335 + int32(52)
	v360 = v335 + int32(32)
	v362 = *(*int32)(unsafe.Add(mBase, _consts[707]))
	if v362 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v365 = int32(4121424)
	*(*int32)(unsafe.Add(mBase, _consts[708])) = v365
	v369 = v365
	goto L67
L66:
	;
	v369 = v362
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v360)+20)) = int32(4121424)
	*(*int32)(unsafe.Add(mBase, uint32(v360)+24)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = v358
	*(*int32)(unsafe.Add(mBase, _consts[707])) = v358
	v376 = int32(2)
	v377 = v320 + v376
	v379 = v322 + v376
	if v379 != v290&int32(2147483646) {
		v320 = v377
		v322 = v379
		v330 = v355
		goto L60
	} else {
		goto L68
	}
L68:
	;
	goto L61
L69:
	;
	v401 = v202 + v384<<(uint(int32(5))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v401)+8)) = v286 + v394
	v405 = v401 + int32(20)
	v407 = *(*int32)(unsafe.Add(mBase, _consts[707]))
	if v407 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v410 = int32(4121424)
	*(*int32)(unsafe.Add(mBase, _consts[708])) = v410
	v414 = v410
	goto L72
L71:
	;
	v414 = v407
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v401)+20)) = int32(4121424)
	*(*int32)(unsafe.Add(mBase, uint32(v401)+24)) = v414
	*(*int32)(unsafe.Add(mBase, uint32(v414))) = v405
	*(*int32)(unsafe.Add(mBase, _consts[707])) = v405
	goto L33
L73:
	;
	F_MemoryContextDelete(m, v438)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_MemoryContextDelete(m, v34)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v37
	*(*int32)(unsafe.Add(mBase, _consts[709])) = v28
	m.G0 = v19 + int32(80)
	return
}
