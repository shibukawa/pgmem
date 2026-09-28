package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TrackNewBufferPin(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	v4 = int32(_a_F_TrackNewBufferPin_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_TrackNewBufferPin[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v5<<(uint(int32(2))%32))+uint32(_c_F_TrackNewBufferPin[1]))) = l0
	v12 = v5 << (uint(int32(4)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_TrackNewBufferPin[2]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_TrackNewBufferPin[3]))) = l0
	*(*int32)(unsafe.Add(mBase, _c_F_TrackNewBufferPin[0])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_TrackNewBufferPin[4])) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_TrackNewBufferPin[5]))) = int32(1)
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_TrackNewBufferPin[6]))
	F_ResourceOwnerRemember(m, v30, base.I64_extend_i32_s(l0), int32(_a_F_TrackNewBufferPin_1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return
	} else {
		return
	}
}
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
								v51 = F_MakeSingleTupleTableSlot(m, v48, int32(_a_F_TransitionTableAddTuple_1))
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
								v51 = F_MakeSingleTupleTableSlot(m, v48, int32(_a_F_TransitionTableAddTuple_1))
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
								v51 = F_MakeSingleTupleTableSlot(m, v48, int32(_a_F_TransitionTableAddTuple_1))
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
	F_errfinish(m, int32(_a_F__tarWriteHeader_5), int32(2052), int32(_a_F__tarWriteHeader_6))
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
	F_errfinish(m, int32(_a_F__tarWriteHeader_5), int32(2059), int32(_a_F__tarWriteHeader_6))
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
	F_errfinish(m, int32(_a_F__tarWriteHeader_5), int32(2062), int32(_a_F__tarWriteHeader_6))
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
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v361 int32
	_ = v361
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v665 int32
	_ = v665
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v917 int32
	_ = v917
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1033 int32
	_ = v1033
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1097 int32
	_ = v1097
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1254 int32
	_ = v1254
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1309 int32
	_ = v1309
	var v1312 int32
	_ = v1312
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1323 int32
	_ = v1323
	var v1328 int32
	_ = v1328
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1349 int32
	_ = v1349
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1386 int32
	_ = v1386
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1442 int32
	_ = v1442
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
	var v1520 int32
	_ = v1520
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1533 int32
	_ = v1533
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1549 int32
	_ = v1549
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1584 int32
	_ = v1584
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1597 int32
	_ = v1597
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1634 int32
	_ = v1634
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1650 int32
	_ = v1650
	var v1656 int32
	_ = v1656
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v2)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = F_r_fix_ending(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v1656
L2:
	;
	return int32(0)
L3:
	;
	if v13 < int32(0) {
		v1656 = v13
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12
	v20 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v21-int32(4))))
	if v29 == v20 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v103 < int32(5) {
		v1656 = v20
		goto L1
	} else {
		goto L21
	}
L6:
	;
	v103 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v34 = v29 & int32(3)
	if base.Ui32(v29) < base.Ui32(int32(4)) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v103 = v92
	goto L5
L10:
	;
	v76 = v70
	v77 = v71
	v81 = v20
	goto L18
L11:
	;
	v70 = v21
	v71 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v41 = v21
	v42 = int32(0)
	v45 = v20
	goto L14
L14:
	;
	v47 = int32(*(*int8)(unsafe.Add(mBase, uint32(v41))))
	v48 = int32(-65)
	v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(v41)+1)))
	v55 = int32(*(*int8)(unsafe.Add(mBase, uint32(v41)+2)))
	v59 = int32(*(*int8)(unsafe.Add(mBase, uint32(v41)+3)))
	v62 = v42 + base.B2i32(v48 < v47) + base.B2i32(v48 < v51) + base.B2i32(v48 < v55) + base.B2i32(v48 < v59)
	v63 = int32(4)
	v64 = v41 + v63
	v66 = v45 + v63
	if v66 != v29&int32(-4) {
		v41 = v64
		v42 = v62
		v45 = v66
		goto L14
	} else {
		goto L16
	}
L15:
	;
	if v34 == int32(0) {
		v92 = v62
		goto L9
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	v70 = v64
	v71 = v62
	goto L10
L18:
	;
	v82 = int32(*(*int8)(unsafe.Add(mBase, uint32(v76))))
	v85 = v77 + base.B2i32(int32(-65) < v82)
	v86 = int32(1)
	v89 = v81 + v86
	if v89 != v34 {
		v76 = v76 + v86
		v77 = v85
		v81 = v89
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v92 = v85
	goto L9
L20:
	;
	goto L19
L21:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v106
	v108 = int32(3)
	v110 = int32(0)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v112-v106 < v108 {
		v122 = v110
		goto L24
	} else {
		goto L25
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106
	v162 = v106 + int32(2)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v163 <= v162 {
		goto L38
	} else {
		goto L39
	}
L23:
	;
	if v122 == int32(0) {
		goto L22
	} else {
		goto L27
	}
L24:
	;
	goto L23
L25:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v118 = F_memcmp(m, v116+v106, int32(_a_F_tamil_UTF_8_stem_0), v108)
	mBase = m.M
	if v118 != 0 {
		v122 = v110
		goto L24
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108 + v106
	v122 = int32(1)
	goto L24
L27:
	;
	v128 = F_find_among(m, l0, int32(_a_F_tamil_UTF_8_stem_1), int32(10), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	if v128 == int32(0) {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v132 = int32(3)
	v134 = int32(0)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v136-v137 < v132 {
		v146 = v134
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v146 == int32(0) {
		goto L22
	} else {
		goto L34
	}
L31:
	;
	goto L30
L32:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v142 = F_memcmp(m, v140+v137, int32(_a_F_tamil_UTF_8_stem_2), v132)
	mBase = m.M
	if v142 != 0 {
		v146 = v134
		goto L31
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v132 + v137
	v146 = int32(1)
	goto L31
L34:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v149
	v151 = F_slice_del(m, l0)
	mBase = m.M
	if v151 < int32(0) {
		v1656 = v151
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v154 = F_r_fix_va_start(m, l0)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	if v154 < int32(0) {
		v1656 = v154
		goto L1
	} else {
		goto L37
	}
L37:
	;
	goto L22
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106
	v222 = int32(0)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v223-int32(4))))
	if v231 == v222 {
		goto L55
	} else {
		goto L56
	}
L39:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+v162))))
	if base.B2i32(v167&int32(224) != int32(128))|base.B2i32(int32(1)<<(uint(v167)%32)&int32(672) == int32(0)) != 0 {
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v182 = F_find_among(m, l0, int32(_a_F_tamil_UTF_8_stem_3), int32(3), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	if v182 == int32(0) {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v189 = F_find_among(m, l0, int32(_a_F_tamil_UTF_8_stem_4), int32(10), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	if v189 == int32(0) {
		goto L38
	} else {
		goto L44
	}
L44:
	;
	v193 = int32(3)
	v195 = int32(0)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v197-v198 < v193 {
		v207 = v195
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v207 == int32(0) {
		goto L38
	} else {
		goto L49
	}
L46:
	;
	goto L45
L47:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v203 = F_memcmp(m, v201+v198, int32(_a_F_tamil_UTF_8_stem_5), v193)
	mBase = m.M
	if v203 != 0 {
		v207 = v195
		goto L46
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v193 + v198
	v207 = int32(1)
	goto L46
L49:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v210
	v212 = F_slice_del(m, l0)
	mBase = m.M
	if v212 < int32(0) {
		v1656 = v212
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v215 = F_r_fix_va_start(m, l0)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	if v215 < int32(0) {
		v1656 = v215
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L38
L53:
	;
	if v361 < int32(0) {
		v1656 = v361
		goto L1
	} else {
		goto L96
	}
L54:
	;
	if v305 < int32(5) {
		v361 = v222
		goto L53
	} else {
		goto L70
	}
L55:
	;
	v305 = int32(0)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v236 = v231 & int32(3)
	if base.Ui32(v231) < base.Ui32(int32(4)) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v305 = v294
	goto L54
L59:
	;
	v278 = v272
	v279 = v273
	v283 = v222
	goto L67
L60:
	;
	v272 = v223
	v273 = int32(0)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v243 = v223
	v244 = int32(0)
	v247 = v222
	goto L63
L63:
	;
	v249 = int32(*(*int8)(unsafe.Add(mBase, uint32(v243))))
	v250 = int32(-65)
	v253 = int32(*(*int8)(unsafe.Add(mBase, uint32(v243)+1)))
	v257 = int32(*(*int8)(unsafe.Add(mBase, uint32(v243)+2)))
	v261 = int32(*(*int8)(unsafe.Add(mBase, uint32(v243)+3)))
	v264 = v244 + base.B2i32(v250 < v249) + base.B2i32(v250 < v253) + base.B2i32(v250 < v257) + base.B2i32(v250 < v261)
	v265 = int32(4)
	v266 = v243 + v265
	v268 = v247 + v265
	if v268 != v231&int32(-4) {
		v243 = v266
		v244 = v264
		v247 = v268
		goto L63
	} else {
		goto L65
	}
L64:
	;
	if v236 == int32(0) {
		v294 = v264
		goto L58
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	v272 = v266
	v273 = v264
	goto L59
L67:
	;
	v284 = int32(*(*int8)(unsafe.Add(mBase, uint32(v278))))
	v287 = v279 + base.B2i32(int32(-65) < v284)
	v288 = int32(1)
	v291 = v283 + v288
	if v291 != v236 {
		v278 = v278 + v288
		v279 = v287
		v283 = v291
		goto L67
	} else {
		goto L69
	}
L68:
	;
	v294 = v287
	goto L58
L69:
	;
	goto L68
L70:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v308
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v310
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v310
	v316 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_6), int32(3), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	if v316 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v318
	v322 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_7))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L2
	} else {
		goto L75
	}
L73:
	;
	v326 = v222
	goto L74
L74:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v327
	v330 = v326
	goto L78
L75:
	;
	if v322 < int32(0) {
		v361 = v322
		goto L53
	} else {
		goto L76
	}
L76:
	;
	v326 = v322
	goto L74
L77:
	;
	v361 = int32(1)
	goto L53
L78:
	;
	v338 = F_r_fix_ending(m, l0)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L2
	} else {
		goto L83
	}
L79:
	;
	if v346 == int32(0) {
		goto L77
	} else {
		goto L94
	}
L80:
	;
	if v338 < int32(0) {
		goto L87
	} else {
		goto L88
	}
L81:
	;
	v346 = int32(2)
	goto L80
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v327
	goto L77
L83:
	;
	v341 = int32(base.Ui32(v338) >> (uint(int32(31)) % 32))
	if v338 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v343 = v341
	goto L86
L85:
	;
	v343 = int32(4)
	goto L86
L86:
	;
	switch v343 {
	case 0:
		goto L81
	default:
		v346 = v341
		goto L80
	case 4:
		goto L82
	}
L87:
	;
	v349 = v338
	goto L89
L88:
	;
	v349 = v330
	goto L89
L89:
	;
	if v338 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v350 = v349
	goto L92
L91:
	;
	v350 = v330
	goto L92
L92:
	;
	if v346 == int32(2) {
		v330 = v350
		goto L78
	} else {
		goto L93
	}
L93:
	;
	goto L79
L94:
	;
	if v350 < int32(0) {
		v361 = v350
		goto L53
	} else {
		goto L95
	}
L95:
	;
	goto L77
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106
	v372 = int32(0)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v373-int32(4))))
	if v381 == v372 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	if v497 < int32(0) {
		v1656 = v497
		goto L1
	} else {
		goto L124
	}
L98:
	;
	if v455 < int32(5) {
		v497 = v372
		goto L97
	} else {
		goto L114
	}
L99:
	;
	v455 = int32(0)
	goto L98
L100:
	;
	goto L101
L101:
	;
	v386 = v381 & int32(3)
	if base.Ui32(v381) < base.Ui32(int32(4)) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v455 = v444
	goto L98
L103:
	;
	v428 = v422
	v429 = v423
	v433 = v372
	goto L111
L104:
	;
	v422 = v373
	v423 = int32(0)
	goto L103
L105:
	;
	goto L106
L106:
	;
	v393 = v373
	v394 = int32(0)
	v397 = v372
	goto L107
L107:
	;
	v399 = int32(*(*int8)(unsafe.Add(mBase, uint32(v393))))
	v400 = int32(-65)
	v403 = int32(*(*int8)(unsafe.Add(mBase, uint32(v393)+1)))
	v407 = int32(*(*int8)(unsafe.Add(mBase, uint32(v393)+2)))
	v411 = int32(*(*int8)(unsafe.Add(mBase, uint32(v393)+3)))
	v414 = v394 + base.B2i32(v400 < v399) + base.B2i32(v400 < v403) + base.B2i32(v400 < v407) + base.B2i32(v400 < v411)
	v415 = int32(4)
	v416 = v393 + v415
	v418 = v397 + v415
	if v418 != v381&int32(-4) {
		v393 = v416
		v394 = v414
		v397 = v418
		goto L107
	} else {
		goto L109
	}
L108:
	;
	if v386 == int32(0) {
		v444 = v414
		goto L102
	} else {
		goto L110
	}
L109:
	;
	goto L108
L110:
	;
	v422 = v416
	v423 = v414
	goto L103
L111:
	;
	v434 = int32(*(*int8)(unsafe.Add(mBase, uint32(v428))))
	v437 = v429 + base.B2i32(int32(-65) < v434)
	v438 = int32(1)
	v441 = v433 + v438
	if v441 != v386 {
		v428 = v428 + v438
		v429 = v437
		v433 = v441
		goto L111
	} else {
		goto L113
	}
L112:
	;
	v444 = v437
	goto L102
L113:
	;
	goto L112
L114:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v458
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v460
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460
	v463 = int32(9)
	v465 = int32(0)
	if v460-v458 < v463 {
		v478 = v465
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if v478 == int32(0) {
		v497 = v372
		goto L97
	} else {
		goto L119
	}
L116:
	;
	goto L115
L117:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v474 = F_memcmp(m, v471+v460-v463, int32(_a_F_tamil_UTF_8_stem_8), v463)
	mBase = m.M
	if v474 != 0 {
		v478 = v465
		goto L116
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v460 - v463
	v478 = int32(1)
	goto L116
L119:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v481
	v485 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_9))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L2
	} else {
		goto L120
	}
L120:
	;
	if v485 < int32(0) {
		v497 = v485
		goto L97
	} else {
		goto L121
	}
L121:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v489
	v491 = F_r_fix_ending(m, l0)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	if v491 < int32(0) {
		v497 = v491
		goto L97
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v489
	v497 = int32(1)
	goto L97
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106
	v502 = int32(0)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v503-int32(4))))
	if v511 == v502 {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	if v665 < int32(0) {
		v1656 = v665
		goto L1
	} else {
		goto L175
	}
L126:
	;
	if v585 < int32(5) {
		v665 = v502
		goto L125
	} else {
		goto L142
	}
L127:
	;
	v585 = int32(0)
	goto L126
L128:
	;
	goto L129
L129:
	;
	v516 = v511 & int32(3)
	if base.Ui32(v511) < base.Ui32(int32(4)) {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	v585 = v574
	goto L126
L131:
	;
	v558 = v552
	v559 = v553
	v563 = v502
	goto L139
L132:
	;
	v552 = v503
	v553 = int32(0)
	goto L131
L133:
	;
	goto L134
L134:
	;
	v523 = v503
	v524 = int32(0)
	v527 = v502
	goto L135
L135:
	;
	v529 = int32(*(*int8)(unsafe.Add(mBase, uint32(v523))))
	v530 = int32(-65)
	v533 = int32(*(*int8)(unsafe.Add(mBase, uint32(v523)+1)))
	v537 = int32(*(*int8)(unsafe.Add(mBase, uint32(v523)+2)))
	v541 = int32(*(*int8)(unsafe.Add(mBase, uint32(v523)+3)))
	v544 = v524 + base.B2i32(v530 < v529) + base.B2i32(v530 < v533) + base.B2i32(v530 < v537) + base.B2i32(v530 < v541)
	v545 = int32(4)
	v546 = v523 + v545
	v548 = v527 + v545
	if v548 != v511&int32(-4) {
		v523 = v546
		v524 = v544
		v527 = v548
		goto L135
	} else {
		goto L137
	}
L136:
	;
	if v516 == int32(0) {
		v574 = v544
		goto L130
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	v552 = v546
	v553 = v544
	goto L131
L139:
	;
	v564 = int32(*(*int8)(unsafe.Add(mBase, uint32(v558))))
	v567 = v559 + base.B2i32(int32(-65) < v564)
	v568 = int32(1)
	v571 = v563 + v568
	if v571 != v516 {
		v558 = v558 + v568
		v559 = v567
		v563 = v571
		goto L139
	} else {
		goto L141
	}
L140:
	;
	v574 = v567
	goto L130
L141:
	;
	goto L140
L142:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v588
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v590
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v590
	v596 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_10), int32(26), int32(0))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L2
	} else {
		goto L143
	}
