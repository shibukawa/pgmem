package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TransitionTableAddTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l5 != 0 {
		if l4 != 0 {
			v64 = l4
			F_tuplestore_puttupleslot(m, l5, v64)
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				m.G0 = v10 + int32(16)
				return
			}
		} else {
			v12 = F_ExecGetChildToRootMap(m, l2)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				if v12 == int32(0) {
					v64 = l3
					F_tuplestore_puttupleslot(m, l5, v64)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						m.G0 = v10 + int32(16)
						return
					}
				} else {
					switch l0 - int32(1) {
					case 0:
						v35 = int32(16)
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l1+v35)))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
						if v38 == int32(0) {
							v41 = int32(_a_F_TransitionTableAddTuple_0)
							v42 = *(*int32)(unsafe.Add(mBase, _c_F_TransitionTableAddTuple[0]))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							v46 = *(*int32)(unsafe.Add(mBase, _c_F_TransitionTableAddTuple[1]))
							*(*int32)(unsafe.Add(mBase, _c_F_TransitionTableAddTuple[0])) = v46
							v48 = F_CreateTupleDescCopy(m, v43)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								v51 = F_MakeTupleTableSlot(m, v48, int32(_a_F_TransitionTableAddTuple_1))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v51
									*(*int32)(unsafe.Add(mBase, _c_F_TransitionTableAddTuple[0])) = v42
									v57 = v51
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									v59 = F_execute_attr_map_slot(m, v58, l3, v57)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										v64 = v57
										F_tuplestore_puttupleslot(m, l5, v64)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											m.G0 = v10 + int32(16)
											return
										}
									}
								}
							}
						} else {
							v57 = v38
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v59 = F_execute_attr_map_slot(m, v58, l3, v57)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								v64 = v57
								F_tuplestore_puttupleslot(m, l5, v64)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							}
						}
					case 1:
						v35 = int32(12)
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l1+v35)))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
						if v38 == int32(0) {
							v41 = int32(_a_F_TransitionTableAddTuple_0)
							v42 = *(*int32)(unsafe.Add(mBase, _c_F_TransitionTableAddTuple[0]))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							v46 = *(*int32)(unsafe.Add(mBase, _c_F_TransitionTableAddTuple[1]))
							*(*int32)(unsafe.Add(mBase, _c_F_TransitionTableAddTuple[0])) = v46
							v48 = F_CreateTupleDescCopy(m, v43)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								v51 = F_MakeTupleTableSlot(m, v48, int32(_a_F_TransitionTableAddTuple_1))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v51
									*(*int32)(unsafe.Add(mBase, _c_F_TransitionTableAddTuple[0])) = v42
									v57 = v51
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									v59 = F_execute_attr_map_slot(m, v58, l3, v57)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										v64 = v57
										F_tuplestore_puttupleslot(m, l5, v64)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											m.G0 = v10 + int32(16)
											return
										}
									}
								}
							}
						} else {
							v57 = v38
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v59 = F_execute_attr_map_slot(m, v58, l3, v57)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								v64 = v57
								F_tuplestore_puttupleslot(m, l5, v64)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							}
						}
					case 2:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(3)
							F_errmsg_internal(m, int32(_a_F_TransitionTableAddTuple_2), v10)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_TransitionTableAddTuple_3), int32(_a_F_TransitionTableAddTuple_4), int32(_a_F_TransitionTableAddTuple_5))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					default:
						v35 = int32(8)
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l1+v35)))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
						if v38 == int32(0) {
							v41 = int32(_a_F_TransitionTableAddTuple_0)
							v42 = *(*int32)(unsafe.Add(mBase, _c_F_TransitionTableAddTuple[0]))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							v46 = *(*int32)(unsafe.Add(mBase, _c_F_TransitionTableAddTuple[1]))
							*(*int32)(unsafe.Add(mBase, _c_F_TransitionTableAddTuple[0])) = v46
							v48 = F_CreateTupleDescCopy(m, v43)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								v51 = F_MakeTupleTableSlot(m, v48, int32(_a_F_TransitionTableAddTuple_1))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v51
									*(*int32)(unsafe.Add(mBase, _c_F_TransitionTableAddTuple[0])) = v42
									v57 = v51
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									v59 = F_execute_attr_map_slot(m, v58, l3, v57)
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
										return
									} else {
										v64 = v57
										F_tuplestore_puttupleslot(m, l5, v64)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											m.G0 = v10 + int32(16)
											return
										}
									}
								}
							}
						} else {
							v57 = v38
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v59 = F_execute_attr_map_slot(m, v58, l3, v57)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								v64 = v57
								F_tuplestore_puttupleslot(m, l5, v64)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v10 + int32(16)
		return
	}
}
func F___trunctfdf2(m *base.Module, l0 int64, l1 int64) float64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v29 int64
	_ = v29
	var v34 int64
	_ = v34
	var v44 int64
	_ = v44
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int64
	_ = v79
	var v83 int64
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v100 int64
	_ = v100
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v118 int32
	_ = v118
	var v133 int64
	_ = v133
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v142 int64
	_ = v142
	var v145 int64
	_ = v145
	var v148 int64
	_ = v148
	var v152 int64
	_ = v152
	var v162 int64
	_ = v162
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v176 int64
	_ = v176
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = l1 & int64(281474976710655)
	v19 = int64(base.Ui64(l1)>>(uint(int64(48))%64)) & int64(32767)
	v20 = base.I32_wrap_i64(v19)
	if base.Ui32(v20-int32(_a_F___trunctfdf2_0)) <= base.Ui32(int32(2045)) {
		v29 = v15<<(uint(int64(4))%64) | int64(base.Ui64(l0)>>(uint(int64(60))%64))
		v34 = l0 & int64(1152921504606846975)
		if base.Ui64(int64(576460752303423489)) <= base.Ui64(v34) {
			v44 = v29 + int64(1)
		} else {
			if v34 != int64(576460752303423488) {
				v44 = v29
			} else {
				v44 = v29&int64(1) + v29
			}
		}
		v47 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v44))
		if base.Ui64(int64(4503599627370495)) < base.Ui64(v44) {
			v48 = int64(0)
		} else {
			v48 = v44
		}
		v169 = v48
		v176 = base.I64_extend_i32_u(v47) + base.I64_extend_i32_u(v20-int32(_a_F___trunctfdf2_1))
	} else {
		if base.B2i32(l0|v15 == int64(0))|base.B2i32(v19 != int64(32767)) == int32(0) {
			v169 = v15<<(uint(int64(4))%64) | int64(base.Ui64(l0)>>(uint(int64(60))%64)) | int64(2251799813685248)
			v176 = int64(2047)
		} else {
			if base.Ui32(int32(_a_F___trunctfdf2_2)) < base.Ui32(v20) {
				v169 = int64(0)
				v176 = int64(2047)
			} else {
				v74 = base.B2i32(v19 == int64(0))
				if v19 == int64(0) {
					v75 = int32(_a_F___trunctfdf2_1)
				} else {
					v75 = int32(_a_F___trunctfdf2_0)
				}
				v76 = v75 - v20
				if int32(112) < v76 {
					v79 = int64(0)
					v169 = v79
					v176 = v79
				} else {
					if v19 == int64(0) {
						v83 = v15
					} else {
						v83 = v15 | int64(281474976710656)
					}
					if v20 != v75 {
						v87 = v12 + int32(16)
						v89 = int32(128) - v76
						if v89&int32(64) != 0 {
							v108 = int64(0)
							v109 = l0 << (uint(base.I64_extend_i32_u(v89+int32(-64))) % 64)
						} else {
							if v89 == int32(0) {
								v108 = l0
								v109 = v83
							} else {
								v100 = base.I64_extend_i32_u(v89)
								v108 = l0 << (uint(v100) % 64)
								v109 = v83<<(uint(v100)%64) | int64(base.Ui64(l0)>>(uint(base.I64_extend_i32_u(int32(64)-v89))%64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v87))) = v108
						*(*int64)(unsafe.Add(mBase, uint32(v87)+8)) = v109
						v113 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
						v114 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
						v118 = base.B2i32(v113|v114 != int64(0))
					} else {
						v118 = int32(0)
					}
					if v76&int32(64) != 0 {
						v137 = int64(base.Ui64(v83) >> (uint(base.I64_extend_i32_u(v76+int32(-64))) % 64))
						v138 = int64(0)
					} else {
						if v76 == int32(0) {
							v137 = l0
							v138 = v83
						} else {
							v133 = base.I64_extend_i32_u(v76)
							v137 = v83<<(uint(base.I64_extend_i32_u(int32(64)-v76))%64) | int64(base.Ui64(l0)>>(uint(v133)%64))
							v138 = int64(base.Ui64(v83) >> (uint(v133) % 64))
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v12))) = v137
					*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v138
					v142 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
					v145 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
					v148 = v142<<(uint(int64(4))%64) | int64(base.Ui64(v145)>>(uint(int64(60))%64))
					v152 = base.I64_extend_i32_u(v118) | v145&int64(1152921504606846975)
					if base.Ui64(int64(576460752303423489)) <= base.Ui64(v152) {
						v162 = v148 + int64(1)
					} else {
						if v152 != int64(576460752303423488) {
							v162 = v148
						} else {
							v162 = v148&int64(1) + v148
						}
					}
					v166 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v162))
					if base.Ui64(int64(4503599627370495)) < base.Ui64(v162) {
						v167 = v162 ^ int64(4503599627370496)
					} else {
						v167 = v162
					}
					v169 = v167
					v176 = base.I64_extend_i32_u(v166)
				}
			}
		}
	}
	m.G0 = v12 + int32(32)
	return base.F64_reinterpret_i64(l1&int64(-9223372036854775807-1) | v176<<(uint(int64(52))%64) | v169)
}
func F__tarWriteHeader(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v523 int32
	_ = v523
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v543 int64
	_ = v543
	var v546 int64
	_ = v546
	var v549 int64
	_ = v549
	var v552 int64
	_ = v552
	var v555 int64
	_ = v555
	var v558 int64
	_ = v558
	var v561 int64
	_ = v561
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v613 int32
	_ = v613
	var v620 int32
	_ = v620
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v641 int32
	_ = v641
	var v647 int64
	_ = v647
	var v650 int64
	_ = v650
	var v653 int64
	_ = v653
	var v656 int64
	_ = v656
	var v659 int64
	_ = v659
	var v662 int64
	_ = v662
	var v665 int64
	_ = v665
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v819 int64
	_ = v819
	var v822 int32
	_ = v822
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v877 int32
	_ = v877
	var v884 int32
	_ = v884
	var v891 int32
	_ = v891
	var v898 int32
	_ = v898
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v923 int64
	_ = v923
	var v925 int64
	_ = v925
	var v928 int64
	_ = v928
	var v942 int32
	_ = v942
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v957 int32
	_ = v957
	var v962 int32
	_ = v962
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v976 int32
	_ = v976
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	if l4 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l3)+24))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l3)+56))
	v25 = F_strlen(m, l1)
	mBase = m.M
	if base.Ui32(int32(99)) < base.Ui32(v25) {
		v942 = int32(1)
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v14 + int32(48)
	return
L4:
	;
	if v942 != 0 {
		goto L147
	} else {
		goto L148
	}
L5:
	;
	if l2 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v289 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+107)) = uint8(v289)
	v291 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+102)) = uint8(v291)
	v293 = int32(_a_F__tarWriteHeader_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+100)) = uint16(v293)
	v295 = int32(7)
	v298 = v20&v295 | v291
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+106)) = uint8(v298)
	v305 = int32(base.Ui32(v20)>>(uint(int32(3))%32))&v295 | v291
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+105)) = uint8(v305)
	v312 = int32(base.Ui32(v20)>>(uint(int32(6))%32))&v295 | v291
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+104)) = uint8(v312)
	v319 = int32(base.Ui32(v20)>>(uint(int32(9))%32))&v295 | v291
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+103)) = uint8(v319)
	if base.Ui32(v21) <= base.Ui32(int32(_a_F__tarWriteHeader_1)) {
		goto L79
	} else {
		goto L80
	}
L7:
	;
	v280 = int32(99)
	v281 = F_strlen(m, l1)
	mBase = m.M
	if v280 <= v281 {
		goto L75
	} else {
		goto L76
	}
L8:
	;
	v29 = F_strlen(m, l2)
	mBase = m.M
	if base.Ui32(int32(99)) < base.Ui32(v29) {
		v942 = int32(2)
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	base.MemoryFill(m, v18, int32(0), int32(512))
	goto L46
L11:
	;
	base.MemoryFill(m, v18, int32(0), int32(512))
	goto L15
L12:
	;
	goto L7
L13:
	;
	v151 = F_strlen(m, v140)
	mBase = m.M
	goto L12
L15:
	;
	goto L16
L16:
	;
	v41 = int32(99)
	if (v18^l1)&int32(3) != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v144 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v144)
	goto L13
L18:
	;
	v125 = v120
	v126 = v121
	v127 = v122
	goto L39
L19:
	;
	if v115 == int32(0) {
		v140 = v113
		v141 = v114
		goto L17
	} else {
		goto L38
	}
L20:
	;
	v113 = l1
	v114 = v18
	v115 = v41
	goto L19
L21:
	;
	goto L22
L22:
	;
	v45 = int32(0)
	if base.B2i32(l1&int32(3) == v45)|int32(0) == v45 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v81 == int32(0) {
		v140 = v78
		v141 = v79
		goto L17
	} else {
		goto L32
	}
L24:
	;
	v57 = l1
	v58 = v18
	v59 = v41
	goto L27
L25:
	;
	goto L26
L26:
	;
	v78 = l1
	v79 = v18
	v80 = v41
	v81 = int32(1)
	goto L23
L27:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	*(*uint8)(unsafe.Add(mBase, uint32(v58))) = uint8(v61)
	if v61 == int32(0) {
		v120 = v57
		v121 = v58
		v122 = v59
		goto L18
	} else {
		goto L29
	}
L28:
	;
	v78 = v72
	v79 = v66
	v80 = v68
	v81 = v70
	goto L23
L29:
	;
	v65 = int32(1)
	v66 = v58 + v65
	v68 = v59 - v65
	v69 = int32(0)
	v70 = base.B2i32(v68 != v69)
	v72 = v57 + v65
	if v72&int32(3) == v69 {
		v78 = v72
		v79 = v66
		v80 = v68
		v81 = v70
		goto L23
	} else {
		goto L30
	}
L30:
	;
	if v68 != 0 {
		v57 = v72
		v58 = v66
		v59 = v68
		goto L27
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if base.B2i32(v84 == int32(0))|base.B2i32(base.Ui32(v80) < base.Ui32(int32(4))) != 0 {
		v113 = v78
		v114 = v79
		v115 = v80
		goto L19
	} else {
		goto L33
	}
L33:
	;
	v91 = v78
	v92 = v79
	v93 = v80
	goto L34
L34:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v99 = int32(-2139062144)
	if (int32(16843008)-v96|v96)&v99 != v99 {
		v120 = v91
		v121 = v92
		v122 = v93
		goto L18
	} else {
		goto L36
	}
L35:
	;
	v113 = v107
	v114 = v105
	v115 = v109
	goto L19
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = v96
	v104 = int32(4)
	v105 = v92 + v104
	v107 = v91 + v104
	v109 = v93 - v104
	if base.Ui32(int32(3)) < base.Ui32(v109) {
		v91 = v107
		v92 = v105
		v93 = v109
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v120 = v113
	v121 = v114
	v122 = v115
	goto L18
L39:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v129)
	if v129 == int32(0) {
		v140 = v125
		v141 = v126
		goto L17
	} else {
		goto L41
	}
L40:
	;
	v140 = v136
	v141 = v134
	goto L17
L41:
	;
	v133 = int32(1)
	v134 = v126 + v133
	v136 = v125 + v133
	v138 = v127 - v133
	if v138 != 0 {
		v125 = v136
		v126 = v134
		v127 = v138
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	if v20&int32(_a_F__tarWriteHeader_2) != int32(_a_F__tarWriteHeader_3) {
		goto L6
	} else {
		goto L74
	}
L44:
	;
	v273 = F_strlen(m, v262)
	mBase = m.M
	goto L43
L46:
	;
	goto L47
L47:
	;
	v163 = int32(99)
	if (v18^l1)&int32(3) != 0 {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v266 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v263))) = uint8(v266)
	goto L44
L49:
	;
	v247 = v242
	v248 = v243
	v249 = v244
	goto L70
L50:
	;
	if v237 == int32(0) {
		v262 = v235
		v263 = v236
		goto L48
	} else {
		goto L69
	}
L51:
	;
	v235 = l1
	v236 = v18
	v237 = v163
	goto L50
L52:
	;
	goto L53
L53:
	;
	v167 = int32(0)
	if base.B2i32(l1&int32(3) == v167)|int32(0) == v167 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v203 == int32(0) {
		v262 = v200
		v263 = v201
		goto L48
	} else {
		goto L63
	}
L55:
	;
	v179 = l1
	v180 = v18
	v181 = v163
	goto L58
L56:
	;
	goto L57
L57:
	;
	v200 = l1
	v201 = v18
	v202 = v163
	v203 = int32(1)
	goto L54
L58:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v183)
	if v183 == int32(0) {
		v242 = v179
		v243 = v180
		v244 = v181
		goto L49
	} else {
		goto L60
	}
