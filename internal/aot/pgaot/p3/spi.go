package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SPI_cursor_open_with_paramlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_SPI_cursor_open_internal(m, l0, l1, l2, l3)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	v6 = l5
	v8 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = int32(-6)
	if l0 == v8 {
		v145 = v21
		m.G0 = v19 + int32(32)
		return v145
	} else {
		if l6 < int32(0) {
			v145 = v21
			m.G0 = v19 + int32(32)
			return v145
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v26 != int32(569278163) {
				v145 = v21
				m.G0 = v19 + int32(32)
				return v145
			} else {
				if l1 != 0 {
					v34 = *(*int32)(unsafe.Add(mBase, _consts[540]))
					if v34 == int32(0) {
						v145 = int32(-4)
						m.G0 = v19 + int32(32)
						return v145
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, _consts[75]))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
						v42 = *(*int32)(unsafe.Add(mBase, _consts[540]))
						*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v40
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v45
						v47 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v47
						*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v47
						*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v47
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if int32(0) < v53 {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v57 = F_makeParamList(m, v53)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								v71 = int32(0)
								for {
									v82 = v57 + int32(32) + v71*int32(12)
									v84 = v71 << (uint(int32(2)) % 32)
									v86 = *(*int32)(unsafe.Add(mBase, uint32(l1+v84)))
									*(*int32)(unsafe.Add(mBase, uint32(v82))) = v86
									if l2 != 0 {
										v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v71))))
										v93 = base.B2i32(v90 == int32(110))
									} else {
										v93 = int32(0)
									}
									v94 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v82)+6)) = uint16(v94)
									*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)) = uint8(v93)
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v56+v84)))
									*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v98
									v101 = v71 + v94
									if v101 != v53 {
										v71 = v101
										continue
									} else {
										break
									}
									break
								}
								v116 = v57
								*(*uint8)(unsafe.Add(mBase, uint32(v19)+12)) = uint8(v6)
								*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v116
								*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = base.I64_extend_i32_u(l6)
								v126 = F__SPI_execute_plan(m, l0, v19+int32(8), l3, l4, int32(0))
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return int32(0)
								} else {
									v130 = *(*int32)(unsafe.Add(mBase, _consts[540]))
									v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+20))
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v131
									*(*int32)(unsafe.Add(mBase, uint32(v130)+12)) = int32(0)
									v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+24))
									F_MemoryContextReset(m, v135)
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return int32(0)
									} else {
										v145 = v126
										m.G0 = v19 + int32(32)
										return v145
									}
								}
							}
						} else {
							v116 = v8
							*(*uint8)(unsafe.Add(mBase, uint32(v19)+12)) = uint8(v6)
							*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v116
							*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = base.I64_extend_i32_u(l6)
							v126 = F__SPI_execute_plan(m, l0, v19+int32(8), l3, l4, int32(0))
							mBase = m.M
							v127 = m.ExcPending
							if v127 != 0 {
								return int32(0)
							} else {
								v130 = *(*int32)(unsafe.Add(mBase, _consts[540]))
								v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+20))
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v131
								*(*int32)(unsafe.Add(mBase, uint32(v130)+12)) = int32(0)
								v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+24))
								F_MemoryContextReset(m, v135)
								mBase = m.M
								v137 = m.ExcPending
								if v137 != 0 {
									return int32(0)
								} else {
									v145 = v126
									m.G0 = v19 + int32(32)
									return v145
								}
							}
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v29 <= int32(0) {
						v34 = *(*int32)(unsafe.Add(mBase, _consts[540]))
						if v34 == int32(0) {
							v145 = int32(-4)
							m.G0 = v19 + int32(32)
							return v145
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, _consts[75]))
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
							v42 = *(*int32)(unsafe.Add(mBase, _consts[540]))
							*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v40
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v45
							v47 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v47
							*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v47
							*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v47
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if int32(0) < v53 {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								v57 = F_makeParamList(m, v53)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v71 = int32(0)
									for {
										v82 = v57 + int32(32) + v71*int32(12)
										v84 = v71 << (uint(int32(2)) % 32)
										v86 = *(*int32)(unsafe.Add(mBase, uint32(l1+v84)))
										*(*int32)(unsafe.Add(mBase, uint32(v82))) = v86
										if l2 != 0 {
											v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v71))))
											v93 = base.B2i32(v90 == int32(110))
										} else {
											v93 = int32(0)
										}
										v94 = int32(1)
										*(*uint16)(unsafe.Add(mBase, uint32(v82)+6)) = uint16(v94)
										*(*uint8)(unsafe.Add(mBase, uint32(v82)+4)) = uint8(v93)
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v56+v84)))
										*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v98
										v101 = v71 + v94
										if v101 != v53 {
											v71 = v101
											continue
										} else {
											break
										}
										break
									}
									v116 = v57
									*(*uint8)(unsafe.Add(mBase, uint32(v19)+12)) = uint8(v6)
									*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v116
									*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = base.I64_extend_i32_u(l6)
									v126 = F__SPI_execute_plan(m, l0, v19+int32(8), l3, l4, int32(0))
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
										return int32(0)
									} else {
										v130 = *(*int32)(unsafe.Add(mBase, _consts[540]))
										v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+20))
										*(*int32)(unsafe.Add(mBase, _consts[0])) = v131
										*(*int32)(unsafe.Add(mBase, uint32(v130)+12)) = int32(0)
										v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+24))
										F_MemoryContextReset(m, v135)
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return int32(0)
										} else {
											v145 = v126
											m.G0 = v19 + int32(32)
											return v145
										}
									}
								}
							} else {
								v116 = v8
								*(*uint8)(unsafe.Add(mBase, uint32(v19)+12)) = uint8(v6)
								*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v116
								*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = base.I64_extend_i32_u(l6)
								v126 = F__SPI_execute_plan(m, l0, v19+int32(8), l3, l4, int32(0))
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return int32(0)
								} else {
									v130 = *(*int32)(unsafe.Add(mBase, _consts[540]))
									v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+20))
									*(*int32)(unsafe.Add(mBase, _consts[0])) = v131
									*(*int32)(unsafe.Add(mBase, uint32(v130)+12)) = int32(0)
									v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+24))
									F_MemoryContextReset(m, v135)
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return int32(0)
									} else {
										v145 = v126
										m.G0 = v19 + int32(32)
										return v145
									}
								}
							}
						}
					} else {
						v145 = int32(-7)
						m.G0 = v19 + int32(32)
						return v145
					}
				}
			}
		}
	}
}
func F_SPI_getbinval(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	*(*int32)(unsafe.Add(mBase, _consts[504])) = int32(0)
	if l2 < int32(-6) {
		*(*int32)(unsafe.Add(mBase, _consts[504])) = int32(-9)
		v17 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v17)
		return int32(0)
	} else {
		if l2 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[504])) = int32(-9)
			v17 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v17)
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if l2 <= v12 {
				v21 = F_heap_getattr_2(m, l0, l2, l1, l3)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return v21
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[504])) = int32(-9)
				v17 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v17)
				return int32(0)
			}
		}
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
		*(*int32)(unsafe.Add(mBase, _consts[504])) = int32(-6)
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[540]))
		if v15 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[504])) = int32(-4)
			return int32(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v23 != int32(2249) {
				v35 = v15
				v36 = int32(4554240)
				v37 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v39
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
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v37
						return v43
					}
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				if int32(0) <= v26 {
					v35 = v15
					v36 = int32(4554240)
					v37 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v39
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
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v37
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
						v34 = *(*int32)(unsafe.Add(mBase, _consts[540]))
						v35 = v34
						v36 = int32(4554240)
						v37 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v39
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
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v37
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
	var v25 int64
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
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
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v286 int64
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int64
	_ = v361
	var v362 int64
	_ = v362
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int64
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v422 int64
	_ = v422
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int64
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v463 int64
	_ = v463
	var v465 int32
	_ = v465
	var v468 int64
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int64
	_ = v474
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v483 int64
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int64
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v530 int64
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v604 int64
	_ = v604
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v632 int64
	_ = v632
	var v634 int32
	_ = v634
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v660 int64
	_ = v660
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v671 int32
	_ = v671
	var v678 int32
	_ = v678
	v6 = int32(0)
	v25 = int64(0)
	v27 = m.G0
	v29 = v27 + int32(-64)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v33 != int32(1) {
		v47 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = int32(779)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v50
	v54 = int32(4547144)
	v55 = *(*int32)(unsafe.Add(mBase, _consts[52]))
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v27 + int32(-20)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v27 + int32(-8)
	if l2 == v48 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+40)))
	if v39 != 0 {
		v47 = int32(0)
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[75]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
	goto L4
L4:
	;
	v47 = base.B2i32(int32(1) < v42) ^ int32(1)
	goto L1
L5:
	;
	v76 = base.B2i32(l2 != int32(0))
	v78 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	if v31 != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v66 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_PushActiveSnapshot(m, l2)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
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
	v74 = m.ExcPending
	if v74 != 0 {
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
	v79 = v31
	goto L15
L14:
	;
	v79 = v78
	goto L15
L15:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v81 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v82 = v79
	goto L18
L17:
	;
	v82 = int32(0)
	goto L18
L18:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
	if v84 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	if v647 != 0 {
		goto L210
	} else {
		goto L211
	}
L20:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L10
	} else {
		goto L209
	}
L21:
	;
	if v592 == int32(0) {
		v644 = v588
		v647 = v591
		v650 = v594
		v660 = v604
		goto L19
	} else {
		goto L208
	}
L22:
	;
	v588 = v6
	v591 = v6
	v592 = v76
	v594 = v6
	v604 = v25
	goto L21
L23:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v103 <= int32(0) {
		goto L33
	} else {
		goto L34
	}
L24:
	;
	if v83 != 0 {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v83 != 0 {
		goto L23
	} else {
		goto L32
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	F_errmsg(m, int32(174114), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(521262), int32(2496), int32(296151))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	goto L22
L33:
	;
	goto L22
L34:
	;
	goto L35
L35:
	;
	if l4 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v108 = int32(0)
	goto L38
L37:
	;
	v108 = int32(32)
	goto L38
L38:
	;
	if v47 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v111 = int32(2)
	goto L41
L40:
	;
	v111 = int32(1)
	goto L41
L41:
	;
	v121 = v6
	v125 = v76
	v127 = v6
	v130 = v6
	v137 = v25
	goto L45
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L10
	} else {
		goto L205
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L10
	} else {
		goto L200
	}
L44:
	;
	v543 = int32(-2)
	if v254 == int32(0) {
		v644 = v543
		v647 = v213
		v650 = v276
		v660 = v286
		goto L19
	} else {
		goto L199
	}
L45:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v138+v130<<(uint(int32(2))%32))))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v143
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v145 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v588 = v514
	v591 = int32(0)
	v592 = v518
	v594 = v520
	v604 = v530
	goto L21
