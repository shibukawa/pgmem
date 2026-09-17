package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SPI_execute_snapshot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	v6 = l5
	v8 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = int32(-6)
	if base.B2i32(l0 == v8)|base.B2i32(l6 < v8) != 0 {
		v146 = v21
		m.G0 = v19 + int32(32)
		return v146
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v27 != int32(569278163) {
			v146 = v21
			m.G0 = v19 + int32(32)
			return v146
		} else {
			if l1 != 0 {
				v35 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_snapshot[0]))
				if v35 == int32(0) {
					v146 = int32(-4)
					m.G0 = v19 + int32(32)
					return v146
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_snapshot[1]))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_snapshot[0]))
					*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v41
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
					*(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_snapshot[2])) = v46
					v48 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v48
					*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v48
					*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v48
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if int32(0) < v54 {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v58 = F_makeParamList(m, v54)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v72 = int32(0)
							for {
								v83 = v58 + int32(32) + v72*int32(12)
								v85 = v72 << (uint(int32(2)) % 32)
								v87 = *(*int32)(unsafe.Add(mBase, uint32(l1+v85)))
								*(*int32)(unsafe.Add(mBase, uint32(v83))) = v87
								if l2 != 0 {
									v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v72))))
									v94 = base.B2i32(v91 == int32(110))
								} else {
									v94 = int32(0)
								}
								v95 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v83)+6)) = uint16(v95)
								*(*uint8)(unsafe.Add(mBase, uint32(v83)+4)) = uint8(v94)
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v57+v85)))
								*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v99
								v102 = v72 + v95
								if v102 != v54 {
									v72 = v102
									continue
								} else {
									break
								}
								break
							}
							v117 = v58
							*(*uint8)(unsafe.Add(mBase, uint32(v19)+12)) = uint8(v6)
							*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v117
							*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = base.I64_extend_i32_u(l6)
							v127 = F__SPI_execute_plan(m, l0, v19+int32(8), l3, l4, int32(0))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return int32(0)
							} else {
								v131 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_snapshot[0]))
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+20))
								*(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_snapshot[2])) = v132
								*(*int32)(unsafe.Add(mBase, uint32(v131)+12)) = int32(0)
								v136 = *(*int32)(unsafe.Add(mBase, uint32(v131)+24))
								F_MemoryContextReset(m, v136)
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
									return int32(0)
								} else {
									v146 = v127
									m.G0 = v19 + int32(32)
									return v146
								}
							}
						}
					} else {
						v117 = v8
						*(*uint8)(unsafe.Add(mBase, uint32(v19)+12)) = uint8(v6)
						*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v117
						*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = base.I64_extend_i32_u(l6)
						v127 = F__SPI_execute_plan(m, l0, v19+int32(8), l3, l4, int32(0))
						mBase = m.M
						v128 = m.ExcPending
						if v128 != 0 {
							return int32(0)
						} else {
							v131 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_snapshot[0]))
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+20))
							*(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_snapshot[2])) = v132
							*(*int32)(unsafe.Add(mBase, uint32(v131)+12)) = int32(0)
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v131)+24))
							F_MemoryContextReset(m, v136)
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return int32(0)
							} else {
								v146 = v127
								m.G0 = v19 + int32(32)
								return v146
							}
						}
					}
				}
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v30 <= int32(0) {
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_snapshot[0]))
					if v35 == int32(0) {
						v146 = int32(-4)
						m.G0 = v19 + int32(32)
						return v146
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_snapshot[1]))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_snapshot[0]))
						*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v41
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
						*(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_snapshot[2])) = v46
						v48 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v48
						*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v48
						*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v48
						v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if int32(0) < v54 {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v58 = F_makeParamList(m, v54)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v72 = int32(0)
								for {
									v83 = v58 + int32(32) + v72*int32(12)
									v85 = v72 << (uint(int32(2)) % 32)
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l1+v85)))
									*(*int32)(unsafe.Add(mBase, uint32(v83))) = v87
									if l2 != 0 {
										v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v72))))
										v94 = base.B2i32(v91 == int32(110))
									} else {
										v94 = int32(0)
									}
									v95 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v83)+6)) = uint16(v95)
									*(*uint8)(unsafe.Add(mBase, uint32(v83)+4)) = uint8(v94)
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v57+v85)))
									*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v99
									v102 = v72 + v95
									if v102 != v54 {
										v72 = v102
										continue
									} else {
										break
									}
									break
								}
								v117 = v58
								*(*uint8)(unsafe.Add(mBase, uint32(v19)+12)) = uint8(v6)
								*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v117
								*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = base.I64_extend_i32_u(l6)
								v127 = F__SPI_execute_plan(m, l0, v19+int32(8), l3, l4, int32(0))
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return int32(0)
								} else {
									v131 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_snapshot[0]))
									v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+20))
									*(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_snapshot[2])) = v132
									*(*int32)(unsafe.Add(mBase, uint32(v131)+12)) = int32(0)
									v136 = *(*int32)(unsafe.Add(mBase, uint32(v131)+24))
									F_MemoryContextReset(m, v136)
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return int32(0)
									} else {
										v146 = v127
										m.G0 = v19 + int32(32)
										return v146
									}
								}
							}
						} else {
							v117 = v8
							*(*uint8)(unsafe.Add(mBase, uint32(v19)+12)) = uint8(v6)
							*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v117
							*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = base.I64_extend_i32_u(l6)
							v127 = F__SPI_execute_plan(m, l0, v19+int32(8), l3, l4, int32(0))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return int32(0)
							} else {
								v131 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_snapshot[0]))
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+20))
								*(*int32)(unsafe.Add(mBase, _c_F_SPI_execute_snapshot[2])) = v132
								*(*int32)(unsafe.Add(mBase, uint32(v131)+12)) = int32(0)
								v136 = *(*int32)(unsafe.Add(mBase, uint32(v131)+24))
								F_MemoryContextReset(m, v136)
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
									return int32(0)
								} else {
									v146 = v127
									m.G0 = v19 + int32(32)
									return v146
								}
							}
						}
					}
				} else {
					v146 = int32(-7)
					m.G0 = v19 + int32(32)
					return v146
				}
			}
		}
	}
}
func F_SPI_getbinval(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v6 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SPI_getbinval[0])) = v6
	if base.B2i32(l2 == v6)|base.B2i32(l2 < int32(-6)) == v6 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if l2 <= v15 {
			v24 = F_heap_getattr_2(m, l0, l2, l1, l3)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				return v24
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_SPI_getbinval[0])) = int32(-9)
			v20 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v20)
			return int32(0)
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_SPI_getbinval[0])) = int32(-9)
		v20 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v20)
		return int32(0)
	}
}
func F_SPI_returntuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	if l1 != 0 {
		v6 = l0
	} else {
		v6 = int32(0)
	}
	if v6 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_SPI_returntuple[0])) = int32(-6)
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_returntuple[1]))
		if v15 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_SPI_returntuple[0])) = int32(-4)
			return int32(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v23 != int32(2249) {
				v35 = v15
				v36 = int32(_a_F_SPI_returntuple_0)
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_returntuple[2]))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
				*(*int32)(unsafe.Add(mBase, _c_F_SPI_returntuple[2])) = v39
				v41 = F_heap_copy_tuple_as_datum(m, l0, l1)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v43 = F_pg_detoast_datum(m, v41)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_SPI_returntuple[2])) = v37
						return v43
					}
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				if int32(0) <= v26 {
					v35 = v15
					v36 = int32(_a_F_SPI_returntuple_0)
					v37 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_returntuple[2]))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
					*(*int32)(unsafe.Add(mBase, _c_F_SPI_returntuple[2])) = v39
					v41 = F_heap_copy_tuple_as_datum(m, l0, l1)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v43 = F_pg_detoast_datum(m, v41)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_SPI_returntuple[2])) = v37
							return v43
						}
					}
				} else {
					F_assign_record_type_typmod(m, l1)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_returntuple[1]))
						v35 = v34
						v36 = int32(_a_F_SPI_returntuple_0)
						v37 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_returntuple[2]))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
						*(*int32)(unsafe.Add(mBase, _c_F_SPI_returntuple[2])) = v39
						v41 = F_heap_copy_tuple_as_datum(m, l0, l1)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v43 = F_pg_detoast_datum(m, v41)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_SPI_returntuple[2])) = v37
								return v43
							}
						}
					}
				}
			}
		}
	}
}
func F__SPI_execute_plan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v282 int64
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int64
	_ = v374
	var v375 int64
	_ = v375
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int64
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v435 int64
	_ = v435
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int64
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v474 int64
	_ = v474
	var v476 int32
	_ = v476
	var v479 int64
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int64
	_ = v485
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v494 int64
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int64
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v550 int32
	_ = v550
	var v554 int64
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v587 int32
	_ = v587
	var v591 int64
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	v6 = int32(0)
	v24 = int64(0)
	v26 = m.G0
	v28 = v26 + int32(-64)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v32 != int32(1) {
		v46 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v47 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+56)) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = int32(779)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+60)) = v49
	v53 = int32(_a_F__SPI_execute_plan_0)
	v54 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[0]))
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[0])) = v26 + int32(-20)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+44)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v28)+52)) = v26 + int32(-8)
	if l2 == v47 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[1]))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+40)))
	if v38 != 0 {
		v46 = int32(0)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[2]))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+28))
	goto L4
