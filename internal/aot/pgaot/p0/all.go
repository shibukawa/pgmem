package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AllTablesyncsReady(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v1 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)) = uint8(v1)
	v12 = F_FetchTableStates(m, v6+int32(15))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
		if v16 == int32(1) {
			F_CommitTransactionCommand(m)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v22 = F_pgstat_report_stat(m, int32(1))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _consts[533]))
					m.G0 = v6 + int32(16)
					return v12 & base.B2i32(v25 == int32(0))
				}
			}
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, _consts[533]))
			m.G0 = v6 + int32(16)
			return v12 & base.B2i32(v25 == int32(0))
		}
	}
}
func F_ExecGetAllUpdatedCols(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+152))
	if v5 == int32(0) {
		v8 = F_MakePerTupleExprContext(m, l1)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = v8
			v13 = int32(4562080)
			v14 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
			v18 = F_ExecGetUpdatedCols(m, l0, l1)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v20 == int32(0) {
					F_ExecInitGenerated(m, l0, l1, int32(2))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v27 = F_bms_union(m, v18, v26)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v14
							return v27
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v27 = F_bms_union(m, v18, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v14
						return v27
					}
				}
			}
		}
	} else {
		v12 = v5
		v13 = int32(4562080)
		v14 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
		v18 = F_ExecGetUpdatedCols(m, l0, l1)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v20 == int32(0) {
				F_ExecInitGenerated(m, l0, l1, int32(2))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v27 = F_bms_union(m, v18, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v14
						return v27
					}
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v27 = F_bms_union(m, v18, v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v14
					return v27
				}
			}
		}
	}
}
func F_ExecStoreAllNullTuple(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	m.T0[v5].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v11 = v9 << (uint(int32(2)) % 32)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v12&int32(3) != 0 {
			v35 = F__emscripten_memset_bulkmem(m, v12, base.I32_extend8_s(int32(0)), v11)
			mBase = m.M
		} else {
			if base.Ui32(int32(1024)) < base.Ui32(v11) {
				v35 = F__emscripten_memset_bulkmem(m, v12, base.I32_extend8_s(int32(0)), v11)
				mBase = m.M
			} else {
				v17 = v12 + v11
				if base.Ui32(v17) <= base.Ui32(v12) {
				} else {
					v23 = v12 + int32(4)
					if base.Ui32(v23) < base.Ui32(v17) {
						v25 = v17
					} else {
						v25 = v23
					}
					v32 = F__emscripten_memset_bulkmem(m, v12, base.I32_extend8_s(int32(0)), (v12^int32(-1)+v25)&int32(-4)+int32(4))
					mBase = m.M
				}
			}
		}
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
		v43 = F__emscripten_memset_bulkmem(m, v38, base.I32_extend8_s(int32(1)), v41)
		mBase = m.M
		v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
		v46 = v44 & int32(65533)
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v46)
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
		*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v49)
		return
	}
}
func F_all_rows_selectable(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v463 int32
	_ = v463
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(l1) < base.Ui32(v9) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11+l1<<(uint(int32(2))%32))))
	v17 = v15
	goto L3
L2:
	;
	v17 = int32(0)
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v17 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v30 = v18 + l1<<(uint(int32(2))%32)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v30 = v24 + l1<<(uint(int32(2))%32) - int32(4)
	goto L4
L8:
	;
	v42 = v17 + int32(160)
	goto L10
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	v36 = F_getRTEPermissionInfo(m, v35, v31)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v43 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	return int32(0)
L12:
	;
	v42 = v36 + int32(24)
	goto L10
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v48 = v47
	goto L15
L14:
	;
	v48 = v43
	goto L15
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v49 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	return int32(0)
L17:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v312)+128))
	if v317 != 0 {
		goto L16
	} else {
		goto L82
	}
L18:
	;
	v312 = v31
	v313 = l2
	goto L17
L19:
	;
	goto L20
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v49+l1<<(uint(int32(2))%32))))
	if v56 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v293 != 0 {
		goto L79
	} else {
		goto L80
	}
L22:
	;
	v289 = l1
	v292 = l2
	v293 = v52
	goto L21
L23:
	;
	goto L24
L24:
	;
	v60 = l1
	v61 = l2
	v62 = v56
	v64 = v52
	goto L25
