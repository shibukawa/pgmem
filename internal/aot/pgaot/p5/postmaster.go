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
					F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(241), int32(_a_F_ReleasePostmasterChildSlot_2))
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
					F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(249), int32(_a_F_ReleasePostmasterChildSlot_2))
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
									F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(262), int32(_a_F_ReleasePostmasterChildSlot_2))
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
										F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(262), int32(_a_F_ReleasePostmasterChildSlot_2))
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
								v83 = v77 + v78<<(uint(int32(2))%32) + int32(44)
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
							F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(262), int32(_a_F_ReleasePostmasterChildSlot_2))
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
								F_errfinish(m, int32(_a_F_ReleasePostmasterChildSlot_1), int32(262), int32(_a_F_ReleasePostmasterChildSlot_2))
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
						v83 = v77 + v78<<(uint(int32(2))%32) + int32(44)
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
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
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
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int64
	_ = v219
	var v222 int64
	_ = v222
	var v225 int64
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int64
	_ = v246
	var v249 int64
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v492 int32
	_ = v492
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
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
	v36 = l3 + int32(3332)
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
	v42 = v37 + int32(3188)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0*int32(12))+uint32(_c_F_postmaster_child_launch[0])))
	if l4 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+3324)) = v55
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[1]))
	goto L13
L7:
	;
	base.MemoryFill(m, v42, int32(0), int32(136))
	v55 = int32(-1)
	goto L6
L8:
	;
	goto L9
L9:
	;
	base.MemoryCopy(m, v42, l4, int32(136))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v55 = v54
	goto L6
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+3184)) = l1
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1024)) = v180
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1028)) = v183
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1032)) = v186
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1036)) = v189
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1040)) = v192
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1044)) = v195
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1048)) = v198
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1052)) = v201
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1056)) = v204
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1060)) = v207
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1064)) = v210
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1068)) = v213
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1072)) = v216
	v219 = *(*int64)(unsafe.Add(mBase, _c_F_postmaster_child_launch[15]))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+1080)) = v219
	v222 = *(*int64)(unsafe.Add(mBase, _c_F_postmaster_child_launch[16]))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+1088)) = v222
	v225 = *(*int64)(unsafe.Add(mBase, _c_F_postmaster_child_launch[17]))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+1096)) = v225
	v228 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postmaster_child_launch[18])))
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+1104)) = uint8(v228)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postmaster_child_launch[19])))
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+1105)) = uint8(v231)
	v234 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postmaster_child_launch[20])))
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+1106)) = uint8(v234)
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1108)) = v237
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1112)) = v240
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[23]))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+1116)) = v243
	v246 = *(*int64)(unsafe.Add(mBase, _c_F_postmaster_child_launch[24]))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+1120)) = v246
	v249 = *(*int64)(unsafe.Add(mBase, _c_F_postmaster_child_launch[25]))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+1128)) = v249
	v252 = v37 + int32(1136)
	v253 = int32(_a_F_postmaster_child_launch_0)
	goto L44
L11:
	;
	v175 = F_strlen(m, v164)
	mBase = m.M
	goto L10
L13:
	;
	goto L14
L14:
	;
	v65 = int32(1023)
	if (v37^v58)&int32(3) != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v165))) = uint8(v168)
	goto L11
L16:
	;
	v149 = v144
	v150 = v145
	v151 = v146
	goto L37
L17:
	;
	if v139 == int32(0) {
		v164 = v137
		v165 = v138
		goto L15
	} else {
		goto L36
	}
L18:
	;
	v137 = v58
	v138 = v37
	v139 = v65
	goto L17
L19:
	;
	goto L20
L20:
	;
	v69 = int32(0)
	if base.B2i32(v58&int32(3) == v69)|int32(0) == v69 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v105 == int32(0) {
		v164 = v102
		v165 = v103
		goto L15
	} else {
		goto L30
	}
L22:
	;
	v81 = v58
	v82 = v37
	v83 = v65
	goto L25
L23:
	;
	goto L24
L24:
	;
	v102 = v58
	v103 = v37
	v104 = v65
	v105 = int32(1)
	goto L21
L25:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	*(*uint8)(unsafe.Add(mBase, uint32(v82))) = uint8(v85)
	if v85 == int32(0) {
		v144 = v81
		v145 = v82
		v146 = v83
		goto L16
	} else {
		goto L27
	}
L26:
	;
	v102 = v96
	v103 = v90
	v104 = v92
	v105 = v94
	goto L21
L27:
	;
	v89 = int32(1)
	v90 = v82 + v89
	v92 = v83 - v89
	v93 = int32(0)
	v94 = base.B2i32(v92 != v93)
	v96 = v81 + v89
	if v96&int32(3) == v93 {
		v102 = v96
		v103 = v90
		v104 = v92
		v105 = v94
		goto L21
	} else {
		goto L28
	}
