package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitPostmasterChildSlots(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[0])) = int32(32)
	v13 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[1])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[2])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[3])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[4])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[5])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[6])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[7])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[8])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[9])) = v13
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[10])) = v13
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[12])) = v44
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[13]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[14])) = v48
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[15]))
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[16]))
	v57 = (v52 + v54) << (uint(v13) % 32)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[17])) = v57
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[18]))
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[19]))
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[20]))
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[21]))
	v75 = v61 + (v63 + (v48 + (v44 + (v65 + (v67 + v57))))) + int32(42)
	*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[22])) = v75
	v79 = F_palloc(m, v75*int32(28))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		return
	} else {
		v84 = int32(0)
		v86 = int32(0)
		for {
			v91 = v86 << (uint(int32(4)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[23]))) = v84 + int32(1)
			v96 = v91 + int32(_a_F_InitPostmasterChildSlots_0)
			*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[24]))) = v96
			*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[25]))) = v96
			v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[21])))
			if int32(0) < v99 {
				v107 = v84
				v110 = int32(0)
				for {
					v115 = v79 + v107*int32(28)
					*(*int64)(unsafe.Add(mBase, uint32(v115)+8)) = int64(0)
					v119 = v107 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v115)+4)) = v119
					v121 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v115))) = v121
					*(*uint8)(unsafe.Add(mBase, uint32(v115)+16)) = uint8(v121)
					v125 = *(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[24])))
					if v125 == v121 {
						*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[24]))) = v96
						*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[25]))) = v91 + int32(_a_F_InitPostmasterChildSlots_0)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v115)+24)) = v96
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[25])))
					*(*int32)(unsafe.Add(mBase, uint32(v115)+20)) = v133
					v136 = v115 + int32(20)
					*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = v136
					*(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[25]))) = v136
					v140 = v110 + int32(1)
					v141 = *(*int32)(unsafe.Add(mBase, uint32(v91)+uint32(_c_F_InitPostmasterChildSlots[21])))
					if v140 < v141 {
						v107 = v119
						v110 = v140
						continue
					} else {
						break
					}
					break
				}
				v145 = v119
			} else {
				v145 = v84
			}
			v152 = v86 + int32(1)
			if v152 != int32(18) {
				v84 = v145
				v86 = v152
				continue
			} else {
				break
			}
			break
		}
		v156 = int32(_a_F_InitPostmasterChildSlots_1)
		*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[26])) = v156
		*(*int32)(unsafe.Add(mBase, _c_F_InitPostmasterChildSlots[27])) = v156
		return
	}
}
func F_ReleasePostmasterChildSlot(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if v15 == int32(2) {
			if v18 != 0 {
				F_errmsg_internal(m, int32(_a_F_ReleasePostmasterChildSlot_0), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(257), int32(_a_F_ReleasePostmasterChildSlot_2))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_pfree(m, l0)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v93 = int32(1)
							m.G0 = v8 + int32(32)
							return v93
						}
					}
				}
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v93 = int32(1)
					m.G0 = v8 + int32(32)
					return v93
				}
			}
		} else {
			if v18 != 0 {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v36
				F_errmsg_internal(m, int32(_a_F_ReleasePostmasterChildSlot_3), v8+int32(16))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(265), int32(_a_F_ReleasePostmasterChildSlot_2))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v50 == int32(6) {
							v57 = int32(_a_F_ReleasePostmasterChildSlot_4)
						} else {
							v57 = v50<<(uint(int32(4))%32) + int32(_a_F_ReleasePostmasterChildSlot_5)
						}
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
						if v48 < v58 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v102
								F_errmsg_internal(m, int32(_a_F_ReleasePostmasterChildSlot_6), v8)
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(278), int32(_a_F_ReleasePostmasterChildSlot_2))
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
							if v60+v58 <= v48 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v102
									F_errmsg_internal(m, int32(_a_F_ReleasePostmasterChildSlot_6), v8)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(278), int32(_a_F_ReleasePostmasterChildSlot_2))
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v64 = v57 + int32(8)
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
								if v65 == int32(0) {
									*(*int32)(unsafe.Add(mBase, uint32(v64))) = v64
									v69 = v64
								} else {
									v69 = v65
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v64
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v69
								v73 = l0 + int32(20)
								*(*int32)(unsafe.Add(mBase, uint32(v69))) = v73
								*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v73
								v77 = *(*int32)(unsafe.Add(mBase, _c_F_ReleasePostmasterChildSlot[0]))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v83 = v77 + v78<<(uint(int32(2))%32) + int32(48)
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
								*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
								v93 = base.B2i32(v84 == int32(1))
								m.G0 = v8 + int32(32)
								return v93
							}
						}
					}
				}
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v50 == int32(6) {
					v57 = int32(_a_F_ReleasePostmasterChildSlot_4)
				} else {
					v57 = v50<<(uint(int32(4))%32) + int32(_a_F_ReleasePostmasterChildSlot_5)
				}
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
				if v48 < v58 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v102
						F_errmsg_internal(m, int32(_a_F_ReleasePostmasterChildSlot_6), v8)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(278), int32(_a_F_ReleasePostmasterChildSlot_2))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
					if v60+v58 <= v48 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v102
							F_errmsg_internal(m, int32(_a_F_ReleasePostmasterChildSlot_6), v8)
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(278), int32(_a_F_ReleasePostmasterChildSlot_2))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v64 = v57 + int32(8)
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
						if v65 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v64))) = v64
							v69 = v64
						} else {
							v69 = v65
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v64
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v69
						v73 = l0 + int32(20)
						*(*int32)(unsafe.Add(mBase, uint32(v69))) = v73
						*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v73
						v77 = *(*int32)(unsafe.Add(mBase, _c_F_ReleasePostmasterChildSlot[0]))
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v83 = v77 + v78<<(uint(int32(2))%32) + int32(48)
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
						*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
						v93 = base.B2i32(v84 == int32(1))
						m.G0 = v8 + int32(32)
						return v93
					}
				}
			}
		}
	}
}
func F_postmaster_child_launch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int64
	_ = v201
	var v204 int64
	_ = v204
	var v207 int64
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int64
	_ = v231
	var v234 int64
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	v10 = m.G0
	v12 = v10 - int32(2128)
	m.G0 = v12
	switch l0 - int32(1) {
	case 0, 5:
		goto L2
	default:
		goto L1
	}
