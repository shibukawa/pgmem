package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_spgFormInnerTuple(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) int32 {
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
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	if l1 != 0 {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
		if v20 != 0 {
			v49 = int32(8)
		} else {
			v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+48)))
			if int32(0) < v21 {
				v49 = v21
			} else {
				v24 = base.I32_wrap_i64(l2)
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				if v25 == int32(1) {
					v29 = int32(18)
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
					if v31 == v29 {
						v34 = v29
					} else {
						v34 = int32(2)
					}
					if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v41 = int32(6)
					} else {
						v41 = v34
					}
					v49 = v41
				} else {
					if v25&int32(1) != 0 {
						v49 = int32(base.Ui32(v25) >> (uint(int32(1)) % 32))
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
						v49 = int32(base.Ui32(v46) >> (uint(int32(2)) % 32))
					}
				}
			}
		}
		v57 = (v49 + int32(7)) & int32(-8)
	} else {
		v57 = v6
	}
	v59 = v57 + int32(8)
	if l3 <= int32(0) {
		v161 = v59
	} else {
		v63 = l3 & int32(3)
		if base.Ui32(l3) < base.Ui32(int32(4)) {
			v119 = int32(0)
			v120 = v59
			v133 = v119
			v134 = v120
			v139 = v6
			for {
				v145 = *(*int32)(unsafe.Add(mBase, uint32(l4+v133<<(uint(int32(2))%32))))
				v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+6)))
				v149 = v134 + v146&int32(_a_F_spgFormInnerTuple_0)
				v150 = int32(1)
				v153 = v139 + v150
				if v153 != v63 {
					v133 = v133 + v150
					v134 = v149
					v139 = v153
					continue
				} else {
					break
				}
				break
			}
			v161 = v149
		} else {
			v75 = int32(0)
			v76 = v59
			v83 = v6
			for {
				v86 = l4 + v75<<(uint(int32(2))%32)
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
				v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+6)))
				v89 = int32(_a_F_spgFormInnerTuple_0)
				v92 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
				v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92)+6)))
				v97 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
				v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97)+6)))
				v102 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
				v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+6)))
				v106 = v76 + v88&v89 + v93&v89 + v98&v89 + v103&v89
				v107 = int32(4)
				v108 = v75 + v107
				v110 = v83 + v107
				if v110 != l3&int32(2147483644) {
					v75 = v108
					v76 = v106
					v83 = v110
					continue
				} else {
					break
				}
				break
			}
			if v63 == int32(0) {
				v161 = v106
			} else {
				v119 = v108
				v120 = v106
				v133 = v119
				v134 = v120
				v139 = v6
				for {
					v145 = *(*int32)(unsafe.Add(mBase, uint32(l4+v133<<(uint(int32(2))%32))))
					v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+6)))
					v149 = v134 + v146&int32(_a_F_spgFormInnerTuple_0)
					v150 = int32(1)
					v153 = v139 + v150
					if v153 != v63 {
						v133 = v133 + v150
						v134 = v149
						v139 = v153
						continue
					} else {
						break
					}
					break
				}
				v161 = v149
			}
		}
	}
	v169 = int32(16)
	if base.Ui32(v161) <= base.Ui32(v169) {
		v172 = v169
	} else {
		v172 = v161
	}
	if base.Ui32(v161) < base.Ui32(int32(_a_F_spgFormInnerTuple_1)) {
		if base.B2i32(int32(_a_F_spgFormInnerTuple_0) < l3)|base.B2i32(base.Ui32(int32(_a_F_spgFormInnerTuple_2)) <= base.Ui32(v57)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v368 = m.ExcPending
			if v368 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_spgFormInnerTuple_3), int32(0))
				mBase = m.M
				v372 = m.ExcPending
				if v372 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_spgFormInnerTuple_4), int32(1047), int32(_a_F_spgFormInnerTuple_5))
					mBase = m.M
					v377 = m.ExcPending
					if v377 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v180 = F_palloc0(m, v172)
			mBase = m.M
			v183 = m.ExcPending
			if v183 != 0 {
				return int32(0)
			} else {
				*(*uint16)(unsafe.Add(mBase, uint32(v180)+4)) = uint16(v172)
				v189 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
				*(*int32)(unsafe.Add(mBase, uint32(v180))) = l3<<(uint(int32(3))%32)&int32(_a_F_spgFormInnerTuple_6) | v189&int32(7) | v57<<(uint(int32(16))%32)
				if l1 == int32(0) {
				} else {
					v200 = v180 + int32(8)
					v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
					if v201 == int32(1) {
						*(*int64)(unsafe.Add(mBase, uint32(v200))) = l2
					} else {
						v205 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+48)))
						if int32(0) < v205 {
							v234 = base.I32_wrap_i64(l2)
							v236 = v205
						} else {
							v209 = base.I32_wrap_i64(l2)
							v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
							if v210 == int32(1) {
								v214 = int32(18)
								v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
								if v216 == v214 {
									v219 = v214
								} else {
									v219 = int32(2)
								}
								if base.Ui32((v216-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v226 = int32(6)
								} else {
									v226 = v219
								}
								v234 = v209
								v236 = v226
							} else {
								if v210&int32(1) != 0 {
									v234 = v209
									v236 = int32(base.Ui32(v210) >> (uint(int32(1)) % 32))
								} else {
									v231 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
									v234 = v209
									v236 = int32(base.Ui32(v231) >> (uint(int32(2)) % 32))
								}
							}
						}
						if v236 == int32(0) {
						} else {
							if v57 != 0 {
								v240 = v200
							} else {
								v240 = int32(0)
							}
							base.MemoryCopy(m, v240, v234, v236)
						}
					}
				}
				if l3 <= int32(0) {
				} else {
					v250 = v180 + v57 + int32(8)
					v251 = int32(0)
					if l3 != int32(1) {
						v264 = v251
						v265 = v250
						v270 = int32(0)
						for {
							v275 = l4 + v264<<(uint(int32(2))%32)
							v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
							v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v276)+6)))
							v279 = v277 & int32(_a_F_spgFormInnerTuple_0)
							if v279 != 0 {
								base.MemoryCopy(m, v265, v276, v279)
							} else {
							}
							v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v276)+6)))
							v282 = int32(_a_F_spgFormInnerTuple_0)
							v284 = v265 + v281&v282
							v285 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
							v286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285)+6)))
							v288 = v286 & v282
							if v288 != 0 {
								base.MemoryCopy(m, v284, v285, v288)
							} else {
							}
							v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285)+6)))
							v293 = v284 + v290&int32(_a_F_spgFormInnerTuple_0)
							v294 = int32(2)
							v295 = v264 + v294
							v297 = v270 + v294
							if v297 != l3&int32(2147483646) {
								v264 = v295
								v265 = v293
								v270 = v297
								continue
							} else {
								break
							}
							break
						}
						if l3&int32(1) == int32(0) {
						} else {
							v306 = v295
							v307 = v293
							v318 = *(*int32)(unsafe.Add(mBase, uint32(l4+v306<<(uint(int32(2))%32))))
							v319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v318)+6)))
							v321 = v319 & int32(_a_F_spgFormInnerTuple_0)
							if v321 == int32(0) {
							} else {
								base.MemoryCopy(m, v307, v318, v321)
							}
						}
					} else {
						v306 = v251
						v307 = v250
						v318 = *(*int32)(unsafe.Add(mBase, uint32(l4+v306<<(uint(int32(2))%32))))
						v319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v318)+6)))
						v321 = v319 & int32(_a_F_spgFormInnerTuple_0)
						if v321 == int32(0) {
						} else {
							base.MemoryCopy(m, v307, v318, v321)
						}
					}
				}
				m.G0 = v17 + int32(16)
				return v180
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v346 = m.ExcPending
		if v346 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v349 = m.ExcPending
			if v349 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = int32(_a_F_spgFormInnerTuple_7)
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = v172
				F_errmsg(m, int32(_a_F_spgFormInnerTuple_8), v17)
				mBase = m.M
				v355 = m.ExcPending
				if v355 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(_a_F_spgFormInnerTuple_9), int32(0))
					mBase = m.M
					v359 = m.ExcPending
					if v359 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_spgFormInnerTuple_4), int32(1038), int32(_a_F_spgFormInnerTuple_5))
						mBase = m.M
						v364 = m.ExcPending
						if v364 != 0 {
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+14)) = uint16(v5)
	if v16 < int32(2) {
		v41 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v48 = F_heap_compute_data_size(m, v15, l2, l3)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v22 = v5
	goto L3
L3:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+l3))))
	if v33 != 0 {
		v41 = v33
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v41 = v33
	goto L1
L5:
	;
	v35 = v22 + int32(1)
	if v35 != v16 {
		v22 = v35
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
L7:
	;
	return int32(0)
L8:
	;
	v53 = v48 + int32(23)
	if base.Ui32(v53) <= base.Ui32(int32(16)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v56 = int32(16)
	goto L11
L10:
	;
	v56 = v53
	goto L11
L11:
	;
	v59 = F_palloc0(m, v56&int32(-8))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+4)))
	v63 = v61 & int32(-16384)
	*(*uint16)(unsafe.Add(mBase, uint32(v59)+4)) = uint16(v63)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v56<<(uint(int32(2))%32)&int32(-32) | v69&int32(3)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v59)+6)) = v74
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v59)+10)) = uint16(v76)
	if v41 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	m.G0 = v13 + int32(16)
	return v59
L14:
	;
	F_heap_fill_tuple(m, v15, l2, l3, v59+int32(16), v13+int32(14), v87)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L7
	} else {
		goto L20
	}