L47:
	;
	v148 = int32(0)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v149 == v148 {
		v167 = v148
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
	if v179 != int32(1) {
		goto L58
	} else {
		goto L59
	}
L50:
	;
	v168 = int32(0)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_CompleteCachedPlan(m, v142, v167, v168, v169, v170, v171, v172, v173, v168)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L10
	} else {
		goto L57
	}
L51:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v152 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v155 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+36))
	v157 = F_pg_analyze_and_rewrite_withcb(m, v149, v143, v152, v153, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L10
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v162 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+36))
	v164 = F_pg_analyze_and_rewrite_fixedparams(m, v149, v143, v159, v160, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L10
	} else {
		goto L56
	}
L55:
	;
	v167 = v157
	goto L50
L56:
	;
	v167 = v164
	goto L50
L57:
	;
	goto L49
L58:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v211 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+36))
	v213 = F_GetCachedPlan(m, v142, v209, v82, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L10
	} else {
		goto L69
	}
L59:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v142)+52))
	if v182 != 0 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
	if v184 != int32(179) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v184<<(uint(int32(3))%32))+uint32(_consts[541])))
	goto L64
L62:
	;
	v192 = int32(551570)
	goto L63
L63:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L10
	} else {
		goto L65
	}
L64:
	;
	v192 = v191
	goto L63
