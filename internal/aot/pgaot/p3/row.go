package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SendRowDescriptionMessage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
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
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v280 int32
	_ = v280
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v24 = v22
	goto L3
L2:
	;
	v24 = int32(0)
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
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = int32(8)
	v40 = v21<<(uint(v34)%32) | int32(base.Ui32(v21&int32(_a_F_SendRowDescriptionMessage_0))>>(uint(v34)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v31+v32))) = uint16(v40)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v31 + int32(2)
	F_enlargeStringInfo(m, l0, v21*int32(274))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if int32(0) < v21 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v57 = v24
	v64 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_pq_endmessage_reuse(m, l0)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L5
	} else {
		goto L42
	}
L11:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v73 = l1 + v67<<(uint(int32(4))%32) + v64*int32(100)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+88))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v75
	v79 = F_getBaseTypeAndTypmod(m, v74, v19+int32(12))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	if v57 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if l3 != 0 {
		goto L26
	} else {
		goto L27
	}
L15:
	;
	v153 = int32(0)
	v164 = v153
	v168 = v153
	v172 = v153
	goto L14
L16:
	;
	v89 = v57
	goto L17
L17:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+26)))
	if v100 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+24)))
	v112 = int32(8)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v99)+20))
	v118 = int32(16711935)
	v128 = v89 + int32(4)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v128) < base.Ui32(v130+v131<<(uint(int32(2))%32)) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v104 = v89 + int32(4)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v104) < base.Ui32(v105+v106<<(uint(int32(2))%32)) {
		v89 = v104
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	goto L18
L22:
	;
	goto L15
L23:
	;
	v136 = v128
	goto L25
L24:
	;
	v136 = int32(0)
	goto L25
L25:
	;
	v164 = base.I32_rotr(v117&v118, v112) | base.I32_rotr(v117, int32(24))&v118
	v168 = v111<<(uint(v112)%32) | int32(base.Ui32(v111)>>(uint(v112)%32))
	v172 = v136
	goto L14
L26:
	;
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3+v64<<(uint(int32(1))%32)))))
	v179 = int32(8)
	v186 = v178<<(uint(v179)%32) | int32(base.Ui32(v178)>>(uint(v179)%32))
	goto L28
L27:
	;
	v186 = int32(0)
	goto L28
L28:
	;
	v188 = v73 + int32(24)
	v189 = F_strlen(m, v188)
	mBase = m.M
	v190 = F_pg_server_to_client(m, v188, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L30
	}
L29:
	;
	v210 = v208 + v209
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v210+v211))) = v164
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint16)(unsafe.Add(mBase, uint32(v214+v210)+4)) = uint16(v168)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v219 = int32(24)
	v221 = int32(16711935)
	v225 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v217+v210)+6)) = base.I32_rotr(v79, v219)&v221 | base.I32_rotr(v79&v221, v225)
	v230 = v210 + int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v230
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73+int32(20))+72)))
	v239 = v234<<(uint(v225)%32) | int32(base.Ui32(v234)>>(uint(v225)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v232+v230))) = uint16(v239)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v241+v210)+12)) = base.I32_rotr(v243&v221, v225) | base.I32_rotr(v243, v219)&v221
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint16)(unsafe.Add(mBase, uint32(v254+v210)+16)) = uint16(v186)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v210 + int32(18)
	v261 = v64 + int32(1)
	if v261 != v21 {
		v57 = v172
		v64 = v261
		goto L11
	} else {
		goto L41
	}