L59:
	;
	v200 = v194
	v201 = v188
	v202 = v190
	v203 = v192
	goto L54
L60:
	;
	v187 = int32(1)
	v188 = v180 + v187
	v190 = v181 - v187
	v191 = int32(0)
	v192 = base.B2i32(v190 != v191)
	v194 = v179 + v187
	if v194&int32(3) == v191 {
		v200 = v194
		v201 = v188
		v202 = v190
		v203 = v192
		goto L54
	} else {
		goto L61
	}
L61:
	;
	if v190 != 0 {
		v179 = v194
		v180 = v188
		v181 = v190
		goto L58
	} else {
		goto L62
	}
L62:
	;
	goto L59
L63:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
	if base.B2i32(v206 == int32(0))|base.B2i32(base.Ui32(v202) < base.Ui32(int32(4))) != 0 {
		v235 = v200
		v236 = v201
		v237 = v202
		goto L50
	} else {
		goto L64
	}
L64:
	;
	v213 = v200
	v214 = v201
	v215 = v202
	goto L65
L65:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v221 = int32(-2139062144)
	if (int32(16843008)-v218|v218)&v221 != v221 {
		v242 = v213
		v243 = v214
		v244 = v215
		goto L49
	} else {
		goto L67
	}
L66:
	;
	v235 = v229
	v236 = v227
	v237 = v231
	goto L50
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = v218
	v226 = int32(4)
	v227 = v214 + v226
	v229 = v213 + v226
	v231 = v215 - v226
	if base.Ui32(int32(3)) < base.Ui32(v231) {
		v213 = v229
		v214 = v227
		v215 = v231
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v242 = v235
	v243 = v236
	v244 = v237
	goto L49
L70:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	*(*uint8)(unsafe.Add(mBase, uint32(v248))) = uint8(v251)
	if v251 == int32(0) {
		v262 = v247
		v263 = v248
		goto L48
	} else {
		goto L72
	}
L71:
	;
	v262 = v258
	v263 = v256
	goto L48
L72:
	;
	v255 = int32(1)
	v256 = v248 + v255
	v258 = v247 + v255
	v260 = v249 - v255
	if v260 != 0 {
		v247 = v258
		v248 = v256
		v249 = v260
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	goto L7
L75:
	;
	v284 = v280
	goto L77
L76:
	;
	v284 = v281
	goto L77
L77:
	;
	v286 = int32(47)
	*(*uint16)(unsafe.Add(mBase, uint32(v18+v284))) = uint16(v286)
	goto L6
L78:
	;
	if base.Ui32(v22) <= base.Ui32(int32(_a_F__tarWriteHeader_1)) {
		goto L83
	} else {
		goto L84
	}
L79:
	;
	v323 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+115)) = uint8(v323)
	v325 = int32(7)
	v327 = int32(48)
	v328 = v21&v325 | v327
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+114)) = uint8(v328)
	v333 = int32(base.Ui32(v21)>>(uint(int32(18))%32)) | v327
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+108)) = uint8(v333)
	v340 = int32(base.Ui32(v21)>>(uint(int32(3))%32))&v325 | v327
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+113)) = uint8(v340)
	v347 = int32(base.Ui32(v21)>>(uint(int32(6))%32))&v325 | v327
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+112)) = uint8(v347)
	v354 = int32(base.Ui32(v21)>>(uint(int32(9))%32))&v325 | v327
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+111)) = uint8(v354)
	v361 = int32(base.Ui32(v21)>>(uint(int32(12))%32))&v325 | v327
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+110)) = uint8(v361)
	v368 = int32(base.Ui32(v21)>>(uint(int32(15))%32))&v325 | v327
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+109)) = uint8(v368)
	goto L78
L80:
	;
	goto L81
L81:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+115)) = uint8(v21)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = int32(128)
	v374 = int32(base.Ui32(v21) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+114)) = uint8(v374)
	v377 = int32(base.Ui32(v21) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+113)) = uint8(v377)
	v380 = int32(base.Ui32(v21) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+112)) = uint8(v380)
	goto L78
L82:
	;
	v443 = int32(0)
	v446 = v20 & int32(_a_F__tarWriteHeader_2)
	if base.B2i32(l2 == v443)&base.B2i32(v446 != int32(_a_F__tarWriteHeader_3)) == v443 {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	v384 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+123)) = uint8(v384)
	v386 = int32(7)
	v388 = int32(48)
	v389 = v22&v386 | v388
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+122)) = uint8(v389)
	v394 = int32(base.Ui32(v22)>>(uint(int32(18))%32)) | v388
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+116)) = uint8(v394)
	v401 = int32(base.Ui32(v22)>>(uint(int32(3))%32))&v386 | v388
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+121)) = uint8(v401)
	v408 = int32(base.Ui32(v22)>>(uint(int32(6))%32))&v386 | v388
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+120)) = uint8(v408)
	v415 = int32(base.Ui32(v22)>>(uint(int32(9))%32))&v386 | v388
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+119)) = uint8(v415)
	v422 = int32(base.Ui32(v22)>>(uint(int32(12))%32))&v386 | v388
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+118)) = uint8(v422)
	v429 = int32(base.Ui32(v22)>>(uint(int32(15))%32))&v386 | v388
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+117)) = uint8(v429)
	goto L82
L84:
	;
	goto L85
L85:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+123)) = uint8(v22)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = int32(128)
	v435 = int32(base.Ui32(v22) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+122)) = uint8(v435)
	v438 = int32(base.Ui32(v22) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+121)) = uint8(v438)
	v441 = int32(base.Ui32(v22) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+120)) = uint8(v441)
	goto L82
L86:
	;
	if base.Ui64(v23) <= base.Ui64(int64(8589934591)) {
		goto L94
	} else {
		goto L95
	}
L87:
	;
	v452 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+135)) = uint8(v452)
	v455 = v18 + int32(124)
	*(*int32)(unsafe.Add(mBase, uint32(v455)+7)) = int32(808464432)
	*(*int64)(unsafe.Add(mBase, uint32(v455))) = int64(3472328296227680304)
	goto L86
L88:
	;
	goto L89
L89:
	;
	if base.Ui64(v19) <= base.Ui64(int64(8589934591)) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v462 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+135)) = uint8(v462)
	v464 = base.I32_wrap_i64(v19)
	v465 = int32(7)
	v467 = int32(48)
	v468 = v464&v465 | v467
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+134)) = uint8(v468)
	v474 = base.I32_wrap_i64(int64(base.Ui64(v19)>>(uint(int64(30))%64))) | v467
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+124)) = uint8(v474)
	v481 = int32(base.Ui32(v464)>>(uint(int32(3))%32))&v465 | v467
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+133)) = uint8(v481)
	v488 = int32(base.Ui32(v464)>>(uint(int32(6))%32))&v465 | v467
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+132)) = uint8(v488)
	v495 = int32(base.Ui32(v464)>>(uint(int32(9))%32))&v465 | v467
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+131)) = uint8(v495)
	v502 = int32(base.Ui32(v464)>>(uint(int32(12))%32))&v465 | v467
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+130)) = uint8(v502)
	v509 = int32(base.Ui32(v464)>>(uint(int32(15))%32))&v465 | v467
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+129)) = uint8(v509)
	v516 = int32(base.Ui32(v464)>>(uint(int32(18))%32))&v465 | v467
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+128)) = uint8(v516)
	v523 = int32(base.Ui32(v464)>>(uint(int32(21))%32))&v465 | v467
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+127)) = uint8(v523)
	v530 = int32(base.Ui32(v464)>>(uint(int32(24))%32))&v465 | v467
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+126)) = uint8(v530)
	v537 = int32(base.Ui32(v464)>>(uint(int32(27))%32))&v465 | v467
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+125)) = uint8(v537)
	goto L86
L91:
	;
	goto L92
L92:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+135)) = uint8(v19)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = int32(128)
	v543 = int64(base.Ui64(v19) >> (uint(int64(8)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+134)) = uint8(v543)
	v546 = int64(base.Ui64(v19) >> (uint(int64(16)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+133)) = uint8(v546)
	v549 = int64(base.Ui64(v19) >> (uint(int64(24)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+132)) = uint8(v549)
	v552 = int64(base.Ui64(v19) >> (uint(int64(32)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+131)) = uint8(v552)
	v555 = int64(base.Ui64(v19) >> (uint(int64(40)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+130)) = uint8(v555)
	v558 = int64(base.Ui64(v19) >> (uint(int64(48)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+129)) = uint8(v558)
	v561 = int64(base.Ui64(v19) >> (uint(int64(56)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+128)) = uint8(v561)
	goto L86
L93:
	;
	if l2 != 0 {
		goto L98
	} else {
		goto L99
	}
L94:
	;
	v566 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+147)) = uint8(v566)
	v568 = base.I32_wrap_i64(v23)
	v569 = int32(7)
	v571 = int32(48)
	v572 = v568&v569 | v571
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+146)) = uint8(v572)
	v578 = base.I32_wrap_i64(int64(base.Ui64(v23)>>(uint(int64(30))%64))) | v571
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+136)) = uint8(v578)
	v585 = int32(base.Ui32(v568)>>(uint(int32(3))%32))&v569 | v571
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+145)) = uint8(v585)
	v592 = int32(base.Ui32(v568)>>(uint(int32(6))%32))&v569 | v571
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+144)) = uint8(v592)
	v599 = int32(base.Ui32(v568)>>(uint(int32(9))%32))&v569 | v571
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+143)) = uint8(v599)
	v606 = int32(base.Ui32(v568)>>(uint(int32(12))%32))&v569 | v571
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+142)) = uint8(v606)
	v613 = int32(base.Ui32(v568)>>(uint(int32(15))%32))&v569 | v571
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+141)) = uint8(v613)
	v620 = int32(base.Ui32(v568)>>(uint(int32(18))%32))&v569 | v571
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+140)) = uint8(v620)
	v627 = int32(base.Ui32(v568)>>(uint(int32(21))%32))&v569 | v571
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+139)) = uint8(v627)
	v634 = int32(base.Ui32(v568)>>(uint(int32(24))%32))&v569 | v571
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+138)) = uint8(v634)
	v641 = int32(base.Ui32(v568)>>(uint(int32(27))%32))&v569 | v571
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+137)) = uint8(v641)
	goto L93
L95:
	;
	goto L96
L96:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+147)) = uint8(v23)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+136)) = int32(128)
	v647 = int64(base.Ui64(v23) >> (uint(int64(8)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+146)) = uint8(v647)
	v650 = int64(base.Ui64(v23) >> (uint(int64(16)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+145)) = uint8(v650)
	v653 = int64(base.Ui64(v23) >> (uint(int64(24)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+144)) = uint8(v653)
	v656 = int64(base.Ui64(v23) >> (uint(int64(32)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+143)) = uint8(v656)
	v659 = int64(base.Ui64(v23) >> (uint(int64(40)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+142)) = uint8(v659)
	v662 = int64(base.Ui64(v23) >> (uint(int64(48)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+141)) = uint8(v662)
	v665 = int64(base.Ui64(v23) >> (uint(int64(56)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+140)) = uint8(v665)
	goto L93
L97:
	;
	v797 = int32(_a_F__tarWriteHeader_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+263)) = uint16(v797)
	v799 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+336)) = uint8(v799)
	v801 = int32(808464432)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+329)) = v801
	*(*int32)(unsafe.Add(mBase, uint32(v18)+332)) = v801
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+344)) = uint8(v799)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+340)) = v801
	*(*int32)(unsafe.Add(mBase, uint32(v18)+337)) = v801
	v813 = int32(*(*uint16)(unsafe.Add(mBase, _c_F__tarWriteHeader[0])))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+261)) = uint16(v813)
	v816 = *(*int32)(unsafe.Add(mBase, _c_F__tarWriteHeader[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+257)) = v816
	v819 = *(*int64)(unsafe.Add(mBase, _c_F__tarWriteHeader[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+265)) = v819
	v822 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__tarWriteHeader[3])))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+273)) = uint8(v822)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+297)) = v819
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+305)) = uint8(v822)
	v830 = int32(256)
	v832 = int32(0)
	goto L135
L98:
	;
	v668 = int32(50)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+156)) = uint8(v668)
	v671 = v18 + int32(157)
	goto L104
L99:
	;
	goto L100
L100:
	;
	if v446 == int32(_a_F__tarWriteHeader_3) {
		goto L132
	} else {
		goto L133
	}
L101:
	;
	goto L97
L102:
	;
	v788 = F_strlen(m, v777)
	mBase = m.M
	goto L101
L104:
	;
	goto L105
L105:
	;
	v678 = int32(99)
	if (v671^l2)&int32(3) != 0 {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	v781 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v778))) = uint8(v781)
	goto L102
L107:
	;
	v762 = v757
	v763 = v758
	v764 = v759
	goto L128
L108:
	;
	if v752 == int32(0) {
		v777 = v750
		v778 = v751
		goto L106
	} else {
		goto L127
	}
L109:
	;
	v750 = l2
	v751 = v671
	v752 = v678
	goto L108
L110:
	;
	goto L111
L111:
	;
	v682 = int32(0)
	if base.B2i32(l2&int32(3) == v682)|int32(0) == v682 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	if v718 == int32(0) {
		v777 = v715
		v778 = v716
		goto L106
	} else {
		goto L121
	}
L113:
	;
	v694 = l2
	v695 = v671
	v696 = v678
	goto L116
L114:
	;
	goto L115
L115:
	;
	v715 = l2
	v716 = v671
	v717 = v678
	v718 = int32(1)
	goto L112
L116:
	;
	v698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694))))
	*(*uint8)(unsafe.Add(mBase, uint32(v695))) = uint8(v698)
	if v698 == int32(0) {
		v757 = v694
		v758 = v695
		v759 = v696
		goto L107
	} else {
		goto L118
	}
L117:
	;
	v715 = v709
	v716 = v703
	v717 = v705
	v718 = v707
	goto L112
L118:
	;
	v702 = int32(1)
	v703 = v695 + v702
	v705 = v696 - v702
	v706 = int32(0)
	v707 = base.B2i32(v705 != v706)
	v709 = v694 + v702
	if v709&int32(3) == v706 {
		v715 = v709
		v716 = v703
		v717 = v705
		v718 = v707
		goto L112
	} else {
		goto L119
	}
L119:
	;
	if v705 != 0 {
		v694 = v709
		v695 = v703
		v696 = v705
		goto L116
	} else {
		goto L120
	}
L120:
	;
	goto L117
L121:
	;
	v721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v715))))
	if base.B2i32(v721 == int32(0))|base.B2i32(base.Ui32(v717) < base.Ui32(int32(4))) != 0 {
		v750 = v715
		v751 = v716
		v752 = v717
		goto L108
	} else {
		goto L122
	}
L122:
	;
	v728 = v715
	v729 = v716
	v730 = v717
	goto L123
L123:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v728)))
	v736 = int32(-2139062144)
	if (int32(16843008)-v733|v733)&v736 != v736 {
		v757 = v728
		v758 = v729
		v759 = v730
		goto L107
	} else {
		goto L125
	}
L124:
	;
	v750 = v744
	v751 = v742
	v752 = v746
	goto L108
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v729))) = v733
	v741 = int32(4)
	v742 = v729 + v741
	v744 = v728 + v741
	v746 = v730 - v741
	if base.Ui32(int32(3)) < base.Ui32(v746) {
		v728 = v744
		v729 = v742
		v730 = v746
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	v757 = v750
	v758 = v751
	v759 = v752
	goto L107
L128:
	;
	v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762))))
	*(*uint8)(unsafe.Add(mBase, uint32(v763))) = uint8(v766)
	if v766 == int32(0) {
		v777 = v762
		v778 = v763
		goto L106
	} else {
		goto L130
	}
L129:
	;
	v777 = v773
	v778 = v771
	goto L106
L130:
	;
	v770 = int32(1)
	v771 = v763 + v770
	v773 = v762 + v770
	v775 = v764 - v770
	if v775 != 0 {
		v762 = v773
		v763 = v771
		v764 = v775
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v793 = int32(53)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+156)) = uint8(v793)
	goto L97
L133:
	;
	goto L134
L134:
	;
	v795 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+156)) = uint8(v795)
	goto L97
L135:
	;
	if base.Ui32(v832-int32(156)) <= base.Ui32(int32(-9)) {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	if base.Ui32(v853) <= base.Ui32(int32(_a_F__tarWriteHeader_1)) {
		goto L144
	} else {
		goto L145
	}
L137:
	;
	v843 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v832))))
	v845 = v830 + v843
	goto L139
L138:
	;
	v845 = v830
	goto L139
L139:
	;
	if base.Ui32(v832-int32(155)) <= base.Ui32(int32(-9)) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v832)+1)))
	v853 = v845 + v851
	goto L142
L141:
	;
	v853 = v845
	goto L142
L142:
	;
	v855 = v832 + int32(2)
	if v855 != int32(512) {
		v830 = v853
		v832 = v855
		goto L135
	} else {
		goto L143
	}
L143:
	;
	goto L136