L143:
	;
	if v596 == int32(0) {
		v665 = v502
		goto L125
	} else {
		goto L144
	}
L144:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v600
	switch v596 - int32(1) {
	case 0:
		goto L148
	case 1:
		goto L147
	case 2:
		goto L146
	default:
		v629 = v502
		goto L145
	}
L145:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v631
	v634 = v629
	goto L157
L146:
	;
	v626 = F_slice_del(m, l0)
	mBase = m.M
	if v626 < int32(0) {
		v665 = v626
		goto L125
	} else {
		goto L155
	}
L147:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v614 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_11), int32(8), int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L2
	} else {
		goto L151
	}
L148:
	;
	v606 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_12))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L2
	} else {
		goto L149
	}
L149:
	;
	if int32(0) <= v606 {
		v629 = v606
		goto L145
	} else {
		goto L150
	}
L150:
	;
	v665 = v606
	goto L125
L151:
	;
	if v614 != 0 {
		v665 = v502
		goto L125
	} else {
		goto L152
	}
L152:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v616 + (v600 - v610)
	v622 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_13))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L2
	} else {
		goto L153
	}
L153:
	;
	if int32(0) <= v622 {
		v629 = v622
		goto L145
	} else {
		goto L154
	}
L154:
	;
	v665 = v622
	goto L125
L155:
	;
	v629 = v626
	goto L145
L156:
	;
	v665 = int32(1)
	goto L125
L157:
	;
	v642 = F_r_fix_ending(m, l0)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L2
	} else {
		goto L162
	}
L158:
	;
	if v650 == int32(0) {
		goto L156
	} else {
		goto L173
	}
L159:
	;
	if v642 < int32(0) {
		goto L166
	} else {
		goto L167
	}
L160:
	;
	v650 = int32(2)
	goto L159
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v631
	goto L156
L162:
	;
	v645 = int32(base.Ui32(v642) >> (uint(int32(31)) % 32))
	if v642 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v647 = v645
	goto L165
L164:
	;
	v647 = int32(4)
	goto L165
L165:
	;
	switch v647 {
	case 0:
		goto L160
	default:
		v650 = v645
		goto L159
	case 4:
		goto L161
	}
L166:
	;
	v653 = v642
	goto L168
L167:
	;
	v653 = v634
	goto L168
L168:
	;
	if v642 != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v654 = v653
	goto L171
L170:
	;
	v654 = v634
	goto L171
L171:
	;
	if v650 == int32(2) {
		v634 = v654
		goto L157
	} else {
		goto L172
	}
L172:
	;
	goto L158
L173:
	;
	if v654 < int32(0) {
		v665 = v654
		goto L125
	} else {
		goto L174
	}
L174:
	;
	goto L156
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106
	v676 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v676)
	v679 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v679-int32(4))))
	if v687 == v676 {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	if v1097 < int32(0) {
		v1656 = v1097
		goto L1
	} else {
		goto L291
	}
L177:
	;
	if v761 < int32(5) {
		v1097 = v676
		goto L176
	} else {
		goto L193
	}
L178:
	;
	v761 = int32(0)
	goto L177
L179:
	;
	goto L180
L180:
	;
	v692 = v687 & int32(3)
	if base.Ui32(v687) < base.Ui32(int32(4)) {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	v761 = v750
	goto L177
L182:
	;
	v734 = v728
	v735 = v729
	v739 = v676
	goto L190
L183:
	;
	v728 = v679
	v729 = int32(0)
	goto L182
L184:
	;
	goto L185
L185:
	;
	v699 = v679
	v700 = int32(0)
	v703 = v676
	goto L186
L186:
	;
	v705 = int32(*(*int8)(unsafe.Add(mBase, uint32(v699))))
	v706 = int32(-65)
	v709 = int32(*(*int8)(unsafe.Add(mBase, uint32(v699)+1)))
	v713 = int32(*(*int8)(unsafe.Add(mBase, uint32(v699)+2)))
	v717 = int32(*(*int8)(unsafe.Add(mBase, uint32(v699)+3)))
	v720 = v700 + base.B2i32(v706 < v705) + base.B2i32(v706 < v709) + base.B2i32(v706 < v713) + base.B2i32(v706 < v717)
	v721 = int32(4)
	v722 = v699 + v721
	v724 = v703 + v721
	if v724 != v687&int32(-4) {
		v699 = v722
		v700 = v720
		v703 = v724
		goto L186
	} else {
		goto L188
	}
L187:
	;
	if v692 == int32(0) {
		v750 = v720
		goto L181
	} else {
		goto L189
	}
L188:
	;
	goto L187
L189:
	;
	v728 = v722
	v729 = v720
	goto L182
L190:
	;
	v740 = int32(*(*int8)(unsafe.Add(mBase, uint32(v734))))
	v743 = v735 + base.B2i32(int32(-65) < v740)
	v744 = int32(1)
	v747 = v739 + v744
	if v747 != v692 {
		v734 = v734 + v744
		v735 = v743
		v739 = v747
		goto L190
	} else {
		goto L192
	}
L191:
	;
	v750 = v743
	goto L181
L192:
	;
	goto L191
L193:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v764
	v766 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v766
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v766
	if v766-int32(2) <= v764 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v1033 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v1033)
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1035
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1035
	v1038 = int32(9)
	v1040 = int32(0)
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1035-v1043 < v1038 {
		v1053 = v1040
		goto L264
	} else {
		goto L265
	}
L195:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v963
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v963
	v966 = int32(3)
	v968 = int32(0)
	v971 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v963-v971 < v966 {
		v981 = v968
		goto L246
	} else {
		goto L247
	}
L196:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v774 = int32(1)
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v772+v766-v774))))
	if base.B2i32(v776&int32(224) != int32(128))|base.B2i32(v774<<(uint(v776)%32)&int32(-2147475197) == int32(0)) != 0 {
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v791 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_14), int32(22), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L2
	} else {
		goto L198
	}
L198:
	;
	if v791 == int32(0) {
		goto L195
	} else {
		goto L199
	}
L199:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v795
	switch v791 - int32(1) {
	case 0:
		goto L206
	case 1:
		goto L205
	case 2:
		goto L204
	case 3:
		goto L203
	case 4:
		goto L202
	case 5:
		goto L201
	case 6:
		goto L200
	default:
		goto L194
	}
L200:
	;
	v957 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_15))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L2
	} else {
		goto L243
	}
L201:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v946 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_16), int32(8), int32(0))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L2
	} else {
		goto L240
	}
L202:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v930 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_17), int32(8), int32(0))
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L2
	} else {
		goto L236
	}
L203:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v836 = int32(0)
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v835-int32(4))))
	if v843 == v836 {
		goto L218
	} else {
		goto L219
	}
L204:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v809 = int32(3)
	v811 = int32(0)
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v813-v814 < v809 {
		v824 = v811
		goto L211
	} else {
		goto L212
	}
L205:
	;
	v804 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_18))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L2
	} else {
		goto L208
	}
L206:
	;
	v799 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v799 {
		goto L194
	} else {
		goto L207
	}
L207:
	;
	v1097 = v799
	goto L176
L208:
	;
	if int32(0) <= v804 {
		goto L194
	} else {
		goto L209
	}
L209:
	;
	v1097 = v804
	goto L176
L210:
	;
	if v824 != 0 {
		goto L195
	} else {
		goto L214
	}
L211:
	;
	goto L210
L212:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v820 = F_memcmp(m, v817+v813-v809, int32(_a_F_tamil_UTF_8_stem_19), v809)
	mBase = m.M
	if v820 != 0 {
		v824 = v811
		goto L211
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v813 - v809
	v824 = int32(1)
	goto L211
L214:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v825 + (v795 - v808)
	v831 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_20))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L2
	} else {
		goto L215
	}
L215:
	;
	if int32(0) <= v831 {
		goto L194
	} else {
		goto L216
	}
L216:
	;
	v1097 = v831
	goto L176
L217:
	;
	if v917 < int32(7) {
		goto L195
	} else {
		goto L233
	}
L218:
	;
	v917 = int32(0)
	goto L217
L219:
	;
	goto L220
L220:
	;
	v848 = v843 & int32(3)
	if base.Ui32(v843) < base.Ui32(int32(4)) {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	v917 = v906
	goto L217
L222:
	;
	v890 = v884
	v891 = v885
	v895 = v836
	goto L230
L223:
	;
	v884 = v835
	v885 = int32(0)
	goto L222
L224:
	;
	goto L225
L225:
	;
	v855 = v835
	v856 = int32(0)
	v859 = v836
	goto L226
L226:
	;
	v861 = int32(*(*int8)(unsafe.Add(mBase, uint32(v855))))
	v862 = int32(-65)
	v865 = int32(*(*int8)(unsafe.Add(mBase, uint32(v855)+1)))
	v869 = int32(*(*int8)(unsafe.Add(mBase, uint32(v855)+2)))
	v873 = int32(*(*int8)(unsafe.Add(mBase, uint32(v855)+3)))
	v876 = v856 + base.B2i32(v862 < v861) + base.B2i32(v862 < v865) + base.B2i32(v862 < v869) + base.B2i32(v862 < v873)
	v877 = int32(4)
	v878 = v855 + v877
	v880 = v859 + v877
	if v880 != v843&int32(-4) {
		v855 = v878
		v856 = v876
		v859 = v880
		goto L226
	} else {
		goto L228
	}
L227:
	;
	if v848 == int32(0) {
		v906 = v876
		goto L221
	} else {
		goto L229
	}
L228:
	;
	goto L227
L229:
	;
	v884 = v878
	v885 = v876
	goto L222
L230:
	;
	v896 = int32(*(*int8)(unsafe.Add(mBase, uint32(v890))))
	v899 = v891 + base.B2i32(int32(-65) < v896)
	v900 = int32(1)
	v903 = v895 + v900
	if v903 != v848 {
		v890 = v890 + v900
		v891 = v899
		v895 = v903
		goto L230
	} else {
		goto L232
	}
L231:
	;
	v906 = v899
	goto L221
L232:
	;
	goto L231
L233:
	;
	v922 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_21))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L2
	} else {
		goto L234
	}
L234:
	;
	if int32(0) <= v922 {
		goto L194
	} else {
		goto L235
	}
L235:
	;
	v1097 = v922
	goto L176
L236:
	;
	if v930 != 0 {
		goto L195
	} else {
		goto L237
	}
L237:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v932 + (v795 - v926)
	v938 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_22))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L2
	} else {
		goto L238
	}
L238:
	;
	if int32(0) <= v938 {
		goto L194
	} else {
		goto L239
	}
L239:
	;
	v1097 = v938
	goto L176
L240:
	;
	if v946 != 0 {
		goto L195
	} else {
		goto L241
	}
L241:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v948 + (v795 - v942)
	v952 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v952 {
		goto L194
	} else {
		goto L242
	}
L242:
	;
	v1097 = v952
	goto L176
L243:
	;
	if int32(0) <= v957 {
		goto L194
	} else {
		goto L244
	}
L244:
	;
	v1097 = v957
	goto L176
L245:
	;
	if v981 == int32(0) {
		v1097 = v676
		goto L176
	} else {
		goto L249
	}
L246:
	;
	goto L245
L247:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v977 = F_memcmp(m, v974+v963-v966, int32(_a_F_tamil_UTF_8_stem_23), v966)
	mBase = m.M
	if v977 != 0 {
		v981 = v968
		goto L246
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v963 - v966
	v981 = int32(1)
	goto L246
L249:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v986 = v984 - v985
	v990 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_24), int32(6), int32(0))
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L2
	} else {
		goto L250
	}
L250:
	;
	if v990 != 0 {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v992 - v986
	v998 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_25), int32(6), int32(0))
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L2
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1021 = v1020 - v986
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1021
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1021
	v1026 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_26))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L2
	} else {
		goto L261
	}
L254:
	;
	if v998 == int32(0) {
		v1097 = v676
		goto L176
	} else {
		goto L255
	}
L255:
	;
	v1002 = int32(3)
	v1004 = int32(0)
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1006-v1007 < v1002 {
		v1017 = v1004
		goto L257
	} else {
		goto L258
	}
L256:
	;
	if v1017 == int32(0) {
		v1097 = v676
		goto L176
	} else {
		goto L260
	}
L257:
	;
	goto L256
L258:
	;
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1013 = F_memcmp(m, v1010+v1006-v1002, int32(_a_F_tamil_UTF_8_stem_27), v1002)
	mBase = m.M
	if v1013 != 0 {
		v1017 = v1004
		goto L257
	} else {
		goto L259
	}
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1006 - v1002
	v1017 = int32(1)
	goto L257
L260:
	;
	goto L253
L261:
	;
	if v1026 < int32(0) {
		v1097 = v1026
		goto L176
	} else {
		goto L262
	}
L262:
	;
	goto L194
L263:
	;
	if v1053 != 0 {
		goto L267
	} else {
		goto L268
	}
L264:
	;
	goto L263
L265:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1049 = F_memcmp(m, v1046+v1035-v1038, int32(_a_F_tamil_UTF_8_stem_28), v1038)
	mBase = m.M
	if v1049 != 0 {
		v1053 = v1040
		goto L264
	} else {
		goto L266
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1035 - v1038
	v1053 = int32(1)
	goto L264
L267:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1054
	v1058 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_29))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L2
	} else {
		goto L270
	}
L268:
	;
	v1062 = v1035
	goto L269
L269:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1063
	v1066 = v1062
	goto L273
L270:
	;
	if v1058 < int32(0) {
		v1097 = v1058
		goto L176
	} else {
		goto L271
	}
L271:
	;
	v1062 = v1058
	goto L269
L272:
	;
	v1097 = int32(1)
	goto L176
L273:
	;
	v1074 = F_r_fix_ending(m, l0)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L2
	} else {
		goto L278
	}
L274:
	;
	if v1082 == int32(0) {
		goto L272
	} else {
		goto L289
	}
L275:
	;
	if v1074 < int32(0) {
		goto L282
	} else {
		goto L283
	}
L276:
	;
	v1082 = int32(2)
	goto L275
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1063
	goto L272
L278:
	;
	v1077 = int32(base.Ui32(v1074) >> (uint(int32(31)) % 32))
	if v1074 != 0 {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v1079 = v1077
	goto L281
L280:
	;
	v1079 = int32(4)
	goto L281
L281:
	;
	switch v1079 {
	case 0:
		goto L276
	default:
		v1082 = v1077
		goto L275
	case 4:
		goto L277
	}
L282:
	;
	v1085 = v1074
	goto L284
L283:
	;
	v1085 = v1066
	goto L284
L284:
	;
	if v1074 != 0 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1086 = v1085
	goto L287
L286:
	;
	v1086 = v1066
	goto L287
L287:
	;
	if v1082 == int32(2) {
		v1066 = v1086
		goto L273
	} else {
		goto L288
	}
L288:
	;
	goto L274
L289:
	;
	if v1086 < int32(0) {
		v1097 = v1086
		goto L176
	} else {
		goto L290
	}
L290:
	;
	goto L272
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106
	v1108 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v106
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1111
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1111
	if v1111-int32(8) <= v106 {
		v1177 = v1108
		goto L292
	} else {
		goto L293
	}
L292:
	;
	if v1177 < int32(0) {
		v1656 = v1177
		goto L1
	} else {
		goto L315
	}
L293:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1117+v1111-int32(1)))))
	if v1121 != int32(141) {
		v1177 = v1108
		goto L292
	} else {
		goto L294
	}