L4:
	;
	v46 = base.B2i32(int32(1) < v41) ^ int32(1)
	goto L1
L5:
	;
	v75 = base.B2i32(l2 != int32(0))
	v77 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[3]))
	if v30 != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v65 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_PushActiveSnapshot(m, l2)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	F_PushCopiedSnapshot(m, l2)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L5
L12:
	;
	goto L5
L13:
	;
	v78 = v30
	goto L15
L14:
	;
	v78 = v77
	goto L15
L15:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v80 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v81 = v78
	goto L18
L17:
	;
	v81 = int32(0)
	goto L18
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
	if v83 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if v577 != 0 {
		goto L204
	} else {
		goto L205
	}
L20:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v102 <= int32(0) {
		v577 = v75
		v579 = v6
		v580 = v6
		v587 = v6
		v591 = v24
		goto L19
	} else {
		goto L30
	}
L21:
	;
	if v82 != 0 {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v82 != 0 {
		goto L20
	} else {
		goto L29
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(_a_F__SPI_execute_plan_1), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F__SPI_execute_plan_2), int32(2496), int32(_a_F__SPI_execute_plan_3))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	v577 = v75
	v579 = v6
	v580 = v6
	v587 = v6
	v591 = v24
	goto L19
L30:
	;
	if l4 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v107 = int32(0)
	goto L33
