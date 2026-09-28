package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RI_FKey_noaction_upd(m *base.Module, l0 int32) int64 {
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
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_restrict(m, v8, int32(1))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			return int64(0)
		}
	}
}
func F_RI_FKey_setnull_upd(m *base.Module, l0 int32) int64 {
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
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		F_ri_set(m, v8, int32(1), int32(2))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			return int64(0)
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	v8 = m.G0
	v10 = v8 - int32(192)
	m.G0 = v10
	if l1 != 0 {
		v14 = F_SearchSysCache1(m, int32(16), base.I64_extend_i32_u(l1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			if v14 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
					F_errmsg_internal(m, int32(_a_F_ri_GenerateQualCollation_0), v10)
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ri_GenerateQualCollation_1), int32(2199), int32(_a_F_ri_GenerateQualCollation_2))
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
				v20 = v18 + v19
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+68))
				v22 = F_get_namespace_name(m, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v24 = int32(34)
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)) = uint8(v24)
					v29 = v22
					v31 = v10 + int32(48)
					for {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
						if v35 != int32(34) {
						} else {
							v42 = int32(34)
							*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)) = uint8(v42)
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
							v47 = v44
							v48 = v31 + int32(2)
							*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v47)
							v29 = v29 + int32(1)
							v31 = v48
							continue
						}
						if v35 == int32(0) {
							break
						} else {
							v47 = v35
							v48 = v31 + int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v47)
							v29 = v29 + int32(1)
							v31 = v48
							continue
						}
						break
					}
					v52 = int32(34)
					*(*uint16)(unsafe.Add(mBase, uint32(v31)+1)) = uint16(v52)
					v55 = v10 + int32(48)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v55
					F_appendStringInfo(m, l0, int32(_a_F_ri_GenerateQualCollation_3), v10+int32(32))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						v62 = int32(34)
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)) = uint8(v62)
						v67 = v20 + int32(4)
						v69 = v55
						for {
							v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
							if v73 != int32(34) {
							} else {
								v80 = int32(34)
								*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)) = uint8(v80)
								v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
								v85 = v82
								v86 = v69 + int32(2)
								*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v85)
								v67 = v67 + int32(1)
								v69 = v86
								continue
							}
							if v73 == int32(0) {
								break
							} else {
								v85 = v73
								v86 = v69 + int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v85)
								v67 = v67 + int32(1)
								v69 = v86
								continue
							}
							break
						}
						v90 = int32(34)
						*(*uint16)(unsafe.Add(mBase, uint32(v69)+1)) = uint16(v90)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v10 + int32(48)
						F_appendStringInfo(m, l0, int32(_a_F_ri_GenerateQualCollation_4), v10+int32(16))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							F_ReleaseCatCache(m, v14)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v134 int32
	_ = v134
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int64
	_ = v181
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int64
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v255 int32
	_ = v255
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v333 int64
	_ = v333
	var v336 int32
	_ = v336
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	v11 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(624)
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
	v109 = v28
	goto L7
L7:
	;
	v117 = int32(0)
	if base.B2i32(l5 == v117)|base.B2i32(v109 <= v117) != 0 {
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
	v109 = v70
	goto L7
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	m.T0[v64].(func(*base.Module, int32, int32))(m, l6, v60)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v70 = v49
	goto L15
L15:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l6)+16))
	v73 = v60 - int32(1)
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v71+v73<<(uint(int32(3))%32))))
	v78 = int32(32)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v73))))
	if v85 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	return int32(0)
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v70 = v69
	goto L15
L18:
	;
	v86 = int32(110)
	goto L20
L19:
	;
	v86 = v78
	goto L20
L20:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23+v78+v48))) = uint8(v86)
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(96)+v48<<(uint(int32(3))%32)))) = v77
	v95 = v48 + int32(1)
	if v95 < v70 {
		v48 = v95
		v49 = v70
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
	v134 = int32(236)
	goto L25
L24:
	;
	v134 = int32(172)
	goto L25
L25:
	;
	v148 = int32(0)
	v149 = v109
	goto L26
L26:
	;
	v160 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0+v134+v148<<(uint(int32(1))%32)))))
	v161 = int32(*(*int16)(unsafe.Add(mBase, uint32(l5)+6)))
	if v161 < v160 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L1