L15:
	;
	v79 = v63 | int32(_a_F_spgFormLeafTuple_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v59)+4)) = uint16(v79)
	v87 = v59 + int32(12)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v83 = int32(0)
	if int32(1) < v16 {
		v87 = v83
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v86 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v87 = v83
	goto L14
L20:
	;
	goto L13
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	var v103 int32
	_ = v103
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int64
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int64
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int64
	_ = v254
	var v256 int64
	_ = v256
	var v258 int64
	_ = v258
	var v260 int64
	_ = v260
	var v262 int64
	_ = v262
	var v264 int64
	_ = v264
	var v266 int64
	_ = v266
	var v268 int64
	_ = v268
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
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
	v303 = m.ExcPending
	if v303 != 0 {
		goto L6
	} else {
		goto L85
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L6
	} else {
		goto L81
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
	v276 = v12
	goto L5
L5:
	;
	m.G0 = v10 + int32(16)
	return v276
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v147
	v152 = int32(1)
	v154 = F_index_getprocinfo(m, l0, v152, v152)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L6
	} else {
		goto L53
	}
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+48)))
	if v44 != 0 {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	switch v22 - int32(2277) {
	case 0, 6:
		goto L9
	case 1, 2, 3, 4, 5:
		v147 = v22
		goto L8
	default:
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if base.B2i32(base.Ui32(v22-int32(_a_F_spgGetCache_0)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v22-int32(_a_F_spgGetCache_1)) < base.Ui32(int32(2))) != 0 {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	if base.B2i32(v22 == int32(2776))|base.B2i32(v22 == int32(3500)) != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v147 = v22
	goto L8
L15:
	;
	if v22 != int32(3831) {
		v147 = v22
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L9
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v46 = F_get_atttype(m, v45, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	if v50 != 0 {
		v55 = v50
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v48 = F_getBaseType(m, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v147 = v48
	goto L8
L22:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59)+10)))
	if v60 <= int32(0) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	v57 = v55
	v58 = v56
	goto L22
L24:
	;
	v51 = F_RelationGetIndexExpressions(m, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	if v51 != 0 {
		v55 = v51
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v53 = int32(0)
	v57 = v53
	v58 = v53
	goto L22
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L6
	} else {
		goto L50
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L47
	}
L29:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+48)))
	if v63 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v58 == int32(0) {
		goto L27
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v60 == int32(1) {
		goto L28
	} else {
		goto L36
	}
L33:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v69 = F_exprType(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v71 = F_getBaseType(m, v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v147 = v71
	goto L8
L36:
	;
	v79 = v58
	v80 = int32(2)
	goto L37
L37:
	;
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59+int32(46)+v80<<(uint(int32(1))%32)))))
	if v88 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L28
L39:
	;
	if v79 == int32(0) {
		goto L27
	} else {
		goto L42
	}
L40:
	;
	v103 = v79
	goto L41
L41:
	;
	if v80 != v60 {
		v79 = v103
		v80 = v80 + int32(1)
		goto L37
	} else {
		goto L46
	}
L42:
	;
	v94 = v79 + int32(4)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if base.Ui32(v94) < base.Ui32(v96+v97<<(uint(int32(2))%32)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v102 = v94
	goto L45
L44:
	;
	v102 = int32(0)
	goto L45
L45:
	;
	v103 = v102
	goto L41
L46:
	;
	goto L38
L47:
	;
	F_errmsg_internal(m, int32(_a_F_spgGetCache_2), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_spgGetCache_3), int32(161), int32(_a_F_spgGetCache_4))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	F_errmsg_internal(m, int32(_a_F_spgGetCache_2), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L6
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_spgGetCache_3), int32(155), int32(_a_F_spgGetCache_4))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v162 = F_FunctionCall2Coll(m, v154, v157, base.I64_extend_i32_u(v10+int32(12)), base.I64_extend_i32_u(v17))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v164 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v180 = v17 + int32(16)
	F_fillTypeDesc(m, v180, v147)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L6
	} else {
		goto L60
	}
L56:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v165+v166<<(uint(int32(3))%32))+96))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v170
	if v147 == v170 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v173 = F_IsBinaryCoercible(m, v170, v147)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	if v173 == int32(0) {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v147
	goto L55
L60:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v147 != v183 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	F_fillTypeDesc(m, v17+int32(40), v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L6
	} else {
		goto L68
	}
L62:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188)+6)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v187+v189*int32(0)<<(uint(int32(2))%32)+int32(24)-int32(4))))
	goto L65
L63:
	;
	goto L64
L64:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v209
	v211 = *(*int64)(unsafe.Add(mBase, uint32(v180)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+28)) = v211
	goto L61
L65:
	;
	if v201 == int32(0) {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_fillTypeDesc(m, v17+int32(28), v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	goto L61
L68:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	F_fillTypeDesc(m, v17+int32(52), v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+119)))
	if v224 != int32(73) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v228 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L6
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v17
	v276 = v17
	goto L5
L73:
	;
	F_LockBufferInternal(m, v228, int32(1))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	if v228 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+24))
	if v251 != int32(-1173640210) {
		goto L1
	} else {
		goto L79
	}
L76:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_spgGetCache[0]))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v236+(v228^int32(-1))<<(uint(int32(2))%32))))
	v250 = v242
	goto L75
L77:
	;
	goto L78
L78:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_spgGetCache[1]))
	v250 = v244 + v228<<(uint(int32(13))%32) + int32(-8192)
	goto L75
L79:
	;
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v250)+84))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = v254
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v250)+76))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = v256
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v250)+68))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+104)) = v258
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v250)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+96)) = v260
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v250)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+88)) = v262
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v250)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = v264
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v250)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+72)) = v266
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v250)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = v268
	F_UnlockReleaseBuffer(m, v228)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	goto L72
