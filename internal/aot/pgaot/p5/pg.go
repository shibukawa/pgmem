package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetPGVariable(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
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
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
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
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v18 = l0
	v19 = int32(_a_F_GetPGVariable_0)
	goto L3
L1:
	;
	m.G0 = v13 + int32(32)
	return
L2:
	;
	if v64 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L3:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v64 = base.I32_extend8_s(v44) - base.I32_extend8_s(v53)
	goto L2
L5:
	;
	v32 = int32(1)
	if base.Ui32((v23-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	if v22&int32(255) != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if v22&int32(255) != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v64 = int32(1)
	goto L2
L10:
	;
	v31 = int32(-1)
	goto L12
L11:
	;
	v31 = int32(0)
	goto L12
L12:
	;
	v64 = v31
	goto L2
L13:
	;
	v44 = v23 | int32(32)
	goto L15
L14:
	;
	v44 = v23
	goto L15
L15:
	;
	if base.Ui32((v22-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v53 = v22 | int32(32)
	goto L18
L17:
	;
	v53 = v22
	goto L18
L18:
	;
	if v44 == v53&int32(255) {
		v18 = v18 + v32
		v19 = v19 + v32
		goto L3
	} else {
		goto L19
	}
L19:
	;
	goto L4
L20:
	;
	v67 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v67)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)) = uint16(v67)
	v74 = F_get_guc_variables(m, v13+int32(28))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v196 = F_GetConfigOptionByName(m, l0, v13+int32(16), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L23
	} else {
		goto L63
	}
L23:
	;
	return
L24:
	;
	v77 = F_CreateTemplateTupleDesc(m, int32(3))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_TupleDescInitBuiltinEntry(m, v77, int32(1), int32(_a_F_GetPGVariable_1), int32(25))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	F_TupleDescInitBuiltinEntry(m, v77, int32(2), int32(_a_F_GetPGVariable_2), int32(25))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	F_TupleDescInitBuiltinEntry(m, v77, int32(3), int32(_a_F_GetPGVariable_3), int32(25))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v95 = F_begin_tup_output_tupdesc(m, l1, v77, int32(_a_F_GetPGVariable_4))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	if int32(0) < v97 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v100 = v67
	goto L33
L31:
	;
	goto L32
L32:
	;
	F_end_tup_output(m, v95)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L23
	} else {
		goto L62
	}
L33:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v74+v100<<(uint(int32(2))%32))))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+20))
	if v114&int32(4) != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v178 = v100 + int32(1)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	if v178 < v179 {
		v100 = v178
		goto L33
	} else {
		goto L61
	}
L36:
	;
	if v114&int32(1024) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_GetPGVariable[0]))
	v122 = F_has_privs_of_role(m, v120, int32(3374))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L23
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v127 = F_cstring_to_text(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L23
	} else {
		goto L42
	}
L40:
	;
	if v122 == int32(0) {
		goto L35
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v127
	v130 = int32(1)
	v131 = int32(0)
	v135 = F_ShowGUCOption(m, v113, v130)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L23
	} else {
		goto L43
	}
L43:
	;
	if v135 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v138 = F_cstring_to_text(m, v135)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L23
	} else {
		goto L47
	}
L45:
	;
	v140 = v131
	v141 = v130
	goto L46
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v141)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v140
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	if v144 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v140 = v138
	v141 = int32(0)
	goto L46
L48:
	;
	v146 = F_cstring_to_text(m, v144)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L23
	} else {
		goto L51
	}
L49:
	;
	v148 = v131
	v149 = v130
	goto L50
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+14)) = uint8(v149)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v148
	F_do_tup_output(m, v95, v13+int32(16), v13+int32(12))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L23
	} else {
		goto L52
	}
L51:
	;
	v148 = v146
	v149 = int32(0)
	goto L50
L52:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	F_pfree(m, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L23
	} else {
		goto L53
	}
L53:
	;
	if v135 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_pfree(m, v135)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L23
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	if v166 == int32(0) {
		goto L35
	} else {
		goto L59
	}
L57:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	F_pfree(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L23
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	F_pfree(m, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L23
	} else {
		goto L60
	}
L60:
	;
	goto L35
L61:
	;
	goto L34
L62:
	;
	goto L1
L63:
	;
	v199 = F_CreateTemplateTupleDesc(m, int32(1))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L23
	} else {
		goto L64
	}
L64:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	F_TupleDescInitBuiltinEntry(m, v199, int32(1), v202, int32(25))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L23
	} else {
		goto L65
	}
L65:
	;
	v207 = F_begin_tup_output_tupdesc(m, l1, v199, int32(_a_F_GetPGVariable_4))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L23
	} else {
		goto L66
	}
L66:
	;
	v209 = F_cstring_to_text(m, v196)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L23
	} else {
		goto L67
	}
L67:
	;
	v211 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+12)) = uint8(v211)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v209
	F_do_tup_output(m, v207, v13+int32(28), v13+int32(12))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L23
	} else {
		goto L68
	}
L68:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	F_pfree(m, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L23
	} else {
		goto L69
	}
L69:
	;
	F_end_tup_output(m, v207)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L23
	} else {
		goto L70
	}
L70:
	;
	goto L1
}
func F_do_pg_abort_backup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	if l1 == int32(0) {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_abort_backup[0])))
		if v6&int32(1) == int32(0) {
			return
		} else {
			F_WALInsertLockAcquireExclusive(m)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_abort_backup[1]))
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+164))
				*(*int32)(unsafe.Add(mBase, uint32(v14)+164)) = v15 - int32(1)
				v20 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_abort_backup[0])) = uint8(v20)
				F_WALInsertLockRelease(m)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					if l1 != 0 {
						return
					} else {
						v26 = F_errstart(m, int32(19), int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							if v26 == int32(0) {
								return
							} else {
								F_errmsg(m, int32(_a_F_do_pg_abort_backup_0), int32(0))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_do_pg_abort_backup_1), int32(_a_F_do_pg_abort_backup_2), int32(_a_F_do_pg_abort_backup_3))
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_WALInsertLockAcquireExclusive(m)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_abort_backup[1]))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+164))
			*(*int32)(unsafe.Add(mBase, uint32(v14)+164)) = v15 - int32(1)
			v20 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_abort_backup[0])) = uint8(v20)
			F_WALInsertLockRelease(m)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				if l1 != 0 {
					return
				} else {
					v26 = F_errstart(m, int32(19), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						if v26 == int32(0) {
							return
						} else {
							F_errmsg(m, int32(_a_F_do_pg_abort_backup_0), int32(0))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_do_pg_abort_backup_1), int32(_a_F_do_pg_abort_backup_2), int32(_a_F_do_pg_abort_backup_3))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_do_pg_backup_start(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
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
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v115 int32
	_ = v115
	var v129 int32
	_ = v129
	var v144 int32
	_ = v144
	var v159 int32
	_ = v159
	var v175 int32
	_ = v175
	var v187 int32
	_ = v187
	var v204 int32
	_ = v204
	var v218 int32
	_ = v218
	var v236 int32
	_ = v236
	var v252 int32
	_ = v252
	var v270 int32
	_ = v270
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v410 int32
	_ = v410
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
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
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v466 int32
	_ = v466
	var v480 int64
	_ = v480
	var v481 int32
	_ = v481
	var v523 int32
	_ = v523
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int64
	_ = v544
	var v546 int64
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v601 int64
	_ = v601
	var v602 int64
	_ = v602
	var v618 int32
	_ = v618
	var v632 int32
	_ = v632
	var v647 int32
	_ = v647
	var v662 int32
	_ = v662
	var v678 int32
	_ = v678
	var v693 int32
	_ = v693
	var v694 int64
	_ = v694
	var v696 int32
	_ = v696
	var v697 int64
	_ = v697
	var v712 int32
	_ = v712
	var v725 int32
	_ = v725
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v774 int32
	_ = v774
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v834 int32
	_ = v834
	var v837 int64
	_ = v837
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v896 int32
	_ = v896
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	var v956 int32
	_ = v956
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1031 int32
	_ = v1031
	var v1048 int32
	_ = v1048
	var v1061 int32
	_ = v1061
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1128 int32
	_ = v1128
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1185 int32
	_ = v1185
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1225 int32
	_ = v1225
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1255 int32
	_ = v1255
	var v1270 int32
	_ = v1270
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1355 int32
	_ = v1355
	var v1367 int64
	_ = v1367
	var v1383 int32
	_ = v1383
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1417 int32
	_ = v1417
	var v1431 int32
	_ = v1431
	var v1444 int32
	_ = v1444
	var v1465 int32
	_ = v1465
	var v1474 int32
	_ = v1474
	var v1475 int64
	_ = v1475
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	v6 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(128)
	m.G0 = v32
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v36 = int32(44)
	goto L3
L2:
	;
	v36 = int32(40)
	goto L3
L3:
	;
	v39 = l1
	v44 = v6
	v45 = v6
	v46 = v6
	v47 = v6
	v48 = v6
	v49 = v6
	v50 = v6
	v51 = v6
	v52 = v6
	v53 = v6
	v54 = int32(-1)
	v55 = v6
	v58 = v32
	goto L4
L4:
	;
	goto L7
L5:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L6:
	;
	goto L5
L7:
	;
	if v54 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v1474 = int32(m.ExcTag)
	v1475 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1474 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L10:
	;
	v70 = v58 - int32(160)
	m.G0 = v70
	v73 = v70 - int32(1040)
	m.G0 = v73
	v76 = v73 - int32(1024)
	m.G0 = v76
	v78 = int32(16)
	v79 = v76 - v78
	m.G0 = v79
	v82 = v79 - v78
	m.G0 = v82
	v86 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_backup_start[0])))
	if v86 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v436 = v44
	v437 = v45
	v438 = v46
	v439 = v47
	v440 = v48
	v441 = v49
	v442 = v50
	v443 = v51
	v444 = v55
	v445 = v58
	goto L12
L12:
	;
	if v443 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L13:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[1]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+316))
	v94 = base.B2i32(v92 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_backup_start[0])) = uint8(v94)
	v96 = v94
	goto L15
L14:
	;
	v96 = int32(0)
	goto L15
L15:
	;
	if v96 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v70
	v187 = F_strlen(m, l0)
	mBase = m.M
	if base.Ui32(int32(1025)) <= base.Ui32(v187) {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[2]))
	if int32(0) < v98 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v70
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		v1465 = v82
		goto L9
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v70
	F_errcode(m, int32(325))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		v1465 = v82
		goto L9
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v70
	F_errmsg(m, int32(_a_F_do_pg_backup_start_0), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		v1465 = v82
		goto L9
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v70
	F_errhint(m, int32(_a_F_do_pg_backup_start_1), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		v1465 = v82
		goto L9
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v70
	F_errfinish(m, int32(_a_F_do_pg_backup_start_2), int32(_a_F_do_pg_backup_start_3), int32(_a_F_do_pg_backup_start_4))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		v1465 = v82
		goto L9
	} else {
		goto L23
	}
L23:
	;
	goto L6
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v70
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		v1465 = v82
		goto L9
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v70
	goto L34
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v70
	F_errcode(m, int32(50856066))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		v1465 = v82
		goto L9
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = int32(1024)
	F_errmsg(m, int32(_a_F_do_pg_backup_start_5), v32-int32(-64))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		v1465 = v82
		goto L9
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v70
	F_errfinish(m, int32(_a_F_do_pg_backup_start_2), int32(_a_F_do_pg_backup_start_6), int32(_a_F_do_pg_backup_start_4))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		v1465 = v82
		goto L9
	} else {
		goto L30
	}
L30:
	;
	goto L6
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v70
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		v1465 = v82
		goto L9
	} else {
		goto L63
	}
L32:
	;
	v376 = F_strlen(m, v365)
	mBase = m.M
	goto L31
L34:
	;
	goto L35
L35:
	;
	v270 = int32(1024)
	if (l3^l0)&int32(3) != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v369 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v366))) = uint8(v369)
	goto L32
L37:
	;
	v350 = v345
	v351 = v346
	v352 = v347
	goto L59
L38:
	;
	if v340 == int32(0) {
		v365 = v338
		v366 = v339
		goto L36
	} else {
		goto L58
	}
L39:
	;
	v338 = l0
	v339 = l3
	v340 = v270
	goto L38
L40:
	;
	goto L41
L41:
	;
	if l0&int32(3) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v307 == int32(0) {
		v365 = v304
		v366 = v305
		goto L36
	} else {
		goto L51
	}
L43:
	;
	v304 = l0
	v305 = l3
	v306 = v270
	v307 = int32(1)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v283 = l0
	v284 = l3
	v285 = v270
	goto L46
L46:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283))))
	*(*uint8)(unsafe.Add(mBase, uint32(v284))) = uint8(v287)
	if v287 == int32(0) {
		v345 = v283
		v346 = v284
		v347 = v285
		goto L37
	} else {
		goto L48
	}
L47:
	;
	v304 = v298
	v305 = v292
	v306 = v294
	v307 = v296
	goto L42
L48:
	;
	v291 = int32(1)
	v292 = v284 + v291
	v294 = v285 - v291
	v295 = int32(0)
	v296 = base.B2i32(v294 != v295)
	v298 = v283 + v291
	if v298&int32(3) == v295 {
		v304 = v298
		v305 = v292
		v306 = v294
		v307 = v296
		goto L42
	} else {
		goto L49
	}
L49:
	;
	if v294 != 0 {
		v283 = v298
		v284 = v292
		v285 = v294
		goto L46
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	if v310 == int32(0) {
		v338 = v304
		v339 = v305
		v340 = v306
		goto L38
	} else {
		goto L52
	}
L52:
	;
	if base.Ui32(v306) < base.Ui32(int32(4)) {
		v338 = v304
		v339 = v305
		v340 = v306
		goto L38
	} else {
		goto L53
	}
L53:
	;
	v316 = v304
	v317 = v305
	v318 = v306
	goto L54
L54:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	v324 = int32(-2139062144)
	if (int32(16843008)-v321|v321)&v324 != v324 {
		v345 = v316
		v346 = v317
		v347 = v318
		goto L37
	} else {
		goto L56
	}
L55:
	;
	v338 = v332
	v339 = v330
	v340 = v334
	goto L38
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v317))) = v321
	v329 = int32(4)
	v330 = v317 + v329
	v332 = v316 + v329
	v334 = v318 - v329
	if base.Ui32(int32(3)) < base.Ui32(v334) {
		v316 = v332
		v317 = v330
		v318 = v334
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v345 = v338
	v346 = v339
	v347 = v340
	goto L37
L59:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350))))
	*(*uint8)(unsafe.Add(mBase, uint32(v351))) = uint8(v354)
	if v354 == int32(0) {
		v365 = v350
		v366 = v351
		goto L36
	} else {
		goto L61
	}
L60:
	;
	v365 = v361
	v366 = v359
	goto L36
L61:
	;
	v358 = int32(1)
	v359 = v351 + v358
	v361 = v350 + v358
	v363 = v352 - v358
	if v363 != 0 {
		v350 = v361
		v351 = v359
		v352 = v363
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v393 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[1]))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v393)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v393)+164)) = v394 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v70
	F_WALInsertLockRelease(m)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		v1465 = v82
		goto L9
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v50
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v70
	F_before_shmem_exit(m, int32(407), int32(1))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		v1465 = v82
		goto L9
	} else {
		goto L65
	}
L65:
	;
	v428 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[3]))
	v430 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[4]))
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v32 + int32(80)
	goto L69
L67:
	;
	v436 = v76
	v437 = v82
	v438 = v73
	v439 = v70
	v440 = v79
	v441 = v428
	v442 = v430
	v443 = int32(0)
	v444 = v96
	v445 = v82
	goto L12
L69:
	;
	goto L67
L70:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[4])) = v439
	v451 = v444 & int32(1)
	if v451 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[3])) = v441
	*(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[4])) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	v1411 = int32(1)
	v1412 = v444 & v1411
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v1412)
	F_cancel_before_shmem_exit(m, int32(407), v1411)
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L179
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	F_XLogBeginInsert(m)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	goto L78
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	v480 = F_XLogInsert(m, int32(0), int32(64))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	F_RequestCheckpoint(m, v36)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L80
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	v740 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[5]))
	v741 = F_strlen(m, v740)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	v754 = F_AllocateDir(m, int32(_a_F_do_pg_backup_start_7))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L106
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	v536 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[6]))
	v540 = F_LWLockAcquire(m, v536+int32(1152), int32(1))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L81
	}
L81:
	;
	v543 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[7]))
	v544 = *(*int64)(unsafe.Add(mBase, uint32(v543)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l3)+1048)) = v544
	v546 = *(*int64)(unsafe.Add(mBase, uint32(v543)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l3)+1032)) = v546
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v543)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+1040)) = v548
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v543)+56)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	v563 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[6]))
	F_LWLockRelease(m, v563+int32(1152))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L82
	}
L82:
	;
	if v451 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L98
	}
L84:
	;
	v571 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[1]))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v571)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v571)+440)) = int32(1)
	if v572 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	v587 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[1]))
	F_s_lock(m, v587+int32(440), int32(_a_F_do_pg_backup_start_2), int32(_a_F_do_pg_backup_start_8), int32(_a_F_do_pg_backup_start_4))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v596 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v596)+440)) = int32(0)
	if v550&int32(1) != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L87
L89:
	;
	v601 = *(*int64)(unsafe.Add(mBase, uint32(l3)+1032))
	v602 = *(*int64)(unsafe.Add(mBase, uint32(v596)+432))
	if base.Ui64(v602) < base.Ui64(v601) {
		goto L83
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	F_errcode(m, int32(325))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	F_errmsg(m, int32(_a_F_do_pg_backup_start_9), int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	F_errhint(m, int32(_a_F_do_pg_backup_start_10), int32(0))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	F_errfinish(m, int32(_a_F_do_pg_backup_start_2), int32(_a_F_do_pg_backup_start_11), int32(_a_F_do_pg_backup_start_4))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L97
	}
L97:
	;
	goto L6
L98:
	;
	v694 = *(*int64)(unsafe.Add(mBase, uint32(l3)+1032))
	v696 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[1]))
	v697 = *(*int64)(unsafe.Add(mBase, uint32(v696)+168))
	if base.Ui64(v697) < base.Ui64(v694) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L79
L100:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v696)+168)) = v694
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	F_WALInsertLockRelease(m)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	F_WALInsertLockRelease(m)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L104
	}
L103:
	;
	goto L99
L104:
	;
	if v451 == int32(0) {
		goto L78
	} else {
		goto L105
	}
L105:
	;
	goto L99
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	v768 = F_ReadDir(m, v754, int32(_a_F_do_pg_backup_start_7))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L107
	}
L107:
	;
	if v768 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v770 = v436 + v741
	v774 = v39
	v787 = v52
	v788 = v53
	v789 = v768
	goto L111
L109:
	;
	v1315 = v39
	v1328 = v52
	v1329 = v53
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	F_FreeDir(m, v754)
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L177
	}
L111:
	;
	v802 = int32(*(*int8)(unsafe.Add(mBase, uint32(v789)+19)))
	if v802 < int32(49) {
		v1285 = v787
		v1286 = v788
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v1315 = v1312
	v1328 = v1285
	v1329 = v1286
	goto L110
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v1285
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v1286
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	v1312 = F_ReadDir(m, v754, int32(_a_F_do_pg_backup_start_7))
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L175
	}
L114:
	;
	v805 = int32(*(*int8)(unsafe.Add(mBase, uint32(v789)+20)))
	if int32(57) < v805 {
		v1285 = v787
		v1286 = v788
		goto L113
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	*(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[8])) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	v834 = v789 + int32(19)
	v837 = F_strtox_2(m, v834, v440, int32(10), int64(4294967295))
	mBase = m.M
	goto L116
L116:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v839))))
	if v840 != 0 {
		v1285 = v787
		v1286 = v788
		goto L113
	} else {
		goto L117
	}
L117:
	;
	v842 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[8]))
	if v842 == int32(28) {
		v1285 = v787
		v1286 = v788
		goto L113
	} else {
		goto L118
	}
L118:
	;
	if v842 == int32(68) {
		v1285 = v787
		v1286 = v788
		goto L113
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v834
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = int32(_a_F_do_pg_backup_start_7)
	v865 = F_pg_snprintf(m, v438, int32(1034), int32(_a_F_do_pg_backup_start_12), v32+int32(48))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	v880 = F_get_dirent_type(m, v438, v789, int32(0), int32(21))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L125
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v438
	F_errmsg(m, v1240, v32)
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L173
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v1176
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v1177
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	v1203 = F_palloc(m, int32(24))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L169
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v834
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = int32(_a_F_do_pg_backup_start_7)
	v1147 = F_pg_snprintf(m, v436, int32(1024), int32(_a_F_do_pg_backup_start_12), v32+int32(32))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L167
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	v896 = F_readlink(m, v438, v436, int32(1024))
	mBase = m.M
	if v896 < int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	switch v880 - int32(3) {
	case 0:
		goto L123
	case 1:
		goto L124
	default:
		v1285 = v787
		v1286 = v788
		goto L113
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	v912 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	if base.Ui32(int32(1024)) <= base.Ui32(v896) {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	if v912 == int32(0) {
		v1285 = v787
		v1286 = v788
		goto L113
	} else {
		goto L130
	}
L130:
	;
	v1240 = int32(_a_F_do_pg_backup_start_13)
	v1241 = int32(_a_F_do_pg_backup_start_14)
	goto L121
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	v933 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v939 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v436+v896))) = uint8(v939)
	if v896 <= v741 {
		v1017 = v788
		v1018 = v939
		goto L136
	} else {
		goto L137
	}
L134:
	;
	if v933 == int32(0) {
		v1285 = v787
		v1286 = v788
		goto L113
	} else {
		goto L135
	}
L135:
	;
	v1240 = int32(_a_F_do_pg_backup_start_15)
	v1241 = int32(_a_F_do_pg_backup_start_16)
	goto L121
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v1017
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	F_initStringInfo(m, v437)
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L155
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	v956 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[5]))
	if v741 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	if v1000 != 0 {
		v1017 = v788
		v1018 = v939
		goto L136
	} else {
		goto L152
	}
L139:
	;
	v1000 = int32(0)
	goto L138
L140:
	;
	goto L141
L141:
	;
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436))))
	if v962 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v963 = v436
	v964 = v956
	v965 = v741
	v966 = v962
	goto L146
L143:
	;
	v988 = v956
	v992 = int32(0)
	goto L144
L144:
	;
	v993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v988))))
	v1000 = v992 - v993
	goto L138
L145:
	;
	v988 = v983
	v992 = v985
	goto L144
L146:
	;
	v968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v964))))
	if v966 != v968 {
		v983 = v964
		v985 = v966
		goto L145
	} else {
		goto L148
	}
L147:
	;
	v983 = v977
	v985 = int32(0)
	goto L145
L148:
	;
	if v968 == int32(0) {
		v983 = v964
		v985 = v966
		goto L145
	} else {
		goto L149
	}
L149:
	;
	v973 = v965 - int32(1)
	if v973 == int32(0) {
		v983 = v964
		v985 = v966
		goto L145
	} else {
		goto L150
	}
L150:
	;
	v976 = int32(1)
	v977 = v964 + v976
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v963)+1)))
	if v978 != 0 {
		v963 = v963 + v976
		v964 = v977
		v965 = v973
		v966 = v978
		goto L146
	} else {
		goto L151
	}
L151:
	;
	goto L147
L152:
	;
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770))))
	if v1001 != int32(47) {
		v1017 = v788
		v1018 = v939
		goto L136
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	v1015 = F_pstrdup(m, v770+int32(1))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L154
	}
L154:
	;
	v1017 = v1015
	v1018 = v1015
	goto L136
L155:
	;
	v1048 = v436
	goto L156
L156:
	;
	v1061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1048))))
	switch v1061 {
	case 0:
		goto L158
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12:
		v1079 = v1061
		goto L159
	case 10, 13:
		goto L160
	default:
		goto L161
	}
L157:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v1017
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v1096
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v834
	F_appendStringInfo(m, l4, int32(_a_F_do_pg_backup_start_17), v32+int32(16))
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L165
	}
L158:
	;
	goto L157
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v1017
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	F_appendStringInfoChar(m, v437, base.I32_extend8_s(v1079))
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L164
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v1017
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	F_appendStringInfoChar(m, v437, int32(92))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L163
	}
L161:
	;
	if v1061 != int32(92) {
		v1079 = v1061
		goto L159
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	v1078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1048))))
	v1079 = v1078
	goto L159
L164:
	;
	v1048 = v1048 + int32(1)
	goto L156
L165:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v1017
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	F_pfree(m, v1115)
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L166
	}
L166:
	;
	v1176 = v787
	v1177 = v1017
	v1185 = v1018
	goto L122
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	v1160 = F_pstrdup(m, v436)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L168
	}
L168:
	;
	v1176 = v1160
	v1177 = v788
	v1185 = v1160
	goto L122
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1203))) = base.I32_wrap_i64(v837)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v1176
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v1177
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	v1217 = F_pstrdup(m, v436)
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L170
	}
L170:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1203)+16)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v1203)+8)) = v1185
	*(*int32)(unsafe.Add(mBase, uint32(v1203)+4)) = v1217
	if l2 == int32(0) {
		v1285 = v1176
		v1286 = v1177
		goto L113
	} else {
		goto L171
	}
L171:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v1176
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v1177
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	v1237 = F_lappend(m, v1225, v1203)
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1237
	v1285 = v1176
	v1286 = v1177
	goto L113
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	F_errfinish(m, int32(_a_F_do_pg_backup_start_2), v1241, int32(_a_F_do_pg_backup_start_4))
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L174
	}
L174:
	;
	v1285 = v787
	v1286 = v788
	goto L113
L175:
	;
	if v1312 != 0 {
		v774 = v1312
		v787 = v1285
		v788 = v1286
		v789 = v1312
		goto L111
	} else {
		goto L176
	}
L176:
	;
	goto L112
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	v1367 = F___time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l3)+1056)) = v1367
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v1328
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v1315
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v1329
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v451)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	F_cancel_before_shmem_exit(m, int32(407), int32(1))
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[3])) = v441
	*(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_start[4])) = v442
	v1388 = int32(1)
	v1389 = v444 & v1388
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+1064)) = uint8(v1389)
	*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_backup_start[9])) = uint8(v1388)
	m.G0 = v32 + int32(128)
	return
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v1412)
	F_do_pg_abort_backup(m, v32, int32(1))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+88)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v32)+92)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v32)+100)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v32)+108)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v32)+116)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v32)+120)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v32)+124)) = v439
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)) = uint8(v1412)
	F_pg_re_throw(m)
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		v1465 = v445
		goto L9
	} else {
		goto L181
	}
L181:
	;
	goto L8
L182:
	;
	v1479 = int32(v1475)
	m.G0 = v1465
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v1479)+4))
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v1479)))
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(v1482)))
	if v32+int32(80) == v1486 {
		goto L185
	} else {
		goto L186
	}
L183:
	;
	m.ExcPending = 1
	goto L191
L184:
	;
	if v1489 != 0 {
		goto L188
	} else {
		goto L189
	}
L185:
	;
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1482)+4))
	v1489 = v1488
	goto L187
L186:
	;
	v1489 = int32(0)
	goto L187
L187:
	;
	goto L184
L188:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v32)+124))
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v32)+120))
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v32)+116))
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v32)+112))
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v32)+108))
	v1495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+107)))
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v32)+100))
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v32)+96))
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v32)+92))
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v32)+88))
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v32)+84))
	v39 = v1500
	v44 = v1492
	v45 = v1494
	v46 = v1491
	v47 = v1490
	v48 = v1493
	v49 = v1497
	v50 = v1496
	v51 = v1481
	v52 = v1499
	v53 = v1498
	v54 = v1489
	v55 = v1495
	v58 = v1465
	goto L4
L189:
	;
	goto L190
L190:
	;
	F___wasm_longjmp(m, v1482, v1481)
	mBase = m.M
	v1502 = m.ExcPending
	if v1502 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	return
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_armor(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
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
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	switch v50 - int32(1) {
	case 0:
		v185 = v2
		v188 = v2
		v189 = v2
		goto L24
	default:
		goto L14
	case 2:
		goto L25
	}
L2:
	;
	return int32(0)
L3:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v19 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = int32(4)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v24&int32(254) == int32(2) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v37 = int32(1)
	if v19&v37 != 0 {
		v49 = int32(base.Ui32(v19)>>(uint(v37)%32)) - v37
		goto L1
	} else {
		goto L13
	}
L7:
	;
	v33 = v22
	goto L9
L8:
	;
	v33 = base.B2i32(v24 == int32(18)) << (uint(v22) % 32)
	goto L9
L9:
	;
	if v24 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v36 = v22
	goto L12
L11:
	;
	v36 = v33
	goto L12
L12:
	;
	v49 = v36
	goto L1
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L2
	} else {
		goto L118
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L2
	} else {
		goto L114
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L2
	} else {
		goto L110
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L2
	} else {
		goto L106
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L2
	} else {
		goto L102
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L2
	} else {
		goto L98
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L2
	} else {
		goto L94
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L2
	} else {
		goto L90
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L2
	} else {
		goto L86
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L2
	} else {
		goto L82
	}
L24:
	;
	F_initStringInfo(m, v12+int32(12))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L2
	} else {
		goto L67
	}
L25:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v54 = F_pg_detoast_datum(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v57 = F_pg_detoast_datum(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if int32(1) < v59 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v59 != v62 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	if v59 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v185 = int32(0)
	v188 = v2
	v189 = v2
	goto L24
L31:
	;
	goto L32
L32:
	;
	F_deconstruct_array_builtin(m, v54, int32(25), v12+int32(12), v12+int32(40), v12+int32(32))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	F_deconstruct_array_builtin(m, v57, int32(25), v12+int32(44), v12+int32(36), v12+int32(28))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	if v85 != v86 {
		goto L22
	} else {
		goto L35
	}
L35:
	;
	v90 = F_palloc(m, v85<<(uint(int32(2))%32))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v95 = F_palloc(m, v92<<(uint(int32(2))%32))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v97 <= int32(0) {
		v185 = v97
		v188 = v90
		v189 = v95
		goto L24
	} else {
		goto L38
	}
L38:
	;
	v103 = int32(0)
	goto L39
L39:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+v103))))
	if v112 == int32(1) {
		goto L21
	} else {
		goto L41
	}
L40:
	;
	v185 = v180
	v188 = v90
	v189 = v95
	goto L24
L41:
	;
	v116 = v103 << (uint(int32(2)) % 32)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116+v117)))
	v120 = F_text_to_cstring(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v123 = v120
	goto L44
L43:
	;
	if base.B2i32(v125 == int32(0)) == int32(0) {
		goto L20
	} else {
		goto L47
	}
L44:
	;
	v125 = int32(*(*int8)(unsafe.Add(mBase, uint32(v123))))
	if int32(0) < v125 {
		v123 = v123 + int32(1)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	goto L43
L46:
	;
	goto L45
L47:
	;
	v135 = F_strstr(m, v120, int32(_a_F_pg_armor_0))
	mBase = m.M
	if v135 != 0 {
		goto L19
	} else {
		goto L48
	}
L48:
	;
	v136 = int32(10)
	v137 = F___strchrnul(m, v120, v136)
	mBase = m.M
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v139 == v136 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v143 != 0 {
		goto L18
	} else {
		goto L53
	}
L50:
	;
	v143 = v137
	goto L52
L51:
	;
	v143 = int32(0)
	goto L52
L52:
	;
	goto L49
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116+v90))) = v120
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146+v103))))
	if v148 == int32(1) {
		goto L17
	} else {
		goto L54
	}
L54:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v151+v116)))
	v154 = F_text_to_cstring(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v157 = v154
	goto L57
L56:
	;
	if base.B2i32(v159 == int32(0)) == int32(0) {
		goto L16
	} else {
		goto L60
	}
L57:
	;
	v159 = int32(*(*int8)(unsafe.Add(mBase, uint32(v157))))
	if int32(0) < v159 {
		v157 = v157 + int32(1)
		goto L57
	} else {
		goto L59
	}
L58:
	;
	goto L56
L59:
	;
	goto L58
L60:
	;
	v168 = int32(10)
	v169 = F___strchrnul(m, v154, v168)
	mBase = m.M
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if v171 == v168 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v175 != 0 {
		goto L15
	} else {
		goto L65
	}
L62:
	;
	v175 = v169
	goto L64
L63:
	;
	v175 = int32(0)
	goto L64
L64:
	;
	goto L61
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116+v95))) = v154
	v179 = v103 + int32(1)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v179 < v180 {
		v103 = v179
		goto L39
	} else {
		goto L66
	}
L66:
	;
	goto L40
L67:
	;
	v195 = int32(1)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v197&v195 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v200 = v195
	goto L70
L69:
	;
	v200 = int32(4)
	goto L70
L70:
	;
	F_pgp_armor_encode(m, v15+v200, v49, v12+int32(12), v185, v188, v189)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v209 = F_palloc(m, v206+int32(4))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L72
	}
L72:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v211<<(uint(int32(2))%32) + int32(16)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v220 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	F_pfree(m, v219)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L2
	} else {
		goto L77
	}
L74:
	;
	v221 = F__emscripten_memcpy_bulkmem(m, v209+int32(4), v219, v220)
	mBase = m.M
	goto L76
L75:
	;
	goto L76
L76:
	;
	goto L73
L77:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v225 != v15 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	F_pfree(m, v15)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L2
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	m.G0 = v12 + int32(48)
	return v209
L81:
	;
	goto L80
L82:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(_a_F_pg_armor_1), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(775), int32(_a_F_pg_armor_3))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L2
	} else {
		goto L87
	}
L87:
	;
	F_errmsg(m, int32(_a_F_pg_armor_4), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L2
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(785), int32(_a_F_pg_armor_3))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L2
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
	F_errcode(m, int32(67108994))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L2
	} else {
		goto L91
	}
L91:
	;
	F_errmsg(m, int32(_a_F_pg_armor_5), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L2
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(798), int32(_a_F_pg_armor_3))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(_a_F_pg_armor_6), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(805), int32(_a_F_pg_armor_3))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L2
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L2
	} else {
		goto L99
	}
L99:
	;
	F_errmsg(m, int32(_a_F_pg_armor_7), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L2
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(809), int32(_a_F_pg_armor_3))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L2
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L2
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(_a_F_pg_armor_8), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L2
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(813), int32(_a_F_pg_armor_3))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L2
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
	F_errcode(m, int32(67108994))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L2
	} else {
		goto L107
	}
L107:
	;
	F_errmsg(m, int32(_a_F_pg_armor_9), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L2
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(820), int32(_a_F_pg_armor_3))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L2
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(_a_F_pg_armor_10), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(827), int32(_a_F_pg_armor_3))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L2
	} else {
		goto L115
	}
L115:
	;
	F_errmsg(m, int32(_a_F_pg_armor_11), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L2
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(831), int32(_a_F_pg_armor_3))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L2
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	v417 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v417
	F_errmsg_internal(m, int32(_a_F_pg_armor_12), v12)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L2
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_pg_armor_2), int32(863), int32(_a_F_pg_armor_13))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L2
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_available_wal_summaries(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+2)) = uint8(v16)
	*(*uint16)(unsafe.Add(mBase, uint32(v8))) = uint16(v16)
	v21 = int64(0)
	v23 = F_GetWalSummaries(m, v16, v21, v21)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v8 + int32(16)
	return int32(0)
L4:
	;
	if v23 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v27 <= int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v32 = v2
	goto L7
L7:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v32<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_pg_available_wal_summaries[0]))
	if v41 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L3
L9:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v44 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v39)+16)))
	v45 = F_Int64GetDatum(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v45
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v39)))
	v49 = F_Int64GetDatum(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v49
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
	v53 = F_Int64GetDatum(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v53
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v59 = F_heap_form_tuple(m, v56, v8+int32(4), v8)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	F_tuplestore_puttuple(m, v61, v59)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v65 = v32 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v65 < v66 {
		v32 = v65
		goto L7
	} else {
		goto L18
	}
L18:
	;
	goto L8
}
func F_pg_b64_enc_len(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = int32(2)
	v5 = base.I32_div_s(l0+v2, int32(3))
	return v5 << (uint(v2) % 32)
}
func F_pg_b64_encode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
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
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v127 int32
	_ = v127
	v8 = l0 + l1
	if base.Ui32(l0) < base.Ui32(v8) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v127 = F__emscripten_memset_bulkmem(m, l2, base.I32_extend8_s(int32(0)), l3)
	mBase = m.M
	goto L18
L2:
	;
	if l3 < v67-l2+int32(4) {
		goto L1
	} else {
		goto L14
	}
L3:
	;
	v12 = l0
	v13 = int32(0)
	v16 = l2
	v17 = int32(2)
	goto L6
L4:
	;
	v78 = l2
	goto L5
L5:
	;
	return v78 - l2
L6:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v23 = v19<<(uint(v17<<(uint(int32(3))%32))%32) | v13
	if int32(0) < v17 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v68 != int32(2) {
		goto L2
	} else {
		goto L13
	}
L8:
	;
	v66 = v23
	v67 = v16
	v68 = v17 - int32(1)
	goto L10
L9:
	;
	if l3 < v16-l2+int32(4) {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v70 = v12 + int32(1)
	if v70 != v8 {
		v12 = v70
		v13 = v66
		v16 = v67
		v17 = v68
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v32 = int32(63)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23&v32)+uint32(_c_F_pg_b64_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+3)) = uint8(v36)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v23)>>(uint(int32(6))%32))&v32)+uint32(_c_F_pg_b64_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+2)) = uint8(v44)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v23)>>(uint(int32(12))%32))&v32)+uint32(_c_F_pg_b64_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)) = uint8(v52)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v23)>>(uint(int32(18))%32))&v32)+uint32(_c_F_pg_b64_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v60)
	v66 = int32(0)
	v67 = v16 + int32(4)
	v68 = int32(2)
	goto L10
