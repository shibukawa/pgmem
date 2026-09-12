package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RI_FKey_noaction_upd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	F_ri_CheckTrigger(m, l0, int32(441051), int32(2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_restrict(m, v8, int32(1))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_RI_FKey_setnull_upd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	F_ri_CheckTrigger(m, l0, int32(441072), int32(2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_set(m, v8, int32(1), int32(2))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_ri_GenerateQualCollation(m *base.Module, l0 int32, l1 int32) {
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	v8 = m.G0
	v10 = v8 - int32(192)
	m.G0 = v10
	if l1 != 0 {
		v13 = F_SearchSysCache1(m, int32(16), l1)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			if v13 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
					F_errmsg_internal(m, int32(49449), v10)
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return
					} else {
						F_errfinish(m, int32(515506), int32(2108), int32(274766))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
				v19 = v17 + v18
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
				v21 = F_get_namespace_name(m, v20)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v23 = int32(34)
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)) = uint8(v23)
					v28 = v21
					v30 = v10 + int32(48)
					for {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
						if v34 != int32(34) {
						} else {
							v41 = int32(34)
							*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)) = uint8(v41)
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
							v46 = v43
							v47 = v30 + int32(2)
							*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v46)
							v28 = v28 + int32(1)
							v30 = v47
							continue
						}
						if v34 == int32(0) {
							break
						} else {
							v46 = v34
							v47 = v30 + int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v46)
							v28 = v28 + int32(1)
							v30 = v47
							continue
						}
						break
					}
					v51 = int32(34)
					*(*uint16)(unsafe.Add(mBase, uint32(v30)+1)) = uint16(v51)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v10 + int32(48)
					F_appendStringInfo(m, l0, int32(208146), v10+int32(32))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						v61 = int32(34)
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)) = uint8(v61)
						v68 = v19 + int32(4)
						v70 = v10 + int32(48)
						for {
							v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
							if v74 != int32(34) {
							} else {
								v81 = int32(34)
								*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)) = uint8(v81)
								v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
								v86 = v83
								v87 = v70 + int32(2)
								*(*uint8)(unsafe.Add(mBase, uint32(v87))) = uint8(v86)
								v68 = v68 + int32(1)
								v70 = v87
								continue
							}
							if v74 == int32(0) {
								break
							} else {
								v86 = v74
								v87 = v70 + int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v87))) = uint8(v86)
								v68 = v68 + int32(1)
								v70 = v87
								continue
							}
							break
						}
						v91 = int32(34)
						*(*uint16)(unsafe.Add(mBase, uint32(v70)+1)) = uint16(v91)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v10 + int32(48)
						F_appendStringInfo(m, l0, int32(187127), v10+int32(16))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return
						} else {
							F_ReleaseCatCache(m, v13)
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return
							} else {
								m.G0 = v10 + int32(192)
								return
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v10 + int32(192)
		return
	}
}
func F_ri_PerformCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v131 int32
	_ = v131
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
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
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v248 int32
	_ = v248
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v326 int64
	_ = v326
	var v329 int32
	_ = v329
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	v11 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(368)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if l6 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v25 < int32(3) {
		goto L51
	} else {
		goto L52
	}
L2:
	;
	if int32(0) < v28 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	if v28 <= int32(0) {
		goto L1
	} else {
		goto L37
	}
L5:
	;
	if v25 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v107 = v28
	goto L7
L7:
	;
	if l5 == int32(0) {
		goto L1
	} else {
		goto L22
	}
L8:
	;
	v35 = int32(236)
	goto L10
L9:
	;
	v35 = int32(172)
	goto L10
L10:
	;
	v47 = v11
	v49 = v28
	goto L11
L11:
	;
	v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0+v35+v47<<(uint(int32(1))%32)))))
	v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(l6)+6)))
	if v61 < v60 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v107 = v68
	goto L7
L13:
	;
	F_slot_getsomeattrs_int(m, l6, v60)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v68 = v49
	goto L15