L294:
	;
	v1127 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_30), int32(4), int32(0))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L2
	} else {
		goto L295
	}
L295:
	;
	if v1127 == int32(0) {
		v1177 = v1108
		goto L292
	} else {
		goto L296
	}
L296:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1131
	switch v1127 - int32(1) {
	case 0:
		goto L301
	case 1:
		goto L300
	case 2:
		goto L299
	case 3:
		goto L298
	default:
		goto L297
	}
L297:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1174
	v1177 = int32(1)
	goto L292
L298:
	;
	v1169 = F_slice_del(m, l0)
	mBase = m.M
	if v1169 < int32(0) {
		v1177 = v1169
		goto L292
	} else {
		goto L314
	}
L299:
	;
	v1165 = F_slice_from_s(m, l0, int32(6), int32(_a_F_tamil_UTF_8_stem_31))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L2
	} else {
		goto L312
	}
L300:
	;
	v1159 = F_slice_from_s(m, l0, int32(6), int32(_a_F_tamil_UTF_8_stem_32))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L2
	} else {
		goto L310
	}
L301:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1139 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_33), int32(6), int32(0))
	mBase = m.M
	v1140 = m.ExcPending
	if v1140 != 0 {
		goto L2
	} else {
		goto L302
	}
L302:
	;
	if v1139 != 0 {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v1143 = F_slice_from_s(m, l0, int32(9), int32(_a_F_tamil_UTF_8_stem_34))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L2
	} else {
		goto L306
	}
L304:
	;
	goto L305
L305:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1147 + (v1131 - v1135)
	v1153 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_35))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L2
	} else {
		goto L308
	}
L306:
	;
	if int32(0) <= v1143 {
		goto L297
	} else {
		goto L307
	}
L307:
	;
	v1177 = v1143
	goto L292
L308:
	;
	if int32(0) <= v1153 {
		goto L297
	} else {
		goto L309
	}
L309:
	;
	v1177 = v1153
	goto L292
L310:
	;
	if int32(0) <= v1159 {
		goto L297
	} else {
		goto L311
	}
L311:
	;
	v1177 = v1159
	goto L292
L312:
	;
	if int32(0) <= v1165 {
		goto L297
	} else {
		goto L313
	}
L313:
	;
	v1177 = v1165
	goto L292
L314:
	;
	goto L297
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106
	v1182 = int32(0)
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1183-int32(4))))
	if v1191 == v1182 {
		goto L318
	} else {
		goto L319
	}
L316:
	;
	if v1299 < int32(0) {
		v1656 = v1299
		goto L1
	} else {
		goto L339
	}
L317:
	;
	if v1265 < int32(5) {
		v1299 = v1182
		goto L316
	} else {
		goto L333
	}
L318:
	;
	v1265 = int32(0)
	goto L317
L319:
	;
	goto L320
L320:
	;
	v1196 = v1191 & int32(3)
	if base.Ui32(v1191) < base.Ui32(int32(4)) {
		goto L323
	} else {
		goto L324
	}
L321:
	;
	v1265 = v1254
	goto L317
L322:
	;
	v1238 = v1232
	v1239 = v1233
	v1243 = v1182
	goto L330
L323:
	;
	v1232 = v1183
	v1233 = int32(0)
	goto L322
L324:
	;
	goto L325
L325:
	;
	v1203 = v1183
	v1204 = int32(0)
	v1207 = v1182
	goto L326
L326:
	;
	v1209 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1203))))
	v1210 = int32(-65)
	v1213 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1203)+1)))
	v1217 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1203)+2)))
	v1221 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1203)+3)))
	v1224 = v1204 + base.B2i32(v1210 < v1209) + base.B2i32(v1210 < v1213) + base.B2i32(v1210 < v1217) + base.B2i32(v1210 < v1221)
	v1225 = int32(4)
	v1226 = v1203 + v1225
	v1228 = v1207 + v1225
	if v1228 != v1191&int32(-4) {
		v1203 = v1226
		v1204 = v1224
		v1207 = v1228
		goto L326
	} else {
		goto L328
	}
L327:
	;
	if v1196 == int32(0) {
		v1254 = v1224
		goto L321
	} else {
		goto L329
	}
L328:
	;
	goto L327
L329:
	;
	v1232 = v1226
	v1233 = v1224
	goto L322
L330:
	;
	v1244 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1238))))
	v1247 = v1239 + base.B2i32(int32(-65) < v1244)
	v1248 = int32(1)
	v1251 = v1243 + v1248
	if v1251 != v1196 {
		v1238 = v1238 + v1248
		v1239 = v1247
		v1243 = v1251
		goto L330
	} else {
		goto L332
	}
L331:
	;
	v1254 = v1247
	goto L321
L332:
	;
	goto L331
L333:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1268
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1270
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1270
	if v1270-int32(5) <= v1268 {
		v1299 = v1182
		goto L316
	} else {
		goto L334
	}
L334:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1276+v1270-int32(1)))))
	if v1280 != int32(191) {
		v1299 = v1182
		goto L316
	} else {
		goto L335
	}
L335:
	;
	v1286 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_36), int32(2), int32(0))
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L2
	} else {
		goto L336
	}
L336:
	;
	if v1286 == int32(0) {
		v1299 = v1182
		goto L316
	} else {
		goto L337
	}
L337:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1290
	v1292 = F_slice_del(m, l0)
	mBase = m.M
	if v1292 < int32(0) {
		v1299 = v1292
		goto L316
	} else {
		goto L338
	}
L338:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1295
	v1299 = int32(1)
	goto L316
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106
	v1309 = int32(0)
	v1312 = v2
	goto L341
L340:
	;
	if v1650 < int32(0) {
		v1656 = v1650
		goto L1
	} else {
		goto L441
	}
L341:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1316 = int32(0)
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1315-int32(4))))
	if v1323 == v1316 {
		goto L347
	} else {
		goto L348
	}
L342:
	;
	v1650 = v1638
	goto L340
L343:
	;
	if v1610 < int32(0) {
		goto L434
	} else {
		goto L435
	}
L344:
	;
	v1634 = int32(2)
	goto L343
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1314
	v1650 = int32(1)
	goto L340
L346:
	;
	if v1397 < int32(5) {
		goto L345
	} else {
		goto L362
	}
L347:
	;
	v1397 = int32(0)
	goto L346
L348:
	;
	goto L349
L349:
	;
	v1328 = v1323 & int32(3)
	if base.Ui32(v1323) < base.Ui32(int32(4)) {
		goto L352
	} else {
		goto L353
	}
L350:
	;
	v1397 = v1386
	goto L346
L351:
	;
	v1370 = v1364
	v1371 = v1365
	v1375 = v1316
	goto L359
L352:
	;
	v1364 = v1315
	v1365 = int32(0)
	goto L351
L353:
	;
	goto L354
L354:
	;
	v1335 = v1315
	v1336 = int32(0)
	v1339 = v1316
	goto L355
L355:
	;
	v1341 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1335))))
	v1342 = int32(-65)
	v1345 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1335)+1)))
	v1349 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1335)+2)))
	v1353 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1335)+3)))
	v1356 = v1336 + base.B2i32(v1342 < v1341) + base.B2i32(v1342 < v1345) + base.B2i32(v1342 < v1349) + base.B2i32(v1342 < v1353)
	v1357 = int32(4)
	v1358 = v1335 + v1357
	v1360 = v1339 + v1357
	if v1360 != v1323&int32(-4) {
		v1335 = v1358
		v1336 = v1356
		v1339 = v1360
		goto L355
	} else {
		goto L357
	}
L356:
	;
	if v1328 == int32(0) {
		v1386 = v1356
		goto L350
	} else {
		goto L358
	}
L357:
	;
	goto L356
L358:
	;
	v1364 = v1358
	v1365 = v1356
	goto L351
L359:
	;
	v1376 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1370))))
	v1379 = v1371 + base.B2i32(int32(-65) < v1376)
	v1380 = int32(1)
	v1383 = v1375 + v1380
	if v1383 != v1328 {
		v1370 = v1370 + v1380
		v1371 = v1379
		v1375 = v1383
		goto L359
	} else {
		goto L361
	}
L360:
	;
	v1386 = v1379
	goto L350
L361:
	;
	goto L360
L362:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1400
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1402
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1402
	v1405 = int32(0)
	v1409 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_37), int32(46), v1405)
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L2
	} else {
		goto L366
	}
L363:
	;
	v1619 = int32(base.Ui32(v1610) >> (uint(int32(31)) % 32))
	if v1610 != 0 {
		goto L431
	} else {
		goto L432
	}
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1542
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1542
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1542-int32(8) <= v1549 {
		v1577 = v1543
		goto L403
	} else {
		goto L404
	}
L365:
	;
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1542 = v1541
	v1543 = v1537
	v1545 = v1539
	goto L364
L366:
	;
	if v1409 == int32(0) {
		v1537 = v1405
		v1539 = v1309
		goto L365
	} else {
		goto L367
	}
L367:
	;
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1413
	v1415 = int32(1)
	switch v1409 - v1415 {
	case 0:
		goto L373
	case 1:
		goto L372
	case 2:
		goto L371
	case 3:
		goto L370
	case 4:
		goto L369
	case 5:
		goto L368
	default:
		v1537 = v1415
		v1539 = v1413
		goto L365
	}
L368:
	;
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1509 = int32(3)
	v1511 = int32(0)
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1513-v1514 < v1509 {
		v1524 = v1511
		goto L396
	} else {
		goto L397
	}
L369:
	;
	v1504 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_38))
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L2
	} else {
		goto L393
	}
L370:
	;
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1474 = int32(3)
	v1476 = int32(0)
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1478-v1479 < v1474 {
		v1489 = v1476
		goto L385
	} else {
		goto L386
	}
L371:
	;
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1459 = int32(0)
	v1463 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_39), int32(8), v1459)
	mBase = m.M
	v1464 = m.ExcPending
	if v1464 != 0 {
		goto L2
	} else {
		goto L381
	}
L372:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1413-int32(2) <= v1422 {
		v1449 = v1421
		goto L375
	} else {
		goto L376
	}
L373:
	;
	v1418 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1418 {
		v1537 = v1415
		v1539 = v1413
		goto L365
	} else {
		goto L374
	}
L374:
	;
	v1610 = v1418
	goto L363
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1413 - v1421 + v1449
	v1455 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1455 {
		v1537 = int32(1)
		v1539 = v1413
		goto L365
	} else {
		goto L380
	}
L376:
	;
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1428 = int32(1)
	v1430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1426+v1413-v1428))))
	if base.B2i32(v1430&int32(224) != int32(128))|base.B2i32(v1428<<(uint(v1430)%32)&int32(_a_F_tamil_UTF_8_stem_40) == int32(0)) != 0 {
		v1449 = v1421
		goto L375
	} else {
		goto L377
	}
L377:
	;
	v1442 = int32(0)
	v1446 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_41), int32(12), v1442)
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L2
	} else {
		goto L378
	}
L378:
	;
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1446 != 0 {
		v1542 = v1448
		v1543 = v1442
		v1545 = v1413
		goto L364
	} else {
		goto L379
	}
L379:
	;
	v1449 = v1448
	goto L375
L380:
	;
	v1610 = v1455
	goto L363
L381:
	;
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1463 != 0 {
		v1542 = v1465
		v1543 = v1459
		v1545 = v1413
		goto L364
	} else {
		goto L382
	}
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1413 - v1458 + v1465
	v1470 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1470 {
		v1537 = int32(1)
		v1539 = v1413
		goto L365
	} else {
		goto L383
	}
L383:
	;
	v1610 = v1470
	goto L363
L384:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1489 != 0 {
		goto L388
	} else {
		goto L389
	}
L385:
	;
	goto L384
L386:
	;
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1485 = F_memcmp(m, v1482+v1478-v1474, int32(_a_F_tamil_UTF_8_stem_42), v1474)
	mBase = m.M
	if v1485 != 0 {
		v1489 = v1476
		goto L385
	} else {
		goto L387
	}
L387:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1478 - v1474
	v1489 = int32(1)
	goto L385
L388:
	;
	v1542 = v1490
	v1543 = int32(0)
	v1545 = v1413
	goto L364
L389:
	;
	goto L390
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1413 - v1473 + v1490
	v1498 = F_slice_from_s(m, l0, int32(3), int32(_a_F_tamil_UTF_8_stem_43))
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L2
	} else {
		goto L391
	}
L391:
	;
	if int32(0) <= v1498 {
		v1537 = int32(1)
		v1539 = v1413
		goto L365
	} else {
		goto L392
	}
L392:
	;
	v1610 = v1498
	goto L363
L393:
	;
	if int32(0) <= v1504 {
		v1537 = v1415
		v1539 = v1413
		goto L365
	} else {
		goto L394
	}
L394:
	;
	v1610 = v1504
	goto L363
L395:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1524 == int32(0) {
		goto L399
	} else {
		goto L400
	}
L396:
	;
	goto L395
L397:
	;
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1520 = F_memcmp(m, v1517+v1513-v1509, int32(_a_F_tamil_UTF_8_stem_44), v1509)
	mBase = m.M
	if v1520 != 0 {
		v1524 = v1511
		goto L396
	} else {
		goto L398
	}
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1513 - v1509
	v1524 = int32(1)
	goto L396
L399:
	;
	v1542 = v1525
	v1543 = int32(0)
	v1545 = v1413
	goto L364
L400:
	;
	goto L401
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1413 - v1508 + v1525
	v1533 = F_slice_del(m, l0)
	mBase = m.M
	if v1533 < int32(0) {
		v1610 = v1533
		goto L363
	} else {
		goto L402
	}
L402:
	;
	v1537 = int32(1)
	v1539 = v1413
	goto L365
L403:
	;
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1578
	v1584 = v1545
	goto L409
L404:
	;
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1553+v1542-int32(1)))))
	if base.B2i32(v1557 != int32(177))&base.B2i32(v1557 != int32(141)) != 0 {
		v1577 = v1543
		goto L403
	} else {
		goto L405
	}
L405:
	;
	v1566 = F_find_among_b(m, l0, int32(_a_F_tamil_UTF_8_stem_45), int32(6), int32(0))
	mBase = m.M
	v1567 = m.ExcPending
	if v1567 != 0 {
		goto L2
	} else {
		goto L406
	}
L406:
	;
	if v1566 == int32(0) {
		v1577 = v1543
		goto L403
	} else {
		goto L407
	}
L407:
	;
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1570
	v1573 = F_slice_del(m, l0)
	mBase = m.M
	if v1573 < int32(0) {
		v1610 = v1573
		goto L363
	} else {
		goto L408
	}
L408:
	;
	v1577 = int32(1)
	goto L403
L409:
	;
	v1589 = F_r_fix_ending(m, l0)
	mBase = m.M
	v1590 = m.ExcPending
	if v1590 != 0 {
		goto L2
	} else {
		goto L414
	}
L410:
	;
	if v1597 == int32(0) {
		goto L425
	} else {
		goto L426
	}
L411:
	;
	if v1589 < int32(0) {
		goto L418
	} else {
		goto L419
	}
L412:
	;
	v1597 = int32(2)
	goto L411
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1578
	v1610 = v1577
	goto L363