L12:
	;
	goto L7
L13:
	;
	v78 = v67
	goto L5
L14:
	;
	v89 = int32(63)
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v66)>>(uint(int32(12))%32))&v89)+uint32(_c_F_pg_b64_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)) = uint8(v93)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v66)>>(uint(int32(18))%32))&v89)+uint32(_c_F_pg_b64_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v101)
	if v68 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v66)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_pg_b64_encode[0]))))
	v113 = v112
	goto L17
L16:
	;
	v113 = int32(61)
	goto L17
L17:
	;
	v114 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+3)) = uint8(v114)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+2)) = uint8(v113)
	return v67 + int32(4) - l2
L18:
	;
	return int32(-1)
}
func F_pg_backup_start(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_backup_start[0])))
		v13 = F_text_to_cstring(m, v6)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			if v12 != int32(1) {
				v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[1]))
				if v18 == int32(0) {
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[2]))
					v28 = F_AllocSetContextCreateInternal(m, v23, int32(_a_F_pg_backup_start_0), int32(0), int32(1024), int32(_a_F_pg_backup_start_1))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[1])) = v28
						v41 = v28
						v42 = int32(_a_F_pg_backup_start_2)
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[3]))
						*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[3])) = v41
						v48 = F_palloc0(m, int32(1112))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4])) = v48
							v51 = F_makeStringInfo(m)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[3])) = v43
								*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[5])) = v51
								v58 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_backup_start[6])))
								if v58 == int32(0) {
									F_before_shmem_exit(m, int32(407), int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										v66 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_backup_start[6])) = uint8(v66)
										v68 = int32(0)
										v72 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4]))
										v74 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[5]))
										F_do_pg_backup_start(m, v13, base.B2i32(v10 != v68), v68, v72, v74)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4]))
											v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+1032))
											v80 = F_Int64GetDatum(m, v79)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												return v80
											}
										}
									}
								} else {
									v68 = int32(0)
									v72 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4]))
									v74 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[5]))
									F_do_pg_backup_start(m, v13, base.B2i32(v10 != v68), v68, v72, v74)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4]))
										v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+1032))
										v80 = F_Int64GetDatum(m, v79)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											return v80
										}
									}
								}
							}
						}
					}
				} else {
					v32 = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[5])) = v32
					*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4])) = v32
					F_MemoryContextReset(m, v18)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[1]))
						v41 = v40
						v42 = int32(_a_F_pg_backup_start_2)
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[3]))
						*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[3])) = v41
						v48 = F_palloc0(m, int32(1112))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4])) = v48
							v51 = F_makeStringInfo(m)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[3])) = v43
								*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[5])) = v51
								v58 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_backup_start[6])))
								if v58 == int32(0) {
									F_before_shmem_exit(m, int32(407), int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										v66 = int32(1)
										*(*uint8)(unsafe.Add(mBase, _c_F_pg_backup_start[6])) = uint8(v66)
										v68 = int32(0)
										v72 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4]))
										v74 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[5]))
										F_do_pg_backup_start(m, v13, base.B2i32(v10 != v68), v68, v72, v74)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4]))
											v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+1032))
											v80 = F_Int64GetDatum(m, v79)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												return v80
											}
										}
									}
								} else {
									v68 = int32(0)
									v72 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4]))
									v74 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[5]))
									F_do_pg_backup_start(m, v13, base.B2i32(v10 != v68), v68, v72, v74)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_start[4]))
										v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+1032))
										v80 = F_Int64GetDatum(m, v79)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return int32(0)
										} else {
											return v80
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
				v86 = m.ExcPending
				if v86 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_pg_backup_start_3), int32(0))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_backup_start_4), int32(70), int32(_a_F_pg_backup_start_5))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
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
func F_pg_base64_decode_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v164 int32
	_ = v164
	v4 = int32(0)
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v164
L2:
	;
	v11 = l0 + l1
	v12 = l0
	v15 = l2
	v17 = v4
	v18 = v4
	v19 = v4
	goto L5
L3:
	;
	v148 = l2
	goto L4
L4:
	;
	v164 = v148 - l2
	goto L1
L5:
	;
	v23 = v12
	goto L14
L6:
	;
	if v142 != 0 {
		v164 = int32(-101)
		goto L1
	} else {
		goto L39
	}
L7:
	;
	goto L6
L8:
	;
	if base.Ui32(v33) < base.Ui32(v11) {
		v12 = v33
		v15 = v130
		v17 = v132
		v18 = v133
		v19 = v134
		goto L5
	} else {
		goto L38
	}
L9:
	;
	v116 = int32(0)
	if base.B2i32(v111 == v116)&base.B2i32(base.Ui32(v112) < base.Ui32(int32(3))) == v116 {
		goto L35
	} else {
		goto L36
	}
L10:
	;
	v106 = int32(base.Ui32(v100) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v106)
	v110 = v100
	v111 = v101
	v112 = v102
	v115 = v15 + int32(2)
	goto L9
L11:
	;
	v94 = int32(base.Ui32(v18) >> (uint(int32(10)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v94)
	v100 = v18 << (uint(int32(6)) % 32)
	v101 = int32(0)
	v102 = int32(2)
	goto L10
L12:
	;
	v76 = v72 + v18<<(uint(int32(6))%32)
	v78 = v19 + int32(1)
	if v78 != int32(4) {
		goto L30
	} else {
		goto L31
	}
L13:
	;
	v72 = int32(62)
	goto L12
L14:
	;
	v33 = v23 + int32(1)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v36 = v34 - int32(65)
	if base.Ui32(v36&int32(255)) <= base.Ui32(int32(25)) {
		v72 = v36
		goto L12
	} else {
		goto L16
	}
L15:
	;
	if v17 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	if base.Ui32((v34-int32(97))&int32(255)) <= base.Ui32(int32(25)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v72 = v34 - int32(71)
	goto L12
L18:
	;
	goto L19
L19:
	;
	if base.Ui32((v34-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v72 = (v34 + int32(4)) & int32(255)
	goto L12
L21:
	;
	goto L22
L22:
	;
	v60 = int32(-101)
	switch v34 - int32(9) {
	case 0, 1, 4, 23:
		goto L24
	default:
		v164 = v60
		goto L1
	case 34:
		goto L13
	case 38:
		v72 = int32(63)
		goto L12
	case 52:
		goto L23
	}
L23:
	;
	goto L15
L24:
	;
	if base.Ui32(v33) < base.Ui32(v11) {
		v23 = v33
		goto L14
	} else {
		goto L25
	}
L25:
	;
	v138 = v15
	v142 = v19
	goto L7
L26:
	;
	v72 = int32(0)
	goto L12
L27:
	;
	goto L28
L28:
	;
	switch v19 - int32(2) {
	case 0:
		goto L29
	case 1:
		goto L11
	default:
		v164 = v60
		goto L1
	}
L29:
	;
	v130 = v15
	v132 = int32(1)
	v133 = v18 << (uint(int32(6)) % 32)
	v134 = int32(3)
	goto L8
L30:
	;
	v130 = v15
	v132 = v17
	v133 = v76
	v134 = v78
	goto L8
L31:
	;
	goto L32
L32:
	;
	v82 = int32(base.Ui32(v76) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v82)
	v85 = base.B2i32(v17 == int32(0))
	if v17 == int32(0) {
		v100 = v76
		v101 = v85
		v102 = v17
		goto L10
	} else {
		goto L33
	}
L33:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v17) {
		v100 = v76
		v101 = v85
		v102 = v17
		goto L10
	} else {
		goto L34
	}
L34:
	;
	v110 = v76
	v111 = int32(0)
	v112 = v17
	v115 = v15 + int32(1)
	goto L9
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v110)
	v127 = v115 + int32(1)
	goto L37
L36:
	;
	v127 = v115
	goto L37
L37:
	;
	v130 = v127
	v132 = v112
	v133 = v116
	v134 = int32(0)
	goto L8
L38:
	;
	v138 = v130
	v142 = v134
	goto L7
L39:
	;
	v148 = v138
	goto L4
}
func F_pg_class_aclcheck_ext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_pg_class_aclmask_ext(m, l0, l1, l2, int32(1), l3)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 == int64(0))
	}
}
func F_pg_class_aclmask_ext(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v137 int64
	_ = v137
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v153
L2:
	;
	return int64(0)
L3:
	;
	if v17 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l4 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+22)))
	v44 = v42 + v43
	if l2&int64(285) == int64(0) {
		v69 = l2
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v23 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v23)
	v153 = int64(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg(m, int32(_a_F_pg_class_aclmask_ext_0), v14)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_pg_class_aclmask_ext_1), int32(3310), int32(_a_F_pg_class_aclmask_ext_2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	v70 = F_superuser_arg(m, l1)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L26
	}
L15:
	;
	v50 = int32(1)
	if base.Ui32(l0) < base.Ui32(int32(_a_F_pg_class_aclmask_ext_6)) {
		v58 = v50
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v58 == int32(0) {
		v69 = l2
		goto L14
	} else {
		goto L20
	}
L17:
	;
	goto L16
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)+68))
	if v53 == int32(99) {
		v58 = v50
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v56 = F_isTempToastNamespace(m, v53)
	mBase = m.M
	v58 = v56
	goto L17
L20:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+119)))
	if v61 == int32(118) {
		v69 = l2
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v66 = F_superuser_arg(m, l1)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	if v66 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v68 = l2
	goto L25
L24:
	;
	v68 = l2 & int64(-286)
	goto L25
L25:
	;
	v69 = v68
	goto L14
L26:
	;
	if v70 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_ReleaseCatCache(m, v17)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L2
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	v79 = F_SysCacheGetAttr(m, int32(57), v17, int32(32), v14+int32(15))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L31
	}
L30:
	;
	v153 = v69
	goto L1
L31:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
	if v81 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v99 = F_aclmask(m, v98, l1, v74, v69, l3)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L2
	} else {
		goto L42
	}
L33:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+119)))
	if v84 == int32(83) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v95 = F_pg_detoast_datum(m, v79)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L41
	}
L36:
	;
	v89 = F_acldefault(m, int32(37), v74)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v93 = F_acldefault(m, int32(41), v74)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L2
	} else {
		goto L40
	}
L39:
	;
	v97 = int32(0)
	v98 = v89
	goto L32
L40:
	;
	v97 = int32(0)
	v98 = v93
	goto L32
L41:
	;
	v97 = v79
	v98 = v95
	goto L32
L42:
	;
	if v98 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_ReleaseCatCache(m, v17)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L47
	}
L44:
	;
	if v98 == v97 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	F_pfree(m, v98)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	if v69&int64(2) == int64(0) {
		v122 = v99
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v124 = v69 & int64(13)
	if v124 == int64(0) {
		v137 = v122
		goto L55
	} else {
		goto L56
	}
L49:
	;
	if v99&int64(2) != int64(0) {
		v122 = v99
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v119 = F_has_privs_of_role(m, l1, int32(_a_F_pg_class_aclmask_ext_5))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	if v119 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v121 = v99 | int64(2)
	goto L54
L53:
	;
	v121 = v99
	goto L54
L54:
	;
	v122 = v121
	goto L48
L55:
	;
	if v69&int64(16384) == int64(0) {
		goto L62
	} else {
		goto L63
	}
L56:
	;
	if v122&int64(13) != int64(0) {
		v137 = v122
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v133 = F_has_privs_of_role(m, l1, int32(_a_F_pg_class_aclmask_ext_4))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	if v133 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v135 = v124
	goto L61
L60:
	;
	v135 = int64(0)
	goto L61
L61:
	;
	v137 = v135 | v122
	goto L55
L62:
	;
	v153 = v137
	goto L1
L63:
	;
	goto L64
L64:
	;
	if v137&int64(16384) != int64(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v153 = v137
	goto L1
L66:
	;
	goto L67
L67:
	;
	v149 = F_has_privs_of_role(m, l1, int32(_a_F_pg_class_aclmask_ext_3))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	if v149 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v151 = v137 | int64(16384)
	goto L71
L70:
	;
	v151 = v137
	goto L71
L71:
	;
	v153 = v151
	goto L1
}
func F_pg_collation_actual_version(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v10 == int32(100) {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_pg_collation_actual_version[0]))
		v16 = F_SearchSysCache1(m, int32(21), v15)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v16 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, _c_F_pg_collation_actual_version[0]))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v79
						F_errmsg(m, int32(_a_F_pg_collation_actual_version_0), v8)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_collation_actual_version_1), int32(524), int32(_a_F_pg_collation_actual_version_2))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25+v26)+76)))
				if v28 == int32(99) {
					v31 = int32(13)
				} else {
					v31 = int32(15)
				}
				v32 = F_SysCacheGetAttrNotNull(m, int32(21), v16, v31)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v51 = v16
					v52 = v28
					v53 = v32
					v54 = F_text_to_cstring(m, v53)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v51)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v59 = F_get_collation_actual_version(m, base.I32_extend8_s(v52), v54)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								if v59 != 0 {
									v61 = F_cstring_to_text(m, v59)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										v66 = v61
										m.G0 = v8 + int32(32)
										return v66
									}
								} else {
									v63 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v63)
									v66 = int32(0)
									m.G0 = v8 + int32(32)
									return v66
								}
							}
						}
					}
				}
			}
		}
	} else {
		v35 = F_SearchSysCache1(m, int32(16), v10)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			if v35 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
						F_errmsg(m, int32(_a_F_pg_collation_actual_version_3), v8+int32(16))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_collation_actual_version_1), int32(550), int32(_a_F_pg_collation_actual_version_2))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+22)))
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v43)+76)))
				if v45 == int32(99) {
					v48 = int32(8)
				} else {
					v48 = int32(10)
				}
				v49 = F_SysCacheGetAttrNotNull(m, int32(16), v35, v48)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					v51 = v35
					v52 = v45
					v53 = v49
					v54 = F_text_to_cstring(m, v53)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v51)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v59 = F_get_collation_actual_version(m, base.I32_extend8_s(v52), v54)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								if v59 != 0 {
									v61 = F_cstring_to_text(m, v59)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										v66 = v61
										m.G0 = v8 + int32(32)
										return v66
									}
								} else {
									v63 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v63)
									v66 = int32(0)
									m.G0 = v8 + int32(32)
									return v66
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_collation_is_visible(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_SearchSysCache1(m, int32(16), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v8 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v74 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v74)
	return int32(0)
L6:
	;
	v16 = v12 + v13
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
	if v17 != int32(11) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	F_ReleaseCatCache(m, v8)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L26
	}
L8:
	;
	v20 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_collation_is_visible[0]))
	if v22 == v20 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	v67 = F_CollationGetCollid(m, v16+int32(4))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L25
	}
L11:
	;
	if v61 == int32(0) {
		v70 = v20
		goto L7
	} else {
		goto L24
	}
L12:
	;
	v61 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v29 <= int32(0) {
		v54 = v20
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v61 = v54
	goto L11
L16:
	;
	v32 = int32(0)
	if v32 < v29 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v35 = v29
	goto L19
L18:
	;
	v35 = v32
	goto L19
L19:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v38 = int32(0)
	goto L20
L20:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v36+v38<<(uint(int32(2))%32))))
	v47 = base.B2i32(v46 == v17)
	if v46 == v17 {
		v54 = v47
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v54 = v47
	goto L15
L22:
	;
	v49 = v38 + int32(1)
	if v49 != v35 {
		v38 = v49
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L10
L25:
	;
	v70 = base.B2i32(v67 == v7)
	goto L7
L26:
	;
	return v70
}
func F_pg_column_compression(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
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
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v11 == int32(0) {
		v15 = F_get_fn_expr_argtype(m, v10, int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = F_get_typlen(m, v15)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v114 = m.ExcPending
					if v114 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
						F_errmsg_internal(m, int32(_a_F_pg_column_compression_0), v8)
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_column_compression_1), int32(_a_F_pg_column_compression_2), int32(_a_F_pg_column_compression_3))
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
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
					v26 = F_MemoryContextAlloc(m, v24, int32(4))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v26
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = v19
						v35 = v19
						if v35 != int32(-1) {
							v38 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
							v103 = int32(0)
							m.G0 = v8 + int32(32)
							return v103
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
							if v43 == int32(1) {
								v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
								if v47 != int32(18) {
									v70 = int32(2)
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+6))
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+2))
									if base.Ui32(v54-int32(4)) <= base.Ui32(v51&int32(1073741823)) {
										v60 = int32(2)
									} else {
										v60 = int32(base.Ui32(v51) >> (uint(int32(30)) % 32))
									}
									v70 = v60
								}
							} else {
								v61 = int32(2)
								if v43&int32(3) != v61 {
									v70 = v61
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
									v70 = int32(base.Ui32(v66) >> (uint(int32(30)) % 32))
								}
							}
							switch v70 {
							case 0:
								v90 = int32(_a_F_pg_column_compression_4)
								v91 = F_strlen(m, v90)
								mBase = m.M
								v93 = v91 + int32(4)
								v94 = F_palloc(m, v93)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v94))) = v93 << (uint(int32(2)) % 32)
									if v91 != 0 {
										v101 = F__emscripten_memcpy_bulkmem(m, v94+int32(4), v90, v91)
										mBase = m.M
									} else {
									}
									v103 = v94
									m.G0 = v8 + int32(32)
									return v103
								}
							case 1:
								v90 = int32(_a_F_pg_column_compression_5)
								v91 = F_strlen(m, v90)
								mBase = m.M
								v93 = v91 + int32(4)
								v94 = F_palloc(m, v93)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v94))) = v93 << (uint(int32(2)) % 32)
									if v91 != 0 {
										v101 = F__emscripten_memcpy_bulkmem(m, v94+int32(4), v90, v91)
										mBase = m.M
									} else {
									}
									v103 = v94
									m.G0 = v8 + int32(32)
									return v103
								}
							case 2:
								v71 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v71)
								v103 = int32(0)
								m.G0 = v8 + int32(32)
								return v103
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
									F_errmsg_internal(m, int32(_a_F_pg_column_compression_6), v8+int32(16))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_pg_column_compression_1), int32(_a_F_pg_column_compression_7), int32(_a_F_pg_column_compression_3))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
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
	} else {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		v35 = v33
		if v35 != int32(-1) {
			v38 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
			v103 = int32(0)
			m.G0 = v8 + int32(32)
			return v103
		} else {
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
			if v43 == int32(1) {
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
				if v47 != int32(18) {
					v70 = int32(2)
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+6))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+2))
					if base.Ui32(v54-int32(4)) <= base.Ui32(v51&int32(1073741823)) {
						v60 = int32(2)
					} else {
						v60 = int32(base.Ui32(v51) >> (uint(int32(30)) % 32))
					}
					v70 = v60
				}
			} else {
				v61 = int32(2)
				if v43&int32(3) != v61 {
					v70 = v61
				} else {
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
					v70 = int32(base.Ui32(v66) >> (uint(int32(30)) % 32))
				}
			}
			switch v70 {
			case 0:
				v90 = int32(_a_F_pg_column_compression_4)
				v91 = F_strlen(m, v90)
				mBase = m.M
				v93 = v91 + int32(4)
				v94 = F_palloc(m, v93)
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v94))) = v93 << (uint(int32(2)) % 32)
					if v91 != 0 {
						v101 = F__emscripten_memcpy_bulkmem(m, v94+int32(4), v90, v91)
						mBase = m.M
					} else {
					}
					v103 = v94
					m.G0 = v8 + int32(32)
					return v103
				}
			case 1:
				v90 = int32(_a_F_pg_column_compression_5)
				v91 = F_strlen(m, v90)
				mBase = m.M
				v93 = v91 + int32(4)
				v94 = F_palloc(m, v93)
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v94))) = v93 << (uint(int32(2)) % 32)
					if v91 != 0 {
						v101 = F__emscripten_memcpy_bulkmem(m, v94+int32(4), v90, v91)
						mBase = m.M
					} else {
					}
					v103 = v94
					m.G0 = v8 + int32(32)
					return v103
				}
			case 2:
				v71 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v71)
				v103 = int32(0)
				m.G0 = v8 + int32(32)
				return v103
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
					F_errmsg_internal(m, int32(_a_F_pg_column_compression_6), v8+int32(16))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_column_compression_1), int32(_a_F_pg_column_compression_7), int32(_a_F_pg_column_compression_3))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
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
func F_pg_copy_physical_replication_slot_a(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_copy_replication_slot(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pg_cryptohash_update(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	if l0 == int32(0) {
		return int32(-1)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = m.Env.Pgmem_hash_update(m, v8, l1, l2)
		mBase = m.M
		return int32(0)
	}
}
func F_pg_ddl_command_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(_a_F_pg_ddl_command_recv_0)
			F_errmsg(m, int32(_a_F_pg_ddl_command_recv_1), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_pg_ddl_command_recv_2), int32(359), int32(_a_F_pg_ddl_command_recv_3))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_pg_decrypt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
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
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L2
	} else {
		goto L94
	}
L2:
	;
	return int32(0)
L3:
	;
	v19 = int32(1)
	v20 = v15 + v19
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v25 = v23 & v19
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = v20
	goto L6
L5:
	;
	v26 = v15 + int32(4)
	goto L6
L6:
	;
	if v23 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v56 = F_downcase_truncate_identifier(m, v26, v54, int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L18
	}
L8:
	;
	v29 = int32(4)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v31&int32(254) == int32(2) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v44 = int32(1)
	if v25 != 0 {
		v54 = int32(base.Ui32(v23)>>(uint(v44)%32)) - v44
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v40 = v29
	goto L13
L12:
	;
	v40 = base.B2i32(v31 == int32(18)) << (uint(v29) % 32)
	goto L13
L13:
	;
	if v31 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v43 = v29
	goto L16
L15:
	;
	v43 = v40
	goto L16
L16:
	;
	v54 = v43
	goto L7
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v60 = F_px_find_combo(m, v56, v12+int32(28))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v60 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_pfree(m, v56)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L2
	} else {
		goto L76
	}
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v68 = F_pg_detoast_datum_packed(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v71 = F_pg_detoast_datum_packed(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v73 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v104 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L27:
	;
	v76 = int32(4)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	if v78&int32(254) == int32(2) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v91 = int32(1)
	if v73&v91 != 0 {
		v103 = int32(base.Ui32(v73)>>(uint(v91)%32)) - v91
		goto L26
	} else {
		goto L36
	}
L30:
	;
	v87 = v76
	goto L32
L31:
	;
	v87 = base.B2i32(v78 == int32(18)) << (uint(v76) % 32)
	goto L32
L32:
	;
	if v78 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v90 = v76
	goto L35
L34:
	;
	v90 = v87
	goto L35
L35:
	;
	v103 = v90
	goto L26
L36:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v103 = int32(base.Ui32(v97)>>(uint(int32(2))%32)) - int32(4)
	goto L26
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	v136 = m.T0[v135].(func(*base.Module, int32, int32) int32)(m, v66, v103)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L2
	} else {
		goto L48
	}
L38:
	;
	v107 = int32(4)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	if v109&int32(254) == int32(2) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v122 = int32(1)
	if v104&v122 != 0 {
		v134 = int32(base.Ui32(v104)>>(uint(v122)%32)) - v122
		goto L37
	} else {
		goto L47
	}
L41:
	;
	v118 = v107
	goto L43
L42:
	;
	v118 = base.B2i32(v109 == int32(18)) << (uint(v107) % 32)
	goto L43
L43:
	;
	if v109 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v121 = v107
	goto L46
L45:
	;
	v121 = v118
	goto L46
L46:
	;
	v134 = v121
	goto L37
L47:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v134 = int32(base.Ui32(v128)>>(uint(int32(2))%32)) - int32(4)
	goto L37
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v136
	v141 = F_palloc(m, v136+int32(4))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v143 = int32(1)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v145&v143 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v148 = v143
	goto L52
L51:
	;
	v148 = int32(4)
	goto L52
L52:
	;
	v150 = int32(0)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v153 = m.T0[v152].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v66, v71+v148, v134, v150, v150)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	if v153 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	m.T0[v155].(func(*base.Module, int32))(m, v66)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v158 = int32(1)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v160&v158 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v254 = v153
	goto L1
L58:
	;
	v163 = v158
	goto L60
L59:
	;
	v163 = int32(4)
	goto L60
L60:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	v170 = m.T0[v169].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v66, v68+v163, v103, v141+int32(4), v12+int32(28))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	m.T0[v172].(func(*base.Module, int32))(m, v66)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	if v170 != 0 {
		v254 = v170
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v175<<(uint(int32(2))%32) + int32(16)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v181 != v68 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	F_pfree(m, v68)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v185 != v71 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L66
L68:
	;
	F_pfree(m, v71)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L2
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v189 != v15 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L70
L72:
	;
	F_pfree(m, v15)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L2
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	m.G0 = v12 + int32(32)
	return v141
L75:
	;
	goto L74
L76:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	if v60 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v237
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v56
	F_errmsg(m, int32(_a_F_pg_decrypt_0), v12+int32(16))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L2
	} else {
		goto L92
	}
L79:
	;
	v237 = int32(_a_F_pg_decrypt_1)
	goto L78
L80:
	;
	goto L81
L81:
	;
	v214 = int32(_a_F_pg_decrypt_2)
	goto L83
L82:
	;
	v237 = v232
	goto L78
L83:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v214)+8))
	if v60 != v217 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	v232 = v229
	goto L82
L85:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v214)+20))
	if v220 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	goto L84
L88:
	;
	v237 = int32(_a_F_pg_decrypt_3)
	goto L78
L89:
	;
	goto L90
L90:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v214)+16))
	if v60 != v225 {
		v214 = v214 + int32(16)
		goto L83
	} else {
		goto L91
	}
L91:
	;
	v232 = v220
	goto L82
L92:
	;
	F_errfinish(m, int32(_a_F_pg_decrypt_4), int32(513), int32(_a_F_pg_decrypt_5))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L2
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	F_errcode(m, int32(579))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L2
	} else {
		goto L95
	}
L95:
	;
	if v254 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v295
	F_errmsg(m, int32(_a_F_pg_decrypt_6), v12)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L2
	} else {
		goto L110
	}
L97:
	;
	v295 = int32(_a_F_pg_decrypt_1)
	goto L96
L98:
	;
	goto L99
L99:
	;
	v272 = int32(_a_F_pg_decrypt_2)
	goto L101
L100:
	;
	v295 = v290
	goto L96
L101:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v272)+8))
	if v254 != v275 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
	v290 = v287
	goto L100
L103:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v272)+20))
	if v278 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	goto L102
L106:
	;
	v295 = int32(_a_F_pg_decrypt_3)
	goto L96
L107:
	;
	goto L108
L108:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v272)+16))
	if v254 != v283 {
		v272 = v272 + int32(16)
		goto L101
	} else {
		goto L109
	}
L109:
	;
	v290 = v278
	goto L100
L110:
	;
	F_errfinish(m, int32(_a_F_pg_decrypt_4), int32(334), int32(_a_F_pg_decrypt_7))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_decrypt_iv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
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
	var v73 int32
	_ = v73
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
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L2
	} else {
		goto L113
	}
L2:
	;
	return int32(0)
L3:
	;
	v21 = int32(1)
	v22 = v17 + v21
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v27 = v25 & v21
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = v22
	goto L6
L5:
	;
	v28 = v17 + int32(4)
	goto L6
L6:
	;
	if v25 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v58 = F_downcase_truncate_identifier(m, v28, v56, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L18
	}
L8:
	;
	v31 = int32(4)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v33&int32(254) == int32(2) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v46 = int32(1)
	if v27 != 0 {
		v56 = int32(base.Ui32(v25)>>(uint(v46)%32)) - v46
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v42 = v31
	goto L13
L12:
	;
	v42 = base.B2i32(v33 == int32(18)) << (uint(v31) % 32)
	goto L13
L13:
	;
	if v33 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v45 = v31
	goto L16
L15:
	;
	v45 = v42
	goto L16
L16:
	;
	v56 = v45
	goto L7
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v56 = int32(base.Ui32(v50)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v62 = F_px_find_combo(m, v58, v14+int32(28))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v62 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_pfree(m, v58)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L2
	} else {
		goto L95
	}
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v70 = F_pg_detoast_datum_packed(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v73 = F_pg_detoast_datum_packed(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v76 = F_pg_detoast_datum_packed(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v78 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v109 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v81 = int32(4)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v83&int32(254) == int32(2) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v96 = int32(1)
	if v78&v96 != 0 {
		v108 = int32(base.Ui32(v78)>>(uint(v96)%32)) - v96
		goto L27
	} else {
		goto L37
	}
L31:
	;
	v92 = v81
	goto L33
L32:
	;
	v92 = base.B2i32(v83 == int32(18)) << (uint(v81) % 32)
	goto L33
L33:
	;
	if v83 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v95 = v81
	goto L36
L35:
	;
	v95 = v92
	goto L36
L36:
	;
	v108 = v95
	goto L27
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v108 = int32(base.Ui32(v102)>>(uint(int32(2))%32)) - int32(4)
	goto L27
L38:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v140 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L39:
	;
	v112 = int32(4)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	if v114&int32(254) == int32(2) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v127 = int32(1)
	if v109&v127 != 0 {
		v139 = int32(base.Ui32(v109)>>(uint(v127)%32)) - v127
		goto L38
	} else {
		goto L48
	}
L42:
	;
	v123 = v112
	goto L44
L43:
	;
	v123 = base.B2i32(v114 == int32(18)) << (uint(v112) % 32)
	goto L44
L44:
	;
	if v114 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v126 = v112
	goto L47
L46:
	;
	v126 = v123
	goto L47
L47:
	;
	v139 = v126
	goto L38
L48:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v139 = int32(base.Ui32(v133)>>(uint(int32(2))%32)) - int32(4)
	goto L38
L49:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	v172 = m.T0[v171].(func(*base.Module, int32, int32) int32)(m, v68, v108)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L2
	} else {
		goto L60
	}
L50:
	;
	v143 = int32(4)
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	if v145&int32(254) == int32(2) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v158 = int32(1)
	if v140&v158 != 0 {
		v170 = int32(base.Ui32(v140)>>(uint(v158)%32)) - v158
		goto L49
	} else {
		goto L59
	}
L53:
	;
	v154 = v143
	goto L55
L54:
	;
	v154 = base.B2i32(v145 == int32(18)) << (uint(v143) % 32)
	goto L55
L55:
	;
	if v145 == int32(1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v157 = v143
	goto L58
L57:
	;
	v157 = v154
	goto L58
L58:
	;
	v170 = v157
	goto L49
L59:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v170 = int32(base.Ui32(v164)>>(uint(int32(2))%32)) - int32(4)
	goto L49
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v172
	v177 = F_palloc(m, v172+int32(4))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v179 = int32(1)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if v181&v179 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v184 = v179
	goto L64
L63:
	;
	v184 = int32(4)
	goto L64
L64:
	;
	v186 = int32(1)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v188&v186 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v191 = v186
	goto L67
L66:
	;
	v191 = int32(4)
	goto L67
L67:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v194 = m.T0[v193].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v68, v73+v184, v139, v76+v191, v170)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	if v194 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	m.T0[v196].(func(*base.Module, int32))(m, v68)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L2
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v199 = int32(1)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v201&v199 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v299 = v194
	goto L1
L73:
	;
	v204 = v199
	goto L75
L74:
	;
	v204 = int32(4)
	goto L75
L75:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	v211 = m.T0[v210].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v68, v70+v204, v108, v177+int32(4), v14+int32(28))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	m.T0[v213].(func(*base.Module, int32))(m, v68)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	if v211 != 0 {
		v299 = v211
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = v216<<(uint(int32(2))%32) + int32(16)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v222 != v70 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	F_pfree(m, v70)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L2
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v226 != v73 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	F_pfree(m, v73)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L2
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v230 != v76 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	F_pfree(m, v76)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v234 != v17 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L89
L91:
	;
	F_pfree(m, v17)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L2
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	m.G0 = v14 + int32(32)
	return v177
L94:
	;
	goto L93
L95:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	if v62 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v58
	F_errmsg(m, int32(_a_F_pg_decrypt_iv_0), v14+int32(16))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L2
	} else {
		goto L111
	}
L98:
	;
	v282 = int32(_a_F_pg_decrypt_iv_1)
	goto L97
L99:
	;
	goto L100
L100:
	;
	v259 = int32(_a_F_pg_decrypt_iv_2)
	goto L102
L101:
	;
	v282 = v277
	goto L97
L102:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v259)+8))
	if v62 != v262 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v277 = v274
	goto L101
L104:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v259)+20))
	if v265 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	goto L103
L107:
	;
	v282 = int32(_a_F_pg_decrypt_iv_3)
	goto L97
L108:
	;
	goto L109
L109:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v259)+16))
	if v62 != v270 {
		v259 = v259 + int32(16)
		goto L102
	} else {
		goto L110
	}
L110:
	;
	v277 = v265
	goto L101
L111:
	;
	F_errfinish(m, int32(_a_F_pg_decrypt_iv_4), int32(513), int32(_a_F_pg_decrypt_iv_5))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errcode(m, int32(579))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	if v299 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v340
	F_errmsg(m, int32(_a_F_pg_decrypt_iv_6), v14)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L2
	} else {
		goto L129
	}
L116:
	;
	v340 = int32(_a_F_pg_decrypt_iv_1)
	goto L115
L117:
	;
	goto L118
L118:
	;
	v317 = int32(_a_F_pg_decrypt_iv_2)
	goto L120
L119:
	;
	v340 = v335
	goto L115
L120:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v317)+8))
	if v299 != v320 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v317)+12))
	v335 = v332
	goto L119
L122:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v317)+20))
	if v323 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	goto L121
L125:
	;
	v340 = int32(_a_F_pg_decrypt_iv_3)
	goto L115
L126:
	;
	goto L127
L127:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v317)+16))
	if v299 != v328 {
		v317 = v317 + int32(16)
		goto L120
	} else {
		goto L128
	}
L128:
	;
	v335 = v323
	goto L119
L129:
	;
	F_errfinish(m, int32(_a_F_pg_decrypt_iv_4), int32(441), int32(_a_F_pg_decrypt_iv_7))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L2
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_eucjp_increment(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	v3 = int32(0)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v6 - int32(142) {
	case 0:
		if l1 != int32(2) {
			return int32(0)
		} else {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if base.Ui32(int32(223)) <= base.Ui32(v13) {
				v16 = int32(_a_F_pg_eucjp_increment_0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v16)
				return int32(1)
			} else {
				if base.Ui32(v13) <= base.Ui32(int32(160)) {
					v22 = int32(161)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v22)
					return int32(1)
				} else {
					v26 = int32(1)
					v27 = v13 + v26
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v27)
					return v26
				}
			}
		}
	case 1:
		if l1 != int32(3) {
			v92 = v3
			return v92
		} else {
			v34 = l0 + int32(2)
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
			if base.Ui32(v35) < base.Ui32(int32(161)) {
				v95 = v34
				v97 = int32(161)
				*(*uint8)(unsafe.Add(mBase, uint32(v95))) = uint8(v97)
				return int32(1)
			} else {
				if base.Ui32(int32(253)) < base.Ui32(v35) {
					v77 = l0 + int32(1)
					v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
					if base.Ui32(v78) < base.Ui32(int32(161)) {
						v95 = v77
						v97 = int32(161)
						*(*uint8)(unsafe.Add(mBase, uint32(v95))) = uint8(v97)
						return int32(1)
					} else {
						if base.Ui32(int32(254)) <= base.Ui32(v78) {
							v92 = v3
							return v92
						} else {
							v83 = v77
							v84 = v78
							v85 = int32(1)
							v86 = v84 + v85
							*(*uint8)(unsafe.Add(mBase, uint32(v83))) = uint8(v86)
							return v85
						}
					}
				} else {
					v83 = v34
					v84 = v35
					v85 = int32(1)
					v86 = v84 + v85
					*(*uint8)(unsafe.Add(mBase, uint32(v83))) = uint8(v86)
					return v85
				}
			}
		}
	default:
		if base.I32_extend8_s(v6) < int32(0) {
			if l1 != int32(2) {
				v92 = v3
				return v92
			} else {
				v46 = l0 + int32(1)
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
				if base.Ui32(int32(161)) <= base.Ui32(v47) {
					if base.Ui32(v47) <= base.Ui32(int32(253)) {
						v69 = v46
						v70 = v47
						v71 = int32(1)
						v72 = v70 + v71
						*(*uint8)(unsafe.Add(mBase, uint32(v69))) = uint8(v72)
						return v71
					} else {
						if base.Ui32(int32(161)) <= base.Ui32(v6) {
							if base.Ui32(int32(254)) <= base.Ui32(v6) {
								v92 = v3
								return v92
							} else {
								v69 = l0
								v70 = v6
								v71 = int32(1)
								v72 = v70 + v71
								*(*uint8)(unsafe.Add(mBase, uint32(v69))) = uint8(v72)
								return v71
							}
						} else {
							v54 = l0
							v55 = int32(161)
							*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v55)
							return int32(1)
						}
					}
				} else {
					v54 = v46
					v55 = int32(161)
					*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v55)
					return int32(1)
				}
			}
		} else {
			if v6 == int32(127) {
				return int32(0)
			} else {
				v63 = int32(1)
				v65 = v6 + v63
				*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v65)
				v92 = v63
				return v92
			}
		}
	}
}
func F_pg_eucjp_verifychar(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v5 - int32(142) {
	case 0:
		v8 = int32(2)
		v9 = int32(-1)
		if l1 < v8 {
			v61 = v9
		} else {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if base.Ui32(int32(193)) <= base.Ui32((v12+int32(32))&int32(255)) {
				v58 = v8
				v61 = v58
			} else {
				v61 = v9
			}
		}
	case 1:
		v19 = int32(-1)
		if l1 < int32(3) {
			v61 = v19
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if base.Ui32(int32(93)) < base.Ui32((v22+int32(95))&int32(255)) {
				v61 = v19
			} else {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
				if base.Ui32(int32(94)) <= base.Ui32((v29+int32(95))&int32(255)) {
					v61 = v19
				} else {
					v58 = int32(3)
					v61 = v58
				}
			}
		}
	default:
		v37 = int32(-1)
		v38 = base.I32_extend8_s(v5)
		if int32(0) <= v38 {
			v58 = int32(1)
			v61 = v58
		} else {
			if l1 < int32(2) {
				v61 = v37
			} else {
				if base.Ui32(int32(93)) < base.Ui32((v38+int32(95))&int32(255)) {
					v61 = v37
				} else {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
					if base.Ui32(int32(93)) < base.Ui32((v50+int32(95))&int32(255)) {
						v61 = v37
					} else {
						v58 = int32(2)
						v61 = v58
					}
				}
			}
		}
	}
	return v61
}
func F_pg_euckr_verifystr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	if l1 <= int32(0) {
		v40 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v40 - l0
L2:
	;
	v8 = l1
	v9 = l0
	goto L3
L3:
	;
	v11 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9))))
	if int32(0) <= v11 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v40 = v34
	goto L1
