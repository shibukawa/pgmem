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
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v313 int32
	_ = v313
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
	v313 = m.ExcPending
	if v313 != 0 {
		goto L5
	} else {
		goto L46
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
	v176 = F_strlen(m, v175)
	mBase = m.M
	v177 = F_pg_server_to_client(m, v175, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L5
	} else {
		goto L32
	}
L31:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v204 = int32(24)
	v206 = int32(65280)
	v208 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v201+v202))) = v156<<(uint(v204)%32) | v156&v206<<(uint(v208)%32) | (int32(base.Ui32(v156)>>(uint(v208)%32))&v206 | int32(base.Ui32(v156)>>(uint(v204)%32)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v228 = v162<<(uint(v208)%32) | int32(base.Ui32(v162&v206)>>(uint(v208)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v220+v201)+4)) = uint16(v228)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v230+v201)+6)) = v83<<(uint(v204)%32) | v83&v206<<(uint(v208)%32) | (int32(base.Ui32(v83)>>(uint(v208)%32))&v206 | int32(base.Ui32(v83)>>(uint(v204)%32)))
	v249 = v201 + int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v249
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+72)))
	v258 = v253<<(uint(v208)%32) | int32(base.Ui32(v253)>>(uint(v208)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v251+v249))) = uint16(v258)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v260+v201)+12)) = v262<<(uint(v204)%32) | v262&v206<<(uint(v208)%32) | (int32(base.Ui32(v262)>>(uint(v208)%32))&v206 | int32(base.Ui32(v262)>>(uint(v204)%32)))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v287 = v173<<(uint(v208)%32) | int32(base.Ui32(v173&v206)>>(uint(v208)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v279+v201)+16)) = uint16(v287)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v201 + int32(18)
	v293 = v63 + int32(1)
	if v293 != v22 {
		v58 = v165
		v63 = v293
		goto L11
	} else {
		goto L45
	}
L32:
	;
	if v177 != v175 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v183 = F_strlen(m, v177)
	mBase = m.M
	v185 = v183 + int32(1)
	if v185 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L35
L35:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v195 = v176 + int32(1)
	if v195 != 0 {
		goto L42
	} else {
		goto L43
	}
L36:
	;
	F_pfree(m, v177)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L5
	} else {
		goto L40
	}
L37:
	;
	v186 = F__emscripten_memcpy_bulkmem(m, v180+v181, v177, v185)
	mBase = m.M
	goto L39
L38:
	;
	goto L39
L39:
	;
	goto L36
L40:
	;
	v201 = v180 + v185
	goto L31
L41:
	;
	v201 = v191 + v195
	goto L31
L42:
	;
	v196 = F__emscripten_memcpy_bulkmem(m, v191+v192, v177, v195)
	mBase = m.M
	goto L44
L43:
	;
	goto L44
L44:
	;
	goto L41
L45:
	;
	goto L12
L46:
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
									v48 = int32(4489152)
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
