package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pq_getkeepalivesinterval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 == v2 {
		v21 = v2
	} else {
		v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
		if v10 == int32(1) {
			v21 = v2
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+404))
			if v13 != 0 {
				v21 = v13
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+388))
				if v14 != 0 {
					v21 = v14
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(4)
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(388))))
					v21 = v20
				}
			}
		}
	}
	m.G0 = v6 + int32(16)
	return v21
}
func F_pq_getmsgbytes(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	if int32(0) <= l1 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if l1 <= v6-v7 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1 + v7
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			return v31 + v7
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16908800))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pq_getmsgbytes_0), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pq_getmsgbytes_1), int32(515), int32(_a_F_pq_getmsgbytes_2))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
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
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_pq_getmsgbytes_0), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pq_getmsgbytes_1), int32(515), int32(_a_F_pq_getmsgbytes_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
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
func F_pq_getmsgstring(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = v5 + v6
	v8 = F_strlen(m, v7)
	mBase = m.M
	v9 = v8 + v5
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v10 <= v9 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_pq_getmsgstring_0), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pq_getmsgstring_1), int32(595), int32(_a_F_pq_getmsgstring_2))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9 + int32(1)
		v33 = F_pg_client_to_server(m, v7, v8)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			return v33
		}
	}
}
func F_pq_getmsgtext(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	if l1 < int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_pq_getmsgtext_0), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pq_getmsgtext_1), int32(554), int32(_a_F_pq_getmsgtext_2))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
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
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v7-v8 < l1 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16908800))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pq_getmsgtext_0), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pq_getmsgtext_1), int32(554), int32(_a_F_pq_getmsgtext_2))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
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
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1 + v8
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v14 = v13 + v8
			v15 = F_pg_client_to_server(m, v14, l1)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v15 != v14 {
					v20 = F_strlen(m, v15)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v20
					return v15
				} else {
					v25 = F_palloc(m, l1+int32(1))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if l1 != 0 {
							base.MemoryCopy(m, v25, v14, l1)
						} else {
						}
						v29 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v25+l1))) = uint8(v29)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = l1
						return v25
					}
				}
			}
		}
	}
}
func F_pq_parse_errornotice(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
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
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	base.MemoryFill(m, l1+int32(4), int32(0), int32(96))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(21)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_pq_parse_errornotice[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v18
	v20 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L2
	} else {
		goto L134
	}
L2:
	;
	return
L3:
	;
	v23 = v20 << (uint(int32(24)) % 32)
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = v23
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_pq_getmsgend(m, l0)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L2
	} else {
		goto L133
	}
L7:
	;
	v29 = F_pq_getmsgrawstring(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v32 = v28 >> (uint(int32(24)) % 32)
	switch v32 - int32(67) {
	case 0:
		goto L27
	case 1:
		goto L26
	default:
		goto L12
	case 3:
		goto L15
	case 5:
		goto L25
	case 9:
		goto L14
	case 10:
		goto L11
	case 13:
		goto L24
	case 15:
		goto L13
	case 16:
		goto L10
	case 19:
		goto L28
	case 20:
		goto L21
	case 32:
		goto L18
	case 33:
		goto L17
	case 43:
		goto L16
	case 45:
		goto L23
	case 46:
		goto L22
	case 48:
		goto L20
	case 49:
		goto L19
	}
L10:
	;
	v397 = F_pq_getmsgbyte(m, l0)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L2
	} else {
		goto L131
	}
L11:
	;
	v394 = F_pstrdup(m, v29)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L2
	} else {
		goto L130
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L2
	} else {
		goto L127
	}
L13:
	;
	v378 = F_pstrdup(m, v29)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L2
	} else {
		goto L126
	}
L14:
	;
	v375 = F_pg_strtoint32(m, v29)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L2
	} else {
		goto L125
	}
L15:
	;
	v372 = F_pstrdup(m, v29)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L2
	} else {
		goto L124
	}
L16:
	;
	v369 = F_pstrdup(m, v29)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L2
	} else {
		goto L123
	}
L17:
	;
	v366 = F_pstrdup(m, v29)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L2
	} else {
		goto L122
	}
L18:
	;
	v363 = F_pstrdup(m, v29)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L2
	} else {
		goto L121
	}
L19:
	;
	v360 = F_pstrdup(m, v29)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L2
	} else {
		goto L120
	}
L20:
	;
	v357 = F_pstrdup(m, v29)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L2
	} else {
		goto L119
	}
L21:
	;
	v354 = F_pstrdup(m, v29)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L2
	} else {
		goto L118
	}
L22:
	;
	v351 = F_pstrdup(m, v29)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L2
	} else {
		goto L117
	}
L23:
	;
	v348 = F_pg_strtoint32(m, v29)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L2
	} else {
		goto L116
	}
L24:
	;
	v345 = F_pg_strtoint32(m, v29)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L2
	} else {
		goto L115
	}
L25:
	;
	v342 = F_pstrdup(m, v29)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L2
	} else {
		goto L114
	}
