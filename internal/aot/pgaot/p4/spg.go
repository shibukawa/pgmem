package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_spgFormInnerTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	if l1 != 0 {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
		if v20 != 0 {
			v50 = int32(4)
		} else {
			v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+48)))
			if int32(0) < v21 {
				v50 = v21
			} else {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
				if v24 == int32(1) {
					v27 = int32(6)
					v29 = int32(18)
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
					if v31 == v29 {
						v34 = v29
					} else {
						v34 = int32(2)
					}
					if v31&int32(254) == int32(2) {
						v39 = v27
					} else {
						v39 = v34
					}
					if v31 == int32(1) {
						v42 = v27
					} else {
						v42 = v39
					}
					v50 = v42
				} else {
					if v24&int32(1) != 0 {
						v50 = int32(base.Ui32(v24) >> (uint(int32(1)) % 32))
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v50 = int32(base.Ui32(v47) >> (uint(int32(2)) % 32))
					}
				}
			}
		}
		v56 = (v50 + int32(7)) & int32(-8)
	} else {
		v56 = v6
	}
	v58 = v56 + int32(8)
	if l3 <= int32(0) {
		v160 = v58
	} else {
		v62 = l3 & int32(3)
		if base.Ui32(l3) < base.Ui32(int32(4)) {
			v116 = int32(0)
			v117 = v58
		} else {
			v74 = int32(0)
			v75 = v58
			v82 = v6
			for {
				v85 = l4 + v74<<(uint(int32(2))%32)
				v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
				v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+6)))
				v88 = int32(8191)
				v91 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
				v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+6)))
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
				v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v96)+6)))
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
				v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+6)))
				v105 = v75 + v87&v88 + v92&v88 + v97&v88 + v102&v88
				v106 = int32(4)
				v107 = v74 + v106
				v109 = v82 + v106
				if v109 != l3&int32(2147483644) {
					v74 = v107
					v75 = v105
					v82 = v109
					continue
				} else {
					break
				}
				break
			}
			v116 = v107
			v117 = v105
		}
		if v62 == int32(0) {
			v160 = v117
		} else {
			v132 = v116
			v133 = v117
			v137 = v6
			for {
				v144 = *(*int32)(unsafe.Add(mBase, uint32(l4+v132<<(uint(int32(2))%32))))
				v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144)+6)))
				v148 = v133 + v145&int32(8191)
				v149 = int32(1)
				v152 = v137 + v149
				if v152 != v62 {
					v132 = v132 + v149
					v133 = v148
					v137 = v152
					continue
				} else {
					break
				}
				break
			}
			v160 = v148
		}
	}
	v168 = int32(16)
	if base.Ui32(v160) <= base.Ui32(v168) {
		v171 = v168
	} else {
		v171 = v160
	}
	if base.Ui32(v160) < base.Ui32(int32(8157)) {
		if int32(8191) < l3 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v362 = m.ExcPending
			if v362 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(299697), int32(0))
				mBase = m.M
				v366 = m.ExcPending
				if v366 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(485466), int32(1048), int32(378330))
					mBase = m.M
					v371 = m.ExcPending
					if v371 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			if base.Ui32(int32(65536)) <= base.Ui32(v56) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v362 = m.ExcPending
				if v362 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(299697), int32(0))
					mBase = m.M
					v366 = m.ExcPending
					if v366 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(485466), int32(1048), int32(378330))
						mBase = m.M
						v371 = m.ExcPending
						if v371 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v178 = F_palloc0(m, v171)
				mBase = m.M
				v181 = m.ExcPending
				if v181 != 0 {
					return int32(0)
				} else {
					*(*uint16)(unsafe.Add(mBase, uint32(v178)+4)) = uint16(v171)
					v187 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
					*(*int32)(unsafe.Add(mBase, uint32(v178))) = l3<<(uint(int32(3))%32)&int32(65528) | v187&int32(7) | v56<<(uint(int32(16))%32)
					if l1 == int32(0) {
					} else {
						if v56 != 0 {
							v200 = v178 + int32(8)
						} else {
							v200 = int32(0)
						}
						v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
						if v201 == int32(1) {
							*(*int32)(unsafe.Add(mBase, uint32(v200))) = l2
						} else {
							v205 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+48)))
							if int32(0) < v205 {
								v234 = v205
							} else {
								v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
								if v208 == int32(1) {
									v211 = int32(6)
									v213 = int32(18)
									v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
									if v215 == v213 {
										v218 = v213
									} else {
										v218 = int32(2)
									}
									if v215&int32(254) == int32(2) {
										v223 = v211
									} else {
										v223 = v218
									}
									if v215 == int32(1) {
										v226 = v211
									} else {
										v226 = v223
									}
									v234 = v226
								} else {
									if v208&int32(1) != 0 {
										v234 = int32(base.Ui32(v208) >> (uint(int32(1)) % 32))
									} else {
										v231 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
										v234 = int32(base.Ui32(v231) >> (uint(int32(2)) % 32))
									}
								}
							}
							if v234 != 0 {
								v235 = F__emscripten_memcpy_bulkmem(m, v200, l2, v234)
								mBase = m.M
							} else {
							}
						}
					}
					if l3 <= int32(0) {
					} else {
						v241 = int32(1)
						v245 = v178 + v56 + int32(8)
						v246 = int32(0)
						if l3 != v241 {
							v257 = v246
							v260 = v245
							v262 = int32(0)
							for {
								v268 = l4 + v257<<(uint(int32(2))%32)
								v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
								v270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v269)+6)))
								v272 = v270 & int32(8191)
								if v272 != 0 {
									v273 = F__emscripten_memcpy_bulkmem(m, v260, v269, v272)
									mBase = m.M
									v274 = v273
								} else {
									v274 = v260
								}
								v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v269)+6)))
								v276 = int32(8191)
								v278 = v274 + v275&v276
								v279 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
								v280 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v279)+6)))
								v282 = v280 & v276
								if v282 != 0 {
									v283 = F__emscripten_memcpy_bulkmem(m, v278, v279, v282)
									mBase = m.M
									v284 = v283
								} else {
									v284 = v278
								}
								v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v279)+6)))
								v288 = v284 + v285&int32(8191)
								v289 = int32(2)
								v290 = v257 + v289
								v292 = v262 + v289
								if v292 != l3&int32(2147483646) {
									v257 = v290
									v260 = v288
									v262 = v292
									continue
								} else {
									break
								}
								break
							}
							v299 = v290
							v302 = v288
						} else {
							v299 = v246
							v302 = v245
						}
						if l3&v241 == int32(0) {
						} else {
							v313 = *(*int32)(unsafe.Add(mBase, uint32(l4+v299<<(uint(int32(2))%32))))
							v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v313)+6)))
							v316 = v314 & int32(8191)
							if v316 != 0 {
								v317 = F__emscripten_memcpy_bulkmem(m, v302, v313, v316)
								mBase = m.M
							} else {
							}
						}
					}
					m.G0 = v17 + int32(16)
					return v178
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v340 = m.ExcPending
		if v340 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v343 = m.ExcPending
			if v343 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = int32(8156)
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = v171
				F_errmsg(m, int32(36331), v17)
				mBase = m.M
				v349 = m.ExcPending
				if v349 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(621401), int32(0))
					mBase = m.M
					v353 = m.ExcPending
					if v353 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(485466), int32(1039), int32(378330))
						mBase = m.M
						v358 = m.ExcPending
						if v358 != 0 {
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
	}
}
func F_spgFormLeafTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
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
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	v5 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v14)+14)) = uint16(v5)
	if v17 < int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v51 = F_heap_compute_data_size(m, v16, l2, l3)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v43 = v5
	goto L1
L3:
	;
	goto L4
L4:
	;
	v23 = v5
	goto L5
L5:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+l3))))
	if v35 != 0 {
		v43 = v35
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v43 = v35
	goto L1
L7:
	;
	v37 = v23 + int32(1)
	if v37 != v17 {
		v23 = v37
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return int32(0)
L10:
	;
	v56 = v51 + int32(23)
	if base.Ui32(v56) <= base.Ui32(int32(16)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v59 = int32(16)
	goto L13
L12:
	;
	v59 = v56
	goto L13
L13:
	;
	v62 = F_palloc0(m, v59&int32(-8))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+4)))
	v66 = v64 & int32(-16384)
	*(*uint16)(unsafe.Add(mBase, uint32(v62)+4)) = uint16(v66)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v59<<(uint(int32(2))%32)&int32(-32) | v72&int32(3)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+6)) = v77
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v62)+10)) = uint16(v79)
	if v43 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	m.G0 = v14 + int32(16)
	return v62
