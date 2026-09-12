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
							v41 = int32(4442576)
							v42 = *(*int32)(unsafe.Add(mBase, _consts[0]))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							v46 = *(*int32)(unsafe.Add(mBase, _consts[300]))
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v46
							v48 = F_CreateTupleDescCopy(m, v43)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								v51 = F_MakeSingleTupleTableSlot(m, v48, int32(1575684))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v51
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v42
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
							v41 = int32(4442576)
							v42 = *(*int32)(unsafe.Add(mBase, _consts[0]))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							v46 = *(*int32)(unsafe.Add(mBase, _consts[300]))
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v46
							v48 = F_CreateTupleDescCopy(m, v43)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								v51 = F_MakeSingleTupleTableSlot(m, v48, int32(1575684))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v51
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v42
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
							F_errmsg_internal(m, int32(464135), v10)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								F_errfinish(m, int32(472360), int32(5622), int32(366120))
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
							v41 = int32(4442576)
							v42 = *(*int32)(unsafe.Add(mBase, _consts[0]))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							v46 = *(*int32)(unsafe.Add(mBase, _consts[300]))
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v46
							v48 = F_CreateTupleDescCopy(m, v43)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								v51 = F_MakeSingleTupleTableSlot(m, v48, int32(1575684))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v37)+32)) = v51
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v42
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v95 int64
	_ = v95
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v122 int64
	_ = v122
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v131 int64
	_ = v131
	var v134 int64
	_ = v134
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v148 int64
	_ = v148
	var v158 int64
	_ = v158
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v172 int64
	_ = v172
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v15 = l1 & int64(281474976710655)
	v19 = int64(base.Ui64(l1)>>(uint(int64(48))%64)) & int64(32767)
	v20 = base.I32_wrap_i64(v19)
	if base.Ui32(v20-int32(15361)) <= base.Ui32(int32(2045)) {
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
		v165 = v48
		v172 = base.I64_extend_i32_u(v47) + base.I64_extend_i32_u(v20-int32(15360))
	} else {
		if l0|v15 == int64(0) {
			if base.Ui32(int32(17406)) < base.Ui32(v20) {
				v165 = int64(0)
				v172 = int64(2047)
			} else {
				v71 = base.B2i32(v19 == int64(0))
				if v19 == int64(0) {
					v72 = int32(15360)
				} else {
					v72 = int32(15361)
				}
				v73 = v72 - v20
				if int32(112) < v73 {
					v76 = int64(0)
					v165 = v76
					v172 = v76
				} else {
					v79 = v12 + int32(16)
					if v19 == int64(0) {
						v82 = v15
					} else {
						v82 = v15 | int64(281474976710656)
					}
					v84 = int32(128) - v73
					if v84&int32(64) != 0 {
						v103 = int64(0)
						v104 = l0 << (uint(base.I64_extend_i32_u(v84+int32(-64))) % 64)
					} else {
						if v84 == int32(0) {
							v103 = l0
							v104 = v82
						} else {
							v95 = base.I64_extend_i32_u(v84)
							v103 = l0 << (uint(v95) % 64)
							v104 = v82<<(uint(v95)%64) | int64(base.Ui64(l0)>>(uint(base.I64_extend_i32_u(int32(64)-v84))%64))
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v79))) = v103
					*(*int64)(unsafe.Add(mBase, uint32(v79)+8)) = v104
					if v73&int32(64) != 0 {
						v126 = int64(base.Ui64(v82) >> (uint(base.I64_extend_i32_u(v73+int32(-64))) % 64))
						v127 = int64(0)
					} else {
						if v73 == int32(0) {
							v126 = l0
							v127 = v82
						} else {
							v122 = base.I64_extend_i32_u(v73)
							v126 = v82<<(uint(base.I64_extend_i32_u(int32(64)-v73))%64) | int64(base.Ui64(l0)>>(uint(v122)%64))
							v127 = int64(base.Ui64(v82) >> (uint(v122) % 64))
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v12))) = v126
					*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v127
					v131 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
					v134 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
					v137 = v131<<(uint(int64(4))%64) | int64(base.Ui64(v134)>>(uint(int64(60))%64))
					v139 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
					v140 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
					v148 = base.I64_extend_i32_u(base.B2i32(v20 != v72)&base.B2i32(v139|v140 != int64(0))) | v134&int64(1152921504606846975)
					if base.Ui64(int64(576460752303423489)) <= base.Ui64(v148) {
						v158 = v137 + int64(1)
					} else {
						if v148 != int64(576460752303423488) {
							v158 = v137
						} else {
							v158 = v137&int64(1) + v137
						}
					}
					v162 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v158))
					if base.Ui64(int64(4503599627370495)) < base.Ui64(v158) {
						v163 = v158 ^ int64(4503599627370496)
					} else {
						v163 = v158
					}
					v165 = v163
					v172 = base.I64_extend_i32_u(v162)
				}
			}
		} else {
			if v19 != int64(32767) {
				if base.Ui32(int32(17406)) < base.Ui32(v20) {
					v165 = int64(0)
					v172 = int64(2047)
				} else {
					v71 = base.B2i32(v19 == int64(0))
					if v19 == int64(0) {
						v72 = int32(15360)
					} else {
						v72 = int32(15361)
					}
					v73 = v72 - v20
					if int32(112) < v73 {
						v76 = int64(0)
						v165 = v76
						v172 = v76
					} else {
						v79 = v12 + int32(16)
						if v19 == int64(0) {
							v82 = v15
						} else {
							v82 = v15 | int64(281474976710656)
						}
						v84 = int32(128) - v73
						if v84&int32(64) != 0 {
							v103 = int64(0)
							v104 = l0 << (uint(base.I64_extend_i32_u(v84+int32(-64))) % 64)
						} else {
							if v84 == int32(0) {
								v103 = l0
								v104 = v82
							} else {
								v95 = base.I64_extend_i32_u(v84)
								v103 = l0 << (uint(v95) % 64)
								v104 = v82<<(uint(v95)%64) | int64(base.Ui64(l0)>>(uint(base.I64_extend_i32_u(int32(64)-v84))%64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v79))) = v103
						*(*int64)(unsafe.Add(mBase, uint32(v79)+8)) = v104
						if v73&int32(64) != 0 {
							v126 = int64(base.Ui64(v82) >> (uint(base.I64_extend_i32_u(v73+int32(-64))) % 64))
							v127 = int64(0)
						} else {
							if v73 == int32(0) {
								v126 = l0
								v127 = v82
							} else {
								v122 = base.I64_extend_i32_u(v73)
								v126 = v82<<(uint(base.I64_extend_i32_u(int32(64)-v73))%64) | int64(base.Ui64(l0)>>(uint(v122)%64))
								v127 = int64(base.Ui64(v82) >> (uint(v122) % 64))
							}
						}
						*(*int64)(unsafe.Add(mBase, uint32(v12))) = v126
						*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v127
						v131 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
						v134 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
						v137 = v131<<(uint(int64(4))%64) | int64(base.Ui64(v134)>>(uint(int64(60))%64))
						v139 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
						v140 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
						v148 = base.I64_extend_i32_u(base.B2i32(v20 != v72)&base.B2i32(v139|v140 != int64(0))) | v134&int64(1152921504606846975)
						if base.Ui64(int64(576460752303423489)) <= base.Ui64(v148) {
							v158 = v137 + int64(1)
						} else {
							if v148 != int64(576460752303423488) {
								v158 = v137
							} else {
								v158 = v137&int64(1) + v137
							}
						}
						v162 = base.B2i32(base.Ui64(int64(4503599627370495)) < base.Ui64(v158))
						if base.Ui64(int64(4503599627370495)) < base.Ui64(v158) {
							v163 = v158 ^ int64(4503599627370496)
						} else {
							v163 = v158
						}
						v165 = v163
						v172 = base.I64_extend_i32_u(v162)
					}
				}
			} else {
				v165 = v15<<(uint(int64(4))%64) | int64(base.Ui64(l0)>>(uint(int64(60))%64)) | int64(2251799813685248)
				v172 = int64(2047)
			}
		}
	}
	m.G0 = v12 + int32(32)
	return base.F64_reinterpret_i64(l1&int64(-9223372036854775807-1) | v172<<(uint(int64(52))%64) | v165)
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
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
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
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
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
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
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
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v563 int32
	_ = v563
	var v570 int32
	_ = v570
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v643 int32
	_ = v643
	var v650 int32
	_ = v650
	var v657 int32
	_ = v657
	var v664 int32
	_ = v664
	var v671 int32
	_ = v671
	var v678 int32
	_ = v678
	var v685 int32
	_ = v685
	var v692 int32
	_ = v692
	var v699 int32
	_ = v699
	var v705 int64
	_ = v705
	var v708 int64
	_ = v708
	var v711 int64
	_ = v711
	var v714 int64
	_ = v714
	var v717 int64
	_ = v717
	var v720 int64
	_ = v720
	var v723 int64
	_ = v723
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v747 int32
	_ = v747
	var v754 int32
	_ = v754
	var v761 int32
	_ = v761
	var v768 int32
	_ = v768
	var v775 int32
	_ = v775
	var v782 int32
	_ = v782
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v803 int32
	_ = v803
	var v809 int64
	_ = v809
	var v812 int64
	_ = v812
	var v815 int64
	_ = v815
	var v818 int64
	_ = v818
	var v821 int64
	_ = v821
	var v824 int64
	_ = v824
	var v827 int64
	_ = v827
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v840 int32
	_ = v840
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v946 int32
	_ = v946
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v965 int64
	_ = v965
	var v968 int32
	_ = v968
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1028 int32
	_ = v1028
	var v1035 int32
	_ = v1035
	var v1042 int32
	_ = v1042
	var v1049 int32
	_ = v1049
	var v1056 int32
	_ = v1056
	var v1063 int32
	_ = v1063
	var v1067 int32
	_ = v1067
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1081 int64
	_ = v1081
	var v1083 int64
	_ = v1083
	var v1086 int64
	_ = v1086
	var v1100 int32
	_ = v1100
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1115 int32
	_ = v1115
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1134 int32
	_ = v1134
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
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
	if l1&int32(3) == int32(0) {
		v48 = l1
		goto L7
	} else {
		goto L8
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
	if v1100 != 0 {
		goto L203
	} else {
		goto L204
	}
L5:
	;
	if base.Ui32(int32(99)) < base.Ui32(v81) {
		v1100 = int32(1)
		goto L4
	} else {
		goto L22
	}
L6:
	;
	v81 = v73 - l1
	goto L5
L7:
	;
	v52 = v48
	goto L16
L8:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v32 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v81 = int32(0)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v37 = l1
	goto L12
L12:
	;
	v41 = v37 + int32(1)
	if v41&int32(3) == int32(0) {
		v48 = v41
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v73 = v41
	goto L6
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v46 != 0 {
		v37 = v41
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v61 = int32(-2139062144)
	if (int32(16843008)-v58|v58)&v61 == v61 {
		v52 = v52 + int32(4)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v67 = v52
	goto L19
L18:
	;
	goto L17
L19:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v71 != 0 {
		v67 = v67 + int32(1)
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v73 = v67
	goto L6
L21:
	;
	goto L20
L22:
	;
	if l2 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v451 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+107)) = uint8(v451)
	v453 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+102)) = uint8(v453)
	v455 = int32(12336)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+100)) = uint16(v455)
	v457 = int32(7)
	v460 = v20&v457 | v453
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+106)) = uint8(v460)
	v467 = int32(base.Ui32(v20)>>(uint(int32(3))%32))&v457 | v453
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+105)) = uint8(v467)
	v474 = int32(base.Ui32(v20)>>(uint(int32(6))%32))&v457 | v453
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+104)) = uint8(v474)
	v481 = int32(base.Ui32(v20)>>(uint(int32(9))%32))&v457 | v453
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+103)) = uint8(v481)
	if base.Ui32(v21) <= base.Ui32(int32(2097151)) {
		goto L134
	} else {
		goto L135
	}
L24:
	;
	if l1&int32(3) == int32(0) {
		v410 = l1
		goto L115
	} else {
		goto L116
	}
L25:
	;
	if l2&int32(3) == int32(0) {
		v108 = l2
		goto L30
	} else {
		goto L31
	}
L26:
	;
	goto L27
L27:
	;
	v266 = F__emscripten_memset_bulkmem(m, v18, base.I32_extend8_s(int32(0)), int32(512))
	mBase = m.M
	goto L79
L28:
	;
	if base.Ui32(int32(99)) < base.Ui32(v141) {
		v1100 = int32(2)
		goto L4
	} else {
		goto L45
	}
L29:
	;
	v141 = v133 - l2
	goto L28
L30:
	;
	v112 = v108
	goto L39
L31:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v92 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v141 = int32(0)
	goto L28
L33:
	;
	goto L34
L34:
	;
	v97 = l2
	goto L35
L35:
	;
	v101 = v97 + int32(1)
	if v101&int32(3) == int32(0) {
		v108 = v101
		goto L30
	} else {
		goto L37
	}
L36:
	;
	v133 = v101
	goto L29
L37:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v106 != 0 {
		v97 = v101
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v121 = int32(-2139062144)
	if (int32(16843008)-v118|v118)&v121 == v121 {
		v112 = v112 + int32(4)
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v127 = v112
	goto L42
L41:
	;
	goto L40
L42:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v131 != 0 {
		v127 = v127 + int32(1)
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v133 = v127
	goto L29
L44:
	;
	goto L43
L45:
	;
	v147 = F__emscripten_memset_bulkmem(m, v18, base.I32_extend8_s(int32(0)), int32(512))
	mBase = m.M
	goto L46
L46:
	;
	goto L50
L47:
	;
	goto L24
L48:
	;
	v260 = F_strlen(m, v249)
	mBase = m.M
	goto L47
L50:
	;
	goto L51
L51:
	;
	v154 = int32(99)
	if (v147^l1)&int32(3) != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v253 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v250))) = uint8(v253)
	goto L48
L53:
	;
	v234 = v229
	v235 = v230
	v236 = v231
	goto L75
L54:
	;
	if v224 == int32(0) {
		v249 = v222
		v250 = v223
		goto L52
	} else {
		goto L74
	}
L55:
	;
	v222 = l1
	v223 = v147
	v224 = v154
	goto L54
L56:
	;
	goto L57
L57:
	;
	if l1&int32(3) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v191 == int32(0) {
		v249 = v188
		v250 = v189
		goto L52
	} else {
		goto L67
	}
L59:
	;
	v188 = l1
	v189 = v147
	v190 = v154
	v191 = int32(1)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v167 = l1
	v168 = v147
	v169 = v154
	goto L62
L62:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	*(*uint8)(unsafe.Add(mBase, uint32(v168))) = uint8(v171)
	if v171 == int32(0) {
		v229 = v167
		v230 = v168
		v231 = v169
		goto L53
	} else {
		goto L64
	}
L63:
	;
	v188 = v182
	v189 = v176
	v190 = v178
	v191 = v180
	goto L58
L64:
	;
	v175 = int32(1)
	v176 = v168 + v175
	v178 = v169 - v175
	v179 = int32(0)
	v180 = base.B2i32(v178 != v179)
	v182 = v167 + v175
	if v182&int32(3) == v179 {
		v188 = v182
		v189 = v176
		v190 = v178
		v191 = v180
		goto L58
	} else {
		goto L65
	}
L65:
	;
	if v178 != 0 {
		v167 = v182
		v168 = v176
		v169 = v178
		goto L62
	} else {
		goto L66
	}
L66:
	;
	goto L63
L67:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v194 == int32(0) {
		v222 = v188
		v223 = v189
		v224 = v190
		goto L54
	} else {
		goto L68
	}
L68:
	;
	if base.Ui32(v190) < base.Ui32(int32(4)) {
		v222 = v188
		v223 = v189
		v224 = v190
		goto L54
	} else {
		goto L69
	}
L69:
	;
	v200 = v188
	v201 = v189
	v202 = v190
	goto L70
L70:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	v208 = int32(-2139062144)
	if (int32(16843008)-v205|v205)&v208 != v208 {
		v229 = v200
		v230 = v201
		v231 = v202
		goto L53
	} else {
		goto L72
	}
L71:
	;
	v222 = v216
	v223 = v214
	v224 = v218
	goto L54
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v205
	v213 = int32(4)
	v214 = v201 + v213
	v216 = v200 + v213
	v218 = v202 - v213
	if base.Ui32(int32(3)) < base.Ui32(v218) {
		v200 = v216
		v201 = v214
		v202 = v218
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v229 = v222
	v230 = v223
	v231 = v224
	goto L53
L75:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234))))
	*(*uint8)(unsafe.Add(mBase, uint32(v235))) = uint8(v238)
	if v238 == int32(0) {
		v249 = v234
		v250 = v235
		goto L52
	} else {
		goto L77
	}