L30:
	;
	if v190 != v188 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v194 = F_strlen(m, v190)
	mBase = m.M
	v196 = v194 + int32(1)
	if v196 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v204 = v189 + int32(1)
	if v204 != 0 {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	base.MemoryCopy(m, v197+v193, v190, v196)
	goto L36
L35:
	;
	goto L36
L36:
	;
	F_pfree(m, v190)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v208 = v193
	v209 = v196
	goto L29
L38:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	base.MemoryCopy(m, v205+v202, v190, v204)
	goto L40
L39:
	;
	goto L40
L40:
	;
	v208 = v202
	v209 = v204
	goto L29
L41:
	;
	goto L12
L42:
	;
	m.G0 = v19 + int32(16)
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
						if v14&int32(_a_F_row_is_in_frame_0) != 0 {
							v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
							if l1 != v96 {
								v120 = int32(1)
							} else {
								v120 = v4
							}
							m.G0 = v12 + int32(16)
							return v120
						} else {
							if v14&int32(_a_F_row_is_in_frame_1) == int32(0) {
								if v14&int32(_a_F_row_is_in_frame_2) == int32(0) {
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
						if v14&int32(_a_F_row_is_in_frame_0) != 0 {
							v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
							if l1 != v96 {
								v120 = int32(1)
							} else {
								v120 = v4
							}
							m.G0 = v12 + int32(16)
							return v120
						} else {
							if v14&int32(_a_F_row_is_in_frame_1) == int32(0) {
								if v14&int32(_a_F_row_is_in_frame_2) == int32(0) {
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
							if v14&int32(_a_F_row_is_in_frame_0) != 0 {
								v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
								if l1 != v96 {
									v120 = int32(1)
								} else {
									v120 = v4
								}
								m.G0 = v12 + int32(16)
								return v120
							} else {
								if v14&int32(_a_F_row_is_in_frame_1) == int32(0) {
									if v14&int32(_a_F_row_is_in_frame_2) == int32(0) {
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
								if v14&int32(_a_F_row_is_in_frame_0) != 0 {
									v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
									if l1 != v96 {
										v120 = int32(1)
									} else {
										v120 = v4
									}
									m.G0 = v12 + int32(16)
									return v120
								} else {
									if v14&int32(_a_F_row_is_in_frame_1) == int32(0) {
										if v14&int32(_a_F_row_is_in_frame_2) == int32(0) {
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
										if v14&int32(_a_F_row_is_in_frame_0) != 0 {
											v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
											if l1 != v96 {
												v120 = int32(1)
											} else {
												v120 = v4
											}
											m.G0 = v12 + int32(16)
											return v120
										} else {
											if v14&int32(_a_F_row_is_in_frame_1) == int32(0) {
												if v14&int32(_a_F_row_is_in_frame_2) == int32(0) {
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
									v48 = int32(_a_F_row_is_in_frame_3)
									v49 = *(*int32)(unsafe.Add(mBase, _c_F_row_is_in_frame[0]))
									v51 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
									*(*int32)(unsafe.Add(mBase, _c_F_row_is_in_frame[0])) = v51
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
									v56 = m.T0[v55].(func(*base.Module, int32, int32, int32) int32)(m, v42, v38, v12+int32(15))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_row_is_in_frame[0])) = v49
										v60 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
										F_MemoryContextReset(m, v60)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											if v56 != 0 {
												if v14&int32(_a_F_row_is_in_frame_0) != 0 {
													v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
													if l1 != v96 {
														v120 = int32(1)
													} else {
														v120 = v4
													}
													m.G0 = v12 + int32(16)
													return v120
												} else {
													if v14&int32(_a_F_row_is_in_frame_1) == int32(0) {
														if v14&int32(_a_F_row_is_in_frame_2) == int32(0) {
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
				if v14&int32(_a_F_row_is_in_frame_4) == int32(0) {
					if v14&int32(_a_F_row_is_in_frame_0) != 0 {
						v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
						if l1 != v96 {
							v120 = int32(1)
						} else {
							v120 = v4
						}
						m.G0 = v12 + int32(16)
						return v120
					} else {
						if v14&int32(_a_F_row_is_in_frame_1) == int32(0) {
							if v14&int32(_a_F_row_is_in_frame_2) == int32(0) {
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
						if v14&int32(_a_F_row_is_in_frame_5) != 0 {
							v77 = int64(0) - v73
						} else {
							v77 = v73
						}
						if l1 <= v70+v77 {
							if v14&int32(_a_F_row_is_in_frame_0) != 0 {
								v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
								if l1 != v96 {
									v120 = int32(1)
								} else {
									v120 = v4
								}
								m.G0 = v12 + int32(16)
								return v120
							} else {
								if v14&int32(_a_F_row_is_in_frame_1) == int32(0) {
									if v14&int32(_a_F_row_is_in_frame_2) == int32(0) {
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
							if v14&int32(_a_F_row_is_in_frame_0) != 0 {
								v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
								if l1 != v96 {
									v120 = int32(1)
								} else {
									v120 = v4
								}
								m.G0 = v12 + int32(16)
								return v120
							} else {
								if v14&int32(_a_F_row_is_in_frame_1) == int32(0) {
									if v14&int32(_a_F_row_is_in_frame_2) == int32(0) {
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
									if v14&int32(_a_F_row_is_in_frame_0) != 0 {
										v96 = *(*int64)(unsafe.Add(mBase, uint32(l0)+176))
										if l1 != v96 {
											v120 = int32(1)
										} else {
											v120 = v4
										}
										m.G0 = v12 + int32(16)
										return v120
									} else {
										if v14&int32(_a_F_row_is_in_frame_1) == int32(0) {
											if v14&int32(_a_F_row_is_in_frame_2) == int32(0) {
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if base.B2i32(v18 == int32(0))|base.B2i32(v18 != v21) != 0 {
		v39 = v18
		v40 = v21
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return v39 - v40
L8:
	;
	goto L7
L9:
	;
	v24 = v6
	v25 = v4
	goto L10
L10:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v29 == int32(0) {
		v39 = v29
		v40 = v28
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v39 = v29
	v40 = v28
	goto L8
L12:
	;
	v32 = int32(1)
	if v29 == v28 {
		v24 = v24 + v32
		v25 = v25 + v32
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
}
