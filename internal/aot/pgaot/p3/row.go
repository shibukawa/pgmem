package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SendRowDescriptionMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v332 int32
	_ = v332
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v425 int32
	_ = v425
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v25 = v23
	goto L3
L2:
	;
	v25 = int32(0)
	goto L3
L3:
	;
	F_resetStringInfo(m, l0)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(84)
	goto L4
L4:
	;
	F_enlargeStringInfo(m, l0, int32(2))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = int32(8)
	v41 = v22<<(uint(v35)%32) | int32(base.Ui32(v22&int32(65280))>>(uint(v35)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v32+v33))) = uint16(v41)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v32 + int32(2)
	F_enlargeStringInfo(m, l0, v22*int32(274))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if int32(0) < v22 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v58 = v25
	v63 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_pq_endmessage_reuse(m, l0)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L5
	} else {
		goto L80
	}
L11:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v77 = l1 + int32(20) + v71<<(uint(int32(4))%32) + v63*int32(100)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+68))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v79
	v83 = F_getBaseTypeAndTypmod(m, v78, v20+int32(12))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	if v58 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if l3 != 0 {
		goto L28
	} else {
		goto L29
	}
L15:
	;
	v140 = v92 + int32(4)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v140) < base.Ui32(v142+v143<<(uint(int32(2))%32)) {
		goto L25
	} else {
		goto L26
	}
L16:
	;
	v137 = int32(0)
	v156 = v137
	v162 = v137
	v165 = v134
	goto L14
L17:
	;
	v134 = int32(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v92 = v58
	goto L20
L20:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+26)))
	if v106 != int32(1) {
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v156 = v118
	v162 = v118
	v165 = v109
	goto L14
L22:
	;
	v109 = int32(0)
	v111 = v92 + int32(4)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v112+v113<<(uint(int32(2))%32)) <= base.Ui32(v111) {
		v134 = v109
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v118 = int32(0)
	if v111 != 0 {
		v92 = v111
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v148 = v140
	goto L27
L26:
	;
	v148 = int32(0)
	goto L27
L27:
	;
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+24)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v105)+20))
	v156 = v150
	v162 = v149
	v165 = v148
	goto L14
L28:
	;
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3+v63<<(uint(int32(1))%32)))))
	v173 = v171
	goto L30
L29:
	;
	v173 = int32(0)
	goto L30
L30:
	;
	v175 = v77 + int32(4)
	if v175&int32(3) == int32(0) {
		v199 = v175
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v316 = int32(24)
	v318 = int32(65280)
	v320 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v313+v314))) = v156<<(uint(v316)%32) | v156&v318<<(uint(v320)%32) | (int32(base.Ui32(v156)>>(uint(v320)%32))&v318 | int32(base.Ui32(v156)>>(uint(v316)%32)))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v340 = v162<<(uint(v320)%32) | int32(base.Ui32(v162&v318)>>(uint(v320)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v332+v313)+4)) = uint16(v340)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v342+v313)+6)) = v83<<(uint(v316)%32) | v83&v318<<(uint(v320)%32) | (int32(base.Ui32(v83)>>(uint(v320)%32))&v318 | int32(base.Ui32(v83)>>(uint(v316)%32)))
	v361 = v313 + int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v361
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+72)))
	v370 = v365<<(uint(v320)%32) | int32(base.Ui32(v365)>>(uint(v320)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v363+v361))) = uint16(v370)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v372+v313)+12)) = v374<<(uint(v316)%32) | v374&v318<<(uint(v320)%32) | (int32(base.Ui32(v374)>>(uint(v320)%32))&v318 | int32(base.Ui32(v374)>>(uint(v316)%32)))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v399 = v173<<(uint(v320)%32) | int32(base.Ui32(v173&v318)>>(uint(v320)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v391+v313)+16)) = uint16(v399)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v313 + int32(18)
	v405 = v63 + int32(1)
	if v405 != v22 {
		v58 = v165
		v63 = v405
		goto L11
	} else {
		goto L79
	}
L32:
	;
	v233 = F_pg_server_to_client(m, v175, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L5
	} else {
		goto L49
	}
L33:
	;
	v232 = v224 - v175
	goto L32
L34:
	;
	v203 = v199
	goto L43
L35:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v183 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v232 = int32(0)
	goto L32
L37:
	;
	goto L38
L38:
	;
	v188 = v175
	goto L39