L414:
	;
	v1592 = int32(base.Ui32(v1589) >> (uint(int32(31)) % 32))
	if v1589 != 0 {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	v1594 = v1592
	goto L417
L416:
	;
	v1594 = int32(4)
	goto L417
L417:
	;
	switch v1594 {
	case 0:
		goto L412
	default:
		v1597 = v1592
		goto L411
	case 4:
		goto L413
	}
L418:
	;
	v1600 = v1589
	goto L420
L419:
	;
	v1600 = v1584
	goto L420
L420:
	;
	if v1589 != 0 {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v1601 = v1600
	goto L423
L422:
	;
	v1601 = v1584
	goto L423
L423:
	;
	if v1597 == int32(2) {
		v1584 = v1601
		goto L409
	} else {
		goto L424
	}
L424:
	;
	goto L410
L425:
	;
	v1610 = v1577
	goto L363
L426:
	;
	goto L427
L427:
	;
	if v1601 < int32(0) {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v1608 = v1601
	goto L430
L429:
	;
	v1608 = v1577
	goto L430
L430:
	;
	v1610 = v1608
	goto L363
L431:
	;
	v1621 = v1619
	goto L433
L432:
	;
	v1621 = int32(4)
	goto L433
L433:
	;
	switch v1621 {
	case 0:
		goto L344
	default:
		v1634 = v1619
		goto L343
	case 4:
		goto L345
	}
L434:
	;
	v1637 = v1610
	goto L436
L435:
	;
	v1637 = v1312
	goto L436
L436:
	;
	if v1610 != 0 {
		goto L437
	} else {
		goto L438
	}
L437:
	;
	v1638 = v1637
	goto L439
L438:
	;
	v1638 = v1312
	goto L439
L439:
	;
	if v1634 != int32(1) {
		v1309 = v1634
		v1312 = v1638
		goto L341
	} else {
		goto L440
	}
L440:
	;
	goto L342
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v106
	v1656 = int32(1)
	goto L1
}
func F_texteqname(m *base.Module, l0 int32) int64 {
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	F_errfinish(m, int32(_a_F_texteqname_2), int32(1337), int32(_a_F_texteqname_3))
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
	return base.I64_extend_i32_u(v150)
L54:
	;
	goto L53
}
func F_texticregexeq_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_like_regex_support(m, v2, int32(3))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_textlen(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_text_length(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_s(v3)
	}
}
func F_textlike(m *base.Module, l0 int32) int64 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v16 == int32(1) {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
				if v22 == int32(18) {
					v25 = int32(16)
				} else {
					v25 = int32(0)
				}
				if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v32 = int32(4)
				} else {
					v32 = v25
				}
				v45 = v32
			} else {
				v33 = int32(1)
				if v16&v33 != 0 {
					v45 = int32(base.Ui32(v16)>>(uint(v33)%32)) - v33
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v46 == int32(1) {
				v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
				if v52 == int32(18) {
					v55 = int32(16)
				} else {
					v55 = int32(0)
				}
				if base.Ui32((v52-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v62 = int32(4)
				} else {
					v62 = v55
				}
				v75 = v62
			} else {
				v63 = int32(1)
				if v46&v63 != 0 {
					v75 = int32(base.Ui32(v46)>>(uint(v63)%32)) - v63
				} else {
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v75 = int32(base.Ui32(v69)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v76 = int32(1)
			if v16&v76 != 0 {
				v80 = v76
			} else {
				v80 = int32(4)
			}
			v82 = int32(1)
			if v46&v82 != 0 {
				v86 = v82
			} else {
				v86 = int32(4)
			}
			v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v89 = F_GenericMatchText(m, v9+v80, v45, v14+v86, v75, v88)
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(base.B2i32(v89 == int32(1)))
			}
		}
	}
}
func F_textnlike(m *base.Module, l0 int32) int64 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v16 == int32(1) {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
				if v22 == int32(18) {
					v25 = int32(16)
				} else {
					v25 = int32(0)
				}
				if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v32 = int32(4)
				} else {
					v32 = v25
				}
				v45 = v32
			} else {
				v33 = int32(1)
				if v16&v33 != 0 {
					v45 = int32(base.Ui32(v16)>>(uint(v33)%32)) - v33
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v46 == int32(1) {
				v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
				if v52 == int32(18) {
					v55 = int32(16)
				} else {
					v55 = int32(0)
				}
				if base.Ui32((v52-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v62 = int32(4)
				} else {
					v62 = v55
				}
				v75 = v62
			} else {
				v63 = int32(1)
				if v46&v63 != 0 {
					v75 = int32(base.Ui32(v46)>>(uint(v63)%32)) - v63
				} else {
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v75 = int32(base.Ui32(v69)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v76 = int32(1)
			if v16&v76 != 0 {
				v80 = v76
			} else {
				v80 = int32(4)
			}
			v82 = int32(1)
			if v46&v82 != 0 {
				v86 = v82
			} else {
				v86 = int32(4)
			}
			v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v89 = F_GenericMatchText(m, v9+v80, v45, v14+v86, v75, v88)
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(base.B2i32(v89 != int32(1)))
			}
		}
	}
}
func F_tidgt(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return base.I64_extend_i32_u(base.B2i32(int32(0) < v27))
}
func F_tidin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
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
	var v149 int64
	_ = v149
	v7 = int64(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = v13
	v16 = int32(0)
	goto L4
L1:
	;
	m.G0 = v10 - int32(-64)
	return v149
L2:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	v109 = F_strtox_2(m, v104, v8+int32(-12), int32(10), int64(4294967295))
	mBase = m.M
	v110 = base.I32_wrap_i64(v109)
	goto L29
L3:
	;
	v84 = F_errsave_start(m, v12)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L18
	} else {
		goto L24
	}
L4:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	switch v21 - int32(41) {
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
	*(*int32)(unsafe.Add(mBase, _c_F_tidin[0])) = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
	v52 = F_strtox_2(m, v47, v8+int32(-12), int32(10), int64(4294967295))
	mBase = m.M
	v53 = base.I32_wrap_i64(v52)
	goto L13
L6:
	;
	v40 = int32(1)
	if v39 <= v40 {
		v14 = v14 + v40
		v16 = v39
		goto L4
	} else {
		goto L12
	}
L7:
	;
	v34 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8+int32(-8)+v16<<(uint(int32(2))%32)))) = v14 + v34
	v39 = v16 + v34
	goto L6
L8:
	;
	if base.B2i32(v21 != int32(40))|v16 != 0 {
		v39 = v16
		goto L6
	} else {
		goto L11
	}
L9:
	;
	if v21 == int32(0) {
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
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_tidin[0]))
	if v55 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v59 == int32(44) {
		goto L2
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v62 = F_errsave_start(m, v12)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	return int64(0)
L19:
	;
	if v62 == int32(0) {
		v149 = v7
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_tidin_0)
	F_errmsg(m, int32(_a_F_tidin_1), v8+int32(-48))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	F_errsave_finish(m, v12, int32(_a_F_tidin_2), int32(80), int32(_a_F_tidin_3))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v149 = v7
	goto L1
L24:
	;
	if v84 == int32(0) {
		v149 = v7
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(_a_F_tidin_0)
	F_errmsg(m, int32(_a_F_tidin_1), v8+int32(-32))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	F_errsave_finish(m, v12, int32(_a_F_tidin_2), int32(72), int32(_a_F_tidin_3))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	v149 = v7
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
	v119 = F_errsave_start(m, v12)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L18
	} else {
		goto L35
	}
L32:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
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
		v149 = v7
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_tidin_0)
	F_errmsg(m, int32(_a_F_tidin_1), v10)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L18
	} else {
		goto L38
	}
L38:
	;
	F_errsave_finish(m, v12, int32(_a_F_tidin_2), int32(103), int32(_a_F_tidin_3))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L18
	} else {
		goto L39
	}
L39:
	;
	v149 = v7
	goto L1
L40:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+4)) = uint16(v110)
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+2)) = uint16(v53)
	v143 = int32(base.Ui32(v53) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v138))) = uint16(v143)
	v149 = base.I64_extend_i32_u(v138)
	goto L1
}
func F_tidne(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return base.I64_extend_i32_u(base.B2i32(v27 != int32(0)))
}
func F_timestamptz2date_safe(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	if l0 == int64(-9223372036854775807-1) {
		v83 = int32(-2147483648)
		m.G0 = v7 - int32(-64)
		return v83
	} else {
		if l0 == int64(9223372036854775807) {
			v83 = int32(2147483647)
			m.G0 = v7 - int32(-64)
			return v83
		} else {
			v21 = int32(0)
			v23 = F_timestamp2tm(m, l0, v5+int32(-52), v5+int32(-44), v5+int32(-48), v21, v21)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v23 != 0 {
					v31 = base.I32_wrap_i64(l0>>(uint(int64(63))%64)) ^ int32(2147483647)
					v32 = F_errsave_start(m, l1)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						if v32 == int32(0) {
							v83 = v31
							m.G0 = v7 - int32(-64)
							return v83
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_timestamptz2date_safe_0), int32(0))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, l1, int32(_a_F_timestamptz2date_safe_1), int32(1460), int32(_a_F_timestamptz2date_safe_2))
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										v83 = v31
										m.G0 = v7 - int32(-64)
										return v83
									}
								}
							}
						}
					}
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
					v55 = base.B2i32(int32(2) < v49)
					if int32(2) < v49 {
						v56 = int32(_a_F_timestamptz2date_safe_3)
					} else {
						v56 = int32(_a_F_timestamptz2date_safe_4)
					}
					v57 = v56 + v48
					v62 = base.I32_div_s(v57, int32(4))
					v65 = base.I32_div_s(v57, int32(-100))
					v68 = base.I32_div_s(v57, int32(400))
					if int32(2) < v49 {
						v72 = int32(1)
					} else {
						v72 = int32(13)
					}
					v77 = base.I32_div_s((v72+v49)*int32(_a_F_timestamptz2date_safe_5), int32(256))
					v83 = v50 + v57*int32(365) + v62 + v65 + v68 + v77 - int32(_a_F_timestamptz2date_safe_6) - int32(_a_F_timestamptz2date_safe_7)
					m.G0 = v7 - int32(-64)
					return v83
				}
			}
		}
	}
}
func F_toupper_libc_mb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		v24 = F_casemap(m, l0, int32(1))
		mBase = m.M
		return v24
	} else {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
		if v5&int32(1) == int32(0) {
			v24 = F_casemap(m, l0, int32(1))
			mBase = m.M
			return v24
		} else {
			if base.Ui32((l0-int32(97))&int32(255)) < base.Ui32(int32(26)) {
				v18 = l0 + int32(224)
			} else {
				v18 = l0
			}
			return v18 & int32(255)
		}
	}
}
func F_toupper_libc_sb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		if base.Ui32(l0) <= base.Ui32(int32(255)) {
			if base.Ui32(l0-int32(97)) < base.Ui32(int32(26)) {
				v31 = l0 & int32(95)
			} else {
				v31 = l0
			}
			v32 = v31
		} else {
			v32 = l0
		}
		return v32
	} else {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
		if v5&int32(1) == int32(0) {
			if base.Ui32(l0) <= base.Ui32(int32(255)) {
				if base.Ui32(l0-int32(97)) < base.Ui32(int32(26)) {
					v31 = l0 & int32(95)
				} else {
					v31 = l0
				}
				v32 = v31
			} else {
				v32 = l0
			}
			return v32
		} else {
			if base.Ui32((l0-int32(97))&int32(255)) < base.Ui32(int32(26)) {
				v18 = l0 + int32(224)
			} else {
				v18 = l0
			}
			return v18 & int32(255)
		}
	}
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
	var v46 int32
	_ = v46
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
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v244 int32
	_ = v244
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
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
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
	v244 = m.ExcPending
	if v244 != 0 {
		goto L10
	} else {
		goto L63
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
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
	v185 = v7
	goto L5
L5:
	;
	m.G0 = v17 + int32(48)
	return v185
L6:
	;
	F_checkExprIsVarFree(m, l0, v169, v164)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
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
	v164 = int32(_a_F_transformFrameOffset_0)
	v169 = v31
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
	v149 = int32(0)
	if l1&int32(8) == v149 {
		v164 = v149
		v169 = v7
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
	v46 = F_SearchSysCacheList(m, int32(5), int32(2), base.I64_extend_i32_u(l2), base.I64_extend_i32_u(l3), int64(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+56))
	if v48 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_ReleaseCatCacheList(m, v46)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
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
	v55 = l3
	goto L25
L24:
	;
	v55 = v38
	goto L25
L25:
	;
	v61 = int32(0)
	v66 = v7
	v69 = v7
	v70 = v7
	v72 = v7
	goto L26
L26:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v46-int32(-64)+v66<<(uint(int32(2))%32))))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+72))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+22)))
	v79 = v77 + v78
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+16)))
	if v80 != int32(3) {
		v100 = v61
		v101 = v69
		v102 = v70
		v103 = v72
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_ReleaseCatCacheList(m, v46)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L10
	} else {
		goto L36
	}
L28:
	;
	v105 = v66 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v46)+56))
	if v105 < v106 {
		v61 = v100
		v66 = v105
		v69 = v101
		v70 = v102
		v72 = v103
		goto L26
	} else {
		goto L35
	}
L29:
	;
	v83 = int32(1)
	v84 = v61 + v83
	v91 = F_can_coerce_type(m, v83, v17+int32(44), v79+int32(12), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	if v91 == int32(0) {
		v100 = v84
		v101 = v69
		v102 = v70
		v103 = v72
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v96 = v70 + int32(1)
	if v55 == v69 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v100 = v84
	v101 = v55
	v102 = v96
	v103 = v72
	goto L28
L33:
	;
	goto L34
L34:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v100 = v84
	v101 = v99
	v102 = v96
	v103 = v98
	goto L28
L35:
	;
	goto L27
L36:
	;
	if v100 == int32(0) {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	switch v102 {
	case 0:
		goto L40
	case 1:
		goto L38
	default:
		goto L39
	}
L38:
	;
	v144 = int32(_a_F_transformFrameOffset_1)
	v146 = F_coerce_to_specific_type(m, l0, v36, v101, v144)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L10
	} else {
		goto L50
	}
L39:
	;
	if v55 != v101 {
		goto L2
	} else {
		goto L49
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L10
	} else {
		goto L42
	}
L42:
	;
	v119 = F_format_type_be(m, l3)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	v122 = F_format_type_be(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v119
	F_errmsg(m, int32(_a_F_transformFrameOffset_2), v17+int32(32))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	F_errhint(m, int32(_a_F_transformFrameOffset_3), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	v135 = F_exprLocation(m, v36)
	mBase = m.M
	F_parser_errposition(m, l0, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_transformFrameOffset_4), int32(3797), int32(_a_F_transformFrameOffset_5))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v103
	v164 = v144
	v169 = v146
	goto L6
L51:
	;
	v156 = F_transformExpr(m, l0, l5, int32(13))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L10
	} else {
		goto L52
	}
L52:
	;
	v160 = F_coerce_to_specific_type(m, l0, v156, int32(20), int32(_a_F_transformFrameOffset_6))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	v164 = int32(_a_F_transformFrameOffset_6)
	v169 = v160
	goto L6
L54:
	;
	v185 = v169
	goto L5
L55:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	v203 = F_format_type_be(m, l3)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L10
	} else {
		goto L57
	}
L57:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	v206 = F_format_type_be(m, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v206
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v203
	F_errmsg(m, int32(_a_F_transformFrameOffset_7), v17+int32(16))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	F_errhint(m, int32(_a_F_transformFrameOffset_8), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	v219 = F_exprLocation(m, v36)
	mBase = m.M
	F_parser_errposition(m, l0, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_transformFrameOffset_4), int32(3805), int32(_a_F_transformFrameOffset_5))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
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
	v247 = m.ExcPending
	if v247 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	v248 = F_format_type_be(m, l3)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L10
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v248
	F_errmsg(m, int32(_a_F_transformFrameOffset_9), v17)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	v254 = F_exprLocation(m, v36)
	mBase = m.M
	F_parser_errposition(m, l0, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L10
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_transformFrameOffset_4), int32(3789), int32(_a_F_transformFrameOffset_5))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
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
func F_trigram_qsort_unsigned(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
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
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v679 int32
	_ = v679
	var v689 int32
	_ = v689
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v796 int32
	_ = v796
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v903 int32
	_ = v903
	var v911 int32
	_ = v911
	var v925 int32
	_ = v925
	var v930 int32
	_ = v930
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v994 int32
	_ = v994
	if base.Ui32(l1) < base.Ui32(int32(7)) {
		v885 = l0
		v886 = l1
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	if base.Ui32(v886) < base.Ui32(int32(2)) {
		goto L1
	} else {
		goto L316
	}
L3:
	;
	v19 = l0
	v20 = l1
	goto L4
L4:
	;
	v36 = v19 + int32(3)
	v38 = v20
	goto L6
L5:
	;
	v885 = v19
	v886 = v882
	goto L2
L6:
	;
	v54 = v38 * int32(3)
	if base.Ui32(v54) < base.Ui32(int32(4)) {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v57 = v19 + v54
	v60 = v36
	goto L9
L9:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60-int32(3)))))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v76 != v77 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v102 = v19 + int32(base.Ui32(v38)>>(uint(int32(1))%32))*int32(3)
	if v38 != int32(7) {
		goto L23
	} else {
		goto L24
	}