L5:
	;
	v34 = v33 + v9
	v35 = v8 - v33
	if int32(0) < v35 {
		v8 = v35
		v9 = v34
		goto L3
	} else {
		goto L13
	}
L6:
	;
	if v11 == int32(0) {
		v40 = v9
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	if v8 == int32(1) {
		v40 = v9
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v33 = int32(1)
	goto L5
L10:
	;
	if base.Ui32(int32(93)) < base.Ui32((v11+int32(95))&int32(255)) {
		v40 = v9
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if base.Ui32(int32(93)) < base.Ui32((v25+int32(95))&int32(255)) {
		v40 = v9
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v33 = int32(2)
	goto L5
L13:
	;
	goto L4
}
func F_pg_event_trigger_ddl_commands(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
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
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int64
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int64
	_ = v350
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int64
	_ = v368
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	v12 = m.G0
	v14 = v12 - int32(160)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_pg_event_trigger_ddl_commands[0]))
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L6
	} else {
		goto L116
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L6
	} else {
		goto L113
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L6
	} else {
		goto L109
	}
L6:
	;
	return int32(0)
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_event_trigger_ddl_commands[0]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	if v26 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	m.G0 = v14 + int32(160)
	return int32(0)
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v29 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v40 = int32(0)
	goto L11
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v40<<(uint(int32(2))%32))))
	v50 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(104)))) = uint8(v50)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+96)) = int64(0)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	switch v54 {
	case 0:
		goto L23
	case 1:
		goto L22
	case 2:
		goto L16
	case 3:
		goto L21
	case 4:
		goto L17
	case 5:
		goto L20
	case 6:
		goto L19
	default:
		goto L14
	}
L12:
	;
	goto L8
L13:
	;
	v310 = v40 + int32(1)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v310 < v311 {
		v40 = v310
		goto L11
	} else {
		goto L108
	}
L14:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	F_tuplestore_putvalues(m, v296, v297, v14+int32(112), v14+int32(96))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L6
	} else {
		goto L107
	}
L15:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v14)+140)) = v288
	goto L14
L16:
	;
	v231 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+98)) = uint8(v231)
	v233 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+96)) = uint16(v233)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	if v238 != 0 {
		goto L81
	} else {
		goto L82
	}
L17:
	;
	v178 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+98)) = uint8(v178)
	v180 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+96)) = uint16(v180)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v183 = F_CreateCommandTag(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L6
	} else {
		goto L57
	}
L18:
	;
	v83 = int32(0)
	v86 = F_getObjectIdentityParts(m, v14+int32(80), v83, v83, int32(1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L25
	}
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v76
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v49)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v78
	goto L18
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v72
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v49)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v74
	goto L18
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v68
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v49)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v70
	goto L18
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v64
	goto L18
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	if v55 == int32(0) {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v58
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v49)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v60
	goto L18
L25:
	;
	if v86 == int32(0) {
		goto L13
	} else {
		goto L26
	}
L26:
	;
	v90 = int32(0)
	v94 = F_getObjectTypeDescription(m, v14+int32(80), int32(1))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	v100 = int32(0)
	goto L30
L28:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v150
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v14)+84))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+116)) = v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v155 = F_CreateCommandTag(m, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L6
	} else {
		goto L47
	}
L29:
	;
	if v107 == int32(0) {
		v147 = v90
		goto L28
	} else {
		goto L36
	}
L30:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v100*int32(40))+uint32(_c_F_pg_event_trigger_ddl_commands[1])))
	v107 = base.B2i32(v106 == v96)
	if v107 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	v111 = v100 + int32(1)
	if v111 != int32(37) {
		v100 = v111
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	goto L31
L35:
	;
	goto L34
L36:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	v118 = F_get_object_attnum_namespace(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	if v118 == int32(0) {
		v147 = v90
		goto L28
	} else {
		goto L38
	}
L38:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	v124 = F_table_open(m, v122, int32(1))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	v127 = F_get_object_attnum_oid(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	v130 = F_get_catalog_object_by_oid(m, v124, v127, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	if v130 == int32(0) {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v124)+52))
	v137 = F_heap_getattr_2(m, v130, v118, v134, v14+int32(112))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+112)))
	if v139 == int32(1) {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v142 = F_get_namespace_name_or_temp(m, v137)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	F_sequence_close(m, v124, int32(1))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	v147 = v142
	goto L28
L47:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v155<<(uint(int32(3))%32))+uint32(_c_F_pg_event_trigger_ddl_commands[2])))
	goto L48
L48:
	;
	v162 = F_cstring_to_text(m, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+124)) = v162
	v165 = F_cstring_to_text(m, v94)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v165
	if v147 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v175 = F_cstring_to_text(m, v86)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L6
	} else {
		goto L56
	}
L52:
	;
	v170 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+101)) = uint8(v170)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v172 = F_cstring_to_text(m, v147)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+132)) = v172
	goto L51
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+136)) = v175
	goto L15
L57:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v183<<(uint(int32(3))%32))+uint32(_c_F_pg_event_trigger_ddl_commands[2])))
	goto L58
L58:
	;
	v190 = F_cstring_to_text(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+124)) = v190
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	switch v194 {
	case 0, 1, 2, 3, 4, 5, 7, 8, 10, 11, 13, 14, 15, 18, 20, 23, 24, 25, 26, 27, 28, 30, 31, 32, 33, 35, 38, 39, 40, 43, 44, 45, 46, 47, 48, 50, 51:
		goto L62
	case 6:
		v225 = int32(_a_F_pg_event_trigger_ddl_commands_0)
		goto L60
	case 9:
		goto L74
	case 12:
		goto L73
	case 16:
		goto L72
	case 17:
		goto L71
	case 19:
		goto L70
	case 21:
		goto L69
	case 22:
		goto L68
	case 29:
		goto L66
	case 34:
		goto L65
	case 36:
		goto L67
	case 37:
		goto L75
	case 41:
		goto L76
	case 42:
		goto L64
	case 49:
		goto L63
	default:
		goto L61
	}
L60:
	;
	v226 = F_cstring_to_text(m, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L6
	} else {
		goto L80
	}
L61:
	;
	v225 = int32(_a_F_pg_event_trigger_ddl_commands_1)
	goto L60
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L6
	} else {
		goto L77
	}
L63:
	;
	v225 = int32(_a_F_pg_event_trigger_ddl_commands_2)
	goto L60
L64:
	;
	v225 = int32(_a_F_pg_event_trigger_ddl_commands_3)
	goto L60
L65:
	;
	v225 = int32(_a_F_pg_event_trigger_ddl_commands_4)
	goto L60
L66:
	;
	v225 = int32(_a_F_pg_event_trigger_ddl_commands_5)
	goto L60
L67:
	;
	v225 = int32(_a_F_pg_event_trigger_ddl_commands_6)
	goto L60
L68:
	;
	v225 = int32(_a_F_pg_event_trigger_ddl_commands_7)
	goto L60
L69:
	;
	v225 = int32(_a_F_pg_event_trigger_ddl_commands_8)
	goto L60
L70:
	;
	v225 = int32(_a_F_pg_event_trigger_ddl_commands_9)
	goto L60
L71:
	;
	v225 = int32(_a_F_pg_event_trigger_ddl_commands_10)
	goto L60
L72:
	;
	v225 = int32(_a_F_pg_event_trigger_ddl_commands_11)
	goto L60
L73:
	;
	v225 = int32(_a_F_pg_event_trigger_ddl_commands_12)
	goto L60
L74:
	;
	v225 = int32(_a_F_pg_event_trigger_ddl_commands_13)
	goto L60
L75:
	;
	v225 = int32(_a_F_pg_event_trigger_ddl_commands_14)
	goto L60
L76:
	;
	v225 = int32(_a_F_pg_event_trigger_ddl_commands_15)
	goto L60
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v194
	F_errmsg_internal(m, int32(_a_F_pg_event_trigger_ddl_commands_16), v14-int32(-64))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_pg_event_trigger_ddl_commands_17), int32(2405), int32(_a_F_pg_event_trigger_ddl_commands_18))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	v228 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+101)) = uint16(v228)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v226
	goto L15
L81:
	;
	v239 = int32(_a_F_pg_event_trigger_ddl_commands_19)
	goto L83
L82:
	;
	v239 = int32(_a_F_pg_event_trigger_ddl_commands_20)
	goto L83
L83:
	;
	v240 = F_cstring_to_text(m, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+124)) = v240
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	switch v245 {
	case 0, 1, 2, 3, 4, 5, 7, 8, 10, 11, 13, 14, 15, 18, 20, 23, 24, 25, 26, 28, 30, 31, 32, 33, 35, 38, 39, 40, 43, 44, 45, 46, 47, 48, 50, 51:
		goto L87
	case 6:
		v277 = int32(_a_F_pg_event_trigger_ddl_commands_21)
		goto L85
	case 9:
		goto L100
	case 12:
		goto L99
	case 16:
		goto L98
	case 17:
		goto L97
	case 19:
		goto L96
	case 21:
		goto L95
	case 22:
		goto L94
	case 27:
		goto L92
	case 29:
		goto L91
	case 34:
		goto L90
	case 36:
		goto L93
	case 37:
		goto L101
	case 41:
		goto L102
	case 42:
		goto L89
	case 49:
		goto L88
	default:
		goto L86
	}
L85:
	;
	v278 = F_cstring_to_text(m, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L6
	} else {
		goto L106
	}
L86:
	;
	v277 = int32(_a_F_pg_event_trigger_ddl_commands_1)
	goto L85
L87:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L6
	} else {
		goto L103
	}
L88:
	;
	v277 = int32(_a_F_pg_event_trigger_ddl_commands_22)
	goto L85
L89:
	;
	v277 = int32(_a_F_pg_event_trigger_ddl_commands_23)
	goto L85
L90:
	;
	v277 = int32(_a_F_pg_event_trigger_ddl_commands_24)
	goto L85
L91:
	;
	v277 = int32(_a_F_pg_event_trigger_ddl_commands_25)
	goto L85
L92:
	;
	v277 = int32(_a_F_pg_event_trigger_ddl_commands_26)
	goto L85
L93:
	;
	v277 = int32(_a_F_pg_event_trigger_ddl_commands_27)
	goto L85
L94:
	;
	v277 = int32(_a_F_pg_event_trigger_ddl_commands_28)
	goto L85
L95:
	;
	v277 = int32(_a_F_pg_event_trigger_ddl_commands_29)
	goto L85
L96:
	;
	v277 = int32(_a_F_pg_event_trigger_ddl_commands_30)
	goto L85
L97:
	;
	v277 = int32(_a_F_pg_event_trigger_ddl_commands_31)
	goto L85
L98:
	;
	v277 = int32(_a_F_pg_event_trigger_ddl_commands_32)
	goto L85
L99:
	;
	v277 = int32(_a_F_pg_event_trigger_ddl_commands_33)
	goto L85
L100:
	;
	v277 = int32(_a_F_pg_event_trigger_ddl_commands_34)
	goto L85
L101:
	;
	v277 = int32(_a_F_pg_event_trigger_ddl_commands_35)
	goto L85
L102:
	;
	v277 = int32(_a_F_pg_event_trigger_ddl_commands_36)
	goto L85
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v245
	F_errmsg_internal(m, int32(_a_F_pg_event_trigger_ddl_commands_16), v14+int32(48))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_pg_event_trigger_ddl_commands_17), int32(2321), int32(_a_F_pg_event_trigger_ddl_commands_37))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L6
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
	v280 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+101)) = uint16(v280)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v278
	goto L15
L107:
	;
	goto L13
L108:
	;
	goto L12
L109:
	;
	F_errcode(m, int32(50463299))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(_a_F_pg_event_trigger_ddl_commands_38)
	F_errmsg(m, int32(_a_F_pg_event_trigger_ddl_commands_39), v14)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_pg_event_trigger_ddl_commands_17), int32(2064), int32(_a_F_pg_event_trigger_ddl_commands_40))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v14)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v350
	F_errmsg_internal(m, int32(_a_F_pg_event_trigger_ddl_commands_41), v14+int32(16))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L6
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_pg_event_trigger_ddl_commands_17), int32(2154), int32(_a_F_pg_event_trigger_ddl_commands_40))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L6
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v366
	v368 = *(*int64)(unsafe.Add(mBase, uint32(v14)+84))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+36)) = v368
	F_errmsg_internal(m, int32(_a_F_pg_event_trigger_ddl_commands_42), v14+int32(32))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_pg_event_trigger_ddl_commands_17), int32(2161), int32(_a_F_pg_event_trigger_ddl_commands_40))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L6
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_event_trigger_dropped_objects(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v49 int64
	_ = v49
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
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
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_pg_event_trigger_dropped_objects[0]))
	if v15 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L39
	}
L2:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)))
	if v18 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_pg_event_trigger_dropped_objects[0]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v35 = v12 + int32(56)
	v37 = v12 + int32(48)
	v40 = v29
	goto L9
L7:
	;
	goto L8
L8:
	;
	m.G0 = v12 + int32(80)
	return int32(0)
L9:
	;
	v49 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12+int32(72)))) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v12-int32(-64)))) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v37))) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(24)))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v49
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v40-int32(40))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v67
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v40-int32(36))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v71
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v40-int32(32))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v75
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40-int32(4)))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v79
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40-int32(3)))))
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v83
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40-int32(2)))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v87
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v40-int32(16))))
	v92 = F_cstring_to_text(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v92
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v40-int32(28))))
	if v97 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v40-int32(24))))
	if v105 != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v98 = F_cstring_to_text(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v101 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+23)) = uint8(v101)
	goto L12
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v98
	goto L12
L17:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v40-int32(20))))
	if v113 != 0 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v106 = F_cstring_to_text(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v109 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)) = uint8(v109)
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v106
	goto L17
L22:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v40-int32(12))))
	if v121 != 0 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v114 = F_cstring_to_text(m, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v117 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+25)) = uint8(v117)
	goto L22
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v114
	goto L22
L27:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	F_tuplestore_putvalues(m, v138, v139, v12+int32(32), v12+int32(16))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L37
	}
L28:
	;
	v122 = F_strlist_to_textarray(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v135 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+26)) = uint16(v135)
	goto L27
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v122
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v40-int32(8))))
	if v127 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v128 = F_strlist_to_textarray(m, v127)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v132 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L36
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v128
	goto L27
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v132
	goto L27
L37:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v146 != 0 {
		v40 = v146
		goto L9
	} else {
		goto L38
	}
L38:
	;
	goto L10
L39:
	;
	F_errcode(m, int32(50463299))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_pg_event_trigger_dropped_objects_0)
	F_errmsg(m, int32(_a_F_pg_event_trigger_dropped_objects_1), v12)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_pg_event_trigger_dropped_objects_2), int32(1537), int32(_a_F_pg_event_trigger_dropped_objects_3))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_export_snapshot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pg_export_snapshot[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = F_ExportSnapshot(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_cstring_to_text(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_pg_flush_data(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	v2 = F_fsync(m, l0)
	return
}
func F_pg_fprintf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	v5 = m.G0
	v7 = v5 - int32(1072)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	if l0 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_pg_fprintf[0])) = int32(28)
		v59 = int32(-1)
		m.G0 = v7 + int32(1072)
		return v59
	} else {
		v15 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+1068)) = uint8(v15)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1064)) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1060)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1056)) = v7 + int32(1040)
		v24 = v7 + int32(16)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1052)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1048)) = v24
		F_dopr(m, v7+int32(1048), l1, l2)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1068)))
			if v35 == int32(0) {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1048))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1052))
				if v38 != v39 {
					v45 = v38 - v39
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1060))
					v47 = F_fwrite(m, v39, int32(1), v45, v46)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						if v47 != v45 {
							v59 = int32(-1)
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1064))
							v59 = v50 + v45
						}
						m.G0 = v7 + int32(1072)
						return v59
					}
				} else {
					if v35 != 0 {
						v59 = int32(-1)
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1064))
						v59 = v43
					}
					m.G0 = v7 + int32(1072)
					return v59
				}
			} else {
				if v35 != 0 {
					v59 = int32(-1)
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1064))
					v59 = v43
				}
				m.G0 = v7 + int32(1072)
				return v59
			}
		}
	}
}
func F_pg_freeaddrinfo_all(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	if l0 == int32(1) {
		if l1 == int32(0) {
		} else {
			v8 = l1
			for {
				v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
				v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
				F_emscripten_builtin_free(m, v10)
				mBase = m.M
				F_emscripten_builtin_free(m, v8)
				mBase = m.M
				if v9 != 0 {
					v8 = v9
					continue
				} else {
					break
				}
				break
			}
		}
	} else {
		if l1 == int32(0) {
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			F_emscripten_builtin_free(m, v15)
			mBase = m.M
			F_emscripten_builtin_free(m, l1)
			mBase = m.M
		}
	}
	return
}
func F_pg_gen_salt(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	v5 = m.G0
	v7 = v5 - int32(160)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_text_to_cstring_buffer(m, v10, v7+int32(16), int32(129))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = v7 + int32(16)
	v24 = F_px_gen_salt(m, v20, v20, int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if int32(0) <= v24 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v28 != v10 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	F_pfree(m, v10)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v34 = F_cstring_to_text_with_len(m, v7+int32(16), v24)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	m.G0 = v7 + int32(160)
	return v34
L13:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v24 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v80
	F_errmsg(m, int32(_a_F_pg_gen_salt_0), v7)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L29
	}
L16:
	;
	v80 = int32(_a_F_pg_gen_salt_1)
	goto L15
L17:
	;
	goto L18
L18:
	;
	v57 = int32(_a_F_pg_gen_salt_2)
	goto L20
L19:
	;
	v80 = v75
	goto L15
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	if v24 != v60 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v75 = v72
	goto L19
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	if v63 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	v80 = int32(_a_F_pg_gen_salt_3)
	goto L15
L26:
	;
	goto L27
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	if v24 != v68 {
		v57 = v57 + int32(16)
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v75 = v63
	goto L19
L29:
	;
	F_errfinish(m, int32(_a_F_pg_gen_salt_4), int32(179), int32(_a_F_pg_gen_salt_5))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_acl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v14|v15 == v2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L35
	} else {
		goto L72
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L35
	} else {
		goto L69
	}
L3:
	;
	m.G0 = v12 + int32(32)
	return v226
L4:
	;
	v222 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v222)
	v226 = int32(0)
	goto L3
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v14 == int32(2613) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v23 = int32(2995)
	goto L8
L7:
	;
	v23 = v14
	goto L8
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_acl[0]))
	if v25 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v85)+28)))
	if v92 == int32(0) {
		goto L4
	} else {
		goto L30
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v26 == v23 {
		v85 = v25
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v30 = v2
	goto L15
L13:
	;
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_get_acl[0])) = v81
	v85 = v81
	goto L9
L15:
	;
	v39 = v30 * int32(40)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_pg_get_acl[1])))
	if v23 != v42 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v81 = v39 + int32(_a_F_pg_get_acl_0)
	goto L14
L17:
	;
	if v30 == int32(36) {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	v49 = (v30 | int32(1)) * int32(40)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+uint32(_c_F_pg_get_acl[1])))
	if v23 == v52 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v81 = v49 + int32(_a_F_pg_get_acl_0)
	goto L14
L22:
	;
	goto L23
L23:
	;
	v59 = (v30 | int32(2)) * int32(40)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v59)+uint32(_c_F_pg_get_acl[1])))
	if v23 == v62 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v81 = v59 + int32(_a_F_pg_get_acl_0)
	goto L14
L25:
	;
	goto L26
L26:
	;
	v69 = (v30 | int32(3)) * int32(40)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)+uint32(_c_F_pg_get_acl[1])))
	if v23 == v72 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v81 = v69 + int32(_a_F_pg_get_acl_0)
	goto L14
L28:
	;
	v30 = v30 + int32(4)
	goto L15
L30:
	;
	if v14 != int32(1259) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)))
	if v210 != int32(1) {
		v226 = v202
		goto L3
	} else {
		goto L68
	}
L32:
	;
	v114 = F_table_open(m, v23, int32(1))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L35
	} else {
		goto L39
	}
L33:
	;
	if v19 == int32(0) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v100 = F_SearchSysCacheCopyAttNum(m, v15, base.I32_extend16_s(v19))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	return int32(0)
L36:
	;
	if v100 == int32(0) {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v110 = F_SysCacheGetAttr(m, int32(7), v100, int32(22), v12+int32(31))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v202 = v110
	goto L31
L39:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_acl[0]))
	if v117 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v184 = int32(*(*int16)(unsafe.Add(mBase, uint32(v177)+20)))
	v186 = F_get_catalog_object_by_oid_extended(m, v114, v184, v15, int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L35
	} else {
		goto L61
	}
L41:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v118 == v23 {
		v177 = v117
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v122 = int32(0)
	goto L46
L44:
	;
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_get_acl[0])) = v173
	v177 = v173
	goto L40
L46:
	;
	v131 = v122 * int32(40)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v131)+uint32(_c_F_pg_get_acl[1])))
	if v23 != v134 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v173 = v131 + int32(_a_F_pg_get_acl_0)
	goto L45
L48:
	;
	if v122 == int32(36) {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	v141 = (v122 | int32(1)) * int32(40)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_c_F_pg_get_acl[1])))
	if v23 == v144 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v173 = v141 + int32(_a_F_pg_get_acl_0)
	goto L45
L53:
	;
	goto L54
L54:
	;
	v151 = (v122 | int32(2)) * int32(40)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v151)+uint32(_c_F_pg_get_acl[1])))
	if v23 == v154 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v173 = v151 + int32(_a_F_pg_get_acl_0)
	goto L45
L56:
	;
	goto L57
L57:
	;
	v161 = (v122 | int32(3)) * int32(40)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v161)+uint32(_c_F_pg_get_acl[1])))
	if v23 == v164 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v173 = v161 + int32(_a_F_pg_get_acl_0)
	goto L45
L59:
	;
	v122 = v122 + int32(4)
	goto L46
L61:
	;
	if v186 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	F_sequence_close(m, v114, int32(1))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L35
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v114)+52))
	v196 = F_heap_getattr_2(m, v186, v92, v193, v12+int32(31))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L35
	} else {
		goto L66
	}
L65:
	;
	goto L4
L66:
	;
	F_sequence_close(m, v114, int32(1))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L35
	} else {
		goto L67
	}
L67:
	;
	v202 = v196
	goto L31
L68:
	;
	goto L4
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v23
	F_errmsg_internal(m, int32(_a_F_pg_get_acl_1), v12+int32(16))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L35
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_pg_get_acl_2), int32(2777), int32(_a_F_pg_get_acl_3))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L35
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v23
	F_errmsg_internal(m, int32(_a_F_pg_get_acl_1), v12)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L35
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_pg_get_acl_2), int32(2777), int32(_a_F_pg_get_acl_3))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L35
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_aios(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v48 int32
	_ = v48
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int64
	_ = v81
	var v99 int64
	_ = v99
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int64
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v339 int64
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int64
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v486 int64
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v630 int64
	_ = v630
	var v632 int32
	_ = v632
	var v633 int64
	_ = v633
	v21 = m.G0
	v23 = v21 - int32(1232)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_aios[0]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	if v33 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v35 = v23 + int32(1156)
	v48 = v32
	v66 = int64(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	m.G0 = v23 + int32(1232)
	return int32(0)
L6:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	v70 = int32(7)
	v72 = v68 + base.I32_wrap_i64(v66)<<(uint(v70)%32)
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_aios[0]))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	goto L8
L7:
	;
	goto L5
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(1224)))) = int32(0)
	v81 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(1216)))) = v81
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(1208)))) = v81
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(1200)))) = v81
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(1192)))) = v81
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(1184)))) = v81
	*(*int64)(unsafe.Add(mBase, uint32(v23)+1176)) = v81
	*(*int64)(unsafe.Add(mBase, uint32(v23)+1168)) = v81
	*(*int64)(unsafe.Add(mBase, uint32(v23)+1159)) = v81
	*(*int64)(unsafe.Add(mBase, uint32(v23)+1152)) = v81
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v72)+48))
	goto L10
L9:
	;
	v630 = v66 + int64(1)
	v632 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_aios[0]))
	v633 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v632)+20)))
	if base.Ui64(v630) < base.Ui64(v633) {
		v48 = v632
		v66 = v630
		goto L6
	} else {
		goto L94
	}
L10:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v120 == int32(0) {
		goto L9
	} else {
		goto L12
	}
L11:
	;
	if v145 != 0 {
		goto L24
	} else {
		goto L25
	}
L12:
	;
	goto L14
L13:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_aios[0]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v23)+1100))
	goto L18
L14:
	;
	v126 = F__emscripten_memcpy_bulkmem(m, v23+int32(1024), v72, int32(128))
	mBase = m.M
	goto L16
L16:
	;
	goto L13
L17:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_aios[1]))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v136)+1040))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v140+v141*int32(640))+44))
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v72)+48))
	if v146 != v99 {
		goto L9
	} else {
		goto L21
	}
L18:
	;
	v136 = F__emscripten_memcpy_bulkmem(m, v23, v130+v131<<(uint(int32(3))%32), int32(1024))
	mBase = m.M
	goto L20
L20:
	;
	goto L17
L21:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v148 != v120 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	goto L11
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+1172)) = (v72 - v75) >> (uint(v70) % 32)
	v154 = F_Int64GetDatum(m, v99)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L27
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+1168)) = v145
	goto L23
L25:
	;
	goto L26
L26:
	;
	v151 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v136)+1152)) = uint8(v151)
	goto L23
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+1176)) = v154
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+int32(1024)))))
	if base.Ui32(v160) <= base.Ui32(int32(7)) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v169 = F_cstring_to_text(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L32
	}
L29:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v160<<(uint(int32(2))%32))+uint32(_c_F_pg_get_aios[2])))
	v168 = v167
	goto L31
L30:
	;
	v168 = int32(0)
	goto L31
L31:
	;
	goto L28
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+1180)) = v169
	if v120 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	F_tuplestore_putvalues(m, v601, v602, v136+int32(1168), v136+int32(1152))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L93
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = int64(72340172838076673)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+7)) = int32(16843009)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+int32(1024))+2)))
	if base.Ui32(v181) <= base.Ui32(int32(2)) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v190 = F_cstring_to_text(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L41
	}
L38:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v181<<(uint(int32(2))%32))+uint32(_c_F_pg_get_aios[3])))
	v189 = v188
	goto L40
L39:
	;
	v189 = int32(0)
	goto L40
L40:
	;
	goto L37
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+1184)) = v190
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1026)))
	switch v193 {
	case 0:
		goto L45
	case 1:
		goto L44
	case 2:
		goto L43
	default:
		goto L42
	}
L42:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+int32(1024))+1)))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v512<<(uint(int32(2))%32))+uint32(_c_F_pg_get_aios[4])))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v517)+8))
	goto L80
L43:
	;
	v343 = *(*int64)(unsafe.Add(mBase, uint32(v136)+1120))
	v344 = F_Int64GetDatum(m, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L63
	}
L44:
	;
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v136)+1120))
	v197 = F_Int64GetDatum(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L46
	}
L45:
	;
	v194 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v136)+1157)) = uint16(v194)
	goto L42
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+1188)) = v197
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+1116)))
	if v200 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v202 = v200 & int32(3)
	v203 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v200) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v339 = int64(0)
	goto L49
L49:
	;
	v340 = F_Int64GetDatum(m, v339)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L62
	}
L50:
	;
	v211 = v203
	v215 = v203
	v218 = int32(0)
	goto L53
L51:
	;
	v247 = v203
	v251 = v203
	goto L52
L52:
	;
	if v202 != 0 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v233 = v136 + v211<<(uint(int32(3))%32)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+28))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v233)+20))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	v241 = v234 + (v235 + (v236 + (v237 + v215)))
	v242 = int32(4)
	v243 = v211 + v242
	v245 = v218 + v242
	if v245 != v200&int32(_a_F_pg_get_aios_0) {
		v211 = v243
		v215 = v241
		v218 = v245
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v247 = v243
	v251 = v241
	goto L52
L55:
	;
	goto L54
L56:
	;
	v267 = v247
	v271 = v251
	v273 = v203
	goto L59
L57:
	;
	v301 = v251
	goto L58
L58:
	;
	v339 = base.I64_extend_i32_u(v301)
	goto L49
L59:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v136+v267<<(uint(int32(3))%32))+4))
	v291 = v290 + v271
	v292 = int32(1)
	v295 = v273 + v292
	if v295 != v202 {
		v267 = v267 + v292
		v271 = v291
		v273 = v295
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v301 = v291
	goto L58
L61:
	;
	goto L60
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+1192)) = v340
	goto L42
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+1188)) = v344
	v347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+1116)))
	if v347 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v349 = v347 & int32(3)
	v350 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v347) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v486 = int64(0)
	goto L66
L66:
	;
	v487 = F_Int64GetDatum(m, v486)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L79
	}
L67:
	;
	v358 = v350
	v362 = v350
	v365 = int32(0)
	goto L70
L68:
	;
	v394 = v350
	v398 = v350
	goto L69
L69:
	;
	if v349 != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v380 = v136 + v358<<(uint(int32(3))%32)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+28))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v380)+20))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v380)+12))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	v388 = v381 + (v382 + (v383 + (v384 + v362)))
	v389 = int32(4)
	v390 = v358 + v389
	v392 = v365 + v389
	if v392 != v347&int32(_a_F_pg_get_aios_0) {
		v358 = v390
		v362 = v388
		v365 = v392
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v394 = v390
	v398 = v388
	goto L69
L72:
	;
	goto L71
L73:
	;
	v414 = v394
	v418 = v398
	v420 = v350
	goto L76
L74:
	;
	v448 = v398
	goto L75
L75:
	;
	v486 = base.I64_extend_i32_u(v448)
	goto L66
L76:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v136+v414<<(uint(int32(3))%32))+4))
	v438 = v437 + v418
	v439 = int32(1)
	v442 = v420 + v439
	if v442 != v349 {
		v414 = v414 + v439
		v418 = v438
		v420 = v442
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v448 = v438
	goto L75
L78:
	;
	goto L77
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+1192)) = v487
	goto L42
L80:
	;
	v519 = F_cstring_to_text(m, v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+1196)) = v519
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1037)))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+1200)) = v522
	if base.Ui32((v120-int32(5))&int32(255)) <= base.Ui32(int32(2)) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v136)+1092))
	v538 = int32(base.Ui32(v534)>>(uint(int32(6))%32)) & int32(7)
	if base.Ui32(v538) <= base.Ui32(int32(4)) {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v136)+1044))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+1204)) = v530
	goto L82
L84:
	;
	goto L85
L85:
	;
	v532 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v136)+1161)) = uint8(v532)
	goto L82
L86:
	;
	v548 = F_cstring_to_text(m, v547)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L90
	}
L87:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v538<<(uint(int32(2))%32))+uint32(_c_F_pg_get_aios[5])))
	v547 = v546
	goto L89
L88:
	;
	v547 = int32(0)
	goto L89
L89:
	;
	goto L86
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+1208)) = v548
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136+int32(1024))+1)))
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v555<<(uint(int32(2))%32))+uint32(_c_F_pg_get_aios[4])))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v560)+4))
	v562 = m.T0[v561].(func(*base.Module, int32) int32)(m, v136+int32(1128))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v564 = F_cstring_to_text(m, v562)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+1212)) = v564
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1027)))
	v568 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v136)+1216)) = v567 & v568
	*(*int32)(unsafe.Add(mBase, uint32(v136)+1224)) = int32(base.Ui32(v567)>>(uint(int32(2))%32)) & v568
	*(*int32)(unsafe.Add(mBase, uint32(v136)+1220)) = int32(base.Ui32(v567)>>(uint(v568)%32)) & v568
	goto L33
L93:
	;
	goto L9
L94:
	;
	goto L7
}
func F_pg_get_constraintdef_ext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v7 != 0 {
		v8 = int32(7)
	} else {
		v8 = int32(2)
	}
	v10 = F_pg_get_constraintdef_worker(m, v3, int32(0), v8, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			v16 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v16)
			return int32(0)
		} else {
			v20 = F_cstring_to_text(m, v10)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v10)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					return v20
				}
			}
		}
	}
}
func F_pg_get_expr_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int64
	_ = v226
	var v233 int32
	_ = v233
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = F_text_to_cstring(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = F_stringToNode(m, v13)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_pfree(m, v13)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v17 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v62 = int32(0)
	v64 = F_pull_varnos(m, v62, v17)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L18
	}
L6:
	;
	v23 = v17
	goto L7
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v31 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L5
L9:
	;
	if v31 != int32(67) {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v53 != 0 {
		v23 = v53
		goto L7
	} else {
		goto L17
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errmsg(m, int32(_a_F_pg_get_expr_worker_0), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_pg_get_expr_worker_1), int32(2740), int32(_a_F_pg_get_expr_worker_2))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	goto L8
L18:
	;
	if l1 != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	m.G0 = v11 - int32(-64)
	return v256
L20:
	;
	F_initStringInfo(m, v9+int32(-16))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L61
	}
L21:
	;
	v158 = F_try_relation_open(m, l1, int32(1))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L50
	}
L22:
	;
	v67 = F_bms_make_singleton(m, int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v64 == int32(0) {
		v213 = v62
		v216 = int32(0)
		goto L20
	} else {
		goto L45
	}
L25:
	;
	v69 = int32(0)
	if v64 == v69 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v122 != 0 {
		goto L21
	} else {
		goto L40
	}
L27:
	;
	v122 = int32(1)
	goto L26
L28:
	;
	goto L29
L29:
	;
	if v67 == int32(0) {
		v113 = v69
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v122 = v113
	goto L26
L31:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v79 < v78 {
		v113 = v69
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v81 = int32(1)
	if v78 <= v81 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v84 = v81
	goto L35
L34:
	;
	v84 = v78
	goto L35
L35:
	;
	v85 = int32(8)
	v90 = int32(0)
	goto L36
L36:
	;
	v97 = v90 << (uint(int32(2)) % 32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v64+v85+v97)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+(v67+v85))))
	v104 = v99 & (v101 ^ int32(-1))
	v106 = base.B2i32(v104 == int32(0))
	if v104 != 0 {
		v113 = v106
		goto L30
	} else {
		goto L38
	}
L37:
	;
	v113 = v106
	goto L30
L38:
	;
	v108 = v90 + int32(1)
	if v108 != v84 {
		v90 = v108
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(_a_F_pg_get_expr_worker_3), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_pg_get_expr_worker_1), int32(2752), int32(_a_F_pg_get_expr_worker_2))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(_a_F_pg_get_expr_worker_4), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_pg_get_expr_worker_1), int32(2759), int32(_a_F_pg_get_expr_worker_2))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	if v158 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v256 = int32(0)
	goto L19