L81:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(_a_F_spgGetCache_5), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_spgGetCache_3), int32(252), int32(_a_F_spgGetCache_6))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v304 + int32(4)
	F_errmsg_internal(m, int32(_a_F_spgGetCache_7), v10)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_spgGetCache_3), int32(281), int32(_a_F_spgGetCache_6))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int64
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int64
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int64
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	v5 = l4
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2+l3<<(uint(int32(2))%32))+20))
	v23 = l2 + v20&int32(_a_F_spgTestLeafTuple_0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v26 = v24 & int32(3)
	if v26 != 0 {
		if l5 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v42 & int32(3)
				F_errmsg_internal(m, int32(_a_F_spgTestLeafTuple_1), v15+int32(16))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_spgTestLeafTuple_2), int32(793), int32(_a_F_spgTestLeafTuple_3))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			switch v26 - int32(1) {
			case 0:
				v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)))
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+40)) = uint16(v30)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+6))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v32
				v262 = int32(2049)
				m.G0 = v15 + int32(96)
				return v262
			case 1:
				v262 = int32(0)
				m.G0 = v15 + int32(96)
				return v262
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v42 & int32(3)
					F_errmsg_internal(m, int32(_a_F_spgTestLeafTuple_1), v15+int32(16))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_spgTestLeafTuple_2), int32(793), int32(_a_F_spgTestLeafTuple_3))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
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
			v56 = int32(0)
			v136 = v56
			v137 = int32(0)
			v138 = int64(0)
			v139 = v56
			v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
			if int32(0) < v140 {
				v143 = int32(_a_F_spgTestLeafTuple_4)
				v144 = *(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0]))
				v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v146
				v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
				if v5 == int32(0) {
					v155 = F_palloc(m, v140<<(uint(int32(3))%32)+int32(48))
					mBase = m.M
					v156 = m.ExcPending
					if v156 != 0 {
						return int32(0)
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v155)+42)) = uint8(v5)
						v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
						if v158 <= int32(0) {
						} else {
							v162 = v158 << (uint(int32(3)) % 32)
							if v162 == int32(0) {
							} else {
								base.MemoryCopy(m, v155+int32(48), v137, v162)
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v155)+32)) = v148
						v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)))
						*(*uint16)(unsafe.Add(mBase, uint32(v155)+40)) = uint16(v170)
						v172 = *(*int32)(unsafe.Add(mBase, uint32(v23)+6))
						*(*int32)(unsafe.Add(mBase, uint32(v155)+36)) = v172
						v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
						if v174 == int32(0) {
							v192 = v155
							*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v192)+16)) = int64(0)
							v231 = v192
							v235 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
							*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
							v240 = v136 & v235
							*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
							v243 = v139 & v235
							*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
							v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							F_pairingheap_add(m, v245, v231)
							mBase = m.M
							v247 = m.ExcPending
							if v247 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
								v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
								v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
								m.G0 = v15 + int32(96)
								return v262
							}
						} else {
							v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
							v178 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
							v179 = F_datumCopy(m, v138, v177, v178)
							mBase = m.M
							v180 = m.ExcPending
							if v180 != 0 {
								return int32(0)
							} else {
								v209 = v155
								v211 = v179
								*(*int64)(unsafe.Add(mBase, uint32(v209)+16)) = v211
								v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
								v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
								if int32(2) <= v214 {
									v217 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
									v220 = F_palloc(m, int32(base.Ui32(v217)>>(uint(int32(2))%32)))
									mBase = m.M
									v221 = m.ExcPending
									if v221 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = v220
										v223 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
										v225 = int32(base.Ui32(v223) >> (uint(int32(2)) % 32))
										if v225 == int32(0) {
											v231 = v209
										} else {
											base.MemoryCopy(m, v220, v23, v225)
											v231 = v209
										}
										v235 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
										*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
										v240 = v136 & v235
										*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
										v243 = v139 & v235
										*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
										v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
										F_pairingheap_add(m, v245, v231)
										mBase = m.M
										v247 = m.ExcPending
										if v247 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
											v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
											v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
											m.G0 = v15 + int32(96)
											return v262
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = int32(0)
									v231 = v209
									v235 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
									*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
									v240 = v136 & v235
									*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
									v243 = v139 & v235
									*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
									v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
									F_pairingheap_add(m, v245, v231)
									mBase = m.M
									v247 = m.ExcPending
									if v247 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
										v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
										v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
										m.G0 = v15 + int32(96)
										return v262
									}
								}
							}
						}
					}
				} else {
					v182 = F_palloc(m, int32(48))
					mBase = m.M
					v183 = m.ExcPending
					if v183 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v182)+32)) = v148
						*(*uint8)(unsafe.Add(mBase, uint32(v182)+42)) = uint8(v5)
						v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)))
						*(*uint16)(unsafe.Add(mBase, uint32(v182)+40)) = uint16(v186)
						v188 = *(*int32)(unsafe.Add(mBase, uint32(v23)+6))
						*(*int32)(unsafe.Add(mBase, uint32(v182)+36)) = v188
						v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
						if v191 != 0 {
							v209 = v182
							v211 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v209)+16)) = v211
							v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
							v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
							if int32(2) <= v214 {
								v217 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
								v220 = F_palloc(m, int32(base.Ui32(v217)>>(uint(int32(2))%32)))
								mBase = m.M
								v221 = m.ExcPending
								if v221 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = v220
									v223 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
									v225 = int32(base.Ui32(v223) >> (uint(int32(2)) % 32))
									if v225 == int32(0) {
										v231 = v209
									} else {
										base.MemoryCopy(m, v220, v23, v225)
										v231 = v209
									}
									v235 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
									*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
									v240 = v136 & v235
									*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
									v243 = v139 & v235
									*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
									v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
									F_pairingheap_add(m, v245, v231)
									mBase = m.M
									v247 = m.ExcPending
									if v247 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
										v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
										v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
										m.G0 = v15 + int32(96)
										return v262
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = int32(0)
								v231 = v209
								v235 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
								*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
								v240 = v136 & v235
								*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
								v243 = v139 & v235
								*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
								v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
								F_pairingheap_add(m, v245, v231)
								mBase = m.M
								v247 = m.ExcPending
								if v247 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
									v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
									v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
									m.G0 = v15 + int32(96)
									return v262
								}
							}
						} else {
							v192 = v182
							*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v192)+16)) = int64(0)
							v231 = v192
							v235 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
							*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
							v240 = v136 & v235
							*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
							v243 = v139 & v235
							*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
							v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							F_pairingheap_add(m, v245, v231)
							mBase = m.M
							v247 = m.ExcPending
							if v247 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
								v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
								v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
								m.G0 = v15 + int32(96)
								return v262
							}
						}
					}
				}
			} else {
				v203 = int32(0)
				m.T0[l7].(func(*base.Module, int32, int32, int64, int32, int32, int32, int32, int32))(m, l0, v23+int32(6), v138, v5, v23, v139&int32(1), v203, v203)
				mBase = m.M
				v206 = m.ExcPending
				if v206 != 0 {
					return int32(0)
				} else {
					v207 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v207)
					v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
					v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
					m.G0 = v15 + int32(96)
					return v262
				}
			}
		} else {
			v58 = int32(_a_F_spgTestLeafTuple_4)
			v59 = *(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0]))
			v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
			*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v61
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v63
			v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v65
			v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v67
			v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v69
			v71 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v15)+64)) = v71
			v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+72)) = v73
			v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+76)) = v75
			v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
			*(*uint8)(unsafe.Add(mBase, uint32(v15)+80)) = uint8(v77)
			v80 = v23 + int32(16)
			v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
			if v81 == int32(1) {
				v84 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+36)))
				if base.I32_popcnt(v84) != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v84
						F_errmsg_internal(m, int32(_a_F_spgTestLeafTuple_6), v15)
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_spgTestLeafTuple_7), int32(123), int32(_a_F_spgTestLeafTuple_8))
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					switch base.I32_ctz(v84) {
					case 0:
						v89 = int64(*(*int8)(unsafe.Add(mBase, uint32(v80))))
						v108 = v89
						v109 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v109
						*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = v108
						*(*uint16)(unsafe.Add(mBase, uint32(v15)+40)) = uint16(v109)
						v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v125 = F_FunctionCall2Coll(m, l0+int32(160), v118, base.I64_extend_i32_u(v15+int32(48)), base.I64_extend_i32_u(v15+int32(32)))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v59
							if v125 == int64(0) {
								v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
								v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
								m.G0 = v15 + int32(96)
								return v262
							} else {
								v131 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
								v132 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
								v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+41)))
								v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+40)))
								v136 = v133
								v137 = v131
								v138 = v132
								v139 = v134
								v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
								if int32(0) < v140 {
									v143 = int32(_a_F_spgTestLeafTuple_4)
									v144 = *(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0]))
									v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
									*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v146
									v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
									if v5 == int32(0) {
										v155 = F_palloc(m, v140<<(uint(int32(3))%32)+int32(48))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v155)+42)) = uint8(v5)
											v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
											if v158 <= int32(0) {
											} else {
												v162 = v158 << (uint(int32(3)) % 32)
												if v162 == int32(0) {
												} else {
													base.MemoryCopy(m, v155+int32(48), v137, v162)
												}
											}
											*(*int32)(unsafe.Add(mBase, uint32(v155)+32)) = v148
											v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)))
											*(*uint16)(unsafe.Add(mBase, uint32(v155)+40)) = uint16(v170)
											v172 = *(*int32)(unsafe.Add(mBase, uint32(v23)+6))
											*(*int32)(unsafe.Add(mBase, uint32(v155)+36)) = v172
											v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
											if v174 == int32(0) {
												v192 = v155
												*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = int32(0)
												*(*int64)(unsafe.Add(mBase, uint32(v192)+16)) = int64(0)
												v231 = v192
												v235 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
												*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
												v240 = v136 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
												v243 = v139 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
												v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v245, v231)
												mBase = m.M
												v247 = m.ExcPending
												if v247 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
													v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
													v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
													m.G0 = v15 + int32(96)
													return v262
												}
											} else {
												v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
												v178 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
												v179 = F_datumCopy(m, v138, v177, v178)
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return int32(0)
												} else {
													v209 = v155
													v211 = v179
													*(*int64)(unsafe.Add(mBase, uint32(v209)+16)) = v211
													v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
													v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
													if int32(2) <= v214 {
														v217 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
														v220 = F_palloc(m, int32(base.Ui32(v217)>>(uint(int32(2))%32)))
														mBase = m.M
														v221 = m.ExcPending
														if v221 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = v220
															v223 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
															v225 = int32(base.Ui32(v223) >> (uint(int32(2)) % 32))
															if v225 == int32(0) {
																v231 = v209
															} else {
																base.MemoryCopy(m, v220, v23, v225)
																v231 = v209
															}
															v235 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
															*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
															v240 = v136 & v235
															*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
															v243 = v139 & v235
															*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
															v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
															F_pairingheap_add(m, v245, v231)
															mBase = m.M
															v247 = m.ExcPending
															if v247 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
																v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
																v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
																m.G0 = v15 + int32(96)
																return v262
															}
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = int32(0)
														v231 = v209
														v235 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
														*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
														v240 = v136 & v235
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
														v243 = v139 & v235
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
														v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
														F_pairingheap_add(m, v245, v231)
														mBase = m.M
														v247 = m.ExcPending
														if v247 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
															v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
															v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
															m.G0 = v15 + int32(96)
															return v262
														}
													}
												}
											}
										}
									} else {
										v182 = F_palloc(m, int32(48))
										mBase = m.M
										v183 = m.ExcPending
										if v183 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v182)+32)) = v148
											*(*uint8)(unsafe.Add(mBase, uint32(v182)+42)) = uint8(v5)
											v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)))
											*(*uint16)(unsafe.Add(mBase, uint32(v182)+40)) = uint16(v186)
											v188 = *(*int32)(unsafe.Add(mBase, uint32(v23)+6))
											*(*int32)(unsafe.Add(mBase, uint32(v182)+36)) = v188
											v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
											if v191 != 0 {
												v209 = v182
												v211 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v209)+16)) = v211
												v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
												if int32(2) <= v214 {
													v217 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
													v220 = F_palloc(m, int32(base.Ui32(v217)>>(uint(int32(2))%32)))
													mBase = m.M
													v221 = m.ExcPending
													if v221 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = v220
														v223 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
														v225 = int32(base.Ui32(v223) >> (uint(int32(2)) % 32))
														if v225 == int32(0) {
															v231 = v209
														} else {
															base.MemoryCopy(m, v220, v23, v225)
															v231 = v209
														}
														v235 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
														*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
														v240 = v136 & v235
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
														v243 = v139 & v235
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
														v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
														F_pairingheap_add(m, v245, v231)
														mBase = m.M
														v247 = m.ExcPending
														if v247 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
															v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
															v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
															m.G0 = v15 + int32(96)
															return v262
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = int32(0)
													v231 = v209
													v235 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
													*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
													v240 = v136 & v235
													*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
													v243 = v139 & v235
													*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
													v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v245, v231)
													mBase = m.M
													v247 = m.ExcPending
													if v247 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
														v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
														v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
														m.G0 = v15 + int32(96)
														return v262
													}
												}
											} else {
												v192 = v182
												*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = int32(0)
												*(*int64)(unsafe.Add(mBase, uint32(v192)+16)) = int64(0)
												v231 = v192
												v235 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
												*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
												v240 = v136 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
												v243 = v139 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
												v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v245, v231)
												mBase = m.M
												v247 = m.ExcPending
												if v247 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
													v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
													v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
													m.G0 = v15 + int32(96)
													return v262
												}
											}
										}
									}
								} else {
									v203 = int32(0)
									m.T0[l7].(func(*base.Module, int32, int32, int64, int32, int32, int32, int32, int32))(m, l0, v23+int32(6), v138, v5, v23, v139&int32(1), v203, v203)
									mBase = m.M
									v206 = m.ExcPending
									if v206 != 0 {
										return int32(0)
									} else {
										v207 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v207)
										v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
										v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
										m.G0 = v15 + int32(96)
										return v262
									}
								}
							}
						}
					case 1:
						v90 = int64(*(*int16)(unsafe.Add(mBase, uint32(v80))))
						v108 = v90
						v109 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v109
						*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = v108
						*(*uint16)(unsafe.Add(mBase, uint32(v15)+40)) = uint16(v109)
						v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v125 = F_FunctionCall2Coll(m, l0+int32(160), v118, base.I64_extend_i32_u(v15+int32(48)), base.I64_extend_i32_u(v15+int32(32)))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v59
							if v125 == int64(0) {
								v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
								v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
								m.G0 = v15 + int32(96)
								return v262
							} else {
								v131 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
								v132 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
								v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+41)))
								v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+40)))
								v136 = v133
								v137 = v131
								v138 = v132
								v139 = v134
								v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
								if int32(0) < v140 {
									v143 = int32(_a_F_spgTestLeafTuple_4)
									v144 = *(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0]))
									v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
									*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v146
									v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
									if v5 == int32(0) {
										v155 = F_palloc(m, v140<<(uint(int32(3))%32)+int32(48))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v155)+42)) = uint8(v5)
											v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
											if v158 <= int32(0) {
											} else {
												v162 = v158 << (uint(int32(3)) % 32)
												if v162 == int32(0) {
												} else {
													base.MemoryCopy(m, v155+int32(48), v137, v162)
												}
											}
											*(*int32)(unsafe.Add(mBase, uint32(v155)+32)) = v148
											v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)))
											*(*uint16)(unsafe.Add(mBase, uint32(v155)+40)) = uint16(v170)
											v172 = *(*int32)(unsafe.Add(mBase, uint32(v23)+6))
											*(*int32)(unsafe.Add(mBase, uint32(v155)+36)) = v172
											v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
											if v174 == int32(0) {
												v192 = v155
												*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = int32(0)
												*(*int64)(unsafe.Add(mBase, uint32(v192)+16)) = int64(0)
												v231 = v192
												v235 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
												*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
												v240 = v136 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
												v243 = v139 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
												v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v245, v231)
												mBase = m.M
												v247 = m.ExcPending
												if v247 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
													v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
													v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
													m.G0 = v15 + int32(96)
													return v262
												}
											} else {
												v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
												v178 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
												v179 = F_datumCopy(m, v138, v177, v178)
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return int32(0)
												} else {
													v209 = v155
													v211 = v179
													*(*int64)(unsafe.Add(mBase, uint32(v209)+16)) = v211
													v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
													v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
													if int32(2) <= v214 {
														v217 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
														v220 = F_palloc(m, int32(base.Ui32(v217)>>(uint(int32(2))%32)))
														mBase = m.M
														v221 = m.ExcPending
														if v221 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = v220
															v223 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
															v225 = int32(base.Ui32(v223) >> (uint(int32(2)) % 32))
															if v225 == int32(0) {
																v231 = v209
															} else {
																base.MemoryCopy(m, v220, v23, v225)
																v231 = v209
															}
															v235 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
															*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
															v240 = v136 & v235
															*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
															v243 = v139 & v235
															*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
															v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
															F_pairingheap_add(m, v245, v231)
															mBase = m.M
															v247 = m.ExcPending
															if v247 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
																v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
																v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
																m.G0 = v15 + int32(96)
																return v262
															}
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = int32(0)
														v231 = v209
														v235 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
														*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
														v240 = v136 & v235
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
														v243 = v139 & v235
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
														v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
														F_pairingheap_add(m, v245, v231)
														mBase = m.M
														v247 = m.ExcPending
														if v247 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
															v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
															v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
															m.G0 = v15 + int32(96)
															return v262
														}
													}
												}
											}
										}
									} else {
										v182 = F_palloc(m, int32(48))
										mBase = m.M
										v183 = m.ExcPending
										if v183 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v182)+32)) = v148
											*(*uint8)(unsafe.Add(mBase, uint32(v182)+42)) = uint8(v5)
											v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)))
											*(*uint16)(unsafe.Add(mBase, uint32(v182)+40)) = uint16(v186)
											v188 = *(*int32)(unsafe.Add(mBase, uint32(v23)+6))
											*(*int32)(unsafe.Add(mBase, uint32(v182)+36)) = v188
											v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
											if v191 != 0 {
												v209 = v182
												v211 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v209)+16)) = v211
												v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
												if int32(2) <= v214 {
													v217 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
													v220 = F_palloc(m, int32(base.Ui32(v217)>>(uint(int32(2))%32)))
													mBase = m.M
													v221 = m.ExcPending
													if v221 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = v220
														v223 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
														v225 = int32(base.Ui32(v223) >> (uint(int32(2)) % 32))
														if v225 == int32(0) {
															v231 = v209
														} else {
															base.MemoryCopy(m, v220, v23, v225)
															v231 = v209
														}
														v235 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
														*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
														v240 = v136 & v235
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
														v243 = v139 & v235
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
														v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
														F_pairingheap_add(m, v245, v231)
														mBase = m.M
														v247 = m.ExcPending
														if v247 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
															v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
															v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
															m.G0 = v15 + int32(96)
															return v262
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = int32(0)
													v231 = v209
													v235 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
													*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
													v240 = v136 & v235
													*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
													v243 = v139 & v235
													*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
													v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v245, v231)
													mBase = m.M
													v247 = m.ExcPending
													if v247 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
														v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
														v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
														m.G0 = v15 + int32(96)
														return v262
													}
												}
											} else {
												v192 = v182
												*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = int32(0)
												*(*int64)(unsafe.Add(mBase, uint32(v192)+16)) = int64(0)
												v231 = v192
												v235 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
												*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
												v240 = v136 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
												v243 = v139 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
												v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v245, v231)
												mBase = m.M
												v247 = m.ExcPending
												if v247 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
													v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
													v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
													m.G0 = v15 + int32(96)
													return v262
												}
											}
										}
									}
								} else {
									v203 = int32(0)
									m.T0[l7].(func(*base.Module, int32, int32, int64, int32, int32, int32, int32, int32))(m, l0, v23+int32(6), v138, v5, v23, v139&int32(1), v203, v203)
									mBase = m.M
									v206 = m.ExcPending
									if v206 != 0 {
										return int32(0)
									} else {
										v207 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v207)
										v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
										v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
										m.G0 = v15 + int32(96)
										return v262
									}
								}
							}
						}
					case 2:
						v91 = int64(*(*int32)(unsafe.Add(mBase, uint32(v80))))
						v108 = v91
						v109 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v109
						*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = v108
						*(*uint16)(unsafe.Add(mBase, uint32(v15)+40)) = uint16(v109)
						v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v125 = F_FunctionCall2Coll(m, l0+int32(160), v118, base.I64_extend_i32_u(v15+int32(48)), base.I64_extend_i32_u(v15+int32(32)))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v59
							if v125 == int64(0) {
								v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
								v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
								m.G0 = v15 + int32(96)
								return v262
							} else {
								v131 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
								v132 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
								v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+41)))
								v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+40)))
								v136 = v133
								v137 = v131
								v138 = v132
								v139 = v134
								v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
								if int32(0) < v140 {
									v143 = int32(_a_F_spgTestLeafTuple_4)
									v144 = *(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0]))
									v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
									*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v146
									v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
									if v5 == int32(0) {
										v155 = F_palloc(m, v140<<(uint(int32(3))%32)+int32(48))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v155)+42)) = uint8(v5)
											v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
											if v158 <= int32(0) {
											} else {
												v162 = v158 << (uint(int32(3)) % 32)
												if v162 == int32(0) {
												} else {
													base.MemoryCopy(m, v155+int32(48), v137, v162)
												}
											}
											*(*int32)(unsafe.Add(mBase, uint32(v155)+32)) = v148
											v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)))
											*(*uint16)(unsafe.Add(mBase, uint32(v155)+40)) = uint16(v170)
											v172 = *(*int32)(unsafe.Add(mBase, uint32(v23)+6))
											*(*int32)(unsafe.Add(mBase, uint32(v155)+36)) = v172
											v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
											if v174 == int32(0) {
												v192 = v155
												*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = int32(0)
												*(*int64)(unsafe.Add(mBase, uint32(v192)+16)) = int64(0)
												v231 = v192
												v235 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
												*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
												v240 = v136 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
												v243 = v139 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
												v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v245, v231)
												mBase = m.M
												v247 = m.ExcPending
												if v247 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
													v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
													v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
													m.G0 = v15 + int32(96)
													return v262
												}
											} else {
												v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
												v178 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
												v179 = F_datumCopy(m, v138, v177, v178)
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return int32(0)
												} else {
													v209 = v155
													v211 = v179
													*(*int64)(unsafe.Add(mBase, uint32(v209)+16)) = v211
													v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
													v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
													if int32(2) <= v214 {
														v217 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
														v220 = F_palloc(m, int32(base.Ui32(v217)>>(uint(int32(2))%32)))
														mBase = m.M
														v221 = m.ExcPending
														if v221 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = v220
															v223 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
															v225 = int32(base.Ui32(v223) >> (uint(int32(2)) % 32))
															if v225 == int32(0) {
																v231 = v209
															} else {
																base.MemoryCopy(m, v220, v23, v225)
																v231 = v209
															}
															v235 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
															*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
															v240 = v136 & v235
															*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
															v243 = v139 & v235
															*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
															v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
															F_pairingheap_add(m, v245, v231)
															mBase = m.M
															v247 = m.ExcPending
															if v247 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
																v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
																v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
																m.G0 = v15 + int32(96)
																return v262
															}
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = int32(0)
														v231 = v209
														v235 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
														*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
														v240 = v136 & v235
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
														v243 = v139 & v235
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
														v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
														F_pairingheap_add(m, v245, v231)
														mBase = m.M
														v247 = m.ExcPending
														if v247 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
															v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
															v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
															m.G0 = v15 + int32(96)
															return v262
														}
													}
												}
											}
										}
									} else {
										v182 = F_palloc(m, int32(48))
										mBase = m.M
										v183 = m.ExcPending
										if v183 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v182)+32)) = v148
											*(*uint8)(unsafe.Add(mBase, uint32(v182)+42)) = uint8(v5)
											v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)))
											*(*uint16)(unsafe.Add(mBase, uint32(v182)+40)) = uint16(v186)
											v188 = *(*int32)(unsafe.Add(mBase, uint32(v23)+6))
											*(*int32)(unsafe.Add(mBase, uint32(v182)+36)) = v188
											v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
											if v191 != 0 {
												v209 = v182
												v211 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v209)+16)) = v211
												v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
												if int32(2) <= v214 {
													v217 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
													v220 = F_palloc(m, int32(base.Ui32(v217)>>(uint(int32(2))%32)))
													mBase = m.M
													v221 = m.ExcPending
													if v221 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = v220
														v223 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
														v225 = int32(base.Ui32(v223) >> (uint(int32(2)) % 32))
														if v225 == int32(0) {
															v231 = v209
														} else {
															base.MemoryCopy(m, v220, v23, v225)
															v231 = v209
														}
														v235 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
														*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
														v240 = v136 & v235
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
														v243 = v139 & v235
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
														v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
														F_pairingheap_add(m, v245, v231)
														mBase = m.M
														v247 = m.ExcPending
														if v247 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
															v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
															v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
															m.G0 = v15 + int32(96)
															return v262
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = int32(0)
													v231 = v209
													v235 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
													*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
													v240 = v136 & v235
													*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
													v243 = v139 & v235
													*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
													v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v245, v231)
													mBase = m.M
													v247 = m.ExcPending
													if v247 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
														v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
														v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
														m.G0 = v15 + int32(96)
														return v262
													}
												}
											} else {
												v192 = v182
												*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = int32(0)
												*(*int64)(unsafe.Add(mBase, uint32(v192)+16)) = int64(0)
												v231 = v192
												v235 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
												*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
												v240 = v136 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
												v243 = v139 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
												v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v245, v231)
												mBase = m.M
												v247 = m.ExcPending
												if v247 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
													v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
													v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
													m.G0 = v15 + int32(96)
													return v262
												}
											}
										}
									}
								} else {
									v203 = int32(0)
									m.T0[l7].(func(*base.Module, int32, int32, int64, int32, int32, int32, int32, int32))(m, l0, v23+int32(6), v138, v5, v23, v139&int32(1), v203, v203)
									mBase = m.M
									v206 = m.ExcPending
									if v206 != 0 {
										return int32(0)
									} else {
										v207 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v207)
										v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
										v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
										m.G0 = v15 + int32(96)
										return v262
									}
								}
							}
						}
					case 3:
						v92 = *(*int64)(unsafe.Add(mBase, uint32(v80)))
						v108 = v92
						v109 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v109
						*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = v108
						*(*uint16)(unsafe.Add(mBase, uint32(v15)+40)) = uint16(v109)
						v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v125 = F_FunctionCall2Coll(m, l0+int32(160), v118, base.I64_extend_i32_u(v15+int32(48)), base.I64_extend_i32_u(v15+int32(32)))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v59
							if v125 == int64(0) {
								v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
								v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
								m.G0 = v15 + int32(96)
								return v262
							} else {
								v131 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
								v132 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
								v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+41)))
								v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+40)))
								v136 = v133
								v137 = v131
								v138 = v132
								v139 = v134
								v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
								if int32(0) < v140 {
									v143 = int32(_a_F_spgTestLeafTuple_4)
									v144 = *(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0]))
									v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
									*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v146
									v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
									if v5 == int32(0) {
										v155 = F_palloc(m, v140<<(uint(int32(3))%32)+int32(48))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v155)+42)) = uint8(v5)
											v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
											if v158 <= int32(0) {
											} else {
												v162 = v158 << (uint(int32(3)) % 32)
												if v162 == int32(0) {
												} else {
													base.MemoryCopy(m, v155+int32(48), v137, v162)
												}
											}
											*(*int32)(unsafe.Add(mBase, uint32(v155)+32)) = v148
											v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)))
											*(*uint16)(unsafe.Add(mBase, uint32(v155)+40)) = uint16(v170)
											v172 = *(*int32)(unsafe.Add(mBase, uint32(v23)+6))
											*(*int32)(unsafe.Add(mBase, uint32(v155)+36)) = v172
											v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
											if v174 == int32(0) {
												v192 = v155
												*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = int32(0)
												*(*int64)(unsafe.Add(mBase, uint32(v192)+16)) = int64(0)
												v231 = v192
												v235 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
												*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
												v240 = v136 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
												v243 = v139 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
												v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v245, v231)
												mBase = m.M
												v247 = m.ExcPending
												if v247 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
													v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
													v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
													m.G0 = v15 + int32(96)
													return v262
												}
											} else {
												v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
												v178 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
												v179 = F_datumCopy(m, v138, v177, v178)
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return int32(0)
												} else {
													v209 = v155
													v211 = v179
													*(*int64)(unsafe.Add(mBase, uint32(v209)+16)) = v211
													v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
													v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
													if int32(2) <= v214 {
														v217 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
														v220 = F_palloc(m, int32(base.Ui32(v217)>>(uint(int32(2))%32)))
														mBase = m.M
														v221 = m.ExcPending
														if v221 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = v220
															v223 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
															v225 = int32(base.Ui32(v223) >> (uint(int32(2)) % 32))
															if v225 == int32(0) {
																v231 = v209
															} else {
																base.MemoryCopy(m, v220, v23, v225)
																v231 = v209
															}
															v235 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
															*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
															v240 = v136 & v235
															*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
															v243 = v139 & v235
															*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
															v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
															F_pairingheap_add(m, v245, v231)
															mBase = m.M
															v247 = m.ExcPending
															if v247 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
																v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
																v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
																m.G0 = v15 + int32(96)
																return v262
															}
														}
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = int32(0)
														v231 = v209
														v235 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
														*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
														v240 = v136 & v235
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
														v243 = v139 & v235
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
														v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
														F_pairingheap_add(m, v245, v231)
														mBase = m.M
														v247 = m.ExcPending
														if v247 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
															v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
															v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
															m.G0 = v15 + int32(96)
															return v262
														}
													}
												}
											}
										}
									} else {
										v182 = F_palloc(m, int32(48))
										mBase = m.M
										v183 = m.ExcPending
										if v183 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v182)+32)) = v148
											*(*uint8)(unsafe.Add(mBase, uint32(v182)+42)) = uint8(v5)
											v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)))
											*(*uint16)(unsafe.Add(mBase, uint32(v182)+40)) = uint16(v186)
											v188 = *(*int32)(unsafe.Add(mBase, uint32(v23)+6))
											*(*int32)(unsafe.Add(mBase, uint32(v182)+36)) = v188
											v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
											if v191 != 0 {
												v209 = v182
												v211 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v209)+16)) = v211
												v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
												if int32(2) <= v214 {
													v217 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
													v220 = F_palloc(m, int32(base.Ui32(v217)>>(uint(int32(2))%32)))
													mBase = m.M
													v221 = m.ExcPending
													if v221 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = v220
														v223 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
														v225 = int32(base.Ui32(v223) >> (uint(int32(2)) % 32))
														if v225 == int32(0) {
															v231 = v209
														} else {
															base.MemoryCopy(m, v220, v23, v225)
															v231 = v209
														}
														v235 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
														*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
														v240 = v136 & v235
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
														v243 = v139 & v235
														*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
														v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
														F_pairingheap_add(m, v245, v231)
														mBase = m.M
														v247 = m.ExcPending
														if v247 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
															v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
															v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
															m.G0 = v15 + int32(96)
															return v262
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = int32(0)
													v231 = v209
													v235 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
													*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
													v240 = v136 & v235
													*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
													v243 = v139 & v235
													*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
													v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v245, v231)
													mBase = m.M
													v247 = m.ExcPending
													if v247 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
														v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
														v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
														m.G0 = v15 + int32(96)
														return v262
													}
												}
											} else {
												v192 = v182
												*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = int32(0)
												*(*int64)(unsafe.Add(mBase, uint32(v192)+16)) = int64(0)
												v231 = v192
												v235 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
												*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
												v240 = v136 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
												v243 = v139 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
												v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v245, v231)
												mBase = m.M
												v247 = m.ExcPending
												if v247 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
													v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
													v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
													m.G0 = v15 + int32(96)
													return v262
												}
											}
										}
									}
								} else {
									v203 = int32(0)
									m.T0[l7].(func(*base.Module, int32, int32, int64, int32, int32, int32, int32, int32))(m, l0, v23+int32(6), v138, v5, v23, v139&int32(1), v203, v203)
									mBase = m.M
									v206 = m.ExcPending
									if v206 != 0 {
										return int32(0)
									} else {
										v207 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v207)
										v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
										v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
										m.G0 = v15 + int32(96)
										return v262
									}
								}
							}
						}
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v15))) = v84
							F_errmsg_internal(m, int32(_a_F_spgTestLeafTuple_6), v15)
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_spgTestLeafTuple_7), int32(123), int32(_a_F_spgTestLeafTuple_8))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
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
				v108 = base.I64_extend_i32_u(v80)
				v109 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v109
				*(*int64)(unsafe.Add(mBase, uint32(v15)+32)) = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v15)+88)) = v108
				*(*uint16)(unsafe.Add(mBase, uint32(v15)+40)) = uint16(v109)
				v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v125 = F_FunctionCall2Coll(m, l0+int32(160), v118, base.I64_extend_i32_u(v15+int32(48)), base.I64_extend_i32_u(v15+int32(32)))
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v59
					if v125 == int64(0) {
						v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
						v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
						m.G0 = v15 + int32(96)
						return v262
					} else {
						v131 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
						v132 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
						v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+41)))
						v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+40)))
						v136 = v133
						v137 = v131
						v138 = v132
						v139 = v134
						v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
						if int32(0) < v140 {
							v143 = int32(_a_F_spgTestLeafTuple_4)
							v144 = *(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0]))
							v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v146
							v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
							if v5 == int32(0) {
								v155 = F_palloc(m, v140<<(uint(int32(3))%32)+int32(48))
								mBase = m.M
								v156 = m.ExcPending
								if v156 != 0 {
									return int32(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v155)+42)) = uint8(v5)
									v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
									if v158 <= int32(0) {
									} else {
										v162 = v158 << (uint(int32(3)) % 32)
										if v162 == int32(0) {
										} else {
											base.MemoryCopy(m, v155+int32(48), v137, v162)
										}
									}
									*(*int32)(unsafe.Add(mBase, uint32(v155)+32)) = v148
									v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)))
									*(*uint16)(unsafe.Add(mBase, uint32(v155)+40)) = uint16(v170)
									v172 = *(*int32)(unsafe.Add(mBase, uint32(v23)+6))
									*(*int32)(unsafe.Add(mBase, uint32(v155)+36)) = v172
									v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
									if v174 == int32(0) {
										v192 = v155
										*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = int32(0)
										*(*int64)(unsafe.Add(mBase, uint32(v192)+16)) = int64(0)
										v231 = v192
										v235 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
										*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
										v240 = v136 & v235
										*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
										v243 = v139 & v235
										*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
										v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
										F_pairingheap_add(m, v245, v231)
										mBase = m.M
										v247 = m.ExcPending
										if v247 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
											v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
											v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
											m.G0 = v15 + int32(96)
											return v262
										}
									} else {
										v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
										v178 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
										v179 = F_datumCopy(m, v138, v177, v178)
										mBase = m.M
										v180 = m.ExcPending
										if v180 != 0 {
											return int32(0)
										} else {
											v209 = v155
											v211 = v179
											*(*int64)(unsafe.Add(mBase, uint32(v209)+16)) = v211
											v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
											if int32(2) <= v214 {
												v217 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
												v220 = F_palloc(m, int32(base.Ui32(v217)>>(uint(int32(2))%32)))
												mBase = m.M
												v221 = m.ExcPending
												if v221 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = v220
													v223 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
													v225 = int32(base.Ui32(v223) >> (uint(int32(2)) % 32))
													if v225 == int32(0) {
														v231 = v209
													} else {
														base.MemoryCopy(m, v220, v23, v225)
														v231 = v209
													}
													v235 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
													*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
													v240 = v136 & v235
													*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
													v243 = v139 & v235
													*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
													v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v245, v231)
													mBase = m.M
													v247 = m.ExcPending
													if v247 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
														v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
														v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
														m.G0 = v15 + int32(96)
														return v262
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = int32(0)
												v231 = v209
												v235 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
												*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
												v240 = v136 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
												v243 = v139 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
												v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v245, v231)
												mBase = m.M
												v247 = m.ExcPending
												if v247 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
													v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
													v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
													m.G0 = v15 + int32(96)
													return v262
												}
											}
										}
									}
								}
							} else {
								v182 = F_palloc(m, int32(48))
								mBase = m.M
								v183 = m.ExcPending
								if v183 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v182)+32)) = v148
									*(*uint8)(unsafe.Add(mBase, uint32(v182)+42)) = uint8(v5)
									v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)))
									*(*uint16)(unsafe.Add(mBase, uint32(v182)+40)) = uint16(v186)
									v188 = *(*int32)(unsafe.Add(mBase, uint32(v23)+6))
									*(*int32)(unsafe.Add(mBase, uint32(v182)+36)) = v188
									v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
									if v191 != 0 {
										v209 = v182
										v211 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v209)+16)) = v211
										v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
										v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
										if int32(2) <= v214 {
											v217 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
											v220 = F_palloc(m, int32(base.Ui32(v217)>>(uint(int32(2))%32)))
											mBase = m.M
											v221 = m.ExcPending
											if v221 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = v220
												v223 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
												v225 = int32(base.Ui32(v223) >> (uint(int32(2)) % 32))
												if v225 == int32(0) {
													v231 = v209
												} else {
													base.MemoryCopy(m, v220, v23, v225)
													v231 = v209
												}
												v235 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
												*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
												v240 = v136 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
												v243 = v139 & v235
												*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
												v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v245, v231)
												mBase = m.M
												v247 = m.ExcPending
												if v247 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
													v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
													v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
													m.G0 = v15 + int32(96)
													return v262
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v209)+24)) = int32(0)
											v231 = v209
											v235 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
											*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
											v240 = v136 & v235
											*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
											v243 = v139 & v235
											*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
											v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
											F_pairingheap_add(m, v245, v231)
											mBase = m.M
											v247 = m.ExcPending
											if v247 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
												v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
												v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
												m.G0 = v15 + int32(96)
												return v262
											}
										}
									} else {
										v192 = v182
										*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = int32(0)
										*(*int64)(unsafe.Add(mBase, uint32(v192)+16)) = int64(0)
										v231 = v192
										v235 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v231)+43)) = uint8(v235)
										*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = int32(0)
										v240 = v136 & v235
										*(*uint8)(unsafe.Add(mBase, uint32(v231)+45)) = uint8(v240)
										v243 = v139 & v235
										*(*uint8)(unsafe.Add(mBase, uint32(v231)+44)) = uint8(v243)
										v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
										F_pairingheap_add(m, v245, v231)
										mBase = m.M
										v247 = m.ExcPending
										if v247 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v144
											v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
											v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
											m.G0 = v15 + int32(96)
											return v262
										}
									}
								}
							}
						} else {
							v203 = int32(0)
							m.T0[l7].(func(*base.Module, int32, int32, int64, int32, int32, int32, int32, int32))(m, l0, v23+int32(6), v138, v5, v23, v139&int32(1), v203, v203)
							mBase = m.M
							v206 = m.ExcPending
							if v206 != 0 {
								return int32(0)
							} else {
								v207 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v207)
								v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+4)))
								v262 = v259 & int32(_a_F_spgTestLeafTuple_5)
								m.G0 = v15 + int32(96)
								return v262
							}
						}
					}
				}
			}
		}
	}
}
func F_spg_kd_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
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
	var v62 int32
	_ = v62
	var v70 int64
	_ = v70
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
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = F_palloc_mul(m, int32(8), v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		if int32(0) < v16 {
			v20 = int32(0)
			for {
				v28 = v20 << (uint(int32(3)) % 32)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				v31 = *(*int64)(unsafe.Add(mBase, uint32(v28+v29)))
				v32 = v28 + v12
				*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v20
				*(*uint32)(unsafe.Add(mBase, uint32(v32))) = uint32(v31)
				v36 = v20 + int32(1)
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				if v36 < v37 {
					v20 = v36
					continue
				} else {
					break
				}
				break
			}
			v41 = v37
		} else {
			v41 = v16
		}
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		if v49&int32(1) != 0 {
			v52 = int32(254)
		} else {
			v52 = int32(255)
		}
		F_pg_qsort(m, v12, v41, int32(8), v52)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int64(0)
		} else {
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v56 = int32(1)
			v57 = v55 >> (uint(v56) % 32)
			v58 = int32(3)
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v12+v57<<(uint(v58)%32))))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			v70 = *(*int64)(unsafe.Add(mBase, uint32(v61+(v62^int32(-1))<<(uint(v58)%32)&int32(8))))
			*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = int64(2)
			*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v70
			*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v56)
			v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v78 = F_palloc_mul(m, int32(4), v77)
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v78
				v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v83 = F_palloc_mul(m, int32(8), v82)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v83
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					if int32(0) < v86 {
						v90 = int32(0)
						for {
							v97 = int32(3)
							v99 = v12 + v90<<(uint(v97)%32)
							v100 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v99))))
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v101+v102<<(uint(int32(2))%32)))) = base.B2i32(v57 <= v90)
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v108+v102<<(uint(v97)%32)))) = v100
							v114 = v90 + int32(1)
							v115 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							if v114 < v115 {
								v90 = v114
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					return int64(0)
				}
			}
		}
	}
}
func F_spg_quad_choose(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v5))))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+20)))
	if v7 == int32(1) {
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(1)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+16)) = v6
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(0)
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(1)
		v16 = F_getQuadrant(m, v12, base.I32_wrap_i64(v6))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v16 - int32(1)
			*(*int64)(unsafe.Add(mBase, uint32(v4)+16)) = v6
			*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(0)
			return int64(0)
		}
	}
}
func F_spg_quad_leaf_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v115 int64
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int64
	_ = v122
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+40)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v4)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v4 < v19 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v122
L2:
	;
	v23 = int32(0)
	goto L5