L15:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l6)+16))
	v71 = v60 - int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69+v71<<(uint(int32(2))%32))))
	v76 = int32(32)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81+v71))))
	if v83 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	return int32(0)
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v68 = v67
	goto L15
L18:
	;
	v84 = int32(110)
	goto L20
L19:
	;
	v84 = v76
	goto L20
L20:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23+v76+v47))) = uint8(v84)
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(96)+v47<<(uint(int32(2))%32)))) = v75
	v93 = v47 + int32(1)
	if v93 < v68 {
		v47 = v93
		v49 = v68
		goto L11
	} else {
		goto L21
	}
L21:
	;
	goto L12
L22:
	;
	if v107 <= int32(0) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v25 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v131 = int32(236)
	goto L26
L25:
	;
	v131 = int32(172)
	goto L26
L26:
	;
	v144 = int32(0)
	v146 = v107
	goto L27
L27:
	;
	v157 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0+v131+v144<<(uint(int32(1))%32)))))
	v158 = int32(*(*int16)(unsafe.Add(mBase, uint32(l5)+6)))
	if v158 < v157 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L1
L29:
	;
	F_slot_getsomeattrs_int(m, l5, v157)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L16
	} else {
		goto L32
	}
L30:
	;
	v163 = v146
	goto L31
L31:
	;
	v165 = v157 - int32(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+v166))))
	v169 = int32(2)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v172+v165<<(uint(v169)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(96)+v107<<(uint(int32(2))%32)+v144<<(uint(v169)%32)))) = v176
	if v168 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v163 = v162
	goto L31
L33:
	;
	v181 = int32(110)
	goto L35
L34:
	;
	v181 = int32(32)
	goto L35
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v144+(v23+int32(32)+v107)))) = uint8(v181)
	v184 = v144 + int32(1)
	if v184 < v163 {
		v144 = v184
		v146 = v163
		goto L27
	} else {
		goto L36
	}
L36:
	;
	goto L28
L37:
	;
	if v25 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v192 = int32(236)
	goto L40
L39:
	;
	v192 = int32(172)
	goto L40
L40:
	;
	v204 = v11
	v206 = v28
	goto L41
L41:
	;
	v217 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0+v192+v204<<(uint(int32(1))%32)))))
	v218 = int32(*(*int16)(unsafe.Add(mBase, uint32(l5)+6)))
	if v218 < v217 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L1
L43:
	;
	F_slot_getsomeattrs_int(m, l5, v217)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L16
	} else {
		goto L46
	}
L44:
	;
	v223 = v206
	goto L45
L45:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v226 = v217 - int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v224+v226<<(uint(int32(2))%32))))
	v231 = int32(32)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236+v226))))
	if v238 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v223 = v222
	goto L45
L47:
	;
	v239 = int32(110)
	goto L49
L48:
	;
	v239 = v231
	goto L49
L49:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23+v231+v204))) = uint8(v239)
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(96)+v204<<(uint(int32(2))%32)))) = v230
	v248 = v204 + int32(1)
	if v248 < v223 {
		v204 = v248
		v206 = v223
		goto L41
	} else {
		goto L50
	}
L50:
	;
	goto L42
L51:
	;
	v270 = l4
	goto L53
L52:
	;
	v270 = l3
	goto L53
L53:
	;
	v271 = int32(0)
	if l8 == v271 {
		v286 = v271
		v287 = v271
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(364)))) = v293
	v296 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(360)))) = v296
	goto L60
L55:
	;
	v277 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	if v277 < int32(2) {
		v286 = v271
		v287 = int32(0)
		goto L54
	} else {
		goto L56
	}
L56:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L16
	} else {
		goto L57
	}
L57:
	;
	v282 = F_GetLatestSnapshot(m)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L16
	} else {
		goto L58
	}
L58:
	;
	v284 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L16
	} else {
		goto L59
	}
L59:
	;
	v286 = v282
	v287 = v284
	goto L54