L144:
	;
	v860 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+155)) = uint8(v860)
	v862 = int32(7)
	v864 = int32(48)
	v865 = v853&v862 | v864
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+154)) = uint8(v865)
	v870 = int32(base.Ui32(v853)>>(uint(int32(18))%32)) | v864
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+148)) = uint8(v870)
	v877 = int32(base.Ui32(v853)>>(uint(int32(3))%32))&v862 | v864
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+153)) = uint8(v877)
	v884 = int32(base.Ui32(v853)>>(uint(int32(6))%32))&v862 | v864
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+152)) = uint8(v884)
	v891 = int32(base.Ui32(v853)>>(uint(int32(9))%32))&v862 | v864
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+151)) = uint8(v891)
	v898 = int32(base.Ui32(v853)>>(uint(int32(12))%32))&v862 | v864
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+150)) = uint8(v898)
	v905 = int32(base.Ui32(v853)>>(uint(int32(15))%32))&v862 | v864
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+149)) = uint8(v905)
	v942 = int32(0)
	goto L4
L145:
	;
	goto L146
L146:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+155)) = uint8(v853)
	v909 = int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+148)) = uint8(v909)
	v912 = int32(base.Ui32(v853) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+154)) = uint8(v912)
	v915 = int32(base.Ui32(v853) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+153)) = uint8(v915)
	v918 = int32(base.Ui32(v853) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+152)) = uint8(v918)
	v921 = v853 >> (uint(int32(31)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+151)) = uint8(v921)
	v923 = base.I64_extend_i32_s(v853)
	v925 = int64(base.Ui64(v923) >> (uint(int64(40)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+150)) = uint8(v925)
	v928 = int64(base.Ui64(v923) >> (uint(int64(48)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+149)) = uint8(v928)
	v942 = int32(0)
	goto L4
L147:
	;
	switch v942 - int32(1) {
	case 0:
		goto L152
	case 1:
		goto L151
	default:
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v996)+8))
	m.T0[v997].(func(*base.Module, int32, int32))(m, l0, int32(512))
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L153
	} else {
		goto L165
	}
L150:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L153
	} else {
		goto L162
	}
L151:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L153
	} else {
		goto L158
	}
L152:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	return
L154:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l1
	F_errmsg(m, int32(_a_F__tarWriteHeader_4), v14+int32(16))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L153
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F__tarWriteHeader_5), int32(2053), int32(_a_F__tarWriteHeader_6))
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L153
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L153
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l1
	F_errmsg(m, int32(_a_F__tarWriteHeader_7), v14+int32(32))
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L153
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F__tarWriteHeader_5), int32(2060), int32(_a_F__tarWriteHeader_6))
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L153
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v942
	F_errmsg_internal(m, int32(_a_F__tarWriteHeader_8), v14)
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L153
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F__tarWriteHeader_5), int32(2063), int32(_a_F__tarWriteHeader_6))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L153
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	goto L3
}
func F_tamil_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v655 int32
	_ = v655
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1084 int32
	_ = v1084
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1174 int32
	_ = v1174
	var v1179 int32
	_ = v1179
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1200 int32
	_ = v1200
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1259 int32
	_ = v1259
	var v1263 int32
	_ = v1263
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1288 int32
	_ = v1288
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1311 int32
	_ = v1311
	var v1316 int32
	_ = v1316
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1465 int32
	_ = v1465
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1498 int32
	_ = v1498
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
	var v1521 int32
	_ = v1521
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1532 int32
	_ = v1532
	var v1536 int32
	_ = v1536
	var v1540 int32
	_ = v1540
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1558 int32
	_ = v1558
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1567 int32
	_ = v1567
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1598 int32
	_ = v1598
	var v1601 int32
	_ = v1601
	var v1605 int32
	_ = v1605
	var v1615 int32
	_ = v1615
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = F_r_fix_ending(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v1615
L2:
	;
	return int32(0)
L3:
	;
	if v12 < int32(0) {
		v1615 = v12
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v11
	v19 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v20-int32(4))))
	if v28 == v19 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v102 < int32(5) {
		v1615 = v19
		goto L1
	} else {
		goto L21
	}
L6:
	;
	v102 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v33 = v28 & int32(3)
	if base.Ui32(v28) < base.Ui32(int32(4)) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v102 = v91
	goto L5
L10:
	;
	v75 = v69
	v76 = v70
	v80 = v19
	goto L18
L11:
	;
	v69 = v20
	v70 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v40 = v20
	v41 = int32(0)
	v44 = v19
	goto L14
L14:
	;
	v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(v40))))
	v47 = int32(-65)
	v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(v40)+1)))
	v54 = int32(*(*int8)(unsafe.Add(mBase, uint32(v40)+2)))
	v58 = int32(*(*int8)(unsafe.Add(mBase, uint32(v40)+3)))
	v61 = v41 + base.B2i32(v47 < v46) + base.B2i32(v47 < v50) + base.B2i32(v47 < v54) + base.B2i32(v47 < v58)
	v62 = int32(4)
	v63 = v40 + v62
	v65 = v44 + v62
	if v65 != v28&int32(-4) {
		v40 = v63
		v41 = v61
		v44 = v65
		goto L14
	} else {
		goto L16
	}
L15:
	;
	if v33 == int32(0) {
		v91 = v61
		goto L9
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	v69 = v63
	v70 = v61
	goto L10
L18:
	;
	v81 = int32(*(*int8)(unsafe.Add(mBase, uint32(v75))))
	v84 = v76 + base.B2i32(int32(-65) < v81)
	v85 = int32(1)
	v88 = v80 + v85
	if v88 != v33 {
		v75 = v75 + v85
		v76 = v84
		v80 = v88
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v91 = v84
	goto L9
L20:
	;
	goto L19
L21:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v105
	v107 = int32(3)
	v109 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v111-v105 < v107 {
		v121 = v109
		goto L24
	} else {
		goto L25
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
	v161 = v105 + int32(2)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v162 <= v161 {
		goto L39
	} else {
		goto L40
	}
L23:
	;
	if v121 == int32(0) {
		goto L22
	} else {
		goto L27
	}
L24:
	;
	goto L23
L25:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v117 = F_memcmp(m, v115+v105, int32(_a_F_tamil_UTF_8_stem_0), v107)
	mBase = m.M
	if v117 != 0 {
		v121 = v109
		goto L24
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v107 + v105
	v121 = int32(1)
	goto L24
L27:
	;
	v126 = F_find_among(m, l0, int32(_a_F_tamil_UTF_8_stem_1), int32(10))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	if v126 == int32(0) {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v130 = int32(3)
	v132 = int32(0)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v134-v135 < v130 {
		v144 = v132
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v144 == int32(0) {
		goto L22
	} else {
		goto L34
	}
L31:
	;
	goto L30
L32:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v140 = F_memcmp(m, v138+v135, int32(_a_F_tamil_UTF_8_stem_2), v130)
	mBase = m.M
	if v140 != 0 {
		v144 = v132
		goto L31
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v130 + v135
	v144 = int32(1)
	goto L31
L34:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v147
	v149 = F_slice_del(m, l0)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	if v149 < int32(0) {
		v1615 = v149
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v153 = F_r_fix_va_start(m, l0)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	if v153 < int32(0) {
		v1615 = v153
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L22
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
	v220 = int32(0)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v221-int32(4))))
	if v229 == v220 {
		goto L57
	} else {
		goto L58
	}
L40:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164+v161))))
	if base.B2i32(v166&int32(224) != int32(128))|base.B2i32(int32(1)<<(uint(v166)%32)&int32(672) == int32(0)) != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v180 = F_find_among(m, l0, int32(_a_F_tamil_UTF_8_stem_3), int32(3))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	if v180 == int32(0) {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v186 = F_find_among(m, l0, int32(_a_F_tamil_UTF_8_stem_4), int32(10))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	if v186 == int32(0) {
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v190 = int32(3)
	v192 = int32(0)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v194-v195 < v190 {
		v204 = v192
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v204 == int32(0) {
		goto L39
	} else {
		goto L50
	}
L47:
	;
	goto L46
L48:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v200 = F_memcmp(m, v198+v195, int32(_a_F_tamil_UTF_8_stem_5), v190)
	mBase = m.M
	if v200 != 0 {
		v204 = v192
		goto L47
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v190 + v195
	v204 = int32(1)
	goto L47
L50:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v207
	v209 = F_slice_del(m, l0)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	if v209 < int32(0) {
		v1615 = v209
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v213 = F_r_fix_va_start(m, l0)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	if v213 < int32(0) {
		v1615 = v213
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L39
L55:
	;
	if v356 < int32(0) {
		v1615 = v356
		goto L1
	} else {
		goto L98
	}
L56:
	;
	if v303 < int32(5) {
		v356 = v220
		goto L55
	} else {
		goto L72
	}
L57:
	;
	v303 = int32(0)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v234 = v229 & int32(3)
	if base.Ui32(v229) < base.Ui32(int32(4)) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v303 = v292
	goto L56
L61:
	;
	v276 = v270
	v277 = v271
	v281 = v220
	goto L69
L62:
	;
	v270 = v221
	v271 = int32(0)
	goto L61
L63:
	;
	goto L64
L64:
	;
	v241 = v221
	v242 = int32(0)
	v245 = v220
	goto L65
L65:
	;
	v247 = int32(*(*int8)(unsafe.Add(mBase, uint32(v241))))
	v248 = int32(-65)
	v251 = int32(*(*int8)(unsafe.Add(mBase, uint32(v241)+1)))
	v255 = int32(*(*int8)(unsafe.Add(mBase, uint32(v241)+2)))
	v259 = int32(*(*int8)(unsafe.Add(mBase, uint32(v241)+3)))
	v262 = v242 + base.B2i32(v248 < v247) + base.B2i32(v248 < v251) + base.B2i32(v248 < v255) + base.B2i32(v248 < v259)
	v263 = int32(4)
	v264 = v241 + v263
	v266 = v245 + v263
	if v266 != v229&int32(-4) {
		v241 = v264
		v242 = v262
		v245 = v266
		goto L65
	} else {
		goto L67
	}
L66:
	;
	if v234 == int32(0) {
		v292 = v262
		goto L60
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	v270 = v264
	v271 = v262
	goto L61
L69:
	;
	v282 = int32(*(*int8)(unsafe.Add(mBase, uint32(v276))))
	v285 = v277 + base.B2i32(int32(-65) < v282)
	v286 = int32(1)
	v289 = v281 + v286
	if v289 != v234 {
		v276 = v276 + v286
		v277 = v285
		v281 = v289
		goto L69
	} else {
		goto L71
	}
L70:
	;
	v292 = v285
	goto L60
L71:
	;
	goto L70
L72:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v306
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v308
	v313 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_6), int32(3))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	if v313 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v315
	v319 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_7))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L2
	} else {
		goto L77
	}
L75:
	;
	v323 = v220
	goto L76
L76:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v324
	v327 = v323
	goto L80
L77:
	;
	if v319 < int32(0) {
		v356 = v319
		goto L55
	} else {
		goto L78
	}
L78:
	;
	v323 = v319
	goto L76
L79:
	;
	v356 = int32(1)
	goto L55
L80:
	;
	v333 = F_r_fix_ending(m, l0)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L2
	} else {
		goto L85
	}
L81:
	;
	if v341 == int32(0) {
		goto L79
	} else {
		goto L96
	}
L82:
	;
	if v333 < int32(0) {
		goto L89
	} else {
		goto L90
	}
L83:
	;
	v341 = int32(2)
	goto L82
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v324
	goto L79
L85:
	;
	v336 = int32(base.Ui32(v333) >> (uint(int32(31)) % 32))
	if v333 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v338 = v336
	goto L88
L87:
	;
	v338 = int32(4)
	goto L88
L88:
	;
	switch v338 {
	case 0:
		goto L83
	default:
		v341 = v336
		goto L82
	case 4:
		goto L84
	}
L89:
	;
	v344 = v333
	goto L91
L90:
	;
	v344 = v327
	goto L91
L91:
	;
	if v333 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v345 = v344
	goto L94
L93:
	;
	v345 = v327
	goto L94
L94:
	;
	if v341 == int32(2) {
		v327 = v345
		goto L80
	} else {
		goto L95
	}
L95:
	;
	goto L81
L96:
	;
	if v345 < int32(0) {
		v356 = v345
		goto L55
	} else {
		goto L97
	}
L97:
	;
	goto L79
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
	v365 = int32(0)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v366-int32(4))))
	if v374 == v365 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	if v490 < int32(0) {
		v1615 = v490
		goto L1
	} else {
		goto L126
	}
L100:
	;
	if v448 < int32(5) {
		v490 = v365
		goto L99
	} else {
		goto L116
	}
L101:
	;
	v448 = int32(0)
	goto L100
L102:
	;
	goto L103
L103:
	;
	v379 = v374 & int32(3)
	if base.Ui32(v374) < base.Ui32(int32(4)) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v448 = v437
	goto L100
L105:
	;
	v421 = v415
	v422 = v416
	v426 = v365
	goto L113
L106:
	;
	v415 = v366
	v416 = int32(0)
	goto L105
L107:
	;
	goto L108
L108:
	;
	v386 = v366
	v387 = int32(0)
	v390 = v365
	goto L109
L109:
	;
	v392 = int32(*(*int8)(unsafe.Add(mBase, uint32(v386))))
	v393 = int32(-65)
	v396 = int32(*(*int8)(unsafe.Add(mBase, uint32(v386)+1)))
	v400 = int32(*(*int8)(unsafe.Add(mBase, uint32(v386)+2)))
	v404 = int32(*(*int8)(unsafe.Add(mBase, uint32(v386)+3)))
	v407 = v387 + base.B2i32(v393 < v392) + base.B2i32(v393 < v396) + base.B2i32(v393 < v400) + base.B2i32(v393 < v404)
	v408 = int32(4)
	v409 = v386 + v408
	v411 = v390 + v408
	if v411 != v374&int32(-4) {
		v386 = v409
		v387 = v407
		v390 = v411
		goto L109
	} else {
		goto L111
	}
L110:
	;
	if v379 == int32(0) {
		v437 = v407
		goto L104
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	v415 = v409
	v416 = v407
	goto L105
L113:
	;
	v427 = int32(*(*int8)(unsafe.Add(mBase, uint32(v421))))
	v430 = v422 + base.B2i32(int32(-65) < v427)
	v431 = int32(1)
	v434 = v426 + v431
	if v434 != v379 {
		v421 = v421 + v431
		v422 = v430
		v426 = v434
		goto L113
	} else {
		goto L115
	}
L114:
	;
	v437 = v430
	goto L104
L115:
	;
	goto L114
L116:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v451
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v453
	v456 = int32(9)
	v458 = int32(0)
	if v453-v451 < v456 {
		v471 = v458
		goto L118
	} else {
		goto L119
	}
L117:
	;
	if v471 == int32(0) {
		v490 = v365
		goto L99
	} else {
		goto L121
	}
L118:
	;
	goto L117
L119:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v467 = F_memcmp(m, v464+v453-v456, int32(_a_F_tamil_UTF_8_stem_8), v456)
	mBase = m.M
	if v467 != 0 {
		v471 = v458
		goto L118
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v453 - v456
	v471 = int32(1)
	goto L118
L121:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v474
	v478 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_9))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	if v478 < int32(0) {
		v490 = v478
		goto L99
	} else {
		goto L123
	}
L123:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v482
	v484 = F_r_fix_ending(m, l0)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L2
	} else {
		goto L124
	}
L124:
	;
	if v484 < int32(0) {
		v490 = v484
		goto L99
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v482
	v490 = int32(1)
	goto L99
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
	v495 = int32(0)
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v496-int32(4))))
	if v504 == v495 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	if v655 < int32(0) {
		v1615 = v655
		goto L1
	} else {
		goto L178
	}
L128:
	;
	if v578 < int32(5) {
		v655 = v495
		goto L127
	} else {
		goto L144
	}
L129:
	;
	v578 = int32(0)
	goto L128
L130:
	;
	goto L131
L131:
	;
	v509 = v504 & int32(3)
	if base.Ui32(v504) < base.Ui32(int32(4)) {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	v578 = v567
	goto L128
L133:
	;
	v551 = v545
	v552 = v546
	v556 = v495
	goto L141
L134:
	;
	v545 = v496
	v546 = int32(0)
	goto L133
L135:
	;
	goto L136
L136:
	;
	v516 = v496
	v517 = int32(0)
	v520 = v495
	goto L137
L137:
	;
	v522 = int32(*(*int8)(unsafe.Add(mBase, uint32(v516))))
	v523 = int32(-65)
	v526 = int32(*(*int8)(unsafe.Add(mBase, uint32(v516)+1)))
	v530 = int32(*(*int8)(unsafe.Add(mBase, uint32(v516)+2)))
	v534 = int32(*(*int8)(unsafe.Add(mBase, uint32(v516)+3)))
	v537 = v517 + base.B2i32(v523 < v522) + base.B2i32(v523 < v526) + base.B2i32(v523 < v530) + base.B2i32(v523 < v534)
	v538 = int32(4)
	v539 = v516 + v538
	v541 = v520 + v538
	if v541 != v504&int32(-4) {
		v516 = v539
		v517 = v537
		v520 = v541
		goto L137
	} else {
		goto L139
	}
L138:
	;
	if v509 == int32(0) {
		v567 = v537
		goto L132
	} else {
		goto L140
	}
L139:
	;
	goto L138
L140:
	;
	v545 = v539
	v546 = v537
	goto L133
L141:
	;
	v557 = int32(*(*int8)(unsafe.Add(mBase, uint32(v551))))
	v560 = v552 + base.B2i32(int32(-65) < v557)
	v561 = int32(1)
	v564 = v556 + v561
	if v564 != v509 {
		v551 = v551 + v561
		v552 = v560
		v556 = v564
		goto L141
	} else {
		goto L143
	}
L142:
	;
	v567 = v560
	goto L132
L143:
	;
	goto L142
L144:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v581
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v583
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v583
	v588 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_10), int32(26))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L2
	} else {
		goto L145
	}