L16:
	;
	F_heap_fill_tuple(m, v16, l2, l3, v62+int32(16), v14+int32(14), v90)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L9
	} else {
		goto L22
	}
L17:
	;
	v82 = v66 | int32(32768)
	*(*uint16)(unsafe.Add(mBase, uint32(v62)+4)) = uint16(v82)
	v90 = v62 + int32(12)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v86 = int32(0)
	if int32(1) < v17 {
		v90 = v86
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v89 != 0 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v90 = v86
	goto L16
L22:
	;
	goto L15
}
func F_spgGetCache(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int64
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int64
	_ = v255
	var v257 int64
	_ = v257
	var v259 int64
	_ = v259
	var v261 int64
	_ = v261
	var v263 int64
	_ = v263
	var v265 int64
	_ = v265
	var v267 int64
	_ = v267
	var v269 int64
	_ = v269
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v12 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L6
	} else {
		goto L87
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L6
	} else {
		goto L83
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v17 = F_MemoryContextAllocZero(m, v15, int32(128))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v277 = v12
	goto L5
L5:
	;
	m.G0 = v10 + int32(16)
	return v277
L6:
	;
	return int32(0)
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v22 <= int32(3830) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v150
	v155 = int32(1)
	v157 = F_index_getprocinfo(m, l0, v155, v155)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L6
	} else {
		goto L55
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+48)))
	if v42 != 0 {
		goto L19
	} else {
		goto L20
	}
L10:
	;
	switch v22 - int32(2277) {
	case 0, 6:
		goto L9
	case 1, 2, 3, 4, 5:
		v150 = v22
		goto L8
	default:
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(v22-int32(5077)) < base.Ui32(int32(4)) {
		goto L9
	} else {
		goto L16
	}
L13:
	;
	if v22 == int32(2776) {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	if v22 == int32(3500) {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v150 = v22
	goto L8
L16:
	;
	if base.Ui32(v22-int32(4537)) < base.Ui32(int32(2)) {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	if v22 != int32(3831) {
		v150 = v22
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L9
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v44 = F_get_atttype(m, v43, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	if v48 != 0 {
		v53 = v48
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v46 = F_getBaseType(m, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v150 = v46
	goto L8
L24:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v57)+10)))
	if v58 <= int32(0) {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v55 = v53
	v56 = v54
	goto L24
L26:
	;
	v49 = F_RelationGetIndexExpressions(m, l0)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	if v49 != 0 {
		v53 = v49
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v51 = int32(0)
	v55 = v51
	v56 = v51
	goto L24
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L6
	} else {
		goto L52
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L49
	}
L31:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+48)))
	if v61 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v56 == int32(0) {
		goto L29
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v58 == int32(1) {
		goto L30
	} else {
		goto L38
	}
L35:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v67 = F_exprType(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	v69 = F_getBaseType(m, v67)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v150 = v69
	goto L8
L38:
	;
	v77 = v56
	v78 = int32(2)
	goto L39
L39:
	;
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78<<(uint(int32(1))%32)+(v57+int32(48))-int32(2)))))
	if v88 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L30
L41:
	;
	if v77 == int32(0) {
		goto L29
	} else {
		goto L44
	}
L42:
	;
	v104 = v77
	goto L43
L43:
	;
	if base.B2i32(v78 == v58) == int32(0) {
		v77 = v104
		v78 = v78 + int32(1)
		goto L39
	} else {
		goto L48
	}
L44:
	;
	v94 = v77 + int32(4)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if base.Ui32(v94) < base.Ui32(v96+v97<<(uint(int32(2))%32)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v102 = v94
	goto L47
L46:
	;
	v102 = int32(0)
	goto L47
L47:
	;
	v104 = v102
	goto L43
L48:
	;
	goto L40
L49:
	;
	F_errmsg_internal(m, int32(142702), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(485466), int32(160), int32(365124))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	F_errmsg_internal(m, int32(142702), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(485466), int32(154), int32(365124))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v163 = F_FunctionCall2Coll(m, v157, v160, v10+int32(12), v17)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v165 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v181 = v17 + int32(16)
	F_fillTypeDesc(m, v181, v150)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L6
	} else {
		goto L62
	}
L58:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v166+v167<<(uint(int32(4))%32))+88))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v171
	if v150 == v171 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v174 = F_IsBinaryCoercible(m, v171, v150)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	if v174 == int32(0) {
		goto L57
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v150
	goto L57
L62:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v150 != v184 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	F_fillTypeDesc(m, v17+int32(40), v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L6
	} else {
		goto L70
	}
L64:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189)+6)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v188+v190*int32(0)<<(uint(int32(2))%32)+int32(24)-int32(4))))
	goto L67
L65:
	;
	goto L66
L66:
	;
	v210 = *(*int64)(unsafe.Add(mBase, uint32(v181)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+28)) = v210
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v181)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v212
	goto L63
L67:
	;
	if v202 == int32(0) {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_fillTypeDesc(m, v17+int32(28), v207)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	goto L63
L70:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	F_fillTypeDesc(m, v17+int32(52), v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+119)))
	if v225 != int32(73) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v229 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L6
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v17
	v277 = v17
	goto L5
L75:
	;
	F_LockBuffer(m, v229, int32(1))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L6
	} else {
		goto L76
	}
L76:
	;
	if v229 < int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+24))
	if v252 != int32(-1173640210) {
		goto L1
	} else {
		goto L81
	}
L78:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v237+(v229^int32(-1))<<(uint(int32(2))%32))))
	v251 = v243
	goto L77
L79:
	;
	goto L80
L80:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v251 = v245 + v229<<(uint(int32(13))%32) + int32(-8192)
	goto L77