L11:
	;
	goto L10
L12:
	;
	v94 = v60 + int32(3)
	if base.Ui32(v94) < base.Ui32(v57) {
		v60 = v94
		goto L9
	} else {
		goto L22
	}
L13:
	;
	if base.Ui32(v76) < base.Ui32(v77) {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60-int32(2)))))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	if v82 != v83 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L11
L17:
	;
	if base.Ui32(v83) <= base.Ui32(v82) {
		goto L11
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60-int32(1)))))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	if base.Ui32(v89) < base.Ui32(v88) {
		goto L11
	} else {
		goto L21
	}
L20:
	;
	goto L12
L21:
	;
	goto L12
L22:
	;
	goto L1
L23:
	;
	v106 = v57 - int32(3)
	if base.Ui32(v38) < base.Ui32(int32(41)) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v431 = v102
	goto L25
L25:
	;
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431))))
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v436)
	*(*uint8)(unsafe.Add(mBase, uint32(v431))) = uint8(v435)
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)) = uint8(v440)
	*(*uint8)(unsafe.Add(mBase, uint32(v431)+1)) = uint8(v439)
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)) = uint8(v444)
	*(*uint8)(unsafe.Add(mBase, uint32(v431)+2)) = uint8(v443)
	v448 = v57 - int32(3)
	v451 = v36
	v452 = v36
	v453 = v448
	v456 = v448
	goto L242
L26:
	;
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350))))
	if v358 != v359 {
		goto L193
	} else {
		goto L194
	}
L27:
	;
	v349 = v19
	v350 = v102
	v352 = v106
	goto L26
L28:
	;
	goto L29
L29:
	;
	v109 = int32(3)
	v110 = int32(base.Ui32(v38) >> (uint(v109) % 32))
	v112 = v110 * v109
	v113 = v19 + v112
	v115 = v110 * int32(6)
	v116 = v19 + v115
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v121 != v122 {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v193 = v102 - v112
	v194 = v102 + v112
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v199 != v200 {
		goto L87
	} else {
		goto L88
	}
L31:
	;
	v192 = v182
	goto L30
L32:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v159 != v122 {
		goto L64
	} else {
		goto L65
	}
L33:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v134 != v122 {
		goto L44
	} else {
		goto L45
	}
L34:
	;
	if base.Ui32(v121) < base.Ui32(v122) {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
	if v125 != v126 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L32
L38:
	;
	if base.Ui32(v126) <= base.Ui32(v125) {
		goto L32
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+2)))
	if base.Ui32(v130) <= base.Ui32(v129) {
		goto L32
	} else {
		goto L42
	}
L41:
	;
	goto L33
L42:
	;
	goto L33
L43:
	;
	if v121 != v134 {
		goto L54
	} else {
		goto L55
	}
L44:
	;
	if base.Ui32(v134) <= base.Ui32(v122) {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	if v137 != v138 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v182 = v113
	goto L31
L48:
	;
	if base.Ui32(v138) <= base.Ui32(v137) {
		goto L43
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+2)))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+2)))
	if base.Ui32(v141) < base.Ui32(v142) {
		v182 = v113
		goto L31
	} else {
		goto L52
	}
L51:
	;
	v182 = v113
	goto L31
L52:
	;
	goto L43
L53:
	;
	v192 = v19
	goto L30
L54:
	;
	if base.Ui32(v134) <= base.Ui32(v121) {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	if v148 != v149 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v182 = v116
	goto L31
L58:
	;
	if base.Ui32(v149) <= base.Ui32(v148) {
		goto L53
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+2)))
	if base.Ui32(v152) < base.Ui32(v153) {
		v182 = v116
		goto L31
	} else {
		goto L62
	}
L61:
	;
	v182 = v116
	goto L31
L62:
	;
	goto L53
L63:
	;
	if v121 != v159 {
		goto L74
	} else {
		goto L75
	}
L64:
	;
	if base.Ui32(v122) < base.Ui32(v159) {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	if v162 != v163 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v182 = v113
	goto L31
L68:
	;
	if base.Ui32(v162) < base.Ui32(v163) {
		goto L63
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+2)))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+2)))
	if base.Ui32(v167) < base.Ui32(v166) {
		v182 = v113
		goto L31
	} else {
		goto L72
	}
L71:
	;
	v182 = v113
	goto L31
L72:
	;
	goto L63
L73:
	;
	v182 = v116
	goto L31
L74:
	;
	if base.Ui32(v159) <= base.Ui32(v121) {
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	if v173 != v174 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v182 = v19
	goto L31
L78:
	;
	if base.Ui32(v174) <= base.Ui32(v173) {
		goto L73
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+2)))
	if base.Ui32(v177) < base.Ui32(v178) {
		v182 = v19
		goto L31
	} else {
		goto L82
	}
L81:
	;
	v182 = v19
	goto L31
L82:
	;
	goto L73
L83:
	;
	v271 = v106 - v115
	v272 = v106 - v112
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	if v277 != v278 {
		goto L140
	} else {
		goto L141
	}
L84:
	;
	v270 = v260
	goto L83
L85:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v237 != v200 {
		goto L117
	} else {
		goto L118
	}
L86:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v212 != v200 {
		goto L97
	} else {
		goto L98
	}
L87:
	;
	if base.Ui32(v199) < base.Ui32(v200) {
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	if v203 != v204 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L85
L91:
	;
	if base.Ui32(v204) <= base.Ui32(v203) {
		goto L85
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+2)))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+2)))
	if base.Ui32(v208) <= base.Ui32(v207) {
		goto L85
	} else {
		goto L95
	}
L94:
	;
	goto L86
L95:
	;
	goto L86
L96:
	;
	if v199 != v212 {
		goto L107
	} else {
		goto L108
	}
L97:
	;
	if base.Ui32(v212) <= base.Ui32(v200) {
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)))
	if v215 != v216 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v260 = v102
	goto L84
L101:
	;
	if base.Ui32(v216) <= base.Ui32(v215) {
		goto L96
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+2)))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+2)))
	if base.Ui32(v219) < base.Ui32(v220) {
		v260 = v102
		goto L84
	} else {
		goto L105
	}
L104:
	;
	v260 = v102
	goto L84
L105:
	;
	goto L96
L106:
	;
	v270 = v193
	goto L83
L107:
	;
	if base.Ui32(v212) <= base.Ui32(v199) {
		goto L106
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)))
	if v226 != v227 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v260 = v194
	goto L84
L111:
	;
	if base.Ui32(v227) <= base.Ui32(v226) {
		goto L106
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+2)))
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+2)))
	if base.Ui32(v230) < base.Ui32(v231) {
		v260 = v194
		goto L84
	} else {
		goto L115
	}
L114:
	;
	v260 = v194
	goto L84
L115:
	;
	goto L106
L116:
	;
	if v199 != v237 {
		goto L127
	} else {
		goto L128
	}
L117:
	;
	if base.Ui32(v200) < base.Ui32(v237) {
		goto L116
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)))
	if v240 != v241 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v260 = v102
	goto L84
L121:
	;
	if base.Ui32(v240) < base.Ui32(v241) {
		goto L116
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+2)))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+2)))
	if base.Ui32(v245) < base.Ui32(v244) {
		v260 = v102
		goto L84
	} else {
		goto L125
	}
L124:
	;
	v260 = v102
	goto L84
L125:
	;
	goto L116
L126:
	;
	v260 = v194
	goto L84
L127:
	;
	if base.Ui32(v237) <= base.Ui32(v199) {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)))
	if v251 != v252 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v260 = v193
	goto L84
L131:
	;
	if base.Ui32(v252) <= base.Ui32(v251) {
		goto L126
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+2)))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+2)))
	if base.Ui32(v255) < base.Ui32(v256) {
		v260 = v193
		goto L84
	} else {
		goto L135
	}
L134:
	;
	v260 = v193
	goto L84
L135:
	;
	goto L126
L136:
	;
	v349 = v192
	v350 = v270
	v352 = v348
	goto L26
L137:
	;
	v348 = v338
	goto L136
L138:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v315 != v278 {
		goto L170
	} else {
		goto L171
	}
L139:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v290 != v278 {
		goto L150
	} else {
		goto L151
	}
L140:
	;
	if base.Ui32(v277) < base.Ui32(v278) {
		goto L139
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+1)))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+1)))
	if v281 != v282 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	goto L138
L144:
	;
	if base.Ui32(v282) <= base.Ui32(v281) {
		goto L138
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+2)))
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+2)))
	if base.Ui32(v286) <= base.Ui32(v285) {
		goto L138
	} else {
		goto L148
	}
L147:
	;
	goto L139
L148:
	;
	goto L139
L149:
	;
	if v277 != v290 {
		goto L160
	} else {
		goto L161
	}
L150:
	;
	if base.Ui32(v290) <= base.Ui32(v278) {
		goto L149
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+1)))
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v293 != v294 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	v338 = v272
	goto L137
L154:
	;
	if base.Ui32(v294) <= base.Ui32(v293) {
		goto L149
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+2)))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+2)))
	if base.Ui32(v297) < base.Ui32(v298) {
		v338 = v272
		goto L137
	} else {
		goto L158
	}
L157:
	;
	v338 = v272
	goto L137
L158:
	;
	goto L149
L159:
	;
	v348 = v271
	goto L136
L160:
	;
	if base.Ui32(v290) <= base.Ui32(v277) {
		goto L159
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+1)))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v304 != v305 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v338 = v106
	goto L137
L164:
	;
	if base.Ui32(v305) <= base.Ui32(v304) {
		goto L159
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+2)))
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+2)))
	if base.Ui32(v308) < base.Ui32(v309) {
		v338 = v106
		goto L137
	} else {
		goto L168
	}
L167:
	;
	v338 = v106
	goto L137
L168:
	;
	goto L159
L169:
	;
	if v277 != v315 {
		goto L180
	} else {
		goto L181
	}
L170:
	;
	if base.Ui32(v278) < base.Ui32(v315) {
		goto L169
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+1)))
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v318 != v319 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v338 = v272
	goto L137
L174:
	;
	if base.Ui32(v318) < base.Ui32(v319) {
		goto L169
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+2)))
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+2)))
	if base.Ui32(v323) < base.Ui32(v322) {
		v338 = v272
		goto L137
	} else {
		goto L178
	}
L177:
	;
	v338 = v272
	goto L137
L178:
	;
	goto L169
L179:
	;
	v338 = v106
	goto L137
L180:
	;
	if base.Ui32(v315) <= base.Ui32(v277) {
		goto L179
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+1)))
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v329 != v330 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v338 = v271
	goto L137
L184:
	;
	if base.Ui32(v330) <= base.Ui32(v329) {
		goto L179
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+2)))
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+2)))
	if base.Ui32(v333) < base.Ui32(v334) {
		v338 = v271
		goto L137
	} else {
		goto L188
	}
L187:
	;
	v338 = v271
	goto L137
L188:
	;
	goto L179
L189:
	;
	v431 = v429
	goto L25
L190:
	;
	v429 = v419
	goto L189
L191:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	if v396 != v359 {
		goto L223
	} else {
		goto L224
	}
L192:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352))))
	if v371 != v359 {
		goto L203
	} else {
		goto L204
	}
L193:
	;
	if base.Ui32(v358) < base.Ui32(v359) {
		goto L192
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+1)))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)))
	if v362 != v363 {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	goto L191
L197:
	;
	if base.Ui32(v363) <= base.Ui32(v362) {
		goto L191
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+2)))
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)))
	if base.Ui32(v367) <= base.Ui32(v366) {
		goto L191
	} else {
		goto L201
	}
L200:
	;
	goto L192
L201:
	;
	goto L192
L202:
	;
	if v358 != v371 {
		goto L213
	} else {
		goto L214
	}
L203:
	;
	if base.Ui32(v371) <= base.Ui32(v359) {
		goto L202
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+1)))
	if v374 != v375 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v419 = v350
	goto L190
L207:
	;
	if base.Ui32(v375) <= base.Ui32(v374) {
		goto L202
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)))
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+2)))
	if base.Ui32(v378) < base.Ui32(v379) {
		v419 = v350
		goto L190
	} else {
		goto L211
	}
L210:
	;
	v419 = v350
	goto L190
L211:
	;
	goto L202
L212:
	;
	v429 = v349
	goto L189
L213:
	;
	if base.Ui32(v371) <= base.Ui32(v358) {
		goto L212
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+1)))
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+1)))
	if v385 != v386 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v419 = v352
	goto L190
L217:
	;
	if base.Ui32(v386) <= base.Ui32(v385) {
		goto L212
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+2)))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+2)))
	if base.Ui32(v389) < base.Ui32(v390) {
		v419 = v352
		goto L190
	} else {
		goto L221
	}
L220:
	;
	v419 = v352
	goto L190
L221:
	;
	goto L212
L222:
	;
	if v358 != v396 {
		goto L233
	} else {
		goto L234
	}
L223:
	;
	if base.Ui32(v359) < base.Ui32(v396) {
		goto L222
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+1)))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+1)))
	if v399 != v400 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v419 = v350
	goto L190
L227:
	;
	if base.Ui32(v399) < base.Ui32(v400) {
		goto L222
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+2)))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+2)))
	if base.Ui32(v404) < base.Ui32(v403) {
		v419 = v350
		goto L190
	} else {
		goto L231
	}
L230:
	;
	v419 = v350
	goto L190
L231:
	;
	goto L222
L232:
	;
	v419 = v352
	goto L190
L233:
	;
	if base.Ui32(v396) <= base.Ui32(v358) {
		goto L232
	} else {
		goto L236
	}
L234:
	;
	goto L235
L235:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+1)))
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+1)))
	if v410 != v411 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v419 = v349
	goto L190
L237:
	;
	if base.Ui32(v411) <= base.Ui32(v410) {
		goto L232
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+2)))
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+2)))
	if base.Ui32(v414) < base.Ui32(v415) {
		v419 = v349
		goto L190
	} else {
		goto L241
	}
L240:
	;
	v419 = v349
	goto L190
L241:
	;
	goto L232
L242:
	;
	if base.Ui32(v453) < base.Ui32(v451) {
		v515 = v451
		v516 = v452
		goto L244
	} else {
		goto L245
	}
L243:
	;
	v878 = int32(3)
	v879 = base.I32_div_u_s(v721, v878)
	F_trigram_qsort_unsigned(m, v57-v721, v879)
	mBase = m.M
	v882 = base.I32_div_u_s(v594, v878)
	if base.Ui32(int32(21)) <= base.Ui32(v594) {
		v38 = v882
		goto L6
	} else {
		goto L315
	}
L244:
	;
	if base.Ui32(v515) <= base.Ui32(v453) {
		goto L264
	} else {
		goto L265
	}