L76:
	;
	v249 = v245
	v250 = v243
	goto L52
L77:
	;
	v242 = int32(1)
	v243 = v235 + v242
	v245 = v234 + v242
	v247 = v236 - v242
	if v247 != 0 {
		v234 = v245
		v235 = v243
		v236 = v247
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	goto L83
L80:
	;
	if v20&int32(61440) != int32(16384) {
		goto L23
	} else {
		goto L112
	}
L81:
	;
	v379 = F_strlen(m, v368)
	mBase = m.M
	goto L80
L83:
	;
	goto L84
L84:
	;
	v273 = int32(99)
	if (v266^l1)&int32(3) != 0 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v372 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v369))) = uint8(v372)
	goto L81
L86:
	;
	v353 = v348
	v354 = v349
	v355 = v350
	goto L108
L87:
	;
	if v343 == int32(0) {
		v368 = v341
		v369 = v342
		goto L85
	} else {
		goto L107
	}
L88:
	;
	v341 = l1
	v342 = v266
	v343 = v273
	goto L87
L89:
	;
	goto L90
L90:
	;
	if l1&int32(3) == int32(0) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v310 == int32(0) {
		v368 = v307
		v369 = v308
		goto L85
	} else {
		goto L100
	}
L92:
	;
	v307 = l1
	v308 = v266
	v309 = v273
	v310 = int32(1)
	goto L91
L93:
	;
	goto L94
L94:
	;
	v286 = l1
	v287 = v266
	v288 = v273
	goto L95
L95:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	*(*uint8)(unsafe.Add(mBase, uint32(v287))) = uint8(v290)
	if v290 == int32(0) {
		v348 = v286
		v349 = v287
		v350 = v288
		goto L86
	} else {
		goto L97
	}
L96:
	;
	v307 = v301
	v308 = v295
	v309 = v297
	v310 = v299
	goto L91
L97:
	;
	v294 = int32(1)
	v295 = v287 + v294
	v297 = v288 - v294
	v298 = int32(0)
	v299 = base.B2i32(v297 != v298)
	v301 = v286 + v294
	if v301&int32(3) == v298 {
		v307 = v301
		v308 = v295
		v309 = v297
		v310 = v299
		goto L91
	} else {
		goto L98
	}
L98:
	;
	if v297 != 0 {
		v286 = v301
		v287 = v295
		v288 = v297
		goto L95
	} else {
		goto L99
	}
L99:
	;
	goto L96
L100:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	if v313 == int32(0) {
		v341 = v307
		v342 = v308
		v343 = v309
		goto L87
	} else {
		goto L101
	}
L101:
	;
	if base.Ui32(v309) < base.Ui32(int32(4)) {
		v341 = v307
		v342 = v308
		v343 = v309
		goto L87
	} else {
		goto L102
	}
L102:
	;
	v319 = v307
	v320 = v308
	v321 = v309
	goto L103
L103:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	v327 = int32(-2139062144)
	if (int32(16843008)-v324|v324)&v327 != v327 {
		v348 = v319
		v349 = v320
		v350 = v321
		goto L86
	} else {
		goto L105
	}
L104:
	;
	v341 = v335
	v342 = v333
	v343 = v337
	goto L87
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320))) = v324
	v332 = int32(4)
	v333 = v320 + v332
	v335 = v319 + v332
	v337 = v321 - v332
	if base.Ui32(int32(3)) < base.Ui32(v337) {
		v319 = v335
		v320 = v333
		v321 = v337
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	v348 = v341
	v349 = v342
	v350 = v343
	goto L86
L108:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353))))
	*(*uint8)(unsafe.Add(mBase, uint32(v354))) = uint8(v357)
	if v357 == int32(0) {
		v368 = v353
		v369 = v354
		goto L85
	} else {
		goto L110
	}
L109:
	;
	v368 = v364
	v369 = v362
	goto L85
L110:
	;
	v361 = int32(1)
	v362 = v354 + v361
	v364 = v353 + v361
	v366 = v355 - v361
	if v366 != 0 {
		v353 = v364
		v354 = v362
		v355 = v366
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	goto L24
L113:
	;
	if int32(99) <= v443 {
		goto L130
	} else {
		goto L131
	}
L114:
	;
	v443 = v435 - l1
	goto L113
L115:
	;
	v414 = v410
	goto L124
L116:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v394 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v443 = int32(0)
	goto L113
L118:
	;
	goto L119
L119:
	;
	v399 = l1
	goto L120
L120:
	;
	v403 = v399 + int32(1)
	if v403&int32(3) == int32(0) {
		v410 = v403
		goto L115
	} else {
		goto L122
	}
L121:
	;
	v435 = v403
	goto L114
L122:
	;
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403))))
	if v408 != 0 {
		v399 = v403
		goto L120
	} else {
		goto L123
	}
L123:
	;
	goto L121
L124:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v414)))
	v423 = int32(-2139062144)
	if (int32(16843008)-v420|v420)&v423 == v423 {
		v414 = v414 + int32(4)
		goto L124
	} else {
		goto L126
	}
L125:
	;
	v429 = v414
	goto L127
L126:
	;
	goto L125
L127:
	;
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v429))))
	if v433 != 0 {
		v429 = v429 + int32(1)
		goto L127
	} else {
		goto L129
	}
L128:
	;
	v435 = v429
	goto L114
L129:
	;
	goto L128
L130:
	;
	v446 = int32(99)
	goto L132
L131:
	;
	v446 = v443
	goto L132
L132:
	;
	v448 = int32(47)
	*(*uint16)(unsafe.Add(mBase, uint32(v18+v446))) = uint16(v448)
	goto L23
L133:
	;
	if base.Ui32(v22) <= base.Ui32(int32(2097151)) {
		goto L138
	} else {
		goto L139
	}
L134:
	;
	v485 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+115)) = uint8(v485)
	v487 = int32(7)
	v489 = int32(48)
	v490 = v21&v487 | v489
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+114)) = uint8(v490)
	v495 = int32(base.Ui32(v21)>>(uint(int32(18))%32)) | v489
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+108)) = uint8(v495)
	v502 = int32(base.Ui32(v21)>>(uint(int32(3))%32))&v487 | v489
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+113)) = uint8(v502)
	v509 = int32(base.Ui32(v21)>>(uint(int32(6))%32))&v487 | v489
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+112)) = uint8(v509)
	v516 = int32(base.Ui32(v21)>>(uint(int32(9))%32))&v487 | v489
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+111)) = uint8(v516)
	v523 = int32(base.Ui32(v21)>>(uint(int32(12))%32))&v487 | v489
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+110)) = uint8(v523)
	v530 = int32(base.Ui32(v21)>>(uint(int32(15))%32))&v487 | v489
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+109)) = uint8(v530)
	goto L133
L135:
	;
	goto L136
L136:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+115)) = uint8(v21)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = int32(128)
	v536 = int32(base.Ui32(v21) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+114)) = uint8(v536)
	v539 = int32(base.Ui32(v21) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+113)) = uint8(v539)
	v542 = int32(base.Ui32(v21) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+112)) = uint8(v542)
	goto L133
L137:
	;
	v605 = int32(0)
	v608 = v20 & int32(61440)
	if base.B2i32(l2 == v605)&base.B2i32(v608 != int32(16384)) == v605 {
		goto L142
	} else {
		goto L143
	}
L138:
	;
	v546 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+123)) = uint8(v546)
	v548 = int32(7)
	v550 = int32(48)
	v551 = v22&v548 | v550
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+122)) = uint8(v551)
	v556 = int32(base.Ui32(v22)>>(uint(int32(18))%32)) | v550
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+116)) = uint8(v556)
	v563 = int32(base.Ui32(v22)>>(uint(int32(3))%32))&v548 | v550
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+121)) = uint8(v563)
	v570 = int32(base.Ui32(v22)>>(uint(int32(6))%32))&v548 | v550
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+120)) = uint8(v570)
	v577 = int32(base.Ui32(v22)>>(uint(int32(9))%32))&v548 | v550
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+119)) = uint8(v577)
	v584 = int32(base.Ui32(v22)>>(uint(int32(12))%32))&v548 | v550
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+118)) = uint8(v584)
	v591 = int32(base.Ui32(v22)>>(uint(int32(15))%32))&v548 | v550
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+117)) = uint8(v591)
	goto L137
L139:
	;
	goto L140
L140:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+123)) = uint8(v22)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = int32(128)
	v597 = int32(base.Ui32(v22) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+122)) = uint8(v597)
	v600 = int32(base.Ui32(v22) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+121)) = uint8(v600)
	v603 = int32(base.Ui32(v22) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+120)) = uint8(v603)
	goto L137
L141:
	;
	if base.Ui64(v23) <= base.Ui64(int64(8589934591)) {
		goto L149
	} else {
		goto L150
	}
L142:
	;
	v614 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+135)) = uint8(v614)
	v617 = v18 + int32(124)
	*(*int32)(unsafe.Add(mBase, uint32(v617)+7)) = int32(808464432)
	*(*int64)(unsafe.Add(mBase, uint32(v617))) = int64(3472328296227680304)
	goto L141
L143:
	;
	goto L144
L144:
	;
	if base.Ui64(v19) <= base.Ui64(int64(8589934591)) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v624 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+135)) = uint8(v624)
	v626 = base.I32_wrap_i64(v19)
	v627 = int32(7)
	v629 = int32(48)
	v630 = v626&v627 | v629
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+134)) = uint8(v630)
	v636 = base.I32_wrap_i64(int64(base.Ui64(v19)>>(uint(int64(30))%64))) | v629
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+124)) = uint8(v636)
	v643 = int32(base.Ui32(v626)>>(uint(int32(3))%32))&v627 | v629
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+133)) = uint8(v643)
	v650 = int32(base.Ui32(v626)>>(uint(int32(6))%32))&v627 | v629
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+132)) = uint8(v650)
	v657 = int32(base.Ui32(v626)>>(uint(int32(9))%32))&v627 | v629
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+131)) = uint8(v657)
	v664 = int32(base.Ui32(v626)>>(uint(int32(12))%32))&v627 | v629
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+130)) = uint8(v664)
	v671 = int32(base.Ui32(v626)>>(uint(int32(15))%32))&v627 | v629
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+129)) = uint8(v671)
	v678 = int32(base.Ui32(v626)>>(uint(int32(18))%32))&v627 | v629
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+128)) = uint8(v678)
	v685 = int32(base.Ui32(v626)>>(uint(int32(21))%32))&v627 | v629
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+127)) = uint8(v685)
	v692 = int32(base.Ui32(v626)>>(uint(int32(24))%32))&v627 | v629
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+126)) = uint8(v692)
	v699 = int32(base.Ui32(v626)>>(uint(int32(27))%32))&v627 | v629
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+125)) = uint8(v699)
	goto L141
L146:
	;
	goto L147
L147:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+135)) = uint8(v19)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = int32(128)
	v705 = int64(base.Ui64(v19) >> (uint(int64(8)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+134)) = uint8(v705)
	v708 = int64(base.Ui64(v19) >> (uint(int64(16)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+133)) = uint8(v708)
	v711 = int64(base.Ui64(v19) >> (uint(int64(24)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+132)) = uint8(v711)
	v714 = int64(base.Ui64(v19) >> (uint(int64(32)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+131)) = uint8(v714)
	v717 = int64(base.Ui64(v19) >> (uint(int64(40)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+130)) = uint8(v717)
	v720 = int64(base.Ui64(v19) >> (uint(int64(48)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+129)) = uint8(v720)
	v723 = int64(base.Ui64(v19) >> (uint(int64(56)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+128)) = uint8(v723)
	goto L141
L148:
	;
	if l2 != 0 {
		goto L153
	} else {
		goto L154
	}
L149:
	;
	v728 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+147)) = uint8(v728)
	v730 = base.I32_wrap_i64(v23)
	v731 = int32(7)
	v733 = int32(48)
	v734 = v730&v731 | v733
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+146)) = uint8(v734)
	v740 = base.I32_wrap_i64(int64(base.Ui64(v23)>>(uint(int64(30))%64))) | v733
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+136)) = uint8(v740)
	v747 = int32(base.Ui32(v730)>>(uint(int32(3))%32))&v731 | v733
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+145)) = uint8(v747)
	v754 = int32(base.Ui32(v730)>>(uint(int32(6))%32))&v731 | v733
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+144)) = uint8(v754)
	v761 = int32(base.Ui32(v730)>>(uint(int32(9))%32))&v731 | v733
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+143)) = uint8(v761)
	v768 = int32(base.Ui32(v730)>>(uint(int32(12))%32))&v731 | v733
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+142)) = uint8(v768)
	v775 = int32(base.Ui32(v730)>>(uint(int32(15))%32))&v731 | v733
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+141)) = uint8(v775)
	v782 = int32(base.Ui32(v730)>>(uint(int32(18))%32))&v731 | v733
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+140)) = uint8(v782)
	v789 = int32(base.Ui32(v730)>>(uint(int32(21))%32))&v731 | v733
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+139)) = uint8(v789)
	v796 = int32(base.Ui32(v730)>>(uint(int32(24))%32))&v731 | v733
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+138)) = uint8(v796)
	v803 = int32(base.Ui32(v730)>>(uint(int32(27))%32))&v731 | v733
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+137)) = uint8(v803)
	goto L148
L150:
	;
	goto L151
L151:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+147)) = uint8(v23)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+136)) = int32(128)
	v809 = int64(base.Ui64(v23) >> (uint(int64(8)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+146)) = uint8(v809)
	v812 = int64(base.Ui64(v23) >> (uint(int64(16)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+145)) = uint8(v812)
	v815 = int64(base.Ui64(v23) >> (uint(int64(24)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+144)) = uint8(v815)
	v818 = int64(base.Ui64(v23) >> (uint(int64(32)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+143)) = uint8(v818)
	v821 = int64(base.Ui64(v23) >> (uint(int64(40)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+142)) = uint8(v821)
	v824 = int64(base.Ui64(v23) >> (uint(int64(48)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+141)) = uint8(v824)
	v827 = int64(base.Ui64(v23) >> (uint(int64(56)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+140)) = uint8(v827)
	goto L148
L152:
	;
	v957 = int32(*(*uint16)(unsafe.Add(mBase, _consts[191])))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+261)) = uint16(v957)
	v960 = *(*int32)(unsafe.Add(mBase, _consts[192]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+257)) = v960
	v962 = int32(12336)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+263)) = uint16(v962)
	v965 = *(*int64)(unsafe.Add(mBase, _consts[97]))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+265)) = v965
	v968 = int32(*(*uint8)(unsafe.Add(mBase, _consts[193])))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+273)) = uint8(v968)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+297)) = v965
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+305)) = uint8(v968)
	v972 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+336)) = uint8(v972)
	v974 = int32(808464432)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+329)) = v974
	*(*int32)(unsafe.Add(mBase, uint32(v18)+332)) = v974
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+344)) = uint8(v972)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+340)) = v974
	*(*int32)(unsafe.Add(mBase, uint32(v18)+337)) = v974
	v990 = int32(0)
	v991 = int32(256)
	goto L191
L153:
	;
	v830 = int32(50)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+156)) = uint8(v830)
	v833 = v18 + int32(157)
	goto L159
L154:
	;
	goto L155
L155:
	;
	if v608 == int32(16384) {
		goto L188
	} else {
		goto L189
	}
L156:
	;
	goto L152
L157:
	;
	v946 = F_strlen(m, v935)
	mBase = m.M
	goto L156
L159:
	;
	goto L160
L160:
	;
	v840 = int32(99)
	if (v833^l2)&int32(3) != 0 {
		goto L164
	} else {
		goto L165
	}
L161:
	;
	v939 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v936))) = uint8(v939)
	goto L157
L162:
	;
	v920 = v915
	v921 = v916
	v922 = v917
	goto L184
L163:
	;
	if v910 == int32(0) {
		v935 = v908
		v936 = v909
		goto L161
	} else {
		goto L183
	}
L164:
	;
	v908 = l2
	v909 = v833
	v910 = v840
	goto L163
L165:
	;
	goto L166
L166:
	;
	if l2&int32(3) == int32(0) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	if v877 == int32(0) {
		v935 = v874
		v936 = v875
		goto L161
	} else {
		goto L176
	}
L168:
	;
	v874 = l2
	v875 = v833
	v876 = v840
	v877 = int32(1)
	goto L167
L169:
	;
	goto L170
L170:
	;
	v853 = l2
	v854 = v833
	v855 = v840
	goto L171
L171:
	;
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v853))))
	*(*uint8)(unsafe.Add(mBase, uint32(v854))) = uint8(v857)
	if v857 == int32(0) {
		v915 = v853
		v916 = v854
		v917 = v855
		goto L162
	} else {
		goto L173
	}