L145:
	;
	if v588 == int32(0) {
		v655 = v495
		goto L127
	} else {
		goto L146
	}
L146:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v592
	switch v588 - int32(1) {
	case 0:
		goto L150
	case 1:
		goto L149
	case 2:
		goto L148
	default:
		v621 = v495
		goto L147
	}
L147:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v623
	v626 = v621
	goto L160
L148:
	;
	v617 = F_slice_del(m, l0)
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L2
	} else {
		goto L157
	}
L149:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v605 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_11), int32(8))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L2
	} else {
		goto L153
	}
L150:
	;
	v598 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_12))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L2
	} else {
		goto L151
	}
L151:
	;
	if int32(0) <= v598 {
		v621 = v598
		goto L147
	} else {
		goto L152
	}
L152:
	;
	v655 = v598
	goto L127
L153:
	;
	if v605 != 0 {
		v655 = v495
		goto L127
	} else {
		goto L154
	}
L154:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v607 + (v592 - v602)
	v613 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_13))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L2
	} else {
		goto L155
	}
L155:
	;
	if int32(0) <= v613 {
		v621 = v613
		goto L147
	} else {
		goto L156
	}
L156:
	;
	v655 = v613
	goto L127
L157:
	;
	if v617 < int32(0) {
		v655 = v617
		goto L127
	} else {
		goto L158
	}
L158:
	;
	v621 = v617
	goto L147
L159:
	;
	v655 = int32(1)
	goto L127
L160:
	;
	v632 = F_r_fix_ending(m, l0)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L2
	} else {
		goto L165
	}
L161:
	;
	if v640 == int32(0) {
		goto L159
	} else {
		goto L176
	}
L162:
	;
	if v632 < int32(0) {
		goto L169
	} else {
		goto L170
	}
L163:
	;
	v640 = int32(2)
	goto L162
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v623
	goto L159
L165:
	;
	v635 = int32(base.Ui32(v632) >> (uint(int32(31)) % 32))
	if v632 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v637 = v635
	goto L168
L167:
	;
	v637 = int32(4)
	goto L168
L168:
	;
	switch v637 {
	case 0:
		goto L163
	default:
		v640 = v635
		goto L162
	case 4:
		goto L164
	}
L169:
	;
	v643 = v632
	goto L171
L170:
	;
	v643 = v626
	goto L171
L171:
	;
	if v632 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v644 = v643
	goto L174
L173:
	;
	v644 = v626
	goto L174
L174:
	;
	if v640 == int32(2) {
		v626 = v644
		goto L160
	} else {
		goto L175
	}
L175:
	;
	goto L161
L176:
	;
	if v644 < int32(0) {
		v655 = v644
		goto L127
	} else {
		goto L177
	}
L177:
	;
	goto L159
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
	v664 = int32(0)
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v665))) = v664
	v668 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v668-int32(4))))
	if v676 == v664 {
		goto L181
	} else {
		goto L182
	}
L179:
	;
	if v1084 < int32(0) {
		v1615 = v1084
		goto L1
	} else {
		goto L296
	}
L180:
	;
	if v750 < int32(5) {
		v1084 = v664
		goto L179
	} else {
		goto L196
	}
L181:
	;
	v750 = int32(0)
	goto L180
L182:
	;
	goto L183
L183:
	;
	v681 = v676 & int32(3)
	if base.Ui32(v676) < base.Ui32(int32(4)) {
		goto L186
	} else {
		goto L187
	}
L184:
	;
	v750 = v739
	goto L180
L185:
	;
	v723 = v717
	v724 = v718
	v728 = v664
	goto L193
L186:
	;
	v717 = v668
	v718 = int32(0)
	goto L185
L187:
	;
	goto L188
L188:
	;
	v688 = v668
	v689 = int32(0)
	v692 = v664
	goto L189
L189:
	;
	v694 = int32(*(*int8)(unsafe.Add(mBase, uint32(v688))))
	v695 = int32(-65)
	v698 = int32(*(*int8)(unsafe.Add(mBase, uint32(v688)+1)))
	v702 = int32(*(*int8)(unsafe.Add(mBase, uint32(v688)+2)))
	v706 = int32(*(*int8)(unsafe.Add(mBase, uint32(v688)+3)))
	v709 = v689 + base.B2i32(v695 < v694) + base.B2i32(v695 < v698) + base.B2i32(v695 < v702) + base.B2i32(v695 < v706)
	v710 = int32(4)
	v711 = v688 + v710
	v713 = v692 + v710
	if v713 != v676&int32(-4) {
		v688 = v711
		v689 = v709
		v692 = v713
		goto L189
	} else {
		goto L191
	}
L190:
	;
	if v681 == int32(0) {
		v739 = v709
		goto L184
	} else {
		goto L192
	}
L191:
	;
	goto L190
L192:
	;
	v717 = v711
	v718 = v709
	goto L185
L193:
	;
	v729 = int32(*(*int8)(unsafe.Add(mBase, uint32(v723))))
	v732 = v724 + base.B2i32(int32(-65) < v729)
	v733 = int32(1)
	v736 = v728 + v733
	if v736 != v681 {
		v723 = v723 + v733
		v724 = v732
		v728 = v736
		goto L193
	} else {
		goto L195
	}
L194:
	;
	v739 = v732
	goto L184
L195:
	;
	goto L194
L196:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v753
	v755 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v755
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v755
	if v755-int32(2) <= v753 {
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1019
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1021))) = int32(1)
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1024
	v1026 = int32(9)
	v1028 = int32(0)
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1024-v1031 < v1026 {
		v1041 = v1028
		goto L269
	} else {
		goto L270
	}
L198:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v951
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v951
	v954 = int32(3)
	v956 = int32(0)
	v959 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v951-v959 < v954 {
		v969 = v956
		goto L251
	} else {
		goto L252
	}
L199:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v763 = int32(1)
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v761+v755-v763))))
	if base.B2i32(v765&int32(224) != int32(128))|base.B2i32(v763<<(uint(v765)%32)&int32(-2147475197) == int32(0)) != 0 {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v779 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_14), int32(22))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L2
	} else {
		goto L201
	}
L201:
	;
	if v779 == int32(0) {
		goto L198
	} else {
		goto L202
	}
L202:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v783
	switch v779 - int32(1) {
	case 0:
		goto L209
	case 1:
		goto L208
	case 2:
		goto L207
	case 3:
		goto L206
	case 4:
		goto L205
	case 5:
		goto L204
	case 6:
		goto L203
	default:
		v1017 = v664
		goto L197
	}
L203:
	;
	v945 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_15))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L2
	} else {
		goto L248
	}
L204:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v933 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_16), int32(8))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L2
	} else {
		goto L244
	}
L205:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v918 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_17), int32(8))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L2
	} else {
		goto L240
	}
L206:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v825 = int32(0)
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v824-int32(4))))
	if v832 == v825 {
		goto L222
	} else {
		goto L223
	}
L207:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v798 = int32(3)
	v800 = int32(0)
	v802 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v802-v803 < v798 {
		v813 = v800
		goto L215
	} else {
		goto L216
	}
L208:
	;
	v793 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_18))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L2
	} else {
		goto L212
	}
L209:
	;
	v787 = F_slice_del(m, l0)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L2
	} else {
		goto L210
	}
L210:
	;
	if int32(0) <= v787 {
		v1017 = v787
		goto L197
	} else {
		goto L211
	}
L211:
	;
	v1084 = v787
	goto L179
L212:
	;
	if int32(0) <= v793 {
		v1017 = v793
		goto L197
	} else {
		goto L213
	}
L213:
	;
	v1084 = v793
	goto L179
L214:
	;
	if v813 != 0 {
		goto L198
	} else {
		goto L218
	}
L215:
	;
	goto L214
L216:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v809 = F_memcmp(m, v806+v802-v798, int32(_a_F_tamil_UTF_8_stem_19), v798)
	mBase = m.M
	if v809 != 0 {
		v813 = v800
		goto L215
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v802 - v798
	v813 = int32(1)
	goto L215
L218:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v814 + (v783 - v797)
	v820 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_20))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L2
	} else {
		goto L219
	}
L219:
	;
	if int32(0) <= v820 {
		v1017 = v820
		goto L197
	} else {
		goto L220
	}
L220:
	;
	v1084 = v820
	goto L179
L221:
	;
	if v906 < int32(7) {
		goto L198
	} else {
		goto L237
	}
L222:
	;
	v906 = int32(0)
	goto L221
L223:
	;
	goto L224
L224:
	;
	v837 = v832 & int32(3)
	if base.Ui32(v832) < base.Ui32(int32(4)) {
		goto L227
	} else {
		goto L228
	}
L225:
	;
	v906 = v895
	goto L221
L226:
	;
	v879 = v873
	v880 = v874
	v884 = v825
	goto L234
L227:
	;
	v873 = v824
	v874 = int32(0)
	goto L226
L228:
	;
	goto L229
L229:
	;
	v844 = v824
	v845 = int32(0)
	v848 = v825
	goto L230
L230:
	;
	v850 = int32(*(*int8)(unsafe.Add(mBase, uint32(v844))))
	v851 = int32(-65)
	v854 = int32(*(*int8)(unsafe.Add(mBase, uint32(v844)+1)))
	v858 = int32(*(*int8)(unsafe.Add(mBase, uint32(v844)+2)))
	v862 = int32(*(*int8)(unsafe.Add(mBase, uint32(v844)+3)))
	v865 = v845 + base.B2i32(v851 < v850) + base.B2i32(v851 < v854) + base.B2i32(v851 < v858) + base.B2i32(v851 < v862)
	v866 = int32(4)
	v867 = v844 + v866
	v869 = v848 + v866
	if v869 != v832&int32(-4) {
		v844 = v867
		v845 = v865
		v848 = v869
		goto L230
	} else {
		goto L232
	}
L231:
	;
	if v837 == int32(0) {
		v895 = v865
		goto L225
	} else {
		goto L233
	}
L232:
	;
	goto L231
L233:
	;
	v873 = v867
	v874 = v865
	goto L226
L234:
	;
	v885 = int32(*(*int8)(unsafe.Add(mBase, uint32(v879))))
	v888 = v880 + base.B2i32(int32(-65) < v885)
	v889 = int32(1)
	v892 = v884 + v889
	if v892 != v837 {
		v879 = v879 + v889
		v880 = v888
		v884 = v892
		goto L234
	} else {
		goto L236
	}
L235:
	;
	v895 = v888
	goto L225
L236:
	;
	goto L235
L237:
	;
	v911 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_21))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L2
	} else {
		goto L238
	}
L238:
	;
	if int32(0) <= v911 {
		v1017 = v911
		goto L197
	} else {
		goto L239
	}
L239:
	;
	v1084 = v911
	goto L179
L240:
	;
	if v918 != 0 {
		goto L198
	} else {
		goto L241
	}
L241:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v920 + (v783 - v915)
	v926 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_22))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L2
	} else {
		goto L242
	}
L242:
	;
	if int32(0) <= v926 {
		v1017 = v926
		goto L197
	} else {
		goto L243
	}
L243:
	;
	v1084 = v926
	goto L179
L244:
	;
	if v933 != 0 {
		goto L198
	} else {
		goto L245
	}
L245:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v935 + (v783 - v930)
	v939 = F_slice_del(m, l0)
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L2
	} else {
		goto L246
	}
L246:
	;
	if int32(0) <= v939 {
		v1017 = v939
		goto L197
	} else {
		goto L247
	}
L247:
	;
	v1084 = v939
	goto L179
L248:
	;
	if int32(0) <= v945 {
		v1017 = v945
		goto L197
	} else {
		goto L249
	}
L249:
	;
	v1084 = v945
	goto L179
L250:
	;
	if v969 == int32(0) {
		v1084 = v664
		goto L179
	} else {
		goto L254
	}
L251:
	;
	goto L250
L252:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v965 = F_memcmp(m, v962+v951-v954, int32(_a_F_tamil_UTF_8_stem_23), v954)
	mBase = m.M
	if v965 != 0 {
		v969 = v956
		goto L251
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v951 - v954
	v969 = int32(1)
	goto L251
L254:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v974 = v972 - v973
	v977 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_24), int32(6))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L2
	} else {
		goto L255
	}
L255:
	;
	if v977 != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v979 - v974
	v984 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_25), int32(6))
	mBase = m.M
	v985 = m.ExcPending
	if v985 != 0 {
		goto L2
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1007 = v1006 - v974
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1007
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1007
	v1012 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_26))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L2
	} else {
		goto L266
	}
L259:
	;
	if v984 == int32(0) {
		v1084 = v664
		goto L179
	} else {
		goto L260
	}
L260:
	;
	v988 = int32(3)
	v990 = int32(0)
	v992 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v993 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v992-v993 < v988 {
		v1003 = v990
		goto L262
	} else {
		goto L263
	}
L261:
	;
	if v1003 == int32(0) {
		v1084 = v664
		goto L179
	} else {
		goto L265
	}
L262:
	;
	goto L261
L263:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v999 = F_memcmp(m, v996+v992-v988, int32(_a_F_tamil_UTF_8_stem_27), v988)
	mBase = m.M
	if v999 != 0 {
		v1003 = v990
		goto L262
	} else {
		goto L264
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v992 - v988
	v1003 = int32(1)
	goto L262
L265:
	;
	goto L258
L266:
	;
	if v1012 < int32(0) {
		v1084 = v1012
		goto L179
	} else {
		goto L267
	}
L267:
	;
	v1017 = v1012
	goto L197
L268:
	;
	if v1041 != 0 {
		goto L272
	} else {
		goto L273
	}
L269:
	;
	goto L268
L270:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1037 = F_memcmp(m, v1034+v1024-v1026, int32(_a_F_tamil_UTF_8_stem_28), v1026)
	mBase = m.M
	if v1037 != 0 {
		v1041 = v1028
		goto L269
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1024 - v1026
	v1041 = int32(1)
	goto L269
L272:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1042
	v1046 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_29))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L2
	} else {
		goto L275
	}
L273:
	;
	v1050 = v1017
	goto L274
L274:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1051
	v1055 = v1050
	goto L278
L275:
	;
	if v1046 < int32(0) {
		v1084 = v1046
		goto L179
	} else {
		goto L276
	}
L276:
	;
	v1050 = v1046
	goto L274
L277:
	;
	v1084 = int32(1)
	goto L179
L278:
	;
	v1060 = F_r_fix_ending(m, l0)
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L2
	} else {
		goto L283
	}
L279:
	;
	if v1068 == int32(0) {
		goto L277
	} else {
		goto L294
	}
L280:
	;
	if v1060 < int32(0) {
		goto L287
	} else {
		goto L288
	}
L281:
	;
	v1068 = int32(2)
	goto L280
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1051
	goto L277
L283:
	;
	v1063 = int32(base.Ui32(v1060) >> (uint(int32(31)) % 32))
	if v1060 != 0 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1065 = v1063
	goto L286
L285:
	;
	v1065 = int32(4)
	goto L286
L286:
	;
	switch v1065 {
	case 0:
		goto L281
	default:
		v1068 = v1063
		goto L280
	case 4:
		goto L282
	}
L287:
	;
	v1071 = v1060
	goto L289
L288:
	;
	v1071 = v1055
	goto L289
L289:
	;
	if v1060 != 0 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1072 = v1071
	goto L292
L291:
	;
	v1072 = v1055
	goto L292
L292:
	;
	if v1068 == int32(2) {
		v1055 = v1072
		goto L278
	} else {
		goto L293
	}
L293:
	;
	goto L279
L294:
	;
	if v1072 < int32(0) {
		v1084 = v1072
		goto L179
	} else {
		goto L295
	}
L295:
	;
	goto L277
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
	v1092 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v105
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1095
	if v1095-int32(8) <= v105 {
		v1160 = v1092
		goto L297
	} else {
		goto L298
	}
L297:
	;
	if v1160 < int32(0) {
		v1615 = v1160
		goto L1
	} else {
		goto L321
	}
L298:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1101+v1095-int32(1)))))
	if v1105 != int32(141) {
		v1160 = v1092
		goto L297
	} else {
		goto L299
	}
L299:
	;
	v1110 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_30), int32(4))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L2
	} else {
		goto L300
	}
L300:
	;
	if v1110 == int32(0) {
		v1160 = v1092
		goto L297
	} else {
		goto L301
	}
L301:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1114
	switch v1110 - int32(1) {
	case 0:
		goto L306
	case 1:
		goto L305
	case 2:
		goto L304
	case 3:
		goto L303
	default:
		goto L302
	}
L302:
	;
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1157
	v1160 = int32(1)
	goto L297
L303:
	;
	v1151 = F_slice_del(m, l0)
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L2
	} else {
		goto L319
	}
L304:
	;
	v1147 = F_slice_from_s(m, l0, int32(6), int32(_a_F_tamil_UTF_8_stem_31))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L2
	} else {
		goto L317
	}
L305:
	;
	v1141 = F_slice_from_s(m, l0, int32(6), int32(_a_F_tamil_UTF_8_stem_32))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L2
	} else {
		goto L315
	}