L3:
	;
	goto L4
L4:
	;
	v111 = int64(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v112 <= int32(0) {
		v122 = v111
		goto L1
	} else {
		goto L32
	}
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v33 = v30 + v23*int32(56)
	v34 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v33)+48)))
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+6)))
	switch v35 - int32(1) {
	case 0:
		goto L14
	default:
		goto L9
	case 4:
		goto L8
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
	v101 = v23 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v101 < v102 {
		v23 = v101
		goto L5
	} else {
		goto L31
	}
L8:
	;
	v95 = F_DirectFunctionCall2Coll(m, int32(258), int32(0), v13, v34)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L15
	} else {
		goto L29
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L15
	} else {
		goto L26
	}
L10:
	;
	v70 = F_DirectFunctionCall2Coll(m, int32(262), int32(0), v34, v13)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L15
	} else {
		goto L24
	}
L11:
	;
	v63 = F_DirectFunctionCall2Coll(m, int32(256), int32(0), v13, v34)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L15
	} else {
		goto L22
	}
L12:
	;
	v56 = F_DirectFunctionCall2Coll(m, int32(260), int32(0), v13, v34)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L15
	} else {
		goto L20
	}
L13:
	;
	v49 = F_DirectFunctionCall2Coll(m, int32(263), int32(0), v13, v34)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L15
	} else {
		goto L18
	}