L81:
	;
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v251)+84))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = v255
	v257 = *(*int64)(unsafe.Add(mBase, uint32(v251)+76))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = v257
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v251)+68))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+104)) = v259
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v251)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+96)) = v261
	v263 = *(*int64)(unsafe.Add(mBase, uint32(v251)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+88)) = v263
	v265 = *(*int64)(unsafe.Add(mBase, uint32(v251)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = v265
	v267 = *(*int64)(unsafe.Add(mBase, uint32(v251)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+72)) = v267
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v251)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = v269
	F_UnlockReleaseBuffer(m, v229)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	goto L74
L83:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	F_errmsg(m, int32(361373), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L6
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(485466), int32(251), int32(393299))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v305 + int32(4)
	F_errmsg_internal(m, int32(28467), v10)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(485466), int32(280), int32(393299))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L6
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_spgTestLeafTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	v5 = l4
	v9 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32)+l2)+20))
	v24 = l2 + v21&int32(32767)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v27 = v25 & int32(3)
	if v27 != 0 {
		if l5 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v43 & int32(3)
				F_errmsg_internal(m, int32(475646), v16+int32(16))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(488373), int32(798), int32(378516))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			switch v27 - int32(1) {
			case 0:
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v31
				v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v33)
				v252 = int32(2049)
				m.G0 = v16 + int32(80)
				return v252
			case 1:
				v252 = int32(0)
				m.G0 = v16 + int32(80)
				return v252
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
					*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v43 & int32(3)
					F_errmsg_internal(m, int32(475646), v16+int32(16))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(488373), int32(798), int32(378516))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
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
	} else {
		if v5 != 0 {
			v57 = int32(0)
			v120 = v57
			v121 = v9
			v122 = v9
			v123 = v57
			v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
			if int32(0) < v124 {
				v127 = int32(4470560)
				v128 = *(*int32)(unsafe.Add(mBase, _consts[28]))
				v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				*(*int32)(unsafe.Add(mBase, _consts[28])) = v130
				v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
				if v5 == int32(0) {
					v139 = F_palloc(m, v124<<(uint(int32(3))%32)+int32(40))
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return int32(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v139)+34)) = uint8(v5)
						v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
						if int32(0) < v142 {
							v148 = v142 << (uint(int32(3)) % 32)
							if v148 != 0 {
								v149 = F__emscripten_memcpy_bulkmem(m, v139+int32(40), v122, v148)
								mBase = m.M
							} else {
							}
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = v132
						v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
						*(*uint16)(unsafe.Add(mBase, uint32(v139)+32)) = uint16(v152)
						v154 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
						*(*int32)(unsafe.Add(mBase, uint32(v139)+28)) = v154
						v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
						if v156 == int32(0) {
							v173 = v139
							*(*int64)(unsafe.Add(mBase, uint32(v173)+12)) = int64(0)
							v221 = v173
							v224 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
							*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
							v229 = v120 & v224
							*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
							v232 = v123 & v224
							*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
							v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							F_pairingheap_add(m, v234, v221)
							mBase = m.M
							v236 = m.ExcPending
							if v236 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
								v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
								v252 = v249 & int32(16383)
								m.G0 = v16 + int32(80)
								return v252
							}
						} else {
							v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
							v160 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
							v161 = F_datumCopy(m, v121, v159, v160)
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return int32(0)
							} else {
								v200 = v139
								v202 = v161
								*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = v202
								v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
								v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
								if int32(2) <= v205 {
									v208 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
									v211 = F_palloc(m, int32(base.Ui32(v208)>>(uint(int32(2))%32)))
									mBase = m.M
									v212 = m.ExcPending
									if v212 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = v211
										v214 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
										v216 = int32(base.Ui32(v214) >> (uint(int32(2)) % 32))
										if v216 != 0 {
											v217 = F__emscripten_memcpy_bulkmem(m, v211, v24, v216)
											mBase = m.M
										} else {
										}
										v221 = v200
										v224 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
										*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
										v229 = v120 & v224
										*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
										v232 = v123 & v224
										*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
										v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
										F_pairingheap_add(m, v234, v221)
										mBase = m.M
										v236 = m.ExcPending
										if v236 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
											v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
											v252 = v249 & int32(16383)
											m.G0 = v16 + int32(80)
											return v252
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = int32(0)
									v221 = v200
									v224 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
									*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
									v229 = v120 & v224
									*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
									v232 = v123 & v224
									*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
									v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
									F_pairingheap_add(m, v234, v221)
									mBase = m.M
									v236 = m.ExcPending
									if v236 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
										v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
										v252 = v249 & int32(16383)
										m.G0 = v16 + int32(80)
										return v252
									}
								}
							}
						}
					}
				} else {
					v164 = F_palloc(m, int32(40))
					mBase = m.M
					v165 = m.ExcPending
					if v165 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v164)+24)) = v132
						*(*uint8)(unsafe.Add(mBase, uint32(v164)+34)) = uint8(v5)
						v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
						*(*uint16)(unsafe.Add(mBase, uint32(v164)+32)) = uint16(v168)
						v170 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
						*(*int32)(unsafe.Add(mBase, uint32(v164)+28)) = v170
						v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
						if v172 != 0 {
							v200 = v164
							v202 = v9
							*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = v202
							v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
							v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
							if int32(2) <= v205 {
								v208 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
								v211 = F_palloc(m, int32(base.Ui32(v208)>>(uint(int32(2))%32)))
								mBase = m.M
								v212 = m.ExcPending
								if v212 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = v211
									v214 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
									v216 = int32(base.Ui32(v214) >> (uint(int32(2)) % 32))
									if v216 != 0 {
										v217 = F__emscripten_memcpy_bulkmem(m, v211, v24, v216)
										mBase = m.M
									} else {
									}
									v221 = v200
									v224 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
									*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
									v229 = v120 & v224
									*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
									v232 = v123 & v224
									*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
									v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
									F_pairingheap_add(m, v234, v221)
									mBase = m.M
									v236 = m.ExcPending
									if v236 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
										v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
										v252 = v249 & int32(16383)
										m.G0 = v16 + int32(80)
										return v252
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = int32(0)
								v221 = v200
								v224 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
								*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
								v229 = v120 & v224
								*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
								v232 = v123 & v224
								*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
								v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
								F_pairingheap_add(m, v234, v221)
								mBase = m.M
								v236 = m.ExcPending
								if v236 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
									v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
									v252 = v249 & int32(16383)
									m.G0 = v16 + int32(80)
									return v252
								}
							}
						} else {
							v173 = v164
							*(*int64)(unsafe.Add(mBase, uint32(v173)+12)) = int64(0)
							v221 = v173
							v224 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
							*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
							v229 = v120 & v224
							*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
							v232 = v123 & v224
							*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
							v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							F_pairingheap_add(m, v234, v221)
							mBase = m.M
							v236 = m.ExcPending
							if v236 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
								v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
								v252 = v249 & int32(16383)
								m.G0 = v16 + int32(80)
								return v252
							}
						}
					}
				}
			} else {
				v181 = int32(0)
				m.T0[l7].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32))(m, l0, v24+int32(6), v121, v5, v24, v123&int32(1), v181, v181)
				mBase = m.M
				v184 = m.ExcPending
				if v184 != 0 {
					return int32(0)
				} else {
					v185 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v185)
					v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
					v252 = v249 & int32(16383)
					m.G0 = v16 + int32(80)
					return v252
				}
			}
		} else {
			v59 = int32(4470560)
			v60 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
			*(*int32)(unsafe.Add(mBase, _consts[28])) = v62
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v64
			v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v66
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v68
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v70
			v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v72
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v74
			v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v76
			v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
			*(*uint8)(unsafe.Add(mBase, uint32(v16)+72)) = uint8(v78)
			v81 = v24 + int32(16)
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
			if v82 != int32(1) {
				v94 = v81
				v95 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v95
				*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v95
				*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v94
				*(*uint16)(unsafe.Add(mBase, uint32(v16)+36)) = uint16(v95)
				v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v109 = F_FunctionCall2Coll(m, l0+int32(160), v104, v16+int32(44), v16+int32(32))
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[28])) = v60
					if v109 == int32(0) {
						v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
						v252 = v249 & int32(16383)
						m.G0 = v16 + int32(80)
						return v252
					} else {
						v115 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
						v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+37)))
						v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)))
						v120 = v117
						v121 = v116
						v122 = v115
						v123 = v118
						v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
						if int32(0) < v124 {
							v127 = int32(4470560)
							v128 = *(*int32)(unsafe.Add(mBase, _consts[28]))
							v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, _consts[28])) = v130
							v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							if v5 == int32(0) {
								v139 = F_palloc(m, v124<<(uint(int32(3))%32)+int32(40))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v139)+34)) = uint8(v5)
									v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
									if int32(0) < v142 {
										v148 = v142 << (uint(int32(3)) % 32)
										if v148 != 0 {
											v149 = F__emscripten_memcpy_bulkmem(m, v139+int32(40), v122, v148)
											mBase = m.M
										} else {
										}
									} else {
									}
									*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = v132
									v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
									*(*uint16)(unsafe.Add(mBase, uint32(v139)+32)) = uint16(v152)
									v154 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
									*(*int32)(unsafe.Add(mBase, uint32(v139)+28)) = v154
									v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
									if v156 == int32(0) {
										v173 = v139
										*(*int64)(unsafe.Add(mBase, uint32(v173)+12)) = int64(0)
										v221 = v173
										v224 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
										*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
										v229 = v120 & v224
										*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
										v232 = v123 & v224
										*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
										v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
										F_pairingheap_add(m, v234, v221)
										mBase = m.M
										v236 = m.ExcPending
										if v236 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
											v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
											v252 = v249 & int32(16383)
											m.G0 = v16 + int32(80)
											return v252
										}
									} else {
										v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
										v160 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
										v161 = F_datumCopy(m, v121, v159, v160)
										mBase = m.M
										v162 = m.ExcPending
										if v162 != 0 {
											return int32(0)
										} else {
											v200 = v139
											v202 = v161
											*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = v202
											v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
											if int32(2) <= v205 {
												v208 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
												v211 = F_palloc(m, int32(base.Ui32(v208)>>(uint(int32(2))%32)))
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = v211
													v214 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
													v216 = int32(base.Ui32(v214) >> (uint(int32(2)) % 32))
													if v216 != 0 {
														v217 = F__emscripten_memcpy_bulkmem(m, v211, v24, v216)
														mBase = m.M
													} else {
													}
													v221 = v200
													v224 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
													*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
													v229 = v120 & v224
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
													v232 = v123 & v224
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
													v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v234, v221)
													mBase = m.M
													v236 = m.ExcPending
													if v236 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
														v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
														v252 = v249 & int32(16383)
														m.G0 = v16 + int32(80)
														return v252
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = int32(0)
												v221 = v200
												v224 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
												*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
												v229 = v120 & v224
												*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
												v232 = v123 & v224
												*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
												v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v234, v221)
												mBase = m.M
												v236 = m.ExcPending
												if v236 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
													v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
													v252 = v249 & int32(16383)
													m.G0 = v16 + int32(80)
													return v252
												}
											}
										}
									}
								}
							} else {
								v164 = F_palloc(m, int32(40))
								mBase = m.M
								v165 = m.ExcPending
								if v165 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v164)+24)) = v132
									*(*uint8)(unsafe.Add(mBase, uint32(v164)+34)) = uint8(v5)
									v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
									*(*uint16)(unsafe.Add(mBase, uint32(v164)+32)) = uint16(v168)
									v170 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
									*(*int32)(unsafe.Add(mBase, uint32(v164)+28)) = v170
									v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
									if v172 != 0 {
										v200 = v164
										v202 = v9
										*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = v202
										v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
										v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
										if int32(2) <= v205 {
											v208 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
											v211 = F_palloc(m, int32(base.Ui32(v208)>>(uint(int32(2))%32)))
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = v211
												v214 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
												v216 = int32(base.Ui32(v214) >> (uint(int32(2)) % 32))
												if v216 != 0 {
													v217 = F__emscripten_memcpy_bulkmem(m, v211, v24, v216)
													mBase = m.M
												} else {
												}
												v221 = v200
												v224 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
												*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
												v229 = v120 & v224
												*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
												v232 = v123 & v224
												*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
												v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v234, v221)
												mBase = m.M
												v236 = m.ExcPending
												if v236 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
													v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
													v252 = v249 & int32(16383)
													m.G0 = v16 + int32(80)
													return v252
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = int32(0)
											v221 = v200
											v224 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
											*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
											v229 = v120 & v224
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
											v232 = v123 & v224
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
											v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
											F_pairingheap_add(m, v234, v221)
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
												v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
												v252 = v249 & int32(16383)
												m.G0 = v16 + int32(80)
												return v252
											}
										}
									} else {
										v173 = v164
										*(*int64)(unsafe.Add(mBase, uint32(v173)+12)) = int64(0)
										v221 = v173
										v224 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
										*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
										v229 = v120 & v224
										*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
										v232 = v123 & v224
										*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
										v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
										F_pairingheap_add(m, v234, v221)
										mBase = m.M
										v236 = m.ExcPending
										if v236 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
											v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
											v252 = v249 & int32(16383)
											m.G0 = v16 + int32(80)
											return v252
										}
									}
								}
							}
						} else {
							v181 = int32(0)
							m.T0[l7].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32))(m, l0, v24+int32(6), v121, v5, v24, v123&int32(1), v181, v181)
							mBase = m.M
							v184 = m.ExcPending
							if v184 != 0 {
								return int32(0)
							} else {
								v185 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v185)
								v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
								v252 = v249 & int32(16383)
								m.G0 = v16 + int32(80)
								return v252
							}
						}
					}
				}
			} else {
				v85 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+36)))
				switch v85&int32(65535) - int32(1) {
				case 0:
					v90 = int32(*(*int8)(unsafe.Add(mBase, uint32(v81))))
					v94 = v90
					v95 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v95
					*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v95
					*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v94
					*(*uint16)(unsafe.Add(mBase, uint32(v16)+36)) = uint16(v95)
					v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v109 = F_FunctionCall2Coll(m, l0+int32(160), v104, v16+int32(44), v16+int32(32))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[28])) = v60
						if v109 == int32(0) {
							v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
							v252 = v249 & int32(16383)
							m.G0 = v16 + int32(80)
							return v252
						} else {
							v115 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
							v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+37)))
							v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)))
							v120 = v117
							v121 = v116
							v122 = v115
							v123 = v118
							v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
							if int32(0) < v124 {
								v127 = int32(4470560)
								v128 = *(*int32)(unsafe.Add(mBase, _consts[28]))
								v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
								*(*int32)(unsafe.Add(mBase, _consts[28])) = v130
								v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								if v5 == int32(0) {
									v139 = F_palloc(m, v124<<(uint(int32(3))%32)+int32(40))
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(v139)+34)) = uint8(v5)
										v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
										if int32(0) < v142 {
											v148 = v142 << (uint(int32(3)) % 32)
											if v148 != 0 {
												v149 = F__emscripten_memcpy_bulkmem(m, v139+int32(40), v122, v148)
												mBase = m.M
											} else {
											}
										} else {
										}
										*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = v132
										v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
										*(*uint16)(unsafe.Add(mBase, uint32(v139)+32)) = uint16(v152)
										v154 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
										*(*int32)(unsafe.Add(mBase, uint32(v139)+28)) = v154
										v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
										if v156 == int32(0) {
											v173 = v139
											*(*int64)(unsafe.Add(mBase, uint32(v173)+12)) = int64(0)
											v221 = v173
											v224 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
											*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
											v229 = v120 & v224
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
											v232 = v123 & v224
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
											v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
											F_pairingheap_add(m, v234, v221)
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
												v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
												v252 = v249 & int32(16383)
												m.G0 = v16 + int32(80)
												return v252
											}
										} else {
											v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
											v160 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
											v161 = F_datumCopy(m, v121, v159, v160)
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return int32(0)
											} else {
												v200 = v139
												v202 = v161
												*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = v202
												v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
												if int32(2) <= v205 {
													v208 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
													v211 = F_palloc(m, int32(base.Ui32(v208)>>(uint(int32(2))%32)))
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = v211
														v214 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
														v216 = int32(base.Ui32(v214) >> (uint(int32(2)) % 32))
														if v216 != 0 {
															v217 = F__emscripten_memcpy_bulkmem(m, v211, v24, v216)
															mBase = m.M
														} else {
														}
														v221 = v200
														v224 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
														*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
														v229 = v120 & v224
														*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
														v232 = v123 & v224
														*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
														v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
														F_pairingheap_add(m, v234, v221)
														mBase = m.M
														v236 = m.ExcPending
														if v236 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
															v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
															v252 = v249 & int32(16383)
															m.G0 = v16 + int32(80)
															return v252
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = int32(0)
													v221 = v200
													v224 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
													*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
													v229 = v120 & v224
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
													v232 = v123 & v224
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
													v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v234, v221)
													mBase = m.M
													v236 = m.ExcPending
													if v236 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
														v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
														v252 = v249 & int32(16383)
														m.G0 = v16 + int32(80)
														return v252
													}
												}
											}
										}
									}
								} else {
									v164 = F_palloc(m, int32(40))
									mBase = m.M
									v165 = m.ExcPending
									if v165 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v164)+24)) = v132
										*(*uint8)(unsafe.Add(mBase, uint32(v164)+34)) = uint8(v5)
										v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
										*(*uint16)(unsafe.Add(mBase, uint32(v164)+32)) = uint16(v168)
										v170 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
										*(*int32)(unsafe.Add(mBase, uint32(v164)+28)) = v170
										v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
										if v172 != 0 {
											v200 = v164
											v202 = v9
											*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = v202
											v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
											if int32(2) <= v205 {
												v208 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
												v211 = F_palloc(m, int32(base.Ui32(v208)>>(uint(int32(2))%32)))
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = v211
													v214 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
													v216 = int32(base.Ui32(v214) >> (uint(int32(2)) % 32))
													if v216 != 0 {
														v217 = F__emscripten_memcpy_bulkmem(m, v211, v24, v216)
														mBase = m.M
													} else {
													}
													v221 = v200
													v224 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
													*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
													v229 = v120 & v224
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
													v232 = v123 & v224
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
													v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v234, v221)
													mBase = m.M
													v236 = m.ExcPending
													if v236 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
														v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
														v252 = v249 & int32(16383)
														m.G0 = v16 + int32(80)
														return v252
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = int32(0)
												v221 = v200
												v224 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
												*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
												v229 = v120 & v224
												*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
												v232 = v123 & v224
												*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
												v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v234, v221)
												mBase = m.M
												v236 = m.ExcPending
												if v236 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
													v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
													v252 = v249 & int32(16383)
													m.G0 = v16 + int32(80)
													return v252
												}
											}
										} else {
											v173 = v164
											*(*int64)(unsafe.Add(mBase, uint32(v173)+12)) = int64(0)
											v221 = v173
											v224 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
											*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
											v229 = v120 & v224
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
											v232 = v123 & v224
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
											v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
											F_pairingheap_add(m, v234, v221)
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
												v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
												v252 = v249 & int32(16383)
												m.G0 = v16 + int32(80)
												return v252
											}
										}
									}
								}
							} else {
								v181 = int32(0)
								m.T0[l7].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32))(m, l0, v24+int32(6), v121, v5, v24, v123&int32(1), v181, v181)
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return int32(0)
								} else {
									v185 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v185)
									v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
									v252 = v249 & int32(16383)
									m.G0 = v16 + int32(80)
									return v252
								}
							}
						}
					}
				case 1:
					v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(v81))))
					v94 = v91
					v95 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v95
					*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v95
					*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v94
					*(*uint16)(unsafe.Add(mBase, uint32(v16)+36)) = uint16(v95)
					v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v109 = F_FunctionCall2Coll(m, l0+int32(160), v104, v16+int32(44), v16+int32(32))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[28])) = v60
						if v109 == int32(0) {
							v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
							v252 = v249 & int32(16383)
							m.G0 = v16 + int32(80)
							return v252
						} else {
							v115 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
							v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+37)))
							v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)))
							v120 = v117
							v121 = v116
							v122 = v115
							v123 = v118
							v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
							if int32(0) < v124 {
								v127 = int32(4470560)
								v128 = *(*int32)(unsafe.Add(mBase, _consts[28]))
								v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
								*(*int32)(unsafe.Add(mBase, _consts[28])) = v130
								v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								if v5 == int32(0) {
									v139 = F_palloc(m, v124<<(uint(int32(3))%32)+int32(40))
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(v139)+34)) = uint8(v5)
										v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
										if int32(0) < v142 {
											v148 = v142 << (uint(int32(3)) % 32)
											if v148 != 0 {
												v149 = F__emscripten_memcpy_bulkmem(m, v139+int32(40), v122, v148)
												mBase = m.M
											} else {
											}
										} else {
										}
										*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = v132
										v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
										*(*uint16)(unsafe.Add(mBase, uint32(v139)+32)) = uint16(v152)
										v154 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
										*(*int32)(unsafe.Add(mBase, uint32(v139)+28)) = v154
										v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
										if v156 == int32(0) {
											v173 = v139
											*(*int64)(unsafe.Add(mBase, uint32(v173)+12)) = int64(0)
											v221 = v173
											v224 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
											*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
											v229 = v120 & v224
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
											v232 = v123 & v224
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
											v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
											F_pairingheap_add(m, v234, v221)
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
												v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
												v252 = v249 & int32(16383)
												m.G0 = v16 + int32(80)
												return v252
											}
										} else {
											v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
											v160 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
											v161 = F_datumCopy(m, v121, v159, v160)
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return int32(0)
											} else {
												v200 = v139
												v202 = v161
												*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = v202
												v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
												if int32(2) <= v205 {
													v208 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
													v211 = F_palloc(m, int32(base.Ui32(v208)>>(uint(int32(2))%32)))
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = v211
														v214 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
														v216 = int32(base.Ui32(v214) >> (uint(int32(2)) % 32))
														if v216 != 0 {
															v217 = F__emscripten_memcpy_bulkmem(m, v211, v24, v216)
															mBase = m.M
														} else {
														}
														v221 = v200
														v224 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
														*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
														v229 = v120 & v224
														*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
														v232 = v123 & v224
														*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
														v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
														F_pairingheap_add(m, v234, v221)
														mBase = m.M
														v236 = m.ExcPending
														if v236 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
															v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
															v252 = v249 & int32(16383)
															m.G0 = v16 + int32(80)
															return v252
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = int32(0)
													v221 = v200
													v224 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
													*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
													v229 = v120 & v224
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
													v232 = v123 & v224
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
													v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v234, v221)
													mBase = m.M
													v236 = m.ExcPending
													if v236 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
														v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
														v252 = v249 & int32(16383)
														m.G0 = v16 + int32(80)
														return v252
													}
												}
											}
										}
									}
								} else {
									v164 = F_palloc(m, int32(40))
									mBase = m.M
									v165 = m.ExcPending
									if v165 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v164)+24)) = v132
										*(*uint8)(unsafe.Add(mBase, uint32(v164)+34)) = uint8(v5)
										v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
										*(*uint16)(unsafe.Add(mBase, uint32(v164)+32)) = uint16(v168)
										v170 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
										*(*int32)(unsafe.Add(mBase, uint32(v164)+28)) = v170
										v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
										if v172 != 0 {
											v200 = v164
											v202 = v9
											*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = v202
											v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
											if int32(2) <= v205 {
												v208 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
												v211 = F_palloc(m, int32(base.Ui32(v208)>>(uint(int32(2))%32)))
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = v211
													v214 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
													v216 = int32(base.Ui32(v214) >> (uint(int32(2)) % 32))
													if v216 != 0 {
														v217 = F__emscripten_memcpy_bulkmem(m, v211, v24, v216)
														mBase = m.M
													} else {
													}
													v221 = v200
													v224 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
													*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
													v229 = v120 & v224
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
													v232 = v123 & v224
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
													v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v234, v221)
													mBase = m.M
													v236 = m.ExcPending
													if v236 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
														v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
														v252 = v249 & int32(16383)
														m.G0 = v16 + int32(80)
														return v252
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = int32(0)
												v221 = v200
												v224 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
												*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
												v229 = v120 & v224
												*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
												v232 = v123 & v224
												*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
												v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v234, v221)
												mBase = m.M
												v236 = m.ExcPending
												if v236 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
													v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
													v252 = v249 & int32(16383)
													m.G0 = v16 + int32(80)
													return v252
												}
											}
										} else {
											v173 = v164
											*(*int64)(unsafe.Add(mBase, uint32(v173)+12)) = int64(0)
											v221 = v173
											v224 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
											*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
											v229 = v120 & v224
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
											v232 = v123 & v224
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
											v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
											F_pairingheap_add(m, v234, v221)
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
												v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
												v252 = v249 & int32(16383)
												m.G0 = v16 + int32(80)
												return v252
											}
										}
									}
								}
							} else {
								v181 = int32(0)
								m.T0[l7].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32))(m, l0, v24+int32(6), v121, v5, v24, v123&int32(1), v181, v181)
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return int32(0)
								} else {
									v185 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v185)
									v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
									v252 = v249 & int32(16383)
									m.G0 = v16 + int32(80)
									return v252
								}
							}
						}
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v190 = m.ExcPending
					if v190 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = v85
						F_errmsg_internal(m, int32(475060), v16)
						mBase = m.M
						v194 = m.ExcPending
						if v194 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(321141), int32(70), int32(66797))
							mBase = m.M
							v199 = m.ExcPending
							if v199 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 3:
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
					v94 = v92
					v95 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v95
					*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v95
					*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v94
					*(*uint16)(unsafe.Add(mBase, uint32(v16)+36)) = uint16(v95)
					v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v109 = F_FunctionCall2Coll(m, l0+int32(160), v104, v16+int32(44), v16+int32(32))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[28])) = v60
						if v109 == int32(0) {
							v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
							v252 = v249 & int32(16383)
							m.G0 = v16 + int32(80)
							return v252
						} else {
							v115 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
							v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+37)))
							v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)))
							v120 = v117
							v121 = v116
							v122 = v115
							v123 = v118
							v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
							if int32(0) < v124 {
								v127 = int32(4470560)
								v128 = *(*int32)(unsafe.Add(mBase, _consts[28]))
								v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
								*(*int32)(unsafe.Add(mBase, _consts[28])) = v130
								v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								if v5 == int32(0) {
									v139 = F_palloc(m, v124<<(uint(int32(3))%32)+int32(40))
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(v139)+34)) = uint8(v5)
										v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
										if int32(0) < v142 {
											v148 = v142 << (uint(int32(3)) % 32)
											if v148 != 0 {
												v149 = F__emscripten_memcpy_bulkmem(m, v139+int32(40), v122, v148)
												mBase = m.M
											} else {
											}
										} else {
										}
										*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = v132
										v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
										*(*uint16)(unsafe.Add(mBase, uint32(v139)+32)) = uint16(v152)
										v154 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
										*(*int32)(unsafe.Add(mBase, uint32(v139)+28)) = v154
										v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
										if v156 == int32(0) {
											v173 = v139
											*(*int64)(unsafe.Add(mBase, uint32(v173)+12)) = int64(0)
											v221 = v173
											v224 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
											*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
											v229 = v120 & v224
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
											v232 = v123 & v224
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
											v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
											F_pairingheap_add(m, v234, v221)
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
												v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
												v252 = v249 & int32(16383)
												m.G0 = v16 + int32(80)
												return v252
											}
										} else {
											v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
											v160 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
											v161 = F_datumCopy(m, v121, v159, v160)
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return int32(0)
											} else {
												v200 = v139
												v202 = v161
												*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = v202
												v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
												if int32(2) <= v205 {
													v208 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
													v211 = F_palloc(m, int32(base.Ui32(v208)>>(uint(int32(2))%32)))
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = v211
														v214 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
														v216 = int32(base.Ui32(v214) >> (uint(int32(2)) % 32))
														if v216 != 0 {
															v217 = F__emscripten_memcpy_bulkmem(m, v211, v24, v216)
															mBase = m.M
														} else {
														}
														v221 = v200
														v224 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
														*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
														v229 = v120 & v224
														*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
														v232 = v123 & v224
														*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
														v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
														F_pairingheap_add(m, v234, v221)
														mBase = m.M
														v236 = m.ExcPending
														if v236 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
															v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
															v252 = v249 & int32(16383)
															m.G0 = v16 + int32(80)
															return v252
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = int32(0)
													v221 = v200
													v224 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
													*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
													v229 = v120 & v224
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
													v232 = v123 & v224
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
													v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v234, v221)
													mBase = m.M
													v236 = m.ExcPending
													if v236 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
														v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
														v252 = v249 & int32(16383)
														m.G0 = v16 + int32(80)
														return v252
													}
												}
											}
										}
									}
								} else {
									v164 = F_palloc(m, int32(40))
									mBase = m.M
									v165 = m.ExcPending
									if v165 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v164)+24)) = v132
										*(*uint8)(unsafe.Add(mBase, uint32(v164)+34)) = uint8(v5)
										v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
										*(*uint16)(unsafe.Add(mBase, uint32(v164)+32)) = uint16(v168)
										v170 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
										*(*int32)(unsafe.Add(mBase, uint32(v164)+28)) = v170
										v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
										if v172 != 0 {
											v200 = v164
											v202 = v9
											*(*int32)(unsafe.Add(mBase, uint32(v200)+12)) = v202
											v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
											if int32(2) <= v205 {
												v208 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
												v211 = F_palloc(m, int32(base.Ui32(v208)>>(uint(int32(2))%32)))
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = v211
													v214 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
													v216 = int32(base.Ui32(v214) >> (uint(int32(2)) % 32))
													if v216 != 0 {
														v217 = F__emscripten_memcpy_bulkmem(m, v211, v24, v216)
														mBase = m.M
													} else {
													}
													v221 = v200
													v224 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
													*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
													v229 = v120 & v224
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
													v232 = v123 & v224
													*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
													v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v234, v221)
													mBase = m.M
													v236 = m.ExcPending
													if v236 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
														v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
														v252 = v249 & int32(16383)
														m.G0 = v16 + int32(80)
														return v252
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v200)+16)) = int32(0)
												v221 = v200
												v224 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
												*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
												v229 = v120 & v224
												*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
												v232 = v123 & v224
												*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
												v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v234, v221)
												mBase = m.M
												v236 = m.ExcPending
												if v236 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
													v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
													v252 = v249 & int32(16383)
													m.G0 = v16 + int32(80)
													return v252
												}
											}
										} else {
											v173 = v164
											*(*int64)(unsafe.Add(mBase, uint32(v173)+12)) = int64(0)
											v221 = v173
											v224 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+35)) = uint8(v224)
											*(*int32)(unsafe.Add(mBase, uint32(v221)+20)) = int32(0)
											v229 = v120 & v224
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+37)) = uint8(v229)
											v232 = v123 & v224
											*(*uint8)(unsafe.Add(mBase, uint32(v221)+36)) = uint8(v232)
											v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
											F_pairingheap_add(m, v234, v221)
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
												v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
												v252 = v249 & int32(16383)
												m.G0 = v16 + int32(80)
												return v252
											}
										}
									}
								}
							} else {
								v181 = int32(0)
								m.T0[l7].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32))(m, l0, v24+int32(6), v121, v5, v24, v123&int32(1), v181, v181)
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return int32(0)
								} else {
									v185 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v185)
									v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
									v252 = v249 & int32(16383)
									m.G0 = v16 + int32(80)
									return v252
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_spg_kd_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 float64
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v14 = F_palloc(m, v11<<(uint(int32(3))%32))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		if int32(0) < v18 {
			v22 = int32(0)
			for {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v22<<(uint(int32(2))%32))))
				v37 = v14 + v22<<(uint(int32(3))%32)
				*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v22
				*(*int32)(unsafe.Add(mBase, uint32(v37))) = v34
				v41 = v22 + int32(1)
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				if v41 < v42 {
					v22 = v41
					continue
				} else {
					break
				}
				break
			}
			v46 = v42
		} else {
			v46 = v18
		}
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		if v55&int32(1) != 0 {
			v58 = int32(246)
		} else {
			v58 = int32(247)
		}
		F_pg_qsort(m, v14, v46, int32(8), v58)
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return int32(0)
		} else {
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v62 = int32(1)
			v63 = v61 >> (uint(v62) % 32)
			v64 = int32(3)
			v67 = *(*int32)(unsafe.Add(mBase, uint32(v14+v63<<(uint(v64)%32))))
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			v76 = *(*float64)(unsafe.Add(mBase, uint32(v67+(v68^int32(-1))<<(uint(v64)%32)&int32(8))))
			*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v62)
			v79 = F_Float8GetDatum(m, v76)
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(2)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v79
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v87 = F_palloc(m, v84<<(uint(int32(2))%32))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v87
					v90 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v93 = F_palloc(m, v90<<(uint(int32(2))%32))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v93
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
						if int32(0) < v96 {
							v100 = int32(0)
							for {
								v110 = v14 + v100<<(uint(int32(3))%32)
								v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
								v112 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
								v114 = v112 << (uint(int32(2)) % 32)
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v114+v115))) = base.B2i32(v63 <= v100)
								v119 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v119+v114))) = v111
								v123 = v100 + int32(1)
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
								if v123 < v124 {
									v100 = v123
									continue
								} else {
									break
								}
								break
							}
						} else {
						}
						return int32(0)
					}
				}
			}
		}
	}
}
func F_spg_quad_choose(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+12)))
	if v7 == int32(1) {
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v6
		v24 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v24
		return v24
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(1)
		v15 = F_getQuadrant(m, v12, v6)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v15 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v6
			v24 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v24
			return v24
		}
	}
}
func F_spg_quad_leaf_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)) = uint8(v2)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v2 < v19 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v112
L2:
	;
	v26 = v2
	goto L5