L52:
	;
	goto L53
L53:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)+48))
	v165 = F_palloc0(m, int32(80))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v168 = F_palloc0(m, int32(136))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+24)) = int32(1)
	v172 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v168)+21)) = uint8(v172)
	*(*int32)(unsafe.Add(mBase, uint32(v168)+16)) = l1
	v175 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v168)+12)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = int32(101)
	v182 = F_makeAlias(m, v163+int32(4), v175)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+8)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v168)+4)) = v182
	v186 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v168)+124)) = uint16(v186)
	v188 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v168)+20)) = uint8(v188)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v168
	v195 = F_list_make1_impl(m, int32(1), v9+int32(-60))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v197 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = v197
	*(*int64)(unsafe.Add(mBase, uint32(v165)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v195
	F_set_rtable_names(m, v165, v197, v197)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_set_simple_column_names(m, v165)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v165
	v211 = F_list_make1_impl(m, int32(1), v11)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v213 = v211
	v216 = v158
	goto L20
L61:
	;
	v222 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v222)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v222
	v226 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v222
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+43)) = uint8(v222)
	v233 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+41)) = uint16(v233)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v9 + int32(-16)
	F_get_rule_expr(m, v17, v9+int32(-56), v222)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	if v216 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F_relation_close(m, v216, int32(1))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v250 = F_cstring_to_text(m, v246)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	F_pfree(m, v246)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v256 = v250
	goto L19
}
func F_pg_get_function_identity_arguments(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v32 int32
	_ = v32
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_SearchSysCache1(m, int32(47), v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			v16 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v16)
			v32 = int32(0)
			m.G0 = v6 + int32(16)
			return v32
		} else {
			F_initStringInfo(m, v6)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = int32(0)
				v23 = F_print_function_arguments(m, v6, v10, v21, v21)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_ReleaseCatCache(m, v10)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
						v28 = F_cstring_to_text(m, v27)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v27)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								v32 = v28
								m.G0 = v6 + int32(16)
								return v32
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_get_function_result(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_SearchSysCache1(m, int32(47), v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
			v42 = int32(0)
			m.G0 = v7 + int32(16)
			return v42
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)))
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v21)+96)))
			if v23 == int32(112) {
				F_ReleaseCatCache(m, v11)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
					v42 = int32(0)
					m.G0 = v7 + int32(16)
					return v42
				}
			} else {
				F_initStringInfo(m, v7)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_print_function_rettype(m, v7, v11)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_ReleaseCatCache(m, v11)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
							v38 = F_cstring_to_text(m, v37)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v37)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v42 = v38
									m.G0 = v7 + int32(16)
									return v42
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_get_partkeydef_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int64
	_ = v215
	var v222 int32
	_ = v222
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	v5 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(160)
	m.G0 = v19
	v22 = F_SearchSysCache1(m, int32(45), l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L6
	} else {
		goto L100
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L6
	} else {
		goto L97
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L6
	} else {
		goto L94
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L91
	}
L5:
	;
	m.G0 = v19 + int32(160)
	return v346
L6:
	;
	return int32(0)
L7:
	;
	if v22 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l3 != 0 {
		v346 = int32(0)
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+22)))
	v46 = F_SysCacheGetAttrNotNull(m, int32(45), v22, int32(6))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L15
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_partkeydef_worker_0), v19)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_pg_get_partkeydef_worker_1), int32(1958), int32(_a_F_pg_get_partkeydef_worker_2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L15:
	;
	v50 = F_SysCacheGetAttrNotNull(m, int32(45), v22, int32(7))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v54 = F_heap_attisnull(m, v22, int32(8), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	if v54 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v60 = F_SysCacheGetAttrNotNull(m, int32(45), v22, int32(8))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	v73 = v5
	v74 = v5
	goto L20
L20:
	;
	v75 = F_get_rel_name(m, l0)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L26
	}
L21:
	;
	v62 = F_text_to_cstring(m, v60)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	v64 = F_stringToNode(m, v62)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v66 != int32(1) {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	F_pfree(m, v62)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v73 = v71
	v74 = v64
	goto L20
L26:
	;
	if v75 == int32(0) {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v80 = F_palloc0(m, int32(80))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v83 = F_palloc0(m, int32(136))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = int32(1)
	v87 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v83)+21)) = uint8(v87)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = l0
	v90 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(101)
	v95 = F_makeAlias(m, v75, v90)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v95
	v99 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v83)+124)) = uint16(v99)
	v101 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v83)+20)) = uint8(v101)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v83
	v108 = F_list_make1_impl(m, int32(1), v19+int32(76))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v110 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+20)) = v110
	*(*int64)(unsafe.Add(mBase, uint32(v80)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v108
	F_set_rtable_names(m, v80, v110, v110)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	F_set_simple_column_names(m, v80)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v80
	v126 = F_list_make1_impl(m, int32(1), v19+int32(72))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	F_initStringInfo(m, v19+int32(88))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v134 = v42 + v43
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+4)))
	switch v135 - int32(104) {
	case 0:
		goto L40
	default:
		goto L2
	case 4:
		goto L39
	case 10:
		goto L38
	}
L36:
	;
	v149 = int32(*(*int16)(unsafe.Add(mBase, uint32(v134)+6)))
	if int32(0) < v149 {
		goto L46
	} else {
		goto L47
	}
L37:
	;
	F_appendStringInfoString(m, v19+int32(88), v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L44
	}
L38:
	;
	if l2 != 0 {
		goto L36
	} else {
		goto L43
	}
L39:
	;
	if l2 != 0 {
		goto L36
	} else {
		goto L42
	}
L40:
	;
	if l2 != 0 {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v141 = int32(_a_F_pg_get_partkeydef_worker_3)
	goto L37
L42:
	;
	v141 = int32(_a_F_pg_get_partkeydef_worker_4)
	goto L37
L43:
	;
	v141 = int32(_a_F_pg_get_partkeydef_worker_5)
	goto L37
L44:
	;
	F_appendStringInfoString(m, v19+int32(88), int32(_a_F_pg_get_partkeydef_worker_6))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	goto L36
L46:
	;
	v152 = int32(24)
	v165 = int32(0)
	v166 = int32(_a_F_pg_get_partkeydef_worker_7)
	v169 = v73
	goto L49
L47:
	;
	goto L48
L48:
	;
	if l2 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L49:
	;
	v179 = int32(*(*int16)(unsafe.Add(mBase, uint32(v134+int32(36)+v165<<(uint(int32(1))%32)))))
	F_appendStringInfoString(m, v19+int32(88), v166)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L6
	} else {
		goto L51
	}
L50:
	;
	goto L48
L51:
	;
	if v179 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v273 = v165 << (uint(int32(2)) % 32)
	if l2 != 0 {
		goto L75
	} else {
		goto L76
	}
L53:
	;
	v187 = F_get_attname(m, l0, v179, int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L6
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if v169 == int32(0) {
		goto L1
	} else {
		goto L60
	}
L56:
	;
	v189 = F_quote_identifier(m, v187)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	F_appendStringInfoString(m, v19+int32(88), v189)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	F_get_atttypetypmodcoll(m, l0, v179, v19+int32(104), v19+int32(84), v19+int32(144))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v19)+144))
	v270 = v201
	v271 = v169
	goto L52
L60:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	F_initStringInfo(m, v19+int32(144))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	v211 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+136)) = uint8(v211)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+120)) = v211
	v215 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+112)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v19)+140)) = v211
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+139)) = uint8(v211)
	v222 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+137)) = uint16(v222)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+128)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v19)+124)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v19 + int32(144)
	F_get_rule_expr(m, v206, v19+int32(104), v211)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	v236 = v169 + int32(4)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v19)+144))
	if v206 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if base.Ui32(v236) < base.Ui32(v204+v205<<(uint(int32(2))%32)) {
		goto L70
	} else {
		goto L71
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v241
	F_appendStringInfo(m, v19+int32(88), int32(_a_F_pg_get_partkeydef_worker_8), v19-int32(-64))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L6
	} else {
		goto L69
	}
L65:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	switch v244 - int32(15) {
	case 0:
		goto L67
	default:
		goto L64
	case 4, 23, 24, 25, 26, 33:
		goto L66
	}
L66:
	;
	F_appendStringInfoString(m, v19+int32(88), v241)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L6
	} else {
		goto L68
	}
L67:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v206)+16))
	switch v247 {
	case 0, 3:
		goto L66
	default:
		goto L64
	}
L68:
	;
	goto L63
L69:
	;
	goto L63
L70:
	;
	v261 = v236
	goto L72
L71:
	;
	v261 = int32(0)
	goto L72
L72:
	;
	v262 = F_exprType(m, v206)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L6
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v262
	v265 = F_exprCollation(m, v206)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v265
	v270 = v265
	v271 = v261
	goto L52
L75:
	;
	if l2 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v273+(v50+v152))))
	if v275 == int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	if v275 == v270 {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v279 = F_generate_collation_name(m, v275)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v279
	F_appendStringInfo(m, v19+int32(88), int32(_a_F_pg_get_partkeydef_worker_9), v19+int32(48))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	goto L75
L81:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v273+(v46+v152))))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	F_get_opclass_name(m, v293, v294, v19+int32(88))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L6
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v301 = v165 + int32(1)
	v302 = int32(*(*int16)(unsafe.Add(mBase, uint32(v134)+6)))
	if v301 < v302 {
		v165 = v301
		v166 = int32(_a_F_pg_get_partkeydef_worker_10)
		v169 = v271
		goto L49
	} else {
		goto L85
	}
L84:
	;
	goto L83
L85:
	;
	goto L50
L86:
	;
	F_appendStringInfoChar(m, v19+int32(88), int32(41))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L6
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	F_ReleaseCatCache(m, v22)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L6
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	v346 = v329
	goto L5
L91:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v355
	F_errmsg_internal(m, int32(_a_F_pg_get_partkeydef_worker_11), v19+int32(80))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L6
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_pg_get_partkeydef_worker_1), int32(1992), int32(_a_F_pg_get_partkeydef_worker_2))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_partkeydef_worker_12), v19+int32(16))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_pg_get_partkeydef_worker_1), int32(_a_F_pg_get_partkeydef_worker_13), int32(_a_F_pg_get_partkeydef_worker_14))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L6
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	v386 = int32(*(*int8)(unsafe.Add(mBase, uint32(v134)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v386
	F_errmsg_internal(m, int32(_a_F_pg_get_partkeydef_worker_15), v19+int32(32))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_pg_get_partkeydef_worker_1), int32(2020), int32(_a_F_pg_get_partkeydef_worker_2))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_errmsg_internal(m, int32(_a_F_pg_get_partkeydef_worker_16), int32(0))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_pg_get_partkeydef_worker_1), int32(2053), int32(_a_F_pg_get_partkeydef_worker_2))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L6
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_ruledef_worker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int64
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v578 int32
	_ = v578
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v649 int32
	_ = v649
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	v11 = m.G0
	v13 = v11 - int32(224)
	m.G0 = v13
	F_initStringInfo(m, v13+int32(72))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_SPI_connect_ext(m, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_ruledef_worker[0]))
	if v25 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L1
	} else {
		goto L200
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L197
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L194
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L1
	} else {
		goto L191
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+100)) = int32(26)
	v34 = F_SPI_prepare(m, int32(_a_F_pg_get_ruledef_worker_0), int32(1), v13+int32(100))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v42 = v25
	goto L10
L10:
	;
	v43 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+91)) = uint8(v43)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+92)) = l0
	v50 = F_SPI_execute_plan(m, v42, v13+int32(92), v13+int32(91))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	if v34 == int32(0) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	F_SPI_keepplan(m, v34)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_get_ruledef_worker[0])) = v34
	v42 = v34
	goto L10
L14:
	;
	if v50 != int32(5) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v55 = *(*int64)(unsafe.Add(mBase, _c_F_pg_get_ruledef_worker[1]))
	if v55 == int64(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_pg_get_ruledef_worker[2]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = int32(_a_F_pg_get_ruledef_worker_1)
	v64 = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v64 < v67 {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	goto L18
L18:
	;
	v660 = F_SPI_finish(m)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L186
	}
L19:
	;
	v115 = F_SPI_getbinval(m, v62, v60, v112, v13+int32(223))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L34
	}
L20:
	;
	v112 = v74 + int32(1)
	goto L19
L21:
	;
	v74 = v64
	v75 = v67
	goto L24
L22:
	;
	goto L23
L23:
	;
	v100 = F_SystemAttributeByName(m, v63)
	mBase = m.M
	if v100 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	v77 = int32(4)
	v82 = v60 + int32(20) + v75<<(uint(v77)%32) + v74*int32(100)
	v85 = F_namestrcmp(m, v82+v77, v63)
	mBase = m.M
	if v85 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L23
L26:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+91)))
	if v88 != int32(1) {
		goto L20
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v92 = v74 + int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v92 < v93 {
		v74 = v92
		v75 = v93
		goto L24
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	goto L25
L31:
	;
	v112 = int32(-9)
	goto L19
L32:
	;
	goto L33
L33:
	;
	v104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+74)))
	v112 = v104
	goto L19
L34:
	;
	v117 = int32(_a_F_pg_get_ruledef_worker_2)
	v118 = int32(0)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v118 < v121 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v169 = F_SPI_getbinval(m, v62, v60, v166, v13+int32(223))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L50
	}
L36:
	;
	v166 = v128 + int32(1)
	goto L35
L37:
	;
	v128 = v118
	v129 = v121
	goto L40
L38:
	;
	goto L39
L39:
	;
	v154 = F_SystemAttributeByName(m, v117)
	mBase = m.M
	if v154 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L40:
	;
	v131 = int32(4)
	v136 = v60 + int32(20) + v129<<(uint(v131)%32) + v128*int32(100)
	v139 = F_namestrcmp(m, v136+v131, v117)
	mBase = m.M
	if v139 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L39
L42:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+91)))
	if v142 != int32(1) {
		goto L36
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v146 = v128 + int32(1)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v146 < v147 {
		v128 = v146
		v129 = v147
		goto L40
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	goto L41
L47:
	;
	v166 = int32(-9)
	goto L35
L48:
	;
	goto L49
L49:
	;
	v158 = int32(*(*int16)(unsafe.Add(mBase, uint32(v154)+74)))
	v166 = v158
	goto L35
L50:
	;
	v171 = int32(_a_F_pg_get_ruledef_worker_3)
	v172 = int32(0)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v172 < v175 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v223 = F_SPI_getbinval(m, v62, v60, v220, v13+int32(223))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L66
	}
L52:
	;
	v220 = v182 + int32(1)
	goto L51
L53:
	;
	v182 = v172
	v183 = v175
	goto L56
L54:
	;
	goto L55
L55:
	;
	v208 = F_SystemAttributeByName(m, v171)
	mBase = m.M
	if v208 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L56:
	;
	v185 = int32(4)
	v190 = v60 + int32(20) + v183<<(uint(v185)%32) + v182*int32(100)
	v193 = F_namestrcmp(m, v190+v185, v171)
	mBase = m.M
	if v193 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L55
L58:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+91)))
	if v196 != int32(1) {
		goto L52
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v200 = v182 + int32(1)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v200 < v201 {
		v182 = v200
		v183 = v201
		goto L56
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	goto L57
L63:
	;
	v220 = int32(-9)
	goto L51
L64:
	;
	goto L65
L65:
	;
	v212 = int32(*(*int16)(unsafe.Add(mBase, uint32(v208)+74)))
	v220 = v212
	goto L51
L66:
	;
	v225 = int32(_a_F_pg_get_ruledef_worker_4)
	v226 = int32(0)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v226 < v229 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v277 = F_SPI_getbinval(m, v62, v60, v274, v13+int32(223))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L82
	}
L68:
	;
	v274 = v236 + int32(1)
	goto L67
L69:
	;
	v236 = v226
	v237 = v229
	goto L72
L70:
	;
	goto L71
L71:
	;
	v262 = F_SystemAttributeByName(m, v225)
	mBase = m.M
	if v262 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L72:
	;
	v239 = int32(4)
	v244 = v60 + int32(20) + v237<<(uint(v239)%32) + v236*int32(100)
	v247 = F_namestrcmp(m, v244+v239, v225)
	mBase = m.M
	if v247 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L71
L74:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+91)))
	if v250 != int32(1) {
		goto L68
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v254 = v236 + int32(1)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v254 < v255 {
		v236 = v254
		v237 = v255
		goto L72
	} else {
		goto L78
	}
L77:
	;
	goto L76
L78:
	;
	goto L73
L79:
	;
	v274 = int32(-9)
	goto L67
L80:
	;
	goto L81
L81:
	;
	v266 = int32(*(*int16)(unsafe.Add(mBase, uint32(v262)+74)))
	v274 = v266
	goto L67
L82:
	;
	v279 = int32(_a_F_pg_get_ruledef_worker_5)
	v280 = int32(0)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v280 < v283 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v329 = F_SPI_getvalue(m, v62, v60, v328)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L98
	}
L84:
	;
	v328 = v290 + int32(1)
	goto L83
L85:
	;
	v290 = v280
	v291 = v283
	goto L88
L86:
	;
	goto L87
L87:
	;
	v316 = F_SystemAttributeByName(m, v279)
	mBase = m.M
	if v316 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L88:
	;
	v293 = int32(4)
	v298 = v60 + int32(20) + v291<<(uint(v293)%32) + v290*int32(100)
	v301 = F_namestrcmp(m, v298+v293, v279)
	mBase = m.M
	if v301 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L87
L90:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298)+91)))
	if v304 != int32(1) {
		goto L84
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v308 = v290 + int32(1)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v308 < v309 {
		v290 = v308
		v291 = v309
		goto L88
	} else {
		goto L94
	}
L93:
	;
	goto L92
L94:
	;
	goto L89
L95:
	;
	v328 = int32(-9)
	goto L83
L96:
	;
	goto L97
L97:
	;
	v320 = int32(*(*int16)(unsafe.Add(mBase, uint32(v316)+74)))
	v328 = v320
	goto L83
L98:
	;
	v331 = int32(_a_F_pg_get_ruledef_worker_6)
	v332 = int32(0)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v332 < v335 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	v381 = F_SPI_getvalue(m, v62, v60, v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L114
	}
L100:
	;
	v380 = v342 + int32(1)
	goto L99
L101:
	;
	v342 = v332
	v343 = v335
	goto L104
L102:
	;
	goto L103
L103:
	;
	v368 = F_SystemAttributeByName(m, v331)
	mBase = m.M
	if v368 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L104:
	;
	v345 = int32(4)
	v350 = v60 + int32(20) + v343<<(uint(v345)%32) + v342*int32(100)
	v353 = F_namestrcmp(m, v350+v345, v331)
	mBase = m.M
	if v353 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L103
L106:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+91)))
	if v356 != int32(1) {
		goto L100
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v360 = v342 + int32(1)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v360 < v361 {
		v342 = v360
		v343 = v361
		goto L104
	} else {
		goto L110
	}
L109:
	;
	goto L108
L110:
	;
	goto L105
L111:
	;
	v380 = int32(-9)
	goto L99
L112:
	;
	goto L113
L113:
	;
	v372 = int32(*(*int16)(unsafe.Add(mBase, uint32(v368)+74)))
	v380 = v372
	goto L99
L114:
	;
	v383 = F_stringToNode(m, v381)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	if v383 == int32(0) {
		goto L5
	} else {
		goto L116
	}
L116:
	;
	v388 = F_table_open(m, v223, int32(1))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v390 = F_quote_identifier(m, v115)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v390
	F_appendStringInfo(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_7), v13+int32(48))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v405 = l1 & int32(2)
	if v405 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v406 = int32(_a_F_pg_get_ruledef_worker_8)
	goto L122
L121:
	;
	v406 = int32(_a_F_pg_get_ruledef_worker_9)
	goto L122
L122:
	;
	F_appendStringInfoString(m, v13+int32(72), v406)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	switch v169&int32(255) - int32(49) {
	case 0:
		goto L125
	case 1:
		goto L129
	case 2:
		goto L128
	case 3:
		goto L127
	default:
		goto L126
	}
L124:
	;
	if base.Ui32(int32(4)) <= base.Ui32(l1) {
		goto L139
	} else {
		goto L140
	}
L125:
	;
	F_appendStringInfoString(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_10))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L137
	}
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L133
	}
L127:
	;
	F_appendStringInfoString(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_11))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L132
	}
L128:
	;
	F_appendStringInfoString(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_12))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	F_appendStringInfoString(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_13))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v457 = int32(0)
	goto L124
L131:
	;
	v457 = int32(0)
	goto L124
L132:
	;
	v457 = int32(0)
	goto L124
L133:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = base.I32_extend8_s(v169)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v115
	F_errmsg(m, int32(_a_F_pg_get_ruledef_worker_14), v13+int32(16))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_pg_get_ruledef_worker_15), int32(_a_F_pg_get_ruledef_worker_16), int32(_a_F_pg_get_ruledef_worker_17))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v388)+52))
	v457 = v456
	goto L124
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v465
	F_appendStringInfo(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_18), v13+int32(32))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L144
	}
L139:
	;
	v461 = F_generate_relation_name(m, v223, int32(0))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v463 = F_generate_qualified_relation_name(m, v223)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L143
	}
L142:
	;
	v465 = v461
	goto L138
L143:
	;
	v465 = v463
	goto L138
L144:
	;
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329))))
	if v474 != int32(60) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	F_appendStringInfoString(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_19))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L164
	}
L146:
	;
	if v405 != 0 {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+1)))
	if v477 != int32(62) {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+2)))
	if v480 == int32(0) {
		goto L145
	} else {
		goto L149
	}
L149:
	;
	goto L146
L150:
	;
	F_appendStringInfoString(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_20))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	F_appendStringInfoString(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_21))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L154
	}
L153:
	;
	goto L152
L154:
	;
	v493 = F_stringToNode(m, v329)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v383)+12))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v495)))
	v498 = F_getInsertSelectQuery(m, v496, int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v500 = int32(0)
	F_AcquireRewriteLocks(m, v498, v500, v500)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v505 = v13 + int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v13 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v505
	v516 = F_list_make1_impl(m, int32(1), v13+int32(28))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+196)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+188)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v516
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v498)+52))
	if v524 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v524)+4))
	v528 = base.B2i32(v525 != int32(1))
	goto L161
L160:
	;
	v528 = int32(1)
	goto L161
L161:
	;
	v529 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+216)) = v529
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+215)) = uint8(v529)
	v533 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+213)) = uint16(v533)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+204)) = int64(34359738368)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+212)) = uint8(v528)
	F_set_deparse_for_query(m, v13+int32(100), v498, v529)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	F_get_rule_expr(m, v493, v13+int32(180), int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	goto L145
L164:
	;
	if v277 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	F_appendStringInfoString(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_22))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	if int32(2) <= v563 {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	goto L167
L169:
	;
	F_sequence_close(m, v388, int32(1))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L185
	}
L170:
	;
	F_appendStringInfoChar(m, v13+int32(72), int32(40))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v383)+12))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	v626 = int32(0)
	F_get_query_def(m, v623, v13+int32(72), v626, v457, int32(1), l1, v626, v626)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L183
	}
L173:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	if int32(0) < v571 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v578 = int32(0)
	goto L177
L175:
	;
	goto L176
L176:
	;
	F_appendStringInfoString(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_23))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L182
	}
L177:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v383)+12))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v585+v578<<(uint(int32(2))%32))))
	v592 = int32(0)
	F_get_query_def(m, v589, v13+int32(72), v592, v457, int32(1), l1, v592, v592)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L179
	}
L178:
	;
	goto L176
L179:
	;
	F_appendStringInfoString(m, v13+int32(72), int32(_a_F_pg_get_ruledef_worker_24))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v604 = v578 + int32(1)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	if v604 < v605 {
		v578 = v604
		goto L177
	} else {
		goto L181
	}
L181:
	;
	goto L178
L182:
	;
	goto L169
L183:
	;
	F_appendStringInfoChar(m, v13+int32(72), int32(59))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	goto L169
L185:
	;
	goto L18
L186:
	;
	if v660 != int32(2) {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	m.G0 = v13 + int32(224)
	if v664 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v670 = v665
	goto L190
L189:
	;
	v670 = int32(0)
	goto L190
L190:
	;
	return v670
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_pg_get_ruledef_worker_0)
	F_errmsg_internal(m, int32(_a_F_pg_get_ruledef_worker_25), v13)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_pg_get_ruledef_worker_15), int32(629), int32(_a_F_pg_get_ruledef_worker_26))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_ruledef_worker_27), v13-int32(-64))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(_a_F_pg_get_ruledef_worker_15), int32(641), int32(_a_F_pg_get_ruledef_worker_26))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	F_errmsg_internal(m, int32(_a_F_pg_get_ruledef_worker_28), int32(0))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_pg_get_ruledef_worker_15), int32(_a_F_pg_get_ruledef_worker_29), int32(_a_F_pg_get_ruledef_worker_17))
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L200:
	;
	F_errmsg_internal(m, int32(_a_F_pg_get_ruledef_worker_30), int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_pg_get_ruledef_worker_15), int32(663), int32(_a_F_pg_get_ruledef_worker_26))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_serial_sequence(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(160)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = F_textToQualifiedNameList(m, v13)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = F_makeRangeVarFromNameList(m, v20)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v24 = int32(0)
	v28 = F_RangeVarGetRelidExtended(m, v22, v24, v24, v24, v24)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v30 = F_text_to_cstring(m, v18)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v32 = F_get_attnum(m, v28, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v32 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v36 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L39
	}
L12:
	;
	F_ScanKeyInit(m, v10+int32(16), int32(4), int32(3), int32(184), int32(1259))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_ScanKeyInit(m, v10-int32(-64), int32(5), int32(3), int32(184), v28)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_ScanKeyInit(m, v10+int32(112), int32(6), int32(3), int32(65), v32)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v66 = F_systable_beginscan(m, v36, int32(2674), int32(1), int32(0), int32(3), v10+int32(16))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	F_systable_endscan(m, v66)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L30
	}
L17:
	;
	v68 = F_systable_getnext(m, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v68 == int32(0) {
		v102 = v2
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v73 = v68
	goto L20
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+22)))
	v81 = v79 + v80
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if v82 != int32(1259) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v102 = v2
	goto L16
L22:
	;
	v95 = F_systable_getnext(m, v66)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L28
	}
L23:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v85 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+24)))
	switch v86 - int32(97) {
	case 0, 8:
		goto L25
	default:
		goto L22
	}
L25:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v90 = F_get_rel_relkind(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v90 != int32(83) {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v102 = v94
	goto L16
L28:
	;
	if v95 != 0 {
		v73 = v95
		goto L20
	} else {
		goto L29
	}
L29:
	;
	goto L21
L30:
	;
	F_sequence_close(m, v36, int32(1))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v102 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	m.G0 = v10 + int32(160)
	return v118
L33:
	;
	v109 = F_generate_qualified_relation_name(m, v102)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v115 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v115)
	v118 = int32(0)
	goto L32
L36:
	;
	v111 = F_cstring_to_text(m, v109)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_pfree(m, v109)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v118 = v111
	goto L32
L39:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v30
	F_errmsg(m, int32(_a_F_pg_get_serial_sequence_0), v10)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_pg_get_serial_sequence_1), int32(2859), int32(_a_F_pg_get_serial_sequence_2))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_viewdef(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_get_viewdef_worker(m, v3, int32(2), int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v12 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
			return int32(0)
		} else {
			v16 = F_cstring_to_text(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v6)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					return v16
				}
			}
		}
	}
}
func F_pg_get_viewdef_ext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v6 != 0 {
		v7 = int32(7)
	} else {
		v7 = int32(2)
	}
	v9 = F_pg_get_viewdef_worker(m, v3, v7, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			v15 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v15)
			return int32(0)
		} else {
			v19 = F_cstring_to_text(m, v9)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v9)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					return v19
				}
			}
		}
	}
}
func F_pg_get_viewdef_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
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
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_textToQualifiedNameList(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = F_makeRangeVarFromNameList(m, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				v12 = int32(0)
				v16 = F_RangeVarGetRelidExtended(m, v10, v12, v12, v12, v12)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v20 = F_pg_get_viewdef_worker(m, v16, int32(2), int32(0))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						if v20 == int32(0) {
							v24 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
							return int32(0)
						} else {
							v28 = F_cstring_to_text(m, v20)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v20)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return int32(0)
								} else {
									return v28
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_getnameinfo_all(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
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
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
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
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	if l0 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v310
L2:
	;
	if l4 != 0 {
		goto L62
	} else {
		goto L63
	}
L3:
	;
	v77 = int32(_a_F_pg_getnameinfo_all_0)
	if l3 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L4:
	;
	if l2 == int32(0) {
		v194 = v71
		goto L2
	} else {
		goto L29
	}
L5:
	;
	v310 = int32(0)
	goto L1
L6:
	;
	v65 = m.Env.Getnameinfo(m, l0, l1, l2, l3, l4, l5, l6)
	mBase = m.M
	if v65 != 0 {
		v71 = v65
		goto L4
	} else {
		goto L28
	}
L7:
	;
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v15 != int32(1) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if l2|l4 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v310 = int32(-4)
	goto L1
L10:
	;
	goto L11
L11:
	;
	if l2 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(_a_F_pg_getnameinfo_all_1)
	v24 = int32(-10)
	v28 = F_pg_snprintf(m, l2, l3, int32(_a_F_pg_getnameinfo_all_2), v11+int32(32))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if l4 == int32(0) {
		goto L5
	} else {
		goto L19
	}
L15:
	;
	return int32(0)
L16:
	;
	if v28 < int32(0) {
		v76 = v24
		goto L3
	} else {
		goto L17
	}
L17:
	;
	if l3 <= v28 {
		v76 = v24
		goto L3
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if v39 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v60 = int32(0)
	v61 = int32(-10)
	if v59 < v60 {
		v71 = v61
		goto L4
	} else {
		goto L26
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0 + int32(2)
	v56 = F_pg_snprintf(m, l4, l5, int32(_a_F_pg_getnameinfo_all_2), v11)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L15
	} else {
		goto L25
	}
L22:
	;
	v41 = l0 + int32(3)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v42 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v41
	v49 = F_pg_snprintf(m, l4, l5, int32(_a_F_pg_getnameinfo_all_3), v11+int32(16))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L15
	} else {
		goto L24
	}
L24:
	;
	v59 = v49
	goto L20
L25:
	;
	v59 = v56
	goto L20
L26:
	;
	if v59 < l5 {
		v310 = v60
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v71 = v61
	goto L4
L28:
	;
	goto L5
L29:
	;
	v76 = v71
	goto L3
L30:
	;
	v194 = v76
	goto L2
L31:
	;
	v189 = F_strlen(m, v185)
	mBase = m.M
	goto L30
L32:
	;
	v185 = v77
	goto L31
L33:
	;
	goto L34
L34:
	;
	v83 = l3 - int32(1)
	if (l2^v77)&int32(3) != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v182 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v179))) = uint8(v182)
	v185 = v178
	goto L31
L36:
	;
	v163 = v158
	v164 = v159
	v165 = v160
	goto L58
L37:
	;
	if v153 == int32(0) {
		v178 = v151
		v179 = v152
		goto L35
	} else {
		goto L57
	}
L38:
	;
	v151 = v77
	v152 = l2
	v153 = v83
	goto L37
L39:
	;
	goto L40
L40:
	;
	goto L43
L41:
	;
	if v120 == int32(0) {
		v178 = v117
		v179 = v118
		goto L35
	} else {
		goto L50
	}
L42:
	;
	v117 = v77
	v118 = l2
	v119 = v83
	v120 = base.B2i32(v83 != int32(0))
	goto L41
L43:
	;
	if v83 == int32(0) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v96 = v77
	v97 = l2
	v98 = v83
	goto L45
L45:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v100)
	if v100 == int32(0) {
		v158 = v96
		v159 = v97
		v160 = v98
		goto L36
	} else {
		goto L47
	}
L46:
	;
	v117 = v111
	v118 = v105
	v119 = v107
	v120 = v109
	goto L41
L47:
	;
	v104 = int32(1)
	v105 = v97 + v104
	v107 = v98 - v104
	v108 = int32(0)
	v109 = base.B2i32(v107 != v108)
	v111 = v96 + v104
	if v111&int32(3) == v108 {
		v117 = v111
		v118 = v105
		v119 = v107
		v120 = v109
		goto L41
	} else {
		goto L48
	}
L48:
	;
	if v107 != 0 {
		v96 = v111
		v97 = v105
		v98 = v107
		goto L45
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v123 == int32(0) {
		v151 = v117
		v152 = v118
		v153 = v119
		goto L37
	} else {
		goto L51
	}
L51:
	;
	if base.Ui32(v119) < base.Ui32(int32(4)) {
		v151 = v117
		v152 = v118
		v153 = v119
		goto L37
	} else {
		goto L52
	}
L52:
	;
	v129 = v117
	v130 = v118
	v131 = v119
	goto L53
L53:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v137 = int32(-2139062144)
	if (int32(16843008)-v134|v134)&v137 != v137 {
		v158 = v129
		v159 = v130
		v160 = v131
		goto L36
	} else {
		goto L55
	}
L54:
	;
	v151 = v145
	v152 = v143
	v153 = v147
	goto L37
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v134
	v142 = int32(4)
	v143 = v130 + v142
	v145 = v129 + v142
	v147 = v131 - v142
	if base.Ui32(int32(3)) < base.Ui32(v147) {
		v129 = v145
		v130 = v143
		v131 = v147
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v158 = v151
	v159 = v152
	v160 = v153
	goto L36
L58:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v167)
	if v167 == int32(0) {
		v178 = v163
		v179 = v164
		goto L35
	} else {
		goto L60
	}
L59:
	;
	v178 = v174
	v179 = v172
	goto L35
L60:
	;
	v171 = int32(1)
	v172 = v164 + v171
	v174 = v163 + v171
	v176 = v165 - v171
	if v176 != 0 {
		v163 = v174
		v164 = v172
		v165 = v176
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v195 = int32(_a_F_pg_getnameinfo_all_0)
	if l5 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	goto L64
L64:
	;
	v310 = v194
	goto L1
L65:
	;
	goto L64
L66:
	;
	v307 = F_strlen(m, v303)
	mBase = m.M
	goto L65
L67:
	;
	v303 = v195
	goto L66
L68:
	;
	goto L69
L69:
	;
	v201 = l5 - int32(1)
	if (l4^v195)&int32(3) != 0 {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v300 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v297))) = uint8(v300)
	v303 = v296
	goto L66
L71:
	;
	v281 = v276
	v282 = v277
	v283 = v278
	goto L93
L72:
	;
	if v271 == int32(0) {
		v296 = v269
		v297 = v270
		goto L70
	} else {
		goto L92
	}
L73:
	;
	v269 = v195
	v270 = l4
	v271 = v201
	goto L72
L74:
	;
	goto L75
L75:
	;
	goto L78
L76:
	;
	if v238 == int32(0) {
		v296 = v235
		v297 = v236
		goto L70
	} else {
		goto L85
	}
L77:
	;
	v235 = v195
	v236 = l4
	v237 = v201
	v238 = base.B2i32(v201 != int32(0))
	goto L76
L78:
	;
	if v201 == int32(0) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v214 = v195
	v215 = l4
	v216 = v201
	goto L80
L80:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
	*(*uint8)(unsafe.Add(mBase, uint32(v215))) = uint8(v218)
	if v218 == int32(0) {
		v276 = v214
		v277 = v215
		v278 = v216
		goto L71
	} else {
		goto L82
	}
L81:
	;
	v235 = v229
	v236 = v223
	v237 = v225
	v238 = v227
	goto L76
L82:
	;
	v222 = int32(1)
	v223 = v215 + v222
	v225 = v216 - v222
	v226 = int32(0)
	v227 = base.B2i32(v225 != v226)
	v229 = v214 + v222
	if v229&int32(3) == v226 {
		v235 = v229
		v236 = v223
		v237 = v225
		v238 = v227
		goto L76
	} else {
		goto L83
	}
L83:
	;
	if v225 != 0 {
		v214 = v229
		v215 = v223
		v216 = v225
		goto L80
	} else {
		goto L84
	}
L84:
	;
	goto L81
L85:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235))))
	if v241 == int32(0) {
		v269 = v235
		v270 = v236
		v271 = v237
		goto L72
	} else {
		goto L86
	}
L86:
	;
	if base.Ui32(v237) < base.Ui32(int32(4)) {
		v269 = v235
		v270 = v236
		v271 = v237
		goto L72
	} else {
		goto L87
	}
L87:
	;
	v247 = v235
	v248 = v236
	v249 = v237
	goto L88
L88:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	v255 = int32(-2139062144)
	if (int32(16843008)-v252|v252)&v255 != v255 {
		v276 = v247
		v277 = v248
		v278 = v249
		goto L71
	} else {
		goto L90
	}