L60:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v270)+48))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+80))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v23)+360))
	*(*int32)(unsafe.Add(mBase, _consts[4])) = v300 | int32(5)
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v299
	goto L61
L61:
	;
	v314 = F_SPI_execute_snapshot(m, l2, v23+int32(96), v23+int32(32), v286, v287, int32(0), base.B2i32(l9 == int32(5)))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L16
	} else {
		goto L62
	}
L62:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v23)+364))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v23)+360))
	*(*int32)(unsafe.Add(mBase, _consts[4])) = v317
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v316
	goto L63
L63:
	;
	if int32(0) <= v314 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	if l6 != 0 {
		goto L83
	} else {
		goto L84
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L16
	} else {
		goto L78
	}
L66:
	;
	if l9 != v314 {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L16
	} else {
		goto L74
	}
L69:
	;
	v326 = *(*int64)(unsafe.Add(mBase, _consts[502]))
	if l9 != int32(5) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	m.G0 = v23 + int32(368)
	return base.B2i32(v326 != int64(0))
L71:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v329 == int32(2) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	if base.B2i32(v326 == int64(0)) != base.B2i32(v329 != int32(1)) {
		goto L64
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	v348 = F_SPI_result_code_string(m, v314)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L16
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v348
	F_errmsg_internal(m, int32(206843), v23)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L16
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(515506), int32(2594), int32(332793))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L16
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L16
	} else {
		goto L79
	}
L79:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l4)+48))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = l0 + int32(20)
	v371 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v367 + v371
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v366 + v371
	F_errmsg(m, int32(104300), v23+int32(16))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L16
	} else {
		goto L80
	}
L80:
	;
	F_errhint(m, int32(599212), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L16
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(515506), int32(2603), int32(332793))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L16
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	v391 = l6
	goto L85
L84:
	;
	v391 = l5
	goto L85
L85:
	;
	v392 = int32(0)
	F_ri_ReportViolation(m, l0, l4, l3, v391, v392, v329, l7, v392)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L16
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ri_PlanCheck(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v19 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(-52)))) = v19
	v22 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(-56)))) = v22
	if v13 < int32(3) {
		v26 = l5
	} else {
		v26 = l4
	}
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, _consts[4])) = v29 | int32(5)
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v28
	v36 = F_SPI_prepare(m, l0, l1, l2)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return int32(0)
	} else {
		if v36 != 0 {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			*(*int32)(unsafe.Add(mBase, _consts[4])) = v41
			*(*int32)(unsafe.Add(mBase, _consts[3])) = v40
			F_SPI_keepplan(m, v36)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, _consts[1045]))
				if v49 != 0 {
					v90 = v49
					v94 = F_hash_search(m, v90, l3, int32(1), v9+int32(-48))
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v36
						m.G0 = v11 - int32(-64)
						return v36
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = int64(3023656976388)
					v58 = F_hash_create(m, int32(416783), int32(64), v9+int32(-48), int32(40))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[1046])) = v58
						F_CacheRegisterSyscacheCallback(m, int32(19), int32(1501), int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = int64(51539607560)
							v74 = F_hash_create(m, int32(416684), int32(256), v9+int32(-48), int32(40))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[1045])) = v74
								*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = int64(292057776136)
								v85 = F_hash_create(m, int32(417097), int32(256), v9+int32(-48), int32(40))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[1047])) = v85
									v89 = *(*int32)(unsafe.Add(mBase, _consts[1045]))
									v90 = v89
									v94 = F_hash_search(m, v90, l3, int32(1), v9+int32(-48))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v36
										m.G0 = v11 - int32(-64)
										return v36
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
			v104 = m.ExcPending
			if v104 != 0 {
				return int32(0)
			} else {
				v106 = *(*int32)(unsafe.Add(mBase, _consts[504]))
				v107 = F_SPI_result_code_string(m, v106)
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v107
					F_errmsg_internal(m, int32(191332), v11)
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(515506), int32(2468), int32(332780))
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
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
func F_ri_ReportViolation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v255 int32
	_ = v255
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	v9 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(224)
	m.G0 = v18
	if l5 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if l7 != 0 {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+52))
	v34 = v33
	v36 = v31
	v37 = v32
	goto L1