L3:
	;
	goto L4
L4:
	;
	v101 = int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v102 <= int32(0) {
		v112 = v101
		goto L1
	} else {
		goto L32
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v32 = v29 + v26*int32(48)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+6)))
	switch v34 - int32(1) {
	case 0:
		goto L8
	default:
		goto L9
	case 4:
		goto L14
	case 5:
		goto L13
	case 7:
		goto L10
	case 9, 28:
		goto L12
	case 10, 29:
		goto L11
	}
L6:
	;
	goto L4
L7:
	;
	v91 = v26 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v91 < v92 {
		v26 = v91
		goto L5
	} else {
		goto L31
	}
L8:
	;
	v82 = int32(0)
	v85 = F_DirectFunctionCall2Coll(m, int32(253), v82, v13, v33)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L15
	} else {
		goto L29
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L15
	} else {
		goto L26
	}
L10:
	;
	v59 = int32(0)
	v62 = F_DirectFunctionCall2Coll(m, int32(254), v59, v33, v13)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L15
	} else {
		goto L24
	}
L11:
	;
	v54 = int32(0)
	v57 = F_DirectFunctionCall2Coll(m, int32(248), v54, v13, v33)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L15
	} else {
		goto L22
	}
L12:
	;
	v49 = int32(0)
	v52 = F_DirectFunctionCall2Coll(m, int32(252), v49, v13, v33)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L15
	} else {
		goto L20
	}
