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
		v19 = F_psprintf(m, int32(_a_F_GetDatabasePath_0), v6+int32(16))
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
		v11 = F_pstrdup(m, int32(_a_F_GetDatabasePath_1))
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
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = int32(_a_F_GetDatabasePath_2)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_GetDatabasePath_3)
		v28 = F_psprintf(m, int32(_a_F_GetDatabasePath_4), v6)
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13916(m, l0, int32(_a_F_has_database_privilege_id_id_0), int32(1262))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
					v20 = F_convert_any_priv_string(m, v11, int32(_a_F_has_database_privilege_id_name_0))
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
						v22 = F_convert_any_priv_string(m, v11, int32(_a_F_has_database_privilege_name_name_0))
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v212 int32
	_ = v212
	var v213 int64
	_ = v213
	var v215 int64
	_ = v215
	var v217 int64
	_ = v217
	var v219 int64
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int64
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int64
	_ = v267
	var v268 int64
	_ = v268
	var v276 int64
	_ = v276
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v297 int64
	_ = v297
	var v301 int32
	_ = v301
	var v302 int64
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int64
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v358 int64
	_ = v358
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
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = l0
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[0]))
	v25 = F_AllocSetContextCreateInternal(m, v20, int32(_a_F_rebuild_database_list_0), v2, int32(_a_F_rebuild_database_list_1), int32(_a_F_rebuild_database_list_2))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v31 = F_AllocSetContextCreateInternal(m, v25, int32(_a_F_rebuild_database_list_3), int32(0), int32(_a_F_rebuild_database_list_1), int32(_a_F_rebuild_database_list_2))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = int32(_a_F_rebuild_database_list_4)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[1])) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v16)+44)) = int64(137438953476)
	v45 = F_hash_create(m, int32(_a_F_rebuild_database_list_5), int32(20), v16+int32(28), int32(1064))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l0 == int32(0) {
		v62 = v2
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[2]))
	v65 = int32(0)
	if base.B2i32(v64 == v65)|base.B2i32(v64 == int32(_a_F_rebuild_database_list_6)) == v65 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v49 = F_pgstat_fetch_stat_dbentry(m, l0)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v49 == int32(0) {
		v62 = v2
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v53 = int32(1)
	v58 = F_hash_search(m, v45, v16+int32(76), v53, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = int32(0)
	v62 = v53
	goto L5
L10:
	;
	v72 = v64
	v73 = v62
	goto L13
L11:
	;
	v106 = v62
	goto L12
L12:
	;
	v118 = F_get_database_list(m)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L22
	}
L13:
	;
	v86 = v72 - int32(20)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v88 = F_pgstat_fetch_stat_dbentry(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v106 = v101
	goto L12
L15:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	if v102 != int32(_a_F_rebuild_database_list_6) {
		v72 = v102
		v73 = v101
		goto L13
	} else {
		goto L20
	}
L16:
	;
	if v88 == int32(0) {
		v101 = v73
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v95 = F_hash_search(m, v45, v86, int32(1), v16+int32(8))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)))
	if v97 != 0 {
		v101 = v73
		goto L15
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95)+16)) = v73
	v101 = v73 + int32(1)
	goto L15
L20:
	;
	goto L14
L21:
	;
	v177 = int32(_a_F_rebuild_database_list_6)
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[2])) = v177
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[3])) = v177
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[1])) = v25
	if v164 <= int32(0) {
		goto L33
	} else {
		goto L34
	}
L22:
	;
	if v118 == int32(0) {
		v164 = v106
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	if v122 <= int32(0) {
		v164 = v106
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v126 = int32(0)
	v127 = v106
	goto L25
L25:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v118)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139+v126<<(uint(int32(2))%32))))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v145 = F_pgstat_fetch_stat_dbentry(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v164 = v158
	goto L21
L27:
	;
	v160 = v126 + int32(1)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	if v160 < v161 {
		v126 = v160
		v127 = v158
		goto L25
	} else {
		goto L32
	}
L28:
	;
	if v145 == int32(0) {
		v158 = v127
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v152 = F_hash_search(m, v45, v143, int32(1), v16+int32(8))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)))
	if v154 != 0 {
		v158 = v127
		goto L27
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+16)) = v127
	v158 = v127 + int32(1)
	goto L27
L32:
	;
	goto L26
L33:
	;
	v396 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[4]))
	if v396 != 0 {
		goto L67
	} else {
		goto L68
	}