L172:
	;
	v874 = v868
	v875 = v862
	v876 = v864
	v877 = v866
	goto L167
L173:
	;
	v861 = int32(1)
	v862 = v854 + v861
	v864 = v855 - v861
	v865 = int32(0)
	v866 = base.B2i32(v864 != v865)
	v868 = v853 + v861
	if v868&int32(3) == v865 {
		v874 = v868
		v875 = v862
		v876 = v864
		v877 = v866
		goto L167
	} else {
		goto L174
	}
L174:
	;
	if v864 != 0 {
		v853 = v868
		v854 = v862
		v855 = v864
		goto L171
	} else {
		goto L175
	}
L175:
	;
	goto L172
L176:
	;
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v874))))
	if v880 == int32(0) {
		v908 = v874
		v909 = v875
		v910 = v876
		goto L163
	} else {
		goto L177
	}
L177:
	;
	if base.Ui32(v876) < base.Ui32(int32(4)) {
		v908 = v874
		v909 = v875
		v910 = v876
		goto L163
	} else {
		goto L178
	}
L178:
	;
	v886 = v874
	v887 = v875
	v888 = v876
	goto L179
L179:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v886)))
	v894 = int32(-2139062144)
	if (int32(16843008)-v891|v891)&v894 != v894 {
		v915 = v886
		v916 = v887
		v917 = v888
		goto L162
	} else {
		goto L181
	}
L180:
	;
	v908 = v902
	v909 = v900
	v910 = v904
	goto L163
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v887))) = v891
	v899 = int32(4)
	v900 = v887 + v899
	v902 = v886 + v899
	v904 = v888 - v899
	if base.Ui32(int32(3)) < base.Ui32(v904) {
		v886 = v902
		v887 = v900
		v888 = v904
		goto L179
	} else {
		goto L182
	}
L182:
	;
	goto L180
L183:
	;
	v915 = v908
	v916 = v909
	v917 = v910
	goto L162
L184:
	;
	v924 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920))))
	*(*uint8)(unsafe.Add(mBase, uint32(v921))) = uint8(v924)
	if v924 == int32(0) {
		v935 = v920
		v936 = v921
		goto L161
	} else {
		goto L186
	}
L185:
	;
	v935 = v931
	v936 = v929
	goto L161
L186:
	;
	v928 = int32(1)
	v929 = v921 + v928
	v931 = v920 + v928
	v933 = v922 - v928
	if v933 != 0 {
		v920 = v931
		v921 = v929
		v922 = v933
		goto L184
	} else {
		goto L187
	}
L187:
	;
	goto L185
L188:
	;
	v951 = int32(53)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+156)) = uint8(v951)
	goto L152
L189:
	;
	goto L190
L190:
	;
	v953 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+156)) = uint8(v953)
	goto L152
L191:
	;
	if base.Ui32(v990-int32(156)) <= base.Ui32(int32(-9)) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	if base.Ui32(v1011) <= base.Ui32(int32(2097151)) {
		goto L200
	} else {
		goto L201
	}
L193:
	;
	v1001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v990))))
	v1003 = v991 + v1001
	goto L195
L194:
	;
	v1003 = v991
	goto L195
L195:
	;
	if base.Ui32(v990-int32(155)) <= base.Ui32(int32(-9)) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v1009 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v990)+1)))
	v1011 = v1003 + v1009
	goto L198
L197:
	;
	v1011 = v1003
	goto L198
L198:
	;
	v1013 = v990 + int32(2)
	if v1013 != int32(512) {
		v990 = v1013
		v991 = v1011
		goto L191
	} else {
		goto L199
	}
L199:
	;
	goto L192
L200:
	;
	v1018 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+155)) = uint8(v1018)
	v1020 = int32(7)
	v1022 = int32(48)
	v1023 = v1011&v1020 | v1022
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+154)) = uint8(v1023)
	v1028 = int32(base.Ui32(v1011)>>(uint(int32(18))%32)) | v1022
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+148)) = uint8(v1028)
	v1035 = int32(base.Ui32(v1011)>>(uint(int32(3))%32))&v1020 | v1022
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+153)) = uint8(v1035)
	v1042 = int32(base.Ui32(v1011)>>(uint(int32(6))%32))&v1020 | v1022
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+152)) = uint8(v1042)
	v1049 = int32(base.Ui32(v1011)>>(uint(int32(9))%32))&v1020 | v1022
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+151)) = uint8(v1049)
	v1056 = int32(base.Ui32(v1011)>>(uint(int32(12))%32))&v1020 | v1022
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+150)) = uint8(v1056)
	v1063 = int32(base.Ui32(v1011)>>(uint(int32(15))%32))&v1020 | v1022
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+149)) = uint8(v1063)
	v1100 = int32(0)
	goto L4
L201:
	;
	goto L202
L202:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+155)) = uint8(v1011)
	v1067 = int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+148)) = uint8(v1067)
	v1070 = int32(base.Ui32(v1011) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+154)) = uint8(v1070)
	v1073 = int32(base.Ui32(v1011) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+153)) = uint8(v1073)
	v1076 = int32(base.Ui32(v1011) >> (uint(int32(24)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+152)) = uint8(v1076)
	v1079 = v1011 >> (uint(int32(31)) % 32)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+151)) = uint8(v1079)
	v1081 = base.I64_extend_i32_s(v1011)
	v1083 = int64(base.Ui64(v1081) >> (uint(int64(40)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+150)) = uint8(v1083)
	v1086 = int64(base.Ui64(v1081) >> (uint(int64(48)) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+149)) = uint8(v1086)
	v1100 = int32(0)
	goto L4
L203:
	;
	switch v1100 - int32(1) {
	case 0:
		goto L208
	case 1:
		goto L207
	default:
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1154)+8))
	m.T0[v1155].(func(*base.Module, int32, int32))(m, l0, int32(512))
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L209
	} else {
		goto L221
	}
L206:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L209
	} else {
		goto L218
	}
L207:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L209
	} else {
		goto L214
	}
L208:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	return
L210:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L209
	} else {
		goto L211
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l1
	F_errmsg(m, int32(683064), v14+int32(16))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L209
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(472560), int32(2053), int32(216852))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L209
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L209
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l1
	F_errmsg(m, int32(657678), v14+int32(32))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L209
	} else {
		goto L216
	}
L216:
	;
	F_errfinish(m, int32(472560), int32(2060), int32(216852))
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L209
	} else {
		goto L217
	}
L217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v1100
	F_errmsg_internal(m, int32(459637), v14)
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L209
	} else {
		goto L219
	}
L219:
	;
	F_errfinish(m, int32(472560), int32(2063), int32(216852))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L209
	} else {
		goto L220
	}
L220:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L221:
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v632 int32
	_ = v632
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v878 int32
	_ = v878
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v924 int32
	_ = v924
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1058 int32
	_ = v1058
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1075 int32
	_ = v1075
	var v1079 int32
	_ = v1079
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1088 int32
	_ = v1088
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1149 int32
	_ = v1149
	var v1154 int32
	_ = v1154
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1232 int32
	_ = v1232
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1255 int32
	_ = v1255
	var v1261 int32
	_ = v1261
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1284 int32
	_ = v1284
	var v1289 int32
	_ = v1289
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1310 int32
	_ = v1310
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1356 int32
	_ = v1356
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1465 int32
	_ = v1465
	var v1468 int32
	_ = v1468
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1487 int32
	_ = v1487
	var v1491 int32
	_ = v1491
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1506 int32
	_ = v1506
	var v1510 int32
	_ = v1510
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1528 int32
	_ = v1528
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1537 int32
	_ = v1537
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1575 int32
	_ = v1575
	var v1585 int32
	_ = v1585
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
	return v1585
L2:
	;
	return int32(0)
L3:
	;
	if v12 < int32(0) {
		v1585 = v12
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
	if v100 < int32(5) {
		v1585 = v19
		goto L1
	} else {
		goto L22
	}
L6:
	;
	v100 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v33 = v28 & int32(3)
	if base.Ui32(v28) < base.Ui32(int32(4)) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v33 != 0 {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	v67 = v20
	v68 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v40 = v20
	v41 = int32(0)
	v44 = v19
	goto L13
L13:
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
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v67 = v63
	v68 = v61
	goto L9
L15:
	;
	goto L14
L16:
	;
	v73 = v67
	v74 = v68
	v76 = v19
	goto L19
L17:
	;
	v89 = v68
	goto L18
L18:
	;
	v100 = v89
	goto L5
L19:
	;
	v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73))))
	v82 = v74 + base.B2i32(int32(-65) < v79)
	v83 = int32(1)
	v86 = v76 + v83
	if v86 != v33 {
		v73 = v73 + v83
		v74 = v82
		v76 = v86
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v89 = v82
	goto L18
L21:
	;
	goto L20
L22:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v103
	v105 = int32(3)
	v107 = int32(0)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v109-v103 < v105 {
		v119 = v107
		goto L25
	} else {
		goto L26
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103
	v159 = v103 + int32(2)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v160 <= v159 {
		goto L40
	} else {
		goto L41
	}
L24:
	;
	if v119 == int32(0) {
		goto L23
	} else {
		goto L28
	}
L25:
	;
	goto L24
L26:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v115 = F_memcmp(m, v113+v103, int32(2172190), v105)
	mBase = m.M
	if v115 != 0 {
		v119 = v107
		goto L25
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v105 + v103
	v119 = int32(1)
	goto L25
L28:
	;
	v124 = F_find_among(m, l0, int32(4318720), int32(10))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	if v124 == int32(0) {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v128 = int32(3)
	v130 = int32(0)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v132-v133 < v128 {
		v142 = v130
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v142 == int32(0) {
		goto L23
	} else {
		goto L35
	}
L32:
	;
	goto L31
L33:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v138 = F_memcmp(m, v136+v133, int32(2172193), v128)
	mBase = m.M
	if v138 != 0 {
		v142 = v130
		goto L32
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v128 + v133
	v142 = int32(1)
	goto L32
L35:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v145
	v147 = F_slice_del(m, l0)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	if v147 < int32(0) {
		v1585 = v147
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v151 = F_r_fix_va_start(m, l0)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	if v151 < int32(0) {
		v1585 = v151
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L23
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v217 = int32(0)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v216-int32(4))))
	if v224 == v217 {
		goto L59
	} else {
		goto L60
	}
L41:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+v159))))
	if v164&int32(224) != int32(128) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	if int32(1)<<(uint(v164)%32)&int32(672) == int32(0) {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v177 = F_find_among(m, l0, int32(4319008), int32(3))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	if v177 == int32(0) {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	v183 = F_find_among(m, l0, int32(4319072), int32(10))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	if v183 == int32(0) {
		goto L40
	} else {
		goto L47
	}
L47:
	;
	v187 = int32(3)
	v189 = int32(0)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v191-v192 < v187 {
		v201 = v189
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v201 == int32(0) {
		goto L40
	} else {
		goto L52
	}
L49:
	;
	goto L48
L50:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v197 = F_memcmp(m, v195+v192, int32(2172262), v187)
	mBase = m.M
	if v197 != 0 {
		v201 = v189
		goto L49
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v187 + v192
	v201 = int32(1)
	goto L49
L52:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v204
	v206 = F_slice_del(m, l0)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	if v206 < int32(0) {
		v1585 = v206
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v210 = F_r_fix_va_start(m, l0)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	if v210 < int32(0) {
		v1585 = v210
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L40
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v353 = int32(0)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v352-int32(4))))
	if v360 == v353 {
		goto L101
	} else {
		goto L102
	}
L58:
	;
	if v296 < int32(5) {
		goto L57
	} else {
		goto L75
	}
L59:
	;
	v296 = int32(0)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v229 = v224 & int32(3)
	if base.Ui32(v224) < base.Ui32(int32(4)) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if v229 != 0 {
		goto L69
	} else {
		goto L70
	}
L63:
	;
	v263 = v216
	v264 = int32(0)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v236 = v216
	v237 = int32(0)
	v240 = v217
	goto L66
L66:
	;
	v242 = int32(*(*int8)(unsafe.Add(mBase, uint32(v236))))
	v243 = int32(-65)
	v246 = int32(*(*int8)(unsafe.Add(mBase, uint32(v236)+1)))
	v250 = int32(*(*int8)(unsafe.Add(mBase, uint32(v236)+2)))
	v254 = int32(*(*int8)(unsafe.Add(mBase, uint32(v236)+3)))
	v257 = v237 + base.B2i32(v243 < v242) + base.B2i32(v243 < v246) + base.B2i32(v243 < v250) + base.B2i32(v243 < v254)
	v258 = int32(4)
	v259 = v236 + v258
	v261 = v240 + v258
	if v261 != v224&int32(-4) {
		v236 = v259
		v237 = v257
		v240 = v261
		goto L66
	} else {
		goto L68
	}
L67:
	;
	v263 = v259
	v264 = v257
	goto L62
L68:
	;
	goto L67
L69:
	;
	v269 = v263
	v270 = v264
	v272 = v217
	goto L72
L70:
	;
	v285 = v264
	goto L71
L71:
	;
	v296 = v285
	goto L58
L72:
	;
	v275 = int32(*(*int8)(unsafe.Add(mBase, uint32(v269))))
	v278 = v270 + base.B2i32(int32(-65) < v275)
	v279 = int32(1)
	v282 = v272 + v279
	if v282 != v229 {
		v269 = v269 + v279
		v270 = v278
		v272 = v282
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v285 = v278
	goto L71
L74:
	;
	goto L73
L75:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v299
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v301
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v301
	v306 = F_find_among_b(m, l0, int32(4319280), int32(3))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	if v306 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v308
	v312 = F_slice_from_s(m, l0, int32(3), int32(2172304))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L2
	} else {
		goto L80
	}
L78:
	;
	v316 = v301
	goto L79
L79:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v317
	v320 = v316
	goto L82
L80:
	;
	if v312 < int32(0) {
		v1585 = v312
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v316 = v312
	goto L79
L82:
	;
	v326 = F_r_fix_ending(m, l0)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L2
	} else {
		goto L86
	}
L83:
	;
	if v333 == int32(0) {
		goto L57
	} else {
		goto L97
	}
L84:
	;
	if v326 < int32(0) {
		goto L90
	} else {
		goto L91
	}
L85:
	;
	v333 = int32(2)
	goto L84
L86:
	;
	v329 = int32(base.Ui32(v326) >> (uint(int32(31)) % 32))
	if v326 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v331 = v329
	goto L89
L88:
	;
	v331 = int32(4)
	goto L89
L89:
	;
	switch v331 {
	case 0:
		goto L85
	default:
		v333 = v329
		goto L84
	case 4:
		goto L57
	}
L90:
	;
	v336 = v326
	goto L92
L91:
	;
	v336 = v320
	goto L92
L92:
	;
	if v326 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v337 = v336
	goto L95
L94:
	;
	v337 = v320
	goto L95
L95:
	;
	if v333 == int32(2) {
		v320 = v337
		goto L82
	} else {
		goto L96
	}
L96:
	;
	goto L83
L97:
	;
	if v337 < int32(0) {
		v1585 = v337
		goto L1
	} else {
		goto L98
	}
L98:
	;
	goto L57
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103
	v474 = int32(0)
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v475-int32(4))))
	if v483 == v474 {
		goto L129
	} else {
		goto L130
	}
L100:
	;
	if v432 < int32(5) {
		goto L99
	} else {
		goto L117
	}
L101:
	;
	v432 = int32(0)
	goto L100
L102:
	;
	goto L103
L103:
	;
	v365 = v360 & int32(3)
	if base.Ui32(v360) < base.Ui32(int32(4)) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	if v365 != 0 {
		goto L111
	} else {
		goto L112
	}
L105:
	;
	v399 = v352
	v400 = int32(0)
	goto L104
L106:
	;
	goto L107
L107:
	;
	v372 = v352
	v373 = int32(0)
	v376 = v353
	goto L108
L108:
	;
	v378 = int32(*(*int8)(unsafe.Add(mBase, uint32(v372))))
	v379 = int32(-65)
	v382 = int32(*(*int8)(unsafe.Add(mBase, uint32(v372)+1)))
	v386 = int32(*(*int8)(unsafe.Add(mBase, uint32(v372)+2)))
	v390 = int32(*(*int8)(unsafe.Add(mBase, uint32(v372)+3)))
	v393 = v373 + base.B2i32(v379 < v378) + base.B2i32(v379 < v382) + base.B2i32(v379 < v386) + base.B2i32(v379 < v390)
	v394 = int32(4)
	v395 = v372 + v394
	v397 = v376 + v394
	if v397 != v360&int32(-4) {
		v372 = v395
		v373 = v393
		v376 = v397
		goto L108
	} else {
		goto L110
	}
L109:
	;
	v399 = v395
	v400 = v393
	goto L104
L110:
	;
	goto L109
L111:
	;
	v405 = v399
	v406 = v400
	v408 = v353
	goto L114
L112:
	;
	v421 = v400
	goto L113
L113:
	;
	v432 = v421
	goto L100
L114:
	;
	v411 = int32(*(*int8)(unsafe.Add(mBase, uint32(v405))))
	v414 = v406 + base.B2i32(int32(-65) < v411)
	v415 = int32(1)
	v418 = v408 + v415
	if v418 != v365 {
		v405 = v405 + v415
		v406 = v414
		v408 = v418
		goto L114
	} else {
		goto L116
	}
L115:
	;
	v421 = v414
	goto L113
L116:
	;
	goto L115
L117:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v435
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v437
	v440 = int32(9)
	v442 = int32(0)
	if v437-v435 < v440 {
		v455 = v442
		goto L119
	} else {
		goto L120
	}
L118:
	;
	if v455 == int32(0) {
		goto L99
	} else {
		goto L122
	}
L119:
	;
	goto L118
L120:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v451 = F_memcmp(m, v448+v437-v440, int32(2172316), v440)
	mBase = m.M
	if v451 != 0 {
		v455 = v442
		goto L119
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v437 - v440
	v455 = int32(1)
	goto L119
L122:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v458
	v462 = F_slice_from_s(m, l0, int32(3), int32(2172325))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L2
	} else {
		goto L123
	}
L123:
	;
	if v462 < int32(0) {
		v1585 = v462
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v466
	v468 = F_r_fix_ending(m, l0)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L2
	} else {
		goto L125
	}
L125:
	;
	if v468 < int32(0) {
		v1585 = v468
		goto L1
	} else {
		goto L126
	}
L126:
	;
	goto L99
L127:
	;
	if v632 < int32(0) {
		v1585 = v632
		goto L1
	} else {
		goto L179
	}
L128:
	;
	if v555 < int32(5) {
		v632 = v474
		goto L127
	} else {
		goto L145
	}
L129:
	;
	v555 = int32(0)
	goto L128
L130:
	;
	goto L131
L131:
	;
	v488 = v483 & int32(3)
	if base.Ui32(v483) < base.Ui32(int32(4)) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	if v488 != 0 {
		goto L139
	} else {
		goto L140
	}
L133:
	;
	v522 = v475
	v523 = int32(0)
	goto L132
L134:
	;
	goto L135
L135:
	;
	v495 = v475
	v496 = int32(0)
	v499 = v474
	goto L136
L136:
	;
	v501 = int32(*(*int8)(unsafe.Add(mBase, uint32(v495))))
	v502 = int32(-65)
	v505 = int32(*(*int8)(unsafe.Add(mBase, uint32(v495)+1)))
	v509 = int32(*(*int8)(unsafe.Add(mBase, uint32(v495)+2)))
	v513 = int32(*(*int8)(unsafe.Add(mBase, uint32(v495)+3)))
	v516 = v496 + base.B2i32(v502 < v501) + base.B2i32(v502 < v505) + base.B2i32(v502 < v509) + base.B2i32(v502 < v513)
	v517 = int32(4)
	v518 = v495 + v517
	v520 = v499 + v517
	if v520 != v483&int32(-4) {
		v495 = v518
		v496 = v516
		v499 = v520
		goto L136
	} else {
		goto L138
	}
L137:
	;
	v522 = v518
	v523 = v516
	goto L132
L138:
	;
	goto L137
L139:
	;
	v528 = v522
	v529 = v523
	v531 = v474
	goto L142
L140:
	;
	v544 = v523
	goto L141
L141:
	;
	v555 = v544
	goto L128
L142:
	;
	v534 = int32(*(*int8)(unsafe.Add(mBase, uint32(v528))))
	v537 = v529 + base.B2i32(int32(-65) < v534)
	v538 = int32(1)
	v541 = v531 + v538
	if v541 != v488 {
		v528 = v528 + v538
		v529 = v537
		v531 = v541
		goto L142
	} else {
		goto L144
	}
L143:
	;
	v544 = v537
	goto L141
L144:
	;
	goto L143
L145:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v558
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v560
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v560
	v565 = F_find_among_b(m, l0, int32(4319344), int32(26))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L2
	} else {
		goto L146
	}