L28:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	m.T0[v164].(func(*base.Module, int32, int32))(m, l5, v160)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L16
	} else {
		goto L31
	}
L29:
	;
	v168 = v149
	goto L30
L30:
	;
	v170 = v160 - int32(1)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170+v171))))
	v174 = int32(3)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v177+v170<<(uint(v174)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(96)+v109<<(uint(int32(3))%32)+v148<<(uint(v174)%32)))) = v181
	if v173 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v168 = v167
	goto L30
L32:
	;
	v186 = int32(110)
	goto L34
L33:
	;
	v186 = int32(32)
	goto L34
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v148+(v23+int32(32)+v109)))) = uint8(v186)
	v189 = v148 + int32(1)
	if v189 < v168 {
		v148 = v189
		v149 = v168
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
	v197 = int32(236)
	goto L39
L38:
	;
	v197 = int32(172)
	goto L39
L39:
	;
	v210 = v11
	v211 = v28
	goto L40
L40:
	;
	v222 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0+v197+v210<<(uint(int32(1))%32)))))
	v223 = int32(*(*int16)(unsafe.Add(mBase, uint32(l5)+6)))
	if v223 < v222 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L1
L42:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+16))
	m.T0[v226].(func(*base.Module, int32, int32))(m, l5, v222)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L16
	} else {
		goto L45
	}
L43:
	;
	v230 = v211
	goto L44
L44:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v233 = v222 - int32(1)
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v231+v233<<(uint(int32(3))%32))))
	v238 = int32(32)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243+v233))))
	if v245 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v230 = v229
	goto L44
L46:
	;
	v246 = int32(110)
	goto L48
L47:
	;
	v246 = v238
	goto L48
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23+v238+v210))) = uint8(v246)
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(96)+v210<<(uint(int32(3))%32)))) = v237
	v255 = v210 + int32(1)
	if v255 < v230 {
		v210 = v255
		v211 = v230
		goto L40
	} else {
		goto L49
	}
L49:
	;
	goto L41
L50:
	;
	v277 = l4
	goto L52
L51:
	;
	v277 = l3
	goto L52
L52:
	;
	v278 = int32(0)
	if l8 == v278 {
		v293 = v278
		v294 = v278
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _c_F_ri_PerformCheck[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(620)))) = v300
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_ri_PerformCheck[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v23+int32(616)))) = v303
	goto L59
L54:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_ri_PerformCheck[2]))
	if v284 < int32(2) {
		v293 = v278
		v294 = int32(0)
		goto L53
	} else {
		goto L55
	}
L55:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L16
	} else {
		goto L56
	}
L56:
	;
	v289 = F_GetLatestSnapshot(m)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L16
	} else {
		goto L57
	}
L57:
	;
	v291 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L16
	} else {
		goto L58
	}
L58:
	;
	v293 = v289
	v294 = v291
	goto L53
L59:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v277)+48))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+80))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v23)+616))
	*(*int32)(unsafe.Add(mBase, _c_F_ri_PerformCheck[1])) = v307 | int32(5)
	*(*int32)(unsafe.Add(mBase, _c_F_ri_PerformCheck[0])) = v306
	goto L60
L60:
	;
	v321 = F_SPI_execute_snapshot(m, l2, v23+int32(96), v23+int32(32), v293, v294, int32(0), base.B2i32(l9 == int32(5)))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L16
	} else {
		goto L61
	}
L61:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v23)+620))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v23)+616))
	*(*int32)(unsafe.Add(mBase, _c_F_ri_PerformCheck[1])) = v324
	*(*int32)(unsafe.Add(mBase, _c_F_ri_PerformCheck[0])) = v323
	goto L62