L89:
	;
	v269 = v263
	v270 = v261
	v271 = v265
	goto L72
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248))) = v252
	v260 = int32(4)
	v261 = v248 + v260
	v263 = v247 + v260
	v265 = v249 - v260
	if base.Ui32(int32(3)) < base.Ui32(v265) {
		v247 = v263
		v248 = v261
		v249 = v265
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v276 = v269
	v277 = v270
	v278 = v271
	goto L71
L93:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	*(*uint8)(unsafe.Add(mBase, uint32(v282))) = uint8(v285)
	if v285 == int32(0) {
		v296 = v281
		v297 = v282
		goto L70
	} else {
		goto L95
	}
L94:
	;
	v296 = v292
	v297 = v290
	goto L70
L95:
	;
	v289 = int32(1)
	v290 = v282 + v289
	v292 = v281 + v289
	v294 = v283 - v289
	if v294 != 0 {
		v281 = v292
		v282 = v290
		v283 = v294
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
}
func F_pg_has_role_id_name(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v21 = F_GetSysCacheOid(m, int32(10), v10, v18, v18, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			if v21 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
						F_errmsg(m, int32(_a_F_pg_has_role_id_name_0), v8)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_has_role_id_name_1), int32(_a_F_pg_has_role_id_name_2), int32(_a_F_pg_has_role_id_name_3))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v42 = F_convert_any_priv_string(m, v13, int32(_a_F_pg_has_role_id_name_4))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					v44 = F_pg_role_aclcheck(m, v21, v11, v42)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return v44 ^ int32(1)
					}
				}
			}
		}
	}
}
func F_pg_has_role_name(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_pg_has_role_name[0]))
		v19 = int32(0)
		v22 = F_GetSysCacheOid(m, int32(10), v10, v19, v19, v19)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if v22 == int32(0) {
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
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
						F_errmsg(m, int32(_a_F_pg_has_role_name_0), v8)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_has_role_name_1), int32(_a_F_pg_has_role_name_2), int32(_a_F_pg_has_role_name_3))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v43 = F_convert_any_priv_string(m, v12, int32(_a_F_pg_has_role_name_4))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v45 = F_pg_role_aclcheck(m, v22, v17, v43)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return v45 ^ int32(1)
					}
				}
			}
		}
	}
}
func F_pg_has_role_name_name(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v21 = F_GetSysCacheOid(m, int32(10), v11, v18, v18, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			if v21 != 0 {
				v24 = int32(0)
				v27 = F_GetSysCacheOid(m, int32(10), v10, v24, v24, v24)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
								F_errmsg(m, int32(_a_F_pg_has_role_name_name_0), v8+int32(16))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_has_role_name_name_1), int32(_a_F_pg_has_role_name_name_2), int32(_a_F_pg_has_role_name_name_3))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v32 = F_convert_any_priv_string(m, v13, int32(_a_F_pg_has_role_name_name_4))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = F_pg_role_aclcheck(m, v27, v21, v32)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(32)
								return v34 ^ int32(1)
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
						F_errmsg(m, int32(_a_F_pg_has_role_name_name_0), v8)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_has_role_name_name_1), int32(_a_F_pg_has_role_name_name_2), int32(_a_F_pg_has_role_name_name_3))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
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
func F_pg_hmac(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = F_pg_detoast_datum_packed(m, v15)
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
	v20 = int32(1)
	v21 = v16 + v20
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v26 = v24 & v20
	if v26 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v27 = v21
	goto L5
L4:
	;
	v27 = v16 + int32(4)
	goto L5
L5:
	;
	if v24 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v57 = F_downcase_truncate_identifier(m, v27, v55, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v30 = int32(4)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v32&int32(254) == int32(2) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v45 = int32(1)
	if v26 != 0 {
		v55 = int32(base.Ui32(v24)>>(uint(v45)%32)) - v45
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v41 = v30
	goto L12
L11:
	;
	v41 = base.B2i32(v32 == int32(18)) << (uint(v30) % 32)
	goto L12
L12:
	;
	if v32 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v44 = v30
	goto L15
L14:
	;
	v44 = v41
	goto L15
L15:
	;
	v55 = v44
	goto L6
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v61 = F_px_find_hmac(m, v57, v13+int32(12))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v61 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_pfree(m, v57)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L71
	}
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v69 = m.T0[v68].(func(*base.Module, int32) int32)(m, v67)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v72 = v69 + int32(4)
	v73 = F_palloc(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v72 << (uint(int32(2)) % 32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v79 = F_pg_detoast_datum_packed(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v82 = F_pg_detoast_datum_packed(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v84 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v115 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v87 = int32(4)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	if v89&int32(254) == int32(2) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v102 = int32(1)
	if v84&v102 != 0 {
		v114 = int32(base.Ui32(v84)>>(uint(v102)%32)) - v102
		goto L27
	} else {
		goto L37
	}
L31:
	;
	v98 = v87
	goto L33
L32:
	;
	v98 = base.B2i32(v89 == int32(18)) << (uint(v87) % 32)
	goto L33
L33:
	;
	if v89 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v101 = v87
	goto L36
L35:
	;
	v101 = v98
	goto L36
L36:
	;
	v114 = v101
	goto L27
L37:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v114 = int32(base.Ui32(v108)>>(uint(int32(2))%32)) - int32(4)
	goto L27
L38:
	;
	v146 = int32(1)
	if v115&v146 != 0 {
		goto L49
	} else {
		goto L50
	}
L39:
	;
	v118 = int32(4)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	if v120&int32(254) == int32(2) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v133 = int32(1)
	if v115&v133 != 0 {
		v145 = int32(base.Ui32(v115)>>(uint(v133)%32)) - v133
		goto L38
	} else {
		goto L48
	}
L42:
	;
	v129 = v118
	goto L44
L43:
	;
	v129 = base.B2i32(v120 == int32(18)) << (uint(v118) % 32)
	goto L44
L44:
	;
	if v120 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v132 = v118
	goto L47
L46:
	;
	v132 = v129
	goto L47
L47:
	;
	v145 = v132
	goto L38
L48:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v145 = int32(base.Ui32(v139)>>(uint(int32(2))%32)) - int32(4)
	goto L38
L49:
	;
	v150 = v146
	goto L51
L50:
	;
	v150 = int32(4)
	goto L51
L51:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
	m.T0[v152].(func(*base.Module, int32, int32, int32))(m, v67, v82+v150, v145)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v155 = int32(1)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v157&v155 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v160 = v155
	goto L55
L54:
	;
	v160 = int32(4)
	goto L55
L55:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	m.T0[v162].(func(*base.Module, int32, int32, int32))(m, v67, v79+v160, v114)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	m.T0[v167].(func(*base.Module, int32, int32))(m, v67, v73+int32(4))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	m.T0[v170].(func(*base.Module, int32))(m, v67)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v173 != v79 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_pfree(m, v79)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v177 != v82 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	F_pfree(m, v82)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v181 != v16 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L65
L67:
	;
	F_pfree(m, v16)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	m.G0 = v13 + int32(16)
	return v73
L70:
	;
	goto L69
L71:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	if v61 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v229
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v57
	F_errmsg(m, int32(_a_F_pg_hmac_0), v13)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L87
	}
L74:
	;
	v229 = int32(_a_F_pg_hmac_1)
	goto L73
L75:
	;
	goto L76
L76:
	;
	v206 = int32(_a_F_pg_hmac_2)
	goto L78
L77:
	;
	v229 = v224
	goto L73
L78:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
	if v61 != v209 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v206)+12))
	v224 = v221
	goto L77
L80:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v206)+20))
	if v212 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	goto L79
L83:
	;
	v229 = int32(_a_F_pg_hmac_3)
	goto L73
L84:
	;
	goto L85
L85:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v206)+16))
	if v61 != v217 {
		v206 = v206 + int32(16)
		goto L78
	} else {
		goto L86
	}
L86:
	;
	v224 = v212
	goto L77
L87:
	;
	F_errfinish(m, int32(_a_F_pg_hmac_4), int32(513), int32(_a_F_pg_hmac_5))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_ident_file_mappings(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v31
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ident_file_mappings[0]))
	v38 = F_open_auth_file(m, v34, int32(21), v31, v31)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ident_file_mappings[0]))
	F_tokenize_auth_file(m, v41, v38, v21+int32(4), int32(12), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ident_file_mappings[1]))
	v54 = F_AllocSetContextCreateInternal(m, v49, int32(_a_F_pg_ident_file_mappings_0), int32(0), int32(1024), int32(_a_F_pg_ident_file_mappings_1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v56 = int32(_a_F_pg_ident_file_mappings_2)
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_pg_ident_file_mappings[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_ident_file_mappings[1])) = v54
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v60 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_free_auth_file(m, v38)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L38
	}
L7:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v63 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v67 = v21 + int32(11)
	v77 = v2
	v79 = v2
	goto L9
L9:
	;
	v90 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91+v77<<(uint(int32(2))%32))))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
	if v96 == v90 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L6
L11:
	;
	v100 = F_parse_ident_line(m, v95, int32(12))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v103 = v96
	v104 = v90
	goto L13
L13:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v107 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21+int32(32)))) = v107
	v109 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(40)))) = v109
	*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v107
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v21)+11)) = v109
	v120 = v79 + int32(1)
	if v103 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
	v103 = v102
	v104 = v100
	goto L13
L15:
	;
	v124 = F_cstring_to_text(m, v106)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L19
	}
L16:
	;
	v121 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)) = uint8(v121)
	goto L15
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v120
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v124
	if v104 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v103 != 0 {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v129 = F_cstring_to_text(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v142 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v67))) = uint16(v142)
	v144 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+2)) = uint8(v144)
	goto L20
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v129
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v134 = F_cstring_to_text(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v134
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v139 = F_cstring_to_text(m, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v139
	goto L20
L27:
	;
	if v103 != 0 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v146 = F_cstring_to_text(m, v103)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v149 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+14)) = uint8(v149)
	goto L27
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v146
	goto L27
L32:
	;
	v151 = v79
	goto L34
L33:
	;
	v151 = v120
	goto L34
L34:
	;
	v156 = F_heap_form_tuple(m, v29, v21+int32(16), v21+int32(8))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_tuplestore_puttuple(m, v30, v156)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v161 = v77 + int32(1)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v161 < v162 {
		v77 = v161
		v79 = v151
		goto L9
	} else {
		goto L37
	}
L37:
	;
	goto L10
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_ident_file_mappings[1])) = v57
	F_MemoryContextDelete(m, v54)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v188 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v188)
	m.G0 = v21 + int32(48)
	return int32(0)
}
func F_pg_index_has_property(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_text_to_cstring(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = int32(0)
			v13 = F_indexam_property(m, l0, v9, v11, v3, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				return v13
			}
		}
	}
}
func F_pg_last_wal_replay_lsn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v4 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int64(0) {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int32(0)
		} else {
			v14 = F_Int64GetDatum(m, v4)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v14
			}
		}
	}
}
func F_pg_mb_radix_conv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	v7 = int32(0)
	switch l1 - int32(1) {
	case 0:
		v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
		if base.Ui32(l5) < base.Ui32(v215) {
			v243 = v7
			return v243
		} else {
			v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
			if base.Ui32(v217) < base.Ui32(l5) {
				v243 = v7
				return v243
			} else {
				v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v219 != 0 {
					v221 = int32(2)
					v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v228 = *(*int32)(unsafe.Add(mBase, uint32(v219+(l5-v215)<<(uint(v221)%32)+v224<<(uint(v221)%32))))
					return v228
				} else {
					v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v232 = int32(1)
					v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v230+(l5-v215)<<(uint(v232)%32)+v235<<(uint(v232)%32)))))
					v243 = v239
					return v243
				}
			}
		}
	case 1:
		v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
		if base.Ui32(l4) < base.Ui32(v168) {
			v243 = v7
			return v243
		} else {
			v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)))
			if base.Ui32(v170) < base.Ui32(l4) {
				v243 = v7
				return v243
			} else {
				v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+22)))
				if base.Ui32(l5) < base.Ui32(v172) {
					v243 = v7
					return v243
				} else {
					v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+23)))
					if base.Ui32(v174) < base.Ui32(l5) {
						v243 = v7
						return v243
					} else {
						v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v177 != 0 {
							v179 = int32(2)
							v189 = *(*int32)(unsafe.Add(mBase, uint32(v177+(l4-v168)<<(uint(v179)%32)+v176<<(uint(v179)%32))))
							v193 = *(*int32)(unsafe.Add(mBase, uint32(v177+(l5-v172)<<(uint(v179)%32)+v189<<(uint(v179)%32))))
							return v193
						} else {
							v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v197 = int32(1)
							v209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195+(l4-v168)<<(uint(v197)%32)+v176&int32(_a_F_pg_mb_radix_conv_0)<<(uint(v197)%32)))))
							v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195+(l5-v172)<<(uint(v197)%32)+v209<<(uint(v197)%32)))))
							return v213
						}
					}
				}
			}
		}
	case 2:
		v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
		if base.Ui32(l3) < base.Ui32(v101) {
			v243 = v7
			return v243
		} else {
			v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
			if base.Ui32(v103) < base.Ui32(l3) {
				v243 = v7
				return v243
			} else {
				v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
				if base.Ui32(l4) < base.Ui32(v105) {
					v243 = v7
					return v243
				} else {
					v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
					if base.Ui32(v107) < base.Ui32(l4) {
						v243 = v7
						return v243
					} else {
						v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
						if base.Ui32(l5) < base.Ui32(v109) {
							v243 = v7
							return v243
						} else {
							v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
							if base.Ui32(v111) < base.Ui32(l5) {
								v243 = v7
								return v243
							} else {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v114 != 0 {
									v116 = int32(2)
									v130 = *(*int32)(unsafe.Add(mBase, uint32(v114+(l3-v101)<<(uint(v116)%32)+v113<<(uint(v116)%32))))
									v134 = *(*int32)(unsafe.Add(mBase, uint32(v114+(l4-v105)<<(uint(v116)%32)+v130<<(uint(v116)%32))))
									v138 = *(*int32)(unsafe.Add(mBase, uint32(v114+(l5-v109)<<(uint(v116)%32)+v134<<(uint(v116)%32))))
									return v138
								} else {
									v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v142 = int32(1)
									v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140+(l3-v101)<<(uint(v142)%32)+v113&int32(_a_F_pg_mb_radix_conv_0)<<(uint(v142)%32)))))
									v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140+(l4-v105)<<(uint(v142)%32)+v158<<(uint(v142)%32)))))
									v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140+(l5-v109)<<(uint(v142)%32)+v162<<(uint(v142)%32)))))
									return v166
								}
							}
						}
					}
				}
			}
		}
	case 3:
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
		if base.Ui32(l2) < base.Ui32(v14) {
			v243 = v7
			return v243
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+41)))
			if base.Ui32(v16) < base.Ui32(l2) {
				v243 = v7
				return v243
			} else {
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+42)))
				if base.Ui32(l3) < base.Ui32(v18) {
					v243 = v7
					return v243
				} else {
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
					if base.Ui32(v20) < base.Ui32(l3) {
						v243 = v7
						return v243
					} else {
						v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)))
						if base.Ui32(l4) < base.Ui32(v22) {
							v243 = v7
							return v243
						} else {
							v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+45)))
							if base.Ui32(v24) < base.Ui32(l4) {
								v243 = v7
								return v243
							} else {
								v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+46)))
								if base.Ui32(l5) < base.Ui32(v26) {
									v243 = v7
									return v243
								} else {
									v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+47)))
									if base.Ui32(v28) < base.Ui32(l5) {
										v243 = v7
										return v243
									} else {
										v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										if v31 != 0 {
											v33 = int32(2)
											v51 = *(*int32)(unsafe.Add(mBase, uint32(v31+(l2-v14)<<(uint(v33)%32)+v30<<(uint(v33)%32))))
											v55 = *(*int32)(unsafe.Add(mBase, uint32(v31+(l3-v18)<<(uint(v33)%32)+v51<<(uint(v33)%32))))
											v59 = *(*int32)(unsafe.Add(mBase, uint32(v31+(l4-v22)<<(uint(v33)%32)+v55<<(uint(v33)%32))))
											v63 = *(*int32)(unsafe.Add(mBase, uint32(v31+(l5-v26)<<(uint(v33)%32)+v59<<(uint(v33)%32))))
											return v63
										} else {
											v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v67 = int32(1)
											v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65+(l2-v14)<<(uint(v67)%32)+v30&int32(_a_F_pg_mb_radix_conv_0)<<(uint(v67)%32)))))
											v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65+(l3-v18)<<(uint(v67)%32)+v87<<(uint(v67)%32)))))
											v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65+(l4-v22)<<(uint(v67)%32)+v91<<(uint(v67)%32)))))
											v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65+(l5-v26)<<(uint(v67)%32)+v95<<(uint(v67)%32)))))
											return v99
										}
									}
								}
							}
						}
					}
				}
			}
		}
	default:
		v243 = v7
		return v243
	}
}
func F_pg_mbstrlen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mbstrlen[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7*int32(28))+uint32(_c_F_pg_mbstrlen[1])))
	if v12 != int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_report_invalid_encoding_db(m, v16, v29, v36)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L20
	}
L2:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v59 = F_strlen(m, l0)
	mBase = m.M
	return v59
L5:
	;
	v16 = l0
	v18 = v2
	goto L8
L6:
	;
	v56 = v2
	goto L7
L7:
	;
	return v56
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mbstrlen[0]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23*int32(28))+uint32(_c_F_pg_mbstrlen[2])))
	v29 = m.T0[v28].(func(*base.Module, int32) int32)(m, v16)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v56 = v51
	goto L7
L10:
	;
	return int32(0)
L11:
	;
	if int32(2) <= v29 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v36 = int32(1)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v51 = v18 + int32(1)
	v52 = v16 + v29
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v53 != 0 {
		v16 = v52
		v18 = v51
		goto L8
	} else {
		goto L19
	}
L15:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v36))))
	if v40 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v44 = v36 + int32(1)
	if v44 != v29 {
		v36 = v44
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	goto L9
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_mcv_list_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(_a_F_pg_mcv_list_in_0)
			F_errmsg(m, int32(_a_F_pg_mcv_list_in_1), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_pg_mcv_list_in_2), int32(1480), int32(_a_F_pg_mcv_list_in_3))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_pg_mule_verifychar(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32((v7+int32(127))&int32(255)) < base.Ui32(int32(13)) {
		v30 = int32(2)
	} else {
		v14 = int32(3)
		v16 = v7 & int32(254)
		if v16 == int32(154) {
			v30 = v14
		} else {
			if base.Ui32((v7+int32(112))&int32(255)) < base.Ui32(int32(10)) {
				v30 = v14
			} else {
				if v16 == int32(156) {
					v29 = int32(4)
				} else {
					v29 = int32(1)
				}
				v30 = v29
			}
		}
	}
	if l1 < v30 {
		v56 = int32(-1)
		return v56
	} else {
		if base.Ui32(v30) < base.Ui32(int32(2)) {
			return v30
		} else {
			v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
			if int32(0) <= v38 {
				v56 = int32(-1)
				return v56
			} else {
				if v30 == int32(2) {
					return v30
				} else {
					v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
					if int32(0) <= v45 {
						v56 = int32(-1)
						return v56
					} else {
						if base.Ui32(v30) < base.Ui32(int32(4)) {
							return v30
						} else {
							v52 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+3)))
							if int32(0) <= v52 {
								v55 = int32(-1)
							} else {
								v55 = v30
							}
							v56 = v55
							return v56
						}
					}
				}
			}
		}
	}
}
func F_pg_ndistinct_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 float64
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = F_statext_ndistinct_deserialize(m, v14)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_initStringInfo(m, v9+int32(-16))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_appendStringInfoChar(m, v9+int32(-16), int32(123))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v34 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	F_appendStringInfoChar(m, v9+int32(-16), int32(125))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L29
	}
L9:
	;
	v42 = v18 + int32(16) + v34<<(uint(int32(4))%32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
	v45 = *(*float64)(unsafe.Add(mBase, uint32(v42)))
	if int32(0) < v34 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	F_appendStringInfoString(m, v9+int32(-16), int32(_a_F_pg_ndistinct_out_0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v44 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	if base.F64_lt(base.F64_abs(v45), float64(2.147483648e+09)) != 0 {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(_a_F_pg_ndistinct_out_1)
	F_appendStringInfo(m, v9+int32(-16), int32(_a_F_pg_ndistinct_out_2), v9+int32(-32))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v66 = int32(1)
	if v44 == v66 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v69 = v66
	goto L19
L19:
	;
	v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43+v69<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_pg_ndistinct_out_0)
	F_appendStringInfo(m, v9+int32(-16), int32(_a_F_pg_ndistinct_out_2), v9+int32(-48))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L15
L21:
	;
	v92 = v69 + int32(1)
	if v92 != v44 {
		v69 = v92
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v107
	F_appendStringInfo(m, v9+int32(-16), int32(_a_F_pg_ndistinct_out_3), v11)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L27
	}
L24:
	;
	v105 = base.I32_trunc_f64_s(v45)
	v107 = v105
	goto L23
L25:
	;
	goto L26
L26:
	;
	v107 = int32(-2147483648)
	goto L23
L27:
	;
	v115 = v34 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if base.Ui32(v115) < base.Ui32(v116) {
		v34 = v115
		goto L9
	} else {
		goto L28
	}
L28:
	;
	goto L10
L29:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	m.G0 = v11 - int32(-64)
	return v131
}
func F_pg_nextoid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_superuser(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 != 0 {
			v21 = F_table_open(m, v15, int32(3))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v24 = F_index_open(m, v13, int32(3))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
					if base.Ui32(v26) < base.Ui32(int32(_a_F_pg_nextoid_0)) {
						v43 = v26
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v24)+192))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
						if v46 != v43 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int32(0)
								} else {
									v123 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
									v124 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
									v125 = int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v124 + v125
									*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v123 + v125
									F_errmsg(m, int32(_a_F_pg_nextoid_1), v11+int32(48))
									mBase = m.M
									v135 = m.ExcPending
									if v135 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(678), int32(_a_F_pg_nextoid_3))
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v48 = F_SearchSysCacheAttName(m, v15, v14)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								if v48 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50360452))
										mBase = m.M
										v147 = m.ExcPending
										if v147 != 0 {
											return int32(0)
										} else {
											v148 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
											*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v148 + int32(4)
											F_errmsg(m, int32(_a_F_pg_nextoid_4), v11)
											mBase = m.M
											v155 = m.ExcPending
											if v155 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(685), int32(_a_F_pg_nextoid_3))
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
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
									v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
									v54 = v52 + v53
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+68))
									if v55 != int32(26) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v14
												F_errmsg(m, int32(_a_F_pg_nextoid_5), v11+int32(32))
												mBase = m.M
												v173 = m.ExcPending
												if v173 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(694), int32(_a_F_pg_nextoid_3))
													mBase = m.M
													v178 = m.ExcPending
													if v178 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)+192))
										v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+10)))
										if v59 != int32(1) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v183 = m.ExcPending
											if v183 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v186 = m.ExcPending
												if v186 != 0 {
													return int32(0)
												} else {
													v187 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v14
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v187 + int32(4)
													F_errmsg(m, int32(_a_F_pg_nextoid_6), v11+int32(16))
													mBase = m.M
													v196 = m.ExcPending
													if v196 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(702), int32(_a_F_pg_nextoid_3))
														mBase = m.M
														v201 = m.ExcPending
														if v201 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+48)))
											v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+74)))
											if v62 != v63&int32(_a_F_pg_nextoid_7) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v183 = m.ExcPending
												if v183 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v186 = m.ExcPending
													if v186 != 0 {
														return int32(0)
													} else {
														v187 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
														*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v14
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v187 + int32(4)
														F_errmsg(m, int32(_a_F_pg_nextoid_6), v11+int32(16))
														mBase = m.M
														v196 = m.ExcPending
														if v196 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(702), int32(_a_F_pg_nextoid_3))
															mBase = m.M
															v201 = m.ExcPending
															if v201 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v67 = F_GetNewOidWithIndex(m, v21, v13, v63)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v48)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return int32(0)
													} else {
														F_sequence_close(m, v21, int32(3))
														mBase = m.M
														v73 = m.ExcPending
														if v73 != 0 {
															return int32(0)
														} else {
															F_relation_close(m, v24, int32(3))
															mBase = m.M
															v76 = m.ExcPending
															if v76 != 0 {
																return int32(0)
															} else {
																m.G0 = v11 + int32(80)
																return v67
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
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
						if v30 == int32(99) {
							v43 = v26
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v24)+192))
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
							if v46 != v43 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int32(0)
									} else {
										v123 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
										v124 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
										v125 = int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v124 + v125
										*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v123 + v125
										F_errmsg(m, int32(_a_F_pg_nextoid_1), v11+int32(48))
										mBase = m.M
										v135 = m.ExcPending
										if v135 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(678), int32(_a_F_pg_nextoid_3))
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v48 = F_SearchSysCacheAttName(m, v15, v14)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									if v48 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(50360452))
											mBase = m.M
											v147 = m.ExcPending
											if v147 != 0 {
												return int32(0)
											} else {
												v148 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
												*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v148 + int32(4)
												F_errmsg(m, int32(_a_F_pg_nextoid_4), v11)
												mBase = m.M
												v155 = m.ExcPending
												if v155 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(685), int32(_a_F_pg_nextoid_3))
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
									} else {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
										v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
										v54 = v52 + v53
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+68))
										if v55 != int32(26) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v14
													F_errmsg(m, int32(_a_F_pg_nextoid_5), v11+int32(32))
													mBase = m.M
													v173 = m.ExcPending
													if v173 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(694), int32(_a_F_pg_nextoid_3))
														mBase = m.M
														v178 = m.ExcPending
														if v178 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)+192))
											v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+10)))
											if v59 != int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v183 = m.ExcPending
												if v183 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v186 = m.ExcPending
													if v186 != 0 {
														return int32(0)
													} else {
														v187 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
														*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v14
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v187 + int32(4)
														F_errmsg(m, int32(_a_F_pg_nextoid_6), v11+int32(16))
														mBase = m.M
														v196 = m.ExcPending
														if v196 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(702), int32(_a_F_pg_nextoid_3))
															mBase = m.M
															v201 = m.ExcPending
															if v201 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+48)))
												v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+74)))
												if v62 != v63&int32(_a_F_pg_nextoid_7) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v186 = m.ExcPending
														if v186 != 0 {
															return int32(0)
														} else {
															v187 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v14
															*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v187 + int32(4)
															F_errmsg(m, int32(_a_F_pg_nextoid_6), v11+int32(16))
															mBase = m.M
															v196 = m.ExcPending
															if v196 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(702), int32(_a_F_pg_nextoid_3))
																mBase = m.M
																v201 = m.ExcPending
																if v201 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v67 = F_GetNewOidWithIndex(m, v21, v13, v63)
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return int32(0)
													} else {
														F_ReleaseCatCache(m, v48)
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return int32(0)
														} else {
															F_sequence_close(m, v21, int32(3))
															mBase = m.M
															v73 = m.ExcPending
															if v73 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v24, int32(3))
																mBase = m.M
																v76 = m.ExcPending
																if v76 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v11 + int32(80)
																	return v67
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
							v35 = *(*int32)(unsafe.Add(mBase, _c_F_pg_nextoid[0]))
							if base.B2i32(v35 != int32(0))&base.B2i32(v30 == v35) == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_pg_nextoid_8), int32(0))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(671), int32(_a_F_pg_nextoid_3))
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
								v43 = v42
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v24)+192))
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
								if v46 != v43 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int32(0)
										} else {
											v123 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
											v124 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
											v125 = int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v124 + v125
											*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v123 + v125
											F_errmsg(m, int32(_a_F_pg_nextoid_1), v11+int32(48))
											mBase = m.M
											v135 = m.ExcPending
											if v135 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(678), int32(_a_F_pg_nextoid_3))
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									v48 = F_SearchSysCacheAttName(m, v15, v14)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										if v48 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v144 = m.ExcPending
											if v144 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50360452))
												mBase = m.M
												v147 = m.ExcPending
												if v147 != 0 {
													return int32(0)
												} else {
													v148 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
													*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v148 + int32(4)
													F_errmsg(m, int32(_a_F_pg_nextoid_4), v11)
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(685), int32(_a_F_pg_nextoid_3))
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
										} else {
											v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
											v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
											v54 = v52 + v53
											v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+68))
											if v55 != int32(26) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v167 = m.ExcPending
													if v167 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v14
														F_errmsg(m, int32(_a_F_pg_nextoid_5), v11+int32(32))
														mBase = m.M
														v173 = m.ExcPending
														if v173 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(694), int32(_a_F_pg_nextoid_3))
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)+192))
												v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+10)))
												if v59 != int32(1) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v183 = m.ExcPending
													if v183 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v186 = m.ExcPending
														if v186 != 0 {
															return int32(0)
														} else {
															v187 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
															*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v14
															*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v187 + int32(4)
															F_errmsg(m, int32(_a_F_pg_nextoid_6), v11+int32(16))
															mBase = m.M
															v196 = m.ExcPending
															if v196 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(702), int32(_a_F_pg_nextoid_3))
																mBase = m.M
																v201 = m.ExcPending
																if v201 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+48)))
													v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+74)))
													if v62 != v63&int32(_a_F_pg_nextoid_7) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v183 = m.ExcPending
														if v183 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(50856066))
															mBase = m.M
															v186 = m.ExcPending
															if v186 != 0 {
																return int32(0)
															} else {
																v187 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
																*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v14
																*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v187 + int32(4)
																F_errmsg(m, int32(_a_F_pg_nextoid_6), v11+int32(16))
																mBase = m.M
																v196 = m.ExcPending
																if v196 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(702), int32(_a_F_pg_nextoid_3))
																	mBase = m.M
																	v201 = m.ExcPending
																	if v201 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													} else {
														v67 = F_GetNewOidWithIndex(m, v21, v13, v63)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return int32(0)
														} else {
															F_ReleaseCatCache(m, v48)
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return int32(0)
															} else {
																F_sequence_close(m, v21, int32(3))
																mBase = m.M
																v73 = m.ExcPending
																if v73 != 0 {
																	return int32(0)
																} else {
																	F_relation_close(m, v24, int32(3))
																	mBase = m.M
																	v76 = m.ExcPending
																	if v76 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v11 + int32(80)
																		return v67
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = int32(_a_F_pg_nextoid_3)
					F_errmsg(m, int32(_a_F_pg_nextoid_9), v11-int32(-64))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_nextoid_2), int32(663), int32(_a_F_pg_nextoid_3))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
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
func F_pg_opclass_is_visible(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_OpclassIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v23 = v2
		} else {
			v23 = v14
		}
		m.G0 = v7 + int32(16)
		return v23
	}
}
func F_pg_partition_root(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
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
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_SearchSysCacheExists(m, int32(57), v5, v2, v2, v2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 != 0 {
			v13 = F_get_rel_relkind(m, v5)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = F_get_rel_relispartition(m, v5)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					if v15 != 0 {
						v28 = F_get_partition_ancestors(m, v5)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							if v28 != 0 {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
								v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
								v37 = *(*int32)(unsafe.Add(mBase, uint32(v30+v31<<(uint(int32(2))%32)-int32(4))))
								F_list_free(m, v28)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return int32(0)
								} else {
									v40 = v37
									return v40
								}
							} else {
								v40 = v5
								return v40
							}
						}
					} else {
						if v13 == int32(73) {
							v28 = F_get_partition_ancestors(m, v5)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								if v28 != 0 {
									v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
									v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
									v37 = *(*int32)(unsafe.Add(mBase, uint32(v30+v31<<(uint(int32(2))%32)-int32(4))))
									F_list_free(m, v28)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										v40 = v37
										return v40
									}
								} else {
									v40 = v5
									return v40
								}
							}
						} else {
							if v13&int32(255) == int32(112) {
								v28 = F_get_partition_ancestors(m, v5)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return int32(0)
								} else {
									if v28 != 0 {
										v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
										v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
										v37 = *(*int32)(unsafe.Add(mBase, uint32(v30+v31<<(uint(int32(2))%32)-int32(4))))
										F_list_free(m, v28)
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											v40 = v37
											return v40
										}
									} else {
										v40 = v5
										return v40
									}
								}
							} else {
								v24 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
								return int32(0)
							}
						}
					}
				}
			}
		} else {
			v24 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
			return int32(0)
		}
	}
}
func F_pg_pwrite_zeros(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v272 int32
	_ = v272
	var v287 int32
	_ = v287
	v12 = m.G0
	v14 = v12 - int32(1024)
	m.G0 = v14
	v17 = l1
	v18 = l2
	v20 = int32(0)
	goto L1
L1:
	;
	v27 = int32(0)
	if v17 == v27 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	m.G0 = v14 + int32(1024)
	return v287
L3:
	;
	goto L2
L4:
	;
	v287 = v20
	goto L3
L5:
	;
	goto L6
L6:
	;
	v31 = v17
	v33 = v27
	goto L7
L7:
	;
	v43 = v14 + v33<<(uint(int32(3))%32)
	v44 = int32(_a_F_pg_pwrite_zeros_0)
	if base.Ui32(v44) <= base.Ui32(v31) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v58 = m.G0
	v60 = v58 - int32(1024)
	m.G0 = v60
	if v52 <= int32(128) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	goto L8
L10:
	;
	v47 = v44
	goto L12
L11:
	;
	v47 = v31
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(_a_F_pg_pwrite_zeros_1)
	v52 = v33 + int32(1)
	v53 = v31 - v47
	if base.Ui32(int32(126)) < base.Ui32(v33) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	if v53 != 0 {
		v31 = v53
		v33 = v52
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L9
L15:
	;
	m.G0 = v60 + int32(1024)
	if int32(0) <= v272 {
		v17 = v53
		v18 = v18 + base.I64_extend_i32_u(v272)
		v20 = v20 + v272
		goto L1
	} else {
		goto L84
	}
L16:
	;
	v67 = v14
	v69 = v52
	v71 = int32(0)
	v74 = v18
	goto L19
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_pwrite_zeros[0])) = int32(28)
	v272 = int32(-1)
	goto L15
L19:
	;
	if v69 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v272 = v85
	goto L15
L21:
	;
	if v81 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v79 = F_pwrite(m, l0, v77, v78, v74)
	mBase = m.M
	v81 = v79
	goto L21
L23:
	;
	goto L24
L24:
	;
	v80 = F_pwritev(m, l0, v67, v69, v74)
	mBase = m.M
	v81 = v80
	goto L21
L25:
	;
	v272 = int32(-1)
	goto L15
L26:
	;
	goto L27
L27:
	;
	v85 = v81 + v71
	v91 = v67
	v93 = v69
	v94 = v81
	goto L28
L28:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if base.Ui32(v99) <= base.Ui32(v94) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v91 != v60 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v105 = v93 - int32(1)
	if v105 != 0 {
		v91 = v91 + int32(8)
		v93 = v105
		v94 = v94 - v99
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	v272 = v85
	goto L15
L34:
	;
	v108 = v93 << (uint(int32(3)) % 32)
	if v60 == v91 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	goto L36
L36:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v253 + v94
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v256 - v94
	if int32(0) < v93 {
		v67 = v60
		v69 = v93
		v71 = v85
		v74 = v74 + base.I64_extend_i32_u(v81)
		goto L19
	} else {
		goto L83
	}
L37:
	;
	goto L36
L38:
	;
	goto L37
L39:
	;
	v112 = v60 + v108
	if base.Ui32(v91-v112) <= base.Ui32(int32(0)-v108<<(uint(int32(1))%32)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v119 = F___memcpy(m, v60, v91, v108)
	mBase = m.M
	goto L37
L41:
	;
	goto L42
L42:
	;
	v122 = (v60 ^ v91) & int32(3)
	if base.Ui32(v60) < base.Ui32(v91) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	if v224 == int32(0) {
		goto L38
	} else {
		goto L79
	}
L44:
	;
	if base.Ui32(v202) <= base.Ui32(int32(3)) {
		v223 = v201
		v224 = v202
		v225 = v203
		goto L43
	} else {
		goto L75
	}
L45:
	;
	if v122 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	if v122 != 0 {
		v184 = v108
		goto L58
	} else {
		goto L59
	}
L48:
	;
	v223 = v91
	v224 = v108
	v225 = v60
	goto L43
L49:
	;
	goto L50
L50:
	;
	if v60&int32(3) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v201 = v91
	v202 = v108
	v203 = v60
	goto L44
L52:
	;
	goto L53
L53:
	;
	v129 = v91
	v130 = v108
	v131 = v60
	goto L54
L54:
	;
	if v130 == int32(0) {
		goto L38
	} else {
		goto L56
	}
L55:
	;
	v201 = v138
	v202 = v140
	v203 = v142
	goto L44
L56:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v135)
	v137 = int32(1)
	v138 = v129 + v137
	v140 = v130 - v137
	v142 = v131 + v137
	if v142&int32(3) != 0 {
		v129 = v138
		v130 = v140
		v131 = v142
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	if v184 == int32(0) {
		goto L38
	} else {
		goto L71
	}
L59:
	;
	if v112&int32(3) != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v149 = v108
	goto L63
L61:
	;
	v164 = v108
	goto L62
L62:
	;
	if base.Ui32(v164) <= base.Ui32(int32(3)) {
		v184 = v164
		goto L58
	} else {
		goto L67
	}
L63:
	;
	if v149 == int32(0) {
		goto L38
	} else {
		goto L65
	}
L64:
	;
	v164 = v155
	goto L62
L65:
	;
	v155 = v149 - int32(1)
	v156 = v60 + v155
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91+v155))))
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v158)
	if v156&int32(3) != 0 {
		v149 = v155
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v171 = v164
	goto L68
L68:
	;
	v175 = v171 - int32(4)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v91+v175)))
	*(*int32)(unsafe.Add(mBase, uint32(v60+v175))) = v178
	if base.Ui32(int32(3)) < base.Ui32(v175) {
		v171 = v175
		goto L68
	} else {
		goto L70
	}