L245:
	;
	v468 = v451
	v469 = v452
	goto L246
L246:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v482 != v483 {
		goto L249
	} else {
		goto L250
	}
L247:
	;
	v515 = v511
	v516 = v508
	goto L244
L248:
	;
	v511 = v468 + int32(3)
	if base.Ui32(v511) <= base.Ui32(v453) {
		v468 = v511
		v469 = v508
		goto L246
	} else {
		goto L261
	}
L249:
	;
	if base.Ui32(v482) < base.Ui32(v483) {
		v508 = v469
		goto L248
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468)+1)))
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v486 != v487 {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	v515 = v468
	v516 = v469
	goto L244
L253:
	;
	if base.Ui32(v486) < base.Ui32(v487) {
		v508 = v469
		goto L248
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468)+2)))
	v491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
	if v490 == v491 {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	v515 = v468
	v516 = v469
	goto L244
L257:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469))))
	*(*uint8)(unsafe.Add(mBase, uint32(v469))) = uint8(v482)
	*(*uint8)(unsafe.Add(mBase, uint32(v468))) = uint8(v493)
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+1)))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+1)) = uint8(v497)
	*(*uint8)(unsafe.Add(mBase, uint32(v468)+1)) = uint8(v496)
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v469)+2)))
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v469)+2)) = uint8(v501)
	*(*uint8)(unsafe.Add(mBase, uint32(v468)+2)) = uint8(v500)
	v508 = v469 + int32(3)
	goto L248
L258:
	;
	goto L259
L259:
	;
	if base.Ui32(v491) <= base.Ui32(v490) {
		v515 = v468
		v516 = v469
		goto L244
	} else {
		goto L260
	}
L260:
	;
	v508 = v469
	goto L248
L261:
	;
	goto L247
L262:
	;
	goto L243
L263:
	;
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515))))
	*(*uint8)(unsafe.Add(mBase, uint32(v515))) = uint8(v546)
	*(*uint8)(unsafe.Add(mBase, uint32(v534))) = uint8(v862)
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)))
	v866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v515)+1)) = uint8(v866)
	*(*uint8)(unsafe.Add(mBase, uint32(v534)+1)) = uint8(v865)
	v869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v515)+2)))
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v515)+2)) = uint8(v870)
	*(*uint8)(unsafe.Add(mBase, uint32(v534)+2)) = uint8(v869)
	v873 = int32(3)
	v451 = v515 + v873
	v452 = v516
	v453 = v534 - v873
	v456 = v537
	goto L242
L264:
	;
	v534 = v453
	v537 = v456
	goto L267
L265:
	;
	v581 = v453
	v584 = v456
	goto L266
L266:
	;
	v593 = v516 - v19
	v594 = v515 - v516
	if v593 < v594 {
		goto L284
	} else {
		goto L285
	}
L267:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534))))
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v546 != v547 {
		goto L270
	} else {
		goto L271
	}
L268:
	;
	v581 = v575
	v584 = v573
	goto L266
L269:
	;
	v575 = v534 - int32(3)
	if base.Ui32(v515) <= base.Ui32(v575) {
		v534 = v575
		v537 = v573
		goto L267
	} else {
		goto L282
	}
L270:
	;
	if base.Ui32(v547) <= base.Ui32(v546) {
		v573 = v537
		goto L269
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+1)))
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v550 != v551 {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	goto L263
L274:
	;
	if base.Ui32(v550) < base.Ui32(v551) {
		goto L263
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+2)))
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
	if v554 == v555 {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	v573 = v537
	goto L269
L278:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v537))))
	*(*uint8)(unsafe.Add(mBase, uint32(v534))) = uint8(v557)
	*(*uint8)(unsafe.Add(mBase, uint32(v537))) = uint8(v546)
	v560 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+1)))
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v537)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v534)+1)) = uint8(v561)
	*(*uint8)(unsafe.Add(mBase, uint32(v537)+1)) = uint8(v560)
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v534)+2)))
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v537)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v534)+2)) = uint8(v565)
	*(*uint8)(unsafe.Add(mBase, uint32(v537)+2)) = uint8(v564)
	v573 = v537 - int32(3)
	goto L269
L279:
	;
	goto L280
L280:
	;
	if base.Ui32(v554) < base.Ui32(v555) {
		goto L263
	} else {
		goto L281
	}
L281:
	;
	v573 = v537
	goto L269
L282:
	;
	goto L268
L283:
	;
	v721 = v584 - v581
	v724 = v57 - v584 - int32(3)
	if base.Ui32(v721) < base.Ui32(v724) {
		goto L299
	} else {
		goto L300
	}
L284:
	;
	v596 = v593
	goto L286
L285:
	;
	v596 = v594
	goto L286
L286:
	;
	if v596 == int32(0) {
		goto L283
	} else {
		goto L287
	}
L287:
	;
	v599 = v515 - v596
	v601 = v596 & int32(3)
	v602 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v596) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v610 = v602
	v612 = int32(0)
	goto L291
L289:
	;
	v663 = v602
	goto L290
L290:
	;
	v679 = v663
	v689 = v602
	goto L295
L291:
	;
	v625 = v19 + v610
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v625))))
	v627 = v610 + v599
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v627))))
	*(*uint8)(unsafe.Add(mBase, uint32(v625))) = uint8(v628)
	*(*uint8)(unsafe.Add(mBase, uint32(v627))) = uint8(v626)
	v632 = v610 | int32(1)
	v633 = v19 + v632
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633))))
	v635 = v632 + v599
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v635))))
	*(*uint8)(unsafe.Add(mBase, uint32(v633))) = uint8(v636)
	*(*uint8)(unsafe.Add(mBase, uint32(v635))) = uint8(v634)
	v640 = v610 | int32(2)
	v641 = v19 + v640
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641))))
	v643 = v640 + v599
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v643))))
	*(*uint8)(unsafe.Add(mBase, uint32(v641))) = uint8(v644)
	*(*uint8)(unsafe.Add(mBase, uint32(v643))) = uint8(v642)
	v648 = v610 | int32(3)
	v649 = v19 + v648
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649))))
	v651 = v648 + v599
	v652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651))))
	*(*uint8)(unsafe.Add(mBase, uint32(v649))) = uint8(v652)
	*(*uint8)(unsafe.Add(mBase, uint32(v651))) = uint8(v650)
	v655 = int32(4)
	v656 = v610 + v655
	v658 = v612 + v655
	if v658 != v596&int32(-4) {
		v610 = v656
		v612 = v658
		goto L291
	} else {
		goto L293
	}
L292:
	;
	if v601 == int32(0) {
		goto L283
	} else {
		goto L294
	}
L293:
	;
	goto L292
L294:
	;
	v663 = v656
	goto L290
L295:
	;
	v694 = v19 + v679
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694))))
	v696 = v679 + v599
	v697 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696))))
	*(*uint8)(unsafe.Add(mBase, uint32(v694))) = uint8(v697)
	*(*uint8)(unsafe.Add(mBase, uint32(v696))) = uint8(v695)
	v700 = int32(1)
	v703 = v689 + v700
	if v703 != v601 {
		v679 = v679 + v700
		v689 = v703
		goto L295
	} else {
		goto L297
	}
L296:
	;
	goto L283
L297:
	;
	goto L296
L298:
	;
	if base.Ui32(v721) < base.Ui32(v594) {
		goto L262
	} else {
		goto L313
	}
L299:
	;
	v726 = v721
	goto L301
L300:
	;
	v726 = v724
	goto L301
L301:
	;
	if v726 == int32(0) {
		goto L298
	} else {
		goto L302
	}
L302:
	;
	v729 = v57 - v726
	v731 = v726 & int32(3)
	v732 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v726) {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	v743 = v732
	v746 = int32(0)
	goto L306
L304:
	;
	v796 = v732
	goto L305
L305:
	;
	v809 = v732
	v812 = v796
	goto L310
L306:
	;
	v755 = v515 + v743
	v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755))))
	v757 = v743 + v729
	v758 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v757))))
	*(*uint8)(unsafe.Add(mBase, uint32(v755))) = uint8(v758)
	*(*uint8)(unsafe.Add(mBase, uint32(v757))) = uint8(v756)
	v762 = v743 | int32(1)
	v763 = v515 + v762
	v764 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v763))))
	v765 = v762 + v729
	v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765))))
	*(*uint8)(unsafe.Add(mBase, uint32(v763))) = uint8(v766)
	*(*uint8)(unsafe.Add(mBase, uint32(v765))) = uint8(v764)
	v770 = v743 | int32(2)
	v771 = v515 + v770
	v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v771))))
	v773 = v770 + v729
	v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v773))))
	*(*uint8)(unsafe.Add(mBase, uint32(v771))) = uint8(v774)
	*(*uint8)(unsafe.Add(mBase, uint32(v773))) = uint8(v772)
	v778 = v743 | int32(3)
	v779 = v515 + v778
	v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v779))))
	v781 = v778 + v729
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v781))))
	*(*uint8)(unsafe.Add(mBase, uint32(v779))) = uint8(v782)
	*(*uint8)(unsafe.Add(mBase, uint32(v781))) = uint8(v780)
	v785 = int32(4)
	v786 = v743 + v785
	v788 = v746 + v785
	if v788 != v726&int32(-4) {
		v743 = v786
		v746 = v788
		goto L306
	} else {
		goto L308
	}
L307:
	;
	if v731 == int32(0) {
		goto L298
	} else {
		goto L309
	}
L308:
	;
	goto L307
L309:
	;
	v796 = v786
	goto L305
L310:
	;
	v824 = v515 + v812
	v825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v824))))
	v826 = v812 + v729
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v826))))
	*(*uint8)(unsafe.Add(mBase, uint32(v824))) = uint8(v827)
	*(*uint8)(unsafe.Add(mBase, uint32(v826))) = uint8(v825)
	v830 = int32(1)
	v833 = v809 + v830
	if v833 != v731 {
		v809 = v833
		v812 = v812 + v830
		goto L310
	} else {
		goto L312
	}
L311:
	;
	goto L298
L312:
	;
	goto L311
L313:
	;
	v852 = int32(3)
	v853 = base.I32_div_u_s(v594, v852)
	F_trigram_qsort_unsigned(m, v19, v853)
	mBase = m.M
	v856 = base.I32_div_u_s(v721, v852)
	v857 = v57 - v721
	if base.Ui32(int32(21)) <= base.Ui32(v721) {
		v19 = v857
		v20 = v856
		goto L4
	} else {
		goto L314
	}
L314:
	;
	v885 = v857
	v886 = v856
	goto L2
L315:
	;
	goto L7
L316:
	;
	v903 = int32(3)
	v911 = v885 + v903
	goto L317
L317:
	;
	if base.Ui32(v911) <= base.Ui32(v885) {
		goto L319
	} else {
		goto L320
	}
L318:
	;
	goto L1
L319:
	;
	v994 = v911 + int32(3)
	if base.Ui32(v994) < base.Ui32(v885+v886*v903) {
		v911 = v994
		goto L317
	} else {
		goto L334
	}
L320:
	;
	v925 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v911))))
	v930 = v911
	goto L321
L321:
	;
	v943 = v930 - int32(3)
	v944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943))))
	if v925 != v944 {
		goto L324
	} else {
		goto L325
	}
L322:
	;
	goto L319
L323:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v930))) = uint8(v944)
	*(*uint8)(unsafe.Add(mBase, uint32(v943))) = uint8(v925)
	*(*uint8)(unsafe.Add(mBase, uint32(v930-int32(2)))) = uint8(v963)
	*(*uint8)(unsafe.Add(mBase, uint32(v930)+1)) = uint8(v962)
	v970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930)+2)))
	v972 = v930 - int32(1)
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v972))))
	*(*uint8)(unsafe.Add(mBase, uint32(v930)+2)) = uint8(v973)
	*(*uint8)(unsafe.Add(mBase, uint32(v972))) = uint8(v970)
	if base.Ui32(v885) < base.Ui32(v943) {
		v930 = v943
		goto L321
	} else {
		goto L333
	}
L324:
	;
	if base.Ui32(v944) < base.Ui32(v925) {
		goto L319
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930-int32(2)))))
	v954 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930)+1)))
	if v953 != v954 {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v949 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930-int32(2)))))
	v950 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930)+1)))
	v962 = v949
	v963 = v950
	goto L323
L328:
	;
	if base.Ui32(v954) <= base.Ui32(v953) {
		v962 = v953
		v963 = v954
		goto L323
	} else {
		goto L331
	}
L329:
	;
	goto L330
L330:
	;
	v959 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930-int32(1)))))
	v960 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930)+2)))
	if base.Ui32(v959) <= base.Ui32(v960) {
		goto L319
	} else {
		goto L332
	}
L331:
	;
	goto L319
L332:
	;
	v962 = v953
	v963 = v953
	goto L323
L333:
	;
	goto L322