L32:
	;
	v107 = int32(32)
	goto L33
L33:
	;
	if v46 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v110 = int32(2)
	goto L36
L35:
	;
	v110 = int32(1)
	goto L36
L36:
	;
	v121 = v75
	v123 = v6
	v130 = v6
	v131 = v6
	v135 = v24
	goto L37
L37:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v130<<(uint(int32(2))%32))))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+56)) = v141
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v143 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v577 = v540
	v579 = v542
	v580 = int32(0)
	v587 = v550
	v591 = v554
	goto L19
L39:
	;
	v146 = int32(0)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v147 == v146 {
		v165 = v146
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
	if v177 != int32(1) {
		goto L50
	} else {
		goto L51
	}
L42:
	;
	v166 = int32(0)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_CompleteCachedPlan(m, v140, v165, v166, v167, v168, v169, v170, v171, v166)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L10
	} else {
		goto L49
	}
L43:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v150 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v153 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[1]))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+36))
	v155 = F_pg_analyze_and_rewrite_withcb(m, v147, v141, v150, v151, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L10
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v160 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[1]))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+36))
	v162 = F_pg_analyze_and_rewrite_fixedparams(m, v147, v141, v157, v158, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L10
	} else {
		goto L48
	}
L47:
	;
	v165 = v155
	goto L42