L13:
	;
	v44 = int32(0)
	v47 = F_DirectFunctionCall2Coll(m, int32(255), v44, v13, v33)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L15
	} else {
		goto L18
	}
L14:
	;
	v37 = int32(0)
	v40 = F_DirectFunctionCall2Coll(m, int32(250), v37, v13, v33)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	if v40 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v112 = v37
	goto L1
L18:
	;
	if v47 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v112 = v44
	goto L1
L20:
	;
	if v52 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v112 = v49
	goto L1
L22:
	;
	if v57 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v112 = v54
	goto L1
L24:
	;
	if v62 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v112 = v59
	goto L1
L26:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68+v26*int32(48))+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v72
	F_errmsg_internal(m, int32(474138), v10)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(491307), int32(457), int32(91314))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L15
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
	if v85 == int32(0) {
		v112 = v82
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L7
L31:
	;
	goto L6
L32:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v108 = F_spg_key_orderbys_distances(m, v105, int32(1), v107, v102)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v108
	v112 = v101
	goto L1
}
func F_spg_range_quad_choose(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
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
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
		if v17 == int32(1) {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(1)
			v98 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v13
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v98
			m.G0 = v8 + int32(48)
			return int32(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v24 = F_range_get_typcache(m, l0, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+13)))
				if v26 == int32(0) {
					v29 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v29
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(v13+int32(base.Ui32(v32)>>(uint(int32(2))%32))-v29))))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = (v38 ^ int32(-1)) & int32(1)
					v98 = v29
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v13
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v98
					m.G0 = v8 + int32(48)
					return int32(0)
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
					v45 = F_pg_detoast_datum(m, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						F_range_deserialize(m, v24, v45, v8+int32(40), v8+int32(32), v8+int32(31))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							F_range_deserialize(m, v24, v13, v8+int32(20), v8+int32(12), v8+int32(11))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+11)))
								if v64 != 0 {
									v91 = int32(5)
									v92 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v91 - v92
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v92
									v98 = v92
									*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v13
									*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v98
									m.G0 = v8 + int32(48)
									return int32(0)
								} else {
									v69 = F_range_cmp_bounds(m, v24, v8+int32(20), v8+int32(40))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										v75 = F_range_cmp_bounds(m, v24, v8+int32(12), v8+int32(32))
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											if int32(0) <= v75 {
												v81 = int32(1)
											} else {
												v81 = int32(2)
											}
											if int32(0) <= v69 {
												v91 = v81
											} else {
												if int32(0) <= v75 {
													v88 = int32(4)
												} else {
													v88 = int32(3)
												}
												v91 = v88
											}
											v92 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v91 - v92
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = v92
											v98 = v92
											*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v13
											*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v98
											m.G0 = v8 + int32(48)
											return int32(0)
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
}
func F_spg_text_choose(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
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
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	v2 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = int32(1)
	v23 = v16 + v22
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v26 = v24 & v22
	if v24 == v22 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v26 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v29 = int32(4)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v31&int32(254) == int32(2) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v44 = int32(1)
	if v26 != 0 {
		v54 = int32(base.Ui32(v24)>>(uint(v44)%32)) - v44
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v40 = v29
	goto L9
L8:
	;
	v40 = base.B2i32(v31 == int32(18)) << (uint(v29) % 32)
	goto L9
L9:
	;
	if v31 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v43 = v29
	goto L12
L11:
	;
	v43 = v40
	goto L12
L12:
	;
	v54 = v43
	goto L3
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L14:
	;
	v55 = v23
	goto L16
L15:
	;
	v55 = v16 + int32(4)
	goto L16
L16:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+13)))
	if v56 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if int32(0) < v269 {
		goto L83
	} else {
		goto L84
	}