L65:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v192
	F_errmsg(m, int32(174149), v29)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L10
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(521262), int32(2570), int32(296151))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L10
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	if l2 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	F_ReleaseCachedPlan(m, v213, v82)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L10
	} else {
		goto L193
	}
L71:
	;
	v255 = int32(0)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v256 <= v255 {
		v514 = v121
		v518 = v254
		v520 = v127
		v530 = v137
		goto L70
	} else {
		goto L93
	}
L72:
	;
	if v215 == int32(0) {
		v514 = v121
		v518 = v125
		v520 = v127
		v530 = v137
		goto L70
	} else {
		goto L92
	}
L73:
	;
	if v215 == int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v218 <= int32(1) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	if v218 != int32(1) {
		v254 = v125
		goto L71
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	F_EnsurePortalSnapshotExists(m)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L10
	} else {
		goto L84
	}
L78:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v224)+88))
	if v227 == int32(0) {
		v235 = int32(1)
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v235 == int32(0) {
		v254 = v125
		goto L71
	} else {
		goto L83
	}
L80:
	;
	goto L79
L81:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	switch v231 - int32(158) {
	case 0, 1, 45, 64, 65, 66, 67, 86, 88, 89:
		v235 = int32(0)
		goto L80
	default:
		goto L82
	}
L82:
	;
	v235 = int32(1)
	goto L80
L83:
	;
	goto L77
L84:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if (v240|v47)&int32(1) != 0 {
		v254 = v125
		goto L71
	} else {
		goto L85
	}
L85:
	;
	if v125 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L10
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v246 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L10
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	F_PushActiveSnapshot(m, v246)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	v254 = int32(1)
	goto L71
L92:
	;
	v254 = v125
	goto L71
L93:
	;
	v269 = v255
	v270 = v121
	v276 = v127
	v286 = v137
	goto L94
L94:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v287+v269<<(uint(int32(2))%32))))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+26)))
	v294 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	v295 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v294)+8)) = v295
	*(*int64)(unsafe.Add(mBase, uint32(v294))) = int64(0)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v291)+88))
	if v299 == v295 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v588 = v476
	v591 = v213
	v592 = v254
	v594 = v497
	v604 = v498
	goto L21