L28:
	;
	if v92 != 0 {
		v81 = v96
		v82 = v90
		v83 = v92
		goto L25
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if base.B2i32(v108 == int32(0))|base.B2i32(base.Ui32(v104) < base.Ui32(int32(4))) != 0 {
		v137 = v102
		v138 = v103
		v139 = v104
		goto L17
	} else {
		goto L31
	}
L31:
	;
	v115 = v102
	v116 = v103
	v117 = v104
	goto L32
L32:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v123 = int32(-2139062144)
	if (int32(16843008)-v120|v120)&v123 != v123 {
		v144 = v115
		v145 = v116
		v146 = v117
		goto L16
	} else {
		goto L34
	}
L33:
	;
	v137 = v131
	v138 = v129
	v139 = v133
	goto L17
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v120
	v128 = int32(4)
	v129 = v116 + v128
	v131 = v115 + v128
	v133 = v117 - v128
	if base.Ui32(int32(3)) < base.Ui32(v133) {
		v115 = v131
		v116 = v129
		v117 = v133
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v144 = v137
	v145 = v138
	v146 = v139
	goto L16
L37:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	*(*uint8)(unsafe.Add(mBase, uint32(v150))) = uint8(v153)
	if v153 == int32(0) {
		v164 = v149
		v165 = v150
		goto L15
	} else {
		goto L39
	}
L38:
	;
	v164 = v160
	v165 = v158
	goto L15
L39:
	;
	v157 = int32(1)
	v158 = v150 + v157
	v160 = v149 + v157
	v162 = v151 - v157
	if v162 != 0 {
		v149 = v160
		v150 = v158
		v151 = v162
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v374 = v37 + int32(2160)
	v375 = int32(_a_F_postmaster_child_launch_1)
	goto L75
L42:
	;
	v370 = F_strlen(m, v359)
	mBase = m.M
	goto L41
L44:
	;
	goto L45
L45:
	;
	v260 = int32(1023)
	if (v252^v253)&int32(3) != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v363 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v360))) = uint8(v363)
	goto L42
L47:
	;
	v344 = v339
	v345 = v340
	v346 = v341
	goto L68
L48:
	;
	if v334 == int32(0) {
		v359 = v332
		v360 = v333
		goto L46
	} else {
		goto L67
	}
L49:
	;
	v332 = v253
	v333 = v252
	v334 = v260
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
	v303 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postmaster_child_launch[26])))
	if base.B2i32(v303 == int32(0))|int32(0) != 0 {
		v332 = v253
		v333 = v252
		v334 = v260
		goto L48
	} else {
		goto L62
	}
L62:
	;
	v310 = v253
	v311 = v252
	v312 = v260
	goto L63
L63:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v318 = int32(-2139062144)
	if (int32(16843008)-v315|v315)&v318 != v318 {
		v339 = v310
		v340 = v311
		v341 = v312
		goto L47
	} else {
		goto L65
	}
L64:
	;
	v332 = v326
	v333 = v324
	v334 = v328
	goto L48
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = v315
	v323 = int32(4)
	v324 = v311 + v323
	v326 = v310 + v323
	v328 = v312 - v323
	if base.Ui32(int32(3)) < base.Ui32(v328) {
		v310 = v326
		v311 = v324
		v312 = v328
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v339 = v332
	v340 = v333
	v341 = v334
	goto L47
L68:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
	*(*uint8)(unsafe.Add(mBase, uint32(v345))) = uint8(v348)
	if v348 == int32(0) {
		v359 = v344
		v360 = v345
		goto L46
	} else {
		goto L70
	}
L69:
	;
	v359 = v355
	v360 = v353
	goto L46
L70:
	;
	v352 = int32(1)
	v353 = v345 + v352
	v355 = v344 + v352
	v357 = v346 - v352
	if v357 != 0 {
		v344 = v355
		v345 = v353
		v346 = v357
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+3328)) = l3
	if l3 != 0 {
		goto L103
	} else {
		goto L104
	}
L73:
	;
	v492 = F_strlen(m, v481)
	mBase = m.M
	goto L72
L75:
	;
	goto L76
L76:
	;
	v382 = int32(1023)
	if (v374^v375)&int32(3) != 0 {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	v485 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v482))) = uint8(v485)
	goto L73
L78:
	;
	v466 = v461
	v467 = v462
	v468 = v463
	goto L99
L79:
	;
	if v456 == int32(0) {
		v481 = v454
		v482 = v455
		goto L77
	} else {
		goto L98
	}
L80:
	;
	v454 = v375
	v455 = v374
	v456 = v382
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
	v425 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_postmaster_child_launch[27])))
	if base.B2i32(v425 == int32(0))|int32(0) != 0 {
		v454 = v375
		v455 = v374
		v456 = v382
		goto L79
	} else {
		goto L93
	}