L25:
	;
	if v64 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v289 = v283
	v292 = v277
	v293 = v281
	goto L21
L27:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	if v82 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v80 = v64 + v67<<(uint(int32(2))%32)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+52))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v80 = v73 + v74<<(uint(int32(2))%32) - int32(4)
	goto L27
L31:
	;
	v289 = v60
	v292 = v61
	v293 = v64
	goto L21
L32:
	;
	goto L33
L33:
	;
	v83 = int32(0)
	if v61 == v83 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	if int32(0) <= v140 {
		goto L45
	} else {
		goto L46
	}
L35:
	;
	v140 = base.I32_ctz(v126) | v127<<(uint(int32(5))%32)
	goto L34
L36:
	;
	v140 = int32(-2)
	goto L34
L37:
	;
	v93 = base.I32_div_s(int32(0), int32(32))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v94 <= v93 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v97 = v61 + int32(8)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97+v93<<(uint(int32(2))%32))))
	v104 = v101 & int32(-1)
	if v104 != 0 {
		v126 = v104
		v127 = v93
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v106 = v93 + int32(1)
	if v106 == v94 {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v109 = v106
	goto L41
L41:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v97+v109<<(uint(int32(2))%32))))
	if v116 != 0 {
		v126 = v116
		v127 = v109
		goto L35
	} else {
		goto L43
	}
L42:
	;
	goto L36
L43:
	;
	v118 = v109 + int32(1)
	if v118 != v94 {
		v109 = v118
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v147 = v83
	v149 = v140
	goto L48
L46:
	;
	v277 = v83
	goto L47
L47:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v282+v283<<(uint(int32(2))%32))))
	if v287 != 0 {
		v60 = v283
		v61 = v277
		v62 = v287
		v64 = v281
		goto L25
	} else {
		goto L78
	}
L48:
	;
	v152 = v149 - int32(7)
	if v152&int32(65535) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v277 = v211
	goto L47
L50:
	;
	if v61 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L51:
	;
	v157 = int32(1)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v62)+24))
	if v159 <= int32(0) {
		v211 = v147
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v188 = base.I32_extend16_s(v152)
	if int32(0) <= v188 {
		goto L60
	} else {
		goto L61
	}
L54:
	;
	v163 = v157
	v166 = v147
	v167 = v157
	goto L55
L55:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v62)+28))
	v176 = int32(*(*int16)(unsafe.Add(mBase, uint32(v170+v163<<(uint(int32(1))%32)-int32(2)))))
	if v176 == int32(0) {
		goto L16
	} else {
		goto L57
	}
L56:
	;
	v211 = v181
	goto L50
L57:
	;
	v181 = F_bms_add_member(m, v166, v176+int32(7))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L11
	} else {
		goto L58
	}
L58:
	;
	v184 = v167 + int32(1)
	v185 = base.I32_extend16_s(v184)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v62)+24))
	if v185 <= v186 {
		v163 = v185
		v166 = v181
		v167 = v184
		goto L55
	} else {
		goto L59
	}
L59:
	;
	goto L56
L60:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v62)+24))
	if v191 < v188 {
		goto L16
	} else {
		goto L63
	}
L61:
	;
	v202 = v188
	goto L62
L62:
	;
	v205 = F_bms_add_member(m, v147, v202+int32(7))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L11
	} else {
		goto L65
	}
L63:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v62)+28))
	v199 = int32(*(*int16)(unsafe.Add(mBase, uint32(v193+v188<<(uint(int32(1))%32)-int32(2)))))
	if v199 == int32(0) {
		goto L16
	} else {
		goto L64
	}
L64:
	;
	v202 = v199
	goto L62
L65:
	;
	v211 = v205
	goto L50
L66:
	;
	if int32(0) <= v270 {
		v147 = v211
		v149 = v270
		goto L48
	} else {
		goto L77
	}
L67:
	;
	v270 = base.I32_ctz(v256) | v257<<(uint(int32(5))%32)
	goto L66
L68:
	;
	v270 = int32(-2)
	goto L66
