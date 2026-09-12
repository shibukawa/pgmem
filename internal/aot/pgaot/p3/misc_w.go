package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WaitLatch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v5 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[840]))
	v13 = int32(1)
	F_ModifyWaitEvent(m, v11, v5, v13, (v5-l1&v13)&l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = int32(*(*uint8)(unsafe.Add(mBase, _consts[104])))
		if v24 == int32(1) {
			v28 = *(*int32)(unsafe.Add(mBase, _consts[840]))
			F_ModifyWaitEvent(m, v28, int32(1), l1&int32(48), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, _consts[840]))
				if l1&int32(8) != 0 {
					v40 = l2
				} else {
					v40 = int32(-1)
				}
				v42 = F_WaitEventSetWait(m, v36, v40, v8, int32(1), l3)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					m.G0 = v8 + int32(16)
					if v42 != 0 {
						v49 = v44
					} else {
						v49 = int32(8)
					}
					return v49
				}
			}
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, _consts[840]))
			if l1&int32(8) != 0 {
				v40 = l2
			} else {
				v40 = int32(-1)
			}
			v42 = F_WaitEventSetWait(m, v36, v40, v8, int32(1), l3)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				m.G0 = v8 + int32(16)
				if v42 != 0 {
					v49 = v44
				} else {
					v49 = int32(8)
				}
				return v49
			}
		}
	}
}
func F_WritebackContextInit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	return
}
func F_win1250_to_latin2(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(29), int32(9))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_local2local(m, v6, v5, v10, int32(29), int32(9), int32(2271568), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
func F_win1251_to_iso(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(23), int32(25))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_local2local(m, v6, v5, v10, int32(23), int32(25), int32(2267040), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
func F_win1251_to_koi8r(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(23), int32(22))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_local2local(m, v6, v5, v10, int32(23), int32(22), int32(2266272), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
func F_win1251_to_mic(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(23), int32(7))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_latin2mic_with_table(m, v6, v5, v10, int32(139), int32(23), int32(2266272), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
func F_write_relcache_init_file(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	v9 = m.G0
	v11 = v9 - int32(2160)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[1095]))
	if v14 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L9
	} else {
		goto L98
	}
L2:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L9
	} else {
		goto L94
	}
L3:
	;
	m.G0 = v11 + int32(2160)
	return
L4:
	;
	if l0 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v68 = v11 + int32(1136)
	v69 = F_unlink(m, v68)
	mBase = m.M
	v73 = F_AllocateFile(m, v68, int32(33848))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L9
	} else {
		goto L14
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(106736)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v18
	v26 = F_pg_snprintf(m, v11+int32(1136), int32(1024), int32(486844), v11+int32(32))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = int32(106736)
	v41 = *(*int32)(unsafe.Add(mBase, _consts[1096]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v41
	v44 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v44
	v52 = F_pg_snprintf(m, v11+int32(1136), int32(1024), int32(486835), v11-int32(-64))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L12
	}
L9:
	;
	return
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(106736)
	v36 = F_pg_snprintf(m, v11+int32(112), int32(1024), int32(186954), v11+int32(16))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	goto L5
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = int32(106736)
	v57 = *(*int32)(unsafe.Add(mBase, _consts[1096]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v57
	v65 = F_pg_snprintf(m, v11+int32(112), int32(1024), int32(186923), v11+int32(48))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L5
L14:
	;
	if v73 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v79 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L9
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = int32(5714534)
	v106 = F_fwrite(m, v11+int32(108), int32(1), int32(4), v73)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L9
	} else {
		goto L24
	}
L18:
	;
	if v79 == int32(0) {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(1136)
	F_errmsg(m, int32(311445), v11)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	F_errdetail(m, int32(652513), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(522071), int32(6635), int32(403855))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	goto L3
L24:
	;
	if v106 != int32(4) {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _consts[1091]))
	F_hash_seq_init(m, v11+int32(88), v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v118 = F_hash_seq_search(m, v11+int32(88))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	if v118 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v122 = v118
	goto L31
L29:
	;
	goto L30
L30:
	;
	v327 = F_FreeFile(m, v73)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L9
	} else {
		goto L84
	}
L31:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+48))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129)+117)))
	if v130 != l0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L30
L33:
	;
	v317 = F_hash_seq_search(m, v11+int32(88))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L9
	} else {
		goto L82
	}
L34:
	;
	if l0 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_write_item(m, v128, int32(276), v73)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L9
	} else {
		goto L54
	}
L36:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+56))
	if base.Ui32(v132-int32(3592)) < base.Ui32(int32(2)) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	if v132 == int32(2671) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	if v132 == int32(2701) {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v141 = int32(0)
	v147 = *(*int32)(unsafe.Add(mBase, _consts[1097]))
	v149 = v147 - int32(1)
	if v149 < v141 {
		v180 = v141
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v180 == int32(0) {
		goto L33
	} else {
		goto L53
	}
L41:
	;
	goto L40
L42:
	;
	v153 = v141
	v154 = v149
	goto L43
L43:
	;
	v159 = int32(2)
	v160 = base.I32_div_s(v154-v153, v159)
	v161 = v160 + v153
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161<<(uint(v159)%32))+uint32(_consts[1098])))
	v167 = base.B2i32(v166 == v132)
	if v166 == v132 {
		v180 = v167
		goto L41
	} else {
		goto L45
	}
L44:
	;
	v180 = v167
	goto L41
L45:
	;
	v170 = base.B2i32(base.Ui32(v166) < base.Ui32(v132))
	if base.Ui32(v166) < base.Ui32(v132) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v171 = v161 + int32(1)
	goto L48