L69:
	;
	v184 = v175
	goto L58
L70:
	;
	goto L69
L71:
	;
	v191 = v184
	goto L72
L72:
	;
	v195 = v191 - int32(1)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91+v195))))
	*(*uint8)(unsafe.Add(mBase, uint32(v60+v195))) = uint8(v198)
	if v195 != 0 {
		v191 = v195
		goto L72
	} else {
		goto L74
	}
L73:
	;
	goto L38
L74:
	;
	goto L73
L75:
	;
	v208 = v201
	v209 = v202
	v210 = v203
	goto L76
L76:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v212
	v214 = int32(4)
	v215 = v208 + v214
	v217 = v210 + v214
	v219 = v209 - v214
	if base.Ui32(int32(3)) < base.Ui32(v219) {
		v208 = v215
		v209 = v219
		v210 = v217
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v223 = v215
	v224 = v219
	v225 = v217
	goto L43
L78:
	;
	goto L77
L79:
	;
	v230 = v223
	v231 = v224
	v232 = v225
	goto L80
L80:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	*(*uint8)(unsafe.Add(mBase, uint32(v232))) = uint8(v234)
	v236 = int32(1)
	v241 = v231 - v236
	if v241 != 0 {
		v230 = v230 + v236
		v231 = v241
		v232 = v232 + v236
		goto L80
	} else {
		goto L82
	}
L81:
	;
	goto L38
L82:
	;
	goto L81
L83:
	;
	goto L20
L84:
	;
	v287 = v272
	goto L3
}
func F_pg_qsort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v484 int32
	_ = v484
	var v498 int32
	_ = v498
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v551 int32
	_ = v551
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v701 int32
	_ = v701
	var v710 int32
	_ = v710
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
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
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v809 int32
	_ = v809
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v958 int32
	_ = v958
	var v981 int32
	_ = v981
	var v989 int32
	_ = v989
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1042 int32
	_ = v1042
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1101 int32
	_ = v1101
	var v1107 int32
	_ = v1107
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1156 int32
	_ = v1156
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1218 int32
	_ = v1218
	var v1236 int32
	_ = v1236
	var v1251 int32
	_ = v1251
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1271 int32
	_ = v1271
	var v1280 int32
	_ = v1280
	var v1283 int32
	_ = v1283
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1336 int32
	_ = v1336
	var v1359 int32
	_ = v1359
	var v1365 int32
	_ = v1365
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1430 int32
	_ = v1430
	v23 = int32(0) - l2
	if base.Ui32(l1) < base.Ui32(int32(7)) {
		v1190 = l0
		v1191 = l1
		v1192 = l2
		v1193 = l3
		v1208 = v23
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v1211 = v1190 + v1192
	v1213 = v1190 + v1191*v1192
	if base.Ui32(v1213) <= base.Ui32(v1211) {
		goto L1
	} else {
		goto L144
	}
L3:
	;
	v30 = l0
	v31 = l1
	v32 = l2
	v33 = l3
	v44 = l2 & int32(3)
	v45 = l2 & int32(-4)
	v48 = v23
	goto L4
L4:
	;
	v51 = v30 + v32
	v53 = v31
	goto L6
L5:
	;
	v1190 = v30
	v1191 = v1187
	v1192 = v32
	v1193 = v33
	v1208 = v48
	goto L2
L6:
	;
	v74 = v30 + v53*v32
	if base.Ui32(v74) <= base.Ui32(v51) {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v80 = v51
	goto L9
L9:
	;
	v98 = m.T0[v33].(func(*base.Module, int32, int32) int32)(m, v80+v48, v80)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v107 = v30 + int32(base.Ui32(v53)>>(uint(int32(1))%32))*v32
	if v53 != int32(7) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	return
L12:
	;
	if v98 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v102 = v32 + v80
	if base.Ui32(v102) < base.Ui32(v74) {
		v80 = v102
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L10
L16:
	;
	goto L1
L17:
	;
	v113 = v30 + (v53-int32(1))*v32
	if base.Ui32(v53) < base.Ui32(int32(41)) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v142 = v107
	goto L19
L19:
	;
	if v32 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	v138 = F_pg_qsort_med3(m, v136, v135, v133, v33)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L11
	} else {
		goto L27
	}
L21:
	;
	v133 = v113
	v135 = v107
	v136 = v30
	goto L20
L22:
	;
	goto L23
L23:
	;
	v118 = int32(base.Ui32(v53)>>(uint(int32(3))%32)) * v32
	v121 = v118 << (uint(int32(1)) % 32)
	v123 = F_pg_qsort_med3(m, v30, v30+v118, v30+v121, v33)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v127 = F_pg_qsort_med3(m, v107-v118, v107, v107+v118, v33)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	v131 = F_pg_qsort_med3(m, v113-v121, v113-v118, v113, v33)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	v133 = v131
	v135 = v127
	v136 = v123
	goto L20
L27:
	;
	v142 = v138
	goto L19
L28:
	;
	v287 = v30 + (v53-int32(1))*v32
	v295 = v51
	v296 = v287
	v298 = v51
	v300 = v287
	goto L40
L29:
	;
	v147 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v32) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v156 = v147
	v162 = v147
	goto L33
L31:
	;
	v212 = v147
	goto L32
L32:
	;
	if v44 == int32(0) {
		goto L28
	} else {
		goto L36
	}
L33:
	;
	v173 = v30 + v156
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	v175 = v156 + v142
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	*(*uint8)(unsafe.Add(mBase, uint32(v173))) = uint8(v176)
	*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v174)
	v180 = v156 | int32(1)
	v181 = v30 + v180
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	v183 = v180 + v142
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	*(*uint8)(unsafe.Add(mBase, uint32(v181))) = uint8(v184)
	*(*uint8)(unsafe.Add(mBase, uint32(v183))) = uint8(v182)
	v188 = v156 | int32(2)
	v189 = v30 + v188
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v191 = v188 + v142
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	*(*uint8)(unsafe.Add(mBase, uint32(v189))) = uint8(v192)
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v190)
	v196 = v156 | int32(3)
	v197 = v30 + v196
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v199 = v196 + v142
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199))))
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v200)
	*(*uint8)(unsafe.Add(mBase, uint32(v199))) = uint8(v198)
	v203 = int32(4)
	v204 = v156 + v203
	v206 = v162 + v203
	if v206 != v45 {
		v156 = v204
		v162 = v206
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v212 = v204
	goto L32
L35:
	;
	goto L34
L36:
	;
	v235 = v212
	v243 = v147
	goto L37
L37:
	;
	v252 = v30 + v235
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	v254 = v235 + v142
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	*(*uint8)(unsafe.Add(mBase, uint32(v252))) = uint8(v255)
	*(*uint8)(unsafe.Add(mBase, uint32(v254))) = uint8(v253)
	v258 = int32(1)
	v261 = v243 + v258
	if v261 != v44 {
		v235 = v235 + v258
		v243 = v261
		goto L37
	} else {
		goto L39
	}
L38:
	;
	goto L28
L39:
	;
	goto L38
L40:
	;
	if base.Ui32(v296) < base.Ui32(v298) {
		v507 = v295
		v510 = v298
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if base.Ui32(v32) < base.Ui32(v881) {
		goto L138
	} else {
		goto L139
	}
L42:
	;
	if base.Ui32(v510) <= base.Ui32(v296) {
		goto L66
	} else {
		goto L67
	}
L43:
	;
	v317 = v295
	v320 = v298
	goto L44
L44:
	;
	v331 = m.T0[v33].(func(*base.Module, int32, int32) int32)(m, v320, v30)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L11
	} else {
		goto L46
	}
L45:
	;
	v507 = v484
	v510 = v498
	goto L42
L46:
	;
	if int32(0) < v331 {
		v507 = v317
		v510 = v320
		goto L42
	} else {
		goto L47
	}
L47:
	;
	if v331 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v32 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v484 = v317
	goto L50
L50:
	;
	v498 = v32 + v320
	if base.Ui32(v498) <= base.Ui32(v296) {
		v317 = v484
		v320 = v498
		goto L44
	} else {
		goto L63
	}
L51:
	;
	v484 = v32 + v317
	goto L50
L52:
	;
	v339 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v32) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v348 = v339
	v350 = v339
	goto L56
L54:
	;
	v404 = v339
	goto L55
L55:
	;
	if v44 == int32(0) {
		goto L51
	} else {
		goto L59
	}
L56:
	;
	v365 = v348 + v317
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365))))
	v367 = v348 + v320
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367))))
	*(*uint8)(unsafe.Add(mBase, uint32(v365))) = uint8(v368)
	*(*uint8)(unsafe.Add(mBase, uint32(v367))) = uint8(v366)
	v372 = v348 | int32(1)
	v373 = v317 + v372
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	v375 = v372 + v320
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375))))
	*(*uint8)(unsafe.Add(mBase, uint32(v373))) = uint8(v376)
	*(*uint8)(unsafe.Add(mBase, uint32(v375))) = uint8(v374)
	v380 = v348 | int32(2)
	v381 = v317 + v380
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381))))
	v383 = v380 + v320
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v383))))
	*(*uint8)(unsafe.Add(mBase, uint32(v381))) = uint8(v384)
	*(*uint8)(unsafe.Add(mBase, uint32(v383))) = uint8(v382)
	v388 = v348 | int32(3)
	v389 = v317 + v388
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389))))
	v391 = v388 + v320
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391))))
	*(*uint8)(unsafe.Add(mBase, uint32(v389))) = uint8(v392)
	*(*uint8)(unsafe.Add(mBase, uint32(v391))) = uint8(v390)
	v395 = int32(4)
	v396 = v348 + v395
	v398 = v350 + v395
	if v398 != v45 {
		v348 = v396
		v350 = v398
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v404 = v396
	goto L55
L58:
	;
	goto L57
L59:
	;
	v427 = v404
	v428 = v339
	goto L60
L60:
	;
	v444 = v427 + v317
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444))))
	v446 = v427 + v320
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446))))
	*(*uint8)(unsafe.Add(mBase, uint32(v444))) = uint8(v447)
	*(*uint8)(unsafe.Add(mBase, uint32(v446))) = uint8(v445)
	v450 = int32(1)
	v453 = v428 + v450
	if v453 != v44 {
		v427 = v427 + v450
		v428 = v453
		goto L60
	} else {
		goto L62
	}
L61:
	;
	goto L51
L62:
	;
	goto L61
L63:
	;
	goto L45
L64:
	;
	goto L41
L65:
	;
	if v32 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L66:
	;
	v530 = v296
	v534 = v300
	goto L69
L67:
	;
	v720 = v296
	v724 = v300
	goto L68
L68:
	;
	v733 = v507 - v30
	v734 = v510 - v507
	if v733 < v734 {
		goto L90
	} else {
		goto L91
	}
L69:
	;
	v543 = m.T0[v33].(func(*base.Module, int32, int32) int32)(m, v530, v30)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L11
	} else {
		goto L71
	}
L70:
	;
	v720 = v710
	v724 = v701
	goto L68
L71:
	;
	if v543 < int32(0) {
		goto L65
	} else {
		goto L72
	}
L72:
	;
	if v543 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if v32 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v701 = v534
	goto L75
L75:
	;
	v710 = v530 + v48
	if base.Ui32(v510) <= base.Ui32(v710) {
		v530 = v710
		v534 = v701
		goto L69
	} else {
		goto L88
	}
L76:
	;
	v701 = v534 + v48
	goto L75
L77:
	;
	v551 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v32) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v560 = v551
	v562 = v551
	goto L81
L79:
	;
	v616 = v551
	goto L80
L80:
	;
	if v44 == int32(0) {
		goto L76
	} else {
		goto L84
	}
L81:
	;
	v577 = v560 + v530
	v578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577))))
	v579 = v560 + v534
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579))))
	*(*uint8)(unsafe.Add(mBase, uint32(v577))) = uint8(v580)
	*(*uint8)(unsafe.Add(mBase, uint32(v579))) = uint8(v578)
	v584 = v560 | int32(1)
	v585 = v530 + v584
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v585))))
	v587 = v584 + v534
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587))))
	*(*uint8)(unsafe.Add(mBase, uint32(v585))) = uint8(v588)
	*(*uint8)(unsafe.Add(mBase, uint32(v587))) = uint8(v586)
	v592 = v560 | int32(2)
	v593 = v530 + v592
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593))))
	v595 = v592 + v534
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595))))
	*(*uint8)(unsafe.Add(mBase, uint32(v593))) = uint8(v596)
	*(*uint8)(unsafe.Add(mBase, uint32(v595))) = uint8(v594)
	v600 = v560 | int32(3)
	v601 = v530 + v600
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601))))
	v603 = v600 + v534
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603))))
	*(*uint8)(unsafe.Add(mBase, uint32(v601))) = uint8(v604)
	*(*uint8)(unsafe.Add(mBase, uint32(v603))) = uint8(v602)
	v607 = int32(4)
	v608 = v560 + v607
	v610 = v562 + v607
	if v610 != v45 {
		v560 = v608
		v562 = v610
		goto L81
	} else {
		goto L83
	}
L82:
	;
	v616 = v608
	goto L80
L83:
	;
	goto L82
L84:
	;
	v639 = v616
	v640 = v551
	goto L85
L85:
	;
	v656 = v639 + v530
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v656))))
	v658 = v639 + v534
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v658))))
	*(*uint8)(unsafe.Add(mBase, uint32(v656))) = uint8(v659)
	*(*uint8)(unsafe.Add(mBase, uint32(v658))) = uint8(v657)
	v662 = int32(1)
	v665 = v640 + v662
	if v665 != v44 {
		v639 = v639 + v662
		v640 = v665
		goto L85
	} else {
		goto L87
	}
L86:
	;
	goto L76
L87:
	;
	goto L86
L88:
	;
	goto L70
L89:
	;
	v881 = v724 - v720
	v883 = v74 - (v32 + v724)
	if base.Ui32(v881) < base.Ui32(v883) {
		goto L105
	} else {
		goto L106
	}
L90:
	;
	v736 = v733
	goto L92
L91:
	;
	v736 = v734
	goto L92
L92:
	;
	if v736 == int32(0) {
		goto L89
	} else {
		goto L93
	}
L93:
	;
	v739 = v510 - v736
	v741 = v736 & int32(3)
	v742 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v736) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v753 = v742
	v755 = int32(0)
	goto L97
L95:
	;
	v809 = v742
	goto L96
L96:
	;
	if v741 == int32(0) {
		goto L89
	} else {
		goto L100
	}
L97:
	;
	v770 = v30 + v753
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770))))
	v772 = v753 + v739
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v772))))
	*(*uint8)(unsafe.Add(mBase, uint32(v770))) = uint8(v773)
	*(*uint8)(unsafe.Add(mBase, uint32(v772))) = uint8(v771)
	v777 = v753 | int32(1)
	v778 = v30 + v777
	v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v778))))
	v780 = v777 + v739
	v781 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v780))))
	*(*uint8)(unsafe.Add(mBase, uint32(v778))) = uint8(v781)
	*(*uint8)(unsafe.Add(mBase, uint32(v780))) = uint8(v779)
	v785 = v753 | int32(2)
	v786 = v30 + v785
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786))))
	v788 = v785 + v739
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v788))))
	*(*uint8)(unsafe.Add(mBase, uint32(v786))) = uint8(v789)
	*(*uint8)(unsafe.Add(mBase, uint32(v788))) = uint8(v787)
	v793 = v753 | int32(3)
	v794 = v30 + v793
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v794))))
	v796 = v793 + v739
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v796))))
	*(*uint8)(unsafe.Add(mBase, uint32(v794))) = uint8(v797)
	*(*uint8)(unsafe.Add(mBase, uint32(v796))) = uint8(v795)
	v800 = int32(4)
	v801 = v753 + v800
	v803 = v755 + v800
	if v803 != v736&int32(-4) {
		v753 = v801
		v755 = v803
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v809 = v801
	goto L96
L99:
	;
	goto L98
L100:
	;
	v832 = v809
	v833 = v742
	goto L101
L101:
	;
	v849 = v30 + v832
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v849))))
	v851 = v832 + v739
	v852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v851))))
	*(*uint8)(unsafe.Add(mBase, uint32(v849))) = uint8(v852)
	*(*uint8)(unsafe.Add(mBase, uint32(v851))) = uint8(v850)
	v855 = int32(1)
	v858 = v833 + v855
	if v858 != v741 {
		v832 = v832 + v855
		v833 = v858
		goto L101
	} else {
		goto L103
	}
L102:
	;
	goto L89
L103:
	;
	goto L102
L104:
	;
	if base.Ui32(v881) < base.Ui32(v734) {
		goto L64
	} else {
		goto L119
	}
L105:
	;
	v885 = v881
	goto L107
L106:
	;
	v885 = v883
	goto L107
L107:
	;
	if v885 == int32(0) {
		goto L104
	} else {
		goto L108
	}
L108:
	;
	v888 = v74 - v885
	v890 = v885 & int32(3)
	v891 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v885) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v902 = v891
	v905 = int32(0)
	goto L112
L110:
	;
	v958 = v891
	goto L111
L111:
	;
	if v890 == int32(0) {
		goto L104
	} else {
		goto L115
	}
L112:
	;
	v919 = v902 + v510
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919))))
	v921 = v902 + v888
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921))))
	*(*uint8)(unsafe.Add(mBase, uint32(v919))) = uint8(v922)
	*(*uint8)(unsafe.Add(mBase, uint32(v921))) = uint8(v920)
	v926 = v902 | int32(1)
	v927 = v510 + v926
	v928 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v927))))
	v929 = v926 + v888
	v930 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v929))))
	*(*uint8)(unsafe.Add(mBase, uint32(v927))) = uint8(v930)
	*(*uint8)(unsafe.Add(mBase, uint32(v929))) = uint8(v928)
	v934 = v902 | int32(2)
	v935 = v510 + v934
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v935))))
	v937 = v934 + v888
	v938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v937))))
	*(*uint8)(unsafe.Add(mBase, uint32(v935))) = uint8(v938)
	*(*uint8)(unsafe.Add(mBase, uint32(v937))) = uint8(v936)
	v942 = v902 | int32(3)
	v943 = v510 + v942
	v944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943))))
	v945 = v942 + v888
	v946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945))))
	*(*uint8)(unsafe.Add(mBase, uint32(v943))) = uint8(v946)
	*(*uint8)(unsafe.Add(mBase, uint32(v945))) = uint8(v944)
	v949 = int32(4)
	v950 = v902 + v949
	v952 = v905 + v949
	if v952 != v885&int32(-4) {
		v902 = v950
		v905 = v952
		goto L112
	} else {
		goto L114
	}
L113:
	;
	v958 = v950
	goto L111
L114:
	;
	goto L113
L115:
	;
	v981 = v958
	v989 = v891
	goto L116
L116:
	;
	v998 = v981 + v510
	v999 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v998))))
	v1000 = v981 + v888
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1000))))
	*(*uint8)(unsafe.Add(mBase, uint32(v998))) = uint8(v1001)
	*(*uint8)(unsafe.Add(mBase, uint32(v1000))) = uint8(v999)
	v1004 = int32(1)
	v1007 = v989 + v1004
	if v1007 != v890 {
		v981 = v981 + v1004
		v989 = v1007
		goto L116
	} else {
		goto L118
	}
L117:
	;
	goto L104
L118:
	;
	goto L117
L119:
	;
	if base.Ui32(v32) < base.Ui32(v734) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v1032 = base.I32_div_u_s(v734, v32)
	F_pg_qsort(m, v30, v1032, v32, v33)
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L11
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	if base.Ui32(v881) <= base.Ui32(v32) {
		goto L1
	} else {
		goto L124
	}
L123:
	;
	goto L122
L124:
	;
	v1036 = v74 - v881
	v1037 = base.I32_div_u_s(v881, v32)
	if base.Ui32(int32(7)) <= base.Ui32(v1037) {
		v30 = v1036
		v31 = v1037
		goto L4
	} else {
		goto L125
	}
L125:
	;
	v1190 = v1036
	v1191 = v1037
	v1192 = v32
	v1193 = v33
	v1208 = v48
	goto L2
L126:
	;
	v295 = v507
	v296 = v530 + v48
	v298 = v32 + v510
	v300 = v534
	goto L40
L127:
	;
	v1042 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v32) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v1051 = v1042
	v1053 = v1042
	goto L131
L129:
	;
	v1107 = v1042
	goto L130
L130:
	;
	if v44 == int32(0) {
		goto L126
	} else {
		goto L134
	}
L131:
	;
	v1068 = v1051 + v510
	v1069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1068))))
	v1070 = v1051 + v530
	v1071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1070))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1068))) = uint8(v1071)
	*(*uint8)(unsafe.Add(mBase, uint32(v1070))) = uint8(v1069)
	v1075 = v1051 | int32(1)
	v1076 = v510 + v1075
	v1077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1076))))
	v1078 = v1075 + v530
	v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1078))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1076))) = uint8(v1079)
	*(*uint8)(unsafe.Add(mBase, uint32(v1078))) = uint8(v1077)
	v1083 = v1051 | int32(2)
	v1084 = v510 + v1083
	v1085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1084))))
	v1086 = v1083 + v530
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1086))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1084))) = uint8(v1087)
	*(*uint8)(unsafe.Add(mBase, uint32(v1086))) = uint8(v1085)
	v1091 = v1051 | int32(3)
	v1092 = v510 + v1091
	v1093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1092))))
	v1094 = v1091 + v530
	v1095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1094))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1092))) = uint8(v1095)
	*(*uint8)(unsafe.Add(mBase, uint32(v1094))) = uint8(v1093)
	v1098 = int32(4)
	v1099 = v1051 + v1098
	v1101 = v1053 + v1098
	if v1101 != v45 {
		v1051 = v1099
		v1053 = v1101
		goto L131
	} else {
		goto L133
	}
L132:
	;
	v1107 = v1099
	goto L130
L133:
	;
	goto L132
L134:
	;
	v1130 = v1107
	v1131 = v1042
	goto L135
L135:
	;
	v1147 = v1130 + v510
	v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1147))))
	v1149 = v1130 + v530
	v1150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1149))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1147))) = uint8(v1150)
	*(*uint8)(unsafe.Add(mBase, uint32(v1149))) = uint8(v1148)
	v1153 = int32(1)
	v1156 = v1131 + v1153
	if v1156 != v44 {
		v1130 = v1130 + v1153
		v1131 = v1156
		goto L135
	} else {
		goto L137
	}
L136:
	;
	goto L126
L137:
	;
	goto L136
L138:
	;
	v1183 = base.I32_div_u_s(v881, v32)
	F_pg_qsort(m, v74-v881, v1183, v32, v33)
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L11
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	if base.Ui32(v734) <= base.Ui32(v32) {
		goto L1
	} else {
		goto L142
	}
L141:
	;
	goto L140
L142:
	;
	v1187 = base.I32_div_u_s(v734, v32)
	if base.Ui32(int32(7)) <= base.Ui32(v1187) {
		v53 = v1187
		goto L6
	} else {
		goto L143
	}
L143:
	;
	goto L7
L144:
	;
	v1218 = v1192 & int32(3)
	v1236 = v1211
	goto L145
L145:
	;
	if base.Ui32(v1236) <= base.Ui32(v1190) {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	goto L1
L147:
	;
	v1430 = v1192 + v1236
	if base.Ui32(v1430) < base.Ui32(v1213) {
		v1236 = v1430
		goto L145
	} else {
		goto L166
	}
L148:
	;
	v1251 = v1236
	goto L149
L149:
	;
	v1264 = v1251 + v1208
	v1265 = m.T0[v1193].(func(*base.Module, int32, int32) int32)(m, v1264, v1251)
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L11
	} else {
		goto L151
	}
L150:
	;
	goto L147
L151:
	;
	if v1265 <= int32(0) {
		goto L147
	} else {
		goto L152
	}
L152:
	;
	if v1192 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	if base.Ui32(v1190) < base.Ui32(v1264) {
		v1251 = v1264
		goto L149
	} else {
		goto L165
	}
L154:
	;
	v1271 = int32(0)
	if base.B2i32(base.Ui32(v1192) < base.Ui32(int32(4))) == v1271 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v1280 = v1271
	v1283 = v1271
	goto L158
L156:
	;
	v1336 = v1271
	goto L157
L157:
	;
	if v1218 == int32(0) {
		goto L153
	} else {
		goto L161
	}
L158:
	;
	v1297 = v1280 + v1251
	v1298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1297))))
	v1299 = v1280 + v1264
	v1300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1299))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1297))) = uint8(v1300)
	*(*uint8)(unsafe.Add(mBase, uint32(v1299))) = uint8(v1298)
	v1304 = v1280 | int32(1)
	v1305 = v1251 + v1304
	v1306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1305))))
	v1307 = v1304 + v1264
	v1308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1307))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1305))) = uint8(v1308)
	*(*uint8)(unsafe.Add(mBase, uint32(v1307))) = uint8(v1306)
	v1312 = v1280 | int32(2)
	v1313 = v1251 + v1312
	v1314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1313))))
	v1315 = v1312 + v1264
	v1316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1315))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1313))) = uint8(v1316)
	*(*uint8)(unsafe.Add(mBase, uint32(v1315))) = uint8(v1314)
	v1320 = v1280 | int32(3)
	v1321 = v1251 + v1320
	v1322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1321))))
	v1323 = v1320 + v1264
	v1324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1323))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1321))) = uint8(v1324)
	*(*uint8)(unsafe.Add(mBase, uint32(v1323))) = uint8(v1322)
	v1327 = int32(4)
	v1328 = v1280 + v1327
	v1330 = v1283 + v1327
	if v1330 != v1192&int32(-4) {
		v1280 = v1328
		v1283 = v1330
		goto L158
	} else {
		goto L160
	}
L159:
	;
	v1336 = v1328
	goto L157
L160:
	;
	goto L159
L161:
	;
	v1359 = v1336
	v1365 = v1271
	goto L162
L162:
	;
	v1376 = v1359 + v1251
	v1377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1376))))
	v1378 = v1359 + v1264
	v1379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1378))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1376))) = uint8(v1379)
	*(*uint8)(unsafe.Add(mBase, uint32(v1378))) = uint8(v1377)
	v1382 = int32(1)
	v1385 = v1365 + v1382
	if v1385 != v1218 {
		v1359 = v1359 + v1382
		v1365 = v1385
		goto L162
	} else {
		goto L164
	}
L163:
	;
	goto L153
L164:
	;
	goto L163
L165:
	;
	goto L150
L166:
	;
	goto L146
}
func F_pg_read_file_common(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	if int64(0) <= l2 {
		v7 = F_convert_and_check_filename(m, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = F_read_binary_file(m, v7, l1, l2, l3)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				if v11 != 0 {
					v13 = int32(4)
					v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					F_pg_verifymbstr(m, v11+v13, int32(base.Ui32(v15)>>(uint(int32(2))%32))-v13)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						return v11
					}
				} else {
					return v11
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_read_file_common_0), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_read_file_common_1), int32(248), int32(_a_F_pg_read_file_common_2))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
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
func F_pg_regexec(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int64
	_ = v112
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
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
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v670 int32
	_ = v670
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v779 int32
	_ = v779
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1025 int32
	_ = v1025
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1065 int32
	_ = v1065
	var v1073 int32
	_ = v1073
	var v1080 int32
	_ = v1080
	var v1088 int32
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1115 int32
	_ = v1115
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1166 int32
	_ = v1166
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1186 int32
	_ = v1186
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1279 int32
	_ = v1279
	v15 = m.G0
	v17 = v15 - int32(_a_F_pg_regexec_0)
	m.G0 = v17
	v19 = int32(16)
	if l0 == int32(0) {
		v1279 = v19
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 + int32(_a_F_pg_regexec_0)
	return v1279
L2:
	;
	if l1 == int32(0) {
		v1279 = v19
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v24 != int32(_a_F_pg_regexec_1) {
		v1279 = v19
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v27 != int32(4) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v1279 = int32(17)
	goto L1
L6:
	;
	goto L7
L7:
	;
	if base.Ui32(l2) < base.Ui32(l3) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v1279 = int32(1)
	goto L1
L9:
	;
	goto L10
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pg_set_regex_collation(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+320)) = l0
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+324)) = v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v41&int32(512) != 0 {
		v1279 = v19
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	if v45&int32(_a_F_pg_regexec_2) != 0 {
		v1279 = int32(1)
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+328)) = int32(0)
	v51 = v45 & int32(1)
	if v51 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v112 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+364)) = v112
	v114 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+372)) = v114
	*(*int64)(unsafe.Add(mBase, uint32(v17)+356)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v17)+340)) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v17)+344)) = l1
	v121 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+352)) = l1 + l2<<(uint(v121)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+348)) = l1 + l3<<(uint(v121)%32)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v107)+68))
	if base.Ui32(int32(41)) <= base.Ui32(v129) {
		goto L41
	} else {
		goto L42
	}
L16:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v103 = v101 + int32(1)
	if base.Ui32(l4) < base.Ui32(v103) {
		goto L35
	} else {
		goto L36
	}
L17:
	;
	v91 = int32(8)
	v99 = F__emscripten_memset_bulkmem(m, l5+v91, base.I32_extend8_s(int32(255)), l4<<(uint(int32(3))%32)-v91)
	mBase = m.M
	goto L34
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if base.Ui32(v52) < base.Ui32(l4) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+336)) = l5
	if base.Ui32(l4) < base.Ui32(int32(2)) {
		goto L16
	} else {
		goto L33
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+336)) = l5
	if l4 != int32(1) {
		goto L17
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v58 = v52 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+332)) = v58
	if base.Ui32(int32(21)) <= base.Ui32(v58) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L16
L25:
	;
	v77 = int32(8)
	v85 = F__emscripten_memset_bulkmem(m, v76+v77, base.I32_extend8_s(int32(255)), v58<<(uint(int32(3))%32)-v77)
	mBase = m.M
	goto L32
L26:
	;
	v65 = F_palloc_extended(m, v58<<(uint(int32(3))%32), int32(2))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L11
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v70 = v17 + int32(160)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+336)) = v70
	if v52 == int32(0) {
		v107 = v39
		v108 = l4
		goto L15
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+336)) = v65
	if v65 != 0 {
		v76 = v65
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v1279 = int32(12)
	goto L1
L31:
	;
	v76 = v70
	goto L25
L32:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v17)+324))
	v107 = v86
	v108 = l4
	goto L15
L33:
	;
	goto L17
L34:
	;
	goto L16
L35:
	;
	v105 = l4
	goto L37
L36:
	;
	v105 = v103
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+332)) = v105
	v107 = v39
	v108 = v105
	goto L15
L38:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v17)+336))
	if v1088 == l5 {
		goto L339
	} else {
		goto L340
	}
L39:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v17)+324))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+428))
	if v247 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L40:
	;
	v143 = int32(3)
	v144 = v129 & v143
	v145 = int32(0)
	if base.Ui32(v143) <= base.Ui32(v129-int32(1)) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	v132 = int32(2)
	v135 = F_palloc_extended(m, v129<<(uint(v132)%32), v132)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L11
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+360)) = v17
	if v129 == int32(0) {
		goto L39
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+360)) = v135
	if v135 != 0 {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	v1080 = int32(12)
	goto L38
L46:
	;
	goto L40
L47:
	;
	v154 = v145
	v157 = int32(0)
	goto L50
L48:
	;
	v191 = v145
	goto L49
L49:
	;
	if v144 == int32(0) {
		goto L39
	} else {
		goto L53
	}
L50:
	;
	v169 = v154 << (uint(int32(2)) % 32)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v17)+360))
	v172 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v169+v170))) = v172
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v17)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v174+v169)+4)) = v172
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v17)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v178+v169)+8)) = v172
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v17)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v182+v169)+12)) = v172
	v186 = int32(4)
	v187 = v154 + v186
	v189 = v157 + v186
	if v189 != v129&int32(-4) {
		v154 = v187
		v157 = v189
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v191 = v187
	goto L49
L52:
	;
	goto L51
L53:
	;
	v207 = v191
	v208 = v145
	goto L54
L54:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v17)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v221+v207<<(uint(int32(2))%32)))) = int32(0)
	v227 = int32(1)
	v230 = v208 + v227
	if v230 != v144 {
		v207 = v207 + v227
		v208 = v230
		goto L54
	} else {
		goto L56
	}
L55:
	;
	goto L39
L56:
	;
	goto L55
L57:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v17)+324))
	v500 = v498 + int32(72)
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v498)+16))
	v503 = v501 + int32(36)
	if v51 != 0 {
		goto L88
	} else {
		goto L89
	}
L58:
	;
	v250 = int32(2)
	v251 = v247 << (uint(v250) % 32)
	v253 = F_palloc_extended(m, v251, v250)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+364)) = v253
	v256 = int32(12)
	if v253 == int32(0) {
		v1080 = v256
		goto L38
	} else {
		goto L60
	}
L60:
	;
	v260 = v247 & int32(3)
	v261 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v247) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v268 = v261
	v270 = int32(0)
	goto L64
L62:
	;
	v305 = v261
	goto L63
L63:
	;
	if v260 != 0 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v283 = v268 << (uint(int32(2)) % 32)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v17)+364))
	v286 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v283+v284))) = v286
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v17)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v288+v283)+4)) = v286
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v17)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v292+v283)+8)) = v286
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v17)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v296+v283)+12)) = v286
	v300 = int32(4)
	v301 = v268 + v300
	v303 = v270 + v300
	if v303 != v247&int32(-4) {
		v268 = v301
		v270 = v303
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v305 = v301
	goto L63
L66:
	;
	goto L65
L67:
	;
	v319 = v305
	v322 = v261
	goto L70
L68:
	;
	goto L69
L69:
	;
	v359 = F_palloc_extended(m, v251, int32(2))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L11
	} else {
		goto L73
	}
L70:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v17)+364))
	*(*int32)(unsafe.Add(mBase, uint32(v333+v319<<(uint(int32(2))%32)))) = int32(0)
	v339 = int32(1)
	v342 = v322 + v339
	if v342 != v260 {
		v319 = v319 + v339
		v322 = v342
		goto L70
	} else {
		goto L72
	}
L71:
	;
	goto L69
L72:
	;
	goto L71
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+368)) = v359
	v363 = F_palloc_extended(m, v251, int32(2))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L11
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+372)) = v363
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v17)+368))
	if v366 == int32(0) {
		v1080 = v256
		goto L38
	} else {
		goto L75
	}
L75:
	;
	if v363 == int32(0) {
		v1080 = v256
		goto L38
	} else {
		goto L76
	}
L76:
	;
	v372 = v247 & int32(3)
	v373 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v247) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v380 = v373
	v383 = int32(0)
	goto L80
L78:
	;
	v439 = v373
	goto L79
L79:
	;
	if v372 == int32(0) {
		goto L57
	} else {
		goto L83
	}
L80:
	;
	v395 = v380 << (uint(int32(2)) % 32)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v17)+368))
	v398 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v395+v396))) = v398
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v17)+372))
	*(*int32)(unsafe.Add(mBase, uint32(v400+v395))) = v398
	v404 = int32(4)
	v405 = v395 | v404
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v17)+368))
	*(*int32)(unsafe.Add(mBase, uint32(v405+v406))) = v398
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v17)+372))
	*(*int32)(unsafe.Add(mBase, uint32(v410+v405))) = v398
	v415 = v395 | int32(8)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v17)+368))
	*(*int32)(unsafe.Add(mBase, uint32(v415+v416))) = v398
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v17)+372))
	*(*int32)(unsafe.Add(mBase, uint32(v420+v415))) = v398
	v425 = v395 | int32(12)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v17)+368))
	*(*int32)(unsafe.Add(mBase, uint32(v425+v426))) = v398
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v17)+372))
	*(*int32)(unsafe.Add(mBase, uint32(v430+v425))) = v398
	v435 = v380 + v404
	v437 = v383 + v404
	if v437 != v247&int32(-4) {
		v380 = v435
		v383 = v437
		goto L80
	} else {
		goto L82
	}
L81:
	;
	v439 = v435
	goto L79
L82:
	;
	goto L81
L83:
	;
	v455 = v439
	v456 = v373
	goto L84
L84:
	;
	v470 = v455 << (uint(int32(2)) % 32)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v17)+368))
	v473 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v470+v471))) = v473
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v17)+372))
	*(*int32)(unsafe.Add(mBase, uint32(v475+v470))) = v473
	v479 = int32(1)
	v482 = v456 + v479
	if v482 != v372 {
		v455 = v455 + v479
		v456 = v482
		goto L84
	} else {
		goto L86
	}
L85:
	;
	goto L57
L86:
	;
	goto L85
L87:
	;
	if v108 == int32(0) {
		v1080 = v1047
		goto L38
	} else {
		goto L327
	}