L96:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v312 == int32(1) {
		goto L103
	} else {
		goto L104
	}
L97:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v299)))
	if v302 != int32(157) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	if v302 != int32(225) {
		goto L96
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v299)+20))
	if v308 == int32(0) {
		goto L44
	} else {
		goto L102
	}
L101:
	;
	v588 = int32(-8)
	v591 = v213
	v592 = v254
	v594 = v276
	v604 = v286
	goto L21
L102:
	;
	goto L96
L103:
	;
	v315 = F_CommandIsReadOnly(m, v291)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L10
	} else {
		goto L106
	}
L104:
	;
	v320 = v312
	goto L105
L105:
	;
	if (v320|base.B2i32(v254 == int32(0)))&int32(1) == int32(0) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	if v315 == int32(0) {
		goto L43
	} else {
		goto L107
	}
L107:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	v320 = v319
	goto L105
L108:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L10
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v331 = v292 & int32(1)
	if v331 != 0 {
		goto L114
	} else {
		goto L115
	}
L111:
	;
	F_UpdateActiveSnapshotCommandId(m)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L10
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v291)+88))
	if v340 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L114:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v332 != 0 {
		v339 = v332
		goto L113
	} else {
		goto L117
	}
L115:
	;
	v336 = int32(0)
	goto L116
L116:
	;
	v337 = F_CreateDestReceiver(m, v336)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L10
	} else {
		goto L118
	}
L117:
	;
	v336 = int32(5)
	goto L116
L118:
	;
	v339 = v337
	goto L113
L119:
	;
	v482 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	if v331 != 0 {
		goto L184
	} else {
		goto L185
	}
L120:
	;
	v343 = int32(0)
	v345 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	goto L123
L121:
	;
	goto L122
L122:
	;
	v436 = v27 + int32(-40)
	*(*int64)(unsafe.Add(mBase, uint32(v436)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v436))) = int32(0)
	goto L168
L123:
	;
	if v345 != v343 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v349 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	goto L127
L125:
	;
	v351 = v343
	goto L126
L126:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v356 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+36))
	v359 = F_CreateQueryDesc(m, v291, v353, v351, l3, v339, v354, v357, int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L10
	} else {
		goto L128
	}
L127:
	;
	v351 = v350
	goto L126
L128:
	;
	if v331 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v361 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	v362 = v361
	goto L131
L130:
	;
	v362 = int64(0)
	goto L131
L131:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v359)))
	switch v364 - int32(1) {
	case 0:
		goto L139
	case 1:
		goto L136
	case 2:
		goto L138
	case 3:
		goto L137
	case 4:
		goto L135
	default:
		v429 = int32(-3)
		goto L132
	}