L146:
	;
	if v565 == int32(0) {
		v632 = v474
		goto L127
	} else {
		goto L147
	}
L147:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v569
	switch v565 - int32(1) {
	case 0:
		goto L151
	case 1:
		goto L150
	case 2:
		goto L149
	default:
		v598 = v474
		goto L148
	}
L148:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v600
	v603 = v598
	goto L161
L149:
	;
	v594 = F_slice_del(m, l0)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L2
	} else {
		goto L158
	}
L150:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v582 = F_find_among_b(m, l0, int32(4319872), int32(8))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L2
	} else {
		goto L154
	}
L151:
	;
	v575 = F_slice_from_s(m, l0, int32(3), int32(2172328))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L2
	} else {
		goto L152
	}
L152:
	;
	if int32(0) <= v575 {
		v598 = v575
		goto L148
	} else {
		goto L153
	}
L153:
	;
	v632 = v575
	goto L127
L154:
	;
	if v582 != 0 {
		v632 = v474
		goto L127
	} else {
		goto L155
	}
L155:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v584 + (v569 - v579)
	v590 = F_slice_from_s(m, l0, int32(3), int32(2172331))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L2
	} else {
		goto L156
	}
L156:
	;
	if int32(0) <= v590 {
		v598 = v590
		goto L148
	} else {
		goto L157
	}
L157:
	;
	v632 = v590
	goto L127
L158:
	;
	if v594 < int32(0) {
		v632 = v594
		goto L127
	} else {
		goto L159
	}
L159:
	;
	v598 = v594
	goto L148
L160:
	;
	v632 = int32(1)
	goto L127
L161:
	;
	v609 = F_r_fix_ending(m, l0)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L2
	} else {
		goto L166
	}
L162:
	;
	if v617 == int32(0) {
		goto L160
	} else {
		goto L177
	}
L163:
	;
	if v609 < int32(0) {
		goto L170
	} else {
		goto L171
	}
L164:
	;
	v617 = int32(2)
	goto L163
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v600
	goto L160
L166:
	;
	v612 = int32(base.Ui32(v609) >> (uint(int32(31)) % 32))
	if v609 != 0 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v614 = v612
	goto L169
L168:
	;
	v614 = int32(4)
	goto L169
L169:
	;
	switch v614 {
	case 0:
		goto L164
	default:
		v617 = v612
		goto L163
	case 4:
		goto L165
	}
L170:
	;
	v620 = v609
	goto L172
L171:
	;
	v620 = v603
	goto L172
L172:
	;
	if v609 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v621 = v620
	goto L175
L174:
	;
	v621 = v603
	goto L175
L175:
	;
	if v617 == int32(2) {
		v603 = v621
		goto L161
	} else {
		goto L176
	}
L176:
	;
	goto L162
L177:
	;
	if v621 < int32(0) {
		v632 = v621
		goto L127
	} else {
		goto L178
	}
L178:
	;
	goto L160
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103
	v641 = int32(0)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v642))) = v641
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v645-int32(4))))
	if v653 == v641 {
		goto L182
	} else {
		goto L183
	}
L180:
	;
	if v1058 < int32(0) {
		v1585 = v1058
		goto L1
	} else {
		goto L300
	}
L181:
	;
	if v725 < int32(5) {
		v1058 = v641
		goto L180
	} else {
		goto L198
	}
L182:
	;
	v725 = int32(0)
	goto L181
L183:
	;
	goto L184
L184:
	;
	v658 = v653 & int32(3)
	if base.Ui32(v653) < base.Ui32(int32(4)) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	if v658 != 0 {
		goto L192
	} else {
		goto L193
	}
L186:
	;
	v692 = v645
	v693 = int32(0)
	goto L185
L187:
	;
	goto L188
L188:
	;
	v665 = v645
	v666 = int32(0)
	v669 = v641
	goto L189
L189:
	;
	v671 = int32(*(*int8)(unsafe.Add(mBase, uint32(v665))))
	v672 = int32(-65)
	v675 = int32(*(*int8)(unsafe.Add(mBase, uint32(v665)+1)))
	v679 = int32(*(*int8)(unsafe.Add(mBase, uint32(v665)+2)))
	v683 = int32(*(*int8)(unsafe.Add(mBase, uint32(v665)+3)))
	v686 = v666 + base.B2i32(v672 < v671) + base.B2i32(v672 < v675) + base.B2i32(v672 < v679) + base.B2i32(v672 < v683)
	v687 = int32(4)
	v688 = v665 + v687
	v690 = v669 + v687
	if v690 != v653&int32(-4) {
		v665 = v688
		v666 = v686
		v669 = v690
		goto L189
	} else {
		goto L191
	}
L190:
	;
	v692 = v688
	v693 = v686
	goto L185
L191:
	;
	goto L190
L192:
	;
	v698 = v692
	v699 = v693
	v701 = v641
	goto L195
L193:
	;
	v714 = v693
	goto L194
L194:
	;
	v725 = v714
	goto L181
L195:
	;
	v704 = int32(*(*int8)(unsafe.Add(mBase, uint32(v698))))
	v707 = v699 + base.B2i32(int32(-65) < v704)
	v708 = int32(1)
	v711 = v701 + v708
	if v711 != v658 {
		v698 = v698 + v708
		v699 = v707
		v701 = v711
		goto L195
	} else {
		goto L197
	}
L196:
	;
	v714 = v707
	goto L194
L197:
	;
	goto L196
L198:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v728
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v730
	if v730-int32(2) <= v728 {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v993
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v995))) = int32(1)
	v998 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v998
	v1000 = int32(9)
	v1002 = int32(0)
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v998-v1005 < v1000 {
		v1015 = v1002
		goto L273
	} else {
		goto L274
	}
L200:
	;
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v924
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v924
	v927 = int32(3)
	v929 = int32(0)
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v924-v932 < v927 {
		v942 = v929
		goto L255
	} else {
		goto L256
	}
L201:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736+v730-int32(1)))))
	if v740&int32(224) != int32(128) {
		goto L200
	} else {
		goto L202
	}
L202:
	;
	if int32(1)<<(uint(v740)%32)&int32(-2147475197) == int32(0) {
		goto L200
	} else {
		goto L203
	}
L203:
	;
	v753 = F_find_among_b(m, l0, int32(4320032), int32(22))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L2
	} else {
		goto L204
	}
L204:
	;
	if v753 == int32(0) {
		goto L200
	} else {
		goto L205
	}
L205:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v757
	switch v753 - int32(1) {
	case 0:
		goto L212
	case 1:
		goto L211
	case 2:
		goto L210
	case 3:
		goto L209
	case 4:
		goto L208
	case 5:
		goto L207
	case 6:
		goto L206
	default:
		v990 = v641
		goto L199
	}
L206:
	;
	v917 = F_slice_from_s(m, l0, int32(3), int32(2172725))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L2
	} else {
		goto L252
	}
L207:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v905 = F_find_among_b(m, l0, int32(4320640), int32(8))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L2
	} else {
		goto L248
	}
L208:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v890 = F_find_among_b(m, l0, int32(4320480), int32(8))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L2
	} else {
		goto L244
	}
L209:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v799 = int32(0)
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v798-int32(4))))
	if v806 == v799 {
		goto L225
	} else {
		goto L226
	}
L210:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v772 = int32(3)
	v774 = int32(0)
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v776-v777 < v772 {
		v787 = v774
		goto L218
	} else {
		goto L219
	}
L211:
	;
	v767 = F_slice_from_s(m, l0, int32(3), int32(2172710))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L2
	} else {
		goto L215
	}
L212:
	;
	v761 = F_slice_del(m, l0)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L2
	} else {
		goto L213
	}
L213:
	;
	if int32(0) <= v761 {
		v990 = v761
		goto L199
	} else {
		goto L214
	}
L214:
	;
	v1058 = v761
	goto L180
L215:
	;
	if int32(0) <= v767 {
		v990 = v767
		goto L199
	} else {
		goto L216
	}
L216:
	;
	v1058 = v767
	goto L180
L217:
	;
	if v787 != 0 {
		goto L200
	} else {
		goto L221
	}
L218:
	;
	goto L217
L219:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v783 = F_memcmp(m, v780+v776-v772, int32(2172713), v772)
	mBase = m.M
	if v783 != 0 {
		v787 = v774
		goto L218
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v776 - v772
	v787 = int32(1)
	goto L218
L221:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v788 + (v757 - v771)
	v794 = F_slice_from_s(m, l0, int32(3), int32(2172716))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L2
	} else {
		goto L222
	}
L222:
	;
	if int32(0) <= v794 {
		v990 = v794
		goto L199
	} else {
		goto L223
	}
L223:
	;
	v1058 = v794
	goto L180
L224:
	;
	if v878 < int32(7) {
		goto L200
	} else {
		goto L241
	}
L225:
	;
	v878 = int32(0)
	goto L224
L226:
	;
	goto L227
L227:
	;
	v811 = v806 & int32(3)
	if base.Ui32(v806) < base.Ui32(int32(4)) {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	if v811 != 0 {
		goto L235
	} else {
		goto L236
	}
L229:
	;
	v845 = v798
	v846 = int32(0)
	goto L228
L230:
	;
	goto L231
L231:
	;
	v818 = v798
	v819 = int32(0)
	v822 = v799
	goto L232
L232:
	;
	v824 = int32(*(*int8)(unsafe.Add(mBase, uint32(v818))))
	v825 = int32(-65)
	v828 = int32(*(*int8)(unsafe.Add(mBase, uint32(v818)+1)))
	v832 = int32(*(*int8)(unsafe.Add(mBase, uint32(v818)+2)))
	v836 = int32(*(*int8)(unsafe.Add(mBase, uint32(v818)+3)))
	v839 = v819 + base.B2i32(v825 < v824) + base.B2i32(v825 < v828) + base.B2i32(v825 < v832) + base.B2i32(v825 < v836)
	v840 = int32(4)
	v841 = v818 + v840
	v843 = v822 + v840
	if v843 != v806&int32(-4) {
		v818 = v841
		v819 = v839
		v822 = v843
		goto L232
	} else {
		goto L234
	}
L233:
	;
	v845 = v841
	v846 = v839
	goto L228
L234:
	;
	goto L233
L235:
	;
	v851 = v845
	v852 = v846
	v854 = v799
	goto L238
L236:
	;
	v867 = v846
	goto L237
L237:
	;
	v878 = v867
	goto L224
L238:
	;
	v857 = int32(*(*int8)(unsafe.Add(mBase, uint32(v851))))
	v860 = v852 + base.B2i32(int32(-65) < v857)
	v861 = int32(1)
	v864 = v854 + v861
	if v864 != v811 {
		v851 = v851 + v861
		v852 = v860
		v854 = v864
		goto L238
	} else {
		goto L240
	}
L239:
	;
	v867 = v860
	goto L237
L240:
	;
	goto L239
L241:
	;
	v883 = F_slice_from_s(m, l0, int32(3), int32(2172719))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L2
	} else {
		goto L242
	}
L242:
	;
	if int32(0) <= v883 {
		v990 = v883
		goto L199
	} else {
		goto L243
	}
L243:
	;
	v1058 = v883
	goto L180
L244:
	;
	if v890 != 0 {
		goto L200
	} else {
		goto L245
	}
L245:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v892 + (v757 - v887)
	v898 = F_slice_from_s(m, l0, int32(3), int32(2172722))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L2
	} else {
		goto L246
	}
L246:
	;
	if int32(0) <= v898 {
		v990 = v898
		goto L199
	} else {
		goto L247
	}
L247:
	;
	v1058 = v898
	goto L180
L248:
	;
	if v905 != 0 {
		goto L200
	} else {
		goto L249
	}
L249:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v907 + (v757 - v902)
	v911 = F_slice_del(m, l0)
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L2
	} else {
		goto L250
	}
L250:
	;
	if int32(0) <= v911 {
		v990 = v911
		goto L199
	} else {
		goto L251
	}
L251:
	;
	v1058 = v911
	goto L180
L252:
	;
	if int32(0) <= v917 {
		v990 = v917
		goto L199
	} else {
		goto L253
	}
L253:
	;
	v1058 = v917
	goto L180
L254:
	;
	if v942 == int32(0) {
		v1058 = v641
		goto L180
	} else {
		goto L258
	}
L255:
	;
	goto L254
L256:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v938 = F_memcmp(m, v935+v924-v927, int32(2172728), v927)
	mBase = m.M
	if v938 != 0 {
		v942 = v929
		goto L255
	} else {
		goto L257
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v924 - v927
	v942 = int32(1)
	goto L255
L258:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v947 = v945 - v946
	v950 = F_find_among_b(m, l0, int32(4320800), int32(6))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L2
	} else {
		goto L259
	}