L18:
	;
	if v100 <= v96 {
		v263 = v96
		v268 = int32(65535)
		goto L17
	} else {
		goto L80
	}
L19:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v60 = F_pg_detoast_datum_packed(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v54 <= v236 {
		v263 = v2
		v268 = int32(65535)
		goto L17
	} else {
		goto L79
	}
L22:
	;
	v64 = int32(1)
	v65 = v60 + v64
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	v68 = v66 & v64
	if v66 == v64 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v68 != 0 {
		goto L34
	} else {
		goto L35
	}
L24:
	;
	v71 = int32(4)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v73&int32(254) == int32(2) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v86 = int32(1)
	if v68 != 0 {
		v96 = int32(base.Ui32(v66)>>(uint(v86)%32)) - v86
		goto L23
	} else {
		goto L33
	}
L27:
	;
	v82 = v71
	goto L29
L28:
	;
	v82 = base.B2i32(v73 == int32(18)) << (uint(v71) % 32)
	goto L29
L29:
	;
	if v73 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v85 = v71
	goto L32
L31:
	;
	v85 = v82
	goto L32
L32:
	;
	v96 = v85
	goto L23
L33:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v96 = int32(base.Ui32(v90)>>(uint(int32(2))%32)) - int32(4)
	goto L23
L34:
	;
	v97 = v65
	goto L36
L35:
	;
	v97 = v60 + int32(4)
	goto L36
L36:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v99 = v55 + v98
	v100 = v54 - v98
	if v100 < v96 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(1)
	v191 = F_palloc(m, int32(4))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L64
	}