L48:
	;
	v165 = v162
	goto L42
L49:
	;
	goto L41
L50:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v207 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[1]))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+36))
	v209 = F_GetCachedPlan(m, v140, v205, v81, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L10
	} else {
		goto L63
	}
L51:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v140)+52))
	if v180 != 0 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
	if v182 != int32(179) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v182<<(uint(int32(3))%32))+uint32(_c_F__SPI_execute_plan[4])))
	goto L56
L54:
	;
	v188 = int32(_a_F__SPI_execute_plan_4)
	goto L55
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L10
	} else {
		goto L57
	}
L56:
	;
	v188 = v187
	goto L55
L57:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v188
	F_errmsg(m, int32(_a_F__SPI_execute_plan_5), v28)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F__SPI_execute_plan_2), int32(2570), int32(_a_F__SPI_execute_plan_3))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_ReleaseCachedPlan(m, v209, v81)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L10
	} else {
		goto L198
	}
L62:
	;
	v254 = int32(0)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	if v255 <= v254 {
		v540 = v253
		v542 = v123
		v550 = v131
		v554 = v135
		goto L61
	} else {
		goto L85
	}
L63:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	v212 = int32(0)
	if l2|base.B2i32(v211 == v212) == v212 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	if v217 <= int32(1) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	if v211 == int32(0) {
		v540 = v121
		v542 = v123
		v550 = v131
		v554 = v135
		goto L61
	} else {
		goto L84
	}
L67:
	;
	if v217 != int32(1) {
		v253 = v121
		goto L62
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	F_EnsurePortalSnapshotExists(m)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L10
	} else {
		goto L76
	}
L70:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v223)+88))
	if v226 == int32(0) {
		v234 = int32(1)
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v234 == int32(0) {
		v253 = v121
		goto L62
	} else {
		goto L75
	}
L72:
	;
	goto L71
L73:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	switch v230 - int32(158) {
	case 0, 1, 45, 64, 65, 66, 67, 86, 88, 89:
		v234 = int32(0)
		goto L72
	default:
		goto L74
	}
L74:
	;
	v234 = int32(1)
	goto L72
L75:
	;
	goto L69
L76:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if (v239|v46)&int32(1) != 0 {
		v253 = v121
		goto L62
	} else {
		goto L77
	}
L77:
	;
	if v121 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L10
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v245 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L10
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	F_PushActiveSnapshot(m, v245)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	v253 = int32(1)
	goto L62
L84:
	;
	v253 = v121
	goto L62
L85:
	;
	v270 = v123
	v274 = v254
	v278 = v131
	v282 = v135
	goto L88
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L10
	} else {
		goto L195
	}
L87:
	;
	v577 = v253
	v579 = int32(-2)
	v580 = v209
	v587 = v278
	v591 = v282
	goto L19
L88:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v211)+12))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v283+v274<<(uint(int32(2))%32))))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+26)))
	v290 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[1]))
	v291 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v290)+8)) = v291
	*(*int64)(unsafe.Add(mBase, uint32(v290))) = int64(0)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v287)+88))
	if v295 == v291 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v577 = v253
	v579 = v487
	v580 = v209
	v587 = v508
	v591 = v509
	goto L19
L90:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v308 == int32(1) {
		goto L99
	} else {
		goto L100
	}
L91:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	if v298 != int32(157) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if v298 != int32(225) {
		goto L90
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v295)+20))
	if v304 == int32(0) {
		goto L87
	} else {
		goto L96
	}
L95:
	;
	v577 = v253
	v579 = int32(-8)
	v580 = v209
	v587 = v278
	v591 = v282
	goto L19
L96:
	;
	goto L90
L97:
	;
	v344 = v288 & int32(1)
	if v344 != 0 {
		goto L116
	} else {
		goto L117
	}