L1:
	;
	v36 = l3 + int32(3312)
	v37 = F_palloc0(m, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v19 = m.G0
	v20 = int32(16)
	v21 = v19 - v20
	m.G0 = v21
	F_gettimeofday(m, v21)
	mBase = m.M
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v25 = int64(*(*int32)(unsafe.Add(mBase, uint32(v21)+8)))
	m.G0 = v21 + v20
	goto L3
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v25 + v24*int64(1000000) - int64(946684800000000)
	goto L1
L4:
	;
	return int32(0)
L5:
	;
	v42 = v37 + int32(3168)
	if l4 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+3304)) = v52
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[0]))
	goto L13
L7:
	;
	base.MemoryFill(m, v42, int32(0), int32(136))
	v52 = int32(-1)
	goto L6
L8:
	;
	goto L9
L9:
	;
	base.MemoryCopy(m, v42, l4, int32(136))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v52 = v51
	goto L6
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+3160)) = l1
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1024)) = v177
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1028)) = v180
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1032)) = v183
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1036)) = v186
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1040)) = v189
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1044)) = v192
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1048)) = v195
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1052)) = v198
	v201 = *(*int64)(unsafe.Add(mBase, _c_F_postmaster_child_launch[9]))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+1056)) = v201
	v204 = *(*int64)(unsafe.Add(mBase, _c_F_postmaster_child_launch[10]))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+1064)) = v204
	v207 = *(*int64)(unsafe.Add(mBase, _c_F_postmaster_child_launch[11]))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+1072)) = v207
	v210 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postmaster_child_launch[12])))
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+1080)) = uint8(v210)
	v213 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postmaster_child_launch[13])))
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+1081)) = uint8(v213)
	v216 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postmaster_child_launch[14])))
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+1082)) = uint8(v216)
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[15]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1084)) = v219
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[16]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1088)) = v222
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[17]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1092)) = v225
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[18]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+3164)) = v228
	v231 = *(*int64)(unsafe.Add(mBase, _c_F_postmaster_child_launch[19]))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+1096)) = v231
	v234 = *(*int64)(unsafe.Add(mBase, _c_F_postmaster_child_launch[20]))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+1104)) = v234
	v237 = v37 + int32(1112)
	v238 = int32(_a_F_postmaster_child_launch_0)
	goto L44