L132:
	;
	F_FreeQueryDesc(m, v359)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L10
	} else {
		goto L167
	}
L133:
	;
	F_ExecutorStart(m, v359, v108)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L10
	} else {
		goto L155
	}
L134:
	;
	v396 = v394
	v398 = int32(0)
	goto L133
L135:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+24)))
	if v392 != 0 {
		goto L152
	} else {
		goto L153
	}
L136:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+24)))
	if v387 != 0 {
		goto L149
	} else {
		goto L150
	}
L137:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+24)))
	if v382 != 0 {
		goto L146
	} else {
		goto L147
	}
L138:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+24)))
	if v377 != 0 {
		goto L143
	} else {
		goto L144
	}
L139:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v359)+20))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)+16))
	if v370 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v371 = int32(5)
	goto L142
L141:
	;
	v371 = int32(4)
	goto L142
L142:
	;
	v396 = v371
	v398 = base.B2i32(v370 != int32(0))
	goto L133
L143:
	;
	v378 = int32(11)
	goto L145
L144:
	;
	v378 = int32(7)
	goto L145
L145:
	;
	v394 = v378
	goto L134
L146:
	;
	v383 = int32(12)
	goto L148
L147:
	;
	v383 = int32(8)
	goto L148
L148:
	;
	v394 = v383
	goto L134
L149:
	;
	v388 = int32(13)
	goto L151
L150:
	;
	v388 = int32(9)
	goto L151
L151:
	;
	v394 = v388
	goto L134
L152:
	;
	v393 = int32(19)
	goto L154
L153:
	;
	v393 = int32(18)
	goto L154
L154:
	;
	v394 = v393
	goto L134
L155:
	;
	F_ExecutorRun(m, v359, int32(1), v362)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L10
	} else {
		goto L156
	}
L156:
	;
	v405 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v359)+40))
	v407 = *(*int64)(unsafe.Add(mBase, uint32(v406)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v405))) = v407
	if v398 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	F_ExecutorFinish(m, v359)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L10
	} else {
		goto L165
	}
L158:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411)+24)))
	if v412 != int32(1) {
		goto L157
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v359)+20))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)+16))
	if v416 != int32(5) {
		goto L157
	} else {
		goto L162
	}
L161:
	;
	goto L160
L162:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v405)+8))
	if v419 == int32(0) {
		goto L42
	} else {
		goto L163
	}
L163:
	;
	v422 = *(*int64)(unsafe.Add(mBase, uint32(v419)+8))
	if v407 != v422 {
		goto L42
	} else {
		goto L164
	}
L164:
	;
	goto L157
L165:
	;
	F_ExecutorEnd(m, v359)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L10
	} else {
		goto L166
	}
L166:
	;
	v429 = v396
	goto L132
L167:
	;
	v476 = v429
	goto L119
L168:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v445 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+36))
	F_ProcessUtility(m, v291, v441, int32(1), v111, v443, v446, v339, v27+int32(-40))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L10
	} else {
		goto L169
	}
L169:
	;
	v452 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v452)+8))
	if v453 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v454 = *(*int64)(unsafe.Add(mBase, uint32(v453)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v452))) = v454
	goto L172
L171:
	;
	goto L172
L172:
	;
	v456 = int32(4)
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v291)+88))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v457)))
	if v458 != int32(157) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	if v458 != int32(242) {
		v476 = v456
		goto L119
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v474 = *(*int64)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v452))) = v474
	v476 = v456
	goto L119
L176:
	;
	v463 = *(*int64)(unsafe.Add(mBase, uint32(v29)+32))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	if v465 == int32(179) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v468 = v463
	goto L179
L178:
	;
	v468 = int64(0)
	goto L179
L179:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v452))) = v468
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457)+16)))
	if v472 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v473 = int32(6)
	goto L182
L181:
	;
	v473 = int32(4)
	goto L182
L182:
	;
	v476 = v473
	goto L119
L183:
	;
	if int32(0) <= v476 {
		goto L189
	} else {
		goto L190
	}
L184:
	;
	v483 = *(*int64)(unsafe.Add(mBase, uint32(v482)))
	F_SPI_freetuptable(m, v276)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L10
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v482)+8))
	F_SPI_freetuptable(m, v489)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L10
	} else {
		goto L188
	}