L306:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1121 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_33), int32(6))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L2
	} else {
		goto L307
	}
L307:
	;
	if v1121 != 0 {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1125 = F_slice_from_s(m, l0, int32(9), int32(_a_F_tamil_UTF_8_stem_34))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L2
	} else {
		goto L311
	}
L309:
	;
	goto L310
L310:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1129 + (v1114 - v1118)
	v1135 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_35))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L2
	} else {
		goto L313
	}
L311:
	;
	if int32(0) <= v1125 {
		goto L302
	} else {
		goto L312
	}
L312:
	;
	v1160 = v1125
	goto L297
L313:
	;
	if int32(0) <= v1135 {
		goto L302
	} else {
		goto L314
	}
L314:
	;
	v1160 = v1135
	goto L297
L315:
	;
	if int32(0) <= v1141 {
		goto L302
	} else {
		goto L316
	}
L316:
	;
	v1160 = v1141
	goto L297
L317:
	;
	if int32(0) <= v1147 {
		goto L302
	} else {
		goto L318
	}
L318:
	;
	v1160 = v1147
	goto L297
L319:
	;
	if v1151 < int32(0) {
		v1160 = v1151
		goto L297
	} else {
		goto L320
	}
L320:
	;
	goto L302
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
	v1165 = int32(0)
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1166-int32(4))))
	if v1174 == v1165 {
		goto L324
	} else {
		goto L325
	}
L322:
	;
	if v1282 < int32(0) {
		v1615 = v1282
		goto L1
	} else {
		goto L346
	}
L323:
	;
	if v1248 < int32(5) {
		v1282 = v1165
		goto L322
	} else {
		goto L339
	}
L324:
	;
	v1248 = int32(0)
	goto L323
L325:
	;
	goto L326
L326:
	;
	v1179 = v1174 & int32(3)
	if base.Ui32(v1174) < base.Ui32(int32(4)) {
		goto L329
	} else {
		goto L330
	}
L327:
	;
	v1248 = v1237
	goto L323
L328:
	;
	v1221 = v1215
	v1222 = v1216
	v1226 = v1165
	goto L336
L329:
	;
	v1215 = v1166
	v1216 = int32(0)
	goto L328
L330:
	;
	goto L331
L331:
	;
	v1186 = v1166
	v1187 = int32(0)
	v1190 = v1165
	goto L332
L332:
	;
	v1192 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1186))))
	v1193 = int32(-65)
	v1196 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1186)+1)))
	v1200 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1186)+2)))
	v1204 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1186)+3)))
	v1207 = v1187 + base.B2i32(v1193 < v1192) + base.B2i32(v1193 < v1196) + base.B2i32(v1193 < v1200) + base.B2i32(v1193 < v1204)
	v1208 = int32(4)
	v1209 = v1186 + v1208
	v1211 = v1190 + v1208
	if v1211 != v1174&int32(-4) {
		v1186 = v1209
		v1187 = v1207
		v1190 = v1211
		goto L332
	} else {
		goto L334
	}
L333:
	;
	if v1179 == int32(0) {
		v1237 = v1207
		goto L327
	} else {
		goto L335
	}
L334:
	;
	goto L333
L335:
	;
	v1215 = v1209
	v1216 = v1207
	goto L328
L336:
	;
	v1227 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1221))))
	v1230 = v1222 + base.B2i32(int32(-65) < v1227)
	v1231 = int32(1)
	v1234 = v1226 + v1231
	if v1234 != v1179 {
		v1221 = v1221 + v1231
		v1222 = v1230
		v1226 = v1234
		goto L336
	} else {
		goto L338
	}
L337:
	;
	v1237 = v1230
	goto L327
L338:
	;
	goto L337
L339:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1251
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1253
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1253
	if v1253-int32(5) <= v1251 {
		v1282 = v1165
		goto L322
	} else {
		goto L340
	}
L340:
	;
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1259+v1253-int32(1)))))
	if v1263 != int32(191) {
		v1282 = v1165
		goto L322
	} else {
		goto L341
	}
L341:
	;
	v1268 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_36), int32(2))
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L2
	} else {
		goto L342
	}
L342:
	;
	if v1268 == int32(0) {
		v1282 = v1165
		goto L322
	} else {
		goto L343
	}
L343:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1272
	v1274 = F_slice_del(m, l0)
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L2
	} else {
		goto L344
	}
L344:
	;
	if v1274 < int32(0) {
		v1282 = v1274
		goto L322
	} else {
		goto L345
	}
L345:
	;
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1278
	v1282 = int32(1)
	goto L322
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1288)+4)) = int32(1)
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1295 = v1288 + int32(4)
	v1296 = int32(0)
	goto L348
L347:
	;
	if v1605 < int32(0) {
		v1615 = v1605
		goto L1
	} else {
		goto L434
	}
L348:
	;
	v1301 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1295))) = v1301
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1303-int32(4))))
	if v1311 == v1301 {
		goto L352
	} else {
		goto L353
	}
L349:
	;
	v1605 = int32(1)
	goto L347
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1293
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v1598)+4))
	if v1601 != 0 {
		v1295 = v1598 + int32(4)
		v1296 = v1598
		goto L348
	} else {
		goto L433
	}
L351:
	;
	if v1385 < int32(5) {
		goto L350
	} else {
		goto L367
	}
L352:
	;
	v1385 = int32(0)
	goto L351
L353:
	;
	goto L354
L354:
	;
	v1316 = v1311 & int32(3)
	if base.Ui32(v1311) < base.Ui32(int32(4)) {
		goto L357
	} else {
		goto L358
	}
L355:
	;
	v1385 = v1374
	goto L351
L356:
	;
	v1358 = v1352
	v1359 = v1353
	v1363 = v1301
	goto L364
L357:
	;
	v1352 = v1303
	v1353 = int32(0)
	goto L356
L358:
	;
	goto L359
L359:
	;
	v1323 = v1303
	v1324 = int32(0)
	v1327 = v1301
	goto L360
L360:
	;
	v1329 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1323))))
	v1330 = int32(-65)
	v1333 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1323)+1)))
	v1337 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1323)+2)))
	v1341 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1323)+3)))
	v1344 = v1324 + base.B2i32(v1330 < v1329) + base.B2i32(v1330 < v1333) + base.B2i32(v1330 < v1337) + base.B2i32(v1330 < v1341)
	v1345 = int32(4)
	v1346 = v1323 + v1345
	v1348 = v1327 + v1345
	if v1348 != v1311&int32(-4) {
		v1323 = v1346
		v1324 = v1344
		v1327 = v1348
		goto L360
	} else {
		goto L362
	}
L361:
	;
	if v1316 == int32(0) {
		v1374 = v1344
		goto L355
	} else {
		goto L363
	}
L362:
	;
	goto L361
L363:
	;
	v1352 = v1346
	v1353 = v1344
	goto L356
L364:
	;
	v1364 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1358))))
	v1367 = v1359 + base.B2i32(int32(-65) < v1364)
	v1368 = int32(1)
	v1371 = v1363 + v1368
	if v1371 != v1316 {
		v1358 = v1358 + v1368
		v1359 = v1367
		v1363 = v1371
		goto L364
	} else {
		goto L366
	}
L365:
	;
	v1374 = v1367
	goto L355
L366:
	;
	goto L365
L367:
	;
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1388
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1390
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1390
	v1395 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_37), int32(46))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L2
	} else {
		goto L369
	}
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1529
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1529
	v1532 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1529-int32(8) <= v1532 {
		v1562 = v1526
		goto L409
	} else {
		goto L410
	}
L369:
	;
	if v1395 != 0 {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1397
	switch v1395 - int32(1) {
	case 0:
		goto L379
	case 1:
		goto L378
	case 2:
		goto L377
	case 3:
		goto L376
	case 4:
		goto L375
	case 5:
		goto L374
	default:
		v1514 = v1397
		goto L373
	}
L371:
	;
	v1521 = v1296
	goto L372
L372:
	;
	v1524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1526 = v1521
	v1529 = v1524
	goto L368
L373:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1517)+4)) = int32(1)
	v1521 = v1514
	goto L372
L374:
	;
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1487 = int32(3)
	v1489 = int32(0)
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1491-v1492 < v1487 {
		v1502 = v1489
		goto L403
	} else {
		goto L404
	}
L375:
	;
	v1482 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_38))
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L2
	} else {
		goto L400
	}
L376:
	;
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1454 = int32(3)
	v1456 = int32(0)
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1458-v1459 < v1454 {
		v1469 = v1456
		goto L394
	} else {
		goto L395
	}
L377:
	;
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1443 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_39), int32(8))
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L2
	} else {
		goto L389
	}
L378:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1397-int32(2) <= v1406 {
		v1431 = v1405
		goto L382
	} else {
		goto L383
	}
L379:
	;
	v1401 = F_slice_del(m, l0)
	mBase = m.M
	v1402 = m.ExcPending
	if v1402 != 0 {
		goto L2
	} else {
		goto L380
	}
L380:
	;
	if int32(0) <= v1401 {
		v1514 = v1401
		goto L373
	} else {
		goto L381
	}
L381:
	;
	v1605 = v1401
	goto L347
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1397 - v1405 + v1431
	v1436 = F_slice_del(m, l0)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L2
	} else {
		goto L387
	}
L383:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1412 = int32(1)
	v1414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1410+v1397-v1412))))
	if base.B2i32(v1414&int32(224) != int32(128))|base.B2i32(v1412<<(uint(v1414)%32)&int32(_a_F_tamil_UTF_8_stem_40) == int32(0)) != 0 {
		v1431 = v1405
		goto L382
	} else {
		goto L384
	}
L384:
	;
	v1428 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_41), int32(12))
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L2
	} else {
		goto L385
	}
L385:
	;
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1428 != 0 {
		v1526 = v1397
		v1529 = v1430
		goto L368
	} else {
		goto L386
	}
L386:
	;
	v1431 = v1430
	goto L382
L387:
	;
	if int32(0) <= v1436 {
		v1514 = v1436
		goto L373
	} else {
		goto L388
	}
L388:
	;
	v1605 = v1436
	goto L347
L389:
	;
	v1445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1443 != 0 {
		v1526 = v1397
		v1529 = v1445
		goto L368
	} else {
		goto L390
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1397 - v1440 + v1445
	v1449 = F_slice_del(m, l0)
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L2
	} else {
		goto L391
	}
L391:
	;
	if int32(0) <= v1449 {
		v1514 = v1449
		goto L373
	} else {
		goto L392
	}
L392:
	;
	v1605 = v1449
	goto L347
L393:
	;
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1469 != 0 {
		v1526 = v1397
		v1529 = v1470
		goto L368
	} else {
		goto L397
	}
L394:
	;
	goto L393
L395:
	;
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1465 = F_memcmp(m, v1462+v1458-v1454, int32(_a_F_tamil_UTF_8_stem_42), v1454)
	mBase = m.M
	if v1465 != 0 {
		v1469 = v1456
		goto L394
	} else {
		goto L396
	}
L396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1458 - v1454
	v1469 = int32(1)
	goto L394
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1397 - v1453 + v1470
	v1476 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_43))
	mBase = m.M
	v1477 = m.ExcPending
	if v1477 != 0 {
		goto L2
	} else {
		goto L398
	}
L398:
	;
	if int32(0) <= v1476 {
		v1514 = v1476
		goto L373
	} else {
		goto L399
	}
L399:
	;
	v1605 = v1476
	goto L347
L400:
	;
	if int32(0) <= v1482 {
		v1514 = v1482
		goto L373
	} else {
		goto L401
	}
L401:
	;
	v1605 = v1482
	goto L347
L402:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1502 == int32(0) {
		v1526 = v1397
		v1529 = v1503
		goto L368
	} else {
		goto L406
	}
L403:
	;
	goto L402
L404:
	;
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1498 = F_memcmp(m, v1495+v1491-v1487, int32(_a_F_tamil_UTF_8_stem_44), v1487)
	mBase = m.M
	if v1498 != 0 {
		v1502 = v1489
		goto L403
	} else {
		goto L405
	}
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1491 - v1487
	v1502 = int32(1)
	goto L403
L406:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1397 - v1486 + v1503
	v1509 = F_slice_del(m, l0)
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L2
	} else {
		goto L407
	}
L407:
	;
	if v1509 < int32(0) {
		v1605 = v1509
		goto L347
	} else {
		goto L408
	}
L408:
	;
	v1514 = v1509
	goto L373
L409:
	;
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1563
	v1567 = v1562
	goto L416
L410:
	;
	v1536 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1536+v1529-int32(1)))))
	if base.B2i32(v1540 != int32(177))&base.B2i32(v1540 != int32(141)) != 0 {
		v1562 = v1526
		goto L409
	} else {
		goto L411
	}
L411:
	;
	v1548 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_45), int32(6))
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L2
	} else {
		goto L412
	}
L412:
	;
	if v1548 == int32(0) {
		v1562 = v1526
		goto L409
	} else {
		goto L413
	}
L413:
	;
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1552
	v1554 = F_slice_del(m, l0)
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L2
	} else {
		goto L414
	}
L414:
	;
	if v1554 < int32(0) {
		v1605 = v1554
		goto L347
	} else {
		goto L415
	}
L415:
	;
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1558)+4)) = int32(1)
	v1562 = v1554
	goto L409
L416:
	;
	v1572 = F_r_fix_ending(m, l0)
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L2
	} else {
		goto L420
	}
L417:
	;
	if v1579 == int32(0) {
		goto L350
	} else {
		goto L431
	}
L418:
	;
	if v1572 < int32(0) {
		goto L424
	} else {
		goto L425
	}
L419:
	;
	v1579 = int32(2)
	goto L418