L39:
	;
	v192 = v188 + int32(1)
	if v192&int32(3) == int32(0) {
		v199 = v192
		goto L34
	} else {
		goto L41
	}
L40:
	;
	v224 = v192
	goto L33
L41:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v197 != 0 {
		v188 = v192
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v212 = int32(-2139062144)
	if (int32(16843008)-v209|v209)&v212 == v212 {
		v203 = v203 + int32(4)
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v218 = v203
	goto L46
L45:
	;
	goto L44
L46:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	if v222 != 0 {
		v218 = v218 + int32(1)
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v224 = v218
	goto L33
L48:
	;
	goto L47
L49:
	;
	if v233 != v175 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v233&int32(3) == int32(0) {
		v262 = v233
		goto L55
	} else {
		goto L56
	}
L51:
	;
	goto L52
L52:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v307 = v232 + int32(1)
	if v307 != 0 {
		goto L76
	} else {
		goto L77
	}
L53:
	;
	v297 = v295 + int32(1)
	if v297 != 0 {
		goto L71
	} else {
		goto L72
	}
L54:
	;
	v295 = v287 - v233
	goto L53
L55:
	;
	v266 = v262
	goto L64
L56:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if v246 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v295 = int32(0)
	goto L53
L58:
	;
	goto L59
L59:
	;
	v251 = v233
	goto L60
L60:
	;
	v255 = v251 + int32(1)
	if v255&int32(3) == int32(0) {
		v262 = v255
		goto L55
	} else {
		goto L62
	}
L61:
	;
	v287 = v255
	goto L54
L62:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255))))
	if v260 != 0 {
		v251 = v255
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v275 = int32(-2139062144)
	if (int32(16843008)-v272|v272)&v275 == v275 {
		v266 = v266 + int32(4)
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v281 = v266
	goto L67
L66:
	;
	goto L65
L67:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	if v285 != 0 {
		v281 = v281 + int32(1)
		goto L67
	} else {
		goto L69
	}
L68:
	;
	v287 = v281
	goto L54
L69:
	;
	goto L68
L70:
	;
	F_pfree(m, v233)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L5
	} else {
		goto L74
	}
L71:
	;
	v298 = F__emscripten_memcpy_bulkmem(m, v236+v237, v233, v297)
	mBase = m.M
	goto L73
L72:
	;
	goto L73
L73:
	;
	goto L70
L74:
	;
	v313 = v236 + v297
	goto L31
L75:
	;
	v313 = v303 + v307
	goto L31
L76:
	;
	v308 = F__emscripten_memcpy_bulkmem(m, v303+v304, v233, v307)
	mBase = m.M
	goto L78
L77:
	;
	goto L78
L78:
	;
	goto L75
L79:
	;
	goto L12
L80:
	;
	m.G0 = v20 + int32(16)
	return
}
func F_row_is_in_frame(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v25 int64
	_ = v25
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v77 int64
	_ = v77
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v96 int64
	_ = v96
	var v106 int64
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int64
	_ = v112
	var v115 int32
	_ = v115
	var v116 int64
	_ = v116
	var v120 int32
	_ = v120
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	F_update_frameheadpos(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+184))
		if l1 < v19 {
			v120 = v4
			m.G0 = v12 + int32(16)
			return v120
		} else {
			if v14&int32(1024) != 0 {
				if v14&int32(4) != 0 {
					v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
					if l1 <= v25 {
						if v14&int32(32768) != 0 {
							v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
							if l1 != v96 {
								v120 = int32(1)
							} else {
								v120 = v4
							}
							m.G0 = v12 + int32(16)
							return v120
						} else {
							if v14&int32(65536) == int32(0) {
								if v14&int32(131072) == int32(0) {
									v120 = int32(1)
									m.G0 = v12 + int32(16)
									return v120
								} else {
									v106 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
									if l1 == v106 {
										v120 = int32(1)
										m.G0 = v12 + int32(16)
										return v120
									} else {
										v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
										if v109 == int32(0) {
											v120 = v4
											m.G0 = v12 + int32(16)
											return v120
										} else {
											v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
											if l1 < v112 {
												v120 = int32(1)
												m.G0 = v12 + int32(16)
												return v120
											} else {
												F_update_grouptailpos(m, l0)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
													if l1 < v116 {
														v120 = v4
													} else {
														v120 = int32(1)
													}
													m.G0 = v12 + int32(16)
													return v120
												}
											}
										}
									}
								}
							} else {
								v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
								if v109 == int32(0) {
									v120 = v4
									m.G0 = v12 + int32(16)
									return v120
								} else {
									v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
									if l1 < v112 {
										v120 = int32(1)
										m.G0 = v12 + int32(16)
										return v120
									} else {
										F_update_grouptailpos(m, l0)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return int32(0)
										} else {
											v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
											if l1 < v116 {
												v120 = v4
											} else {
												v120 = int32(1)
											}
											m.G0 = v12 + int32(16)
											return v120
										}
									}
								}
							}
						}
					} else {
						v120 = int32(-1)
						m.G0 = v12 + int32(16)
						return v120
					}
				} else {
					if v14&int32(10) == int32(0) {
						if v14&int32(32768) != 0 {
							v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
							if l1 != v96 {
								v120 = int32(1)
							} else {
								v120 = v4
							}
							m.G0 = v12 + int32(16)
							return v120
						} else {
							if v14&int32(65536) == int32(0) {
								if v14&int32(131072) == int32(0) {
									v120 = int32(1)
									m.G0 = v12 + int32(16)
									return v120
								} else {
									v106 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
									if l1 == v106 {
										v120 = int32(1)
										m.G0 = v12 + int32(16)
										return v120
									} else {
										v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
										if v109 == int32(0) {
											v120 = v4
											m.G0 = v12 + int32(16)
											return v120
										} else {
											v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
											if l1 < v112 {
												v120 = int32(1)
												m.G0 = v12 + int32(16)
												return v120
											} else {
												F_update_grouptailpos(m, l0)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
													if l1 < v116 {
														v120 = v4
													} else {
														v120 = int32(1)
													}
													m.G0 = v12 + int32(16)
													return v120
												}
											}
										}
									}
								}
							} else {
								v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
								if v109 == int32(0) {
									v120 = v4
									m.G0 = v12 + int32(16)
									return v120
								} else {
									v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
									if l1 < v112 {
										v120 = int32(1)
										m.G0 = v12 + int32(16)
										return v120
									} else {
										F_update_grouptailpos(m, l0)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return int32(0)
										} else {
											v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
											if l1 < v116 {
												v120 = v4
											} else {
												v120 = int32(1)
											}
											m.G0 = v12 + int32(16)
											return v120
										}
									}
								}
							}
						}
					} else {
						v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
						if l1 <= v32 {
							if v14&int32(32768) != 0 {
								v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
								if l1 != v96 {
									v120 = int32(1)
								} else {
									v120 = v4
								}
								m.G0 = v12 + int32(16)
								return v120
							} else {
								if v14&int32(65536) == int32(0) {
									if v14&int32(131072) == int32(0) {
										v120 = int32(1)
										m.G0 = v12 + int32(16)
										return v120
									} else {
										v106 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
										if l1 == v106 {
											v120 = int32(1)
											m.G0 = v12 + int32(16)
											return v120
										} else {
											v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
											if v109 == int32(0) {
												v120 = v4
												m.G0 = v12 + int32(16)
												return v120
											} else {
												v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
												if l1 < v112 {
													v120 = int32(1)
													m.G0 = v12 + int32(16)
													return v120
												} else {
													F_update_grouptailpos(m, l0)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
														if l1 < v116 {
															v120 = v4
														} else {
															v120 = int32(1)
														}
														m.G0 = v12 + int32(16)
														return v120
													}
												}
											}
										}
									}
								} else {
									v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
									if v109 == int32(0) {
										v120 = v4
										m.G0 = v12 + int32(16)
										return v120
									} else {
										v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
										if l1 < v112 {
											v120 = int32(1)
											m.G0 = v12 + int32(16)
											return v120
										} else {
											F_update_grouptailpos(m, l0)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
												if l1 < v116 {
													v120 = v4
												} else {
													v120 = int32(1)
												}
												m.G0 = v12 + int32(16)
												return v120
											}
										}
									}
								}
							}
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+96))
							if v35 == int32(0) {
								if v14&int32(32768) != 0 {
									v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
									if l1 != v96 {
										v120 = int32(1)
									} else {
										v120 = v4
									}
									m.G0 = v12 + int32(16)
									return v120
								} else {
									if v14&int32(65536) == int32(0) {
										if v14&int32(131072) == int32(0) {
											v120 = int32(1)
											m.G0 = v12 + int32(16)
											return v120
										} else {
											v106 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
											if l1 == v106 {
												v120 = int32(1)
												m.G0 = v12 + int32(16)
												return v120
											} else {
												v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
												if v109 == int32(0) {
													v120 = v4
													m.G0 = v12 + int32(16)
													return v120
												} else {
													v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
													if l1 < v112 {
														v120 = int32(1)
														m.G0 = v12 + int32(16)
														return v120
													} else {
														F_update_grouptailpos(m, l0)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return int32(0)
														} else {
															v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
															if l1 < v116 {
																v120 = v4
															} else {
																v120 = int32(1)
															}
															m.G0 = v12 + int32(16)
															return v120
														}
													}
												}
											}
										}
									} else {
										v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
										if v109 == int32(0) {
											v120 = v4
											m.G0 = v12 + int32(16)
											return v120
										} else {
											v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
											if l1 < v112 {
												v120 = int32(1)
												m.G0 = v12 + int32(16)
												return v120
											} else {
												F_update_grouptailpos(m, l0)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
													if l1 < v116 {
														v120 = v4
													} else {
														v120 = int32(1)
													}
													m.G0 = v12 + int32(16)
													return v120
												}
											}
										}
									}
								}
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
								v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
								*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v39
								*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = l2
								v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
								if v42 == int32(0) {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
									F_MemoryContextReset(m, v45)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										if v14&int32(32768) != 0 {
											v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
											if l1 != v96 {
												v120 = int32(1)
											} else {
												v120 = v4
											}
											m.G0 = v12 + int32(16)
											return v120
										} else {
											if v14&int32(65536) == int32(0) {
												if v14&int32(131072) == int32(0) {
													v120 = int32(1)
													m.G0 = v12 + int32(16)
													return v120
												} else {
													v106 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
													if l1 == v106 {
														v120 = int32(1)
														m.G0 = v12 + int32(16)
														return v120
													} else {
														v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
														v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
														if v109 == int32(0) {
															v120 = v4
															m.G0 = v12 + int32(16)
															return v120
														} else {
															v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
															if l1 < v112 {
																v120 = int32(1)
																m.G0 = v12 + int32(16)
																return v120
															} else {
																F_update_grouptailpos(m, l0)
																mBase = m.M
																v115 = m.ExcPending
																if v115 != 0 {
																	return int32(0)
																} else {
																	v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
																	if l1 < v116 {
																		v120 = v4
																	} else {
																		v120 = int32(1)
																	}
																	m.G0 = v12 + int32(16)
																	return v120
																}
															}
														}
													}
												}
											} else {
												v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
												if v109 == int32(0) {
													v120 = v4
													m.G0 = v12 + int32(16)
													return v120
												} else {
													v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
													if l1 < v112 {
														v120 = int32(1)
														m.G0 = v12 + int32(16)
														return v120
													} else {
														F_update_grouptailpos(m, l0)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return int32(0)
														} else {
															v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
															if l1 < v116 {
																v120 = v4
															} else {
																v120 = int32(1)
															}
															m.G0 = v12 + int32(16)
															return v120
														}
													}
												}
											}
										}
									}
								} else {
									v48 = int32(4476144)
									v49 = *(*int32)(unsafe.Add(mBase, _consts[0]))
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v51
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
									v56 = m.T0[v55].(func(*base.Module, int32, int32, int32) int32)(m, v42, v38, v12+int32(15))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v49
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
										F_MemoryContextReset(m, v60)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											if v56 != 0 {
												if v14&int32(32768) != 0 {
													v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
													if l1 != v96 {
														v120 = int32(1)
													} else {
														v120 = v4
													}
													m.G0 = v12 + int32(16)
													return v120
												} else {
													if v14&int32(65536) == int32(0) {
														if v14&int32(131072) == int32(0) {
															v120 = int32(1)
															m.G0 = v12 + int32(16)
															return v120
														} else {
															v106 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
															if l1 == v106 {
																v120 = int32(1)
																m.G0 = v12 + int32(16)
																return v120
															} else {
																v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
																v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
																if v109 == int32(0) {
																	v120 = v4
																	m.G0 = v12 + int32(16)
																	return v120
																} else {
																	v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
																	if l1 < v112 {
																		v120 = int32(1)
																		m.G0 = v12 + int32(16)
																		return v120
																	} else {
																		F_update_grouptailpos(m, l0)
																		mBase = m.M
																		v115 = m.ExcPending
																		if v115 != 0 {
																			return int32(0)
																		} else {
																			v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
																			if l1 < v116 {
																				v120 = v4
																			} else {
																				v120 = int32(1)
																			}
																			m.G0 = v12 + int32(16)
																			return v120
																		}
																	}
																}
															}
														}
													} else {
														v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
														v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
														if v109 == int32(0) {
															v120 = v4
															m.G0 = v12 + int32(16)
															return v120
														} else {
															v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
															if l1 < v112 {
																v120 = int32(1)
																m.G0 = v12 + int32(16)
																return v120
															} else {
																F_update_grouptailpos(m, l0)
																mBase = m.M
																v115 = m.ExcPending
																if v115 != 0 {
																	return int32(0)
																} else {
																	v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
																	if l1 < v116 {
																		v120 = v4
																	} else {
																		v120 = int32(1)
																	}
																	m.G0 = v12 + int32(16)
																	return v120
																}
															}
														}
													}
												}
											} else {
												v120 = int32(-1)
												m.G0 = v12 + int32(16)
												return v120
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				if v14&int32(20480) == int32(0) {
					if v14&int32(32768) != 0 {
						v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
						if l1 != v96 {
							v120 = int32(1)
						} else {
							v120 = v4
						}
						m.G0 = v12 + int32(16)
						return v120
					} else {
						if v14&int32(65536) == int32(0) {
							if v14&int32(131072) == int32(0) {
								v120 = int32(1)
								m.G0 = v12 + int32(16)
								return v120
							} else {
								v106 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
								if l1 == v106 {
									v120 = int32(1)
									m.G0 = v12 + int32(16)
									return v120
								} else {
									v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
									if v109 == int32(0) {
										v120 = v4
										m.G0 = v12 + int32(16)
										return v120
									} else {
										v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
										if l1 < v112 {
											v120 = int32(1)
											m.G0 = v12 + int32(16)
											return v120
										} else {
											F_update_grouptailpos(m, l0)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
												if l1 < v116 {
													v120 = v4
												} else {
													v120 = int32(1)
												}
												m.G0 = v12 + int32(16)
												return v120
											}
										}
									}
								}
							}
						} else {
							v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
							if v109 == int32(0) {
								v120 = v4
								m.G0 = v12 + int32(16)
								return v120
							} else {
								v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
								if l1 < v112 {
									v120 = int32(1)
									m.G0 = v12 + int32(16)
									return v120
								} else {
									F_update_grouptailpos(m, l0)
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
										if l1 < v116 {
											v120 = v4
										} else {
											v120 = int32(1)
										}
										m.G0 = v12 + int32(16)
										return v120
									}
								}
							}
						}
					}
				} else {
					if v14&int32(4) != 0 {
						v70 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
						v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
						if v14&int32(4096) != 0 {
							v77 = int64(0) - v73
						} else {
							v77 = v73
						}
						if l1 <= v70+v77 {
							if v14&int32(32768) != 0 {
								v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
								if l1 != v96 {
									v120 = int32(1)
								} else {
									v120 = v4
								}
								m.G0 = v12 + int32(16)
								return v120
							} else {
								if v14&int32(65536) == int32(0) {
									if v14&int32(131072) == int32(0) {
										v120 = int32(1)
										m.G0 = v12 + int32(16)
										return v120
									} else {
										v106 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
										if l1 == v106 {
											v120 = int32(1)
											m.G0 = v12 + int32(16)
											return v120
										} else {
											v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
											if v109 == int32(0) {
												v120 = v4
												m.G0 = v12 + int32(16)
												return v120
											} else {
												v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
												if l1 < v112 {
													v120 = int32(1)
													m.G0 = v12 + int32(16)
													return v120
												} else {
													F_update_grouptailpos(m, l0)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
														if l1 < v116 {
															v120 = v4
														} else {
															v120 = int32(1)
														}
														m.G0 = v12 + int32(16)
														return v120
													}
												}
											}
										}
									}
								} else {
									v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
									if v109 == int32(0) {
										v120 = v4
										m.G0 = v12 + int32(16)
										return v120
									} else {
										v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
										if l1 < v112 {
											v120 = int32(1)
											m.G0 = v12 + int32(16)
											return v120
										} else {
											F_update_grouptailpos(m, l0)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
												if l1 < v116 {
													v120 = v4
												} else {
													v120 = int32(1)
												}
												m.G0 = v12 + int32(16)
												return v120
											}
										}
									}
								}
							}
						} else {
							v120 = int32(-1)
							m.G0 = v12 + int32(16)
							return v120
						}
					} else {
						if v14&int32(10) == int32(0) {
							if v14&int32(32768) != 0 {
								v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
								if l1 != v96 {
									v120 = int32(1)
								} else {
									v120 = v4
								}
								m.G0 = v12 + int32(16)
								return v120
							} else {
								if v14&int32(65536) == int32(0) {
									if v14&int32(131072) == int32(0) {
										v120 = int32(1)
										m.G0 = v12 + int32(16)
										return v120
									} else {
										v106 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
										if l1 == v106 {
											v120 = int32(1)
											m.G0 = v12 + int32(16)
											return v120
										} else {
											v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
											if v109 == int32(0) {
												v120 = v4
												m.G0 = v12 + int32(16)
												return v120
											} else {
												v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
												if l1 < v112 {
													v120 = int32(1)
													m.G0 = v12 + int32(16)
													return v120
												} else {
													F_update_grouptailpos(m, l0)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
														if l1 < v116 {
															v120 = v4
														} else {
															v120 = int32(1)
														}
														m.G0 = v12 + int32(16)
														return v120
													}
												}
											}
										}
									}
								} else {
									v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
									if v109 == int32(0) {
										v120 = v4
										m.G0 = v12 + int32(16)
										return v120
									} else {
										v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
										if l1 < v112 {
											v120 = int32(1)
											m.G0 = v12 + int32(16)
											return v120
										} else {
											F_update_grouptailpos(m, l0)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
												if l1 < v116 {
													v120 = v4
												} else {
													v120 = int32(1)
												}
												m.G0 = v12 + int32(16)
												return v120
											}
										}
									}
								}
							}
						} else {
							F_update_frametailpos(m, l0)
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return int32(0)
							} else {
								v87 = *(*int64)(unsafe.Add(mBase, uint32(l0)+192))
								if l1 < v87 {
									if v14&int32(32768) != 0 {
										v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
										if l1 != v96 {
											v120 = int32(1)
										} else {
											v120 = v4
										}
										m.G0 = v12 + int32(16)
										return v120
									} else {
										if v14&int32(65536) == int32(0) {
											if v14&int32(131072) == int32(0) {
												v120 = int32(1)
												m.G0 = v12 + int32(16)
												return v120
											} else {
												v106 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
												if l1 == v106 {
													v120 = int32(1)
													m.G0 = v12 + int32(16)
													return v120
												} else {
													v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
													if v109 == int32(0) {
														v120 = v4
														m.G0 = v12 + int32(16)
														return v120
													} else {
														v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
														if l1 < v112 {
															v120 = int32(1)
															m.G0 = v12 + int32(16)
															return v120
														} else {
															F_update_grouptailpos(m, l0)
															mBase = m.M
															v115 = m.ExcPending
															if v115 != 0 {
																return int32(0)
															} else {
																v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
																if l1 < v116 {
																	v120 = v4
																} else {
																	v120 = int32(1)
																}
																m.G0 = v12 + int32(16)
																return v120
															}
														}
													}
												}
											}
										} else {
											v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+96))
											if v109 == int32(0) {
												v120 = v4
												m.G0 = v12 + int32(16)
												return v120
											} else {
												v112 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
												if l1 < v112 {
													v120 = int32(1)
													m.G0 = v12 + int32(16)
													return v120
												} else {
													F_update_grouptailpos(m, l0)
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return int32(0)
													} else {
														v116 = *(*int64)(unsafe.Add(mBase, uint32(l0)+352))
														if l1 < v116 {
															v120 = v4
														} else {
															v120 = int32(1)
														}
														m.G0 = v12 + int32(16)
														return v120
													}
												}
											}
										}
									}
								} else {
									v120 = int32(-1)
									m.G0 = v12 + int32(16)
									return v120
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_row_security_policy_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return base.B2i32(v4 != int32(0))
L2:
	;
	goto L3
L3:
	;
	if v4 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(-1)
L5:
	;
	goto L6
L6:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if v19 == int32(0) {
		v38 = v18
		v39 = v19
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return v39 - v38
L8:
	;
	goto L7
L9:
	;
	if v18 != v19 {
		v38 = v18
		v39 = v19
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v23 = v6
	v24 = v4
	goto L11
L11:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v28 == int32(0) {
		v38 = v27
		v39 = v28
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v38 = v27
	v39 = v28
	goto L8
L13:
	;
	v31 = int32(1)
	if v27 == v28 {
		v23 = v23 + v31
		v24 = v24 + v31
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