L259:
	;
	if v950 != 0 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v952 - v947
	v957 = F_find_among_b(m, l0, int32(4320928), int32(6))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L2
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v980 = v979 - v947
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v980
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v980
	v985 = F_slice_from_s(m, l0, int32(3), int32(2172734))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L2
	} else {
		goto L270
	}
L263:
	;
	if v957 == int32(0) {
		v1058 = v641
		goto L180
	} else {
		goto L264
	}
L264:
	;
	v961 = int32(3)
	v963 = int32(0)
	v965 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v965-v966 < v961 {
		v976 = v963
		goto L266
	} else {
		goto L267
	}
L265:
	;
	if v976 == int32(0) {
		v1058 = v641
		goto L180
	} else {
		goto L269
	}
L266:
	;
	goto L265
L267:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v972 = F_memcmp(m, v969+v965-v961, int32(2172731), v961)
	mBase = m.M
	if v972 != 0 {
		v976 = v963
		goto L266
	} else {
		goto L268
	}
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v965 - v961
	v976 = int32(1)
	goto L266
L269:
	;
	goto L262
L270:
	;
	if v985 < int32(0) {
		v1058 = v985
		goto L180
	} else {
		goto L271
	}
L271:
	;
	v990 = v985
	goto L199
L272:
	;
	if v1015 != 0 {
		goto L276
	} else {
		goto L277
	}
L273:
	;
	goto L272
L274:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1011 = F_memcmp(m, v1008+v998-v1000, int32(2172737), v1000)
	mBase = m.M
	if v1011 != 0 {
		v1015 = v1002
		goto L273
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v998 - v1000
	v1015 = int32(1)
	goto L273
L276:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1016
	v1020 = F_slice_from_s(m, l0, int32(3), int32(2172746))
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L2
	} else {
		goto L279
	}
L277:
	;
	v1024 = v990
	goto L278
L278:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1025
	v1029 = v1024
	goto L282
L279:
	;
	if v1020 < int32(0) {
		v1058 = v1020
		goto L180
	} else {
		goto L280
	}
L280:
	;
	v1024 = v1020
	goto L278
L281:
	;
	v1058 = int32(1)
	goto L180
L282:
	;
	v1034 = F_r_fix_ending(m, l0)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L2
	} else {
		goto L287
	}
L283:
	;
	if v1042 == int32(0) {
		goto L281
	} else {
		goto L298
	}
L284:
	;
	if v1034 < int32(0) {
		goto L291
	} else {
		goto L292
	}
L285:
	;
	v1042 = int32(2)
	goto L284
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1025
	goto L281
L287:
	;
	v1037 = int32(base.Ui32(v1034) >> (uint(int32(31)) % 32))
	if v1034 != 0 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1039 = v1037
	goto L290
L289:
	;
	v1039 = int32(4)
	goto L290
L290:
	;
	switch v1039 {
	case 0:
		goto L285
	default:
		v1042 = v1037
		goto L284
	case 4:
		goto L286
	}
L291:
	;
	v1045 = v1034
	goto L293
L292:
	;
	v1045 = v1029
	goto L293
L293:
	;
	if v1034 != 0 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1046 = v1045
	goto L296
L295:
	;
	v1046 = v1029
	goto L296
L296:
	;
	if v1042 == int32(2) {
		v1029 = v1046
		goto L282
	} else {
		goto L297
	}
L297:
	;
	goto L283
L298:
	;
	if v1046 < int32(0) {
		v1058 = v1046
		goto L180
	} else {
		goto L299
	}
L299:
	;
	goto L281
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103
	v1066 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v103
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1069
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1069
	if v1069-int32(8) <= v103 {
		v1134 = v1066
		goto L301
	} else {
		goto L302
	}
L301:
	;
	if v1134 < int32(0) {
		v1585 = v1134
		goto L1
	} else {
		goto L325
	}
L302:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075+v1069-int32(1)))))
	if v1079 != int32(141) {
		v1134 = v1066
		goto L301
	} else {
		goto L303
	}
L303:
	;
	v1084 = F_find_among_b(m, l0, int32(4321056), int32(4))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L2
	} else {
		goto L304
	}
L304:
	;
	if v1084 == int32(0) {
		v1134 = v1066
		goto L301
	} else {
		goto L305
	}
L305:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1088
	switch v1084 - int32(1) {
	case 0:
		goto L310
	case 1:
		goto L309
	case 2:
		goto L308
	case 3:
		goto L307
	default:
		goto L306
	}
L306:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1131
	v1134 = int32(1)
	goto L301
L307:
	;
	v1125 = F_slice_del(m, l0)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L2
	} else {
		goto L323
	}
L308:
	;
	v1121 = F_slice_from_s(m, l0, int32(6), int32(2173078))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L2
	} else {
		goto L321
	}
L309:
	;
	v1115 = F_slice_from_s(m, l0, int32(6), int32(2173072))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L2
	} else {
		goto L319
	}
L310:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1095 = F_find_among_b(m, l0, int32(4321136), int32(6))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L2
	} else {
		goto L311
	}
L311:
	;
	if v1095 != 0 {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1099 = F_slice_from_s(m, l0, int32(9), int32(2173060))
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L2
	} else {
		goto L315
	}
L313:
	;
	goto L314
L314:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1103 + (v1088 - v1092)
	v1109 = F_slice_from_s(m, l0, int32(3), int32(2173069))
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L2
	} else {
		goto L317
	}
L315:
	;
	if int32(0) <= v1099 {
		goto L306
	} else {
		goto L316
	}
L316:
	;
	v1134 = v1099
	goto L301
L317:
	;
	if int32(0) <= v1109 {
		goto L306
	} else {
		goto L318
	}
L318:
	;
	v1134 = v1109
	goto L301
L319:
	;
	if int32(0) <= v1115 {
		goto L306
	} else {
		goto L320
	}
L320:
	;
	v1134 = v1115
	goto L301
L321:
	;
	if int32(0) <= v1121 {
		goto L306
	} else {
		goto L322
	}
L322:
	;
	v1134 = v1121
	goto L301
L323:
	;
	if v1125 < int32(0) {
		v1134 = v1125
		goto L301
	} else {
		goto L324
	}
L324:
	;
	goto L306
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103
	v1140 = int32(0)
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1141-int32(4))))
	if v1149 == v1140 {
		goto L328
	} else {
		goto L329
	}
L326:
	;
	if v1255 < int32(0) {
		v1585 = v1255
		goto L1
	} else {
		goto L351
	}
L327:
	;
	if v1221 < int32(5) {
		v1255 = v1140
		goto L326
	} else {
		goto L344
	}
L328:
	;
	v1221 = int32(0)
	goto L327
L329:
	;
	goto L330
L330:
	;
	v1154 = v1149 & int32(3)
	if base.Ui32(v1149) < base.Ui32(int32(4)) {
		goto L332
	} else {
		goto L333
	}
L331:
	;
	if v1154 != 0 {
		goto L338
	} else {
		goto L339
	}
L332:
	;
	v1188 = v1141
	v1189 = int32(0)
	goto L331
L333:
	;
	goto L334
L334:
	;
	v1161 = v1141
	v1162 = int32(0)
	v1165 = v1140
	goto L335
L335:
	;
	v1167 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1161))))
	v1168 = int32(-65)
	v1171 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1161)+1)))
	v1175 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1161)+2)))
	v1179 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1161)+3)))
	v1182 = v1162 + base.B2i32(v1168 < v1167) + base.B2i32(v1168 < v1171) + base.B2i32(v1168 < v1175) + base.B2i32(v1168 < v1179)
	v1183 = int32(4)
	v1184 = v1161 + v1183
	v1186 = v1165 + v1183
	if v1186 != v1149&int32(-4) {
		v1161 = v1184
		v1162 = v1182
		v1165 = v1186
		goto L335
	} else {
		goto L337
	}
L336:
	;
	v1188 = v1184
	v1189 = v1182
	goto L331
L337:
	;
	goto L336
L338:
	;
	v1194 = v1188
	v1195 = v1189
	v1197 = v1140
	goto L341
L339:
	;
	v1210 = v1189
	goto L340
L340:
	;
	v1221 = v1210
	goto L327
L341:
	;
	v1200 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1194))))
	v1203 = v1195 + base.B2i32(int32(-65) < v1200)
	v1204 = int32(1)
	v1207 = v1197 + v1204
	if v1207 != v1154 {
		v1194 = v1194 + v1204
		v1195 = v1203
		v1197 = v1207
		goto L341
	} else {
		goto L343
	}
L342:
	;
	v1210 = v1203
	goto L340
L343:
	;
	goto L342
L344:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1224
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1226
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1226
	if v1226-int32(5) <= v1224 {
		v1255 = v1140
		goto L326
	} else {
		goto L345
	}
L345:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1232+v1226-int32(1)))))
	if v1236 != int32(191) {
		v1255 = v1140
		goto L326
	} else {
		goto L346
	}
L346:
	;
	v1241 = F_find_among_b(m, l0, int32(4321264), int32(2))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L2
	} else {
		goto L347
	}
L347:
	;
	if v1241 == int32(0) {
		v1255 = v1140
		goto L326
	} else {
		goto L348
	}
L348:
	;
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1245
	v1247 = F_slice_del(m, l0)
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L2
	} else {
		goto L349
	}
L349:
	;
	if v1247 < int32(0) {
		v1255 = v1247
		goto L326
	} else {
		goto L350
	}
L350:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1251
	v1255 = int32(1)
	goto L326
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1261)+4)) = int32(1)
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1268 = v1261 + int32(4)
	v1269 = int32(0)
	goto L353
L352:
	;
	if v1575 < int32(0) {
		v1585 = v1575
		goto L1
	} else {
		goto L441
	}
L353:
	;
	v1274 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1268))) = v1274
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1276-int32(4))))
	if v1284 == v1274 {
		goto L357
	} else {
		goto L358
	}
L354:
	;
	v1575 = int32(1)
	goto L352
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1266
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1568)+4))
	if v1571 != 0 {
		v1268 = v1568 + int32(4)
		v1269 = v1568
		goto L353
	} else {
		goto L440
	}
L356:
	;
	if v1356 < int32(5) {
		goto L355
	} else {
		goto L373
	}
L357:
	;
	v1356 = int32(0)
	goto L356
L358:
	;
	goto L359
L359:
	;
	v1289 = v1284 & int32(3)
	if base.Ui32(v1284) < base.Ui32(int32(4)) {
		goto L361
	} else {
		goto L362
	}
L360:
	;
	if v1289 != 0 {
		goto L367
	} else {
		goto L368
	}
L361:
	;
	v1323 = v1276
	v1324 = int32(0)
	goto L360
L362:
	;
	goto L363
L363:
	;
	v1296 = v1276
	v1297 = int32(0)
	v1300 = v1274
	goto L364
L364:
	;
	v1302 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1296))))
	v1303 = int32(-65)
	v1306 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1296)+1)))
	v1310 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1296)+2)))
	v1314 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1296)+3)))
	v1317 = v1297 + base.B2i32(v1303 < v1302) + base.B2i32(v1303 < v1306) + base.B2i32(v1303 < v1310) + base.B2i32(v1303 < v1314)
	v1318 = int32(4)
	v1319 = v1296 + v1318
	v1321 = v1300 + v1318
	if v1321 != v1284&int32(-4) {
		v1296 = v1319
		v1297 = v1317
		v1300 = v1321
		goto L364
	} else {
		goto L366
	}
L365:
	;
	v1323 = v1319
	v1324 = v1317
	goto L360
L366:
	;
	goto L365
L367:
	;
	v1329 = v1323
	v1330 = v1324
	v1332 = v1274
	goto L370
L368:
	;
	v1345 = v1324
	goto L369
L369:
	;
	v1356 = v1345
	goto L356
L370:
	;
	v1335 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1329))))
	v1338 = v1330 + base.B2i32(int32(-65) < v1335)
	v1339 = int32(1)
	v1342 = v1332 + v1339
	if v1342 != v1289 {
		v1329 = v1329 + v1339
		v1330 = v1338
		v1332 = v1342
		goto L370
	} else {
		goto L372
	}
L371:
	;
	v1345 = v1338
	goto L369
L372:
	;
	goto L371
L373:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1359
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1361
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1361
	v1366 = F_find_among_b(m, l0, int32(4321312), int32(46))
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L2
	} else {
		goto L375
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1499
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1499
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1499-int32(8) <= v1502 {
		v1532 = v1496
		goto L416
	} else {
		goto L417
	}
L375:
	;
	if v1366 != 0 {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1368
	switch v1366 - int32(1) {
	case 0:
		goto L385
	case 1:
		goto L384
	case 2:
		goto L383
	case 3:
		goto L382
	case 4:
		goto L381
	case 5:
		goto L380
	default:
		v1484 = v1368
		goto L379
	}
L377:
	;
	v1491 = v1269
	goto L378
L378:
	;
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1496 = v1491
	v1499 = v1494
	goto L374
L379:
	;
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1487)+4)) = int32(1)
	v1491 = v1484
	goto L378
L380:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1457 = int32(3)
	v1459 = int32(0)
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1461-v1462 < v1457 {
		v1472 = v1459
		goto L410
	} else {
		goto L411
	}
L381:
	;
	v1452 = F_slice_from_s(m, l0, int32(3), int32(2173188))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L2
	} else {
		goto L407
	}
L382:
	;
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1424 = int32(3)
	v1426 = int32(0)
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1428-v1429 < v1424 {
		v1439 = v1426
		goto L401
	} else {
		goto L402
	}
L383:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1413 = F_find_among_b(m, l0, int32(4322480), int32(8))
	mBase = m.M
	v1414 = m.ExcPending
	if v1414 != 0 {
		goto L2
	} else {
		goto L396
	}
L384:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1368-int32(2) <= v1377 {
		v1401 = v1376
		goto L388
	} else {
		goto L389
	}
L385:
	;
	v1372 = F_slice_del(m, l0)
	mBase = m.M
	v1373 = m.ExcPending
	if v1373 != 0 {
		goto L2
	} else {
		goto L386
	}
L386:
	;
	if int32(0) <= v1372 {
		v1484 = v1372
		goto L379
	} else {
		goto L387
	}
L387:
	;
	v1575 = v1372
	goto L352
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1368 - v1376 + v1401
	v1406 = F_slice_del(m, l0)
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L2
	} else {
		goto L394
	}
L389:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1381+v1368-int32(1)))))
	if v1385&int32(224) != int32(128) {
		v1401 = v1376
		goto L388
	} else {
		goto L390
	}
L390:
	;
	if int32(1)<<(uint(v1385)%32)&int32(1951712) == int32(0) {
		v1401 = v1376
		goto L388
	} else {
		goto L391
	}
L391:
	;
	v1398 = F_find_among_b(m, l0, int32(4322240), int32(12))
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L2
	} else {
		goto L392
	}
L392:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1398 != 0 {
		v1496 = v1368
		v1499 = v1400
		goto L374
	} else {
		goto L393
	}
L393:
	;
	v1401 = v1400
	goto L388
L394:
	;
	if int32(0) <= v1406 {
		v1484 = v1406
		goto L379
	} else {
		goto L395
	}
L395:
	;
	v1575 = v1406
	goto L352
L396:
	;
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1413 != 0 {
		v1496 = v1368
		v1499 = v1415
		goto L374
	} else {
		goto L397
	}
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1368 - v1410 + v1415
	v1419 = F_slice_del(m, l0)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L2
	} else {
		goto L398
	}
L398:
	;
	if int32(0) <= v1419 {
		v1484 = v1419
		goto L379
	} else {
		goto L399
	}
L399:
	;
	v1575 = v1419
	goto L352
L400:
	;
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1439 != 0 {
		v1496 = v1368
		v1499 = v1440
		goto L374
	} else {
		goto L404
	}
L401:
	;
	goto L400
L402:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1435 = F_memcmp(m, v1432+v1428-v1424, int32(2173182), v1424)
	mBase = m.M
	if v1435 != 0 {
		v1439 = v1426
		goto L401
	} else {
		goto L403
	}
L403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1428 - v1424
	v1439 = int32(1)
	goto L401
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1368 - v1423 + v1440
	v1446 = F_slice_from_s(m, l0, int32(3), int32(2173185))
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L2
	} else {
		goto L405
	}
L405:
	;
	if int32(0) <= v1446 {
		v1484 = v1446
		goto L379
	} else {
		goto L406
	}
L406:
	;
	v1575 = v1446
	goto L352
L407:
	;
	if int32(0) <= v1452 {
		v1484 = v1452
		goto L379
	} else {
		goto L408
	}
L408:
	;
	v1575 = v1452
	goto L352
L409:
	;
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1472 == int32(0) {
		v1496 = v1368
		v1499 = v1473
		goto L374
	} else {
		goto L413
	}
L410:
	;
	goto L409