L14:
	;
	v40 = F_DirectFunctionCall2Coll(m, int32(261), int32(0), v13, v34)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int64(0)
L16:
	;
	if v40 != int64(0) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v122 = int64(0)
	goto L1
L18:
	;
	if v49 != int64(0) {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v122 = int64(0)
	goto L1
L20:
	;
	if v56 != int64(0) {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v122 = int64(0)
	goto L1
L22:
	;
	if v63 != int64(0) {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	v122 = int64(0)
	goto L1
L24:
	;
	if v70 != int64(0) {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v122 = int64(0)
	goto L1
L26:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79+v23*int32(56))+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v83
	F_errmsg_internal(m, int32(_a_F_spg_quad_leaf_consistent_0), v10)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_spg_quad_leaf_consistent_1), int32(459), int32(_a_F_spg_quad_leaf_consistent_2))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
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
	if v95 != int64(0) {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	v122 = int64(0)
	goto L1
L31:
	;
	goto L6
L32:
	;
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v12)+40))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v118 = F_spg_key_orderbys_distances(m, v115, int32(1), v117, v112)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v118
	v122 = v111
	goto L1
}
func F_spg_range_quad_choose(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+20)))
		if v20 == int32(1) {
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(1)
			v93 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v93
			*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = base.I64_extend_i32_u(v16)
			m.G0 = v11 + int32(80)
			return int64(0)
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
			v27 = F_range_get_typcache(m, l0, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int64(0)
			} else {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+21)))
				if v29 == int32(0) {
					v32 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v32
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v16+int32(base.Ui32(v35)>>(uint(int32(2))%32))-v32))))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = (v41 ^ int32(-1)) & int32(1)
					v93 = v32
					*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v93
					*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = base.I64_extend_i32_u(v16)
					m.G0 = v11 + int32(80)
					return int64(0)
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
					v48 = F_pg_detoast_datum(m, v47)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int64(0)
					} else {
						v51 = v11 - int32(-64)
						v53 = v11 + int32(48)
						F_range_deserialize(m, v27, v48, v51, v53, v11+int32(47))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int64(0)
						} else {
							v59 = v11 + int32(24)
							v61 = v11 + int32(8)
							F_range_deserialize(m, v27, v16, v59, v61, v11+int32(7))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int64(0)
							} else {
								v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+7)))
								if v67 != 0 {
									v86 = int32(5)
									v87 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v86 - v87
									*(*int32)(unsafe.Add(mBase, uint32(v13))) = v87
									v93 = v87
									*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v93
									*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = base.I64_extend_i32_u(v16)
									m.G0 = v11 + int32(80)
									return int64(0)
								} else {
									v68 = F_range_cmp_bounds(m, v27, v59, v51)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int64(0)
									} else {
										v72 = F_range_cmp_bounds(m, v27, v61, v53)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int64(0)
										} else {
											if int32(0) <= v72 {
												v76 = int32(1)
											} else {
												v76 = int32(2)
											}
											if int32(0) <= v68 {
												v86 = v76
											} else {
												if int32(0) <= v72 {
													v83 = int32(4)
												} else {
													v83 = int32(3)
												}
												v86 = v83
											}
											v87 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v86 - v87
											*(*int32)(unsafe.Add(mBase, uint32(v13))) = v87
											v93 = v87
											*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v93
											*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = base.I64_extend_i32_u(v16)
											m.G0 = v11 + int32(80)
											return int64(0)
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
func F_spg_text_choose(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
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
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int64
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
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
	var v228 int32
	_ = v228
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v328 int64
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	v2 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v25 = v23 & int32(1)
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = int32(1)
	goto L5
L4:
	;
	v26 = int32(4)
	goto L5
L5:
	;
	if v23 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v54 = v19 + v26
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+21)))
	if v55 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v32 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v43 = int32(1)
	if v25 != 0 {
		v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v35 = int32(16)
	goto L12
L11:
	;
	v35 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v42 = int32(4)
	goto L15
L14:
	;
	v42 = v35
	goto L15
L15:
	;
	v53 = v42
	goto L6
L16:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if int32(0) < v275 {
		goto L79
	} else {
		goto L80
	}
L18:
	;
	if v70 <= v99 {
		v268 = v99
		v274 = int32(_a_F_spg_text_choose_0)
		goto L17
	} else {
		goto L76
	}
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v61 = F_pg_detoast_datum_packed(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v53 <= v240 {
		v268 = v2
		v274 = int32(_a_F_spg_text_choose_0)
		goto L17
	} else {
		goto L75
	}
L22:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v63&int32(1) != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v66 = int32(1)
	goto L25
L24:
	;
	v66 = int32(4)
	goto L25
L25:
	;
	v67 = v61 + v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v69 = v54 + v68
	v70 = v53 - v68
	if v63 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = int32(1)
	v193 = F_palloc(m, int32(8))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L63
	}
L27:
	;
	v154 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v154)
	v158 = v129 + int32(4)
	v159 = F_palloc(m, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L55
	}