L62:
	;
	if int32(0) <= v321 {
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
	v369 = m.ExcPending
	if v369 != 0 {
		goto L16
	} else {
		goto L77
	}
L65:
	;
	if v321 != l9 {
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
	v354 = m.ExcPending
	if v354 != 0 {
		goto L16
	} else {
		goto L73
	}
L68:
	;
	v333 = *(*int64)(unsafe.Add(mBase, _c_F_ri_PerformCheck[3]))
	if l9 != int32(5) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	m.G0 = v23 + int32(624)
	return base.B2i32(v333 != int64(0))
L70:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v336 == int32(2) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	if base.B2i32(v333 == int64(0)) != base.B2i32(v336 != int32(1)) {
		goto L63
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	v355 = F_SPI_result_code_string(m, v321)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L16
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v355
	F_errmsg_internal(m, int32(_a_F_ri_PerformCheck_0), v23)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L16
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_ri_PerformCheck_1), int32(2734), int32(_a_F_ri_PerformCheck_2))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
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
	v372 = m.ExcPending
	if v372 != 0 {
		goto L16
	} else {
		goto L78
	}
L78:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l4)+48))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = l0 + int32(20)
	v378 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v374 + v378
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v373 + v378
	F_errmsg(m, int32(_a_F_ri_PerformCheck_3), v23+int32(16))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L16
	} else {
		goto L79
	}
L79:
	;
	F_errhint(m, int32(_a_F_ri_PerformCheck_4), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L16
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_ri_PerformCheck_1), int32(2743), int32(_a_F_ri_PerformCheck_2))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
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
	v398 = l6
	goto L84
L83:
	;
	v398 = l5
	goto L84
L84:
	;
	v399 = int32(0)
	F_ri_ReportViolation(m, l0, l4, l3, v398, v399, v336, l7, v399)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
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
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
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
					v92 = v49
					v96 = F_hash_search(m, v92, l3, int32(1), v9+int32(-48))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v36
						m.G0 = v11 - int32(-64)
						return v36
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(3092376453124)
					v56 = v9 + int32(-48)
					v58 = F_hash_create(m, int32(_a_F_ri_PlanCheck_0), int64(64), v56, int32(40))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_ri_PlanCheck[3])) = v58
						F_CacheRegisterSyscacheCallback(m, int32(19), int32(1698), int64(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							F_CacheRegisterSyscacheCallback(m, int32(3), int32(1698), int64(0))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(51539607560)
								v77 = F_hash_create(m, int32(_a_F_ri_PlanCheck_1), int64(256), v56, int32(40))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_ri_PlanCheck[2])) = v77
									*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(292057776136)
									v86 = F_hash_create(m, int32(_a_F_ri_PlanCheck_2), int64(256), v56, int32(40))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_ri_PlanCheck[4])) = v86
										v90 = *(*int32)(unsafe.Add(mBase, _c_F_ri_PlanCheck[2]))
										v92 = v90
										v96 = F_hash_search(m, v92, l3, int32(1), v9+int32(-48))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v36
											m.G0 = v11 - int32(-64)
											return v36
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
			v106 = m.ExcPending
			if v106 != 0 {
				return int32(0)
			} else {
				v108 = *(*int32)(unsafe.Add(mBase, _c_F_ri_PlanCheck[5]))
				v109 = F_SPI_result_code_string(m, v108)
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v109
					F_errmsg_internal(m, int32(_a_F_ri_PlanCheck_3), v11)
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_ri_PlanCheck_4), int32(2602), int32(_a_F_ri_PlanCheck_5))
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
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
	var v142 int64
	_ = v142
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
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int64
	_ = v217
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v262 int32
	_ = v262
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	v17 = m.G0
	v19 = v17 - int32(224)
	m.G0 = v19
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
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+52))
	v35 = v34
	v37 = v32
	v38 = v33
	goto L1