L411:
	;
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1468 = F_memcmp(m, v1465+v1461-v1457, int32(2173191), v1457)
	mBase = m.M
	if v1468 != 0 {
		v1472 = v1459
		goto L410
	} else {
		goto L412
	}
L412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1461 - v1457
	v1472 = int32(1)
	goto L410
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1368 - v1456 + v1473
	v1479 = F_slice_del(m, l0)
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L2
	} else {
		goto L414
	}
L414:
	;
	if v1479 < int32(0) {
		v1575 = v1479
		goto L352
	} else {
		goto L415
	}
L415:
	;
	v1484 = v1479
	goto L379
L416:
	;
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1533
	v1537 = v1532
	goto L423
L417:
	;
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1506+v1499-int32(1)))))
	if base.B2i32(v1510 != int32(177))&base.B2i32(v1510 != int32(141)) != 0 {
		v1532 = v1496
		goto L416
	} else {
		goto L418
	}
L418:
	;
	v1518 = F_find_among_b(m, l0, int32(4322640), int32(6))
	mBase = m.M
	v1519 = m.ExcPending
	if v1519 != 0 {
		goto L2
	} else {
		goto L419
	}
L419:
	;
	if v1518 == int32(0) {
		v1532 = v1496
		goto L416
	} else {
		goto L420
	}
L420:
	;
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1522
	v1524 = F_slice_del(m, l0)
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L2
	} else {
		goto L421
	}
L421:
	;
	if v1524 < int32(0) {
		v1575 = v1524
		goto L352
	} else {
		goto L422
	}
L422:
	;
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1528)+4)) = int32(1)
	v1532 = v1524
	goto L416
L423:
	;
	v1542 = F_r_fix_ending(m, l0)
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L2
	} else {
		goto L427
	}
L424:
	;
	if v1549 == int32(0) {
		goto L355
	} else {
		goto L438
	}
L425:
	;
	if v1542 < int32(0) {
		goto L431
	} else {
		goto L432
	}
L426:
	;
	v1549 = int32(2)
	goto L425
L427:
	;
	v1545 = int32(base.Ui32(v1542) >> (uint(int32(31)) % 32))
	if v1542 != 0 {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v1547 = v1545
	goto L430
L429:
	;
	v1547 = int32(4)
	goto L430
L430:
	;
	switch v1547 {
	case 0:
		goto L426
	default:
		v1549 = v1545
		goto L425
	case 4:
		goto L355
	}
L431:
	;
	v1552 = v1542
	goto L433
L432:
	;
	v1552 = v1537
	goto L433
L433:
	;
	if v1542 != 0 {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v1553 = v1552
	goto L436
L435:
	;
	v1553 = v1537
	goto L436
L436:
	;
	if v1549 == int32(2) {
		v1537 = v1553
		goto L423
	} else {
		goto L437
	}
L437:
	;
	goto L424
L438:
	;
	if v1553 < int32(0) {
		v1575 = v1553
		goto L352
	} else {
		goto L439
	}
L439:
	;
	goto L355
L440:
	;
	goto L354
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103
	v1585 = int32(1)
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
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
	if v13&int32(3) == int32(0) {
		v68 = v13
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v17 = int32(4)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v19&int32(254) == int32(2) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v32 = int32(1)
	if v14&v32 != 0 {
		v44 = int32(base.Ui32(v14)>>(uint(v32)%32)) - v32
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v28 = v17
	goto L9
L8:
	;
	v28 = base.B2i32(v19 == int32(18)) << (uint(v17) % 32)
	goto L9
L9:
	;
	if v19 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v31 = v17
	goto L12
L11:
	;
	v31 = v28
	goto L12
L12:
	;
	v44 = v31
	goto L3
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L14:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v102 != int32(950) {
		goto L33
	} else {
		goto L34
	}
L15:
	;
	v101 = v93 - v13
	goto L14
L16:
	;
	v72 = v68
	goto L25
L17:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v52 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v101 = int32(0)
	goto L14
L19:
	;
	goto L20
L20:
	;
	v57 = v13
	goto L21
L21:
	;
	v61 = v57 + int32(1)
	if v61&int32(3) == int32(0) {
		v68 = v61
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v93 = v61
	goto L15
L23:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v66 != 0 {
		v57 = v61
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v81 = int32(-2139062144)
	if (int32(16843008)-v78|v78)&v81 == v81 {
		v72 = v72 + int32(4)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v87 = v72
	goto L28
L27:
	;
	goto L26
L28:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v91 != 0 {
		v87 = v87 + int32(1)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v93 = v87
	goto L15
L30:
	;
	goto L29
L31:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v208 != v9 {
		goto L68
	} else {
		goto L69
	}
L32:
	;
	v197 = int32(1)
	if v14&v197 != 0 {
		goto L64
	} else {
		goto L65
	}
L33:
	;
	if v102 != 0 {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v44 != v101 {
		v207 = int32(0)
		goto L31
	} else {
		goto L42
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(232906), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errhint(m, int32(534683), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(476708), int32(1648), int32(98701))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v127 = int32(1)
	if v14&v127 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v131 = v127
	goto L45
L44:
	;
	v131 = int32(4)
	goto L45
L45:
	;
	v132 = v9 + v131
	if base.Ui32(int32(4)) <= base.Ui32(v44) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v207 = base.B2i32(v194 == int32(0))
	goto L31
L47:
	;
	v194 = int32(0)
	goto L46
L48:
	;
	v168 = v163
	v169 = v164
	v170 = v165
	goto L58
L49:
	;
	if (v132|v13)&int32(3) != 0 {
		v163 = v132
		v164 = v13
		v165 = v44
		goto L48
	} else {
		goto L52
	}
L50:
	;
	v156 = v132
	v157 = v13
	v158 = v44
	goto L51
L51:
	;
	if v158 == int32(0) {
		goto L47
	} else {
		goto L57
	}
L52:
	;
	v140 = v132
	v141 = v13
	v142 = v44
	goto L53
L53:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v145 != v146 {
		v163 = v140
		v164 = v141
		v165 = v142
		goto L48
	} else {
		goto L55
	}
L54:
	;
	v156 = v151
	v157 = v149
	v158 = v153
	goto L51
L55:
	;
	v148 = int32(4)
	v149 = v141 + v148
	v151 = v140 + v148
	v153 = v142 - v148
	if base.Ui32(int32(3)) < base.Ui32(v153) {
		v140 = v151
		v141 = v149
		v142 = v153
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v163 = v156
	v164 = v157
	v165 = v158
	goto L48
L58:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if v173 == v174 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v194 = v173 - v174
	goto L46
L60:
	;
	v176 = int32(1)
	v181 = v170 - v176
	if v181 != 0 {
		v168 = v168 + v176
		v169 = v169 + v176
		v170 = v181
		goto L58
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	goto L47
L64:
	;
	v201 = v197
	goto L66
L65:
	;
	v201 = int32(4)
	goto L66
L66:
	;
	v203 = F_varstr_cmp(m, v9+v201, v44, v13, v101, v102)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v207 = base.B2i32(v203 == int32(0))
	goto L31
L68:
	;
	F_pfree(m, v9)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	return v207
L71:
	;
	goto L70
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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9*int32(28))+uint32(_consts[863])))
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
				v36 = int32(4)
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
				if v38&int32(254) == int32(2) {
					v47 = v36
				} else {
					v47 = base.B2i32(v38 == int32(18)) << (uint(v36) % 32)
				}
				if v38 == int32(1) {
					v50 = v36
				} else {
					v50 = v47
				}
				v51 = F_pg_mbstrlen_with_len(m, v33, v50)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					return v51
				}
			} else {
				if v32 != 0 {
					v54 = int32(1)
					v58 = F_pg_mbstrlen_with_len(m, v33, int32(base.Ui32(v30)>>(uint(v54)%32))-v54)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						return v58
					}
				} else {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
					v66 = F_pg_mbstrlen_with_len(m, v33, int32(base.Ui32(v61)>>(uint(int32(2))%32))-int32(4))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						return v66
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
				v25 = int32(4)
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
				if v27&int32(254) == int32(2) {
					v36 = v25
				} else {
					v36 = base.B2i32(v27 == int32(18)) << (uint(v25) % 32)
				}
				if v27 == int32(1) {
					v39 = v25
				} else {
					v39 = v36
				}
				v50 = v39
			} else {
				v40 = int32(1)
				if v21 != 0 {
					v50 = int32(base.Ui32(v19)>>(uint(v40)%32)) - v40
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v51 = int32(1)
			v52 = v15 + v51
			v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			v57 = v55 & v51
			if v57 != 0 {
				v58 = v52
			} else {
				v58 = v15 + int32(4)
			}
			if v55 == int32(1) {
				v61 = int32(4)
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
				if v63&int32(254) == int32(2) {
					v72 = v61
				} else {
					v72 = base.B2i32(v63 == int32(18)) << (uint(v61) % 32)
				}
				if v63 == int32(1) {
					v75 = v61
				} else {
					v75 = v72
				}
				v86 = v75
			} else {
				v76 = int32(1)
				if v57 != 0 {
					v86 = int32(base.Ui32(v55)>>(uint(v76)%32)) - v76
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v86 = int32(base.Ui32(v80)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v88 = F_GenericMatchText(m, v22, v50, v58, v86, v87)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v88 == int32(1))
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
				v25 = int32(4)
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
				if v27&int32(254) == int32(2) {
					v36 = v25
				} else {
					v36 = base.B2i32(v27 == int32(18)) << (uint(v25) % 32)
				}
				if v27 == int32(1) {
					v39 = v25
				} else {
					v39 = v36
				}
				v50 = v39
			} else {
				v40 = int32(1)
				if v21 != 0 {
					v50 = int32(base.Ui32(v19)>>(uint(v40)%32)) - v40
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v51 = int32(1)
			v52 = v15 + v51
			v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
			v57 = v55 & v51
			if v57 != 0 {
				v58 = v52
			} else {
				v58 = v15 + int32(4)
			}
			if v55 == int32(1) {
				v61 = int32(4)
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
				if v63&int32(254) == int32(2) {
					v72 = v61
				} else {
					v72 = base.B2i32(v63 == int32(18)) << (uint(v61) % 32)
				}
				if v63 == int32(1) {
					v75 = v61
				} else {
					v75 = v72
				}
				v86 = v75
			} else {
				v76 = int32(1)
				if v57 != 0 {
					v86 = int32(base.Ui32(v55)>>(uint(v76)%32)) - v76
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v86 = int32(base.Ui32(v80)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v88 = F_GenericMatchText(m, v22, v50, v58, v86, v87)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v88 != int32(1))
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
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int64
	_ = v108
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
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
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
	return v144
L2:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	v108 = F_strtox_2(m, v103, v7+int32(-12), int32(10), int64(4294967295))
	mBase = m.M
	v109 = base.I32_wrap_i64(v108)
	goto L30
L3:
	;
	v82 = int32(0)
	v83 = F_errsave_start(m, v11)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L19
	} else {
		goto L25
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
	v41 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[166])) = v41
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
	v50 = F_strtox_2(m, v45, v7+int32(-12), int32(10), int64(4294967295))
	mBase = m.M
	v51 = base.I32_wrap_i64(v50)
	goto L14
L6:
	;
	v37 = int32(1)
	if v36 <= v37 {
		v13 = v13 + v37
		v15 = v36
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v31 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7+int32(-8)+v15<<(uint(int32(2))%32)))) = v13 + v31
	v36 = v15 + v31
	goto L6
L8:
	;
	if v19 != int32(40) {
		v36 = v15
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
	if v15 != 0 {
		v36 = v15
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L7
L13:
	;
	goto L5
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v53 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v57 == int32(44) {
		goto L2
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v60 = F_errsave_start(m, v11)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	return int32(0)
L20:
	;
	if v60 == int32(0) {
		v144 = v41
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(412423)
	F_errmsg(m, int32(683256), v7+int32(-48))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	F_errsave_finish(m, v11, int32(476320), int32(81), int32(263730))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v144 = v41
	goto L1
L25:
	;
	if v83 == int32(0) {
		v144 = v82
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = int32(412423)
	F_errmsg(m, int32(683256), v7+int32(-32))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	F_errsave_finish(m, v11, int32(476320), int32(73), int32(263730))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	v144 = v82
	goto L1
L30:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v111 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v137 = F_palloc(m, int32(6))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L19
	} else {
		goto L41
	}
L32:
	;
	v118 = F_errsave_start(m, v11)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L19
	} else {
		goto L36
	}
L33:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v9)+52))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v113 != int32(41) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	if base.Ui32(v109) < base.Ui32(int32(65536)) {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	if v118 == int32(0) {
		v144 = v41
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L19
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(412423)
	F_errmsg(m, int32(683256), v9)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L19
	} else {
		goto L39
	}
L39:
	;
	F_errsave_finish(m, v11, int32(476320), int32(104), int32(263730))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	v144 = v41
	goto L1
L41:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v137)+4)) = uint16(v109)
	*(*uint16)(unsafe.Add(mBase, uint32(v137)+2)) = uint16(v51)
	v142 = int32(base.Ui32(v51) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v137))) = uint16(v142)
	v144 = v137
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
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
	return v60
L8:
	;
	v60 = int32(0)
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
	v48 = v16
	v52 = int32(0)
	goto L13
L13:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	v60 = v52 - v53
	goto L7
L14:
	;
	v48 = v43
	v52 = v45
	goto L13
L15:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v26 != v28 {
		v43 = v24
		v45 = v26
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v43 = v37
	v45 = int32(0)
	goto L14
L17:
	;
	if v28 == int32(0) {
		v43 = v24
		v45 = v26
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v33 = v25 - int32(1)
	if v33 == int32(0) {
		v43 = v24
		v45 = v26
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v36 = int32(1)
	v37 = v24 + v36
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v38 != 0 {
		v23 = v23 + v36
		v24 = v37
		v25 = v33
		v26 = v38
		goto L15
	} else {
		goto L20
	}
L20:
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
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
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
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
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
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v198 int32
	_ = v198
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
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
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
	v243 = m.ExcPending
	if v243 != 0 {
		goto L10
	} else {
		goto L63
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
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
	F_checkExprIsVarFree(m, l0, v167, v163)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
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
	v31 = F_coerce_to_specific_type(m, l0, v25, int32(20), int32(498566))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v163 = int32(498566)
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
	v148 = int32(0)
	if l1&int32(8) == v148 {
		v163 = v148
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
	v56 = int32(0)
	v60 = v56
	v64 = v56
	v66 = v7
	v69 = v7
	v70 = v7
	goto L26
L26:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v44+int32(48)+v64<<(uint(int32(2))%32))))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+56))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+22)))
	v78 = v76 + v77
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+16)))
	if v79 != int32(3) {
		v99 = v60
		v100 = v66
		v101 = v69
		v102 = v70
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_ReleaseCatCacheList(m, v44)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L10
	} else {
		goto L36
	}
L28:
	;
	v104 = v64 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	if v104 < v105 {
		v60 = v99
		v64 = v104
		v66 = v100
		v69 = v101
		v70 = v102
		goto L26
	} else {
		goto L35
	}
