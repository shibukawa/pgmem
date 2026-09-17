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
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_noaction_upd_0), int32(2))
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
	F_ri_CheckTrigger(m, l0, int32(_a_F_RI_FKey_setnull_upd_0), int32(2))
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
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
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
				v114 = m.ExcPending
				if v114 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
					F_errmsg_internal(m, int32(_a_F_ri_GenerateQualCollation_0), v10)
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ri_GenerateQualCollation_1), int32(2108), int32(_a_F_ri_GenerateQualCollation_2))
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
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
					v54 = v10 + int32(48)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v54
					F_appendStringInfo(m, l0, int32(_a_F_ri_GenerateQualCollation_3), v10+int32(32))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						v61 = int32(34)
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)) = uint8(v61)
						v66 = v19 + int32(4)
						v68 = v54
						for {
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
							if v72 != int32(34) {
							} else {
								v79 = int32(34)
								*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)) = uint8(v79)
								v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
								v84 = v81
								v85 = v68 + int32(2)
								*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v84)
								v66 = v66 + int32(1)
								v68 = v85
								continue
							}
							if v72 == int32(0) {
								break
							} else {
								v84 = v72
								v85 = v68 + int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v84)
								v66 = v66 + int32(1)
								v68 = v85
								continue
							}
							break
						}
						v89 = int32(34)
						*(*uint16)(unsafe.Add(mBase, uint32(v68)+1)) = uint16(v89)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v10 + int32(48)
						F_appendStringInfo(m, l0, int32(_a_F_ri_GenerateQualCollation_4), v10+int32(16))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							F_ReleaseCatCache(m, v13)
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
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
	var v48 int32
	_ = v48
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
	var v115 int32
	_ = v115
	var v132 int32
	_ = v132
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
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
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v327 int64
	_ = v327
	var v330 int32
	_ = v330
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
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
		goto L50
	} else {
		goto L51
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
		goto L36
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
	v115 = int32(0)
	if base.B2i32(l5 == v115)|base.B2i32(v107 <= v115) != 0 {
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
	v48 = v11
	v49 = v28
	goto L11
L11:
	;
	v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0+v35+v48<<(uint(int32(1))%32)))))
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
	*(*uint8)(unsafe.Add(mBase, uint32(v23+v76+v48))) = uint8(v84)
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(96)+v48<<(uint(int32(2))%32)))) = v75
	v93 = v48 + int32(1)
	if v93 < v68 {
		v48 = v93
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
	if v25 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v132 = int32(236)
	goto L25
L24:
	;
	v132 = int32(172)
	goto L25
L25:
	;
	v146 = int32(0)
	v147 = v107
	goto L26
L26:
	;
	v158 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0+v132+v146<<(uint(int32(1))%32)))))
	v159 = int32(*(*int16)(unsafe.Add(mBase, uint32(l5)+6)))
	if v159 < v158 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L1
L28:
	;
	F_slot_getsomeattrs_int(m, l5, v158)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L16
	} else {
		goto L31
	}
L29:
	;
	v164 = v147
	goto L30
L30:
	;
	v166 = v158 - int32(1)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166+v167))))
	v170 = int32(2)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173+v166<<(uint(v170)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(96)+v107<<(uint(int32(2))%32)+v146<<(uint(v170)%32)))) = v177
	if v169 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v164 = v163
	goto L30
L32:
	;
	v182 = int32(110)
	goto L34
L33:
	;
	v182 = int32(32)
	goto L34
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v146+(v23+int32(32)+v107)))) = uint8(v182)
	v185 = v146 + int32(1)
	if v185 < v164 {
		v146 = v185
		v147 = v164
		goto L26
	} else {
		goto L35
	}
L35:
	;
	goto L27
L36:
	;
	if v25 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v193 = int32(236)
	goto L39
L38:
	;
	v193 = int32(172)
	goto L39
L39:
	;
	v206 = v11
	v207 = v28
	goto L40
L40:
	;
	v218 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0+v193+v206<<(uint(int32(1))%32)))))
	v219 = int32(*(*int16)(unsafe.Add(mBase, uint32(l5)+6)))
	if v219 < v218 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L1
L42:
	;
	F_slot_getsomeattrs_int(m, l5, v218)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L16
	} else {
		goto L45
	}
L43:
	;
	v224 = v207
	goto L44
L44:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v227 = v218 - int32(1)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v225+v227<<(uint(int32(2))%32))))
	v232 = int32(32)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237+v227))))
	if v239 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v224 = v223
	goto L44
L46:
	;
	v240 = int32(110)
	goto L48