L3:
	;
	v23 = l0 + int32(236)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	if l4 == int32(0) {
		v30 = l2
		v31 = v24
		v32 = v23
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v28 = l0 + int32(172)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if l4 != 0 {
		v34 = l4
		v36 = v29
		v37 = v28
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v34 = l4
	v36 = v24
	v37 = v23
	goto L1
L7:
	;
	v30 = l1
	v31 = v29
	v32 = v28
	goto L2
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L13
	} else {
		goto L97
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L13
	} else {
		goto L60
	}
L10:
	;
	if l7 != 0 {
		goto L8
	} else {
		goto L59
	}
L11:
	;
	F_initStringInfo(m, v18+int32(208))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L13
	} else {
		goto L26
	}
L12:
	;
	v40 = F_check_enable_rls(m, v36, int32(0), int32(1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	if v40 == int32(2) {
		v270 = v9
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v47 = F_pg_class_aclcheck(m, v36, v45, int64(2))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v47 == int32(0) {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v51 = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v52 <= v51 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v64 = v51
	goto L19
L19:
	;
	v73 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37+v64<<(uint(int32(1))%32)))))
	v75 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v77 = F_pg_attribute_aclcheck(m, v36, v73, v75, int64(2))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L13
	} else {
		goto L21
	}
L20:
	;
	v255 = v9
	goto L10
L21:
	;
	if v77 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v82 = v64 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v82 < v83 {
		v64 = v82
		goto L19
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L20
L25:
	;
	goto L11
L26:
	;
	F_initStringInfo(m, v18+int32(192))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v109 <= int32(0) {
		v255 = int32(1)
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v113 = v34 + int32(20)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v118 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37))))
	v120 = v118 - int32(1)
	v123 = v113 + v114<<(uint(int32(4))%32) + v120*int32(100)
	v124 = int32(*(*int16)(unsafe.Add(mBase, uint32(l3)+6)))
	if v124 < v118 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_slot_getsomeattrs_int(m, l3, v118)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L13
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v120))))
	if v133 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v120<<(uint(int32(2))%32))))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v123)+68))
	F_getTypeOutputInfo(m, v141, v18+int32(188), v18+int32(187))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L13
	} else {
		goto L36
	}
L34:
	;
	v151 = int32(316837)
	goto L35
L35:
	;
	F_appendStringInfoString(m, v18+int32(208), v123+int32(4))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L13
	} else {
		goto L38
	}
L36:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v18)+188))
	v149 = F_OidOutputFunctionCall(m, v148, v140)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	v151 = v149
	goto L35
L38:
	;
	F_appendStringInfoString(m, v18+int32(192), v151)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L13
	} else {
		goto L39
	}
L39:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v161 < int32(2) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v255 = int32(1)
	goto L10
L41:
	;
	goto L42
L42:
	;
	v174 = int32(1)
	goto L43
L43:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v184 = int32(1)
	v187 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37+v174<<(uint(v184)%32)))))
	v189 = v187 - v184
	v192 = v113 + v180<<(uint(int32(4))%32) + v189*int32(100)
	v193 = int32(*(*int16)(unsafe.Add(mBase, uint32(l3)+6)))
	if v193 < v187 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v255 = v240
	goto L10
L45:
	;
	F_slot_getsomeattrs_int(m, l3, v187)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L13
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+v189))))
	if v200 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v203+v189<<(uint(int32(2))%32))))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v192)+68))
	F_getTypeOutputInfo(m, v208, v18+int32(188), v18+int32(187))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L13
	} else {
		goto L52
	}
L50:
	;
	v219 = int32(316837)
	goto L51
L51:
	;
	F_appendStringInfoString(m, v18+int32(208), int32(778193))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L13
	} else {
		goto L54
	}