L11:
	;
	v172 = F_strlen(m, v161)
	mBase = m.M
	goto L10
L13:
	;
	goto L14
L14:
	;
	v62 = int32(1023)
	if (v37^v55)&int32(3) != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v165 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v162))) = uint8(v165)
	goto L11
L16:
	;
	v146 = v141
	v147 = v142
	v148 = v143
	goto L37
L17:
	;
	if v136 == int32(0) {
		v161 = v134
		v162 = v135
		goto L15
	} else {
		goto L36
	}
L18:
	;
	v134 = v55
	v135 = v37
	v136 = v62
	goto L17
L19:
	;
	goto L20
L20:
	;
	v66 = int32(0)
	if base.B2i32(v55&int32(3) == v66)|int32(0) == v66 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v102 == int32(0) {
		v161 = v99
		v162 = v100
		goto L15
	} else {
		goto L30
	}
L22:
	;
	v78 = v55
	v79 = v37
	v80 = v62
	goto L25
L23:
	;
	goto L24
L24:
	;
	v99 = v55
	v100 = v37
	v101 = v62
	v102 = int32(1)
	goto L21
L25:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v82)
	if v82 == int32(0) {
		v141 = v78
		v142 = v79
		v143 = v80
		goto L16
	} else {
		goto L27
	}
L26:
	;
	v99 = v93
	v100 = v87
	v101 = v89
	v102 = v91
	goto L21
L27:
	;
	v86 = int32(1)
	v87 = v79 + v86
	v89 = v80 - v86
	v90 = int32(0)
	v91 = base.B2i32(v89 != v90)
	v93 = v78 + v86
	if v93&int32(3) == v90 {
		v99 = v93
		v100 = v87
		v101 = v89
		v102 = v91
		goto L21
	} else {
		goto L28
	}
L28:
	;
	if v89 != 0 {
		v78 = v93
		v79 = v87
		v80 = v89
		goto L25
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if base.B2i32(v105 == int32(0))|base.B2i32(base.Ui32(v101) < base.Ui32(int32(4))) != 0 {
		v134 = v99
		v135 = v100
		v136 = v101
		goto L17
	} else {
		goto L31
	}
L31:
	;
	v112 = v99
	v113 = v100
	v114 = v101
	goto L32
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v120 = int32(-2139062144)
	if (int32(16843008)-v117|v117)&v120 != v120 {
		v141 = v112
		v142 = v113
		v143 = v114
		goto L16
	} else {
		goto L34
	}
L33:
	;
	v134 = v128
	v135 = v126
	v136 = v130
	goto L17
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v117
	v125 = int32(4)
	v126 = v113 + v125
	v128 = v112 + v125
	v130 = v114 - v125
	if base.Ui32(int32(3)) < base.Ui32(v130) {
		v112 = v128
		v113 = v126
		v114 = v130
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v141 = v134
	v142 = v135
	v143 = v136
	goto L16
L37:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	*(*uint8)(unsafe.Add(mBase, uint32(v147))) = uint8(v150)
	if v150 == int32(0) {
		v161 = v146
		v162 = v147
		goto L15
	} else {
		goto L39
	}
L38:
	;
	v161 = v157
	v162 = v155
	goto L15
L39:
	;
	v154 = int32(1)
	v155 = v147 + v154
	v157 = v146 + v154
	v159 = v148 - v154
	if v159 != 0 {
		v146 = v157
		v147 = v155
		v148 = v159
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v359 = v37 + int32(2136)
	v360 = int32(_a_F_postmaster_child_launch_1)
	goto L75
L42:
	;
	v355 = F_strlen(m, v344)
	mBase = m.M
	goto L41
L44:
	;
	goto L45
L45:
	;
	v245 = int32(1023)
	if (v237^v238)&int32(3) != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v348 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v345))) = uint8(v348)
	goto L42