L47:
	;
	v240 = v232
	goto L48
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23+v232+v206))) = uint8(v240)
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(96)+v206<<(uint(int32(2))%32)))) = v231
	v249 = v206 + int32(1)
	if v249 < v224 {
		v206 = v249
		v207 = v224
		goto L40
	} else {
		goto L49
	}
L49:
	;
	goto L41
L50:
	;
	v271 = l4
	goto L52
L51:
	;
	v271 = l3
	goto L52
L52:
	;
	v272 = int32(0)
	if l8 == v272 {
		v287 = v272
		v288 = v272
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_ri_PerformCheck[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(364)))) = v294
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_ri_PerformCheck[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(360)))) = v297
	goto L59
L54:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_ri_PerformCheck[2]))
	if v278 < int32(2) {
		v287 = v272
		v288 = int32(0)
		goto L53
	} else {
		goto L55
	}
L55:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L16
	} else {
		goto L56
	}
L56:
	;
	v283 = F_GetLatestSnapshot(m)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L16
	} else {
		goto L57
	}
L57:
	;
	v285 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L16
	} else {
		goto L58
	}
L58:
	;
	v287 = v283
	v288 = v285
	goto L53
L59:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v271)+48))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+80))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v23)+360))
	*(*int32)(unsafe.Add(mBase, _c_F_ri_PerformCheck[1])) = v301 | int32(5)
	*(*int32)(unsafe.Add(mBase, _c_F_ri_PerformCheck[0])) = v300
	goto L60
L60:
	;
	v315 = F_SPI_execute_snapshot(m, l2, v23+int32(96), v23+int32(32), v287, v288, int32(0), base.B2i32(l9 == int32(5)))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L16
	} else {
		goto L61
	}
L61:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v23)+364))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v23)+360))
	*(*int32)(unsafe.Add(mBase, _c_F_ri_PerformCheck[1])) = v318
	*(*int32)(unsafe.Add(mBase, _c_F_ri_PerformCheck[0])) = v317
	goto L62
L62:
	;
	if int32(0) <= v315 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	if l6 != 0 {
		goto L82
	} else {
		goto L83
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L16
	} else {
		goto L77
	}
L65:
	;
	if v315 != l9 {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L16
	} else {
		goto L73
	}
L68:
	;
	v327 = *(*int64)(unsafe.Add(mBase, _c_F_ri_PerformCheck[3]))
	if l9 != int32(5) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	m.G0 = v23 + int32(368)
	return base.B2i32(v327 != int64(0))
L70:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v330 == int32(2) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	if base.B2i32(v327 == int64(0)) != base.B2i32(v330 != int32(1)) {
		goto L63
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	v349 = F_SPI_result_code_string(m, v315)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L16
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v349
	F_errmsg_internal(m, int32(_a_F_ri_PerformCheck_0), v23)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L16
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_ri_PerformCheck_1), int32(2594), int32(_a_F_ri_PerformCheck_2))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L16
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errcode(m, int32(2600))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L16
	} else {
		goto L78
	}
L78:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l4)+48))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = l0 + int32(20)
	v372 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v368 + v372
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v367 + v372
	F_errmsg(m, int32(_a_F_ri_PerformCheck_3), v23+int32(16))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L16
	} else {
		goto L79
	}
L79:
	;
	F_errhint(m, int32(_a_F_ri_PerformCheck_4), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L16
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_ri_PerformCheck_1), int32(2603), int32(_a_F_ri_PerformCheck_2))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L16
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	v392 = l6
	goto L84
L83:
	;
	v392 = l5
	goto L84
L84:
	;
	v393 = int32(0)
	F_ri_ReportViolation(m, l0, l4, l3, v392, v393, v330, l7, v393)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L16
	} else {
		goto L85
	}