L98:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L10
	} else {
		goto L113
	}
L99:
	;
	v311 = F_CommandIsReadOnly(m, v287)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L10
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	if v253 == int32(0) {
		goto L97
	} else {
		goto L112
	}
L102:
	;
	if v311 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v253&(v313^int32(-1)) != 0 {
		goto L98
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L10
	} else {
		goto L107
	}
L106:
	;
	goto L97
L107:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L10
	} else {
		goto L108
	}
L108:
	;
	v324 = F_CreateCommandName(m, v287)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L10
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v324
	F_errmsg(m, int32(_a_F__SPI_execute_plan_6), v26+int32(-48))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L10
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F__SPI_execute_plan_2), int32(2658), int32(_a_F__SPI_execute_plan_3))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L10
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	goto L98
L113:
	;
	F_UpdateActiveSnapshotCommandId(m)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L10
	} else {
		goto L114
	}
L114:
	;
	goto L97
L115:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v287)+88))
	if v353 == int32(0) {
		goto L122
	} else {
		goto L123
	}
L116:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v345 != 0 {
		v352 = v345
		goto L115
	} else {
		goto L119
	}
L117:
	;
	v349 = int32(0)
	goto L118
L118:
	;
	v350 = F_CreateDestReceiver(m, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L10
	} else {
		goto L120
	}
L119:
	;
	v349 = int32(5)
	goto L118
L120:
	;
	v352 = v350
	goto L115
L121:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[1]))
	if v344 != 0 {
		goto L186
	} else {
		goto L187
	}
L122:
	;
	v356 = int32(0)
	v358 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[5]))
	goto L125
L123:
	;
	goto L124
L124:
	;
	v449 = v26 + int32(-40)
	*(*int64)(unsafe.Add(mBase, uint32(v449)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v449))) = int32(0)
	goto L170
L125:
	;
	if v358 != v356 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[5]))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)))
	goto L129
L127:
	;
	v364 = v356
	goto L128
L128:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v369 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[1]))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)+36))
	v372 = F_CreateQueryDesc(m, v287, v366, v364, l3, v352, v367, v370, int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L10
	} else {
		goto L130
	}
L129:
	;
	v364 = v363
	goto L128
L130:
	;
	if v344 != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v374 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v375 = v374
	goto L133
L132:
	;
	v375 = int64(0)
	goto L133
L133:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	switch v377 - int32(1) {
	case 0:
		goto L141
	case 1:
		goto L138
	case 2:
		goto L140
	case 3:
		goto L139
	case 4:
		goto L137
	default:
		v442 = int32(-3)
		goto L134
	}
L134:
	;
	F_FreeQueryDesc(m, v372)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L10
	} else {
		goto L169
	}
L135:
	;
	F_ExecutorStart(m, v372, v107)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L10
	} else {
		goto L157
	}
L136:
	;
	v409 = v407
	v411 = int32(0)
	goto L135
L137:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+24)))
	if v405 != 0 {
		goto L154
	} else {
		goto L155
	}
L138:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+24)))
	if v400 != 0 {
		goto L151
	} else {
		goto L152
	}
L139:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+24)))
	if v395 != 0 {
		goto L148
	} else {
		goto L149
	}
L140:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389)+24)))
	if v390 != 0 {
		goto L145
	} else {
		goto L146
	}
L141:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v372)+20))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+16))
	if v383 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v384 = int32(5)
	goto L144
L143:
	;
	v384 = int32(4)
	goto L144
L144:
	;
	v409 = v384
	v411 = base.B2i32(v383 != int32(0))
	goto L135
L145:
	;
	v391 = int32(11)
	goto L147
L146:
	;
	v391 = int32(7)
	goto L147
L147:
	;
	v407 = v391
	goto L136
L148:
	;
	v396 = int32(12)
	goto L150
L149:
	;
	v396 = int32(8)
	goto L150
L150:
	;
	v407 = v396
	goto L136