L47:
	;
	v329 = v324
	v330 = v325
	v331 = v326
	goto L68
L48:
	;
	if v319 == int32(0) {
		v344 = v317
		v345 = v318
		goto L46
	} else {
		goto L67
	}
L49:
	;
	v317 = v238
	v318 = v237
	v319 = v245
	goto L48
L50:
	;
	goto L51
L51:
	;
	goto L54
L52:
	;
	goto L61
L54:
	;
	goto L55
L55:
	;
	goto L52
L61:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postmaster_child_launch[21])))
	if base.B2i32(v288 == int32(0))|int32(0) != 0 {
		v317 = v238
		v318 = v237
		v319 = v245
		goto L48
	} else {
		goto L62
	}
L62:
	;
	v295 = v238
	v296 = v237
	v297 = v245
	goto L63
L63:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	v303 = int32(-2139062144)
	if (int32(16843008)-v300|v300)&v303 != v303 {
		v324 = v295
		v325 = v296
		v326 = v297
		goto L47
	} else {
		goto L65
	}
L64:
	;
	v317 = v311
	v318 = v309
	v319 = v313
	goto L48
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v296))) = v300
	v308 = int32(4)
	v309 = v296 + v308
	v311 = v295 + v308
	v313 = v297 - v308
	if base.Ui32(int32(3)) < base.Ui32(v313) {
		v295 = v311
		v296 = v309
		v297 = v313
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v324 = v317
	v325 = v318
	v326 = v319
	goto L47
L68:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329))))
	*(*uint8)(unsafe.Add(mBase, uint32(v330))) = uint8(v333)
	if v333 == int32(0) {
		v344 = v329
		v345 = v330
		goto L46
	} else {
		goto L70
	}
L69:
	;
	v344 = v340
	v345 = v338
	goto L46
L70:
	;
	v337 = int32(1)
	v338 = v330 + v337
	v340 = v329 + v337
	v342 = v331 - v337
	if v342 != 0 {
		v329 = v340
		v330 = v338
		v331 = v342
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+3308)) = l3
	if l3 != 0 {
		goto L103
	} else {
		goto L104
	}
L73:
	;
	v477 = F_strlen(m, v466)
	mBase = m.M
	goto L72
L75:
	;
	goto L76
L76:
	;
	v367 = int32(1023)
	if (v359^v360)&int32(3) != 0 {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	v470 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v467))) = uint8(v470)
	goto L73
L78:
	;
	v451 = v446
	v452 = v447
	v453 = v448
	goto L99
L79:
	;
	if v441 == int32(0) {
		v466 = v439
		v467 = v440
		goto L77
	} else {
		goto L98
	}
L80:
	;
	v439 = v360
	v440 = v359
	v441 = v367
	goto L79
L81:
	;
	goto L82
L82:
	;
	goto L85
L83:
	;
	goto L92
L85:
	;
	goto L86
L86:
	;
	goto L83
L92:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postmaster_child_launch[22])))
	if base.B2i32(v410 == int32(0))|int32(0) != 0 {
		v439 = v360
		v440 = v359
		v441 = v367
		goto L79
	} else {
		goto L93
	}
L93:
	;
	v417 = v360
	v418 = v359
	v419 = v367
	goto L94
L94:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	v425 = int32(-2139062144)
	if (int32(16843008)-v422|v422)&v425 != v425 {
		v446 = v417
		v447 = v418
		v448 = v419
		goto L78
	} else {
		goto L96
	}