L38:
	;
	v153 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)) = uint8(v153)
	v157 = v129 + int32(4)
	v158 = F_palloc(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L55
	}
L39:
	;
	v150 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)) = uint8(v150)
	v177 = v150
	goto L37
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(3)
	if v129 != 0 {
		goto L38
	} else {
		goto L54
	}
L41:
	;
	v102 = v100
	goto L43
L42:
	;
	v102 = v96
	goto L43
L43:
	;
	if int32(0) < v102 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v106 = v99
	v107 = int32(0)
	v108 = v97
	goto L48
L45:
	;
	goto L46
L46:
	;
	if v96 == int32(0) {
		goto L18
	} else {
		goto L53
	}
L47:
	;
	if v129 != v96 {
		goto L40
	} else {
		goto L52
	}
L48:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v118 != v119 {
		v129 = v107
		goto L47
	} else {
		goto L50
	}
L49:
	;
	v129 = v102
	goto L47
L50:
	;
	v121 = int32(1)
	v126 = v107 + v121
	if v126 != v102 {
		v106 = v106 + v121
		v107 = v126
		v108 = v108 + v121
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	goto L18
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(3)
	goto L39
L54:
	;
	goto L39
L55:
	;
	if base.Ui32(v129) <= base.Ui32(int32(126)) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v129 != 0 {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v165 = v129<<(uint(int32(1))%32) + int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v165)
	v171 = v153
	goto L56
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v157 << (uint(int32(2)) % 32)
	v171 = int32(4)
	goto L56
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v158
	v177 = v129
	goto L37
L61:
	;
	v173 = F__emscripten_memcpy_bulkmem(m, v158+v171, v97, v129)
	mBase = m.M
	goto L63
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v191
	v194 = v177 + v97
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v195
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(0)
	v199 = v96 - v177
	if v199 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v202 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)) = uint8(v202)
	return v202
L66:
	;
	goto L67
L67:
	;
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)) = uint8(v206)
	v209 = v199 - v206
	v211 = v199 + int32(3)
	v212 = F_palloc(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	if base.Ui32(v199) <= base.Ui32(int32(127)) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v212
	return int32(0)
L70:
	;
	if v209 != 0 {
		goto L76
	} else {
		goto L77
	}
L71:
	;
	v216 = int32(1)
	v219 = v199<<(uint(v216)%32) | v216
	*(*uint8)(unsafe.Add(mBase, uint32(v212))) = uint8(v219)
	if v209 != 0 {
		v226 = v216
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = v211 << (uint(int32(2)) % 32)
	v226 = int32(4)
	goto L70
L74:
	;
	goto L69
L75:
	;
	goto L69
L76:
	;
	v230 = F__emscripten_memcpy_bulkmem(m, v226+v212, v194+int32(1), v209)
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
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236+v55))))
	v263 = v2
	v268 = v239
	goto L17
L80:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+v99))))
	v263 = v96
	v268 = v255
	goto L17
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = base.I32_extend16_s(v268)
	return int32(0)
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v289
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(1)
	v340 = int32(0)
	v342 = v263 + base.B2i32(v340 <= v286)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v342
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v346 = v54 - (v342 + v344)
	if v340 < v346 {
		goto L96
	} else {
		goto L97
	}
L83:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v274 = v269
	v276 = int32(0)
	goto L86
L84:
	;
	v302 = v269
	goto L85
L85:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+12)))
	if v314 != int32(1) {
		goto L81
	} else {
		goto L94
	}
L86:
	;
	v286 = base.I32_extend16_s(v268)
	v289 = (v274 + v276) >> (uint(int32(1)) % 32)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v272+v289<<(uint(int32(2))%32))))
	v294 = base.I32_extend16_s(v293)
	if v286 < v294 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	v302 = v299
	goto L85
L88:
	;
	if v300 < v299 {
		v274 = v299
		v276 = v300
		goto L86
	} else {
		goto L93
	}
L89:
	;
	v299 = v289
	v300 = v276
	goto L88
L90:
	;
	goto L91
L91:
	;
	if v286 <= v294 {
		goto L82
	} else {
		goto L92
	}
L92:
	;
	v299 = v274
	v300 = v289 + int32(1)
	goto L88
L93:
	;
	goto L87
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(3)
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)) = uint8(v319)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v321
	v326 = F_palloc(m, int32(4))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v326
	*(*int32)(unsafe.Add(mBase, uint32(v326))) = int32(-2)
	v331 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)) = uint8(v331)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v331
	return v331
L96:
	;
	v351 = int32(4)
	v353 = v346 + v351
	v354 = F_palloc(m, v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v375 = F_palloc(m, int32(4))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L108
	}
L99:
	;
	if base.Ui32(v346) <= base.Ui32(int32(126)) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	if v346 != 0 {
		goto L105
	} else {
		goto L106
	}
L101:
	;
	v358 = int32(1)
	v362 = v346<<(uint(v358)%32) + int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v354))) = uint8(v362)
	v367 = v358
	goto L100
L102:
	;
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v354))) = v353 << (uint(int32(2)) % 32)
	v367 = v351
	goto L100
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v354
	return int32(0)
L105:
	;
	v369 = F__emscripten_memcpy_bulkmem(m, v367+v354, v344+v55+v342, v346)
	mBase = m.M
	goto L107
L106:
	;
	goto L107
L107:
	;
	goto L104