L93:
	;
	v432 = v375
	v433 = v374
	v434 = v382
	goto L94
L94:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v432)))
	v440 = int32(-2139062144)
	if (int32(16843008)-v437|v437)&v440 != v440 {
		v461 = v432
		v462 = v433
		v463 = v434
		goto L78
	} else {
		goto L96
	}
L95:
	;
	v454 = v448
	v455 = v446
	v456 = v450
	goto L79
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v433))) = v437
	v445 = int32(4)
	v446 = v433 + v445
	v448 = v432 + v445
	v450 = v434 - v445
	if base.Ui32(int32(3)) < base.Ui32(v450) {
		v432 = v448
		v433 = v446
		v434 = v450
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v461 = v454
	v462 = v455
	v463 = v456
	goto L78
L99:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466))))
	*(*uint8)(unsafe.Add(mBase, uint32(v467))) = uint8(v470)
	if v470 == int32(0) {
		v481 = v466
		v482 = v467
		goto L77
	} else {
		goto L101
	}
L100:
	;
	v481 = v477
	v482 = v475
	goto L77
L101:
	;
	v474 = int32(1)
	v475 = v467 + v474
	v477 = v466 + v474
	v479 = v468 - v474
	if v479 != 0 {
		v466 = v477
		v467 = v475
		v468 = v479
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	base.MemoryCopy(m, v37+int32(3332), l2, l3)
	goto L105
L104:
	;
	goto L105
L105:
	;
	v499 = int32(_a_F_postmaster_child_launch_2)
	v501 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[28]))
	v503 = v501 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[28])) = v503
	v505 = int32(_a_F_postmaster_child_launch_3)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v505
	v510 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[29]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v503
	v514 = v12 + int32(1104)
	v519 = F_pg_snprintf(m, v514, int32(1024), int32(_a_F_postmaster_child_launch_4), v12-int32(-64))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	v522 = F_AllocateFile(m, v514, int32(_a_F_postmaster_child_launch_5))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L4
	} else {
		goto L110
	}
L107:
	;
	m.G0 = v12 + int32(2128)
	return v624
L108:
	;
	F_pfree(m, v37)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L4
	} else {
		goto L146
	}
L109:
	;
	v550 = F_fwrite(m, v37, v36, int32(1), v548)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L4
	} else {
		goto L120
	}
L110:
	;
	if v522 != 0 {
		v548 = v522
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v526 = *(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[30]))
	v527 = F_mkdir(m, int32(_a_F_postmaster_child_launch_3), v526)
	mBase = m.M
	goto L112
L112:
	;
	v529 = F_AllocateFile(m, v514, int32(_a_F_postmaster_child_launch_5))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	if v529 != 0 {
		v548 = v529
		goto L109
	} else {
		goto L114
	}
L114:
	;
	v533 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	if v533 == int32(0) {
		goto L108
	} else {
		goto L116
	}
L116:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L4
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v514
	F_errmsg(m, int32(_a_F_postmaster_child_launch_6), v12)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L4
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_postmaster_child_launch_7), int32(355), int32(_a_F_postmaster_child_launch_8))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	goto L108
L120:
	;
	if v550 != int32(1) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v556 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
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
	v576 = m.ExcPending
	if v576 != 0 {
		goto L4
	} else {
		goto L132
	}
L124:
	;
	if v556 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L4
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v573 = F_FreeFile(m, v548)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L4
	} else {
		goto L131
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v12 + int32(1104)
	F_errmsg(m, int32(_a_F_postmaster_child_launch_9), v12+int32(48))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L4
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_postmaster_child_launch_7), int32(365), int32(_a_F_postmaster_child_launch_8))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
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
	v577 = F_FreeFile(m, v548)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	if v577 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v579 = int32(-1)
	v582 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L4
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v45
	v603 = v12 + int32(80)
	v608 = F_pg_snprintf(m, v603, int32(1024), int32(_a_F_postmaster_child_launch_10), v12+int32(16))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L4
	} else {
		goto L142
	}
L137:
	;
	if v582 == int32(0) {
		v624 = v579
		goto L107
	} else {
		goto L138
	}
L138:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L4
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v12 + int32(1104)
	F_errmsg(m, int32(_a_F_postmaster_child_launch_9), v12+int32(32))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L4
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_postmaster_child_launch_7), int32(377), int32(_a_F_postmaster_child_launch_8))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	v624 = v579
	goto L107
L142:
	;
	v613 = m.Env.Pgmem_spawn(m, v603, v12+int32(1104))
	mBase = m.M
	if int32(0) <= v613 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v624 = v613
	goto L107
L144:
	;
	goto L145
L145:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_postmaster_child_launch[31])) = int32(0) - v613
	v624 = int32(-1)
	goto L107
L146:
	;
	v624 = int32(-1)
	goto L107
}