L334:
	;
	goto L318
}
func F_trigram_qsort_unsigned_med3(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v8 != v9 {
		if base.Ui32(v8) < base.Ui32(v9) {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
			if v21 != v9 {
				if base.Ui32(v21) <= base.Ui32(v9) {
					if v8 != v21 {
						if base.Ui32(v21) <= base.Ui32(v8) {
							return l0
						} else {
							v70 = l2
							return v70
						}
					} else {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
						if v35 != v36 {
							if base.Ui32(v36) <= base.Ui32(v35) {
								return l0
							} else {
								v70 = l2
								return v70
							}
						} else {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
							if base.Ui32(v39) < base.Ui32(v40) {
								v70 = l2
								return v70
							} else {
								return l0
							}
						}
					}
				} else {
					v70 = l1
					return v70
				}
			} else {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
				if v24 != v25 {
					if base.Ui32(v25) <= base.Ui32(v24) {
						if v8 != v21 {
							if base.Ui32(v21) <= base.Ui32(v8) {
								return l0
							} else {
								v70 = l2
								return v70
							}
						} else {
							v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
							if v35 != v36 {
								if base.Ui32(v36) <= base.Ui32(v35) {
									return l0
								} else {
									v70 = l2
									return v70
								}
							} else {
								v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
								if base.Ui32(v39) < base.Ui32(v40) {
									v70 = l2
									return v70
								} else {
									return l0
								}
							}
						}
					} else {
						v70 = l1
						return v70
					}
				} else {
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
					if base.Ui32(v28) < base.Ui32(v29) {
						v70 = l1
						return v70
					} else {
						if v8 != v21 {
							if base.Ui32(v21) <= base.Ui32(v8) {
								return l0
							} else {
								v70 = l2
								return v70
							}
						} else {
							v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
							if v35 != v36 {
								if base.Ui32(v36) <= base.Ui32(v35) {
									return l0
								} else {
									v70 = l2
									return v70
								}
							} else {
								v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
								if base.Ui32(v39) < base.Ui32(v40) {
									v70 = l2
									return v70
								} else {
									return l0
								}
							}
						}
					}
				}
			}
		} else {
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
			if v47 != v9 {
				if base.Ui32(v9) < base.Ui32(v47) {
					if v8 != v47 {
						if base.Ui32(v47) <= base.Ui32(v8) {
							v70 = l2
						} else {
							v70 = l0
						}
					} else {
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
						v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
						if v61 != v62 {
							if base.Ui32(v62) <= base.Ui32(v61) {
								v70 = l2
							} else {
								v70 = l0
							}
						} else {
							v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
							v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
							if base.Ui32(v65) < base.Ui32(v66) {
								v70 = l0
							} else {
								v70 = l2
							}
						}
					}
				} else {
					v70 = l1
				}
			} else {
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
				if v50 != v51 {
					if base.Ui32(v50) < base.Ui32(v51) {
						if v8 != v47 {
							if base.Ui32(v47) <= base.Ui32(v8) {
								v70 = l2
							} else {
								v70 = l0
							}
						} else {
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
							if v61 != v62 {
								if base.Ui32(v62) <= base.Ui32(v61) {
									v70 = l2
								} else {
									v70 = l0
								}
							} else {
								v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
								v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
								if base.Ui32(v65) < base.Ui32(v66) {
									v70 = l0
								} else {
									v70 = l2
								}
							}
						}
					} else {
						v70 = l1
					}
				} else {
					v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
					if base.Ui32(v55) < base.Ui32(v54) {
						v70 = l1
					} else {
						if v8 != v47 {
							if base.Ui32(v47) <= base.Ui32(v8) {
								v70 = l2
							} else {
								v70 = l0
							}
						} else {
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
							if v61 != v62 {
								if base.Ui32(v62) <= base.Ui32(v61) {
									v70 = l2
								} else {
									v70 = l0
								}
							} else {
								v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
								v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
								if base.Ui32(v65) < base.Ui32(v66) {
									v70 = l0
								} else {
									v70 = l2
								}
							}
						}
					}
				}
			}
			return v70
		}
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		if v12 != v13 {
			if base.Ui32(v13) <= base.Ui32(v12) {
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
				if v47 != v9 {
					if base.Ui32(v9) < base.Ui32(v47) {
						if v8 != v47 {
							if base.Ui32(v47) <= base.Ui32(v8) {
								v70 = l2
							} else {
								v70 = l0
							}
						} else {
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
							if v61 != v62 {
								if base.Ui32(v62) <= base.Ui32(v61) {
									v70 = l2
								} else {
									v70 = l0
								}
							} else {
								v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
								v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
								if base.Ui32(v65) < base.Ui32(v66) {
									v70 = l0
								} else {
									v70 = l2
								}
							}
						}
					} else {
						v70 = l1
					}
				} else {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
					if v50 != v51 {
						if base.Ui32(v50) < base.Ui32(v51) {
							if v8 != v47 {
								if base.Ui32(v47) <= base.Ui32(v8) {
									v70 = l2
								} else {
									v70 = l0
								}
							} else {
								v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
								if v61 != v62 {
									if base.Ui32(v62) <= base.Ui32(v61) {
										v70 = l2
									} else {
										v70 = l0
									}
								} else {
									v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
									v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
									if base.Ui32(v65) < base.Ui32(v66) {
										v70 = l0
									} else {
										v70 = l2
									}
								}
							}
						} else {
							v70 = l1
						}
					} else {
						v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
						if base.Ui32(v55) < base.Ui32(v54) {
							v70 = l1
						} else {
							if v8 != v47 {
								if base.Ui32(v47) <= base.Ui32(v8) {
									v70 = l2
								} else {
									v70 = l0
								}
							} else {
								v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
								if v61 != v62 {
									if base.Ui32(v62) <= base.Ui32(v61) {
										v70 = l2
									} else {
										v70 = l0
									}
								} else {
									v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
									v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
									if base.Ui32(v65) < base.Ui32(v66) {
										v70 = l0
									} else {
										v70 = l2
									}
								}
							}
						}
					}
				}
				return v70
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
				if v21 != v9 {
					if base.Ui32(v21) <= base.Ui32(v9) {
						if v8 != v21 {
							if base.Ui32(v21) <= base.Ui32(v8) {
								return l0
							} else {
								v70 = l2
								return v70
							}
						} else {
							v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
							if v35 != v36 {
								if base.Ui32(v36) <= base.Ui32(v35) {
									return l0
								} else {
									v70 = l2
									return v70
								}
							} else {
								v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
								if base.Ui32(v39) < base.Ui32(v40) {
									v70 = l2
									return v70
								} else {
									return l0
								}
							}
						}
					} else {
						v70 = l1
						return v70
					}
				} else {
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
					if v24 != v25 {
						if base.Ui32(v25) <= base.Ui32(v24) {
							if v8 != v21 {
								if base.Ui32(v21) <= base.Ui32(v8) {
									return l0
								} else {
									v70 = l2
									return v70
								}
							} else {
								v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
								if v35 != v36 {
									if base.Ui32(v36) <= base.Ui32(v35) {
										return l0
									} else {
										v70 = l2
										return v70
									}
								} else {
									v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
									v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
									if base.Ui32(v39) < base.Ui32(v40) {
										v70 = l2
										return v70
									} else {
										return l0
									}
								}
							}
						} else {
							v70 = l1
							return v70
						}
					} else {
						v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
						v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
						if base.Ui32(v28) < base.Ui32(v29) {
							v70 = l1
							return v70
						} else {
							if v8 != v21 {
								if base.Ui32(v21) <= base.Ui32(v8) {
									return l0
								} else {
									v70 = l2
									return v70
								}
							} else {
								v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
								if v35 != v36 {
									if base.Ui32(v36) <= base.Ui32(v35) {
										return l0
									} else {
										v70 = l2
										return v70
									}
								} else {
									v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
									v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
									if base.Ui32(v39) < base.Ui32(v40) {
										v70 = l2
										return v70
									} else {
										return l0
									}
								}
							}
						}
					}
				}
			}
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			if base.Ui32(v17) <= base.Ui32(v16) {
				v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
				if v47 != v9 {
					if base.Ui32(v9) < base.Ui32(v47) {
						if v8 != v47 {
							if base.Ui32(v47) <= base.Ui32(v8) {
								v70 = l2
							} else {
								v70 = l0
							}
						} else {
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
							if v61 != v62 {
								if base.Ui32(v62) <= base.Ui32(v61) {
									v70 = l2
								} else {
									v70 = l0
								}
							} else {
								v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
								v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
								if base.Ui32(v65) < base.Ui32(v66) {
									v70 = l0
								} else {
									v70 = l2
								}
							}
						}
					} else {
						v70 = l1
					}
				} else {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
					if v50 != v51 {
						if base.Ui32(v50) < base.Ui32(v51) {
							if v8 != v47 {
								if base.Ui32(v47) <= base.Ui32(v8) {
									v70 = l2
								} else {
									v70 = l0
								}
							} else {
								v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
								if v61 != v62 {
									if base.Ui32(v62) <= base.Ui32(v61) {
										v70 = l2
									} else {
										v70 = l0
									}
								} else {
									v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
									v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
									if base.Ui32(v65) < base.Ui32(v66) {
										v70 = l0
									} else {
										v70 = l2
									}
								}
							}
						} else {
							v70 = l1
						}
					} else {
						v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
						if base.Ui32(v55) < base.Ui32(v54) {
							v70 = l1
						} else {
							if v8 != v47 {
								if base.Ui32(v47) <= base.Ui32(v8) {
									v70 = l2
								} else {
									v70 = l0
								}
							} else {
								v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
								if v61 != v62 {
									if base.Ui32(v62) <= base.Ui32(v61) {
										v70 = l2
									} else {
										v70 = l0
									}
								} else {
									v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
									v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
									if base.Ui32(v65) < base.Ui32(v66) {
										v70 = l0
									} else {
										v70 = l2
									}
								}
							}
						}
					}
				}
				return v70
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
				if v21 != v9 {
					if base.Ui32(v21) <= base.Ui32(v9) {
						if v8 != v21 {
							if base.Ui32(v21) <= base.Ui32(v8) {
								return l0
							} else {
								v70 = l2
								return v70
							}
						} else {
							v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
							if v35 != v36 {
								if base.Ui32(v36) <= base.Ui32(v35) {
									return l0
								} else {
									v70 = l2
									return v70
								}
							} else {
								v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
								if base.Ui32(v39) < base.Ui32(v40) {
									v70 = l2
									return v70
								} else {
									return l0
								}
							}
						}
					} else {
						v70 = l1
						return v70
					}
				} else {
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
					if v24 != v25 {
						if base.Ui32(v25) <= base.Ui32(v24) {
							if v8 != v21 {
								if base.Ui32(v21) <= base.Ui32(v8) {
									return l0
								} else {
									v70 = l2
									return v70
								}
							} else {
								v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
								if v35 != v36 {
									if base.Ui32(v36) <= base.Ui32(v35) {
										return l0
									} else {
										v70 = l2
										return v70
									}
								} else {
									v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
									v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
									if base.Ui32(v39) < base.Ui32(v40) {
										v70 = l2
										return v70
									} else {
										return l0
									}
								}
							}
						} else {
							v70 = l1
							return v70
						}
					} else {
						v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
						v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
						if base.Ui32(v28) < base.Ui32(v29) {
							v70 = l1
							return v70
						} else {
							if v8 != v21 {
								if base.Ui32(v21) <= base.Ui32(v8) {
									return l0
								} else {
									v70 = l2
									return v70
								}
							} else {
								v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
								if v35 != v36 {
									if base.Ui32(v36) <= base.Ui32(v35) {
										return l0
									} else {
										v70 = l2
										return v70
									}
								} else {
									v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
									v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
									if base.Ui32(v39) < base.Ui32(v40) {
										v70 = l2
										return v70
									} else {
										return l0
									}
								}
							}
						}
					}
				}
			}
		}
	}
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
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+95)) = uint8(v2)
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
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
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
func F_try_nestloop_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int64, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
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
	var v382 int32
	_ = v382
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 float64
	_ = v406
	var v407 float64
	_ = v407
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 float64
	_ = v446
	var v453 float64
	_ = v453
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	v9 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(96)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v21 = v20
	goto L3
L2:
	;
	v21 = v9
	goto L3
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v24 = v23
	goto L6
L5:
	;
	v24 = v9
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	if v28 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	m.G0 = v17 + int32(96)
	return
L8:
	;
	v29 = F_bms_is_member(m, v28, v21)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)+252))
	if v35 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	return
L12:
	;
	if v29 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	v33 = F_bms_is_member(m, v32, v24)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	if v33 != 0 {
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
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)+252))
	if v39 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v43 = v42
	goto L21
L20:
	;
	v43 = v39
	goto L21
L21:
	;
	if v21 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	F_bms_free(m, v52)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L11
	} else {
		goto L187
	}
L23:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v207 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L24:
	;
	if v52 == int32(0) {
		goto L23
	} else {
		goto L31
	}
L25:
	;
	v46 = F_bms_copy(m, v24)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L11
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v48 = F_bms_union(m, v24, v21)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L11
	} else {
		goto L29
	}
L28:
	;
	v52 = v46
	goto L24
L29:
	;
	v50 = F_bms_del_members(m, v48, v43)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	v52 = v50
	goto L24
L31:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l7)+32))
	v56 = int32(0)
	if base.B2i32(v52 == v56)|base.B2i32(v55 == v56) != 0 {
		v101 = v56
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v101 != 0 {
		goto L23
	} else {
		goto L45
	}
L33:
	;
	goto L32
L34:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v66 < v67 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v69 = v66
	goto L37
L36:
	;
	v69 = v67
	goto L37
L37:
	;
	if v69 <= int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v72 = int32(1)
	goto L40
L39:
	;
	v72 = v69
	goto L40
L40:
	;
	v73 = int32(8)
	v78 = int32(0)
	goto L41
L41:
	;
	v85 = v78 << (uint(int32(2)) % 32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v55+v73+v85)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v52+v73+v85)))
	v90 = v87 & v89
	v92 = base.B2i32(v90 != int32(0))
	if v90 != 0 {
		v101 = v92
		goto L33
	} else {
		goto L43
	}
L42:
	;
	v101 = v92
	goto L33
L43:
	;
	v94 = v78 + int32(1)
	if v94 != v72 {
		v78 = v94
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v102 = int32(0)
	if base.B2i32(v21 == v102)|base.B2i32(v43 == v102) != 0 {
		v147 = v102
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v147 == int32(0) {
		goto L22
	} else {
		goto L59
	}
L47:
	;
	goto L46
L48:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v112 < v113 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v115 = v112
	goto L51
L50:
	;
	v115 = v113
	goto L51
L51:
	;
	if v115 <= int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v118 = int32(1)
	goto L54
L53:
	;
	v118 = v115
	goto L54
L54:
	;
	v119 = int32(8)
	v124 = int32(0)
	goto L55
L55:
	;
	v131 = v124 << (uint(int32(2)) % 32)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v43+v119+v131)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v21+v119+v131)))
	v136 = v133 & v135
	v138 = base.B2i32(v136 != int32(0))
	if v136 != 0 {
		v147 = v138
		goto L47
	} else {
		goto L57
	}
L56:
	;
	v147 = v138
	goto L47
L57:
	;
	v140 = v124 + int32(1)
	if v140 != v118 {
		v124 = v140
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	if v21 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v204 == int32(0) {
		goto L22
	} else {
		goto L74
	}
L61:
	;
	v204 = int32(0)
	goto L60
L62:
	;
	goto L63
L63:
	;
	v157 = int32(1)
	if v43 == int32(0) {
		v194 = v157
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v204 = v194
	goto L60
L65:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v161 < v160 {
		v194 = v157
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v163 = int32(1)
	if v160 <= v163 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v166 = v163
	goto L69
L68:
	;
	v166 = v160
	goto L69
L69:
	;
	v167 = int32(8)
	v172 = int32(0)
	goto L70
L70:
	;
	v179 = v172 << (uint(int32(2)) % 32)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v21+v167+v179)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v43+v167+v179)))
	v186 = v181 & (v183 ^ int32(-1))
	v188 = base.B2i32(v186 != int32(0))
	if v186 != 0 {
		v194 = v188
		goto L64
	} else {
		goto L72
	}
L71:
	;
	v194 = v188
	goto L64
L72:
	;
	v190 = v172 + int32(1)
	if v190 != v166 {
		v172 = v190
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
	F_initial_cost_nestloop(m, l0, v17, l5, l6|int64(262144), l2, l3, l7)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L11
	} else {
		goto L147
	}
L76:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+252))
	v213 = int32(0)
	if base.B2i32(v210 == v213)|base.B2i32(v212 == v213) != 0 {
		v258 = v213
		goto L78
	} else {
		goto L79
	}
L77:
	;
	if v258 == int32(0) {
		goto L75
	} else {
		goto L90
	}
L78:
	;
	goto L77
L79:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v212)+4))
	if v223 < v224 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v226 = v223
	goto L82
L81:
	;
	v226 = v224
	goto L82
L82:
	;
	if v226 <= int32(1) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v229 = int32(1)
	goto L85
L84:
	;
	v229 = v226
	goto L85
L85:
	;
	v230 = int32(8)
	v235 = int32(0)
	goto L86
L86:
	;
	v242 = v235 << (uint(int32(2)) % 32)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v212+v230+v242)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v210+v230+v242)))
	v247 = v244 & v246
	v249 = base.B2i32(v247 != int32(0))
	if v247 != 0 {
		v258 = v249
		goto L78
	} else {
		goto L88
	}
L87:
	;
	v258 = v249
	goto L78
L88:
	;
	v251 = v235 + int32(1)
	if v251 != v229 {
		v235 = v251
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v264 = int32(1)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v265 == int32(0) {
		v392 = v264
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v398 == int32(0) {
		goto L22
	} else {
		goto L146
	}
L92:
	;
	v398 = v392
	goto L91
L93:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v261)+252))
	v270 = F_bms_overlap(m, v268, v269)
	mBase = m.M
	if v270 == int32(0) {
		v392 = v264
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v273 = int32(0)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	switch v274 - int32(282) {
	case 0, 1:
		goto L95
	default:
		v392 = v273
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
	case 16:
		goto L96
	case 18, 19, 20:
		goto L100
	}
L95:
	;
	v392 = int32(1)
	goto L92
L96:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v382 = F_path_is_reparameterizable_by_child(m, v381, v261)
	mBase = m.M
	if v382 == int32(0) {
		v392 = v273
		goto L92
	} else {
		goto L145
	}
L97:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v380 = F_path_is_reparameterizable_by_child(m, v379, v261)
	mBase = m.M
	if v380 != 0 {
		goto L95
	} else {
		goto L144
	}
L98:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v378 = F_path_is_reparameterizable_by_child(m, v377, v261)
	mBase = m.M
	if v378 != 0 {
		goto L95
	} else {
		goto L143
	}
L99:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v355 == int32(0) {
		goto L95
	} else {
		goto L135
	}
L100:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l3)+80))
	v350 = F_path_is_reparameterizable_by_child(m, v349, v261)
	mBase = m.M
	if v350 == int32(0) {
		v392 = v273
		goto L92
	} else {
		goto L133
	}