L69:
	;
	v221 = v149 + int32(1)
	v223 = base.I32_div_s(v221, int32(32))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v224 <= v223 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v227 = v61 + int32(8)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227+v223<<(uint(int32(2))%32))))
	v234 = v231 & (int32(-1) << (uint(v221) % 32))
	if v234 != 0 {
		v256 = v234
		v257 = v223
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v236 = v223 + int32(1)
	if v236 == v224 {
		goto L68
	} else {
		goto L72
	}
L72:
	;
	v239 = v236
	goto L73
L73:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v227+v239<<(uint(int32(2))%32))))
	if v246 != 0 {
		v256 = v246
		v257 = v239
		goto L67
	} else {
		goto L75
	}
L74:
	;
	goto L68
L75:
	;
	v248 = v239 + int32(1)
	if v248 != v224 {
		v239 = v248
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	goto L49
L78:
	;
	goto L26
L79:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v293+v289<<(uint(int32(2))%32))))
	v312 = v299
	v313 = v292
	goto L17
L80:
	;
	goto L81
L81:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)+52))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)+12))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v302+v289<<(uint(int32(2))%32)-int32(4))))
	v312 = v308
	v313 = v292
	goto L17
L82:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v312)+16))
	v320 = F_pg_class_aclcheck(m, v318, v48, int64(2))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L11
	} else {
		goto L84
	}
L83:
	;
	return int32(1)
L84:
	;
	if v320 == int32(0) {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	if v313 == int32(0) {
		goto L16
	} else {
		goto L86
	}
L86:
	;
	if v313 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	if v382 < int32(0) {
		goto L83
	} else {
		goto L98
	}
L88:
	;
	v382 = base.I32_ctz(v368) | v369<<(uint(int32(5))%32)
	goto L87
L89:
	;
	v382 = int32(-2)
	goto L87
L90:
	;
	v335 = base.I32_div_s(int32(0), int32(32))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	if v336 <= v335 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v339 = v313 + int32(8)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v339+v335<<(uint(int32(2))%32))))
	v346 = v343 & int32(-1)
	if v346 != 0 {
		v368 = v346
		v369 = v335
		goto L88
	} else {
		goto L92
	}
L92:
	;
	v348 = v335 + int32(1)
	if v348 == v336 {
		goto L89
	} else {
		goto L93
	}
L93:
	;
	v351 = v348
	goto L94
L94:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v339+v351<<(uint(int32(2))%32))))
	if v358 != 0 {
		v368 = v358
		v369 = v351
		goto L88
	} else {
		goto L96
	}
L95:
	;
	goto L89
L96:
	;
	v360 = v351 + int32(1)
	if v360 != v336 {
		v351 = v360
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v386 = v382
	goto L99
L99:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v312)+16))
	v395 = v386 - int32(7)
	if v395&int32(65535) == int32(0) {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L83
L101:
	;
	if v313 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L102:
	;
	v402 = F_pg_attribute_aclcheck_all(m, v393, v48, int64(2), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L11
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v406 = F_pg_attribute_aclcheck(m, v393, base.I32_extend16_s(v395), v48, int64(2))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L11
	} else {
		goto L107
	}
L105:
	;
	if v402 != 0 {
		goto L16
	} else {
		goto L106
	}
L106:
	;
	goto L101
L107:
	;
	if v406 != 0 {
		goto L16
	} else {
		goto L108
	}
L108:
	;
	goto L101
L109:
	;
	if int32(0) <= v463 {
		v386 = v463
		goto L99
	} else {
		goto L120
	}
L110:
	;
	v463 = base.I32_ctz(v449) | v450<<(uint(int32(5))%32)
	goto L109
L111:
	;
	v463 = int32(-2)
	goto L109
L112:
	;
	v414 = v386 + int32(1)
	v416 = base.I32_div_s(v414, int32(32))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	if v417 <= v416 {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v420 = v313 + int32(8)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v420+v416<<(uint(int32(2))%32))))
	v427 = v424 & (int32(-1) << (uint(v414) % 32))
	if v427 != 0 {
		v449 = v427
		v450 = v416
		goto L110
	} else {
		goto L114
	}
L114:
	;
	v429 = v416 + int32(1)
	if v429 == v417 {
		goto L111
	} else {
		goto L115
	}
L115:
	;
	v432 = v429
	goto L116
L116:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v420+v432<<(uint(int32(2))%32))))
	if v439 != 0 {
		v449 = v439
		v450 = v432
		goto L110
	} else {
		goto L118
	}
L117:
	;
	goto L111
L118:
	;
	v441 = v432 + int32(1)
	if v441 != v417 {
		v432 = v441
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	goto L100
}