L28:
	;
	v151 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v151)
	v178 = v151
	goto L26
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(3)
	if v129 != 0 {
		goto L27
	} else {
		goto L54
	}
L30:
	;
	if v70 < v99 {
		goto L41
	} else {
		goto L42
	}
L31:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v76 == int32(18) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v87 = int32(1)
	if v63&v87 != 0 {
		v99 = int32(base.Ui32(v63)>>(uint(v87)%32)) - v87
		goto L30
	} else {
		goto L40
	}
L34:
	;
	v79 = int32(16)
	goto L36
L35:
	;
	v79 = int32(0)
	goto L36
L36:
	;
	if base.Ui32((v76-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v86 = int32(4)
	goto L39
L38:
	;
	v86 = v79
	goto L39
L39:
	;
	v99 = v86
	goto L30
L40:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v99 = int32(base.Ui32(v93)>>(uint(int32(2))%32)) - int32(4)
	goto L30
L41:
	;
	v101 = v70
	goto L43
L42:
	;
	v101 = v99
	goto L43
L43:
	;
	if int32(0) < v101 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v105 = v69
	v106 = int32(0)
	v108 = v67
	goto L48
L45:
	;
	goto L46
L46:
	;
	if v99 == int32(0) {
		goto L18
	} else {
		goto L53
	}
L47:
	;
	if v129 != v99 {
		goto L29
	} else {
		goto L52
	}
L48:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v118 != v119 {
		v129 = v106
		goto L47
	} else {
		goto L50
	}
L49:
	;
	v129 = v101
	goto L47
L50:
	;
	v121 = int32(1)
	v126 = v106 + v121
	if v126 != v101 {
		v105 = v105 + v121
		v106 = v126
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
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(3)
	goto L28
L54:
	;
	goto L28
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
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v166 = v129<<(uint(int32(1))%32) + int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v159))) = uint8(v166)
	v172 = v154
	goto L56
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159))) = v158 << (uint(int32(2)) % 32)
	v172 = int32(4)
	goto L56