L101:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	if v327 == int32(0) {
		goto L95
	} else {
		goto L125
	}
L102:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v323 == int32(0) {
		goto L95
	} else {
		goto L123
	}
L103:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v301 == int32(0) {
		goto L95
	} else {
		goto L115
	}
L104:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v279 == int32(0) {
		goto L95
	} else {
		goto L107
	}
L105:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v278 = F_path_is_reparameterizable_by_child(m, v277, v261)
	mBase = m.M
	if v278 != 0 {
		goto L95
	} else {
		goto L106
	}
L106:
	;
	v392 = v273
	goto L92
L107:
	;
	v282 = int32(0)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	if v283 <= v282 {
		goto L95
	} else {
		goto L108
	}
L108:
	;
	v286 = v282
	goto L109
L109:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v279)+12))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v290+v286<<(uint(int32(2))%32))))
	v295 = F_path_is_reparameterizable_by_child(m, v294, v261)
	mBase = m.M
	if v295 != 0 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v398 = int32(0)
	goto L91
L111:
	;
	v297 = v286 + int32(1)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	if v297 < v298 {
		v286 = v297
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
	v304 = int32(0)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	if v305 <= v304 {
		goto L95
	} else {
		goto L116
	}
L116:
	;
	v308 = v304
	goto L117
L117:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v301)+12))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v312+v308<<(uint(int32(2))%32))))
	v317 = F_path_is_reparameterizable_by_child(m, v316, v261)
	mBase = m.M
	if v317 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v398 = int32(0)
	goto L91
L119:
	;
	v319 = v308 + int32(1)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v301)+4))
	if v319 < v320 {
		v308 = v319
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
	v326 = F_path_is_reparameterizable_by_child(m, v323, v261)
	mBase = m.M
	if v326 != 0 {
		goto L95
	} else {
		goto L124
	}
L124:
	;
	v392 = v273
	goto L92
L125:
	;
	v330 = int32(0)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if v331 <= v330 {
		goto L95
	} else {
		goto L126
	}
L126:
	;
	v334 = v330
	goto L127
L127:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v327)+12))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v338+v334<<(uint(int32(2))%32))))
	v343 = F_path_is_reparameterizable_by_child(m, v342, v261)
	mBase = m.M
	if v343 != 0 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v398 = int32(0)
	goto L91
L129:
	;
	v345 = v334 + int32(1)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if v345 < v346 {
		v334 = v345
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
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	v354 = F_path_is_reparameterizable_by_child(m, v353, v261)
	mBase = m.M
	if v354 != 0 {
		goto L95
	} else {
		goto L134
	}
L134:
	;
	v392 = v273
	goto L92
L135:
	;
	v358 = int32(0)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v355)+4))
	if v359 <= v358 {
		goto L95
	} else {
		goto L136
	}
L136:
	;
	v362 = v358
	goto L137
L137:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v355)+12))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v366+v362<<(uint(int32(2))%32))))
	v371 = F_path_is_reparameterizable_by_child(m, v370, v261)
	mBase = m.M
	if v371 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v398 = int32(0)
	goto L91
L139:
	;
	v373 = v362 + int32(1)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v355)+4))
	if v373 < v374 {
		v362 = v373
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
	v392 = v273
	goto L92
L144:
	;
	v392 = v273
	goto L92
L145:
	;
	goto L95
L146:
	;
	goto L75
L147:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v406 = *(*float64)(unsafe.Add(mBase, uint32(v17)+8))
	v407 = *(*float64)(unsafe.Add(mBase, uint32(v17)+16))
	v408 = int32(0)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v412 == v408 {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	if v492 == int32(0) {
		goto L22
	} else {
		goto L184
	}
L149:
	;
	v492 = int32(1)
	goto L148
L150:
	;
	goto L151
L151:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	if v416 <= int32(0) {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v492 = int32(1)
	goto L148
L153:
	;
	goto L154
L154:
	;
	if v52 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v421 = int32(0)
	goto L157
L156:
	;
	v421 = l4
	goto L157
L157:
	;
	if v52 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v424 = int32(25)
	goto L160
L159:
	;
	v424 = int32(24)
	goto L160
L160:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v424))))
	v434 = v408
	goto L161
L161:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v412)+12))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v437+v434<<(uint(int32(2))%32))))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)+40))
	if v405 != v442 {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v492 = v480
	goto L148
L163:
	;
	if v426 != 0 {
		goto L171
	} else {
		goto L172
	}
L164:
	;
	if v442 <= v405 {
		goto L163
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v446 = *(*float64)(unsafe.Add(mBase, uint32(v441)+56))
	if base.F64_le(v407, base.F64_mul(v446, float64(1.01))) == int32(0) {
		goto L163
	} else {
		goto L168
	}
L167:
	;
	v492 = int32(1)
	goto L148
L168:
	;
	v492 = int32(1)
	goto L148
L169:
	;
	goto L162
L170:
	;
	v474 = int32(1)
	v476 = v434 + v474
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	if v476 < v477 {
		v434 = v476
		goto L161
	} else {
		goto L183
	}
L171:
	;
	v453 = *(*float64)(unsafe.Add(mBase, uint32(v441)+48))
	if base.F64_gt(v406, base.F64_mul(v453, float64(1.01))) == int32(0) {
		goto L170
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v441)+16))
	if v460 != 0 {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	goto L173
L175:
	;
	v463 = int32(0)
	goto L177
L176:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v441)+64))
	v463 = v462
	goto L177
L177:
	;
	v464 = F_compare_pathkeys(m, v421, v463)
	mBase = m.M
	if v464&int32(-3) != 0 {
		goto L170
	} else {
		goto L178
	}
L178:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v441)+16))
	if v467 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v467)+4))
	v470 = v468
	goto L181
L180:
	;
	v470 = int32(0)
	goto L181
L181:
	;
	v471 = F_bms_equal(m, v52, v470)
	mBase = m.M
	if v471 != 0 {
		v480 = int32(0)
		goto L169
	} else {
		goto L182
	}
L182:
	;
	goto L170
L183:
	;
	v480 = v474
	goto L169
L184:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v496 = F_create_nestloop_path(m, l0, l1, l5, v17, l7, l2, l3, v495, l4, v52)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L11
	} else {
		goto L185
	}
L185:
	;
	F_add_path(m, l1, v496)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
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
func F_tsvectorrecv(m *base.Module, l0 int32) int64 {
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	v2 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = F_pq_getmsgint(m, v19, int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L7
	} else {
		goto L105
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L7
	} else {
		goto L102
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L7
	} else {
		goto L99
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L7
	} else {
		goto L96
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L7
	} else {
		goto L93
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L7
	} else {
		goto L90
	}
L7:
	;
	return int64(0)
L8:
	;
	if base.Ui32(v21) < base.Ui32(int32(_a_F_tsvectorrecv_0)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v28 = v21 << (uint(int32(2)) % 32)
	v30 = v28 + int32(8)
	v32 = v30 << (uint(int32(1)) % 32)
	v33 = F_palloc0(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L7
	} else {
		goto L87
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v21
	if v21 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(32)
	return base.I64_extend_i32_u(v33)
L14:
	;
	goto L15
L15:
	;
	v44 = v32
	v45 = v33
	v46 = v2
	v55 = v2
	v56 = v2
	goto L16
L16:
	;
	v62 = F_pq_getmsgstring(m, v19)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L7
	} else {
		goto L18
	}
L17:
	;
	if int32(_a_F_tsvectorrecv_0) <= v318 {
		goto L1
	} else {
		goto L82
	}
L18:
	;
	v65 = F_pq_getmsgint(m, v19, int32(2))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v67 = F_strlen(m, v62)
	mBase = m.M
	if v67 == int32(0) {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	if base.Ui32(int32(2048)) <= base.Ui32(v67) {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	if int32(_a_F_tsvectorrecv_0) <= v46 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v75 = v65 & int32(_a_F_tsvectorrecv_1)
	if base.Ui32(int32(256)) < base.Ui32(v75) {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v78 = v46 + v67
	v79 = int32(1)
	v82 = (v78 + v79) & int32(-2)
	v84 = v75 << (uint(v79) % 32)
	v86 = v82 + (v28 + int32(10) + v84)
	if base.Ui32(v44) <= base.Ui32(v86) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v88 = v44
	v89 = v45
	goto L27
L25:
	;
	v111 = v44
	v112 = v45
	goto L26
L26:
	;
	v130 = v112 + int32(8)
	v133 = v130 + v55<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v67<<(uint(int32(1))%32) | base.B2i32(v75 != int32(0)) | v46<<(uint(int32(12))%32)
	if v67 != 0 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v107 = v88 << (uint(int32(1)) % 32)
	v108 = F_repalloc(m, v89, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L7
	} else {
		goto L29
	}
L28:
	;
	v111 = v107
	v112 = v108
	goto L26
L29:
	;
	if base.Ui32(v107) <= base.Ui32(v86) {
		v88 = v107
		v89 = v108
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	base.MemoryCopy(m, v130+v143<<(uint(int32(2))%32)+v46, v62, v67)
	goto L33
L32:
	;
	goto L33
L33:
	;
	if v55 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v152 = v130 + v149<<(uint(int32(2))%32)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v154 = int32(12)
	v157 = int32(1)
	v159 = int32(2047)
	v160 = int32(base.Ui32(v153)>>(uint(v157)%32)) & v159
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v133-int32(4))))
	v170 = int32(base.Ui32(v163)>>(uint(v157)%32)) & v159
	if v160 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v202 = v56
	goto L36
L36:
	;
	if v75 != 0 {
		goto L65
	} else {
		goto L66
	}
L37:
	;
	v202 = base.B2i32(v196 <= int32(0)) | v56
	goto L36
L38:
	;
	goto L42
L39:
	;
	goto L40
L40:
	;
	if v170 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	goto L43
L43:
	;
	v176 = int32(0)
	if v176 < v170 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v179 = int32(-1)
	goto L46
L45:
	;
	v179 = v176
	goto L46
L46:
	;
	v196 = v179
	goto L37
L47:
	;
	v196 = base.B2i32(int32(0) < v160)
	goto L37
L48:
	;
	goto L49
L49:
	;
	if base.Ui32(v160) < base.Ui32(v170) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v185 = v160
	goto L52
L51:
	;
	v185 = v170
	goto L52
L52:
	;
	v186 = F_memcmp(m, v152+int32(base.Ui32(v153)>>(uint(v154)%32)), v152+int32(base.Ui32(v163)>>(uint(v154)%32)), v185)
	mBase = m.M
	goto L55
L53:
	;
	v196 = v194
	goto L37
L55:
	;
	goto L56
L56:
	;
	if v186 != 0 {
		v194 = v186
		goto L53
	} else {
		goto L58
	}
L58:
	;
	if v160 == v170 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v196 = int32(0)
	goto L37
L60:
	;
	goto L61
L61:
	;
	if v160 < v170 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v193 = int32(-1)
	goto L64
L63:
	;
	v193 = int32(1)
	goto L64
L64:
	;
	v194 = v193
	goto L53
L65:
	;
	if v78 == v82 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v318 = v78
	goto L67
L67:
	;
	v320 = v55 + int32(1)
	if v320 != v21 {
		v44 = v111
		v45 = v112
		v46 = v318
		v55 = v320
		v56 = v202
		goto L16
	} else {
		goto L81
	}
L68:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v213 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v130+v212<<(uint(v213)%32)+v211))) = uint16(v65)
	v218 = int32(1)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v235 = v130 + v219<<(uint(v213)%32) + (int32(base.Ui32(v223)>>(uint(int32(12))%32))+int32(base.Ui32(v223)>>(uint(v218)%32))&int32(2047)+v218)&int32(_a_F_tsvectorrecv_2)
	v237 = F_pq_getmsgint(m, v19, v213)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L7
	} else {
		goto L72
	}
L69:
	;
	v211 = v78
	goto L68
L70:
	;
	goto L71
L71:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v209 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v130+v204<<(uint(int32(2))%32)+v78))) = uint8(v209)
	v211 = v82
	goto L68
L72:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v235)+2)) = uint16(v237)
	if v75 != int32(1) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v247 = v218
	goto L76
L74:
	;
	goto L75
L75:
	;
	v318 = v211 + v84 + int32(2)
	goto L67
L76:
	;
	v263 = v247 << (uint(int32(1)) % 32)
	v266 = F_pq_getmsgint(m, v19, int32(2))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L7
	} else {
		goto L78
	}
L77:
	;
	goto L75
L78:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v235+int32(2)+v263))) = uint16(v266)
	v269 = int32(_a_F_tsvectorrecv_3)
	v272 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v235+v263))))
	if base.Ui32(v266&v269) <= base.Ui32(v272&v269) {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	v277 = v247 + int32(1)
	if v277 != v75 {
		v247 = v277
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	goto L17
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = (v318 + v30) << (uint(int32(2)) % 32)
	if v202&int32(1) != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v331 = v112 + int32(8)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	F_qsort_arg(m, v331, v332, int32(4), int32(1732), v331+v332<<(uint(int32(2))%32))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L7
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	return base.I64_extend_i32_u(v112)
L86:
	;
	goto L85
L87:
	;
	F_errmsg_internal(m, int32(_a_F_tsvectorrecv_4), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_tsvectorrecv_5), int32(465), int32(_a_F_tsvectorrecv_6))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L7
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	F_errmsg_internal(m, int32(_a_F_tsvectorrecv_7), int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L7
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_tsvectorrecv_5), int32(487), int32(_a_F_tsvectorrecv_6))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L7
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errmsg_internal(m, int32(_a_F_tsvectorrecv_8), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L7
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_tsvectorrecv_5), int32(489), int32(_a_F_tsvectorrecv_6))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L7
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
	F_errmsg_internal(m, int32(_a_F_tsvectorrecv_9), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L7
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_tsvectorrecv_5), int32(492), int32(_a_F_tsvectorrecv_6))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L7
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	F_errmsg_internal(m, int32(_a_F_tsvectorrecv_10), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L7
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_tsvectorrecv_5), int32(495), int32(_a_F_tsvectorrecv_6))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L7
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	F_errmsg_internal(m, int32(_a_F_tsvectorrecv_11), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_tsvectorrecv_5), int32(546), int32(_a_F_tsvectorrecv_6))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L7
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errmsg_internal(m, int32(_a_F_tsvectorrecv_9), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L7
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_tsvectorrecv_5), int32(560), int32(_a_F_tsvectorrecv_6))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tuplehash_lookup_hash_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v97 int32
	_ = v97
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = v16 & l1
	v20 = v15 + v17*int32(12)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v97
L2:
	;
	v24 = v15
	v25 = v16
	v26 = v20
	v29 = v17
	goto L5
L3:
	;
	goto L4
L4:
	;
	v97 = int32(0)
	goto L1
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v32 == l1 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+36))
	v39 = F_ExecStoreMinimalTuple(m, v36, v37, int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v71 = v24
	v72 = v25
	goto L9
L9:
	;
	v77 = v72 & (v29 + int32(1))
	v80 = v71 + v77*int32(12)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v81 != 0 {
		v24 = v71
		v25 = v72
		v26 = v80
		v29 = v77
		goto L5
	} else {
		goto L19
	}
L10:
	;
	return int32(0)
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	if v46 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	F_MemoryContextReset(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v52 = int32(_a_F_tuplehash_lookup_hash_internal_0)
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_tuplehash_lookup_hash_internal[0]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplehash_lookup_hash_internal[0])) = v55
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	v60 = m.T0[v59].(func(*base.Module, int32, int32, int32) int64)(m, v46, v35, v13+int32(15))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	v97 = v26
	goto L1
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_tuplehash_lookup_hash_internal[0])) = v53
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	F_MemoryContextReset(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	if v60 != int64(0) {
		v97 = v26
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v71 = v70
	v72 = v69
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