L151:
	;
	v401 = int32(13)
	goto L153
L152:
	;
	v401 = int32(9)
	goto L153
L153:
	;
	v407 = v401
	goto L136
L154:
	;
	v406 = int32(19)
	goto L156
L155:
	;
	v406 = int32(18)
	goto L156
L156:
	;
	v407 = v406
	goto L136
L157:
	;
	F_ExecutorRun(m, v372, int32(1), v375)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L10
	} else {
		goto L158
	}
L158:
	;
	v418 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[1]))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v372)+40))
	v420 = *(*int64)(unsafe.Add(mBase, uint32(v419)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v418))) = v420
	if v411 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	F_ExecutorFinish(m, v372)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L10
	} else {
		goto L167
	}
L160:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424)+24)))
	if v425 != int32(1) {
		goto L159
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v372)+20))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)+16))
	if v429 != int32(5) {
		goto L159
	} else {
		goto L164
	}
L163:
	;
	goto L162
L164:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v418)+8))
	if v432 == int32(0) {
		goto L86
	} else {
		goto L165
	}
L165:
	;
	v435 = *(*int64)(unsafe.Add(mBase, uint32(v432)+8))
	if v420 != v435 {
		goto L86
	} else {
		goto L166
	}
L166:
	;
	goto L159
L167:
	;
	F_ExecutorEnd(m, v372)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L10
	} else {
		goto L168
	}
L168:
	;
	v442 = v409
	goto L134
L169:
	;
	v487 = v442
	goto L121
L170:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v458 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[1]))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v458)+36))
	F_ProcessUtility(m, v287, v454, int32(1), v110, v456, v459, v352, v449)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L10
	} else {
		goto L171
	}
L171:
	;
	v463 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[1]))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v463)+8))
	if v464 != 0 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v465 = *(*int64)(unsafe.Add(mBase, uint32(v464)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v463))) = v465
	goto L174
L173:
	;
	goto L174
L174:
	;
	v467 = int32(4)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v287)+88))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	if v469 != int32(157) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	if v469 != int32(242) {
		v487 = v467
		goto L121
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v485 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v463))) = v485
	v487 = v467
	goto L121
L178:
	;
	v474 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	if v476 == int32(179) {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v479 = v474
	goto L181
L180:
	;
	v479 = int64(0)
	goto L181
L181:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v463))) = v479
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468)+16)))
	if v483 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v484 = int32(6)
	goto L184
L183:
	;
	v484 = int32(4)
	goto L184
L184:
	;
	v487 = v484
	goto L121
L185:
	;
	if int32(0) <= v487 {
		goto L191
	} else {
		goto L192
	}
L186:
	;
	v494 = *(*int64)(unsafe.Add(mBase, uint32(v493)))
	F_SPI_freetuptable(m, v278)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L10
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v493)+8))
	F_SPI_freetuptable(m, v500)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L10
	} else {
		goto L190
	}
L189:
	;
	v498 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[1]))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)+8))
	v507 = v487
	v508 = v499
	v509 = v494
	goto L185
L190:
	;
	v504 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v504)+8)) = int32(0)
	v507 = v270
	v508 = v278
	v509 = v282
	goto L185
L191:
	;
	v513 = v274 + int32(1)
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	if v514 <= v513 {
		v540 = v253
		v542 = v507
		v550 = v508
		v554 = v509
		goto L61
	} else {
		goto L194
	}
L192:
	;
	goto L193
L193:
	;
	goto L89
L194:
	;
	v270 = v507
	v274 = v513
	v278 = v508
	v282 = v509
	goto L88
L195:
	;
	F_errmsg_internal(m, int32(_a_F__SPI_execute_plan_7), int32(0))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L10
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(_a_F__SPI_execute_plan_2), int32(2940), int32(_a_F__SPI_execute_plan_8))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L10
	} else {
		goto L197
	}
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L198:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v557 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L10
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v563 = v130 + int32(1)
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v563 < v564 {
		v121 = v540
		v123 = v542
		v130 = v563
		v131 = v550
		v135 = v554
		goto L37
	} else {
		goto L203
	}