L60:
	;
	base.MemoryCopy(m, v159+v172, v67, v129)
	goto L62
L61:
	;
	goto L62
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = base.I64_extend_i32_u(v159)
	v178 = v129
	goto L26
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v193
	v196 = v178 + v67
	v197 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	*(*int64)(unsafe.Add(mBase, uint32(v193))) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(0)
	v201 = v99 - v178
	if v201 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v204 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+36)) = uint8(v204)
	return int64(0)
L65:
	;
	goto L66
L66:
	;
	v208 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+36)) = uint8(v208)
	v211 = v201 - v208
	v213 = v201 + int32(3)
	v214 = F_palloc(m, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if base.Ui32(v201) <= base.Ui32(int32(127)) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = base.I64_extend_i32_u(v214)
	return int64(0)
L69:
	;
	if v211 == int32(0) {
		goto L68
	} else {
		goto L74
	}
L70:
	;
	v218 = int32(1)
	v221 = v201<<(uint(v218)%32) | v218
	*(*uint8)(unsafe.Add(mBase, uint32(v214))) = uint8(v221)
	if v211 != 0 {
		v228 = v218
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = v213 << (uint(int32(2)) % 32)
	v228 = int32(4)
	goto L69
L73:
	;
	goto L68
L74:
	;
	base.MemoryCopy(m, v228+v214, v196+int32(1), v211)
	goto L68
L75:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240+v54))))
	v268 = v2
	v274 = v243
	goto L17
L76:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+v99))))
	v268 = v99
	v274 = v260
	goto L17
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = base.I64_extend16_s(base.I64_extend_i32_u(v274))
	return int64(0)
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v296
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(1)
	v347 = int32(0)
	v349 = v268 + base.B2i32(v347 <= v293)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v349
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v353 = v53 - (v349 + v351)
	if v347 < v353 {
		goto L92
	} else {
		goto L93
	}
L79:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	v280 = v275
	v283 = int32(0)
	goto L82
L80:
	;
	v308 = v275
	goto L81
L81:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+20)))
	if v321 != int32(1) {
		goto L77
	} else {
		goto L90
	}
L82:
	;
	v293 = base.I32_extend16_s(v274)
	v296 = (v280 + v283) >> (uint(int32(1)) % 32)
	v300 = int32(*(*int16)(unsafe.Add(mBase, uint32(v278+v296<<(uint(int32(3))%32)))))
	if v293 < v300 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v308 = v305
	goto L81
L84:
	;
	if v306 < v305 {
		v280 = v305
		v283 = v306
		goto L82
	} else {
		goto L89
	}
L85:
	;
	v305 = v296
	v306 = v283
	goto L84
L86:
	;
	goto L87
L87:
	;
	if v293 <= v300 {
		goto L78
	} else {
		goto L88
	}
L88:
	;
	v305 = v280
	v306 = v296 + int32(1)
	goto L84
L89:
	;
	goto L83
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(3)
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+21)))
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v326)
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v328
	v333 = F_palloc(m, int32(8))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v333
	*(*int64)(unsafe.Add(mBase, uint32(v333))) = int64(-2)
	v338 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+36)) = uint8(v338)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v338
	return int64(0)