L3:
	;
	v24 = l0 + int32(236)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	if l4 == int32(0) {
		v31 = l2
		v32 = v25
		v33 = v24
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v29 = l0 + int32(172)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if l4 != 0 {
		v35 = l4
		v37 = v30
		v38 = v29
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v35 = l4
	v37 = v25
	v38 = v24
	goto L1
L7:
	;
	v31 = l1
	v32 = v30
	v33 = v29
	goto L2
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L13
	} else {
		goto L97
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
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
	F_initStringInfo(m, v19+int32(208))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L13
	} else {
		goto L26
	}
L12:
	;
	v39 = int32(0)
	v42 = F_check_enable_rls(m, v37, v39, int32(1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	if v42 == int32(2) {
		v278 = v39
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_ri_ReportViolation[0]))
	v49 = F_pg_class_aclcheck(m, v37, v47, int64(2))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v49 == int32(0) {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v53 <= int32(0) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v65 = v39
	goto L19
L19:
	;
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38+v65<<(uint(int32(1))%32)))))
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_ri_ReportViolation[0]))
	v79 = F_pg_attribute_aclcheck(m, v37, v75, v77, int64(2))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L13
	} else {
		goto L21
	}
L20:
	;
	v262 = int32(0)
	goto L10
L21:
	;
	if v79 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v84 = v65 + int32(1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v84 < v85 {
		v65 = v84
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
	F_initStringInfo(m, v19+int32(192))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v113 <= int32(0) {
		v262 = int32(1)
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v120 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38))))
	v123 = v35 + v116<<(uint(int32(3))%32) + v120*int32(100)
	v124 = int32(*(*int16)(unsafe.Add(mBase, uint32(l3)+6)))
	if v124 < v120 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	m.T0[v127].(func(*base.Module, int32, int32))(m, l3, v120)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L13
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v132 = v120 - int32(1)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132+v133))))
	if v135 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v138+v132<<(uint(int32(3))%32))))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v123-int32(72))+68))
	F_getTypeOutputInfo(m, v145, v19+int32(188), v19+int32(187))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L13
	} else {
		goto L36
	}
L34:
	;
	v155 = int32(_a_F_ri_ReportViolation_0)
	goto L35
L35:
	;
	F_appendStringInfoString(m, v19+int32(208), v123-int32(68))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L38
	}
L36:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v19)+188))
	v153 = F_OidOutputFunctionCall(m, v152, v142)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	v155 = v153
	goto L35
L38:
	;
	F_appendStringInfoString(m, v19+int32(192), v155)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L13
	} else {
		goto L39
	}
L39:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v168 < int32(2) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v262 = int32(1)
	goto L10
L41:
	;
	goto L42
L42:
	;
	v182 = int32(1)
	goto L43
L43:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v195 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38+v182<<(uint(int32(1))%32)))))
	v198 = v35 + v188<<(uint(int32(3))%32) + v195*int32(100)
	v199 = int32(*(*int16)(unsafe.Add(mBase, uint32(l3)+6)))
	if v199 < v195 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v262 = v248
	goto L10
L45:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+16))
	m.T0[v202].(func(*base.Module, int32, int32))(m, l3, v195)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L13
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v207 = v195 - int32(1)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+v208))))
	if v210 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v213+v207<<(uint(int32(3))%32))))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v198-int32(72))+68))
	F_getTypeOutputInfo(m, v220, v19+int32(188), v19+int32(187))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L13
	} else {
		goto L52
	}
L50:
	;
	v230 = int32(_a_F_ri_ReportViolation_0)
	goto L51
L51:
	;
	v235 = v19 + int32(208)
	F_appendStringInfoString(m, v235, int32(_a_F_ri_ReportViolation_1))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L13
	} else {
		goto L54
	}
L52:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v19)+188))
	v228 = F_OidOutputFunctionCall(m, v227, v217)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L13
	} else {
		goto L53
	}
L53:
	;
	v230 = v228
	goto L51