L52:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v18)+188))
	v216 = F_OidOutputFunctionCall(m, v215, v207)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L13
	} else {
		goto L53
	}
L53:
	;
	v219 = v216
	goto L51
L54:
	;
	F_appendStringInfoString(m, v18+int32(192), int32(778193))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	F_appendStringInfoString(m, v18+int32(208), v192+int32(4))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	F_appendStringInfoString(m, v18+int32(192), v219)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	v240 = int32(1)
	v242 = v174 + v240
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v242 < v243 {
		v174 = v242
		goto L43
	} else {
		goto L58
	}
L58:
	;
	goto L44
L59:
	;
	v270 = v255
	goto L9
L60:
	;
	if l5 == int32(1) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	F_errcode(m, int32(50352322))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L13
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v327 = l0 + int32(20)
	if l6 != 0 {
		goto L74
	} else {
		goto L75
	}
L64:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v286 = l0 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v286
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v284 + int32(4)
	F_errmsg(m, int32(730205), v18-int32(-64))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	if v270 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	F_errtableconstraint(m, l2, v286)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L13
	} else {
		goto L72
	}
L67:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v18)+208))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v297
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v18)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v299
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v296 + int32(4)
	F_errdetail(m, int32(695521), v18+int32(32))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L13
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v309 + int32(4)
	F_errdetail(m, int32(695487), v18+int32(48))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L13
	} else {
		goto L71
	}
L70:
	;
	goto L66
L71:
	;
	goto L66
L72:
	;
	F_errfinish(m, int32(515506), int32(2783), int32(274327))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L13
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errcode(m, int32(16777410))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L13
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	F_errcode(m, int32(50352322))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L13
	} else {
		goto L87
	}
L77:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = v327
	v334 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = v332 + v334
	*(*int32)(unsafe.Add(mBase, uint32(v18)+112)) = v331 + v334
	F_errmsg(m, int32(750617), v18+int32(112))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L13
	} else {
		goto L78
	}
L78:
	;
	if v270 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	F_errtableconstraint(m, l2, v327)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L13
	} else {
		goto L85
	}
L80:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v18)+208))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v346
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v18)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v345 + int32(4)
	F_errdetail(m, int32(695600), v18+int32(80))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L13
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v358 + int32(4)
	F_errdetail(m, int32(695565), v18+int32(96))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L13
	} else {
		goto L84
	}
L83:
	;
	goto L79
L84:
	;
	goto L79
L85:
	;
	F_errfinish(m, int32(515506), int32(2797), int32(274327))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L13
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
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+164)) = v327
	v381 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+168)) = v379 + v381
	*(*int32)(unsafe.Add(mBase, uint32(v18)+160)) = v378 + v381
	F_errmsg(m, int32(750535), v18+int32(160))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L13
	} else {
		goto L88
	}
L88:
	;
	if v270 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	F_errtableconstraint(m, l2, v327)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L13
	} else {
		goto L95
	}
L90:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v18)+208))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+128)) = v393
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v18)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+132)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v18)+136)) = v392 + int32(4)
	F_errdetail(m, int32(695686), v18+int32(128))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L13
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+144)) = v405 + int32(4)
	F_errdetail(m, int32(695645), v18+int32(144))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L13
	} else {
		goto L94
	}
L93:
	;
	goto L89
L94:
	;
	goto L89
L95:
	;
	F_errfinish(m, int32(515506), int32(2811), int32(274327))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L13
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
	F_errcode(m, int32(50352322))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L13
	} else {
		goto L98
	}
L98:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v431 = l0 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v429 + int32(4)
	F_errmsg(m, int32(730144), v18+int32(16))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L13
	} else {
		goto L99
	}
L99:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v18)+208))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v442
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v18)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v441 + int32(4)
	F_errdetail(m, int32(695686), v18)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L13
	} else {
		goto L100
	}
L100:
	;
	F_errtableconstraint(m, l2, v431)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L13
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(515506), int32(2770), int32(274327))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L13
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