L47:
	;
	v171 = v153
	goto L48
L48:
	;
	if base.Ui32(v166) < base.Ui32(v132) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v174 = v154
	goto L51
L50:
	;
	v174 = v161 - int32(1)
	goto L51
L51:
	;
	if v171 <= v174 {
		v153 = v171
		v154 = v174
		goto L43
	} else {
		goto L52
	}
L52:
	;
	goto L44
L53:
	;
	goto L35
L54:
	;
	F_write_item(m, v129, int32(144), v73)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	v191 = int32(*(*int16)(unsafe.Add(mBase, uint32(v129)+120)))
	if int32(0) < v191 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v197 = int32(0)
	goto L59
L57:
	;
	goto L58
L58:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v128)+180))
	if v228 != 0 {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v128)+52))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v208 = int32(100)
	F_write_item(m, v203+v204<<(uint(int32(4))%32)+v197*v208+int32(20), v208, v73)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L9
	} else {
		goto L61
	}
L60:
	;
	goto L58
L61:
	;
	v217 = v197 + int32(1)
	v218 = int32(*(*int16)(unsafe.Add(mBase, uint32(v129)+120)))
	if v217 < v218 {
		v197 = v217
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v233 = int32(base.Ui32(v229) >> (uint(int32(2)) % 32))
	goto L65
L64:
	;
	v233 = int32(0)
	goto L65
L65:
	;
	F_write_item(m, v228, v233, v73)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v128)+48))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+119)))
	if v237 != int32(105) {
		goto L33
	} else {
		goto L67
	}
L67:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v128)+196))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	F_write_item(m, v240, v241+int32(24), v73)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v128)+208))
	v247 = int32(*(*int16)(unsafe.Add(mBase, uint32(v129)+120)))
	F_write_item(m, v246, v247<<(uint(int32(2))%32), v73)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v128)+212))
	v253 = int32(*(*int16)(unsafe.Add(mBase, uint32(v129)+120)))
	F_write_item(m, v252, v253<<(uint(int32(2))%32), v73)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v128)+216))
	v259 = int32(*(*int16)(unsafe.Add(mBase, uint32(v129)+120)))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v128)+204))
	v261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v260)+6)))
	F_write_item(m, v258, v259*v261<<(uint(int32(2))%32), v73)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v128)+248))
	v268 = int32(*(*int16)(unsafe.Add(mBase, uint32(v129)+120)))
	F_write_item(m, v267, v268<<(uint(int32(2))%32), v73)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v128)+224))
	v274 = int32(*(*int16)(unsafe.Add(mBase, uint32(v129)+120)))
	F_write_item(m, v273, v274<<(uint(int32(1))%32), v73)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	v279 = int32(*(*int16)(unsafe.Add(mBase, uint32(v129)+120)))
	if v279 <= int32(0) {
		goto L33
	} else {
		goto L74
	}
L74:
	;
	v285 = int32(0)
	goto L75
L75:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v128)+252))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v292+v285<<(uint(int32(2))%32))))
	if v296 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L33
L77:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	v300 = int32(base.Ui32(v297) >> (uint(int32(2)) % 32))
	goto L79
L78:
	;
	v300 = int32(0)
	goto L79
L79:
	;
	F_write_item(m, v296, v300, v73)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	v304 = v285 + int32(1)
	v305 = int32(*(*int16)(unsafe.Add(mBase, uint32(v129)+120)))
	if v304 < v305 {
		v285 = v304
		goto L75
	} else {
		goto L81
	}
L81:
	;
	goto L76
L82:
	;
	if v317 != 0 {
		v122 = v317
		goto L31
	} else {
		goto L83
	}
L83:
	;
	goto L32
L84:
	;
	if v327 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	v334 = F_LWLockAcquire(m, v330+int32(2048), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L9
	} else {
		goto L86
	}
L86:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L9
	} else {
		goto L87
	}
L87:
	;
	v339 = *(*int32)(unsafe.Add(mBase, _consts[1095]))
	if v339 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v353 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	F_LWLockRelease(m, v353+int32(2048))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L9
	} else {
		goto L93
	}
L89:
	;
	v346 = F_rename(m, v11+int32(1136), v11+int32(112))
	mBase = m.M
	if int32(0) <= v346 {
		goto L88
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v351 = F_unlink(m, v11+int32(1136))
	mBase = m.M
	goto L88
L92:
	;
	goto L91
L93:
	;
	goto L3
L94:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L9
	} else {
		goto L95
	}
L95:
	;
	F_errmsg_internal(m, int32(307804), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L9
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(522071), int32(6647), int32(403855))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L9
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L9
	} else {
		goto L99
	}
L99:
	;
	F_errmsg_internal(m, int32(307804), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L9
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(522071), int32(6750), int32(403855))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L9
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_writetup_cluster(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v10 + int32(10)
	F_LogicalTapeWrite(m, l1, v7+int32(12), int32(4))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		F_LogicalTapeWrite(m, l1, v9+int32(4), int32(6))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			F_LogicalTapeWrite(m, l1, v24, v25)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
				if v28&int32(1) != 0 {
					F_LogicalTapeWrite(m, l1, v7+int32(12), int32(4))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			}
		}
	}
}
func F_writetup_index_brin(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v10 + v11
	F_LogicalTapeWrite(m, l1, v7+int32(12), v11)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		F_LogicalTapeWrite(m, l1, v9+int32(4), v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
			if v24&int32(1) != 0 {
				F_LogicalTapeWrite(m, l1, v7+int32(12), int32(4))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	}
}