L29:
	;
	v82 = int32(1)
	v83 = v60 + v82
	v90 = F_can_coerce_type(m, v82, v17+int32(44), v78+int32(12), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	if v90 == int32(0) {
		v99 = v83
		v100 = v66
		v101 = v69
		v102 = v70
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v95 = v69 + int32(1)
	if v66 == v53 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v99 = v83
	v100 = v53
	v101 = v95
	v102 = v70
	goto L28
L33:
	;
	goto L34
L34:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v99 = v83
	v100 = v98
	v101 = v95
	v102 = v97
	goto L28
L35:
	;
	goto L27
L36:
	;
	if v99 == int32(0) {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	switch v101 {
	case 0:
		goto L40
	case 1:
		goto L38
	default:
		goto L39
	}
L38:
	;
	v143 = int32(517355)
	v145 = F_coerce_to_specific_type(m, l0, v36, v100, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L10
	} else {
		goto L50
	}
L39:
	;
	if v100 != v53 {
		goto L2
	} else {
		goto L49
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L10
	} else {
		goto L42
	}
L42:
	;
	v118 = F_format_type_be(m, l3)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	v121 = F_format_type_be(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v118
	F_errmsg(m, int32(178716), v17+int32(32))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	F_errhint(m, int32(592704), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	v134 = F_exprLocation(m, v36)
	mBase = m.M
	F_parser_errposition(m, l0, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(475292), int32(3788), int32(98276))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v102
	v163 = v143
	v167 = v145
	goto L6
L51:
	;
	v155 = F_transformExpr(m, l0, l5, int32(13))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L10
	} else {
		goto L52
	}
L52:
	;
	v159 = F_coerce_to_specific_type(m, l0, v155, int32(20), int32(499136))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	v163 = int32(499136)
	v167 = v159
	goto L6
L54:
	;
	v183 = v167
	goto L5
L55:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	v202 = F_format_type_be(m, l3)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L10
	} else {
		goto L57
	}
L57:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	v205 = F_format_type_be(m, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v202
	F_errmsg(m, int32(178611), v17+int32(16))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	F_errhint(m, int32(592909), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	v218 = F_exprLocation(m, v36)
	mBase = m.M
	F_parser_errposition(m, l0, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(475292), int32(3796), int32(98276))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
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
	v246 = m.ExcPending
	if v246 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	v247 = F_format_type_be(m, l3)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L10
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v247
	F_errmsg(m, int32(181420), v17)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	v253 = F_exprLocation(m, v36)
	mBase = m.M
	F_parser_errposition(m, l0, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L10
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(475292), int32(3780), int32(98276))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
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
	var v39 int32
	_ = v39
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
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v11 != v13 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		if v15 == int32(0) {
			v31 = int32(0)
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			F_expandRTE(m, v12, v32, l2, v31, l3, v31, v31, v9+int32(12))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v41 = F_palloc0(m, int32(24))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(36)
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
					if v47 != 0 {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
						v49 = v48
					} else {
						v49 = v31
					}
					v50 = int32(0)
					if v45 == v50 {
						v58 = v50
					} else {
						if v49 <= int32(0) {
							v58 = v50
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
							if v49 < v55 {
								*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v49
							} else {
							}
							v58 = v45
						}
					}
					*(*int64)(unsafe.Add(mBase, uint32(v41)+8)) = int64(8589936841)
					*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v58
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
					v64 = F_copyObjectImpl(m, v63)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v41)+20)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v64
						v70 = v41
						m.G0 = v9 + int32(16)
						return v70
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
						v70 = v20
						m.G0 = v9 + int32(16)
						return v70
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
					v70 = v20
					m.G0 = v9 + int32(16)
					return v70
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
	var v13 int32
	_ = v13
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
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
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	v3 = int32(0)
	if l1 == v3 {
		v97 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v97
L2:
	;
	if l0 == int32(0) {
		v97 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+120)))
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = int32(104)
	goto L9
L8:
	;
	v26 = int32(100)
	goto L9
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22+v26)))
	if v20 != v28 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L12
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v34 = F_lappend(m, int32(0), v22)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v38 < int32(2) {
		v97 = v34
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v42 = v18 + int32(4)
	if base.Ui32(v42) < base.Ui32(v18+v32<<(uint(int32(2))%32)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v48 = v42
	goto L18
L17:
	;
	v48 = int32(0)
	goto L18
L18:
	;
	v52 = v48
	v53 = int32(1)
	v54 = v34
	v55 = v20
	goto L19
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v53<<(uint(int32(2))%32))))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+120)))
	if v65 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v97 = v87
	goto L1
L21:
	;
	v66 = int32(104)
	goto L23
L22:
	;
	v66 = int32(100)
	goto L23
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v62+v66)))
	if v55 != v68 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v52 == int32(0) {
		v97 = v54
		goto L1
	} else {
		goto L27
	}
L25:
	;
	v84 = v52
	v85 = v55
	goto L26
L26:
	;
	if v85 != v68 {
		v97 = v54
		goto L1
	} else {
		goto L31
	}
L27:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v75 = v52 + int32(4)
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
	v85 = v73
	goto L26
L31:
	;
	v87 = F_lappend(m, v54, v62)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	v90 = v53 + int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v90 < v91 {
		v52 = v84
		v53 = v90
		v54 = v87
		v55 = v68
		goto L19
	} else {
		goto L33
	}
L33:
	;
	goto L20
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
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
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
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
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
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
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v322 int32
	_ = v322
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
	var v336 int32
	_ = v336
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
	v336 = m.ExcPending
	if v336 != 0 {
		goto L5
	} else {
		goto L142
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
	if v40 == v43 {
		v84 = v43
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
		goto L40
	} else {
		goto L41
	}
L25:
	;
	if v84 == int32(0) {
		goto L20
	} else {
		goto L39
	}
L26:
	;
	goto L25
L27:
	;
	if v42 == int32(0) {
		v84 = v43
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v52 < v53 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v55 = v52
	goto L31
L30:
	;
	v55 = v53
	goto L31
L31:
	;
	if v55 <= int32(1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v58 = int32(1)
	goto L34
L33:
	;
	v58 = v55
	goto L34
L34:
	;
	v59 = int32(8)
	v64 = int32(0)
	goto L35
L35:
	;
	v71 = v64 << (uint(int32(2)) % 32)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v42+v59+v71)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+(v40+v59))))
	v76 = v73 & v75
	v78 = base.B2i32(v76 != int32(0))
	if v76 != 0 {
		v84 = v78
		goto L26
	} else {
		goto L37
	}
L36:
	;
	v84 = v78
	goto L26
L37:
	;
	v80 = v64 + int32(1)
	if v80 != v58 {
		v64 = v80
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	goto L24
L40:
	;
	v91 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	v94 = v17 + int32(108)
	if l6 == v92 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v174 = v12
	goto L42
L42:
	;
	if l7 != 0 {
		goto L78
	} else {
		goto L79
	}
L43:
	;
	if v172 != 0 {
		goto L75
	} else {
		goto L76
	}
L44:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v160
	v172 = int32(1)
	goto L43
L45:
	;
	if l6 != 0 {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if l6 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(0)
	v172 = int32(1)
	goto L43
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(0)
	v172 = int32(1)
	goto L43
L50:
	;
	goto L51
L51:
	;
	if v92 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v112 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v112
	v172 = v112
	goto L43
L53:
	;
	goto L54
L54:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	v116 = int32(0)
	if v116 < v115 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v119 = v115
	goto L57
L56:
	;
	v119 = v116
	goto L57
L57:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v124 = v91
	goto L58
L58:
	;
	if v124 < v120 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v136 = v132 + v124<<(uint(int32(2))%32)
	goto L62
L61:
	;
	v136 = int32(0)
	goto L62
L62:
	;
	if v124 == v119 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v119
	v172 = base.B2i32(v136 == int32(0))
	goto L43
L64:
	;
	goto L65
L65:
	;
	v142 = base.B2i32(v136 == int32(0))
	if v136 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v124
	v172 = v142
	goto L43
L67:
	;
	goto L68
L68:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v149 = v146 + v124<<(uint(int32(2))%32)
	if v149 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v124
	v172 = v142
	goto L43
L70:
	;
	goto L71
L71:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	if v153 != v154 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v124
	v172 = int32(0)
	goto L43
L73:
	;
	v124 = v124 + int32(1)
	goto L58
L75:
	;
	v173 = v91
	goto L77
L76:
	;
	v173 = l6
	goto L77
L77:
	;
	v174 = v173
	goto L42
L78:
	;
	v175 = int32(0)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	if l7 == v176 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v231 = int32(0)
	goto L80
L80:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v17)+108))
	F_initial_cost_mergejoin(m, l0, v17+int32(8), l8, l5, l2, l3, v174, v231, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L5
	} else {
		goto L102
	}
L81:
	;
	if v229 != 0 {
		goto L99
	} else {
		goto L100
	}
L82:
	;
	v229 = int32(1)
	goto L81
L83:
	;
	goto L84
L84:
	;
	v185 = v175
	goto L86
L85:
	;
	v229 = v221
	goto L81
L86:
	;
	v189 = int32(0)
	if l7 == v189 {
		v199 = v189
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v221 = int32(0)
	goto L85
L88:
	;
	if v176 != 0 {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l7)+4))
	if v193 <= v185 {
		v199 = int32(0)
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l7)+12))
	v199 = v195 + v185<<(uint(int32(2))%32)
	goto L88
L91:
	;
	v205 = base.B2i32(v199 == int32(0))
	if v199 == int32(0) {
		v221 = v205
		goto L85
	} else {
		goto L96
	}
L92:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v185 < v200 {
		goto L91
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v229 = base.B2i32(v199 == int32(0))
	goto L81
L95:
	;
	goto L94
L96:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v176)+12))
	v211 = v208 + v185<<(uint(int32(2))%32)
	if v211 == int32(0) {
		v221 = v205
		goto L85
	} else {
		goto L97
	}
L97:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	if v216 == v217 {
		v185 = v185 + int32(1)
		goto L86
	} else {
		goto L98
	}
L98:
	;
	goto L87
L99:
	;
	v230 = v175
	goto L101
L100:
	;
	v230 = l7
	goto L101
L101:
	;
	v231 = v230
	goto L80
L102:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v238 = *(*float64)(unsafe.Add(mBase, uint32(v17)+16))
	v239 = *(*float64)(unsafe.Add(mBase, uint32(v17)+24))
	v240 = int32(0)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v244 == v240 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v322 == int32(0) {
		goto L20
	} else {
		goto L139
	}
L104:
	;
	v322 = int32(1)
	goto L103
L105:
	;
	goto L106
L106:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	if v248 <= int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v322 = int32(1)
	goto L103
L108:
	;
	goto L109
L109:
	;
	if v40 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v253 = int32(0)
	goto L112
L111:
	;
	v253 = l4
	goto L112
L112:
	;
	if v40 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v256 = int32(25)
	goto L115
L114:
	;
	v256 = int32(24)
	goto L115
L115:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v256))))
	v266 = v240
	goto L116
L116:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v269+v266<<(uint(int32(2))%32))))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+40))
	if v237 != v274 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	v322 = v310
	goto L103
L118:
	;
	if v258 != 0 {
		goto L126
	} else {
		goto L127
	}
L119:
	;
	if v274 <= v237 {
		goto L118
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v278 = *(*float64)(unsafe.Add(mBase, uint32(v273)+56))
	if base.F64_le(v239, base.F64_mul(v278, float64(1.01))) == int32(0) {
		goto L118
	} else {
		goto L123
	}
L122:
	;
	v322 = int32(1)
	goto L103
L123:
	;
	v322 = int32(1)
	goto L103
L124:
	;
	goto L117
L125:
	;
	v304 = int32(1)
	v306 = v266 + v304
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	if v306 < v307 {
		v266 = v306
		goto L116
	} else {
		goto L138
	}
L126:
	;
	v285 = *(*float64)(unsafe.Add(mBase, uint32(v273)+48))
	if base.F64_gt(v238, base.F64_mul(v285, float64(1.01))) == int32(0) {
		goto L125
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v291 = int32(0)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v273)+16))
	if v292 != 0 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	goto L128
L130:
	;
	v294 = v291
	goto L132
L131:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v273)+64))
	v294 = v293
	goto L132
L132:
	;
	v295 = F_compare_pathkeys(m, v253, v294)
	mBase = m.M
	if v295&int32(-3) != 0 {
		goto L125
	} else {
		goto L133
	}
L133:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v273)+16))
	if v298 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	v300 = v299
	goto L136
L135:
	;
	v300 = v291
	goto L136
L136:
	;
	v301 = F_bms_equal(m, v40, v300)
	mBase = m.M
	if v301 != 0 {
		v310 = v291
		goto L124
	} else {
		goto L137
	}
L137:
	;
	goto L125
L138:
	;
	v310 = v304
	goto L124
L139:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v17)+108))
	v329 = F_create_mergejoin_path(m, l0, l1, l8, v17+int32(8), l9, l2, l3, v327, l4, v40, l5, v174, v231, v328)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L5
	} else {
		goto L140
	}
L140:
	;
	F_add_path(m, l1, v329)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L5
	} else {
		goto L141
	}
L141:
	;
	goto L1
L142:
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
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
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 float64
	_ = v400
	var v401 float64
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v440 float64
	_ = v440
	var v447 float64
	_ = v447
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
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
	v494 = m.ExcPending
	if v494 != 0 {
		goto L11
	} else {
		goto L190
	}
L23:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v204 == int32(0) {
		goto L77
	} else {
		goto L78
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
	if v51 == v55 {
		v96 = v55
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if v96 != 0 {
		goto L23
	} else {
		goto L46
	}
L33:
	;
	goto L32
L34:
	;
	if v54 == int32(0) {
		v96 = v55
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v64 < v65 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v67 = v64
	goto L38
L37:
	;
	v67 = v65
	goto L38
L38:
	;
	if v67 <= int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v70 = int32(1)
	goto L41
L40:
	;
	v70 = v67
	goto L41
L41:
	;
	v71 = int32(8)
	v76 = int32(0)
	goto L42
L42:
	;
	v83 = v76 << (uint(int32(2)) % 32)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v54+v71+v83)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v83+(v51+v71))))
	v88 = v85 & v87
	v90 = base.B2i32(v88 != int32(0))
	if v88 != 0 {
		v96 = v90
		goto L33
	} else {
		goto L44
	}
L43:
	;
	v96 = v90
	goto L33
L44:
	;
	v92 = v76 + int32(1)
	if v92 != v70 {
		v76 = v92
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v100 = int32(0)
	if v20 == v100 {
		v141 = v100
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v141 == int32(0) {
		goto L22
	} else {
		goto L61
	}
L48:
	;
	goto L47
L49:
	;
	if v42 == int32(0) {
		v141 = v100
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v109 < v110 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v112 = v109
	goto L53
L52:
	;
	v112 = v110
	goto L53
L53:
	;
	if v112 <= int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v115 = int32(1)
	goto L56
L55:
	;
	v115 = v112
	goto L56
L56:
	;
	v116 = int32(8)
	v121 = int32(0)
	goto L57
L57:
	;
	v128 = v121 << (uint(int32(2)) % 32)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v42+v116+v128)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128+(v20+v116))))
	v133 = v130 & v132
	v135 = base.B2i32(v133 != int32(0))
	if v133 != 0 {
		v141 = v135
		goto L48
	} else {
		goto L59
	}
L58:
	;
	v141 = v135
	goto L48
L59:
	;
	v137 = v121 + int32(1)
	if v137 != v115 {
		v121 = v137
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	if v20 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if v201 == int32(0) {
		goto L22
	} else {
		goto L76
	}
L63:
	;
	v201 = int32(0)
	goto L62
L64:
	;
	goto L65
L65:
	;
	v154 = int32(1)
	if v42 == int32(0) {
		v192 = v154
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v201 = v192
	goto L62
L67:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v158 < v157 {
		v192 = v154
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v160 = int32(1)
	if v157 <= v160 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v163 = v160
	goto L71
L70:
	;
	v163 = v157
	goto L71
L71:
	;
	v164 = int32(8)
	v169 = int32(0)
	goto L72
L72:
	;
	v176 = v169 << (uint(int32(2)) % 32)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v20+v164+v176)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176+(v42+v164))))
	v183 = v178 & (v180 ^ int32(-1))
	v185 = base.B2i32(v183 != int32(0))
	if v183 != 0 {
		v192 = v185
		goto L66
	} else {
		goto L74
	}
L73:
	;
	v192 = v185
	goto L66
L74:
	;
	v187 = v169 + int32(1)
	if v187 != v163 {
		v169 = v187
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	goto L23
L77:
	;
	F_initial_cost_nestloop(m, l0, v16, l5, l2, l3, l6)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L11
	} else {
		goto L150
	}
L78:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+228))
	v210 = int32(0)
	if v207 == v210 {
		v251 = v210
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v251 == int32(0) {
		goto L77
	} else {
		goto L93
	}
L80:
	;
	goto L79
L81:
	;
	if v209 == int32(0) {
		v251 = v210
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	if v219 < v220 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v222 = v219
	goto L85
L84:
	;
	v222 = v220
	goto L85
L85:
	;
	if v222 <= int32(1) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v225 = int32(1)
	goto L88
L87:
	;
	v225 = v222
	goto L88
L88:
	;
	v226 = int32(8)
	v231 = int32(0)
	goto L89
L89:
	;
	v238 = v231 << (uint(int32(2)) % 32)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v209+v226+v238)))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v238+(v207+v226))))
	v243 = v240 & v242
	v245 = base.B2i32(v243 != int32(0))
	if v243 != 0 {
		v251 = v245
		goto L80
	} else {
		goto L91
	}
L90:
	;
	v251 = v245
	goto L80