L187:
	;
	v487 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)+8))
	v496 = v476
	v497 = v488
	v498 = v483
	goto L183
L188:
	;
	v493 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	*(*int32)(unsafe.Add(mBase, uint32(v493)+8)) = int32(0)
	v496 = v270
	v497 = v276
	v498 = v286
	goto L183
L189:
	;
	v502 = v269 + int32(1)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v503 <= v502 {
		v514 = v496
		v518 = v254
		v520 = v497
		v530 = v498
		goto L70
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	goto L95
L192:
	;
	v269 = v502
	v270 = v496
	v276 = v497
	v286 = v498
	goto L94
L193:
	;
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v533 == int32(0) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L10
	} else {
		goto L197
	}
L195:
	;
	goto L196
L196:
	;
	v539 = v130 + int32(1)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v539 < v540 {
		v121 = v514
		v125 = v518
		v127 = v520
		v130 = v539
		v137 = v530
		goto L45
	} else {
		goto L198
	}
L197:
	;
	goto L196
L198:
	;
	goto L46
L199:
	;
	v616 = v543
	v619 = v213
	v622 = v276
	v632 = v286
	goto L20
L200:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L10
	} else {
		goto L201
	}
L201:
	;
	v553 = F_CreateCommandName(m, v291)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L10
	} else {
		goto L202
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v553
	F_errmsg(m, int32(266064), v27+int32(-48))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L10
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(521262), int32(2658), int32(296151))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L10
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L205:
	;
	F_errmsg_internal(m, int32(474665), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L10
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(521262), int32(2940), int32(15489))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L10
	} else {
		goto L207
	}
L207:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L208:
	;
	v616 = v588
	v619 = v591
	v622 = v594
	v632 = v604
	goto L20
L209:
	;
	v644 = v616
	v647 = v619
	v650 = v622
	v660 = v632
	goto L19
L210:
	;
	F_ReleaseCachedPlan(m, v647, v82)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L10
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v664
	*(*int64)(unsafe.Add(mBase, _consts[502])) = v660
	*(*int32)(unsafe.Add(mBase, _consts[503])) = v650
	v671 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	*(*int32)(unsafe.Add(mBase, uint32(v671)+8)) = int32(0)
	m.G0 = v29 - int32(-64)
	if v644 != 0 {
		goto L214
	} else {
		goto L215
	}
L213:
	;
	goto L212
L214:
	;
	v678 = v644
	goto L216
L215:
	;
	v678 = int32(14)
	goto L216
L216:
	;
	return v678
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
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v91 int32
	_ = v91
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
	v19 = int32(4547144)
	v20 = *(*int32)(unsafe.Add(mBase, _consts[52]))
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v12 + int32(12)
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
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v91
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v99
	m.G0 = v12 + int32(32)
	return
L2:
	;
	return
L3:
	;
	if v29 == int32(0) {
		v91 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v33 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v34 <= v33 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v91 = v3
	goto L1
L6:
	;
	goto L7
L7:
	;
	v41 = v33
	v42 = v3
	goto L8
L8:
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
		goto L10
	}
L9:
	;
	v91 = v80
	goto L1
L10:
	;
	v54 = F_CreateCachedPlan(m, v50, l0, v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v56 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
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
		goto L18
	}
L13:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v59 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+36))
	v61 = F_pg_analyze_and_rewrite_withcb(m, v50, l0, v56, v57, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v66 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+36))
	v68 = F_pg_analyze_and_rewrite_fixedparams(m, v50, l0, v63, v64, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L17
	}
L16:
	;
	v70 = v61
	goto L12
L17:
	;
	v70 = v68
	goto L12
L18:
	;
	v80 = F_lappend(m, v42, v54)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v83 = v41 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v83 < v84 {
		v41 = v83
		v42 = v80
		goto L8
	} else {
		goto L20
	}
L20:
	;
	goto L9
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
	v7 = *(*int32)(unsafe.Add(mBase, _consts[540]))
	if v7 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		if v8 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(242956), int32(0))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(521262), int32(2181), int32(242973))
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
			v11 = int32(4554240)
			v12 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v14
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
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v12
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
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v12
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
			F_errmsg_internal(m, int32(559066), int32(0))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(521262), int32(2177), int32(242973))
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