L85:
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
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_ri_PlanCheck[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(-52)))) = v19
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_ri_PlanCheck[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(-56)))) = v22
	if v13 < int32(3) {
		v26 = l5
	} else {
		v26 = l4
	}
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_ri_PlanCheck[1])) = v29 | int32(5)
	*(*int32)(unsafe.Add(mBase, _c_F_ri_PlanCheck[0])) = v28
	v36 = F_SPI_prepare(m, l0, l1, l2)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return int32(0)
	} else {
		if v36 != 0 {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			*(*int32)(unsafe.Add(mBase, _c_F_ri_PlanCheck[1])) = v41
			*(*int32)(unsafe.Add(mBase, _c_F_ri_PlanCheck[0])) = v40
			F_SPI_keepplan(m, v36)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, _c_F_ri_PlanCheck[2]))
				if v49 != 0 {
					v87 = v49
					v91 = F_hash_search(m, v87, l3, int32(1), v9+int32(-48))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v36
						m.G0 = v11 - int32(-64)
						return v36
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = int64(3023656976388)
					v56 = v9 + int32(-48)
					v58 = F_hash_create(m, int32(_a_F_ri_PlanCheck_0), int32(64), v56, int32(40))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_ri_PlanCheck[3])) = v58
						F_CacheRegisterSyscacheCallback(m, int32(19), int32(1485), int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = int64(51539607560)
							v72 = F_hash_create(m, int32(_a_F_ri_PlanCheck_1), int32(256), v56, int32(40))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_ri_PlanCheck[2])) = v72
								*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = int64(292057776136)
								v81 = F_hash_create(m, int32(_a_F_ri_PlanCheck_2), int32(256), v56, int32(40))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ri_PlanCheck[4])) = v81
									v85 = *(*int32)(unsafe.Add(mBase, _c_F_ri_PlanCheck[2]))
									v87 = v85
									v91 = F_hash_search(m, v87, l3, int32(1), v9+int32(-48))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v36
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
			v101 = m.ExcPending
			if v101 != 0 {
				return int32(0)
			} else {
				v103 = *(*int32)(unsafe.Add(mBase, _c_F_ri_PlanCheck[5]))
				v104 = F_SPI_result_code_string(m, v103)
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v104
					F_errmsg_internal(m, int32(_a_F_ri_PlanCheck_3), v11)
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_ri_PlanCheck_4), int32(2468), int32(_a_F_ri_PlanCheck_5))
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
		}
	}
}
func F_ri_ReportViolation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
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
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v252 int32
	_ = v252
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
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
	v423 = m.ExcPending
	if v423 != 0 {
		goto L13
	} else {
		goto L97
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
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
	v104 = m.ExcPending
	if v104 != 0 {
		goto L13
	} else {
		goto L26
	}
L12:
	;
	v38 = int32(0)
	v41 = F_check_enable_rls(m, v36, v38, int32(1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	if v41 == int32(2) {
		v267 = v38
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_ri_ReportViolation[0]))
	v48 = F_pg_class_aclcheck(m, v36, v46, int64(2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v48 == int32(0) {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v52 <= int32(0) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v64 = v38
	goto L19
L19:
	;
	v73 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37+v64<<(uint(int32(1))%32)))))
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_ri_ReportViolation[0]))
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
	v252 = int32(0)
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
	v108 = m.ExcPending
	if v108 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v110 <= int32(0) {
		v252 = int32(1)
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v117 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37))))
	v120 = v34 + v113<<(uint(int32(4))%32) + v117*int32(100)
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(l3)+6)))
	if v121 < v117 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	F_slot_getsomeattrs_int(m, l3, v117)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L13
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v127 = v117 - int32(1)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127+v128))))
	if v130 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v133+v127<<(uint(int32(2))%32))))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v120-int32(80))+68))
	F_getTypeOutputInfo(m, v140, v18+int32(188), v18+int32(187))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L13
	} else {
		goto L36
	}
L34:
	;
	v150 = int32(_a_F_ri_ReportViolation_0)
	goto L35
L35:
	;
	F_appendStringInfoString(m, v18+int32(208), v120-int32(76))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L13
	} else {
		goto L38
	}
L36:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v18)+188))
	v148 = F_OidOutputFunctionCall(m, v147, v137)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	v150 = v148
	goto L35
L38:
	;
	F_appendStringInfoString(m, v18+int32(192), v150)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L13
	} else {
		goto L39
	}
L39:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v162 < int32(2) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v252 = int32(1)
	goto L10
L41:
	;
	goto L42
L42:
	;
	v176 = int32(1)
	goto L43
L43:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v188 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37+v176<<(uint(int32(1))%32)))))
	v191 = v34 + v181<<(uint(int32(4))%32) + v188*int32(100)
	v192 = int32(*(*int16)(unsafe.Add(mBase, uint32(l3)+6)))
	if v192 < v188 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v252 = v238
	goto L10
L45:
	;
	F_slot_getsomeattrs_int(m, l3, v188)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L13
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v198 = v188 - int32(1)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+v199))))
	if v201 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204+v198<<(uint(int32(2))%32))))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v191-int32(80))+68))
	F_getTypeOutputInfo(m, v211, v18+int32(188), v18+int32(187))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L13
	} else {
		goto L52
	}
L50:
	;
	v221 = int32(_a_F_ri_ReportViolation_0)
	goto L51