L34:
	;
	v188 = F_palloc(m, v164<<(uint(int32(5))%32))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v191 = v16 + int32(8)
	F_hash_seq_init(m, v191, v45)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v194 = F_hash_seq_search(m, v191)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v194 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v197 = v194
	v199 = int32(0)
	goto L41
L39:
	;
	goto L40
L40:
	;
	F_pg_qsort(m, v188, v164, int32(32), int32(920))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L45
	}
L41:
	;
	v212 = v188 + v199<<(uint(int32(5))%32)
	v213 = *(*int64)(unsafe.Add(mBase, uint32(v197)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v212)+24)) = v213
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v197)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v212)+16)) = v215
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v197)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v212)+8)) = v217
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v197)))
	*(*int64)(unsafe.Add(mBase, uint32(v212))) = v219
	v225 = F_hash_seq_search(m, v16+int32(8))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	goto L40
L43:
	;
	if v225 != 0 {
		v197 = v225
		v199 = v199 + int32(1)
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[5]))
	v252 = base.I32_trunc_sat_f64_s(base.F64_div(base.F64_mul(base.F64_convert_i32_s(v246), float64(1000)), base.F64_convert_i32_u(v164)))
	if v252 < int32(101) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v255 = int32(110)
	goto L48
L47:
	;
	v255 = v252
	goto L48
L48:
	;
	v258 = base.I64_extend_i32_u(v255) * int64(1000)
	v262 = m.G0
	v263 = int32(16)
	v264 = v262 - v263
	m.G0 = v264
	F_gettimeofday(m, v264)
	mBase = m.M
	v267 = *(*int64)(unsafe.Add(mBase, uint32(v264)))
	v268 = int64(*(*int32)(unsafe.Add(mBase, uint32(v264)+8)))
	m.G0 = v264 + v263
	v276 = v268 + v267*int64(1000000) - int64(946684800000000)
	goto L49
L49:
	;
	if v164 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v362 = v188 + v349<<(uint(int32(5))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v362)+8)) = v358 + v258
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[2]))
	if v366 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L51:
	;
	v349 = int32(0)
	v358 = v276
	goto L50
L52:
	;
	goto L53
L53:
	;
	v284 = int32(0)
	v288 = v284
	v291 = v284
	v297 = v276
	goto L54
L54:
	;
	v301 = v188 + v288<<(uint(int32(5))%32)
	v302 = v297 + v258
	*(*int64)(unsafe.Add(mBase, uint32(v301)+8)) = v302
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[2]))
	if v305 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v164&int32(1) == int32(0) {
		goto L33
	} else {
		goto L63
	}
L56:
	;
	v308 = int32(_a_F_rebuild_database_list_6)
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[3])) = v308
	v312 = v308
	goto L58
L57:
	;
	v312 = v305
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v301)+20)) = int32(_a_F_rebuild_database_list_6)
	*(*int32)(unsafe.Add(mBase, uint32(v301)+24)) = v312
	v317 = v301 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v312))) = v317
	v319 = int32(_a_F_rebuild_database_list_7)
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[2])) = v317
	v321 = v302 + v258
	*(*int64)(unsafe.Add(mBase, uint32(v301)+40)) = v321
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[2]))
	if v324 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v327 = int32(_a_F_rebuild_database_list_6)
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[3])) = v327
	v331 = v327
	goto L61
L60:
	;
	v331 = v324
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v301)+52)) = int32(_a_F_rebuild_database_list_6)
	*(*int32)(unsafe.Add(mBase, uint32(v301)+56)) = v331
	v336 = v301 + int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = v336
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[2])) = v336
	v340 = int32(2)
	v341 = v288 + v340
	v343 = v291 + v340
	if v343 != v164&int32(2147483646) {
		v288 = v341
		v291 = v343
		v297 = v321
		goto L54
	} else {
		goto L62
	}
L62:
	;
	goto L55
L63:
	;
	v349 = v341
	v358 = v321
	goto L50
L64:
	;
	v369 = int32(_a_F_rebuild_database_list_6)
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[3])) = v369
	v373 = v369
	goto L66
L65:
	;
	v373 = v366
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v362)+20)) = int32(_a_F_rebuild_database_list_6)
	*(*int32)(unsafe.Add(mBase, uint32(v362)+24)) = v373
	v378 = v362 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v373))) = v378
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[2])) = v378
	goto L33
L67:
	;
	F_MemoryContextDelete(m, v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	F_MemoryContextDelete(m, v31)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[1])) = v34
	*(*int32)(unsafe.Add(mBase, _c_F_rebuild_database_list[4])) = v25
	m.G0 = v16 + int32(80)
	return
}