L92:
	;
	v356 = int32(4)
	v358 = v353 + v356
	v359 = F_palloc(m, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v382 = F_palloc(m, int32(4))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L103
	}
L95:
	;
	if base.Ui32(v353) <= base.Ui32(int32(126)) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v353 != 0 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v363 = int32(1)
	v367 = v353<<(uint(v363)%32) + int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v359))) = uint8(v367)
	v372 = v363
	goto L96
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359))) = v358 << (uint(int32(2)) % 32)
	v372 = v356
	goto L96
L100:
	;
	base.MemoryCopy(m, v359+v372, v351+v54+v349, v353)
	goto L102
L101:
	;
	goto L102
L102:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = base.I64_extend_i32_u(v359)
	return int64(0)
L103:
	;
	v384 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v382))) = uint8(v384)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = base.I64_extend_i32_u(v382)
	return int64(0)
}
func F_spg_text_leaf_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int64
	_ = v155
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int64
	_ = v219
	var v221 int64
	_ = v221
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
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
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v371 int64
	_ = v371
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)) = uint8(v2)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v25 = F_pg_detoast_datum_packed(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	v60 = v58 + v20
	v61 = int32(0)
	if v58|base.B2i32(v20 <= v61) == v61 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	return int64(0)
L3:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v29 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v32 = int32(1)
	if v29&v32 != 0 {
		v58 = int32(base.Ui32(v29)>>(uint(v32)%32)) - v32
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v46 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v58 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L8:
	;
	v49 = int32(16)
	goto L10
L9:
	;
	v49 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v46-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v56 = int32(4)
	goto L13
L12:
	;
	v56 = v49
	goto L13
L13:
	;
	v58 = v56
	goto L1
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = base.I64_extend_i32_u(v149)
	v155 = int64(1)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v156 <= int32(0) {
		v371 = v155
		goto L45
	} else {
		goto L46
	}
L15:
	;
	v149 = v59
	v151 = v59 + int32(4)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v69 = v60 + int32(4)
	v70 = F_palloc(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v69 << (uint(int32(2)) % 32)
	v76 = v70 + int32(4)
	if v20 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	base.MemoryCopy(m, v76, v59+int32(4), v20)
	goto L21
L20:
	;
	goto L21
L21:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v80 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v141 == int32(0) {
		v149 = v70
		v151 = v76
		goto L14
	} else {
		goto L44
	}
L23:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if base.B2i32(base.Ui32(int32(18)) < base.Ui32(v83))|base.B2i32(int32(1)<<(uint(v83)%32)&int32(_a_F_spg_text_leaf_consistent_0) == int32(0)) != 0 {
		v149 = v70
		v151 = v76
		goto L14
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v109 = v80 & int32(1)
	if v109 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	if v83 == int32(18) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v98 = int32(16)
	goto L29
L28:
	;
	v98 = int32(0)
	goto L29
L29:
	;
	if base.Ui32((v83-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v105 = int32(4)
	goto L32
L31:
	;
	v105 = v98
	goto L32
L32:
	;
	v141 = v105
	v143 = v25 + int32(1)
	goto L22
L33:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v112&int32(-4) == int32(16) {
		v149 = v70
		v151 = v76
		goto L14
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v80&int32(254) == int32(2) {
		v149 = v70
		v151 = v76
		goto L14
	} else {
		goto L40
	}
L36:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v120 = int32(4)
	if v109 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v124 = int32(1)
	goto L39
L38:
	;
	v124 = v120
	goto L39
L39:
	;
	v141 = int32(base.Ui32(v117)>>(uint(int32(2))%32)) - v120
	v143 = v25 + v124
	goto L22
L40:
	;
	v130 = int32(1)
	if v80&v130 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v138 = v130
	goto L43
L42:
	;
	v138 = int32(4)
	goto L43
L43:
	;
	v141 = int32(base.Ui32(v80)>>(uint(v130)%32)) - v130
	v143 = v25 + v138
	goto L22
L44:
	;
	base.MemoryCopy(m, v76+v20, v143, v141)
	v149 = v70
	v151 = v76
	goto L14
L45:
	;
	m.G0 = v17 + int32(16)
	return v371
L46:
	;
	v164 = int32(0)
	goto L47
L47:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v177 = v174 + v164*int32(56)
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v177)+6)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v177)+48))
	v180 = F_pg_detoast_datum_packed(m, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L50
	}
L48:
	;
	v371 = v155
	goto L45
L49:
	;
	v213 = v178 & int32(_a_F_spg_text_leaf_consistent_1)
	if v213 == int32(28) {
		goto L62
	} else {
		goto L63
	}
L50:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v182 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+1)))
	if v188 == int32(18) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v199 = int32(1)
	if v182&v199 != 0 {
		v211 = int32(base.Ui32(v182)>>(uint(v199)%32)) - v199
		goto L49
	} else {
		goto L60
	}
L54:
	;
	v191 = int32(16)
	goto L56
L55:
	;
	v191 = int32(0)
	goto L56
L56:
	;
	if base.Ui32((v188-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v198 = int32(4)
	goto L59
L58:
	;
	v198 = v191
	goto L59
L59:
	;
	v211 = v198
	goto L49
L60:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	v211 = int32(base.Ui32(v205)>>(uint(int32(2))%32)) - int32(4)
	goto L49
L61:
	;
	v355 = v164 + int32(1)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v355 < v356 {
		v164 = v355
		goto L47
	} else {
		goto L115
	}
L62:
	;
	if v211 <= v20 {
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if base.Ui32(int32(11)) <= base.Ui32(v213) {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v221 = F_DirectFunctionCall2Coll(m, int32(268), v218, v219, base.I64_extend_i32_u(v180))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	if v221 != int64(0) {
		goto L61
	} else {
		goto L67
	}
L67:
	;
	v371 = int64(0)
	goto L45
L68:
	;
	switch v312&int32(_a_F_spg_text_leaf_consistent_1) - int32(1) {
	case 0:
		goto L106
	case 1:
		goto L101
	case 2:
		goto L105
	case 3:
		goto L104
	case 4:
		goto L103
	default:
		goto L102
	}
L69:
	;
	v230 = int32(1)
	if v182&v230 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v239 = int32(1)
	if v182&v239 != 0 {
		goto L76
	} else {
		goto L77
	}
L72:
	;
	v234 = v230
	goto L74
L73:
	;
	v234 = int32(4)
	goto L74
L74:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v237 = F_varstr_cmp(m, v151, v60, v180+v234, v211, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	v311 = v237
	v312 = v178 - int32(10)
	goto L68
L76:
	;
	v243 = v239
	goto L78
L77:
	;
	v243 = int32(4)
	goto L78
L78:
	;
	v244 = v180 + v243
	v245 = base.B2i32(v211 < v60)
	if v211 < v60 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v246 = v211
	goto L81
L80:
	;
	v246 = v60
	goto L81
L81:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v246) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	if v308 != 0 {
		v311 = v308
		v312 = v178
		goto L68
	} else {
		goto L100
	}
L83:
	;
	v308 = int32(0)
	goto L82
L84:
	;
	v282 = v277
	v283 = v278
	v284 = v279
	goto L94
L85:
	;
	if (v151|v244)&int32(3) != 0 {
		v277 = v151
		v278 = v244
		v279 = v246
		goto L84
	} else {
		goto L88
	}
L86:
	;
	v270 = v151
	v271 = v244
	v272 = v246
	goto L87
L87:
	;
	if v272 == int32(0) {
		goto L83
	} else {
		goto L93
	}
L88:
	;
	v254 = v151
	v255 = v244
	v256 = v246
	goto L89
L89:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	if v259 != v260 {
		v277 = v254
		v278 = v255
		v279 = v256
		goto L84
	} else {
		goto L91
	}
L90:
	;
	v270 = v265
	v271 = v263
	v272 = v267
	goto L87
L91:
	;
	v262 = int32(4)
	v263 = v255 + v262
	v265 = v254 + v262
	v267 = v256 - v262
	if base.Ui32(int32(3)) < base.Ui32(v267) {
		v254 = v265
		v255 = v263
		v256 = v267
		goto L89
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v277 = v270
	v278 = v271
	v279 = v272
	goto L84
L94:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283))))
	if v287 == v288 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v308 = v287 - v288
	goto L82
L96:
	;
	v290 = int32(1)
	v295 = v284 - v290
	if v295 != 0 {
		v282 = v282 + v290
		v283 = v283 + v290
		v284 = v295
		goto L94
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	goto L95
L99:
	;
	goto L83
L100:
	;
	v311 = v245 - base.B2i32(v60 < v211)
	v312 = v178
	goto L68
L101:
	;
	if v311 <= int32(0) {
		goto L61
	} else {
		goto L114
	}
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L2
	} else {
		goto L111
	}
L103:
	;
	if int32(0) < v311 {
		goto L61
	} else {
		goto L110
	}
L104:
	;
	if int32(0) <= v311 {
		goto L61
	} else {
		goto L109
	}
L105:
	;
	if v311 == int32(0) {
		goto L61
	} else {
		goto L108
	}
L106:
	;
	if v311 < int32(0) {
		goto L61
	} else {
		goto L107
	}
L107:
	;
	v371 = int64(0)
	goto L45
L108:
	;
	v371 = int64(0)
	goto L45
L109:
	;
	v371 = int64(0)
	goto L45
L110:
	;
	v371 = int64(0)
	goto L45
L111:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v338 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v334+v164*int32(56))+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v338
	F_errmsg_internal(m, int32(_a_F_spg_text_leaf_consistent_2), v17)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_spg_text_leaf_consistent_3), int32(692), int32(_a_F_spg_text_leaf_consistent_4))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	v371 = int64(0)
	goto L45
L115:
	;
	goto L48
}
func F_spg_xlog_cleanup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_spg_xlog_cleanup[0]))
	F_MemoryContextDelete(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_spg_xlog_cleanup[0])) = int32(0)
		return
	}
}