L202:
	;
	goto L201
L203:
	;
	goto L38
L204:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L10
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	if v580 != 0 {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	goto L206
L208:
	;
	F_ReleaseCachedPlan(m, v580, v81)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L10
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v28)+44))
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[0])) = v597
	*(*int64)(unsafe.Add(mBase, _c_F__SPI_execute_plan[6])) = v591
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[7])) = v587
	v604 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_execute_plan[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v604)+8)) = int32(0)
	m.G0 = v28 - int32(-64)
	if v579 != 0 {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	goto L210
L212:
	;
	v611 = v579
	goto L214
L213:
	;
	v611 = int32(14)
	goto L214
L214:
	;
	return v611
}
func F__SPI_prepare_plan(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l0
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(779)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v15
	v19 = int32(_a_F__SPI_prepare_plan_0)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_prepare_plan[0]))
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_prepare_plan[0])) = v12 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v12 + int32(24)
	v29 = F_raw_parser(m, l0, v15)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v95 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)) = uint8(v95)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v92
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	*(*int32)(unsafe.Add(mBase, _c_F__SPI_prepare_plan[0])) = v99
	m.G0 = v12 + int32(32)
	return
L2:
	;
	return
L3:
	;
	if v29 == int32(0) {
		v92 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v33 <= int32(0) {
		v92 = v3
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v41 = int32(0)
	v43 = v3
	goto L6
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+v41<<(uint(int32(2))%32))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v52 = F_CreateCommandTag(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v92 = v80
	goto L1
L8:
	;
	v54 = F_CreateCachedPlan(m, v50, l0, v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v56 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v71 = int32(0)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_CompleteCachedPlan(m, v54, v70, v71, v72, v73, v74, v75, v76, v71)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L2
	} else {
		goto L16
	}
L11:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v59 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_prepare_plan[1]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+36))
	v61 = F_pg_analyze_and_rewrite_withcb(m, v50, l0, v56, v57, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v66 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_prepare_plan[1]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+36))
	v68 = F_pg_analyze_and_rewrite_fixedparams(m, v50, l0, v63, v64, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L15
	}
L14:
	;
	v70 = v61
	goto L10
L15:
	;
	v70 = v68
	goto L10
L16:
	;
	v80 = F_lappend(m, v43, v54)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v83 = v41 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v83 < v84 {
		v41 = v83
		v43 = v80
		goto L6
	} else {
		goto L18
	}
L18:
	;
	goto L7
}
func F_spi_printtup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int64
	_ = v42
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_spi_printtup[0]))
	if v7 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		if v8 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_spi_printtup_0), int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_spi_printtup_1), int32(2181), int32(_a_F_spi_printtup_2))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v11 = int32(_a_F_spi_printtup_3)
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_spi_printtup[1]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_spi_printtup[1])) = v14
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
			if base.Ui64(v16) <= base.Ui64(v17) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				v23 = F_repalloc_huge(m, v19, base.I32_wrap_i64(v16)<<(uint(int32(3))%32))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v16 << (uint(int64(1)) % 64)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v23
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
					v34 = m.T0[v33].(func(*base.Module, int32) int32)(m, l0)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v36+v37<<(uint(int32(2))%32)))) = v34
						v42 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v42 + int64(1)
						*(*int32)(unsafe.Add(mBase, _c_F_spi_printtup[1])) = v12
						return int32(1)
					}
				}
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
				v34 = m.T0[v33].(func(*base.Module, int32) int32)(m, l0)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v36+v37<<(uint(int32(2))%32)))) = v34
					v42 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v42 + int64(1)
					*(*int32)(unsafe.Add(mBase, _c_F_spi_printtup[1])) = v12
					return int32(1)
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_spi_printtup_4), int32(0))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_spi_printtup_1), int32(2177), int32(_a_F_spi_printtup_2))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
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