L51:
	;
	v225 = v18 + int32(208)
	F_appendStringInfoString(m, v225, int32(_a_F_ri_ReportViolation_1))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L13
	} else {
		goto L54
	}
L52:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v18)+188))
	v219 = F_OidOutputFunctionCall(m, v218, v208)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L13
	} else {
		goto L53
	}
L53:
	;
	v221 = v219
	goto L51
L54:
	;
	v230 = v18 + int32(192)
	F_appendStringInfoString(m, v230, int32(_a_F_ri_ReportViolation_1))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	F_appendStringInfoString(m, v225, v191-int32(76))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	F_appendStringInfoString(m, v230, v221)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	v238 = int32(1)
	v240 = v176 + v238
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v240 < v241 {
		v176 = v240
		goto L43
	} else {
		goto L58
	}
L58:
	;
	goto L44
L59:
	;
	v267 = v252
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
	v281 = m.ExcPending
	if v281 != 0 {
		goto L13
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v325 = l0 + int32(20)
	if l6 != 0 {
		goto L74
	} else {
		goto L75
	}
L64:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v284 = l0 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v282 + int32(4)
	F_errmsg(m, int32(_a_F_ri_ReportViolation_2), v18-int32(-64))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	if v267 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	F_errtableconstraint(m, l2, v284)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L13
	} else {
		goto L72
	}
L67:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v18)+208))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v295
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v18)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v297
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v294 + int32(4)
	F_errdetail(m, int32(_a_F_ri_ReportViolation_3), v18+int32(32))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L13
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v307 + int32(4)
	F_errdetail(m, int32(_a_F_ri_ReportViolation_4), v18+int32(48))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
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
	F_errfinish(m, int32(_a_F_ri_ReportViolation_5), int32(2783), int32(_a_F_ri_ReportViolation_6))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
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
	v328 = m.ExcPending
	if v328 != 0 {
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
	v375 = m.ExcPending
	if v375 != 0 {
		goto L13
	} else {
		goto L87
	}
L77:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = v325
	v332 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = v330 + v332
	*(*int32)(unsafe.Add(mBase, uint32(v18)+112)) = v329 + v332
	F_errmsg(m, int32(_a_F_ri_ReportViolation_7), v18+int32(112))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L13
	} else {
		goto L78
	}
L78:
	;
	if v267 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	F_errtableconstraint(m, l2, v325)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L13
	} else {
		goto L85
	}
L80:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v18)+208))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v344
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v18)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v343 + int32(4)
	F_errdetail(m, int32(_a_F_ri_ReportViolation_8), v18+int32(80))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L13
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v356 + int32(4)
	F_errdetail(m, int32(_a_F_ri_ReportViolation_9), v18+int32(96))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
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
	F_errfinish(m, int32(_a_F_ri_ReportViolation_5), int32(2797), int32(_a_F_ri_ReportViolation_6))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
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
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+164)) = v325
	v379 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+168)) = v377 + v379
	*(*int32)(unsafe.Add(mBase, uint32(v18)+160)) = v376 + v379
	F_errmsg(m, int32(_a_F_ri_ReportViolation_10), v18+int32(160))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L13
	} else {
		goto L88
	}
L88:
	;
	if v267 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	F_errtableconstraint(m, l2, v325)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L13
	} else {
		goto L95
	}
L90:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v18)+208))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+128)) = v391
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v18)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+132)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v18)+136)) = v390 + int32(4)
	F_errdetail(m, int32(_a_F_ri_ReportViolation_11), v18+int32(128))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L13
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+144)) = v403 + int32(4)
	F_errdetail(m, int32(_a_F_ri_ReportViolation_12), v18+int32(144))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
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
	F_errfinish(m, int32(_a_F_ri_ReportViolation_5), int32(2811), int32(_a_F_ri_ReportViolation_6))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
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
	v426 = m.ExcPending
	if v426 != 0 {
		goto L13
	} else {
		goto L98
	}
L98:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v429 = l0 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v427 + int32(4)
	F_errmsg(m, int32(_a_F_ri_ReportViolation_13), v18+int32(16))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L13
	} else {
		goto L99
	}
L99:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v18)+208))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v440
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v18)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v439 + int32(4)
	F_errdetail(m, int32(_a_F_ri_ReportViolation_11), v18)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L13
	} else {
		goto L100
	}
L100:
	;
	F_errtableconstraint(m, l2, v429)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L13
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_ri_ReportViolation_5), int32(2770), int32(_a_F_ri_ReportViolation_6))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
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