L420:
	;
	v1575 = int32(base.Ui32(v1572) >> (uint(int32(31)) % 32))
	if v1572 != 0 {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v1577 = v1575
	goto L423
L422:
	;
	v1577 = int32(4)
	goto L423
L423:
	;
	switch v1577 {
	case 0:
		goto L419
	default:
		v1579 = v1575
		goto L418
	case 4:
		goto L350
	}
L424:
	;
	v1582 = v1572
	goto L426
L425:
	;
	v1582 = v1567
	goto L426
L426:
	;
	if v1572 != 0 {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v1583 = v1582
	goto L429
L428:
	;
	v1583 = v1567
	goto L429
L429:
	;
	if v1579 == int32(2) {
		v1567 = v1583
		goto L416
	} else {
		goto L430
	}
L430:
	;
	goto L417
L431:
	;
	if v1583 < int32(0) {
		v1605 = v1583
		goto L347
	} else {
		goto L432
	}
L432:
	;
	goto L350
L433:
	;
	goto L349
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105
	v1615 = int32(1)
	goto L1
}
func F_texteqname(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v14 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v44 = F_strlen(m, v13)
	mBase = m.M
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v45 != int32(950) {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v20 == int32(18) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v31 = int32(1)
	if v14&v31 != 0 {
		v43 = int32(base.Ui32(v14)>>(uint(v31)%32)) - v31
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v23 = int32(16)
	goto L9
L8:
	;
	v23 = int32(0)
	goto L9
L9:
	;
	if base.Ui32((v20-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v30 = int32(4)
	goto L12
L11:
	;
	v30 = v23
	goto L12
L12:
	;
	v43 = v30
	goto L3
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v43 = int32(base.Ui32(v37)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L14:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v151 != v9 {
		goto L51
	} else {
		goto L52
	}
L15:
	;
	v140 = int32(1)
	if v14&v140 != 0 {
		goto L47
	} else {
		goto L48
	}
L16:
	;
	if v45 != 0 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v43 != v44 {
		v150 = int32(0)
		goto L14
	} else {
		goto L25
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(_a_F_texteqname_0), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errhint(m, int32(_a_F_texteqname_1), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_texteqname_2), int32(1648), int32(_a_F_texteqname_3))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	v70 = int32(1)
	if v14&v70 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v74 = v70
	goto L28
L27:
	;
	v74 = int32(4)
	goto L28
L28:
	;
	v75 = v9 + v74
	if base.Ui32(int32(4)) <= base.Ui32(v43) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v150 = base.B2i32(v137 == int32(0))
	goto L14
L30:
	;
	v137 = int32(0)
	goto L29
L31:
	;
	v111 = v106
	v112 = v107
	v113 = v108
	goto L41
L32:
	;
	if (v75|v13)&int32(3) != 0 {
		v106 = v75
		v107 = v13
		v108 = v43
		goto L31
	} else {
		goto L35
	}
L33:
	;
	v99 = v75
	v100 = v13
	v101 = v43
	goto L34
L34:
	;
	if v101 == int32(0) {
		goto L30
	} else {
		goto L40
	}
L35:
	;
	v83 = v75
	v84 = v13
	v85 = v43
	goto L36
L36:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v88 != v89 {
		v106 = v83
		v107 = v84
		v108 = v85
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v99 = v94
	v100 = v92
	v101 = v96
	goto L34
L38:
	;
	v91 = int32(4)
	v92 = v84 + v91
	v94 = v83 + v91
	v96 = v85 - v91
	if base.Ui32(int32(3)) < base.Ui32(v96) {
		v83 = v94
		v84 = v92
		v85 = v96
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v106 = v99
	v107 = v100
	v108 = v101
	goto L31
L41:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v116 == v117 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v137 = v116 - v117
	goto L29
L43:
	;
	v119 = int32(1)
	v124 = v113 - v119
	if v124 != 0 {
		v111 = v111 + v119
		v112 = v112 + v119
		v113 = v124
		goto L41
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	goto L42
L46:
	;
	goto L30
L47:
	;
	v144 = v140
	goto L49
L48:
	;
	v144 = int32(4)
	goto L49
L49:
	;
	v146 = F_varstr_cmp(m, v9+v144, v43, v13, v44, v45)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v150 = base.B2i32(v146 == int32(0))
	goto L14
L51:
	;
	F_pfree(m, v9)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	return v150
L54:
	;
	goto L53
}
func F_texticregexeq_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_like_regex_support(m, v2, int32(3))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_textlen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_textlen[0]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9*int32(28))+uint32(_c_F_textlen[1])))
	if v14 == int32(1) {
		v17 = F_toast_raw_datum_size(m, v6)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			return v17 - int32(4)
		}
	} else {
		v24 = F_pg_detoast_datum_packed(m, v6)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(1)
			v27 = v24 + v26
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
			v32 = v30 & v26
			if v32 != 0 {
				v33 = v27
			} else {
				v33 = v24 + int32(4)
			}
			if v30 == int32(1) {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
				if v39 == int32(18) {
					v42 = int32(16)
				} else {
					v42 = int32(0)
				}
				if base.Ui32((v39-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v49 = int32(4)
				} else {
					v49 = v42
				}
				v50 = F_pg_mbstrlen_with_len(m, v33, v49)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					return v50
				}
			} else {
				if v32 != 0 {
					v53 = int32(1)
					v57 = F_pg_mbstrlen_with_len(m, v33, int32(base.Ui32(v30)>>(uint(v53)%32))-v53)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						return v57
					}
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
					v65 = F_pg_mbstrlen_with_len(m, v33, int32(base.Ui32(v60)>>(uint(int32(2))%32))-int32(4))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						return v65
					}
				}
			}
		}
	}
}
func F_textlike(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = v8 + int32(1)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			v21 = v19 & int32(1)
			if v21 != 0 {
				v22 = v13
			} else {
				v22 = v8 + int32(4)
			}
			if v19 == int32(1) {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
				if v28 == int32(18) {
					v31 = int32(16)
				} else {
					v31 = int32(0)
				}
				if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v38 = int32(4)
				} else {
					v38 = v31
				}
				v49 = v38
			} else {
				v39 = int32(1)
				if v21 != 0 {
					v49 = int32(base.Ui32(v19)>>(uint(v39)%32)) - v39
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v50 = int32(1)
			v51 = v15 + v50
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			v56 = v54 & v50
			if v56 != 0 {
				v57 = v51
			} else {
				v57 = v15 + int32(4)
			}
			if v54 == int32(1) {
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
				if v63 == int32(18) {
					v66 = int32(16)
				} else {
					v66 = int32(0)
				}
				if base.Ui32((v63-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v73 = int32(4)
				} else {
					v73 = v66
				}
				v84 = v73
			} else {
				v74 = int32(1)
				if v56 != 0 {
					v84 = int32(base.Ui32(v54)>>(uint(v74)%32)) - v74
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v84 = int32(base.Ui32(v78)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v86 = F_GenericMatchText(m, v22, v49, v57, v84, v85)
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v86 == int32(1))
			}
		}
	}
}
func F_textnlike(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = v8 + int32(1)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			v21 = v19 & int32(1)
			if v21 != 0 {
				v22 = v13
			} else {
				v22 = v8 + int32(4)
			}
			if v19 == int32(1) {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
				if v28 == int32(18) {
					v31 = int32(16)
				} else {
					v31 = int32(0)
				}
				if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v38 = int32(4)
				} else {
					v38 = v31
				}
				v49 = v38
			} else {
				v39 = int32(1)
				if v21 != 0 {
					v49 = int32(base.Ui32(v19)>>(uint(v39)%32)) - v39
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v50 = int32(1)
			v51 = v15 + v50
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			v56 = v54 & v50
			if v56 != 0 {
				v57 = v51
			} else {
				v57 = v15 + int32(4)
			}
			if v54 == int32(1) {
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
				if v63 == int32(18) {
					v66 = int32(16)
				} else {
					v66 = int32(0)
				}
				if base.Ui32((v63-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v73 = int32(4)
				} else {
					v73 = v66
				}
				v84 = v73
			} else {
				v74 = int32(1)
				if v56 != 0 {
					v84 = int32(base.Ui32(v54)>>(uint(v74)%32)) - v74
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v84 = int32(base.Ui32(v78)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v86 = F_GenericMatchText(m, v22, v49, v57, v84, v85)
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v86 != int32(1))
			}
		}
	}
}
func F_tidgt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+2)))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2))))
	v9 = int32(16)
	v11 = v7 | v8<<(uint(v9)%32)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v16 = v12 | v13<<(uint(v9)%32)
	if base.Ui32(v11) < base.Ui32(v16) {
		v27 = int32(-1)
	} else {
		if base.Ui32(v16) < base.Ui32(v11) {
			v27 = int32(1)
		} else {
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+4)))
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
			if base.Ui32(v21) < base.Ui32(v22) {
				v27 = int32(-1)
			} else {
				v27 = base.B2i32(base.Ui32(v22) < base.Ui32(v21))
			}
		}
	}
	return base.B2i32(int32(0) < v27)
}
func F_tidin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = v12
	v15 = int32(0)
	goto L4
L1:
	;
	m.G0 = v9 - int32(-64)
	return v145
L2:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v109 = F_strtox_2(m, v104, v7+int32(-12), int32(10), int64(4294967295))
	mBase = m.M
	v110 = base.I32_wrap_i64(v109)
	goto L29
L3:
	;
	v83 = int32(0)
	v84 = F_errsave_start(m, v11)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L18
	} else {
		goto L24
	}
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	switch v19 - int32(41) {
	case 0:
		goto L3
	case 1, 2:
		goto L8
	case 3:
		goto L7
	default:
		goto L9
	}
L5:
	;
	v42 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_tidin[0])) = v42
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
	v51 = F_strtox_2(m, v46, v7+int32(-12), int32(10), int64(4294967295))
	mBase = m.M
	v52 = base.I32_wrap_i64(v51)
	goto L13
L6:
	;
	v38 = int32(1)
	if v37 <= v38 {
		v13 = v13 + v38
		v15 = v37
		goto L4
	} else {
		goto L12
	}
L7:
	;
	v32 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7+int32(-8)+v15<<(uint(int32(2))%32)))) = v13 + v32
	v37 = v15 + v32
	goto L6
L8:
	;
	if base.B2i32(v19 != int32(40))|v15 != 0 {
		v37 = v15
		goto L6
	} else {
		goto L11
	}
L9:
	;
	if v19 == int32(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	goto L7
L12:
	;
	goto L5
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_tidin[0]))
	if v54 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v58 == int32(44) {
		goto L2
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v61 = F_errsave_start(m, v11)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	if v61 == int32(0) {
		v145 = v42
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(_a_F_tidin_0)
	F_errmsg(m, int32(_a_F_tidin_1), v7+int32(-48))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	F_errsave_finish(m, v11, int32(_a_F_tidin_2), int32(81), int32(_a_F_tidin_3))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v145 = v42
	goto L1
L24:
	;
	if v84 == int32(0) {
		v145 = v83
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = int32(_a_F_tidin_0)
	F_errmsg(m, int32(_a_F_tidin_1), v7+int32(-32))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	F_errsave_finish(m, v11, int32(_a_F_tidin_2), int32(73), int32(_a_F_tidin_3))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	v145 = v83
	goto L1
L29:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_tidin[0]))
	if v112 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v138 = F_palloc(m, int32(6))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L18
	} else {
		goto L40
	}
L31:
	;
	v119 = F_errsave_start(m, v11)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L18
	} else {
		goto L35
	}
L32:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v114 != int32(41) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	if base.Ui32(v110) < base.Ui32(int32(_a_F_tidin_4)) {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	if v119 == int32(0) {
		v145 = v42
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L18
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_F_tidin_0)
	F_errmsg(m, int32(_a_F_tidin_1), v9)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L18
	} else {
		goto L38
	}
L38:
	;
	F_errsave_finish(m, v11, int32(_a_F_tidin_2), int32(104), int32(_a_F_tidin_3))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L18
	} else {
		goto L39
	}
L39:
	;
	v145 = v42
	goto L1
L40:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+4)) = uint16(v110)
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+2)) = uint16(v52)
	v143 = int32(base.Ui32(v52) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v138))) = uint16(v143)
	v145 = v138
	goto L1
}
func F_tidne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+2)))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2))))
	v9 = int32(16)
	v11 = v7 | v8<<(uint(v9)%32)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v16 = v12 | v13<<(uint(v9)%32)
	if base.Ui32(v11) < base.Ui32(v16) {
		v27 = int32(-1)
	} else {
		if base.Ui32(v16) < base.Ui32(v11) {
			v27 = int32(1)
		} else {
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+4)))
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
			if base.Ui32(v21) < base.Ui32(v22) {
				v27 = int32(-1)
			} else {
				v27 = base.B2i32(base.Ui32(v22) < base.Ui32(v21))
			}
		}
	}
	return base.B2i32(v27 != int32(0))
}
func F_tqueueShutdownReceiver(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v3 != 0 {
		F_shm_mq_detach(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
			return
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
		return
	}
}
func F_trackitem_compare_lexemes(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v8 < v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	goto L3
L3:
	;
	if v6 < v8 {
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v6 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return v61
L8:
	;
	v61 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v22 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v23 = v15
	v24 = v16
	v25 = v6
	v26 = v22
	goto L15
L12:
	;
	v49 = v16
	v53 = int32(0)
	goto L13
L13:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v61 = v53 - v54
	goto L7
L14:
	;
	v49 = v44
	v53 = v46
	goto L13
L15:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if base.B2i32(v26 != v28)|base.B2i32(v28 == int32(0)) != 0 {
		v44 = v24
		v46 = v26
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v44 = v38
	v46 = int32(0)
	goto L14
L17:
	;
	v34 = v25 - int32(1)
	if v34 == int32(0) {
		v44 = v24
		v46 = v26
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v37 = int32(1)
	v38 = v24 + v37
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v39 != 0 {
		v23 = v23 + v37
		v24 = v38
		v25 = v34
		v26 = v39
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
}
func F_transformFrameOffset(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v242 int32
	_ = v242
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
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	v7 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v7
	if l5 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L10
	} else {
		goto L63
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L10
	} else {
		goto L55
	}
L3:
	;
	if l1&int32(4) != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v183 = v7
	goto L5
L5:
	;
	m.G0 = v17 + int32(48)
	return v183
L6:
	;
	F_checkExprIsVarFree(m, l0, v167, v162)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L10
	} else {
		goto L54
	}
L7:
	;
	v25 = F_transformExpr(m, l0, l5, int32(12))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	if l1&int32(2) != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	return int32(0)
L11:
	;
	v31 = F_coerce_to_specific_type(m, l0, v25, int32(20), int32(_a_F_transformFrameOffset_0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v162 = int32(_a_F_transformFrameOffset_0)
	v167 = v31
	goto L6
L13:
	;
	v36 = F_transformExpr(m, l0, l5, int32(11))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L10
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v147 = int32(0)
	if l1&int32(8) == v147 {
		v162 = v147
		v167 = v7
		goto L6
	} else {
		goto L51
	}
L16:
	;
	v38 = F_exprType(m, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v38
	v44 = F_SearchSysCacheList(m, int32(5), int32(2), l2, l3, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	if v46 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_ReleaseCatCacheList(m, v44)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v38 == int32(705) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L1
L23:
	;
	v53 = l3
	goto L25
L24:
	;
	v53 = v38
	goto L25
L25:
	;
	v59 = int32(0)
	v64 = v7
	v67 = v7
	v68 = v7
	v70 = v7
	goto L26
L26:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v44+int32(48)+v64<<(uint(int32(2))%32))))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+56))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+22)))
	v77 = v75 + v76
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+16)))
	if v78 != int32(3) {
		v98 = v59
		v99 = v67
		v100 = v68
		v101 = v70
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_ReleaseCatCacheList(m, v44)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L10
	} else {
		goto L36
	}
L28:
	;
	v103 = v64 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	if v103 < v104 {
		v59 = v98
		v64 = v103
		v67 = v99
		v68 = v100
		v70 = v101
		goto L26
	} else {
		goto L35
	}
L29:
	;
	v81 = int32(1)
	v82 = v59 + v81
	v89 = F_can_coerce_type(m, v81, v17+int32(44), v77+int32(12), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	if v89 == int32(0) {
		v98 = v82
		v99 = v67
		v100 = v68
		v101 = v70
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v94 = v68 + int32(1)
	if v53 == v67 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v98 = v82
	v99 = v53
	v100 = v94
	v101 = v70
	goto L28
L33:
	;
	goto L34
L34:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v98 = v82
	v99 = v97
	v100 = v94
	v101 = v96
	goto L28
L35:
	;
	goto L27
L36:
	;
	if v98 == int32(0) {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	switch v100 {
	case 0:
		goto L40
	case 1:
		goto L38
	default:
		goto L39
	}
L38:
	;
	v142 = int32(_a_F_transformFrameOffset_1)
	v144 = F_coerce_to_specific_type(m, l0, v36, v99, v142)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L10
	} else {
		goto L50
	}
L39:
	;
	if v53 != v99 {
		goto L2
	} else {
		goto L49
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L10
	} else {
		goto L42
	}
L42:
	;
	v117 = F_format_type_be(m, l3)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	v120 = F_format_type_be(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v117
	F_errmsg(m, int32(_a_F_transformFrameOffset_2), v17+int32(32))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	F_errhint(m, int32(_a_F_transformFrameOffset_3), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	v133 = F_exprLocation(m, v36)
	mBase = m.M
	F_parser_errposition(m, l0, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_transformFrameOffset_4), int32(3788), int32(_a_F_transformFrameOffset_5))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L10
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	goto L38
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v101
	v162 = v142
	v167 = v144
	goto L6
L51:
	;
	v154 = F_transformExpr(m, l0, l5, int32(13))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L10
	} else {
		goto L52
	}
L52:
	;
	v158 = F_coerce_to_specific_type(m, l0, v154, int32(20), int32(_a_F_transformFrameOffset_6))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	v162 = int32(_a_F_transformFrameOffset_6)
	v167 = v158
	goto L6
L54:
	;
	v183 = v167
	goto L5
L55:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	v201 = F_format_type_be(m, l3)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L10
	} else {
		goto L57
	}
L57:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	v204 = F_format_type_be(m, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v201
	F_errmsg(m, int32(_a_F_transformFrameOffset_7), v17+int32(16))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	F_errhint(m, int32(_a_F_transformFrameOffset_8), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	v217 = F_exprLocation(m, v36)
	mBase = m.M
	F_parser_errposition(m, l0, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_transformFrameOffset_4), int32(3796), int32(_a_F_transformFrameOffset_5))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	v246 = F_format_type_be(m, l3)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L10
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v246
	F_errmsg(m, int32(_a_F_transformFrameOffset_9), v17)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	v252 = F_exprLocation(m, v36)
	mBase = m.M
	F_parser_errposition(m, l0, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L10
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_transformFrameOffset_4), int32(3780), int32(_a_F_transformFrameOffset_5))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L10
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_transformWholeRowRef(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v11 != v13 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		if v15 == int32(0) {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v32 = int32(0)
			F_expandRTE(m, v12, v31, l2, v32, l3, v32, v32, v9+int32(12))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v40 = F_palloc0(m, int32(24))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(36)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					if v46 != 0 {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
						v49 = v47
					} else {
						v49 = int32(0)
					}
					v50 = int32(0)
					if base.B2i32(v44 == v50)|base.B2i32(v49 <= v50) != 0 {
						v60 = int32(0)
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
						if v49 < v57 {
							*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v49
						} else {
						}
						v60 = v44
					}
					*(*int64)(unsafe.Add(mBase, uint32(v40)+8)) = int64(8589936841)
					*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v60
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
					v66 = F_copyObjectImpl(m, v65)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v66
						v72 = v40
						m.G0 = v9 + int32(16)
						return v72
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v20 = F_makeWholeRowVar(m, v12, v18, l2, int32(1))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v24
				F_markNullableIfNeeded(m, l0, v20)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					F_markVarForSelectPriv(m, l0, v20)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v72 = v20
						m.G0 = v9 + int32(16)
						return v72
					}
				}
			}
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v20 = F_makeWholeRowVar(m, v12, v18, l2, int32(1))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
			*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v24
			F_markNullableIfNeeded(m, l0, v20)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				F_markVarForSelectPriv(m, l0, v20)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v72 = v20
					m.G0 = v9 + int32(16)
					return v72
				}
			}
		}
	}
}
func F_trgm_contained_by(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v77 int32
	_ = v77
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = int32(2)
	v12 = int32(5)
	v14 = int32(3)
	v15 = base.I32_div_u_s(int32(base.Ui32(v9)>>(uint(v10)%32))-v12, v14)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = int32(base.Ui32(v16)>>(uint(v10)%32)) - v12
	v22 = base.I32_div_u_s(v20, v14)
	if base.Ui32(v14) <= base.Ui32(v20) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v77
L2:
	;
	v25 = int32(5)
	v26 = l0 + v25
	v28 = l1 + v25
	v29 = v28
	v31 = v26
	goto L5
L3:
	;
	goto L4
L4:
	;
	v77 = int32(1)
	goto L1
L5:
	;
	v39 = base.I32_div_s(v29-v28, int32(3))
	if v15 <= v39 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	return int32(0)
L8:
	;
	goto L9
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_trgm_contained_by[0]))
	v45 = m.T0[v44].(func(*base.Module, int32, int32) int32)(m, v31, v29)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if v45 < int32(0) {
		v77 = int32(0)
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v51 = int32(3)
	if v45 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v55 = int32(0)
	goto L15
L14:
	;
	v55 = v51
	goto L15
L15:
	;
	v56 = v31 + v55
	v59 = base.I32_div_s(v56-v26, int32(3))
	if v59 < v22 {
		v29 = v29 + v51
		v31 = v56
		goto L5
	} else {
		goto L16
	}
L16:
	;
	goto L6
}
func F_trim_mergeclauses_for_inner_pathkeys(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	v3 = int32(0)
	if base.B2i32(l1 == v3)|base.B2i32(l0 == v3) != 0 {
		v98 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v98
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+120)))
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = int32(104)
	goto L8