L26:
	;
	v339 = F_pstrdup(m, v29)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L2
	} else {
		goto L113
	}
L27:
	;
	v298 = F_strlen(m, v29)
	mBase = m.M
	if v298 != int32(5) {
		goto L1
	} else {
		goto L112
	}
L28:
	;
	v35 = int32(_a_F_pq_parse_errornotice_0)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pq_parse_errornotice[1])))
	if base.B2i32(v38 == int32(0))|base.B2i32(v38 != v41) != 0 {
		v59 = v38
		v60 = v41
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v59-v60 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	v44 = v29
	v45 = v35
	goto L32
L32:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if v49 == int32(0) {
		v59 = v49
		v60 = v48
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v59 = v49
	v60 = v48
	goto L30
L34:
	;
	v52 = int32(1)
	if v49 == v48 {
		v44 = v44 + v52
		v45 = v45 + v52
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(14)
	goto L10
L37:
	;
	goto L38
L38:
	;
	v66 = int32(_a_F_pq_parse_errornotice_1)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pq_parse_errornotice[2])))
	if base.B2i32(v69 == int32(0))|base.B2i32(v69 != v72) != 0 {
		v90 = v69
		v91 = v72
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v90-v91 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	goto L39
L41:
	;
	v75 = v29
	v76 = v66
	goto L42
L42:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	if v80 == int32(0) {
		v90 = v80
		v91 = v79
		goto L40
	} else {
		goto L44
	}
L43:
	;
	v90 = v80
	v91 = v79
	goto L40
L44:
	;
	v83 = int32(1)
	if v80 == v79 {
		v75 = v75 + v83
		v76 = v76 + v83
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(15)
	goto L10
L47:
	;
	goto L48
L48:
	;
	v97 = int32(_a_F_pq_parse_errornotice_2)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pq_parse_errornotice[3])))
	if base.B2i32(v100 == int32(0))|base.B2i32(v100 != v103) != 0 {
		v121 = v100
		v122 = v103
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v121-v122 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	goto L49
L51:
	;
	v106 = v29
	v107 = v97
	goto L52
L52:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v111 == int32(0) {
		v121 = v111
		v122 = v110
		goto L50
	} else {
		goto L54
	}
L53:
	;
	v121 = v111
	v122 = v110
	goto L50
L54:
	;
	v114 = int32(1)
	if v111 == v110 {
		v106 = v106 + v114
		v107 = v107 + v114
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(17)
	goto L10
L57:
	;
	goto L58
L58:
	;
	v128 = int32(_a_F_pq_parse_errornotice_3)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pq_parse_errornotice[4])))
	if base.B2i32(v131 == int32(0))|base.B2i32(v131 != v134) != 0 {
		v152 = v131
		v153 = v134
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v152-v153 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L60:
	;
	goto L59
L61:
	;
	v137 = v29
	v138 = v128
	goto L62
L62:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+1)))
	if v142 == int32(0) {
		v152 = v142
		v153 = v141
		goto L60
	} else {
		goto L64
	}
L63:
	;
	v152 = v142
	v153 = v141
	goto L60
L64:
	;
	v145 = int32(1)
	if v142 == v141 {
		v137 = v137 + v145
		v138 = v138 + v145
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(18)
	goto L10
L67:
	;
	goto L68
L68:
	;
	v159 = int32(_a_F_pq_parse_errornotice_4)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pq_parse_errornotice[5])))
	if base.B2i32(v162 == int32(0))|base.B2i32(v162 != v165) != 0 {
		v183 = v162
		v184 = v165
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v183-v184 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L70:
	;
	goto L69
L71:
	;
	v168 = v29
	v169 = v159
	goto L72
L72:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	if v173 == int32(0) {
		v183 = v173
		v184 = v172
		goto L70
	} else {
		goto L74
	}
L73:
	;
	v183 = v173
	v184 = v172
	goto L70
L74:
	;
	v176 = int32(1)
	if v173 == v172 {
		v168 = v168 + v176
		v169 = v169 + v176
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(19)
	goto L10
L77:
	;
	goto L78
L78:
	;
	v190 = int32(_a_F_pq_parse_errornotice_5)
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pq_parse_errornotice[6])))
	if base.B2i32(v193 == int32(0))|base.B2i32(v193 != v196) != 0 {
		v214 = v193
		v215 = v196
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v214-v215 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L80:
	;
	goto L79
L81:
	;
	v199 = v29
	v200 = v190
	goto L82
L82:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200)+1)))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+1)))
	if v204 == int32(0) {
		v214 = v204
		v215 = v203
		goto L80
	} else {
		goto L84
	}
L83:
	;
	v214 = v204
	v215 = v203
	goto L80
L84:
	;
	v207 = int32(1)
	if v204 == v203 {
		v199 = v199 + v207
		v200 = v200 + v207
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(21)
	goto L10
L87:
	;
	goto L88
L88:
	;
	v221 = int32(_a_F_pq_parse_errornotice_6)
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pq_parse_errornotice[7])))
	if base.B2i32(v224 == int32(0))|base.B2i32(v224 != v227) != 0 {
		v245 = v224
		v246 = v227
		goto L90
	} else {
		goto L91
	}
