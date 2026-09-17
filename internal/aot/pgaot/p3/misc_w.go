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
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLatch[0]))
	v13 = int32(1)
	F_ModifyWaitEvent(m, v11, v5, v13, (v5-l1&v13)&l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitLatch[1])))
		if v24 == int32(1) {
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLatch[0]))
			F_ModifyWaitEvent(m, v28, int32(1), l1&int32(48), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLatch[0]))
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
			v36 = *(*int32)(unsafe.Add(mBase, _c_F_WaitLatch[0]))
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
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13936(m, l0, int32(_a_F_win1250_to_latin2_0), int32(9), int32(29))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_win1251_to_iso(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13936(m, l0, int32(_a_F_win1251_to_iso_0), int32(25), int32(23))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_win1251_to_koi8r(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13936(m, l0, int32(_a_F_win1251_to_koi8r_0), int32(22), int32(23))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_win1251_to_mic(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13929(m, l0, int32(_a_F_win1251_to_mic_0), int32(23), int32(139))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
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
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
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
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	v9 = m.G0
	v11 = v9 - int32(2160)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_write_relcache_init_file[0]))
	if v14 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L9
	} else {
		goto L96
	}
L2:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L9
	} else {
		goto L92
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
	v71 = F_AllocateFile(m, v68, int32(_a_F_write_relcache_init_file_0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L9
	} else {
		goto L14
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(_a_F_write_relcache_init_file_1)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_write_relcache_init_file[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v18
	v26 = F_pg_snprintf(m, v11+int32(1136), int32(1024), int32(_a_F_write_relcache_init_file_2), v11+int32(32))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = int32(_a_F_write_relcache_init_file_1)
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_write_relcache_init_file[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v41
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_write_relcache_init_file[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v44
	v52 = F_pg_snprintf(m, v11+int32(1136), int32(1024), int32(_a_F_write_relcache_init_file_3), v11-int32(-64))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_write_relcache_init_file_1)
	v36 = F_pg_snprintf(m, v11+int32(112), int32(1024), int32(_a_F_write_relcache_init_file_4), v11+int32(16))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = int32(_a_F_write_relcache_init_file_1)
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_write_relcache_init_file[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v57
	v65 = F_pg_snprintf(m, v11+int32(112), int32(1024), int32(_a_F_write_relcache_init_file_5), v11+int32(48))
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
	if v71 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v77 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L9
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = int32(_a_F_write_relcache_init_file_6)
	v102 = F_fwrite(m, v11+int32(108), int32(1), int32(4), v71)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L24
	}
L18:
	;
	if v77 == int32(0) {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v68
	F_errmsg(m, int32(_a_F_write_relcache_init_file_7), v11)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L21
	}
L21:
	;
	F_errdetail(m, int32(_a_F_write_relcache_init_file_8), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_write_relcache_init_file_9), int32(_a_F_write_relcache_init_file_10), int32(_a_F_write_relcache_init_file_11))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	goto L3
L24:
	;
	if v102 != int32(4) {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v107 = v11 + int32(88)
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_write_relcache_init_file[3]))
	F_hash_seq_init(m, v107, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	v112 = F_hash_seq_search(m, v107)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	if v112 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v116 = v112
	goto L31
L29:
	;
	goto L30
L30:
	;
	v323 = F_FreeFile(m, v71)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L9
	} else {
		goto L82
	}
L31:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+48))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+117)))
	if v124 != l0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L30
L33:
	;
	v313 = F_hash_seq_search(m, v11+int32(88))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L9
	} else {
		goto L80
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
	F_write_item(m, v122, int32(276), v71)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L9
	} else {
		goto L52
	}
L36:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)+56))
	if base.B2i32(base.Ui32(v126-int32(3592)) < base.Ui32(int32(2)))|base.B2i32(v126 == int32(2671))|base.B2i32(v126 == int32(2701)) != 0 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v137 = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_write_relcache_init_file[4]))
	v145 = v143 - int32(1)
	if v145 < v137 {
		v177 = v137
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v177 == int32(0) {
		goto L33
	} else {
		goto L51
	}
L39:
	;
	goto L38
L40:
	;
	v149 = v145
	v150 = v137
	goto L41
L41:
	;
	v155 = int32(2)
	v156 = base.I32_div_s(v149-v150, v155)
	v157 = v156 + v150
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v157<<(uint(v155)%32))+uint32(_c_F_write_relcache_init_file[5])))
	v163 = base.B2i32(v162 == v126)
	if v162 == v126 {
		v177 = v163
		goto L39
	} else {
		goto L43
	}
L42:
	;
	v177 = v163
	goto L39
L43:
	;
	v166 = base.B2i32(base.Ui32(v162) < base.Ui32(v126))
	if base.Ui32(v162) < base.Ui32(v126) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v167 = v157 + int32(1)
	goto L46
L45:
	;
	v167 = v150
	goto L46
L46:
	;
	if base.Ui32(v162) < base.Ui32(v126) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v170 = v149
	goto L49
L48:
	;
	v170 = v157 - int32(1)
	goto L49
L49:
	;
	if v167 <= v170 {
		v149 = v170
		v150 = v167
		goto L41
	} else {
		goto L50
	}
L50:
	;
	goto L42
L51:
	;
	goto L35