L7:
	;
	v27 = int32(100)
	goto L8
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23+v27)))
	if v21 != v29 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	goto L11
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v35 = F_lappend(m, int32(0), v23)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v39 < int32(2) {
		v98 = v35
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if int32(1) < v33 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v47 = v19 + int32(4)
	goto L17
L16:
	;
	v47 = int32(0)
	goto L17
L17:
	;
	v51 = v47
	v53 = int32(1)
	v54 = v21
	v55 = v35
	goto L18
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v53<<(uint(int32(2))%32))))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+120)))
	if v64 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v98 = v86
	goto L1
L20:
	;
	v65 = int32(104)
	goto L22
L21:
	;
	v65 = int32(100)
	goto L22
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61+v65)))
	if v54 != v67 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v51 == int32(0) {
		v98 = v55
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v84 = v51
	goto L25
L25:
	;
	v86 = F_lappend(m, v55, v61)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L12
	} else {
		goto L31
	}
L26:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v67 != v72 {
		v98 = v55
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v75 = v51 + int32(4)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v75) < base.Ui32(v77+v78<<(uint(int32(2))%32)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v83 = v75
	goto L30
L29:
	;
	v83 = int32(0)
	goto L30
L30:
	;
	v84 = v83
	goto L25
L31:
	;
	v89 = v53 + int32(1)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v89 < v90 {
		v51 = v84
		v53 = v89
		v54 = v67
		v55 = v86
		goto L18
	} else {
		goto L32
	}
L32:
	;
	goto L19
}
func F_trueConsistentFn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+87)) = uint8(v2)
	return int32(1)
}
func F_try_mergejoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 float64
	_ = v238
	var v239 float64
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 float64
	_ = v278
	var v285 float64
	_ = v285
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	v12 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(112)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+108)) = v12
	if l10 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v17 + int32(112)
	return
L2:
	;
	F_try_partial_mergejoin_path(m, l0, l1, l2, l3, l4, l5, l6, l7, l8, l9)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l9)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return
L6:
	;
	goto L1
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v25 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v40 = F_calc_non_nestloop_required_outer(m, l2, l3)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L21
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v28 = v26
	goto L12
L11:
	;
	v28 = int32(0)
	goto L12
L12:
	;
	v29 = F_bms_is_member(m, v24, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	if v29 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l9)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v33 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v36 = v34
	goto L17
L16:
	;
	v36 = int32(0)
	goto L17
L17:
	;
	v37 = F_bms_is_member(m, v32, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	if v37 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L9
L20:
	;
	F_bms_free(m, v40)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L5
	} else {
		goto L141
	}
L21:
	;
	if v40 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l9)+32))
	v43 = int32(0)
	if base.B2i32(v40 == v43)|base.B2i32(v42 == v43) != 0 {
		v88 = v43
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L24
L24:
	;
	if l6 != 0 {
		goto L39
	} else {
		goto L40
	}
L25:
	;
	if v88 == int32(0) {
		goto L20
	} else {
		goto L38
	}
L26:
	;
	goto L25
L27:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v53 < v54 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v56 = v53
	goto L30
L29:
	;
	v56 = v54
	goto L30
L30:
	;
	if v56 <= int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v59 = int32(1)
	goto L33
L32:
	;
	v59 = v56
	goto L33
L33:
	;
	v60 = int32(8)
	v65 = int32(0)
	goto L34
L34:
	;
	v72 = v65 << (uint(int32(2)) % 32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v42+v60+v72)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v40+v60+v72)))
	v77 = v74 & v76
	v79 = base.B2i32(v77 != int32(0))
	if v77 != 0 {
		v88 = v79
		goto L26
	} else {
		goto L36
	}
L35:
	;
	v88 = v79
	goto L26
L36:
	;
	v81 = v65 + int32(1)
	if v81 != v59 {
		v65 = v81
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	goto L24
L39:
	;
	v91 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	v94 = v17 + int32(108)
	if l6 == v92 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v174 = v12
	goto L41
L41:
	;
	if l7 != 0 {
		goto L77
	} else {
		goto L78
	}
L42:
	;
	if v172 != 0 {
		goto L74
	} else {
		goto L75
	}
L43:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v160
	v172 = int32(1)
	goto L42
L44:
	;
	if l6 != 0 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	if l6 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(0)
	v172 = int32(1)
	goto L42
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(0)
	v172 = int32(1)
	goto L42
L49:
	;
	goto L50
L50:
	;
	if v92 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v112 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v112
	v172 = v112
	goto L42
L52:
	;
	goto L53
L53:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v116 = int32(0)
	if v116 < v115 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v119 = v115
	goto L56
L55:
	;
	v119 = v116
	goto L56
L56:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v125 = v91
	goto L57
L57:
	;
	if v125 < v120 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v136 = v132 + v125<<(uint(int32(2))%32)
	goto L61
L60:
	;
	v136 = int32(0)
	goto L61
L61:
	;
	if v125 == v119 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v119
	v172 = base.B2i32(v136 == int32(0))
	goto L42
L63:
	;
	goto L64
L64:
	;
	v142 = base.B2i32(v136 == int32(0))
	if v136 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v125
	v172 = v142
	goto L42
L66:
	;
	goto L67
L67:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	if v146 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v125
	v172 = v142
	goto L42
L69:
	;
	goto L70
L70:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v146+v125<<(uint(int32(2))%32))))
	if v150 != v154 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v125
	v172 = int32(0)
	goto L42
L72:
	;
	v125 = v125 + int32(1)
	goto L57
L74:
	;
	v173 = v91
	goto L76
L75:
	;
	v173 = l6
	goto L76
L76:
	;
	v174 = v173
	goto L41
L77:
	;
	v175 = int32(0)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	if l7 == v176 {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	v231 = v12
	goto L79
L79:
	;
	v233 = v17 + int32(8)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v17)+108))
	F_initial_cost_mergejoin(m, l0, v233, l8, l5, l2, l3, v174, v231, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L5
	} else {
		goto L101
	}
L80:
	;
	if v229 != 0 {
		goto L98
	} else {
		goto L99
	}
L81:
	;
	v229 = int32(1)
	goto L80
L82:
	;
	goto L83
L83:
	;
	v185 = v175
	goto L85
L84:
	;
	v229 = v221
	goto L80
L85:
	;
	v189 = int32(0)
	if l7 == v189 {
		v199 = v189
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v221 = int32(0)
	goto L84
L87:
	;
	if v176 != 0 {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v193 <= v185 {
		v199 = int32(0)
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v199 = v195 + v185<<(uint(int32(2))%32)
	goto L87
L90:
	;
	v205 = base.B2i32(v199 == int32(0))
	if v199 == int32(0) {
		v221 = v205
		goto L84
	} else {
		goto L95
	}
L91:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v185 < v200 {
		goto L90
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v229 = base.B2i32(v199 == int32(0))
	goto L80
L94:
	;
	goto L93
L95:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	if v208 == int32(0) {
		v221 = v205
		goto L84
	} else {
		goto L96
	}
L96:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v185<<(uint(int32(2))%32)+v208)))
	if v215 == v217 {
		v185 = v185 + int32(1)
		goto L85
	} else {
		goto L97
	}
L97:
	;
	goto L86
L98:
	;
	v230 = v175
	goto L100
L99:
	;
	v230 = l7
	goto L100
L100:
	;
	v231 = v230
	goto L79
L101:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v238 = *(*float64)(unsafe.Add(mBase, uint32(v17)+16))
	v239 = *(*float64)(unsafe.Add(mBase, uint32(v17)+24))
	v240 = int32(0)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v244 == v240 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	if v324 == int32(0) {
		goto L20
	} else {
		goto L138
	}
L103:
	;
	v324 = int32(1)
	goto L102
L104:
	;
	goto L105
L105:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	if v248 <= int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v324 = int32(1)
	goto L102
L107:
	;
	goto L108
L108:
	;
	if v40 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v253 = int32(0)
	goto L111
L110:
	;
	v253 = l4
	goto L111
L111:
	;
	if v40 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v256 = int32(25)
	goto L114
L113:
	;
	v256 = int32(24)
	goto L114
L114:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v256))))
	v266 = v240
	goto L115
L115:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v269+v266<<(uint(int32(2))%32))))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+40))
	if v237 != v274 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v324 = v312
	goto L102
L117:
	;
	if v258 != 0 {
		goto L125
	} else {
		goto L126
	}
L118:
	;
	if v274 <= v237 {
		goto L117
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v278 = *(*float64)(unsafe.Add(mBase, uint32(v273)+56))
	if base.F64_le(v239, base.F64_mul(v278, float64(1.01))) == int32(0) {
		goto L117
	} else {
		goto L122
	}
L121:
	;
	v324 = int32(1)
	goto L102
L122:
	;
	v324 = int32(1)
	goto L102
L123:
	;
	goto L116
L124:
	;
	v306 = int32(1)
	v308 = v266 + v306
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	if v308 < v309 {
		v266 = v308
		goto L115
	} else {
		goto L137
	}
L125:
	;
	v285 = *(*float64)(unsafe.Add(mBase, uint32(v273)+48))
	if base.F64_gt(v238, base.F64_mul(v285, float64(1.01))) == int32(0) {
		goto L124
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v273)+16))
	if v292 != 0 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	goto L127
L129:
	;
	v295 = int32(0)
	goto L131
L130:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v273)+64))
	v295 = v294
	goto L131
L131:
	;
	v296 = F_compare_pathkeys(m, v253, v295)
	mBase = m.M
	if v296&int32(-3) != 0 {
		goto L124
	} else {
		goto L132
	}
L132:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v273)+16))
	if v299 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+4))
	v302 = v300
	goto L135
L134:
	;
	v302 = int32(0)
	goto L135
L135:
	;
	v303 = F_bms_equal(m, v40, v302)
	mBase = m.M
	if v303 != 0 {
		v312 = int32(0)
		goto L123
	} else {
		goto L136
	}
L136:
	;
	goto L124
L137:
	;
	v312 = v306
	goto L123
L138:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v17)+108))
	v329 = F_create_mergejoin_path(m, l0, l1, l8, v233, l9, l2, l3, v327, l4, v40, l5, v174, v231, v328)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L5
	} else {
		goto L139
	}
L139:
	;
	F_add_path(m, l1, v329)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L5
	} else {
		goto L140
	}
L140:
	;
	goto L1
L141:
	;
	goto L1
}
func F_try_nestloop_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 float64
	_ = v403
	var v404 float64
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 float64
	_ = v443
	var v450 float64
	_ = v450
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	v8 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(96)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v20 = v19
	goto L3
L2:
	;
	v20 = v8
	goto L3
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v23 = v22
	goto L6
L5:
	;
	v23 = v8
	goto L6
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	m.G0 = v16 + int32(96)
	return
L8:
	;
	v28 = F_bms_is_member(m, v27, v20)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+228))
	if v34 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	return
L12:
	;
	if v28 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	v32 = F_bms_is_member(m, v31, v23)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	if v32 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	goto L18
L17:
	;
	goto L18
L18:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v24)+228))
	if v38 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v42 = v41
	goto L21
L20:
	;
	v42 = v38
	goto L21
L21:
	;
	if v20 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	F_bms_free(m, v51)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L11
	} else {
		goto L187
	}
L23:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v206 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L24:
	;
	if v51 == int32(0) {
		goto L23
	} else {
		goto L31
	}
L25:
	;
	v45 = F_bms_copy(m, v23)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L11
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v47 = F_bms_union(m, v23, v20)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L11
	} else {
		goto L29
	}
L28:
	;
	v51 = v45
	goto L24
L29:
	;
	v49 = F_bms_del_members(m, v47, v42)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	v51 = v49
	goto L24
L31:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l6)+32))
	v55 = int32(0)
	if base.B2i32(v51 == v55)|base.B2i32(v54 == v55) != 0 {
		v100 = v55
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v100 != 0 {
		goto L23
	} else {
		goto L45
	}
L33:
	;
	goto L32
L34:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v65 < v66 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v68 = v65
	goto L37
L36:
	;
	v68 = v66
	goto L37
L37:
	;
	if v68 <= int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v71 = int32(1)
	goto L40
L39:
	;
	v71 = v68
	goto L40
L40:
	;
	v72 = int32(8)
	v77 = int32(0)
	goto L41
L41:
	;
	v84 = v77 << (uint(int32(2)) % 32)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v54+v72+v84)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v51+v72+v84)))
	v89 = v86 & v88
	v91 = base.B2i32(v89 != int32(0))
	if v89 != 0 {
		v100 = v91
		goto L33
	} else {
		goto L43
	}
L42:
	;
	v100 = v91
	goto L33
L43:
	;
	v93 = v77 + int32(1)
	if v93 != v71 {
		v77 = v93
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v101 = int32(0)
	if base.B2i32(v20 == v101)|base.B2i32(v42 == v101) != 0 {
		v146 = v101
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v146 == int32(0) {
		goto L22
	} else {
		goto L59
	}
L47:
	;
	goto L46
L48:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v111 < v112 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v114 = v111
	goto L51
L50:
	;
	v114 = v112
	goto L51
L51:
	;
	if v114 <= int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v117 = int32(1)
	goto L54
L53:
	;
	v117 = v114
	goto L54
L54:
	;
	v118 = int32(8)
	v123 = int32(0)
	goto L55
L55:
	;
	v130 = v123 << (uint(int32(2)) % 32)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v42+v118+v130)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v20+v118+v130)))
	v135 = v132 & v134
	v137 = base.B2i32(v135 != int32(0))
	if v135 != 0 {
		v146 = v137
		goto L47
	} else {
		goto L57
	}
L56:
	;
	v146 = v137
	goto L47
L57:
	;
	v139 = v123 + int32(1)
	if v139 != v117 {
		v123 = v139
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	if v20 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v203 == int32(0) {
		goto L22
	} else {
		goto L74
	}
L61:
	;
	v203 = int32(0)
	goto L60
L62:
	;
	goto L63
L63:
	;
	v156 = int32(1)
	if v42 == int32(0) {
		v193 = v156
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v203 = v193
	goto L60
L65:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v160 < v159 {
		v193 = v156
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v162 = int32(1)
	if v159 <= v162 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v165 = v162
	goto L69
L68:
	;
	v165 = v159
	goto L69
L69:
	;
	v166 = int32(8)
	v171 = int32(0)
	goto L70
L70:
	;
	v178 = v171 << (uint(int32(2)) % 32)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v20+v166+v178)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v42+v166+v178)))
	v185 = v180 & (v182 ^ int32(-1))
	v187 = base.B2i32(v185 != int32(0))
	if v185 != 0 {
		v193 = v187
		goto L64
	} else {
		goto L72
	}
L71:
	;
	v193 = v187
	goto L64
L72:
	;
	v189 = v171 + int32(1)
	if v189 != v165 {
		v171 = v189
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	goto L23
L75:
	;
	F_initial_cost_nestloop(m, l0, v16, l5, l2, l3, l6)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L11
	} else {
		goto L147
	}
L76:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v206)+4))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+228))
	v212 = int32(0)
	if base.B2i32(v209 == v212)|base.B2i32(v211 == v212) != 0 {
		v257 = v212
		goto L78
	} else {
		goto L79
	}
L77:
	;
	if v257 == int32(0) {
		goto L75
	} else {
		goto L90
	}
L78:
	;
	goto L77
L79:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	if v222 < v223 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v225 = v222
	goto L82
L81:
	;
	v225 = v223
	goto L82
L82:
	;
	if v225 <= int32(1) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v228 = int32(1)
	goto L85
L84:
	;
	v228 = v225
	goto L85
L85:
	;
	v229 = int32(8)
	v234 = int32(0)
	goto L86
L86:
	;
	v241 = v234 << (uint(int32(2)) % 32)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v211+v229+v241)))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v209+v229+v241)))
	v246 = v243 & v245
	v248 = base.B2i32(v246 != int32(0))
	if v246 != 0 {
		v257 = v248
		goto L78
	} else {
		goto L88
	}
L87:
	;
	v257 = v248
	goto L78