L54:
	;
	v240 = v19 + int32(192)
	F_appendStringInfoString(m, v240, int32(_a_F_ri_ReportViolation_1))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	F_appendStringInfoString(m, v235, v198-int32(68))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	F_appendStringInfoString(m, v240, v230)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	v248 = int32(1)
	v250 = v182 + v248
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v250 < v251 {
		v182 = v250
		goto L43
	} else {
		goto L58
	}
L58:
	;
	goto L44
L59:
	;
	v278 = v262
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
	v293 = m.ExcPending
	if v293 != 0 {
		goto L13
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v337 = l0 + int32(20)
	if l6 != 0 {
		goto L74
	} else {
		goto L75
	}
L64:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v296 = l0 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v296
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v294 + int32(4)
	F_errmsg(m, int32(_a_F_ri_ReportViolation_2), v19-int32(-64))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	if v278 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	F_errtableconstraint(m, l2, v296)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L13
	} else {
		goto L72
	}
L67:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v19)+208))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v307
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v19)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v306 + int32(4)
	v317 = F_errdetail(m, int32(_a_F_ri_ReportViolation_3), v19+int32(32))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L13
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v319 + int32(4)
	v326 = F_errdetail(m, int32(_a_F_ri_ReportViolation_4), v19+int32(48))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
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
	F_errfinish(m, int32(_a_F_ri_ReportViolation_5), int32(3482), int32(_a_F_ri_ReportViolation_6))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
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
	v340 = m.ExcPending
	if v340 != 0 {
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
	v387 = m.ExcPending
	if v387 != 0 {
		goto L13
	} else {
		goto L87
	}
L77:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+116)) = v337
	v344 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+120)) = v342 + v344
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v341 + v344
	F_errmsg(m, int32(_a_F_ri_ReportViolation_7), v19+int32(112))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L13
	} else {
		goto L78
	}
L78:
	;
	if v278 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	F_errtableconstraint(m, l2, v337)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L13
	} else {
		goto L85
	}
L80:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v19)+208))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v356
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v19)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+84)) = v358
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v355 + int32(4)
	v366 = F_errdetail(m, int32(_a_F_ri_ReportViolation_8), v19+int32(80))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L13
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v368 + int32(4)
	v375 = F_errdetail(m, int32(_a_F_ri_ReportViolation_9), v19+int32(96))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
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
	F_errfinish(m, int32(_a_F_ri_ReportViolation_5), int32(3496), int32(_a_F_ri_ReportViolation_6))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
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
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+164)) = v337
	v391 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+168)) = v389 + v391
	*(*int32)(unsafe.Add(mBase, uint32(v19)+160)) = v388 + v391
	F_errmsg(m, int32(_a_F_ri_ReportViolation_10), v19+int32(160))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L13
	} else {
		goto L88
	}
L88:
	;
	if v278 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	F_errtableconstraint(m, l2, v337)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L13
	} else {
		goto L95
	}
L90:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v19)+208))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v403
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v19)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+132)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v19)+136)) = v402 + int32(4)
	v413 = F_errdetail(m, int32(_a_F_ri_ReportViolation_11), v19+int32(128))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L13
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = v415 + int32(4)
	v422 = F_errdetail(m, int32(_a_F_ri_ReportViolation_12), v19+int32(144))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
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
	F_errfinish(m, int32(_a_F_ri_ReportViolation_5), int32(3510), int32(_a_F_ri_ReportViolation_6))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
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
	v438 = m.ExcPending
	if v438 != 0 {
		goto L13
	} else {
		goto L98
	}
L98:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v441 = l0 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v439 + int32(4)
	F_errmsg(m, int32(_a_F_ri_ReportViolation_13), v19+int32(16))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L13
	} else {
		goto L99
	}
L99:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v19)+208))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v452
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v19)+192))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v451 + int32(4)
	v460 = F_errdetail(m, int32(_a_F_ri_ReportViolation_11), v19)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L13
	} else {
		goto L100
	}
L100:
	;
	F_errtableconstraint(m, l2, v441)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L13
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_ri_ReportViolation_5), int32(3469), int32(_a_F_ri_ReportViolation_6))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
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