L52:
	;
	F_write_item(m, v123, int32(144), v71)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	v187 = int32(*(*int16)(unsafe.Add(mBase, uint32(v123)+120)))
	if int32(0) < v187 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v193 = int32(0)
	goto L57
L55:
	;
	goto L56
L56:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v122)+180))
	if v224 != 0 {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v122)+52))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v204 = int32(100)
	F_write_item(m, v199+v200<<(uint(int32(4))%32)+v193*v204+int32(20), v204, v71)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L9
	} else {
		goto L59
	}
L58:
	;
	goto L56
L59:
	;
	v213 = v193 + int32(1)
	v214 = int32(*(*int16)(unsafe.Add(mBase, uint32(v123)+120)))
	if v213 < v214 {
		v193 = v213
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v229 = int32(base.Ui32(v225) >> (uint(int32(2)) % 32))
	goto L63
L62:
	;
	v229 = int32(0)
	goto L63
L63:
	;
	F_write_item(m, v224, v229, v71)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v122)+48))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+119)))
	if v233 != int32(105) {
		goto L33
	} else {
		goto L65
	}
L65:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v122)+196))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	F_write_item(m, v236, v237+int32(24), v71)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v122)+208))
	v243 = int32(*(*int16)(unsafe.Add(mBase, uint32(v123)+120)))
	F_write_item(m, v242, v243<<(uint(int32(2))%32), v71)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v122)+212))
	v249 = int32(*(*int16)(unsafe.Add(mBase, uint32(v123)+120)))
	F_write_item(m, v248, v249<<(uint(int32(2))%32), v71)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v122)+216))
	v255 = int32(*(*int16)(unsafe.Add(mBase, uint32(v123)+120)))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v122)+204))
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256)+6)))
	F_write_item(m, v254, v255*v257<<(uint(int32(2))%32), v71)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v122)+248))
	v264 = int32(*(*int16)(unsafe.Add(mBase, uint32(v123)+120)))
	F_write_item(m, v263, v264<<(uint(int32(2))%32), v71)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v122)+224))
	v270 = int32(*(*int16)(unsafe.Add(mBase, uint32(v123)+120)))
	F_write_item(m, v269, v270<<(uint(int32(1))%32), v71)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	v275 = int32(*(*int16)(unsafe.Add(mBase, uint32(v123)+120)))
	if v275 <= int32(0) {
		goto L33
	} else {
		goto L72
	}
L72:
	;
	v281 = int32(0)
	goto L73
L73:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v122)+252))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v288+v281<<(uint(int32(2))%32))))
	if v292 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L33
L75:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v296 = int32(base.Ui32(v293) >> (uint(int32(2)) % 32))
	goto L77
L76:
	;
	v296 = int32(0)
	goto L77
L77:
	;
	F_write_item(m, v292, v296, v71)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L9
	} else {
		goto L78
	}
L78:
	;
	v300 = v281 + int32(1)
	v301 = int32(*(*int16)(unsafe.Add(mBase, uint32(v123)+120)))
	if v300 < v301 {
		v281 = v300
		goto L73
	} else {
		goto L79
	}
L79:
	;
	goto L74
L80:
	;
	if v313 != 0 {
		v116 = v313
		goto L31
	} else {
		goto L81
	}
L81:
	;
	goto L32
L82:
	;
	if v323 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_write_relcache_init_file[6]))
	v330 = F_LWLockAcquire(m, v326+int32(2048), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L9
	} else {
		goto L84
	}
L84:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L9
	} else {
		goto L85
	}
L85:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_write_relcache_init_file[0]))
	if v335 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v349 = *(*int32)(unsafe.Add(mBase, _c_F_write_relcache_init_file[6]))
	F_LWLockRelease(m, v349+int32(2048))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L9
	} else {
		goto L91
	}
L87:
	;
	v342 = F_rename(m, v11+int32(1136), v11+int32(112))
	mBase = m.M
	if int32(0) <= v342 {
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v347 = F_unlink(m, v11+int32(1136))
	mBase = m.M
	goto L86
L90:
	;
	goto L89
L91:
	;
	goto L3
L92:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L9
	} else {
		goto L93
	}
L93:
	;
	F_errmsg_internal(m, int32(_a_F_write_relcache_init_file_12), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L9
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_write_relcache_init_file_9), int32(_a_F_write_relcache_init_file_13), int32(_a_F_write_relcache_init_file_11))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L9
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L9
	} else {
		goto L97
	}
L97:
	;
	F_errmsg_internal(m, int32(_a_F_write_relcache_init_file_12), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L9
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_write_relcache_init_file_9), int32(_a_F_write_relcache_init_file_14), int32(_a_F_write_relcache_init_file_11))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L9
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_writetup_cluster(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v11 + int32(10)
	v16 = v8 + int32(12)
	F_LogicalTapeWrite(m, l1, v16, int32(4))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		F_LogicalTapeWrite(m, l1, v10+int32(4), int32(6))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			F_LogicalTapeWrite(m, l1, v25, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
				if v29&int32(1) != 0 {
					F_LogicalTapeWrite(m, l1, v16, int32(4))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			}
		}
	}
}
func F_writetup_index_brin(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v11 + v12
	v16 = v8 + int32(12)
	F_LogicalTapeWrite(m, l1, v16, v12)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		F_LogicalTapeWrite(m, l1, v10+int32(4), v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
			if v25&int32(1) != 0 {
				F_LogicalTapeWrite(m, l1, v16, int32(4))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