L88:
	;
	v250 = v234 + int32(1)
	if v250 != v228 {
		v234 = v250
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v263 = int32(1)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v264 == int32(0) {
		v391 = v263
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v397 == int32(0) {
		goto L22
	} else {
		goto L146
	}
L92:
	;
	v397 = v391
	goto L91
L93:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v260)+228))
	v269 = F_bms_overlap(m, v267, v268)
	mBase = m.M
	if v269 == int32(0) {
		v391 = v263
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v272 = int32(0)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	switch v273 - int32(279) {
	case 0, 1:
		goto L95
	default:
		v391 = v272
		goto L92
	case 3:
		goto L105
	case 4:
		goto L104
	case 5:
		goto L103
	case 9:
		goto L102
	case 10:
		goto L101
	case 11:
		goto L99
	case 14:
		goto L98
	case 15:
		goto L97
	case 17:
		goto L96
	case 19, 20, 21:
		goto L100
	}
L95:
	;
	v391 = int32(1)
	goto L92
L96:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v381 = F_path_is_reparameterizable_by_child(m, v380, v260)
	mBase = m.M
	if v381 == int32(0) {
		v391 = v272
		goto L92
	} else {
		goto L145
	}
L97:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v379 = F_path_is_reparameterizable_by_child(m, v378, v260)
	mBase = m.M
	if v379 != 0 {
		goto L95
	} else {
		goto L144
	}
L98:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v377 = F_path_is_reparameterizable_by_child(m, v376, v260)
	mBase = m.M
	if v377 != 0 {
		goto L95
	} else {
		goto L143
	}
L99:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v354 == int32(0) {
		goto L95
	} else {
		goto L135
	}
L100:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l3)+80))
	v349 = F_path_is_reparameterizable_by_child(m, v348, v260)
	mBase = m.M
	if v349 == int32(0) {
		v391 = v272
		goto L92
	} else {
		goto L133
	}
L101:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	if v326 == int32(0) {
		goto L95
	} else {
		goto L125
	}
L102:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v322 == int32(0) {
		goto L95
	} else {
		goto L123
	}
L103:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v300 == int32(0) {
		goto L95
	} else {
		goto L115
	}
L104:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v278 == int32(0) {
		goto L95
	} else {
		goto L107
	}
L105:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v277 = F_path_is_reparameterizable_by_child(m, v276, v260)
	mBase = m.M
	if v277 != 0 {
		goto L95
	} else {
		goto L106
	}
L106:
	;
	v391 = v272
	goto L92
L107:
	;
	v281 = int32(0)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	if v282 <= v281 {
		goto L95
	} else {
		goto L108
	}
L108:
	;
	v285 = v281
	goto L109
L109:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v289+v285<<(uint(int32(2))%32))))
	v294 = F_path_is_reparameterizable_by_child(m, v293, v260)
	mBase = m.M
	if v294 != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v397 = int32(0)
	goto L91
L111:
	;
	v296 = v285 + int32(1)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	if v296 < v297 {
		v285 = v296
		goto L109
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	goto L110
L114:
	;
	goto L95
L115:
	;
	v303 = int32(0)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	if v304 <= v303 {
		goto L95
	} else {
		goto L116
	}
L116:
	;
	v307 = v303
	goto L117
L117:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v300)+12))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v311+v307<<(uint(int32(2))%32))))
	v316 = F_path_is_reparameterizable_by_child(m, v315, v260)
	mBase = m.M
	if v316 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v397 = int32(0)
	goto L91
L119:
	;
	v318 = v307 + int32(1)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	if v318 < v319 {
		v307 = v318
		goto L117
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	goto L118
L122:
	;
	goto L95
L123:
	;
	v325 = F_path_is_reparameterizable_by_child(m, v322, v260)
	mBase = m.M
	if v325 != 0 {
		goto L95
	} else {
		goto L124
	}
L124:
	;
	v391 = v272
	goto L92
L125:
	;
	v329 = int32(0)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if v330 <= v329 {
		goto L95
	} else {
		goto L126
	}
L126:
	;
	v333 = v329
	goto L127
L127:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v326)+12))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v337+v333<<(uint(int32(2))%32))))
	v342 = F_path_is_reparameterizable_by_child(m, v341, v260)
	mBase = m.M
	if v342 != 0 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v397 = int32(0)
	goto L91
L129:
	;
	v344 = v333 + int32(1)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if v344 < v345 {
		v333 = v344
		goto L127
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	goto L128
L132:
	;
	goto L95
L133:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	v353 = F_path_is_reparameterizable_by_child(m, v352, v260)
	mBase = m.M
	if v353 != 0 {
		goto L95
	} else {
		goto L134
	}
L134:
	;
	v391 = v272
	goto L92
L135:
	;
	v357 = int32(0)
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v354)+4))
	if v358 <= v357 {
		goto L95
	} else {
		goto L136
	}
L136:
	;
	v361 = v357
	goto L137
L137:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v354)+12))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v365+v361<<(uint(int32(2))%32))))
	v370 = F_path_is_reparameterizable_by_child(m, v369, v260)
	mBase = m.M
	if v370 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v397 = int32(0)
	goto L91
L139:
	;
	v372 = v361 + int32(1)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v354)+4))
	if v372 < v373 {
		v361 = v372
		goto L137
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	goto L138
L142:
	;
	goto L95
L143:
	;
	v391 = v272
	goto L92
L144:
	;
	v391 = v272
	goto L92
L145:
	;
	goto L95
L146:
	;
	goto L75
L147:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v403 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
	v404 = *(*float64)(unsafe.Add(mBase, uint32(v16)+16))
	v405 = int32(0)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v409 == v405 {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	if v489 == int32(0) {
		goto L22
	} else {
		goto L184
	}
L149:
	;
	v489 = int32(1)
	goto L148
L150:
	;
	goto L151
L151:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v409)+4))
	if v413 <= int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v489 = int32(1)
	goto L148
L153:
	;
	goto L154
L154:
	;
	if v51 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v418 = int32(0)
	goto L157
L156:
	;
	v418 = l4
	goto L157
L157:
	;
	if v51 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v421 = int32(25)
	goto L160
L159:
	;
	v421 = int32(24)
	goto L160
L160:
	;
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v421))))
	v431 = v405
	goto L161
L161:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v409)+12))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v434+v431<<(uint(int32(2))%32))))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v438)+40))
	if v402 != v439 {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v489 = v477
	goto L148
L163:
	;
	if v423 != 0 {
		goto L171
	} else {
		goto L172
	}
L164:
	;
	if v439 <= v402 {
		goto L163
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v443 = *(*float64)(unsafe.Add(mBase, uint32(v438)+56))
	if base.F64_le(v404, base.F64_mul(v443, float64(1.01))) == int32(0) {
		goto L163
	} else {
		goto L168
	}
L167:
	;
	v489 = int32(1)
	goto L148
L168:
	;
	v489 = int32(1)
	goto L148
L169:
	;
	goto L162
L170:
	;
	v471 = int32(1)
	v473 = v431 + v471
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v409)+4))
	if v473 < v474 {
		v431 = v473
		goto L161
	} else {
		goto L183
	}
L171:
	;
	v450 = *(*float64)(unsafe.Add(mBase, uint32(v438)+48))
	if base.F64_gt(v403, base.F64_mul(v450, float64(1.01))) == int32(0) {
		goto L170
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v438)+16))
	if v457 != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	goto L173
L175:
	;
	v460 = int32(0)
	goto L177
L176:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v438)+64))
	v460 = v459
	goto L177
L177:
	;
	v461 = F_compare_pathkeys(m, v418, v460)
	mBase = m.M
	if v461&int32(-3) != 0 {
		goto L170
	} else {
		goto L178
	}
L178:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v438)+16))
	if v464 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v464)+4))
	v467 = v465
	goto L181
L180:
	;
	v467 = int32(0)
	goto L181
L181:
	;
	v468 = F_bms_equal(m, v51, v467)
	mBase = m.M
	if v468 != 0 {
		v477 = int32(0)
		goto L169
	} else {
		goto L182
	}
L182:
	;
	goto L170
L183:
	;
	v477 = v471
	goto L169
L184:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v493 = F_create_nestloop_path(m, l0, l1, l5, v16, l6, l2, l3, v492, l4, v51)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L11
	} else {
		goto L185
	}
L185:
	;
	F_add_path(m, l1, v493)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L11
	} else {
		goto L186
	}
L186:
	;
	goto L7
L187:
	;
	goto L7
}
func F_tstoreReceiveSlot_tupmap(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v6 = F_execute_attr_map_slot(m, v4, l0, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
		F_tuplestore_puttupleslot(m, v10, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return int32(1)
		}
	}
}
func F_tsvectorrecv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	v2 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_pq_getmsgint(m, v19, int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L5
	} else {
		goto L95
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L5
	} else {
		goto L92
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L5
	} else {
		goto L89
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L5
	} else {
		goto L86
	}
L5:
	;
	return int32(0)
L6:
	;
	if base.Ui32(v21) < base.Ui32(int32(268435456)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v28 = v21 << (uint(int32(2)) % 32)
	v30 = v28 + int32(8)
	v32 = v30 << (uint(int32(1)) % 32)
	v33 = F_palloc0(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L5
	} else {
		goto L83
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v21
	if v21 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(32)
	return v33
L12:
	;
	goto L13
L13:
	;
	v43 = v32
	v44 = v33
	v46 = v2
	v54 = v2
	v55 = v2
	goto L14
L14:
	;
	v61 = F_pq_getmsgstring(m, v19)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = (v315 + v30) << (uint(int32(2)) % 32)
	if v199&int32(1) != 0 {
		goto L79
	} else {
		goto L80
	}
L16:
	;
	v64 = F_pq_getmsgint(m, v19, int32(2))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v66 = F_strlen(m, v61)
	mBase = m.M
	if base.Ui32(int32(2048)) <= base.Ui32(v66) {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	if int32(_a_F_tsvectorrecv_0) <= v46 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v72 = v64 & int32(_a_F_tsvectorrecv_1)
	if base.Ui32(int32(256)) < base.Ui32(v72) {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v75 = v46 + v66
	v76 = int32(1)
	v79 = (v75 + v76) & int32(-2)
	v81 = v72 << (uint(v76) % 32)
	v83 = v79 + (v28 + int32(10) + v81)
	if base.Ui32(v43) <= base.Ui32(v83) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v85 = v43
	v86 = v44
	goto L24
L22:
	;
	v108 = v43
	v109 = v44
	goto L23
L23:
	;
	v127 = v109 + int32(8)
	v130 = v127 + v54<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v66<<(uint(int32(1))%32) | base.B2i32(v72 != int32(0)) | v46<<(uint(int32(12))%32)
	if v66 != 0 {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v104 = v85 << (uint(int32(1)) % 32)
	v105 = F_repalloc(m, v86, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L26
	}
L25:
	;
	v108 = v104
	v109 = v105
	goto L23
L26:
	;
	if base.Ui32(v104) <= base.Ui32(v83) {
		v85 = v104
		v86 = v105
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	base.MemoryCopy(m, v127+v140<<(uint(int32(2))%32)+v46, v61, v66)
	goto L30
L29:
	;
	goto L30
L30:
	;
	if v54 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v149 = v127 + v146<<(uint(int32(2))%32)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v151 = int32(12)
	v154 = int32(1)
	v156 = int32(2047)
	v157 = int32(base.Ui32(v150)>>(uint(v154)%32)) & v156
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v130-int32(4))))
	v167 = int32(base.Ui32(v160)>>(uint(v154)%32)) & v156
	if v157 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v199 = v55
	goto L33
L33:
	;
	if v72 != 0 {
		goto L62
	} else {
		goto L63
	}
L34:
	;
	v199 = base.B2i32(v193 <= int32(0)) | v55
	goto L33
L35:
	;
	goto L39
L36:
	;
	goto L37
L37:
	;
	if v167 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	goto L40
L40:
	;
	v173 = int32(0)
	if v173 < v167 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v176 = int32(-1)
	goto L43
L42:
	;
	v176 = v173
	goto L43
L43:
	;
	v193 = v176
	goto L34
L44:
	;
	v193 = base.B2i32(int32(0) < v157)
	goto L34
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(v157) < base.Ui32(v167) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v182 = v157
	goto L49
L48:
	;
	v182 = v167
	goto L49
L49:
	;
	v183 = F_memcmp(m, v149+int32(base.Ui32(v150)>>(uint(v151)%32)), v149+int32(base.Ui32(v160)>>(uint(v151)%32)), v182)
	mBase = m.M
	goto L52
L50:
	;
	v193 = v191
	goto L34
L52:
	;
	goto L53
L53:
	;
	if v183 != 0 {
		v191 = v183
		goto L50
	} else {
		goto L55
	}
L55:
	;
	if v157 == v167 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v193 = int32(0)
	goto L34
L57:
	;
	goto L58
L58:
	;
	if v157 < v167 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v190 = int32(-1)
	goto L61
L60:
	;
	v190 = int32(1)
	goto L61
L61:
	;
	v191 = v190
	goto L50
L62:
	;
	if v75 == v79 {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v315 = v75
	goto L64
L64:
	;
	v317 = v54 + int32(1)
	if v317 != v21 {
		v43 = v108
		v44 = v109
		v46 = v315
		v54 = v317
		v55 = v199
		goto L14
	} else {
		goto L78
	}
L65:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v210 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v127+v209<<(uint(v210)%32)+v208))) = uint16(v64)
	v215 = int32(1)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v232 = v127 + v216<<(uint(v210)%32) + (int32(base.Ui32(v220)>>(uint(int32(12))%32))+int32(base.Ui32(v220)>>(uint(v215)%32))&int32(2047)+v215)&int32(_a_F_tsvectorrecv_2)
	v234 = F_pq_getmsgint(m, v19, v210)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L5
	} else {
		goto L69
	}
L66:
	;
	v208 = v75
	goto L65
L67:
	;
	goto L68
L68:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v206 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v127+v201<<(uint(int32(2))%32)+v75))) = uint8(v206)
	v208 = v79
	goto L65
L69:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v232)+2)) = uint16(v234)
	if v72 != int32(1) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v243 = v215
	goto L73
L71:
	;
	goto L72
L72:
	;
	v315 = v208 + v81 + int32(2)
	goto L64
L73:
	;
	v260 = v243 << (uint(int32(1)) % 32)
	v263 = F_pq_getmsgint(m, v19, int32(2))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L5
	} else {
		goto L75
	}
L74:
	;
	goto L72
L75:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v232+int32(2)+v260))) = uint16(v263)
	v266 = int32(_a_F_tsvectorrecv_3)
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v232+v260))))
	if base.Ui32(v263&v266) <= base.Ui32(v269&v266) {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v274 = v243 + int32(1)
	if v274 != v72 {
		v243 = v274
		goto L73
	} else {
		goto L77
	}
L77:
	;
	goto L74
L78:
	;
	goto L15
L79:
	;
	v326 = v109 + int32(8)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	F_qsort_arg(m, v326, v327, int32(4), int32(1519), v326+v327<<(uint(int32(2))%32))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L5
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	return v109
L82:
	;
	goto L81
L83:
	;
	F_errmsg_internal(m, int32(_a_F_tsvectorrecv_4), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_tsvectorrecv_5), int32(462), int32(_a_F_tsvectorrecv_6))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errmsg_internal(m, int32(_a_F_tsvectorrecv_7), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_tsvectorrecv_5), int32(484), int32(_a_F_tsvectorrecv_6))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	F_errmsg_internal(m, int32(_a_F_tsvectorrecv_8), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_tsvectorrecv_5), int32(487), int32(_a_F_tsvectorrecv_6))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errmsg_internal(m, int32(_a_F_tsvectorrecv_9), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L5
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_tsvectorrecv_5), int32(490), int32(_a_F_tsvectorrecv_6))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	F_errmsg_internal(m, int32(_a_F_tsvectorrecv_10), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_tsvectorrecv_5), int32(541), int32(_a_F_tsvectorrecv_6))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tuplehash_lookup_hash_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v92 int32
	_ = v92
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = v15 & l1
	v19 = v14 + v16*int32(12)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v92
L2:
	;
	v23 = v15
	v24 = v14
	v25 = v19
	v28 = v16
	goto L5
L3:
	;
	goto L4
L4:
	;
	v92 = int32(0)
	goto L1
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v30 == l1 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+36))
	v37 = F_ExecStoreMinimalTuple(m, v34, v35, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v67 = v23
	v68 = v24
	goto L9
L9:
	;
	v73 = (v28 + int32(1)) & v67
	v76 = v68 + v73*int32(12)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v77 != 0 {
		v23 = v67
		v24 = v68
		v25 = v76
		v28 = v73
		goto L5
	} else {
		goto L19
	}
L10:
	;
	return int32(0)
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v41
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v32)+48))
	if v44 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	F_MemoryContextReset(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v50 = int32(_a_F_tuplehash_lookup_hash_internal_0)
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_tuplehash_lookup_hash_internal[0]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplehash_lookup_hash_internal[0])) = v53
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	v58 = m.T0[v57].(func(*base.Module, int32, int32, int32) int32)(m, v44, v33, v12+int32(15))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	v92 = v25
	goto L1
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tuplehash_lookup_hash_internal[0])) = v51
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	F_MemoryContextReset(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	if v58 != 0 {
		v92 = v25
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v67 = v65
	v68 = v66
	goto L9
L19:
	;
	goto L6
}
func F_typenameTypeId(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v4 = F_typenameType(m, l0, l1, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v8+v9)))
		F_ReleaseCatCache(m, v4)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return v11
		}
	}
}