L91:
	;
	v247 = v231 + int32(1)
	if v247 != v225 {
		v231 = v247
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v260 = int32(1)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v261 == int32(0) {
		v388 = v260
		goto L95
	} else {
		goto L96
	}
L94:
	;
	if v394 == int32(0) {
		goto L22
	} else {
		goto L149
	}
L95:
	;
	v394 = v388
	goto L94
L96:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v257)+228))
	v266 = F_bms_overlap(m, v264, v265)
	mBase = m.M
	if v266 == int32(0) {
		v388 = v260
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v269 = int32(0)
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	switch v270 - int32(279) {
	case 0, 1:
		goto L98
	default:
		v388 = v269
		goto L95
	case 3:
		goto L108
	case 4:
		goto L107
	case 5:
		goto L106
	case 9:
		goto L105
	case 10:
		goto L104
	case 11:
		goto L102
	case 14:
		goto L101
	case 15:
		goto L100
	case 17:
		goto L99
	case 19, 20, 21:
		goto L103
	}
L98:
	;
	v388 = int32(1)
	goto L95
L99:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v378 = F_path_is_reparameterizable_by_child(m, v377, v257)
	mBase = m.M
	if v378 == int32(0) {
		v388 = v269
		goto L95
	} else {
		goto L148
	}
L100:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v376 = F_path_is_reparameterizable_by_child(m, v375, v257)
	mBase = m.M
	if v376 != 0 {
		goto L98
	} else {
		goto L147
	}
L101:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v374 = F_path_is_reparameterizable_by_child(m, v373, v257)
	mBase = m.M
	if v374 != 0 {
		goto L98
	} else {
		goto L146
	}
L102:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v351 == int32(0) {
		goto L98
	} else {
		goto L138
	}
L103:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l3)+80))
	v346 = F_path_is_reparameterizable_by_child(m, v345, v257)
	mBase = m.M
	if v346 == int32(0) {
		v388 = v269
		goto L95
	} else {
		goto L136
	}
L104:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	if v323 == int32(0) {
		goto L98
	} else {
		goto L128
	}
L105:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v319 == int32(0) {
		goto L98
	} else {
		goto L126
	}
L106:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v297 == int32(0) {
		goto L98
	} else {
		goto L118
	}
L107:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	if v275 == int32(0) {
		goto L98
	} else {
		goto L110
	}
L108:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v274 = F_path_is_reparameterizable_by_child(m, v273, v257)
	mBase = m.M
	if v274 != 0 {
		goto L98
	} else {
		goto L109
	}
L109:
	;
	v388 = v269
	goto L95
L110:
	;
	v278 = int32(0)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v279 <= v278 {
		goto L98
	} else {
		goto L111
	}
L111:
	;
	v282 = v278
	goto L112
L112:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v275)+12))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v286+v282<<(uint(int32(2))%32))))
	v291 = F_path_is_reparameterizable_by_child(m, v290, v257)
	mBase = m.M
	if v291 != 0 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v394 = int32(0)
	goto L94
L114:
	;
	v293 = v282 + int32(1)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v293 < v294 {
		v282 = v293
		goto L112
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	goto L113
L117:
	;
	goto L98
L118:
	;
	v300 = int32(0)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	if v301 <= v300 {
		goto L98
	} else {
		goto L119
	}
L119:
	;
	v304 = v300
	goto L120
L120:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v297)+12))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v308+v304<<(uint(int32(2))%32))))
	v313 = F_path_is_reparameterizable_by_child(m, v312, v257)
	mBase = m.M
	if v313 != 0 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v394 = int32(0)
	goto L94
L122:
	;
	v315 = v304 + int32(1)
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	if v315 < v316 {
		v304 = v315
		goto L120
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	goto L121
L125:
	;
	goto L98
L126:
	;
	v322 = F_path_is_reparameterizable_by_child(m, v319, v257)
	mBase = m.M
	if v322 != 0 {
		goto L98
	} else {
		goto L127
	}
L127:
	;
	v388 = v269
	goto L95
L128:
	;
	v326 = int32(0)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	if v327 <= v326 {
		goto L98
	} else {
		goto L129
	}
L129:
	;
	v330 = v326
	goto L130
L130:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v323)+12))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v334+v330<<(uint(int32(2))%32))))
	v339 = F_path_is_reparameterizable_by_child(m, v338, v257)
	mBase = m.M
	if v339 != 0 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v394 = int32(0)
	goto L94
L132:
	;
	v341 = v330 + int32(1)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v323)+4))
	if v341 < v342 {
		v330 = v341
		goto L130
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	goto L131
L135:
	;
	goto L98
L136:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l3)+84))
	v350 = F_path_is_reparameterizable_by_child(m, v349, v257)
	mBase = m.M
	if v350 != 0 {
		goto L98
	} else {
		goto L137
	}
L137:
	;
	v388 = v269
	goto L95
L138:
	;
	v354 = int32(0)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v351)+4))
	if v355 <= v354 {
		goto L98
	} else {
		goto L139
	}
L139:
	;
	v358 = v354
	goto L140
L140:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v362+v358<<(uint(int32(2))%32))))
	v367 = F_path_is_reparameterizable_by_child(m, v366, v257)
	mBase = m.M
	if v367 != 0 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v394 = int32(0)
	goto L94
L142:
	;
	v369 = v358 + int32(1)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v351)+4))
	if v369 < v370 {
		v358 = v369
		goto L140
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	goto L141
L145:
	;
	goto L98
L146:
	;
	v388 = v269
	goto L95
L147:
	;
	v388 = v269
	goto L95
L148:
	;
	goto L98
L149:
	;
	goto L77
L150:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v400 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
	v401 = *(*float64)(unsafe.Add(mBase, uint32(v16)+16))
	v402 = int32(0)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v406 == v402 {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	if v484 == int32(0) {
		goto L22
	} else {
		goto L187
	}
L152:
	;
	v484 = int32(1)
	goto L151
L153:
	;
	goto L154
L154:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
	if v410 <= int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v484 = int32(1)
	goto L151
L156:
	;
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
	v415 = int32(0)
	goto L160
L159:
	;
	v415 = l4
	goto L160
L160:
	;
	if v51 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v418 = int32(25)
	goto L163
L162:
	;
	v418 = int32(24)
	goto L163
L163:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v418))))
	v428 = v402
	goto L164
L164:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v406)+12))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v431+v428<<(uint(int32(2))%32))))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v435)+40))
	if v399 != v436 {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v484 = v472
	goto L151
L166:
	;
	if v420 != 0 {
		goto L174
	} else {
		goto L175
	}
L167:
	;
	if v436 <= v399 {
		goto L166
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	v440 = *(*float64)(unsafe.Add(mBase, uint32(v435)+56))
	if base.F64_le(v401, base.F64_mul(v440, float64(1.01))) == int32(0) {
		goto L166
	} else {
		goto L171
	}
L170:
	;
	v484 = int32(1)
	goto L151
L171:
	;
	v484 = int32(1)
	goto L151
L172:
	;
	goto L165
L173:
	;
	v466 = int32(1)
	v468 = v428 + v466
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v406)+4))
	if v468 < v469 {
		v428 = v468
		goto L164
	} else {
		goto L186
	}
L174:
	;
	v447 = *(*float64)(unsafe.Add(mBase, uint32(v435)+48))
	if base.F64_gt(v400, base.F64_mul(v447, float64(1.01))) == int32(0) {
		goto L173
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v453 = int32(0)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v435)+16))
	if v454 != 0 {
		goto L178
	} else {
		goto L179
	}
L177:
	;
	goto L176
L178:
	;
	v456 = v453
	goto L180
L179:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v435)+64))
	v456 = v455
	goto L180
L180:
	;
	v457 = F_compare_pathkeys(m, v415, v456)
	mBase = m.M
	if v457&int32(-3) != 0 {
		goto L173
	} else {
		goto L181
	}
L181:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v435)+16))
	if v460 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)+4))
	v462 = v461
	goto L184
L183:
	;
	v462 = v453
	goto L184
L184:
	;
	v463 = F_bms_equal(m, v51, v462)
	mBase = m.M
	if v463 != 0 {
		v472 = v453
		goto L172
	} else {
		goto L185
	}
L185:
	;
	goto L173
L186:
	;
	v472 = v466
	goto L172
L187:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v488 = F_create_nestloop_path(m, l0, l1, l5, v16, l6, l2, l3, v487, l4, v51)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L11
	} else {
		goto L188
	}
L188:
	;
	F_add_path(m, l1, v488)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L11
	} else {
		goto L189
	}
L189:
	;
	goto L7
L190:
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
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
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v305 int32
	_ = v305
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	v2 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = F_pq_getmsgint(m, v20, int32(4))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L5
	} else {
		goto L113
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L5
	} else {
		goto L110
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L5
	} else {
		goto L107
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L5
	} else {
		goto L104
	}
L5:
	;
	return int32(0)
L6:
	;
	if base.Ui32(v22) < base.Ui32(int32(268435456)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v29 = v22 << (uint(int32(2)) % 32)
	v31 = v29 + int32(8)
	v33 = v31 << (uint(int32(1)) % 32)
	v34 = F_palloc0(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
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
	v405 = m.ExcPending
	if v405 != 0 {
		goto L5
	} else {
		goto L101
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v22
	if v22 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(32)
	return v34
L12:
	;
	goto L13
L13:
	;
	v44 = v33
	v45 = v34
	v46 = v2
	v53 = v2
	v54 = v2
	goto L14
L14:
	;
	v63 = F_pq_getmsgstring(m, v20)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = (v379 + v31) << (uint(int32(2)) % 32)
	if v260&int32(1) != 0 {
		goto L97
	} else {
		goto L98
	}
L16:
	;
	v66 = F_pq_getmsgint(m, v20, int32(2))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	if v63&int32(3) == int32(0) {
		v91 = v63
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if base.Ui32(int32(2048)) <= base.Ui32(v124) {
		goto L4
	} else {
		goto L35
	}
L19:
	;
	v124 = v116 - v63
	goto L18
L20:
	;
	v95 = v91
	goto L29
L21:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v75 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v124 = int32(0)
	goto L18
L23:
	;
	goto L24
L24:
	;
	v80 = v63
	goto L25
L25:
	;
	v84 = v80 + int32(1)
	if v84&int32(3) == int32(0) {
		v91 = v84
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v116 = v84
	goto L19
L27:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v89 != 0 {
		v80 = v84
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v104 = int32(-2139062144)
	if (int32(16843008)-v101|v101)&v104 == v104 {
		v95 = v95 + int32(4)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v110 = v95
	goto L32
L31:
	;
	goto L30
L32:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v114 != 0 {
		v110 = v110 + int32(1)
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v116 = v110
	goto L19
L34:
	;
	goto L33
L35:
	;
	if int32(1048576) <= v46 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v130 = v66 & int32(65535)
	if base.Ui32(int32(256)) < base.Ui32(v130) {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v133 = v46 + v124
	v134 = int32(1)
	v137 = (v133 + v134) & int32(-2)
	v139 = v130 << (uint(v134) % 32)
	v141 = v137 + (v29 + int32(10) + v139)
	if base.Ui32(v44) <= base.Ui32(v141) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v143 = v44
	v144 = v45
	goto L41
L39:
	;
	v167 = v44
	v168 = v45
	goto L40
L40:
	;
	v187 = v168 + int32(8)
	v188 = int32(2)
	v190 = v187 + v53<<(uint(v188)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v124<<(uint(int32(1))%32) | base.B2i32(v130 != int32(0)) | v46<<(uint(int32(12))%32)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	if v124 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	v163 = v143 << (uint(int32(1)) % 32)
	v164 = F_repalloc(m, v144, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L5
	} else {
		goto L43
	}
L42:
	;
	v167 = v163
	v168 = v164
	goto L40
L43:
	;
	if base.Ui32(v163) <= base.Ui32(v141) {
		v143 = v163
		v144 = v164
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	if v53 != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v205 = F__emscripten_memcpy_bulkmem(m, v187+v200<<(uint(v188)%32)+v46, v63, v124)
	mBase = m.M
	goto L48
L47:
	;
	goto L48
L48:
	;
	goto L45
L49:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v210 = v187 + v207<<(uint(int32(2))%32)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v212 = int32(12)
	v215 = int32(1)
	v217 = int32(2047)
	v218 = int32(base.Ui32(v211)>>(uint(v215)%32)) & v217
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v190-int32(4))))
	v228 = int32(base.Ui32(v221)>>(uint(v215)%32)) & v217
	if v218 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v260 = v54
	goto L51
L51:
	;
	if v130 != 0 {
		goto L80
	} else {
		goto L81
	}
L52:
	;
	v260 = base.B2i32(v254 <= int32(0)) | v54
	goto L51
L53:
	;
	goto L57
L54:
	;
	goto L55
L55:
	;
	if v228 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L57:
	;
	goto L58
L58:
	;
	v234 = int32(0)
	if v234 < v228 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v237 = int32(-1)
	goto L61
L60:
	;
	v237 = v234
	goto L61
L61:
	;
	v254 = v237
	goto L52
L62:
	;
	v254 = base.B2i32(int32(0) < v218)
	goto L52
L63:
	;
	goto L64
L64:
	;
	if base.Ui32(v218) < base.Ui32(v228) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v243 = v218
	goto L67
L66:
	;
	v243 = v228
	goto L67
L67:
	;
	v244 = F_memcmp(m, v210+int32(base.Ui32(v211)>>(uint(v212)%32)), v210+int32(base.Ui32(v221)>>(uint(v212)%32)), v243)
	mBase = m.M
	goto L70
L68:
	;
	v254 = v252
	goto L52
L70:
	;
	goto L71
L71:
	;
	if v244 != 0 {
		v252 = v244
		goto L68
	} else {
		goto L73
	}
L73:
	;
	if v218 == v228 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v254 = int32(0)
	goto L52
L75:
	;
	goto L76
L76:
	;
	if v218 < v228 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v251 = int32(-1)
	goto L79
L78:
	;
	v251 = int32(1)
	goto L79
L79:
	;
	v252 = v251
	goto L68
L80:
	;
	if v133 == v137 {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v379 = v133
	goto L82
L82:
	;
	v381 = v53 + int32(1)
	if v381 != v22 {
		v44 = v167
		v45 = v168
		v46 = v379
		v53 = v381
		v54 = v260
		goto L14
	} else {
		goto L96
	}
L83:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v271 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v187+v270<<(uint(v271)%32)+v269))) = uint16(v66)
	v276 = int32(1)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v293 = v187 + v277<<(uint(v271)%32) + (int32(base.Ui32(v281)>>(uint(int32(12))%32))+int32(base.Ui32(v281)>>(uint(v276)%32))&int32(2047)+v276)&int32(4194302)
	v295 = F_pq_getmsgint(m, v20, v271)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L5
	} else {
		goto L87
	}
L84:
	;
	v269 = v133
	goto L83
L85:
	;
	goto L86
L86:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v267 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v187+v262<<(uint(int32(2))%32)+v133))) = uint8(v267)
	v269 = v137
	goto L83
L87:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v293)+2)) = uint16(v295)
	if v130 != int32(1) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v305 = v276
	goto L91
L89:
	;
	goto L90
L90:
	;
	v379 = v269 + v139 + int32(2)
	goto L82
L91:
	;
	v322 = v305 << (uint(int32(1)) % 32)
	v325 = F_pq_getmsgint(m, v20, int32(2))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L5
	} else {
		goto L93
	}
L92:
	;
	goto L90
L93:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v293+int32(2)+v322))) = uint16(v325)
	v328 = int32(16383)
	v331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v322+v293))))
	if base.Ui32(v325&v328) <= base.Ui32(v331&v328) {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v336 = v305 + int32(1)
	if v130 != v336 {
		v305 = v336
		goto L91
	} else {
		goto L95
	}
L95:
	;
	goto L92
L96:
	;
	goto L15
L97:
	;
	v390 = v168 + int32(8)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	F_qsort_arg(m, v390, v391, int32(4), int32(1534), v390+v391<<(uint(int32(2))%32))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L5
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	return v168
L100:
	;
	goto L99
L101:
	;
	F_errmsg_internal(m, int32(197865), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(471930), int32(462), int32(33496))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	F_errmsg_internal(m, int32(311592), int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(471930), int32(484), int32(33496))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	F_errmsg_internal(m, int32(440934), int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(471930), int32(487), int32(33496))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errmsg_internal(m, int32(129496), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(471930), int32(490), int32(33496))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errmsg_internal(m, int32(430093), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L5
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(471930), int32(541), int32(33496))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L5
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tuplehash_lookup_hash_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v90 int32
	_ = v90
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v15 = v14 & l1
	v18 = v13 + v15*int32(12)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v19 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v90
L2:
	;
	v22 = v14
	v23 = v13
	v24 = v15
	v25 = v18
	goto L5
L3:
	;
	goto L4
L4:
	;
	v90 = int32(0)
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
	v67 = v22
	v68 = v23
	goto L9
L9:
	;
	v71 = (v24 + int32(1)) & v67
	v74 = v71*int32(12) + v68
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v75 != 0 {
		v22 = v67
		v23 = v68
		v24 = v71
		v25 = v74
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
	v50 = int32(4442576)
	v51 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v53
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	v58 = m.T0[v57].(func(*base.Module, int32, int32, int32) int32)(m, v44, v33, v11+int32(15))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L16
	}
L15:
	;
	v90 = v25
	goto L1
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v51
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
		v90 = v25
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v67 = v66
	v68 = v65
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