L88:
	;
	v504 = m.G0
	v506 = v504 - int32(16)
	m.G0 = v506
	v509 = v17 + int32(320)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	v515 = F_newdfa(m, v509, v510+int32(20), v500, v17+int32(376))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L11
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	v790 = m.G0
	v792 = v790 - int32(16)
	m.G0 = v792
	v795 = v17 + int32(320)
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v795)+4))
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v796)+16))
	v798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v797)+1)))
	v802 = v17 + int32(376)
	v803 = F_newdfa(m, v795, v796+int32(20), v500, v802)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L11
	} else {
		goto L217
	}
L91:
	;
	m.G0 = v506 + int32(16)
	v1047 = v779
	goto L87
L92:
	;
	if v515 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v509)+36))
	v779 = v519
	goto L91
L94:
	;
	goto L95
L95:
	;
	v522 = F_newdfa(m, v509, v503, v500, v17+int32(_a_F_pg_regexec_3))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L11
	} else {
		goto L96
	}
L96:
	;
	if v522 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+69)))
	if v526 != int32(1) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v550)+16))
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+1)))
	*(*int32)(unsafe.Add(mBase, uint32(v506)+12)) = int32(0)
	v556 = v552 & int32(2)
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v509)+32))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v509)+28))
	v559 = v557
	v565 = v558
	goto L121
L100:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+68)))
	if v544 == int32(1) {
		goto L116
	} else {
		goto L117
	}
L101:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v515)+20))
	if v529 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	F_pfree(m, v529)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L11
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v515)+24))
	if v532 != 0 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L104
L106:
	;
	F_pfree(m, v532)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L11
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v515)+32))
	if v535 != 0 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	goto L108
L110:
	;
	F_pfree(m, v535)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L11
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v515)+36))
	if v538 == int32(0) {
		goto L100
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	F_pfree(m, v538)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L11
	} else {
		goto L115
	}
L115:
	;
	goto L100
L116:
	;
	F_pfree(m, v515)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L11
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v509)+36))
	v779 = v549
	goto L91
L119:
	;
	goto L118
L120:
	;
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522)+69)))
	if v707 != int32(1) {
		goto L169
	} else {
		goto L170
	}
L121:
	;
	v576 = F_shortest(m, v509, v515, v565, v565, v559, v506+int32(12), int32(0))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L11
	} else {
		goto L123
	}
L122:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v506)+12))
	v694 = int32(1)
	v706 = v691
	goto L120
L123:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v509)+36))
	if v578 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v506)+12))
	v694 = v578
	v706 = v579
	goto L120
L125:
	;
	goto L126
L126:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v506)+12))
	if v576 == int32(0) {
		v694 = int32(1)
		v706 = v581
		goto L120
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v506)+12)) = int32(0)
	if base.Ui32(v581) <= base.Ui32(v576) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v587 = v581
	goto L131
L129:
	;
	goto L130
L130:
	;
	v687 = v576 + int32(4)
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v509)+32))
	if base.Ui32(v687) < base.Ui32(v688) {
		v559 = v688
		v565 = v687
		goto L121
	} else {
		goto L168
	}
L131:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v509)+32))
	v603 = v587
	v615 = v601
	goto L133
L132:
	;
	goto L130
L133:
	;
	if v556 != 0 {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	v670 = v587 + int32(4)
	if base.Ui32(v670) <= base.Ui32(v576) {
		v587 = v670
		goto L131
	} else {
		goto L167
	}
L135:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v509)+36))
	if v626 != 0 {
		goto L141
	} else {
		goto L142
	}
L136:
	;
	v619 = F_shortest(m, v509, v522, v587, v603, v615, int32(0), v506+int32(8))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L11
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v623 = F_longest(m, v509, v522, v587, v615, v506+int32(8))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L11
	} else {
		goto L140
	}
L139:
	;
	v625 = v619
	goto L135
L140:
	;
	v625 = v623
	goto L135
L141:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v506)+12))
	v694 = v626
	v706 = v627
	goto L120
L142:
	;
	goto L143
L143:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v506)+8))
	if v628 == int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	if v625 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v506)+12))
	if v631 != 0 {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v506)+12)) = v587
	goto L144
L147:
	;
	goto L134
L148:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v635)+16))
	v637 = F_cdissect(m, v509, v636, v587, v625)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L11
	} else {
		goto L149
	}
L149:
	;
	if v637 != int32(1) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	if v637 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	goto L152
L152:
	;
	if v556 != 0 {
		goto L162
	} else {
		goto L163
	}
L153:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v509)+12))
	if v643 != 0 {
		goto L156
	} else {
		goto L157
	}
L154:
	;
	goto L155
L155:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v509)+36))
	if v658 != 0 {
		goto L159
	} else {
		goto L160
	}
L156:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v509)+16))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v509)+24))
	v647 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v644))) = (v587 - v645) >> (uint(v647) % 32)
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v509)+16))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v509)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v650)+4)) = (v625 - v651) >> (uint(v647) % 32)
	goto L158
L157:
	;
	goto L158
L158:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v506)+12))
	v694 = int32(0)
	v706 = v657
	goto L120
L159:
	;
	v659 = v658
	goto L161
L160:
	;
	v659 = v637
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v509)+36)) = v659
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v506)+12))
	v694 = v637
	v706 = v661
	goto L120
L162:
	;
	if v625 == v615 {
		goto L147
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	if v587 == v625 {
		goto L147
	} else {
		goto L166
	}
L165:
	;
	v603 = v625 + int32(4)
	goto L133
L166:
	;
	v615 = v625 - int32(4)
	goto L133
L167:
	;
	goto L132
L168:
	;
	goto L122
L169:
	;
	v725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522)+68)))
	if v725 == int32(1) {
		goto L185
	} else {
		goto L186
	}
L170:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v522)+20))
	if v710 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	F_pfree(m, v710)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L11
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v522)+24))
	if v713 != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	goto L173
L175:
	;
	F_pfree(m, v713)
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L11
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v522)+32))
	if v716 != 0 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	goto L177
L179:
	;
	F_pfree(m, v716)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L11
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v522)+36))
	if v719 == int32(0) {
		goto L169
	} else {
		goto L183
	}
L182:
	;
	goto L181
L183:
	;
	F_pfree(m, v719)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L11
	} else {
		goto L184
	}
L184:
	;
	goto L169
L185:
	;
	F_pfree(m, v522)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L11
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+69)))
	if v730 != int32(1) {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	goto L187
L189:
	;
	v748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+68)))
	if v748 == int32(1) {
		goto L205
	} else {
		goto L206
	}
L190:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v515)+20))
	if v733 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	F_pfree(m, v733)
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L11
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v515)+24))
	if v736 != 0 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	goto L193
L195:
	;
	F_pfree(m, v736)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L11
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v515)+32))
	if v739 != 0 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	goto L197
L199:
	;
	F_pfree(m, v739)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L11
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v515)+36))
	if v742 == int32(0) {
		goto L189
	} else {
		goto L203
	}
L202:
	;
	goto L201
L203:
	;
	F_pfree(m, v742)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L11
	} else {
		goto L204
	}
L204:
	;
	goto L189
L205:
	;
	F_pfree(m, v515)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L11
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v509)+36))
	if v753 != 0 {
		v779 = v753
		goto L91
	} else {
		goto L209
	}
L208:
	;
	goto L207
L209:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v754)+5)))
	if v755&int32(2) != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v509)+20))
	if v706 != 0 {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	goto L212
L212:
	;
	v779 = v694
	goto L91
L213:
	;
	v760 = v706
	goto L215
L214:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v509)+32))
	v760 = v759
	goto L215
L215:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v509)+24))
	v763 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v758))) = (v760 - v761) >> (uint(v763) % 32)
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v509)+20))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v509)+32))
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v509)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v766)+4)) = (v767 - v768) >> (uint(v763) % 32)
	goto L212
L216:
	;
	m.G0 = v792 + int32(16)
	v1047 = v1025
	goto L87
L217:
	;
	if v803 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v795)+36))
	v1025 = v807
	goto L216
L219:
	;
	goto L220
L220:
	;
	v808 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v792)+12)) = v808
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v795)+28))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v795)+32))
	v815 = F_shortest(m, v795, v803, v810, v810, v811, v792+int32(12), v808)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L11
	} else {
		goto L221
	}
L221:
	;
	v817 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+69)))
	if v817 != int32(1) {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v803)+68)))
	if v835 == int32(1) {
		goto L238
	} else {
		goto L239
	}
L223:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v803)+20))
	if v820 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	F_pfree(m, v820)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L11
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v803)+24))
	if v823 != 0 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	goto L226
L228:
	;
	F_pfree(m, v823)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L11
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v803)+32))
	if v826 != 0 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	goto L230
L232:
	;
	F_pfree(m, v826)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L11
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v803)+36))
	if v829 == int32(0) {
		goto L222
	} else {
		goto L236
	}
L235:
	;
	goto L234
L236:
	;
	F_pfree(m, v829)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L11
	} else {
		goto L237
	}
L237:
	;
	goto L222
L238:
	;
	F_pfree(m, v803)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L11
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v795)+36))
	if v840 != 0 {
		v1025 = v840
		goto L216
	} else {
		goto L242
	}
L241:
	;
	goto L240
L242:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v795)+4))
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v841)+5)))
	if v842&int32(2) != 0 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v795)+20))
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v792)+12))
	if v846 != 0 {
		goto L246
	} else {
		goto L247
	}
L244:
	;
	goto L245
L245:
	;
	if v815 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L246:
	;
	v848 = v846
	goto L248
L247:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v795)+32))
	v848 = v847
	goto L248
L248:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v795)+24))
	v851 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v845))) = (v848 - v849) >> (uint(v851) % 32)
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v795)+20))
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v795)+32))
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v795)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v854)+4)) = (v855 - v856) >> (uint(v851) % 32)
	goto L245
L249:
	;
	v1025 = int32(1)
	goto L216
L250:
	;
	goto L251
L251:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v795)+12))
	if v865 == int32(0) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1025 = int32(0)
	goto L216
L253:
	;
	goto L254
L254:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v792)+12))
	v870 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v792)+12)) = v870
	v873 = F_newdfa(m, v795, v503, v500, v802)
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L11
	} else {
		goto L255
	}
L255:
	;
	if v873 != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	if base.Ui32(v815) < base.Ui32(v869) {
		v946 = v869
		v949 = v870
		goto L259
	} else {
		goto L260
	}
L257:
	;
	goto L258
L258:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v795)+36))
	v1025 = v1015
	goto L216
L259:
	;
	v951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873)+69)))
	if v951 != int32(1) {
		goto L297
	} else {
		goto L298
	}
L260:
	;
	v887 = v869
	goto L261
L261:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v795)+32))
	if v798&int32(2) != 0 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	v946 = v934
	v949 = int32(0)
	goto L259
L263:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v795)+36))
	if v903 != 0 {
		goto L269
	} else {
		goto L270
	}
L264:
	;
	v896 = F_shortest(m, v795, v873, v887, v887, v892, int32(0), v792+int32(8))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L11
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v900 = F_longest(m, v795, v873, v887, v892, v792+int32(8))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L11
	} else {
		goto L268
	}
L267:
	;
	v902 = v896
	goto L263
L268:
	;
	v902 = v900
	goto L263
L269:
	;
	v904 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873)+69)))
	if v904 != int32(1) {
		goto L272
	} else {
		goto L273
	}
L270:
	;
	goto L271
L271:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v792)+8))
	if v928 == int32(0) {
		goto L292
	} else {
		goto L293
	}
L272:
	;
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873)+68)))
	if v922 == int32(1) {
		goto L288
	} else {
		goto L289
	}
L273:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v873)+20))
	if v907 != 0 {
		goto L274
	} else {
		goto L275
	}
L274:
	;
	F_pfree(m, v907)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L11
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v873)+24))
	if v910 != 0 {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	goto L276
L278:
	;
	F_pfree(m, v910)
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L11
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v873)+32))
	if v913 != 0 {
		goto L282
	} else {
		goto L283
	}
L281:
	;
	goto L280
L282:
	;
	F_pfree(m, v913)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L11
	} else {
		goto L285
	}
L283:
	;
	goto L284
L284:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v873)+36))
	if v916 == int32(0) {
		goto L272
	} else {
		goto L286
	}
L285:
	;
	goto L284
L286:
	;
	F_pfree(m, v916)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L11
	} else {
		goto L287
	}
L287:
	;
	goto L272
L288:
	;
	F_pfree(m, v873)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L11
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v795)+36))
	v1025 = v927
	goto L216
L291:
	;
	goto L290
L292:
	;
	if v902 != 0 {
		v946 = v887
		v949 = v902
		goto L259
	} else {
		goto L295
	}
L293:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v792)+12))
	if v931 != 0 {
		goto L292
	} else {
		goto L294
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v792)+12)) = v887
	goto L292
L295:
	;
	v934 = v887 + int32(4)
	if base.Ui32(v934) <= base.Ui32(v815) {
		v887 = v934
		goto L261
	} else {
		goto L296
	}
L296:
	;
	goto L262
L297:
	;
	v969 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v873)+68)))
	if v969 == int32(1) {
		goto L313
	} else {
		goto L314
	}
L298:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v873)+20))
	if v954 != 0 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	F_pfree(m, v954)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L11
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v873)+24))
	if v957 != 0 {
		goto L303
	} else {
		goto L304
	}
L302:
	;
	goto L301
L303:
	;
	F_pfree(m, v957)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L11
	} else {
		goto L306
	}
L304:
	;
	goto L305
L305:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v873)+32))
	if v960 != 0 {
		goto L307
	} else {
		goto L308
	}
L306:
	;
	goto L305
L307:
	;
	F_pfree(m, v960)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L11
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v873)+36))
	if v963 == int32(0) {
		goto L297
	} else {
		goto L311
	}
L310:
	;
	goto L309
L311:
	;
	F_pfree(m, v963)
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L11
	} else {
		goto L312
	}
L312:
	;
	goto L297
L313:
	;
	F_pfree(m, v873)
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L11
	} else {
		goto L316
	}
L314:
	;
	goto L315
L315:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v795)+16))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v795)+24))
	v977 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v974))) = (v946 - v975) >> (uint(v977) % 32)
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v795)+16))
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v795)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v980)+4)) = (v949 - v981) >> (uint(v977) % 32)
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v795)+4))
	v987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986)+5)))
	if v987&v977 != 0 {
		goto L317
	} else {
		goto L318
	}
L316:
	;
	goto L315
L317:
	;
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v795)+20))
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v792)+12))
	if v991 != 0 {
		goto L320
	} else {
		goto L321
	}
L318:
	;
	goto L319
L319:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v795)+12))
	if v1007 == int32(1) {
		goto L323
	} else {
		goto L324
	}
L320:
	;
	v993 = v991
	goto L322
L321:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v795)+32))
	v993 = v992
	goto L322
L322:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v795)+24))
	v996 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v990))) = (v993 - v994) >> (uint(v996) % 32)
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v795)+20))
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v795)+32))
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v795)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v999)+4)) = (v1000 - v1001) >> (uint(v996) % 32)
	goto L319
L323:
	;
	v1025 = int32(0)
	goto L216
L324:
	;
	goto L325
L325:
	;
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v795)+4))
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v1011)+16))
	v1013 = F_cdissect(m, v795, v1012, v946, v949)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L11
	} else {
		goto L326
	}
L326:
	;
	v1025 = v1013
	goto L216
L327:
	;
	if v1047 != 0 {
		v1080 = v1047
		goto L38
	} else {
		goto L328
	}
L328:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v17)+336))
	if l5 != v1050 {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1053 = v108 << (uint(int32(3)) % 32)
	if v1053 != 0 {
		goto L333
	} else {
		goto L334
	}
L330:
	;
	goto L331
L331:
	;
	v1056 = int32(0)
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v17)+324))
	v1058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1057)+4)))
	if v1058&int32(16) == v1056 {
		v1080 = v1056
		goto L38
	} else {
		goto L336
	}
L332:
	;
	goto L331
L333:
	;
	v1054 = F__emscripten_memcpy_bulkmem(m, l5, v1050, v1053)
	mBase = m.M
	goto L335
L334:
	;
	goto L335
L335:
	;
	goto L332
L336:
	;
	if v108 == int32(1) {
		v1080 = v1056
		goto L38
	} else {
		goto L337
	}
L337:
	;
	v1065 = int32(8)
	v1073 = F__emscripten_memset_bulkmem(m, l5+v1065, base.I32_extend8_s(int32(255)), v108<<(uint(int32(3))%32)-v1065)
	mBase = m.M
	goto L338
L338:
	;
	v1080 = v1056
	goto L38
L339:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v17)+360))
	if v1095 == int32(0) {
		goto L343
	} else {
		goto L344
	}
L340:
	;
	if v1088 == v17+int32(160) {
		goto L339
	} else {
		goto L341
	}
L341:
	;
	F_pfree(m, v1088)
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L11
	} else {
		goto L342
	}
L342:
	;
	goto L339
L343:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v17)+364))
	if v1181 != 0 {
		goto L373
	} else {
		goto L374
	}
L344:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v17)+324))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1098)+68))
	if v1099 != 0 {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1102 = int32(0)
	goto L348
L346:
	;
	v1150 = v1095
	goto L347
L347:
	;
	if v1150 == v17 {
		goto L343
	} else {
		goto L371
	}
L348:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v17)+360))
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1115+v1102<<(uint(int32(2))%32))))
	if v1119 == int32(0) {
		goto L350
	} else {
		goto L351
	}
L349:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v17)+360))
	v1150 = v1149
	goto L347
L350:
	;
	v1147 = v1102 + int32(1)
	if v1147 != v1099 {
		v1102 = v1147
		goto L348
	} else {
		goto L370
	}
L351:
	;
	v1122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1119)+69)))
	if v1122 != int32(1) {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v1140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1119)+68)))
	if v1140 != int32(1) {
		goto L350
	} else {
		goto L368
	}
L353:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+20))
	if v1125 != 0 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	F_pfree(m, v1125)
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L11
	} else {
		goto L357
	}
L355:
	;
	goto L356
L356:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+24))
	if v1128 != 0 {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	goto L356
L358:
	;
	F_pfree(m, v1128)
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L11
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+32))
	if v1131 != 0 {
		goto L362
	} else {
		goto L363
	}
L361:
	;
	goto L360
L362:
	;
	F_pfree(m, v1131)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L11
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1119)+36))
	if v1134 == int32(0) {
		goto L352
	} else {
		goto L366
	}
L365:
	;
	goto L364
L366:
	;
	F_pfree(m, v1134)
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L11
	} else {
		goto L367
	}
L367:
	;
	goto L352
L368:
	;
	F_pfree(m, v1119)
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L11
	} else {
		goto L369
	}
L369:
	;
	goto L350
L370:
	;
	goto L349
L371:
	;
	F_pfree(m, v1150)
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L11
	} else {
		goto L372
	}
L372:
	;
	goto L343
L373:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v17)+324))
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1182)+428))
	if v1183 != 0 {
		goto L376
	} else {
		goto L377
	}
L374:
	;
	goto L375
L375:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v17)+368))
	if v1265 != 0 {
		goto L403
	} else {
		goto L404
	}
L376:
	;
	v1186 = int32(0)
	goto L379
L377:
	;
	v1248 = v1181
	goto L378
L378:
	;
	F_pfree(m, v1248)
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L11
	} else {
		goto L402
	}
L379:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v17)+364))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1199+v1186<<(uint(int32(2))%32))))
	if v1203 == int32(0) {
		goto L381
	} else {
		goto L382
	}
L380:
	;
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(v17)+364))
	v1248 = v1233
	goto L378
L381:
	;
	v1231 = v1186 + int32(1)
	if v1231 != v1183 {
		v1186 = v1231
		goto L379
	} else {
		goto L401
	}
L382:
	;
	v1206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1203)+69)))
	if v1206 != int32(1) {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v1224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1203)+68)))
	if v1224 != int32(1) {
		goto L381
	} else {
		goto L399
	}
L384:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1203)+20))
	if v1209 != 0 {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	F_pfree(m, v1209)
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L11
	} else {
		goto L388
	}
L386:
	;
	goto L387
L387:
	;
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1203)+24))
	if v1212 != 0 {
		goto L389
	} else {
		goto L390
	}
L388:
	;
	goto L387
L389:
	;
	F_pfree(m, v1212)
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L11
	} else {
		goto L392
	}
L390:
	;
	goto L391
L391:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1203)+32))
	if v1215 != 0 {
		goto L393
	} else {
		goto L394
	}
L392:
	;
	goto L391
L393:
	;
	F_pfree(m, v1215)
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L11
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1203)+36))
	if v1218 == int32(0) {
		goto L383
	} else {
		goto L397
	}
L396:
	;
	goto L395
L397:
	;
	F_pfree(m, v1218)
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L11
	} else {
		goto L398
	}
L398:
	;
	goto L383
L399:
	;
	F_pfree(m, v1203)
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L11
	} else {
		goto L400
	}
L400:
	;
	goto L381
L401:
	;
	goto L380
L402:
	;
	goto L375
L403:
	;
	F_pfree(m, v1265)
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L11
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v17)+372))
	if v1268 == int32(0) {
		v1279 = v1080
		goto L1
	} else {
		goto L407
	}
L406:
	;
	goto L405
L407:
	;
	F_pfree(m, v1268)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L11
	} else {
		goto L408
	}
L408:
	;
	v1279 = v1080
	goto L1
}
func F_pg_safe_snapshot_blocking_pids(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_pg_safe_snapshot_blocking_pids[0]))
	v16 = F_palloc(m, v13<<(uint(int32(2))%32))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_pg_safe_snapshot_blocking_pids[0]))
		v22 = F_GetSafeSnapshotBlockingPids(m, v11, v16, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			if v22 <= int32(0) {
				v111 = v2
				v120 = F_construct_array_builtin(m, v111, v22, int32(23))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
					return int32(0)
				} else {
					return v120
				}
			} else {
				v27 = v22 & int32(3)
				v30 = F_palloc(m, v22<<(uint(int32(2))%32))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v22) {
						v37 = v32
						v45 = v2
						for {
							v48 = v37 << (uint(int32(2)) % 32)
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v48+v16)))
							*(*int32)(unsafe.Add(mBase, uint32(v30+v48))) = v51
							v53 = int32(4)
							v54 = v48 | v53
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v16+v54)))
							*(*int32)(unsafe.Add(mBase, uint32(v30+v54))) = v57
							v60 = v48 | int32(8)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v16+v60)))
							*(*int32)(unsafe.Add(mBase, uint32(v30+v60))) = v63
							v66 = v48 | int32(12)
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v66+v16)))
							*(*int32)(unsafe.Add(mBase, uint32(v30+v66))) = v69
							v72 = v37 + v53
							v74 = v45 + v53
							if v74 != v22&int32(2147483644) {
								v37 = v72
								v45 = v74
								continue
							} else {
								break
							}
							break
						}
						v76 = v72
					} else {
						v76 = v32
					}
					if v27 == int32(0) {
						v111 = v30
					} else {
						v88 = v76
						v95 = v2
						for {
							v99 = v88 << (uint(int32(2)) % 32)
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v99+v16)))
							*(*int32)(unsafe.Add(mBase, uint32(v30+v99))) = v102
							v104 = int32(1)
							v107 = v95 + v104
							if v107 != v27 {
								v88 = v88 + v104
								v95 = v107
								continue
							} else {
								break
							}
							break
						}
						v111 = v30
					}
					v120 = F_construct_array_builtin(m, v111, v22, int32(23))
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						return v120
					}
				}
			}
		}
	}
}
func F_pg_set_noblock(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(2048)
	m.G0 = v5 + int32(16)
	return int32(1)
}
func F_pg_size_bytes(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v131 int32
	_ = v131
	var v136 int64
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v337 int32
	_ = v337
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v382 int32
	_ = v382
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v405 int32
	_ = v405
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v427 int32
	_ = v427
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v472 int32
	_ = v472
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v495 int32
	_ = v495
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v518 int32
	_ = v518
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int64
	_ = v551
	var v552 int64
	_ = v552
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int64
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
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
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum_packed(m, v14)
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
	v19 = F_text_to_cstring(m, v15)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = v19
	goto L5
L4:
	;
	v46 = v43 - int32(48)
	if base.Ui32(v46&int32(255)) <= base.Ui32(int32(9)) {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if base.Ui32(v30-int32(9)) < base.Ui32(int32(5)) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v41 = v24 + int32(1)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v43 = v42
	v44 = v41
	goto L4
L7:
	;
	goto L6
L8:
	;
	v24 = v24 + int32(1)
	goto L5
L9:
	;
	switch v30 - int32(32) {
	case 0:
		goto L8
	default:
		v43 = v30
		v44 = v24
		goto L4
	case 11, 13:
		goto L7
	}
L10:
	;
	v52 = v44
	goto L13
L11:
	;
	v69 = v43
	v70 = v44
	goto L12
L12:
	;
	if v69&int32(255) != int32(46) {
		v109 = v69
		v110 = v70
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v61 = v52 + int32(1)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if base.Ui32((v62-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v52 = v61
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v69 = v62
	v70 = v61
	goto L12
L15:
	;
	goto L14
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L175
	}
L17:
	;
	if (v115|int32(32))&int32(255) == int32(101) {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v46&int32(255)) {
		goto L16
	} else {
		goto L24
	}
L19:
	;
	v83 = v70 + int32(1)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if base.Ui32(int32(9)) < base.Ui32((v84-int32(48))&int32(255)) {
		v109 = v84
		v110 = v83
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v92 = v83
	goto L21
L21:
	;
	v101 = v92 + int32(1)
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if base.Ui32((v102-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v92 = v101
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v115 = v102
	v116 = v101
	goto L17
L23:
	;
	goto L22
L24:
	;
	v115 = v109
	v116 = v110
	goto L17
L25:
	;
	v131 = v116 + int32(1)
	v136 = F_strtox_2(m, v131, v12+int32(44), int32(10), int64(2147483648))
	mBase = m.M
	goto L28
L26:
	;
	v142 = v115
	v143 = v116
	goto L27
L27:
	;
	v145 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v143))) = uint8(v145)
	v151 = F_DirectFunctionCall3Coll(m, int32(408), v145, v24, v145, int32(-1))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L32
	}
L28:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	if base.Ui32(v131) < base.Ui32(v138) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v140 = v138
	goto L31
L30:
	;
	v140 = v116
	goto L31
L31:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	v142 = v141
	v143 = v140
	goto L27
L32:
	;
	v153 = F_pg_detoast_datum(m, v151)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v143))) = uint8(v142)
	v156 = v142
	v157 = v143
	goto L34
L34:
	;
	v166 = v156 & int32(255)
	if base.Ui32(v166-int32(9)) < base.Ui32(int32(5)) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L168
	}
L36:
	;
	goto L35
L37:
	;
	v584 = v157 + int32(1)
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584))))
	v156 = v585
	v157 = v584
	goto L34
L38:
	;
	if v166 == int32(32) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	if v166 == int32(0) {
		v569 = v153
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v574 = F_DirectFunctionCall1Coll(m, int32(1277), int32(0), v569)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L166
	}
L41:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v175 == int32(1) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v207 = v205 + v19
	goto L53
L43:
	;
	v178 = int32(4)
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v180&int32(254) == int32(2) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v193 = int32(1)
	if v175&v193 != 0 {
		v205 = int32(base.Ui32(v175)>>(uint(v193)%32)) - v193
		goto L42
	} else {
		goto L52
	}
L46:
	;
	v189 = v178
	goto L48
L47:
	;
	v189 = base.B2i32(v180 == int32(18)) << (uint(v178) % 32)
	goto L48
L48:
	;
	if v180 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v192 = v178
	goto L51
L50:
	;
	v192 = v189
	goto L51
L51:
	;
	v205 = v192
	goto L42
L52:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v205 = int32(base.Ui32(v199)>>(uint(int32(2))%32)) - int32(4)
	goto L42
L53:
	;
	v217 = v207 - int32(1)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if base.Ui32(v218-int32(9)) < base.Ui32(int32(5)) {
		v207 = v217
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v225 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v207))) = uint8(v225)
	v232 = v157
	v233 = int32(_a_F_pg_size_bytes_0)
	goto L59
L55:
	;
	if v218 == int32(32) {
		v207 = v217
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v551 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v550)+9)))
	v552 = int64(1) << (uint(v551) % 64)
	if v552 < int64(2) {
		v569 = v153
		goto L40
	} else {
		goto L162
	}
L58:
	;
	if v270 == int32(0) {
		v550 = int32(_a_F_pg_size_bytes_1)
		goto L57
	} else {
		goto L71
	}