L95:
	;
	v439 = v433
	v440 = v431
	v441 = v435
	goto L79
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v418))) = v422
	v430 = int32(4)
	v431 = v418 + v430
	v433 = v417 + v430
	v435 = v419 - v430
	if base.Ui32(int32(3)) < base.Ui32(v435) {
		v417 = v433
		v418 = v431
		v419 = v435
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v446 = v439
	v447 = v440
	v448 = v441
	goto L78
L99:
	;
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v451))))
	*(*uint8)(unsafe.Add(mBase, uint32(v452))) = uint8(v455)
	if v455 == int32(0) {
		v466 = v451
		v467 = v452
		goto L77
	} else {
		goto L101
	}
L100:
	;
	v466 = v462
	v467 = v460
	goto L77
L101:
	;
	v459 = int32(1)
	v460 = v452 + v459
	v462 = v451 + v459
	v464 = v453 - v459
	if v464 != 0 {
		v451 = v462
		v452 = v460
		v453 = v464
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	base.MemoryCopy(m, v37+int32(3312), l2, l3)
	goto L105
L104:
	;
	goto L105
L105:
	;
	v484 = int32(_a_F_postmaster_child_launch_2)
	v486 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[23]))
	v488 = v486 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[23])) = v488
	v490 = int32(_a_F_postmaster_child_launch_3)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v490
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[24]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v488
	v499 = v12 + int32(1104)
	v504 = F_pg_snprintf(m, v499, int32(1024), int32(_a_F_postmaster_child_launch_4), v12-int32(-64))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	v507 = F_AllocateFile(m, v499, int32(_a_F_postmaster_child_launch_5))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L4
	} else {
		goto L110
	}
L107:
	;
	m.G0 = v12 + int32(2128)
	return v609
L108:
	;
	F_pfree(m, v37)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L4
	} else {
		goto L146
	}
L109:
	;
	v535 = F_fwrite(m, v37, v36, int32(1), v533)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L4
	} else {
		goto L120
	}
L110:
	;
	if v507 != 0 {
		v533 = v507
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v511 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[25]))
	v512 = F_mkdir(m, int32(_a_F_postmaster_child_launch_3), v511)
	mBase = m.M
	goto L112
L112:
	;
	v514 = F_AllocateFile(m, v499, int32(_a_F_postmaster_child_launch_5))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	if v514 != 0 {
		v533 = v514
		goto L109
	} else {
		goto L114
	}
L114:
	;
	v518 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	if v518 == int32(0) {
		goto L108
	} else {
		goto L116
	}
L116:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v499
	F_errmsg(m, int32(_a_F_postmaster_child_launch_6), v12)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_postmaster_child_launch_7), int32(333), int32(_a_F_postmaster_child_launch_8))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	goto L108
L120:
	;
	if v535 != int32(1) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v541 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L4
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	F_pfree(m, v37)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L4
	} else {
		goto L132
	}
L124:
	;
	if v541 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L4
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v558 = F_FreeFile(m, v533)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L4
	} else {
		goto L131
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v12 + int32(1104)
	F_errmsg(m, int32(_a_F_postmaster_child_launch_9), v12+int32(48))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_postmaster_child_launch_7), int32(343), int32(_a_F_postmaster_child_launch_8))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	goto L127
L131:
	;
	goto L108
L132:
	;
	v562 = F_FreeFile(m, v533)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	if v562 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v564 = int32(-1)
	v567 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L4
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l0
	v588 = v12 + int32(80)
	v593 = F_pg_snprintf(m, v588, int32(1024), int32(_a_F_postmaster_child_launch_10), v12+int32(16))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L4
	} else {
		goto L142
	}
L137:
	;
	if v567 == int32(0) {
		v609 = v564
		goto L107
	} else {
		goto L138
	}
L138:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v12 + int32(1104)
	F_errmsg(m, int32(_a_F_postmaster_child_launch_9), v12+int32(32))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_postmaster_child_launch_7), int32(355), int32(_a_F_postmaster_child_launch_8))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	v609 = v564
	goto L107
L142:
	;
	v598 = m.Env.Pgmem_spawn(m, v588, v12+int32(1104))
	mBase = m.M
	if int32(0) <= v598 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v609 = v598
	goto L107
L144:
	;
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[26])) = int32(0) - v598
	v609 = int32(-1)
	goto L107
L146:
	;
	v609 = int32(-1)
	goto L107
}