L108:
	;
	v377 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v375))) = uint8(v377)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v375
	return int32(0)
}
func F_spg_text_leaf_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
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
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
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
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)) = uint8(v2)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v24 = F_pg_detoast_datum_packed(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v60 = v58 + v19
	if v58 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	return int32(0)
L3:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v28 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v31 = int32(1)
	if v28&v31 != 0 {
		v58 = int32(base.Ui32(v28)>>(uint(v31)%32)) - v31
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v42 = int32(4)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v44&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v58 = int32(base.Ui32(v37)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L8:
	;
	v53 = v42
	goto L10
L9:
	;
	v53 = base.B2i32(v44 == int32(18)) << (uint(v42) % 32)
	goto L10
L10:
	;
	if v44 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v56 = v42
	goto L13
L12:
	;
	v56 = v53
	goto L13
L13:
	;
	v58 = v56
	goto L1
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v139
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v145 <= int32(0) {
		v346 = int32(1)
		goto L47
	} else {
		goto L48
	}
L15:
	;
	v66 = v60 + int32(4)
	v67 = F_palloc(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L18
	}
L16:
	;
	if v19 <= int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v139 = v59
	v142 = v59 + int32(4)
	goto L14
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v66 << (uint(int32(2)) % 32)
	v73 = v67 + int32(4)
	if v19 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v19 != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L21
L21:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v78 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	goto L21
L23:
	;
	v76 = F__emscripten_memcpy_bulkmem(m, v73, v59+int32(4), v19)
	mBase = m.M
	goto L25
L24:
	;
	goto L25
L25:
	;
	goto L22
L26:
	;
	if v134 != 0 {
		goto L44
	} else {
		goto L45
	}
L27:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if base.Ui32(int32(18)) < base.Ui32(v81) {
		v139 = v67
		v142 = v73
		goto L14
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v78&int32(1) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L30:
	;
	if int32(1)<<(uint(v81)%32)&int32(262158) == int32(0) {
		v139 = v67
		v142 = v73
		goto L14
	} else {
		goto L31
	}
L31:
	;
	v92 = int32(4)
	if v81&int32(254) == int32(2) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v102 = v92
	goto L34
L33:
	;
	v102 = base.B2i32(v81 == int32(18)) << (uint(v92) % 32)
	goto L34
L34:
	;
	if v81 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v105 = v92
	goto L37
L36:
	;
	v105 = v102
	goto L37
L37:
	;
	v133 = v24 + int32(1)
	v134 = v105
	goto L26
L38:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v110&int32(-4) == int32(16) {
		v139 = v67
		v142 = v73
		goto L14
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v78&int32(254) == int32(2) {
		v139 = v67
		v142 = v73
		goto L14
	} else {
		goto L42
	}
L41:
	;
	v115 = int32(4)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v133 = v24 + v115
	v134 = int32(base.Ui32(v117)>>(uint(int32(2))%32)) - v115
	goto L26
L42:
	;
	v126 = int32(1)
	v133 = v24 + v126
	v134 = int32(base.Ui32(v78)>>(uint(v126)%32)) - v126
	goto L26
L43:
	;
	v139 = v67
	v142 = v73
	goto L14
L44:
	;
	v136 = F__emscripten_memcpy_bulkmem(m, v19+v73, v133, v134)
	mBase = m.M
	goto L46
L45:
	;
	goto L46
L46:
	;
	goto L43
L47:
	;
	m.G0 = v16 + int32(16)
	return v346
L48:
	;
	v152 = int32(0)
	goto L49
L49:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v165 = v162 + v152*int32(48)
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+6)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v165)+44))
	v168 = F_pg_detoast_datum_packed(m, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L52
	}
L50:
	;
	v346 = v339
	goto L47
L51:
	;
	v202 = v166 & int32(65535)
	if v202 == int32(28) {
		goto L64
	} else {
		goto L65
	}
L52:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v170 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v173 = int32(4)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	if v175&int32(254) == int32(2) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v188 = int32(1)
	if v170&v188 != 0 {
		v200 = int32(base.Ui32(v170)>>(uint(v188)%32)) - v188
		goto L51
	} else {
		goto L62
	}
L56:
	;
	v184 = v173
	goto L58
L57:
	;
	v184 = base.B2i32(v175 == int32(18)) << (uint(v173) % 32)
	goto L58
L58:
	;
	if v175 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v187 = v173
	goto L61
L60:
	;
	v187 = v184
	goto L61
L61:
	;
	v200 = v187
	goto L51
L62:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v200 = int32(base.Ui32(v194)>>(uint(int32(2))%32)) - int32(4)
	goto L51
L63:
	;
	v339 = int32(1)
	v341 = v152 + v339
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v341 < v342 {
		v152 = v341
		goto L49
	} else {
		goto L117
	}
L64:
	;
	if v200 <= v19 {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	if base.Ui32(int32(11)) <= base.Ui32(v202) {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v209 = F_DirectFunctionCall2Coll(m, int32(260), v207, v208, v168)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	if v209 != 0 {
		goto L63
	} else {
		goto L69
	}
L69:
	;
	v346 = int32(0)
	goto L47
L70:
	;
	switch v299&int32(65535) - int32(1) {
	case 0:
		goto L104
	case 1:
		goto L108
	case 2:
		goto L107
	case 3:
		goto L106
	case 4:
		goto L103
	default:
		goto L105
	}
L71:
	;
	v216 = int32(1)
	if v170&v216 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	v225 = int32(1)
	if v170&v225 != 0 {
		goto L78
	} else {
		goto L79
	}
L74:
	;
	v220 = v216
	goto L76
L75:
	;
	v220 = int32(4)
	goto L76
L76:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v223 = F_varstr_cmp(m, v142, v60, v168+v220, v200, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	v298 = v223
	v299 = v166 - int32(10)
	goto L70
L78:
	;
	v229 = v225
	goto L80
L79:
	;
	v229 = int32(4)
	goto L80
L80:
	;
	v230 = v168 + v229
	if v200 < v60 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v232 = v200
	goto L83
L82:
	;
	v232 = v60
	goto L83
L83:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v232) {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	if v294 != 0 {
		v298 = v294
		v299 = v166
		goto L70
	} else {
		goto L102
	}
L85:
	;
	v294 = int32(0)
	goto L84
L86:
	;
	v268 = v263
	v269 = v264
	v270 = v265
	goto L96
L87:
	;
	if (v142|v230)&int32(3) != 0 {
		v263 = v142
		v264 = v230
		v265 = v232
		goto L86
	} else {
		goto L90
	}
L88:
	;
	v256 = v142
	v257 = v230
	v258 = v232
	goto L89
L89:
	;
	if v258 == int32(0) {
		goto L85
	} else {
		goto L95
	}
L90:
	;
	v240 = v142
	v241 = v230
	v242 = v232
	goto L91
L91:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	if v245 != v246 {
		v263 = v240
		v264 = v241
		v265 = v242
		goto L86
	} else {
		goto L93
	}
L92:
	;
	v256 = v251
	v257 = v249
	v258 = v253
	goto L89
L93:
	;
	v248 = int32(4)
	v249 = v241 + v248
	v251 = v240 + v248
	v253 = v242 - v248
	if base.Ui32(int32(3)) < base.Ui32(v253) {
		v240 = v251
		v241 = v249
		v242 = v253
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v263 = v256
	v264 = v257
	v265 = v258
	goto L86
L96:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	if v273 == v274 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v294 = v273 - v274
	goto L84
L98:
	;
	v276 = int32(1)
	v281 = v270 - v276
	if v281 != 0 {
		v268 = v268 + v276
		v269 = v269 + v276
		v270 = v281
		goto L96
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	goto L97
L101:
	;
	goto L85
L102:
	;
	v298 = base.B2i32(v200 < v60) - base.B2i32(v60 < v200)
	v299 = v166
	goto L70
L103:
	;
	v334 = int32(0)
	if v298 <= v334 {
		v346 = v334
		goto L47
	} else {
		goto L116
	}
L104:
	;
	if v298 < int32(0) {
		goto L63
	} else {
		goto L115
	}
L105:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L2
	} else {
		goto L112
	}
L106:
	;
	if int32(0) <= v298 {
		goto L63
	} else {
		goto L111
	}
L107:
	;
	if v298 == int32(0) {
		goto L63
	} else {
		goto L110
	}
L108:
	;
	if v298 <= int32(0) {
		goto L63
	} else {
		goto L109
	}
L109:
	;
	v346 = int32(0)
	goto L47
L110:
	;
	v346 = int32(0)
	goto L47
L111:
	;
	v346 = int32(0)
	goto L47
L112:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v317+v152*int32(48))+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v321
	F_errmsg_internal(m, int32(474138), v16)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(491256), int32(691), int32(91235))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	v346 = int32(0)
	goto L47
L116:
	;
	goto L63
L117:
	;
	goto L50
}
func F_spg_xlog_cleanup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	F_MemoryContextDelete(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[139])) = int32(0)
		return
	}
}