L59:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if v236 == v237 {
		v259 = v236
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v270 = int32(0)
	goto L58
L61:
	;
	v261 = int32(1)
	if v259 != 0 {
		v232 = v232 + v261
		v233 = v233 + v261
		goto L59
	} else {
		goto L70
	}
L62:
	;
	if base.Ui32((v236-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v247 = v236 | int32(32)
	goto L65
L64:
	;
	v247 = v236
	goto L65
L65:
	;
	if base.Ui32((v237-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v256 = v237 | int32(32)
	goto L68
L67:
	;
	v256 = v237
	goto L68
L68:
	;
	if v247 == v256 {
		v259 = v247
		goto L61
	} else {
		goto L69
	}
L69:
	;
	v270 = v247 - v256
	goto L58
L70:
	;
	goto L60
L71:
	;
	v277 = v157
	v278 = int32(_a_F_pg_size_bytes_2)
	goto L73
L72:
	;
	if v315 == int32(0) {
		v550 = int32(_a_F_pg_size_bytes_3)
		goto L57
	} else {
		goto L85
	}
L73:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	if v281 == v282 {
		v304 = v281
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v315 = int32(0)
	goto L72
L75:
	;
	v306 = int32(1)
	if v304 != 0 {
		v277 = v277 + v306
		v278 = v278 + v306
		goto L73
	} else {
		goto L84
	}
L76:
	;
	if base.Ui32((v281-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v292 = v281 | int32(32)
	goto L79
L78:
	;
	v292 = v281
	goto L79
L79:
	;
	if base.Ui32((v282-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v301 = v282 | int32(32)
	goto L82
L81:
	;
	v301 = v282
	goto L82
L82:
	;
	if v292 == v301 {
		v304 = v292
		goto L75
	} else {
		goto L83
	}
L83:
	;
	v315 = v292 - v301
	goto L72
L84:
	;
	goto L74
L85:
	;
	v322 = v157
	v323 = int32(_a_F_pg_size_bytes_4)
	goto L87
L86:
	;
	if v360 == int32(0) {
		v550 = int32(_a_F_pg_size_bytes_5)
		goto L57
	} else {
		goto L99
	}
L87:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
	if v326 == v327 {
		v349 = v326
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v360 = int32(0)
	goto L86
L89:
	;
	v351 = int32(1)
	if v349 != 0 {
		v322 = v322 + v351
		v323 = v323 + v351
		goto L87
	} else {
		goto L98
	}
L90:
	;
	if base.Ui32((v326-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v337 = v326 | int32(32)
	goto L93
L92:
	;
	v337 = v326
	goto L93
L93:
	;
	if base.Ui32((v327-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v346 = v327 | int32(32)
	goto L96
L95:
	;
	v346 = v327
	goto L96
L96:
	;
	if v337 == v346 {
		v349 = v337
		goto L89
	} else {
		goto L97
	}
L97:
	;
	v360 = v337 - v346
	goto L86
L98:
	;
	goto L88
L99:
	;
	v367 = v157
	v368 = int32(_a_F_pg_size_bytes_6)
	goto L101
L100:
	;
	if v405 == int32(0) {
		v550 = int32(_a_F_pg_size_bytes_7)
		goto L57
	} else {
		goto L113
	}
L101:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367))))
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368))))
	if v371 == v372 {
		v394 = v371
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v405 = int32(0)
	goto L100
L103:
	;
	v396 = int32(1)
	if v394 != 0 {
		v367 = v367 + v396
		v368 = v368 + v396
		goto L101
	} else {
		goto L112
	}
L104:
	;
	if base.Ui32((v371-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v382 = v371 | int32(32)
	goto L107
L106:
	;
	v382 = v371
	goto L107
L107:
	;
	if base.Ui32((v372-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v391 = v372 | int32(32)
	goto L110
L109:
	;
	v391 = v372
	goto L110
L110:
	;
	if v382 == v391 {
		v394 = v382
		goto L103
	} else {
		goto L111
	}
L111:
	;
	v405 = v382 - v391
	goto L100
L112:
	;
	goto L102
L113:
	;
	v412 = v157
	v413 = int32(_a_F_pg_size_bytes_8)
	goto L115
L114:
	;
	if v450 == int32(0) {
		v550 = int32(_a_F_pg_size_bytes_9)
		goto L57
	} else {
		goto L127
	}
L115:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412))))
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v413))))
	if v416 == v417 {
		v439 = v416
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v450 = int32(0)
	goto L114
L117:
	;
	v441 = int32(1)
	if v439 != 0 {
		v412 = v412 + v441
		v413 = v413 + v441
		goto L115
	} else {
		goto L126
	}
L118:
	;
	if base.Ui32((v416-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v427 = v416 | int32(32)
	goto L121
L120:
	;
	v427 = v416
	goto L121
L121:
	;
	if base.Ui32((v417-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v436 = v417 | int32(32)
	goto L124
L123:
	;
	v436 = v417
	goto L124
L124:
	;
	if v427 == v436 {
		v439 = v427
		goto L117
	} else {
		goto L125
	}
L125:
	;
	v450 = v427 - v436
	goto L114
L126:
	;
	goto L116
L127:
	;
	v457 = v157
	v458 = int32(_a_F_pg_size_bytes_10)
	goto L129
L128:
	;
	if v495 == int32(0) {
		v550 = int32(_a_F_pg_size_bytes_11)
		goto L57
	} else {
		goto L141
	}
L129:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457))))
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458))))
	if v461 == v462 {
		v484 = v461
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v495 = int32(0)
	goto L128
L131:
	;
	v486 = int32(1)
	if v484 != 0 {
		v457 = v457 + v486
		v458 = v458 + v486
		goto L129
	} else {
		goto L140
	}
L132:
	;
	if base.Ui32((v461-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v472 = v461 | int32(32)
	goto L135
L134:
	;
	v472 = v461
	goto L135
L135:
	;
	if base.Ui32((v462-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v481 = v462 | int32(32)
	goto L138
L137:
	;
	v481 = v462
	goto L138
L138:
	;
	if v472 == v481 {
		v484 = v472
		goto L131
	} else {
		goto L139
	}
L139:
	;
	v495 = v472 - v481
	goto L128
L140:
	;
	goto L130
L141:
	;
	v503 = v157
	v504 = int32(_a_F_pg_size_bytes_12)
	goto L143
L142:
	;
	if v541 != 0 {
		goto L155
	} else {
		goto L156
	}
L143:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503))))
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504))))
	if v507 == v508 {
		v530 = v507
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v541 = int32(0)
	goto L142
L145:
	;
	v532 = int32(1)
	if v530 != 0 {
		v503 = v503 + v532
		v504 = v504 + v532
		goto L143
	} else {
		goto L154
	}
L146:
	;
	if base.Ui32((v507-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v518 = v507 | int32(32)
	goto L149
L148:
	;
	v518 = v507
	goto L149
L149:
	;
	if base.Ui32((v508-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v527 = v508 | int32(32)
	goto L152
L151:
	;
	v527 = v508
	goto L152
L152:
	;
	if v518 == v527 {
		v530 = v518
		goto L145
	} else {
		goto L153
	}
L153:
	;
	v541 = v518 - v527
	goto L142
L154:
	;
	goto L144
L155:
	;
	v542 = int32(_a_F_pg_size_bytes_13)
	goto L157
L156:
	;
	v542 = int32(_a_F_pg_size_bytes_1)
	goto L157
L157:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)))
	if v543 == int32(0) {
		goto L36
	} else {
		goto L158
	}
L158:
	;
	if v541 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v548 = int32(_a_F_pg_size_bytes_13)
	goto L161
L160:
	;
	v548 = int32(_a_F_pg_size_bytes_1)
	goto L161
L161:
	;
	v550 = v548
	goto L57
L162:
	;
	v557 = F_int64_to_numeric(m, v552)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v559 = F_DirectFunctionCall2Coll(m, int32(1278), int32(0), v557, v153)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v561 = F_pg_detoast_datum(m, v559)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v569 = v561
	goto L40
L166:
	;
	v576 = *(*int64)(unsafe.Add(mBase, uint32(v574)))
	v577 = F_Int64GetDatum(m, v576)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	m.G0 = v12 + int32(48)
	return v577
L168:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L169
	}
L169:
	;
	v593 = F_text_to_cstring(m, v15)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v593
	F_errmsg(m, int32(_a_F_pg_size_bytes_14), v12+int32(16))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v157
	F_errdetail(m, int32(_a_F_pg_size_bytes_15), v12)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	F_errhint(m, int32(_a_F_pg_size_bytes_16), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(_a_F_pg_size_bytes_17), int32(860), int32(_a_F_pg_size_bytes_18))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v19
	F_errmsg(m, int32(_a_F_pg_size_bytes_14), v12+int32(32))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_pg_size_bytes_17), int32(782), int32(_a_F_pg_size_bytes_18))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_snapshot_send(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v44 int64
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v87 int32
	_ = v87
	var v90 int64
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v149 int64
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int64
	_ = v156
	var v158 int64
	_ = v158
	var v160 int64
	_ = v160
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_pq_begintypsend(m, v9)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	F_enlargeStringInfo(m, v9, int32(4))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v25 = int32(24)
	v27 = int32(_a_F_pg_snapshot_send_0)
	v29 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v22+v23))) = v18<<(uint(v25)%32) | v18&v27<<(uint(v29)%32) | (int32(base.Ui32(v18)>>(uint(v29)%32))&v27 | int32(base.Ui32(v18)>>(uint(v25)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v22 + int32(4)
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	F_enlargeStringInfo(m, v9, v29)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v51 = int64(56)
	v53 = int64(65280)
	v55 = int64(40)
	v58 = int64(16711680)
	v60 = int64(24)
	v62 = int64(4278190080)
	v64 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v48+v49))) = v44<<(uint(v51)%64) | v44&v53<<(uint(v55)%64) | (v44&v58<<(uint(v60)%64) | v44&v62<<(uint(v64)%64)) | (int64(base.Ui64(v44)>>(uint(v64)%64))&v62 | int64(base.Ui64(v44)>>(uint(v60)%64))&v58 | (int64(base.Ui64(v44)>>(uint(v55)%64))&v53 | int64(base.Ui64(v44)>>(uint(v51)%64))))
	v87 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v48 + v87
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	F_enlargeStringInfo(m, v9, v87)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v97 = int64(56)
	v99 = int64(65280)
	v101 = int64(40)
	v104 = int64(16711680)
	v106 = int64(24)
	v108 = int64(4278190080)
	v110 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v94+v95))) = v90<<(uint(v97)%64) | v90&v99<<(uint(v101)%64) | (v90&v104<<(uint(v106)%64) | v90&v108<<(uint(v110)%64)) | (int64(base.Ui64(v90)>>(uint(v110)%64))&v108 | int64(base.Ui64(v90)>>(uint(v106)%64))&v104 | (int64(base.Ui64(v90)>>(uint(v101)%64))&v99 | int64(base.Ui64(v90)>>(uint(v97)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v94 + int32(8)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v136 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v140 = int32(0)
	goto L10
L8:
	;
	goto L9
L9:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v207 << (uint(int32(2)) % 32)
	goto L14
L10:
	;
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(24)+v140<<(uint(int32(3))%32))))
	F_enlargeStringInfo(m, v9, int32(8))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v156 = int64(56)
	v158 = int64(65280)
	v160 = int64(40)
	v163 = int64(16711680)
	v165 = int64(24)
	v167 = int64(4278190080)
	v169 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v153+v154))) = v149<<(uint(v156)%64) | v149&v158<<(uint(v160)%64) | (v149&v163<<(uint(v165)%64) | v149&v167<<(uint(v169)%64)) | (int64(base.Ui64(v149)>>(uint(v169)%64))&v167 | int64(base.Ui64(v149)>>(uint(v165)%64))&v163 | (int64(base.Ui64(v149)>>(uint(v160)%64))&v158 | int64(base.Ui64(v149)>>(uint(v156)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v153 + int32(8)
	v196 = v140 + int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if base.Ui32(v196) < base.Ui32(v197) {
		v140 = v196
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	m.G0 = v9 + int32(16)
	return v206
}
func F_pg_snapshot_xmax(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int64)(unsafe.Add(mBase, uint32(v3)+16))
		v8 = F_Int64GetDatum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_pg_snapshot_xmin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int64)(unsafe.Add(mBase, uint32(v3)+8))
		v8 = F_Int64GetDatum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_pg_strcasecmp(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v5 = l0
	v6 = l1
	goto L1
L1:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v9 == v10 {
		v33 = v9
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return int32(0)
L3:
	;
	v35 = int32(1)
	if v33 != 0 {
		v5 = v5 + v35
		v6 = v6 + v35
		goto L1
	} else {
		goto L12
	}
L4:
	;
	if base.Ui32((v9-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v20 = v9 | int32(32)
	goto L7
L6:
	;
	v20 = v9
	goto L7
L7:
	;
	if base.Ui32((v10-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v29 = v10 | int32(32)
	goto L10
L9:
	;
	v29 = v10
	goto L10
L10:
	;
	if v20 == v29 {
		v33 = v20
		goto L3
	} else {
		goto L11
	}
L11:
	;
	return v20 - v29
L12:
	;
	goto L2
}
func F_pg_strncasecmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v6 = l0
	v7 = l1
	v8 = l2
	goto L1
L1:
	;
	if v8 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return int32(0)
L3:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v11 == v12 {
		v35 = v11
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	v37 = int32(1)
	if v35 != 0 {
		v6 = v6 + v37
		v7 = v7 + v37
		v8 = v8 - v37
		goto L1
	} else {
		goto L15
	}
L7:
	;
	if base.Ui32((v11-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v22 = v11 | int32(32)
	goto L10
L9:
	;
	v22 = v11
	goto L10
L10:
	;
	if base.Ui32((v12-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = v12 | int32(32)
	goto L13
L12:
	;
	v31 = v12
	goto L13
L13:
	;
	if v22 == v31 {
		v35 = v22
		goto L6
	} else {
		goto L14
	}
L14:
	;
	return v22 - v31
L15:
	;
	goto L5
}
func F_pg_strsignal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v4 = int32(_a_F_pg_strsignal_0)
	if base.Ui32(l0-int32(65)) < base.Ui32(int32(-64)) {
		v21 = v4
	} else {
		v10 = l0
		v11 = v4
		for {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			v15 = v11 + int32(1)
			if v13 != 0 {
				v11 = v15
				continue
			} else {
			}
			v17 = v10 - int32(1)
			if v17 != 0 {
				v10 = v17
				v11 = v15
				continue
			} else {
				break
			}
			break
		}
		v21 = v15
	}
	if v21 != 0 {
		v23 = v21
	} else {
		v23 = int32(_a_F_pg_strsignal_1)
	}
	return v23
}
func F_pg_strtoint64_safe(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int64
	_ = v46
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v81 int64
	_ = v81
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v159 int64
	_ = v159
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v194 int64
	_ = v194
	var v195 int32
	_ = v195
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v234 int64
	_ = v234
	var v235 int32
	_ = v235
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int64
	_ = v272
	var v274 int32
	_ = v274
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int64
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v327 int64
	_ = v327
	var v333 int64
	_ = v333
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v392 int32
	_ = v392
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v425 int64
	_ = v425
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v17 = base.B2i32(v15 == int32(45))
	v18 = l0 + v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v21 = v19 - int32(48)
	if base.Ui32(int32(10)) <= base.Ui32(v21&int32(255)) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v425
L2:
	;
	v425 = int64(0)
	goto L1
L3:
	;
	F_errsave_finish(m, l1, int32(_a_F_pg_strtoint64_safe_0), v403, int32(_a_F_pg_strtoint64_safe_1))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L82
	} else {
		goto L91
	}
L4:
	;
	v378 = F_errsave_start(m, l1)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L82
	} else {
		goto L87
	}
L5:
	;
	v353 = F_errsave_start(m, l1)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L82
	} else {
		goto L83
	}
L6:
	;
	v101 = l0
	v104 = v15
	goto L23
L7:
	;
	v28 = base.I64_extend_i32_u(v21) & int64(255)
	v30 = v18 + int32(1)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	v33 = v31 - int32(48)
	if base.Ui32(v33&int32(255)) <= base.Ui32(int32(9)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v40 = v33
	v41 = v30
	v46 = v28
	goto L11
L9:
	;
	v68 = v31
	v72 = v28
	goto L10
L10:
	;
	if v68 != 0 {
		goto L6
	} else {
		goto L15
	}
L11:
	;
	if base.Ui64(int64(922337203685477580)) < base.Ui64(v46) {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v68 = v57
	v72 = v54
	goto L10
L13:
	;
	v54 = v46*int64(10) + base.I64_extend_i32_u(v40)&int64(255)
	v56 = v41 + int32(1)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	v59 = v57 - int32(48)
	if base.Ui32(v59&int32(255)) < base.Ui32(int32(10)) {
		v40 = v59
		v41 = v56
		v46 = v54
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	if v15 == int32(45) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v75 = int64(0)
	v81 = v75 - v72
	if v75-base.I64_extend_i32_u(base.B2i32(v72 != v75))^v81>>(uint(int64(63))%64) != v75 {
		goto L5
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v72 < int64(0) {
		goto L5
	} else {
		goto L20
	}
L19:
	;
	v425 = v81
	goto L1
L20:
	;
	v425 = v72
	goto L1
L21:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v125 != int32(48) {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	v123 = v101 + int32(1)
	v124 = v17
	goto L21
L23:
	;
	v108 = v104 & int32(255)
	if base.Ui32(v108-int32(9)) < base.Ui32(int32(5)) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v118 = int32(1)
	v123 = v101 + v118
	v124 = v118
	goto L21
L25:
	;
	goto L24
L26:
	;
	v116 = v101 + int32(1)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	v101 = v116
	v104 = v117
	goto L23
L27:
	;
	switch v108 - int32(32) {
	case 0:
		goto L26
	default:
		v123 = v101
		v124 = v17
		goto L21
	case 11:
		goto L22
	case 13:
		goto L25
	}
L28:
	;
	if v304 == v305 {
		goto L4
	} else {
		goto L70
	}
L29:
	;
	v266 = v123
	v268 = v125
	v272 = int64(0)
	goto L61
L30:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
	switch v128 - int32(66) {
	case 0, 32:
		goto L31
	default:
		goto L29
	case 13, 45:
		goto L32
	case 22, 54:
		goto L33
	}
L31:
	;
	v225 = v123 + int32(2)
	v228 = v225
	v234 = int64(0)
	goto L53
L32:
	;
	v185 = v123 + int32(2)
	v188 = v185
	v194 = int64(0)
	goto L45
L33:
	;
	v133 = v123 + int32(2)
	v136 = v133
	v142 = int64(0)
	goto L34
L34:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	goto L36
L35:
	;
	goto L4
L36:
	;
	if base.B2i32(base.Ui32(v143-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v143|int32(32)-int32(97)) < base.Ui32(int32(6))) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if base.Ui64(int64(576460752303423488)) < base.Ui64(v142) {
		goto L5
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if v143 != int32(95) {
		v304 = v136
		v305 = v133
		v306 = v143
		v310 = v142
		goto L28
	} else {
		goto L41
	}
L40:
	;
	v159 = int64(*(*int8)(unsafe.Add(mBase, uint32(v143)+uint32(_c_F_pg_strtoint64_safe[0]))))
	v136 = v136 + int32(1)
	v142 = v159 + v142<<(uint(int64(4))%64)
	goto L34
L41:
	;
	v168 = v136 + int32(1)
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v169 == int32(0) {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	goto L43
L43:
	;
	if base.B2i32(base.Ui32(v169-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v169|int32(32)-int32(97)) < base.Ui32(int32(6))) != 0 {
		v136 = v168
		goto L34
	} else {
		goto L44
	}
L44:
	;
	goto L35
L45:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v195&int32(248) == int32(48) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L4
L47:
	;
	if base.Ui64(int64(1152921504606846976)) < base.Ui64(v194) {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v195 != int32(95) {
		v304 = v188
		v305 = v185
		v306 = v195
		v310 = v194
		goto L28
	} else {
		goto L51
	}
L50:
	;
	v188 = v188 + int32(1)
	v194 = base.I64_extend_i32_u(v195-int32(48))&int64(255) | v194<<(uint(int64(3))%64)
	goto L45
L51:
	;
	v215 = v188 + int32(1)
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if base.Ui32(int32(248)) <= base.Ui32((v216-int32(56))&int32(255)) {
		v188 = v215
		goto L45
	} else {
		goto L52
	}
L52:
	;
	goto L46
L53:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	if v235&int32(254) == int32(48) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L4
L55:
	;
	if base.Ui64(int64(4611686018427387904)) < base.Ui64(v234) {
		goto L5
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if v235 != int32(95) {
		v304 = v228
		v305 = v225
		v306 = v235
		v310 = v234
		goto L28
	} else {
		goto L59
	}
L58:
	;
	v228 = v228 + int32(1)
	v234 = base.I64_extend_i32_u(v235-int32(48))&int64(255) | v234<<(uint(int64(1))%64)
	goto L53
L59:
	;
	v255 = v228 + int32(1)
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	if base.Ui32(int32(254)) <= base.Ui32((v256-int32(50))&int32(255)) {
		v228 = v255
		goto L53
	} else {
		goto L60
	}
L60:
	;
	goto L54
L61:
	;
	v274 = v268 - int32(48)
	if base.Ui32(v274&int32(255)) <= base.Ui32(int32(9)) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L4
L63:
	;
	if base.Ui64(int64(922337203685477580)) < base.Ui64(v272) {
		goto L5
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	if v268 != int32(95) {
		v304 = v266
		v305 = v123
		v306 = v268
		v310 = v272
		goto L28
	} else {
		goto L67
	}
L66:
	;
	v288 = v266 + int32(1)
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288))))
	v266 = v288
	v268 = v289
	v272 = v272*int64(10) + base.I64_extend_i32_u(v274)&int64(255)
	goto L61
L67:
	;
	if v266 == v123 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v294 = v266 + int32(1)
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if base.Ui32((v295-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v266 = v294
		v268 = v295
		goto L61
	} else {
		goto L69
	}
L69:
	;
	goto L62
L70:
	;
	v314 = v304
	v316 = v306
	goto L71
L71:
	;
	if base.Ui32(v316-int32(9)) < base.Ui32(int32(5)) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v342 = v314 + int32(1)
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342))))
	v314 = v342
	v316 = v343
	goto L71
L74:
	;
	if v316 == int32(32) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	if v316 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	if v124 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v327 = int64(0)
	v333 = v327 - v310
	if v327-base.I64_extend_i32_u(base.B2i32(v310 != v327))^v333>>(uint(int64(63))%64) != v327 {
		goto L5
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if int64(0) <= v310 {
		v425 = v310
		goto L1
	} else {
		goto L81
	}
L80:
	;
	v425 = v333
	goto L1
L81:
	;
	goto L5
L82:
	;
	return int64(0)
L83:
	;
	if v353 == int32(0) {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(_a_F_pg_strtoint64_safe_2)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg(m, int32(_a_F_pg_strtoint64_safe_3), v12)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v403 = int32(873)
	goto L3
L87:
	;
	if v378 == int32(0) {
		goto L2
	} else {
		goto L88
	}
L88:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L82
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_pg_strtoint64_safe_2)
	F_errmsg(m, int32(_a_F_pg_strtoint64_safe_4), v12+int32(16))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L82
	} else {
		goto L90
	}
L90:
	;
	v403 = int32(879)
	goto L3
L91:
	;
	goto L2
}
func F_pg_strtok(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_pg_strtok[0]))
	v9 = v7
	goto L5
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_strtok[0])) = v62
	return v60
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	v60 = v56
	v62 = v55
	goto L1
L3:
	;
	v42 = v40 - v9
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v42
	if v42 != int32(2) {
		v60 = v9
		v62 = v40
		goto L1
	} else {
		goto L19
	}
L4:
	;
	v40 = v9 + int32(1)
	goto L3
L5:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	switch v13 {
	case 0:
		v55 = v9
		v56 = int32(0)
		goto L2
	case 1, 2, 3, 4, 5, 6, 7, 8, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33, 34, 35, 36, 37, 38, 39:
		goto L7
	case 9, 10, 32:
		goto L8
	case 40, 41:
		goto L4
	default:
		goto L9
	}
L6:
	;
	v20 = v13
	v21 = v9
	goto L10
L7:
	;
	goto L6
L8:
	;
	v9 = v9 + int32(1)
	goto L5
L9:
	;
	switch v13 - int32(123) {
	case 0, 2:
		goto L4
	default:
		goto L7
	}
L10:
	;
	v24 = v20 & int32(255)
	if v24 != int32(92) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v33 = v32 + v21
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v20 = v34
	v21 = v33
	goto L10
L13:
	;
	v32 = int32(1)
	goto L12
L14:
	;
	switch v24 {
	case 0, 9, 10, 32, 40, 41:
		v40 = v21
		goto L3
	case 1, 2, 3, 4, 5, 6, 7, 8, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33, 34, 35, 36, 37, 38, 39:
		goto L13
	default:
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v30 != 0 {
		v32 = int32(2)
		goto L12
	} else {
		goto L18
	}
L17:
	;
	switch v24 - int32(123) {
	case 0, 2:
		v40 = v21
		goto L3
	default:
		goto L13
	}
L18:
	;
	goto L13
L19:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v46 != int32(60) {
		v60 = v9
		v62 = v40
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v49 != int32(62) {
		v60 = v9
		v62 = v40
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v55 = v40
	v56 = v9
	goto L2
}
func F_pg_tablespace_size_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_SearchSysCacheExists(m, int32(69), v10, v2, v2, v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v18 = F_calculate_tablespace_size(m, v10)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 < int64(0) {
					v22 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
					v27 = int32(0)
					m.G0 = v7 + int32(16)
					return v27
				} else {
					v25 = F_Int64GetDatum(m, v18)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = v25
						m.G0 = v7 + int32(16)
						return v27
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(67137668))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v10
					F_errmsg(m, int32(_a_F_pg_tablespace_size_oid_0), v7)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_tablespace_size_oid_1), int32(293), int32(_a_F_pg_tablespace_size_oid_2))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
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
func F_pg_timezone_abbrev_is_known(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	v6 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l4)+268))
	if v12 <= v6 {
		v124 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v124
L2:
	;
	v16 = l4 + int32(_a_F_pg_timezone_abbrev_is_known_0)
	v26 = v6
	goto L3
L3:
	;
	v28 = v16 + v26
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v32 == int32(0) {
		v51 = v31
		v52 = v32
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l4)+264))
	if v70 <= int32(0) {
		v124 = v6
		goto L1
	} else {
		goto L20
	}
L5:
	;
	if v52-v51 != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	goto L5
L7:
	;
	if v31 != v32 {
		v51 = v31
		v52 = v32
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v36 = l0
	v37 = v28
	goto L9
L9:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v41 == int32(0) {
		v51 = v40
		v52 = v41
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v51 = v40
	v52 = v41
	goto L6
L11:
	;
	v44 = int32(1)
	if v40 == v41 {
		v36 = v36 + v44
		v37 = v37 + v44
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v59 = v26
	goto L16
L14:
	;
	goto L15
L15:
	;
	goto L4
L16:
	;
	v67 = v59 + int32(1)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v16))))
	if v68 != 0 {
		v59 = v67
		goto L16
	} else {
		goto L18
	}
L17:
	;
	if v67 < v12 {
		v26 = v67
		goto L3
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v124 = v6
	goto L1
L20:
	;
	v81 = int32(0)
	v82 = v6
	v84 = v70
	goto L21
L21:
	;
	v89 = l4 + int32(_a_F_pg_timezone_abbrev_is_known_1) + v81<<(uint(int32(4))%32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	if v90 != v26 {
		v113 = v82
		v114 = v84
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v124 = v113
	goto L1
L23:
	;
	v116 = v81 + int32(1)
	if v116 < v114 {
		v81 = v116
		v82 = v113
		v84 = v114
		goto L21
	} else {
		goto L32
	}
L24:
	;
	if v82 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v94 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v94)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v97
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l4)+264))
	v113 = v94
	v114 = v101
	goto L23
L26:
	;
	goto L27
L27:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v102 == v103 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
	if v106 == v107 {
		v113 = int32(1)
		v114 = v84
		goto L23
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v110 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v110)
	v124 = int32(1)
	goto L1
L31:
	;
	goto L30
L32:
	;
	goto L22
}
func F_pg_ts_parser_is_visible(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_TSParserIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v23 = v2
		} else {
			v23 = v14
		}
		m.G0 = v7 + int32(16)
		return v23
	}
}
func F_pg_utf8_increment(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	switch l1 - int32(1) {
	case 0:
		v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		v51 = v50
		v55 = v51 & int32(255)
		v57 = v55 - int32(223)
		if int32(1)<<(uint(v57)%32)&int32(_a_F_pg_utf8_increment_0) != 0 {
			v65 = base.B2i32(base.Ui32(v57) <= base.Ui32(int32(21)))
		} else {
			v65 = int32(0)
		}
		if v65 != 0 {
			return int32(0)
		} else {
			if v55 == int32(127) {
				return int32(0)
			} else {
				v68 = int32(1)
				v69 = v51 + v68
				*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v69)
				return v68
			}
		}
	case 1:
		v32 = l0 + int32(1)
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
		v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		if v37 == int32(244) {
			v40 = int32(143)
		} else {
			v40 = int32(191)
		}
		if v37 == int32(237) {
			v43 = int32(159)
		} else {
			v43 = v40
		}
		if base.Ui32(v43) <= base.Ui32(v33) {
			v51 = v37
			v55 = v51 & int32(255)
			v57 = v55 - int32(223)
			if int32(1)<<(uint(v57)%32)&int32(_a_F_pg_utf8_increment_0) != 0 {
				v65 = base.B2i32(base.Ui32(v57) <= base.Ui32(int32(21)))
			} else {
				v65 = int32(0)
			}
			if v65 != 0 {
				return int32(0)
			} else {
				if v55 == int32(127) {
					return int32(0)
				} else {
					v68 = int32(1)
					v69 = v51 + v68
					*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v69)
					return v68
				}
			}
		} else {
			v45 = int32(1)
			v46 = v33 + v45
			*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v46)
			return v45
		}
	case 2:
		v20 = l0 + int32(2)
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
		if base.Ui32(int32(190)) < base.Ui32(v21) {
			v32 = l0 + int32(1)
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v37 == int32(244) {
				v40 = int32(143)
			} else {
				v40 = int32(191)
			}
			if v37 == int32(237) {
				v43 = int32(159)
			} else {
				v43 = v40
			}
			if base.Ui32(v43) <= base.Ui32(v33) {
				v51 = v37
				v55 = v51 & int32(255)
				v57 = v55 - int32(223)
				if int32(1)<<(uint(v57)%32)&int32(_a_F_pg_utf8_increment_0) != 0 {
					v65 = base.B2i32(base.Ui32(v57) <= base.Ui32(int32(21)))
				} else {
					v65 = int32(0)
				}
				if v65 != 0 {
					return int32(0)
				} else {
					if v55 == int32(127) {
						return int32(0)
					} else {
						v68 = int32(1)
						v69 = v51 + v68
						*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v69)
						return v68
					}
				}
			} else {
				v45 = int32(1)
				v46 = v33 + v45
				*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v46)
				return v45
			}
		} else {
			v24 = int32(1)
			v25 = v21 + v24
			*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v25)
			return v24
		}
	case 3:
		v8 = l0 + int32(3)
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		if base.Ui32(int32(190)) < base.Ui32(v9) {
			v20 = l0 + int32(2)
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
			if base.Ui32(int32(190)) < base.Ui32(v21) {
				v32 = l0 + int32(1)
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				if v37 == int32(244) {
					v40 = int32(143)
				} else {
					v40 = int32(191)
				}
				if v37 == int32(237) {
					v43 = int32(159)
				} else {
					v43 = v40
				}
				if base.Ui32(v43) <= base.Ui32(v33) {
					v51 = v37
					v55 = v51 & int32(255)
					v57 = v55 - int32(223)
					if int32(1)<<(uint(v57)%32)&int32(_a_F_pg_utf8_increment_0) != 0 {
						v65 = base.B2i32(base.Ui32(v57) <= base.Ui32(int32(21)))
					} else {
						v65 = int32(0)
					}
					if v65 != 0 {
						return int32(0)
					} else {
						if v55 == int32(127) {
							return int32(0)
						} else {
							v68 = int32(1)
							v69 = v51 + v68
							*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v69)
							return v68
						}
					}
				} else {
					v45 = int32(1)
					v46 = v33 + v45
					*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v46)
					return v45
				}
			} else {
				v24 = int32(1)
				v25 = v21 + v24
				*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v25)
				return v24
			}
		} else {
			v12 = int32(1)
			v13 = v9 + v12
			*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v13)
			return v12
		}
	default:
		return int32(0)
	}
}
func F_pg_utf8_islegal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	v3 = int32(0)
	switch l1 - int32(1) {
	case 0:
		v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		v39 = v38
		if base.I32_extend8_s(v39) < int32(-62) {
			v52 = v3
		} else {
			v44 = v39
			v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
		}
	case 1:
		v12 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		switch v13 - int32(224) {
		case 0:
			v16 = int32(224)
			if base.Ui32(v16) <= base.Ui32((v12-int32(-64))&int32(255)) {
				v44 = v16
				v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
			} else {
				v52 = v3
			}
		default:
			if v12 <= int32(-65) {
				v39 = v13
				if base.I32_extend8_s(v39) < int32(-62) {
					v52 = v3
				} else {
					v44 = v39
					v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
				}
			} else {
				v52 = v3
			}
		case 13:
			if int32(-97) < v12 {
				v52 = v3
			} else {
				v44 = int32(237)
				v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
			}
		case 16:
			if base.Ui32((v12-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
				v52 = v3
			} else {
				v44 = int32(240)
				v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
			}
		case 20:
			if int32(-113) < v12 {
				v52 = v3
			} else {
				v44 = int32(244)
				v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
			}
		}
	case 2:
		v9 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
		if int32(-65) < v9 {
			v52 = v3
		} else {
			v12 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			switch v13 - int32(224) {
			case 0:
				v16 = int32(224)
				if base.Ui32(v16) <= base.Ui32((v12-int32(-64))&int32(255)) {
					v44 = v16
					v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
				} else {
					v52 = v3
				}
			default:
				if v12 <= int32(-65) {
					v39 = v13
					if base.I32_extend8_s(v39) < int32(-62) {
						v52 = v3
					} else {
						v44 = v39
						v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
					}
				} else {
					v52 = v3
				}
			case 13:
				if int32(-97) < v12 {
					v52 = v3
				} else {
					v44 = int32(237)
					v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
				}
			case 16:
				if base.Ui32((v12-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
					v52 = v3
				} else {
					v44 = int32(240)
					v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
				}
			case 20:
				if int32(-113) < v12 {
					v52 = v3
				} else {
					v44 = int32(244)
					v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
				}
			}
		}
	case 3:
		v6 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+3)))
		if int32(-65) < v6 {
			v52 = v3
		} else {
			v9 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
			if int32(-65) < v9 {
				v52 = v3
			} else {
				v12 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
				v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				switch v13 - int32(224) {
				case 0:
					v16 = int32(224)
					if base.Ui32(v16) <= base.Ui32((v12-int32(-64))&int32(255)) {
						v44 = v16
						v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
					} else {
						v52 = v3
					}
				default:
					if v12 <= int32(-65) {
						v39 = v13
						if base.I32_extend8_s(v39) < int32(-62) {
							v52 = v3
						} else {
							v44 = v39
							v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
						}
					} else {
						v52 = v3
					}
				case 13:
					if int32(-97) < v12 {
						v52 = v3
					} else {
						v44 = int32(237)
						v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
					}
				case 16:
					if base.Ui32((v12-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
						v52 = v3
					} else {
						v44 = int32(240)
						v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
					}
				case 20:
					if int32(-113) < v12 {
						v52 = v3
					} else {
						v44 = int32(244)
						v52 = base.B2i32(base.Ui32(v44&int32(255)) < base.Ui32(int32(245)))
					}
				}
			}
		}
	default:
		v52 = v3
	}
	return v52
}
func F_pg_utf8_verifychar(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v6 = base.I32_extend8_s(v5)
	if int32(0) <= v6 {
		if v6 != 0 {
			v11 = int32(1)
		} else {
			v11 = int32(-1)
		}
		return v11
	} else {
		if v5&int32(224) == int32(192) {
			v31 = int32(2)
		} else {
			if v5&int32(240) == int32(224) {
				v31 = int32(3)
			} else {
				if v5&int32(248) == int32(240) {
					v30 = int32(4)
				} else {
					v30 = int32(1)
				}
				v31 = v30
			}
		}
		if v31 <= l1 {
			v34 = int32(0)
			switch v31 - int32(1) {
			case 0:
				v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				v70 = v69
				if base.I32_extend8_s(v70) < int32(-62) {
					v83 = v34
				} else {
					v75 = v70
					v83 = base.B2i32(base.Ui32(v75&int32(255)) < base.Ui32(int32(245)))
				}
			case 1:
				v43 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				switch v44 - int32(224) {
				case 0:
					v47 = int32(224)
					if base.Ui32(v47) <= base.Ui32((v43-int32(-64))&int32(255)) {
						v75 = v47
						v83 = base.B2i32(base.Ui32(v75&int32(255)) < base.Ui32(int32(245)))
					} else {
						v83 = v34
					}
				default:
					if v43 <= int32(-65) {
						v70 = v44
						if base.I32_extend8_s(v70) < int32(-62) {
							v83 = v34
						} else {
							v75 = v70
							v83 = base.B2i32(base.Ui32(v75&int32(255)) < base.Ui32(int32(245)))
						}
					} else {
						v83 = v34
					}
				case 13:
					if int32(-97) < v43 {
						v83 = v34
					} else {
						v75 = int32(237)
						v83 = base.B2i32(base.Ui32(v75&int32(255)) < base.Ui32(int32(245)))
					}
				case 16:
					if base.Ui32((v43-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
						v83 = v34
					} else {
						v75 = int32(240)
						v83 = base.B2i32(base.Ui32(v75&int32(255)) < base.Ui32(int32(245)))
					}
				case 20:
					if int32(-113) < v43 {
						v83 = v34
					} else {
						v75 = int32(244)
						v83 = base.B2i32(base.Ui32(v75&int32(255)) < base.Ui32(int32(245)))
					}
				}
			case 2:
				v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
				if int32(-65) < v40 {
					v83 = v34
				} else {
					v43 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					switch v44 - int32(224) {
					case 0:
						v47 = int32(224)
						if base.Ui32(v47) <= base.Ui32((v43-int32(-64))&int32(255)) {
							v75 = v47
							v83 = base.B2i32(base.Ui32(v75&int32(255)) < base.Ui32(int32(245)))
						} else {
							v83 = v34
						}
					default:
						if v43 <= int32(-65) {
							v70 = v44
							if base.I32_extend8_s(v70) < int32(-62) {
								v83 = v34
							} else {
								v75 = v70
								v83 = base.B2i32(base.Ui32(v75&int32(255)) < base.Ui32(int32(245)))
							}
						} else {
							v83 = v34
						}
					case 13:
						if int32(-97) < v43 {
							v83 = v34
						} else {
							v75 = int32(237)
							v83 = base.B2i32(base.Ui32(v75&int32(255)) < base.Ui32(int32(245)))
						}
					case 16:
						if base.Ui32((v43-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
							v83 = v34
						} else {
							v75 = int32(240)
							v83 = base.B2i32(base.Ui32(v75&int32(255)) < base.Ui32(int32(245)))
						}
					case 20:
						if int32(-113) < v43 {
							v83 = v34
						} else {
							v75 = int32(244)
							v83 = base.B2i32(base.Ui32(v75&int32(255)) < base.Ui32(int32(245)))
						}
					}
				}
			case 3:
				v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+3)))
				if int32(-65) < v37 {
					v83 = v34
				} else {
					v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+2)))
					if int32(-65) < v40 {
						v83 = v34
					} else {
						v43 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
						switch v44 - int32(224) {
						case 0:
							v47 = int32(224)
							if base.Ui32(v47) <= base.Ui32((v43-int32(-64))&int32(255)) {
								v75 = v47
								v83 = base.B2i32(base.Ui32(v75&int32(255)) < base.Ui32(int32(245)))
							} else {
								v83 = v34
							}
						default:
							if v43 <= int32(-65) {
								v70 = v44
								if base.I32_extend8_s(v70) < int32(-62) {
									v83 = v34
								} else {
									v75 = v70
									v83 = base.B2i32(base.Ui32(v75&int32(255)) < base.Ui32(int32(245)))
								}
							} else {
								v83 = v34
							}
						case 13:
							if int32(-97) < v43 {
								v83 = v34
							} else {
								v75 = int32(237)
								v83 = base.B2i32(base.Ui32(v75&int32(255)) < base.Ui32(int32(245)))
							}
						case 16:
							if base.Ui32((v43-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
								v83 = v34
							} else {
								v75 = int32(240)
								v83 = base.B2i32(base.Ui32(v75&int32(255)) < base.Ui32(int32(245)))
							}
						case 20:
							if int32(-113) < v43 {
								v83 = v34
							} else {
								v75 = int32(244)
								v83 = base.B2i32(base.Ui32(v75&int32(255)) < base.Ui32(int32(245)))
							}
						}
					}
				}
			default:
				v83 = v34
			}
			if v83 != 0 {
				v84 = v31
			} else {
				v84 = int32(-1)
			}
			v85 = v84
		} else {
			v85 = int32(-1)
		}
		return v85
	}
}
func F_pg_xact_status(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int64
	_ = v38
	var v46 int64
	_ = v46
	var v51 int64
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_pg_xact_status[0]))
	v19 = F_LWLockAcquire(m, v15+int32(_a_F_pg_xact_status_0), int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = F_ReadNextFullTransactionId(m)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = base.I32_wrap_i64(v13)
			if v25 == int32(0) {
				v93 = *(*int32)(unsafe.Add(mBase, _c_F_pg_xact_status[0]))
				F_LWLockRelease(m, v93+int32(_a_F_pg_xact_status_0))
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int32(0)
				} else {
					v98 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v98)
					v103 = int32(0)
					m.G0 = v10 + int32(16)
					return v103
				}
			} else {
				if base.Ui32(int32(3)) <= base.Ui32(v25) {
					if base.Ui64(v23) <= base.Ui64(v13) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v10))) = v13
								F_errmsg(m, int32(_a_F_pg_xact_status_1), v10)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_xact_status_2), int32(121), int32(_a_F_pg_xact_status_3))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, _c_F_pg_xact_status[1]))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+64))
						if base.Ui32(v33) <= base.Ui32(int32(2)) {
							v51 = base.I64_extend_i32_u(v33)
						} else {
							v38 = int64(base.Ui64(v23) >> (uint(int64(32)) % 64))
							if base.Ui32(base.I32_wrap_i64(v23)) < base.Ui32(v33) {
								v46 = (v38 - int64(1)) & int64(4294967295)
							} else {
								v46 = v38
							}
							v51 = base.I64_extend_i32_u(v33) | v46<<(uint(int64(32))%64)
						}
						if base.Ui64(v13) < base.Ui64(v51) {
							v93 = *(*int32)(unsafe.Add(mBase, _c_F_pg_xact_status[0]))
							F_LWLockRelease(m, v93+int32(_a_F_pg_xact_status_0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								v98 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v98)
								v103 = int32(0)
								m.G0 = v10 + int32(16)
								return v103
							}
						} else {
							v56 = F_TransactionIdIsInProgress(m, v25)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								if v56 == int32(0) {
									v62 = F_TransactionIdDidCommit(m, v25)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										if v62 != 0 {
											v64 = int32(_a_F_pg_xact_status_4)
										} else {
											v64 = int32(_a_F_pg_xact_status_5)
										}
										v65 = v64
										v67 = *(*int32)(unsafe.Add(mBase, _c_F_pg_xact_status[0]))
										F_LWLockRelease(m, v67+int32(_a_F_pg_xact_status_0))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											v72 = F_cstring_to_text(m, v65)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												v103 = v72
												m.G0 = v10 + int32(16)
												return v103
											}
										}
									}
								} else {
									v65 = int32(_a_F_pg_xact_status_6)
									v67 = *(*int32)(unsafe.Add(mBase, _c_F_pg_xact_status[0]))
									F_LWLockRelease(m, v67+int32(_a_F_pg_xact_status_0))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										v72 = F_cstring_to_text(m, v65)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											v103 = v72
											m.G0 = v10 + int32(16)
											return v103
										}
									}
								}
							}
						}
					}
				} else {
					v56 = F_TransactionIdIsInProgress(m, v25)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						if v56 == int32(0) {
							v62 = F_TransactionIdDidCommit(m, v25)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								if v62 != 0 {
									v64 = int32(_a_F_pg_xact_status_4)
								} else {
									v64 = int32(_a_F_pg_xact_status_5)
								}
								v65 = v64
								v67 = *(*int32)(unsafe.Add(mBase, _c_F_pg_xact_status[0]))
								F_LWLockRelease(m, v67+int32(_a_F_pg_xact_status_0))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									v72 = F_cstring_to_text(m, v65)
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v103 = v72
										m.G0 = v10 + int32(16)
										return v103
									}
								}
							}
						} else {
							v65 = int32(_a_F_pg_xact_status_6)
							v67 = *(*int32)(unsafe.Add(mBase, _c_F_pg_xact_status[0]))
							F_LWLockRelease(m, v67+int32(_a_F_pg_xact_status_0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v72 = F_cstring_to_text(m, v65)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v103 = v72
									m.G0 = v10 + int32(16)
									return v103
								}
							}
						}
					}
				}
			}
		}
	}
}