L89:
	;
	if v245-v246 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L90:
	;
	goto L89
L91:
	;
	v230 = v29
	v231 = v221
	goto L92
L92:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+1)))
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+1)))
	if v235 == int32(0) {
		v245 = v235
		v246 = v234
		goto L90
	} else {
		goto L94
	}
L93:
	;
	v245 = v235
	v246 = v234
	goto L90
L94:
	;
	v238 = int32(1)
	if v235 == v234 {
		v230 = v230 + v238
		v231 = v231 + v238
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(22)
	goto L10
L97:
	;
	goto L98
L98:
	;
	v252 = int32(_a_F_pq_parse_errornotice_7)
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pq_parse_errornotice[8])))
	if base.B2i32(v255 == int32(0))|base.B2i32(v255 != v258) != 0 {
		v276 = v255
		v277 = v258
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v276-v277 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L100:
	;
	goto L99
L101:
	;
	v261 = v29
	v262 = v252
	goto L102
L102:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+1)))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)))
	if v266 == int32(0) {
		v276 = v266
		v277 = v265
		goto L100
	} else {
		goto L104
	}
L103:
	;
	v276 = v266
	v277 = v265
	goto L100
L104:
	;
	v269 = int32(1)
	if v266 == v265 {
		v261 = v261 + v269
		v262 = v262 + v269
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(23)
	goto L10
L107:
	;
	goto L108
L108:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L2
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v29
	F_errmsg_internal(m, int32(_a_F_pq_parse_errornotice_8), v8+int32(16))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L2
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_pq_parse_errornotice_9), int32(272), int32(_a_F_pq_parse_errornotice_10))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L2
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	v302 = int32(16)
	v304 = int32(63)
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)))
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+3)))
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = (v301+v302)&v304 | (v306+v302)&v304<<(uint(int32(6))%32) | (v314+v302)&v304<<(uint(int32(12))%32) | (v322+v302)&v304<<(uint(int32(18))%32) | (v330+v302)&v304<<(uint(int32(24))%32)
	goto L10
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v339
	goto L10
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v342
	goto L10
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v345
	goto L10
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+84)) = v348
	goto L10
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v351
	goto L10
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v354
	goto L10
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+60)) = v357
	goto L10
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v360
	goto L10
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v363
	goto L10
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v366
	goto L10
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v369
	goto L10
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v372
	goto L10
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v375
	goto L10
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v378
	goto L10
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v32
	F_errmsg_internal(m, int32(_a_F_pq_parse_errornotice_11), v8)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L2
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_pq_parse_errornotice_9), int32(326), int32(_a_F_pq_parse_errornotice_10))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L2
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
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v394
	goto L10
L131:
	;
	v400 = v397 << (uint(int32(24)) % 32)
	if v400 != 0 {
		v28 = v400
		goto L7
	} else {
		goto L132
	}
L132:
	;
	goto L8
L133:
	;
	m.G0 = v8 + int32(48)
	return
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v29
	F_errmsg_internal(m, int32(_a_F_pq_parse_errornotice_12), v8+int32(32))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L2
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_pq_parse_errornotice_9), int32(276), int32(_a_F_pq_parse_errornotice_10))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L2
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pq_redirect_to_shm_mq(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	*(*int32)(unsafe.Add(mBase, _c_F_pq_redirect_to_shm_mq[0])) = l1
	*(*int32)(unsafe.Add(mBase, _c_F_pq_redirect_to_shm_mq[1])) = int32(_a_F_pq_redirect_to_shm_mq_0)
	*(*int32)(unsafe.Add(mBase, _c_F_pq_redirect_to_shm_mq[2])) = int32(2)
	*(*int32)(unsafe.Add(mBase, _c_F_pq_redirect_to_shm_mq[3])) = int32(_a_F_pq_redirect_to_shm_mq_1)
	F_on_dsm_detach(m, l0, int32(808), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		return
	}
}
func F_pq_send_ascii_string(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = l1
	v8 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_appendStringInfoChar(m, l0, int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L13
	} else {
		goto L16
	}
L4:
	;
	if base.I32_extend8_s(v8) < int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v14 = int32(63)
	goto L8
L7:
	;
	v14 = v8
	goto L8
L8:
	;
	v15 = base.I32_extend8_s(v14)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 <= v17+int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v36 = v7 + int32(1)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v37 != 0 {
		v7 = v36
		v8 = v37
		goto L4
	} else {
		goto L15
	}
L10:
	;
	F_appendStringInfoChar(m, l0, v15)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23+v17))) = uint8(v15)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = v26 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30+v28))) = uint8(v32)
	goto L9
L13:
	;
	return
L14:
	;
	goto L9
L15:
	;
	goto L5
L16:
	;
	return
}
