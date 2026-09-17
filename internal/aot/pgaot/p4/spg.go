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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
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
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	if l1 != 0 {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
		if v20 != 0 {
			v48 = int32(4)
		} else {
			v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+48)))
			if int32(0) < v21 {
				v48 = v21
			} else {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
				if v24 == int32(1) {
					v28 = int32(18)
					v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
					if v30 == v28 {
						v33 = v28
					} else {
						v33 = int32(2)
					}
					if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v40 = int32(6)
					} else {
						v40 = v33
					}
					v48 = v40
				} else {
					if v24&int32(1) != 0 {
						v48 = int32(base.Ui32(v24) >> (uint(int32(1)) % 32))
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v48 = int32(base.Ui32(v45) >> (uint(int32(2)) % 32))
					}
				}
			}
		}
		v54 = (v48 + int32(7)) & int32(-8)
	} else {
		v54 = v6
	}
	v56 = v54 + int32(8)
	if l3 <= int32(0) {
		v158 = v56
	} else {
		v60 = l3 & int32(3)
		if base.Ui32(l3) < base.Ui32(int32(4)) {
			v116 = int32(0)
			v117 = v56
			v130 = v116
			v131 = v117
			v135 = v6
			for {
				v142 = *(*int32)(unsafe.Add(mBase, uint32(l4+v130<<(uint(int32(2))%32))))
				v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+6)))
				v146 = v131 + v143&int32(_a_F_spgFormInnerTuple_0)
				v147 = int32(1)
				v150 = v135 + v147
				if v150 != v60 {
					v130 = v130 + v147
					v131 = v146
					v135 = v150
					continue
				} else {
					break
				}
				break
			}
			v158 = v146
		} else {
			v72 = int32(0)
			v73 = v56
			v80 = v6
			for {
				v83 = l4 + v72<<(uint(int32(2))%32)
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
				v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+6)))
				v86 = int32(_a_F_spgFormInnerTuple_0)
				v89 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
				v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+6)))
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
				v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v94)+6)))
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
				v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+6)))
				v103 = v73 + v85&v86 + v90&v86 + v95&v86 + v100&v86
				v104 = int32(4)
				v105 = v72 + v104
				v107 = v80 + v104
				if v107 != l3&int32(2147483644) {
					v72 = v105
					v73 = v103
					v80 = v107
					continue
				} else {
					break
				}
				break
			}
			if v60 == int32(0) {
				v158 = v103
			} else {
				v116 = v105
				v117 = v103
				v130 = v116
				v131 = v117
				v135 = v6
				for {
					v142 = *(*int32)(unsafe.Add(mBase, uint32(l4+v130<<(uint(int32(2))%32))))
					v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+6)))
					v146 = v131 + v143&int32(_a_F_spgFormInnerTuple_0)
					v147 = int32(1)
					v150 = v135 + v147
					if v150 != v60 {
						v130 = v130 + v147
						v131 = v146
						v135 = v150
						continue
					} else {
						break
					}
					break
				}
				v158 = v146
			}
		}
	}
	v166 = int32(16)
	if base.Ui32(v158) <= base.Ui32(v166) {
		v169 = v166
	} else {
		v169 = v158
	}
	if base.Ui32(v158) < base.Ui32(int32(_a_F_spgFormInnerTuple_1)) {
		if base.B2i32(int32(_a_F_spgFormInnerTuple_0) < l3)|base.B2i32(base.Ui32(int32(_a_F_spgFormInnerTuple_2)) <= base.Ui32(v54)) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v361 = m.ExcPending
			if v361 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_spgFormInnerTuple_3), int32(0))
				mBase = m.M
				v365 = m.ExcPending
				if v365 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_spgFormInnerTuple_4), int32(1048), int32(_a_F_spgFormInnerTuple_5))
					mBase = m.M
					v370 = m.ExcPending
					if v370 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v177 = F_palloc0(m, v169)
			mBase = m.M
			v180 = m.ExcPending
			if v180 != 0 {
				return int32(0)
			} else {
				*(*uint16)(unsafe.Add(mBase, uint32(v177)+4)) = uint16(v169)
				v186 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
				*(*int32)(unsafe.Add(mBase, uint32(v177))) = l3<<(uint(int32(3))%32)&int32(_a_F_spgFormInnerTuple_6) | v186&int32(7) | v54<<(uint(int32(16))%32)
				if l1 == int32(0) {
				} else {
					v197 = v177 + int32(8)
					v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+50)))
					if v198 == int32(1) {
						*(*int32)(unsafe.Add(mBase, uint32(v197))) = l2
					} else {
						v202 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+48)))
						if int32(0) < v202 {
							v230 = v202
						} else {
							v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
							if v205 == int32(1) {
								v209 = int32(18)
								v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
								if v211 == v209 {
									v214 = v209
								} else {
									v214 = int32(2)
								}
								if base.Ui32((v211-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v221 = int32(6)
								} else {
									v221 = v214
								}
								v230 = v221
							} else {
								if v205&int32(1) != 0 {
									v230 = int32(base.Ui32(v205) >> (uint(int32(1)) % 32))
								} else {
									v226 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
									v230 = int32(base.Ui32(v226) >> (uint(int32(2)) % 32))
								}
							}
						}
						if v230 == int32(0) {
						} else {
							if v54 != 0 {
								v234 = v197
							} else {
								v234 = int32(0)
							}
							base.MemoryCopy(m, v234, l2, v230)
						}
					}
				}
				if l3 <= int32(0) {
				} else {
					v243 = v177 + v54 + int32(8)
					v244 = int32(0)
					if l3 != int32(1) {
						v257 = v244
						v258 = v243
						v262 = int32(0)
						for {
							v268 = l4 + v257<<(uint(int32(2))%32)
							v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
							v270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v269)+6)))
							v272 = v270 & int32(_a_F_spgFormInnerTuple_0)
							if v272 != 0 {
								base.MemoryCopy(m, v258, v269, v272)
							} else {
							}
							v274 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v269)+6)))
							v275 = int32(_a_F_spgFormInnerTuple_0)
							v277 = v258 + v274&v275
							v278 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
							v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v278)+6)))
							v281 = v279 & v275
							if v281 != 0 {
								base.MemoryCopy(m, v277, v278, v281)
							} else {
							}
							v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v278)+6)))
							v286 = v277 + v283&int32(_a_F_spgFormInnerTuple_0)
							v287 = int32(2)
							v288 = v257 + v287
							v290 = v262 + v287
							if v290 != l3&int32(2147483646) {
								v257 = v288
								v258 = v286
								v262 = v290
								continue
							} else {
								break
							}
							break
						}
						if l3&int32(1) == int32(0) {
						} else {
							v299 = v288
							v300 = v286
							v311 = *(*int32)(unsafe.Add(mBase, uint32(l4+v299<<(uint(int32(2))%32))))
							v312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v311)+6)))
							v314 = v312 & int32(_a_F_spgFormInnerTuple_0)
							if v314 == int32(0) {
							} else {
								base.MemoryCopy(m, v300, v311, v314)
							}
						}
					} else {
						v299 = v244
						v300 = v243
						v311 = *(*int32)(unsafe.Add(mBase, uint32(l4+v299<<(uint(int32(2))%32))))
						v312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v311)+6)))
						v314 = v312 & int32(_a_F_spgFormInnerTuple_0)
						if v314 == int32(0) {
						} else {
							base.MemoryCopy(m, v300, v311, v314)
						}
					}
				}
				m.G0 = v17 + int32(16)
				return v177
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v339 = m.ExcPending
		if v339 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v342 = m.ExcPending
			if v342 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = int32(_a_F_spgFormInnerTuple_7)
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = v169
				F_errmsg(m, int32(_a_F_spgFormInnerTuple_8), v17)
				mBase = m.M
				v348 = m.ExcPending
				if v348 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(_a_F_spgFormInnerTuple_9), int32(0))
					mBase = m.M
					v352 = m.ExcPending
					if v352 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_spgFormInnerTuple_4), int32(1039), int32(_a_F_spgFormInnerTuple_5))
						mBase = m.M
						v357 = m.ExcPending
						if v357 != 0 {
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
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int64
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int64
	_ = v252
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
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
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
	v301 = m.ExcPending
	if v301 != 0 {
		goto L6
	} else {
		goto L85
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
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
	v274 = v12
	goto L5
L5:
	;
	m.G0 = v10 + int32(16)
	return v274
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
	F_errfinish(m, int32(_a_F_spgGetCache_3), int32(160), int32(_a_F_spgGetCache_4))
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
	F_errfinish(m, int32(_a_F_spgGetCache_3), int32(154), int32(_a_F_spgGetCache_4))
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
	v160 = F_FunctionCall2Coll(m, v154, v157, v10+int32(12), v17)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v162 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v178 = v17 + int32(16)
	F_fillTypeDesc(m, v178, v147)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L6
	} else {
		goto L60
	}
L56:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163+v164<<(uint(int32(4))%32))+88))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v168
	if v147 == v168 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v171 = F_IsBinaryCoercible(m, v168, v147)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	if v171 == int32(0) {
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
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v147 != v181 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	F_fillTypeDesc(m, v17+int32(40), v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L6
	} else {
		goto L68
	}
L62:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v186)+6)))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v185+v187*int32(0)<<(uint(int32(2))%32)+int32(24)-int32(4))))
	goto L65
L63:
	;
	goto L64
L64:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v207
	v209 = *(*int64)(unsafe.Add(mBase, uint32(v178)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+28)) = v209
	goto L61
L65:
	;
	if v199 == int32(0) {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_fillTypeDesc(m, v17+int32(28), v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	goto L61
L68:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	F_fillTypeDesc(m, v17+int32(52), v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+119)))
	if v222 != int32(73) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v226 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
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
	v274 = v17
	goto L5
L73:
	;
	F_LockBuffer(m, v226, int32(1))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	if v226 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+24))
	if v249 != int32(-1173640210) {
		goto L1
	} else {
		goto L79
	}
L76:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_spgGetCache[0]))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v234+(v226^int32(-1))<<(uint(int32(2))%32))))
	v248 = v240
	goto L75
L77:
	;
	goto L78
L78:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_spgGetCache[1]))
	v248 = v242 + v226<<(uint(int32(13))%32) + int32(-8192)
	goto L75
L79:
	;
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v248)+84))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = v252
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v248)+76))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = v254
	v256 = *(*int64)(unsafe.Add(mBase, uint32(v248)+68))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+104)) = v256
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v248)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+96)) = v258
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v248)+52))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+88)) = v260
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v248)+44))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = v262
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v248)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+72)) = v264
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v248)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = v266
	F_UnlockReleaseBuffer(m, v226)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
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
	v288 = m.ExcPending
	if v288 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(_a_F_spgGetCache_5), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_spgGetCache_3), int32(251), int32(_a_F_spgGetCache_6))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
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
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v302 + int32(4)
	F_errmsg_internal(m, int32(_a_F_spgGetCache_7), v10)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_spgGetCache_3), int32(280), int32(_a_F_spgGetCache_6))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
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
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	v5 = l4
	v9 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2+l3<<(uint(int32(2))%32))+20))
	v24 = l2 + v21&int32(_a_F_spgTestLeafTuple_0)
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
				F_errmsg_internal(m, int32(_a_F_spgTestLeafTuple_1), v16+int32(16))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_spgTestLeafTuple_2), int32(798), int32(_a_F_spgTestLeafTuple_3))
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
				v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v31)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v33
				v256 = int32(2049)
				m.G0 = v16 + int32(80)
				return v256
			case 1:
				v256 = int32(0)
				m.G0 = v16 + int32(80)
				return v256
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
					*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v43 & int32(3)
					F_errmsg_internal(m, int32(_a_F_spgTestLeafTuple_1), v16+int32(16))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_spgTestLeafTuple_2), int32(798), int32(_a_F_spgTestLeafTuple_3))
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
				v127 = int32(_a_F_spgTestLeafTuple_4)
				v128 = *(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0]))
				v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
				*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v130
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
						if v142 <= int32(0) {
						} else {
							v146 = v142 << (uint(int32(3)) % 32)
							if v146 == int32(0) {
							} else {
								base.MemoryCopy(m, v139+int32(40), v122, v146)
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = v132
						v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
						*(*uint16)(unsafe.Add(mBase, uint32(v139)+32)) = uint16(v154)
						v156 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
						*(*int32)(unsafe.Add(mBase, uint32(v139)+28)) = v156
						v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
						if v158 == int32(0) {
							v175 = v139
							*(*int64)(unsafe.Add(mBase, uint32(v175)+12)) = int64(0)
							v224 = v175
							v228 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
							*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
							v233 = v120 & v228
							*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
							v236 = v123 & v228
							*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
							v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							F_pairingheap_add(m, v238, v224)
							mBase = m.M
							v240 = m.ExcPending
							if v240 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
								v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
								v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
								m.G0 = v16 + int32(80)
								return v256
							}
						} else {
							v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
							v162 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
							v163 = F_datumCopy(m, v121, v161, v162)
							mBase = m.M
							v164 = m.ExcPending
							if v164 != 0 {
								return int32(0)
							} else {
								v202 = v139
								v204 = v163
								*(*int32)(unsafe.Add(mBase, uint32(v202)+12)) = v204
								v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
								v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
								if int32(2) <= v207 {
									v210 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
									v213 = F_palloc(m, int32(base.Ui32(v210)>>(uint(int32(2))%32)))
									mBase = m.M
									v214 = m.ExcPending
									if v214 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = v213
										v216 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
										v218 = int32(base.Ui32(v216) >> (uint(int32(2)) % 32))
										if v218 == int32(0) {
											v224 = v202
										} else {
											base.MemoryCopy(m, v213, v24, v218)
											v224 = v202
										}
										v228 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
										*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
										v233 = v120 & v228
										*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
										v236 = v123 & v228
										*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
										v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
										F_pairingheap_add(m, v238, v224)
										mBase = m.M
										v240 = m.ExcPending
										if v240 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
											v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
											v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
											m.G0 = v16 + int32(80)
											return v256
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = int32(0)
									v224 = v202
									v228 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
									*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
									v233 = v120 & v228
									*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
									v236 = v123 & v228
									*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
									v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
									F_pairingheap_add(m, v238, v224)
									mBase = m.M
									v240 = m.ExcPending
									if v240 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
										v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
										v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
										m.G0 = v16 + int32(80)
										return v256
									}
								}
							}
						}
					}
				} else {
					v166 = F_palloc(m, int32(40))
					mBase = m.M
					v167 = m.ExcPending
					if v167 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v166)+24)) = v132
						*(*uint8)(unsafe.Add(mBase, uint32(v166)+34)) = uint8(v5)
						v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
						*(*uint16)(unsafe.Add(mBase, uint32(v166)+32)) = uint16(v170)
						v172 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
						*(*int32)(unsafe.Add(mBase, uint32(v166)+28)) = v172
						v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
						if v174 != 0 {
							v202 = v166
							v204 = v9
							*(*int32)(unsafe.Add(mBase, uint32(v202)+12)) = v204
							v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
							v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
							if int32(2) <= v207 {
								v210 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
								v213 = F_palloc(m, int32(base.Ui32(v210)>>(uint(int32(2))%32)))
								mBase = m.M
								v214 = m.ExcPending
								if v214 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = v213
									v216 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
									v218 = int32(base.Ui32(v216) >> (uint(int32(2)) % 32))
									if v218 == int32(0) {
										v224 = v202
									} else {
										base.MemoryCopy(m, v213, v24, v218)
										v224 = v202
									}
									v228 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
									*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
									v233 = v120 & v228
									*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
									v236 = v123 & v228
									*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
									v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
									F_pairingheap_add(m, v238, v224)
									mBase = m.M
									v240 = m.ExcPending
									if v240 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
										v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
										v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
										m.G0 = v16 + int32(80)
										return v256
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = int32(0)
								v224 = v202
								v228 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
								*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
								v233 = v120 & v228
								*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
								v236 = v123 & v228
								*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
								v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
								F_pairingheap_add(m, v238, v224)
								mBase = m.M
								v240 = m.ExcPending
								if v240 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
									v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
									v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
									m.G0 = v16 + int32(80)
									return v256
								}
							}
						} else {
							v175 = v166
							*(*int64)(unsafe.Add(mBase, uint32(v175)+12)) = int64(0)
							v224 = v175
							v228 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
							*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
							v233 = v120 & v228
							*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
							v236 = v123 & v228
							*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
							v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
							F_pairingheap_add(m, v238, v224)
							mBase = m.M
							v240 = m.ExcPending
							if v240 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
								v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
								v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
								m.G0 = v16 + int32(80)
								return v256
							}
						}
					}
				}
			} else {
				v183 = int32(0)
				m.T0[l7].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32))(m, l0, v24+int32(6), v121, v5, v24, v123&int32(1), v183, v183)
				mBase = m.M
				v186 = m.ExcPending
				if v186 != 0 {
					return int32(0)
				} else {
					v187 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v187)
					v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
					v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
					m.G0 = v16 + int32(80)
					return v256
				}
			}
		} else {
			v59 = int32(_a_F_spgTestLeafTuple_4)
			v60 = *(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0]))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
			*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v62
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
					*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v60
					if v109 == int32(0) {
						v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
						v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
						m.G0 = v16 + int32(80)
						return v256
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
							v127 = int32(_a_F_spgTestLeafTuple_4)
							v128 = *(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0]))
							v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v130
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
									if v142 <= int32(0) {
									} else {
										v146 = v142 << (uint(int32(3)) % 32)
										if v146 == int32(0) {
										} else {
											base.MemoryCopy(m, v139+int32(40), v122, v146)
										}
									}
									*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = v132
									v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
									*(*uint16)(unsafe.Add(mBase, uint32(v139)+32)) = uint16(v154)
									v156 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
									*(*int32)(unsafe.Add(mBase, uint32(v139)+28)) = v156
									v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
									if v158 == int32(0) {
										v175 = v139
										*(*int64)(unsafe.Add(mBase, uint32(v175)+12)) = int64(0)
										v224 = v175
										v228 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
										*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
										v233 = v120 & v228
										*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
										v236 = v123 & v228
										*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
										v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
										F_pairingheap_add(m, v238, v224)
										mBase = m.M
										v240 = m.ExcPending
										if v240 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
											v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
											v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
											m.G0 = v16 + int32(80)
											return v256
										}
									} else {
										v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
										v162 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
										v163 = F_datumCopy(m, v121, v161, v162)
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int32(0)
										} else {
											v202 = v139
											v204 = v163
											*(*int32)(unsafe.Add(mBase, uint32(v202)+12)) = v204
											v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
											if int32(2) <= v207 {
												v210 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
												v213 = F_palloc(m, int32(base.Ui32(v210)>>(uint(int32(2))%32)))
												mBase = m.M
												v214 = m.ExcPending
												if v214 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = v213
													v216 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
													v218 = int32(base.Ui32(v216) >> (uint(int32(2)) % 32))
													if v218 == int32(0) {
														v224 = v202
													} else {
														base.MemoryCopy(m, v213, v24, v218)
														v224 = v202
													}
													v228 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
													*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
													v233 = v120 & v228
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
													v236 = v123 & v228
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
													v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v238, v224)
													mBase = m.M
													v240 = m.ExcPending
													if v240 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
														v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
														v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
														m.G0 = v16 + int32(80)
														return v256
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = int32(0)
												v224 = v202
												v228 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
												*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
												v233 = v120 & v228
												*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
												v236 = v123 & v228
												*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
												v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v238, v224)
												mBase = m.M
												v240 = m.ExcPending
												if v240 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
													v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
													v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
													m.G0 = v16 + int32(80)
													return v256
												}
											}
										}
									}
								}
							} else {
								v166 = F_palloc(m, int32(40))
								mBase = m.M
								v167 = m.ExcPending
								if v167 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v166)+24)) = v132
									*(*uint8)(unsafe.Add(mBase, uint32(v166)+34)) = uint8(v5)
									v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
									*(*uint16)(unsafe.Add(mBase, uint32(v166)+32)) = uint16(v170)
									v172 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
									*(*int32)(unsafe.Add(mBase, uint32(v166)+28)) = v172
									v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
									if v174 != 0 {
										v202 = v166
										v204 = v9
										*(*int32)(unsafe.Add(mBase, uint32(v202)+12)) = v204
										v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
										v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
										if int32(2) <= v207 {
											v210 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
											v213 = F_palloc(m, int32(base.Ui32(v210)>>(uint(int32(2))%32)))
											mBase = m.M
											v214 = m.ExcPending
											if v214 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = v213
												v216 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
												v218 = int32(base.Ui32(v216) >> (uint(int32(2)) % 32))
												if v218 == int32(0) {
													v224 = v202
												} else {
													base.MemoryCopy(m, v213, v24, v218)
													v224 = v202
												}
												v228 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
												*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
												v233 = v120 & v228
												*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
												v236 = v123 & v228
												*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
												v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v238, v224)
												mBase = m.M
												v240 = m.ExcPending
												if v240 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
													v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
													v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
													m.G0 = v16 + int32(80)
													return v256
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = int32(0)
											v224 = v202
											v228 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
											*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
											v233 = v120 & v228
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
											v236 = v123 & v228
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
											v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
											F_pairingheap_add(m, v238, v224)
											mBase = m.M
											v240 = m.ExcPending
											if v240 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
												v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
												v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
												m.G0 = v16 + int32(80)
												return v256
											}
										}
									} else {
										v175 = v166
										*(*int64)(unsafe.Add(mBase, uint32(v175)+12)) = int64(0)
										v224 = v175
										v228 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
										*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
										v233 = v120 & v228
										*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
										v236 = v123 & v228
										*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
										v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
										F_pairingheap_add(m, v238, v224)
										mBase = m.M
										v240 = m.ExcPending
										if v240 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
											v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
											v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
											m.G0 = v16 + int32(80)
											return v256
										}
									}
								}
							}
						} else {
							v183 = int32(0)
							m.T0[l7].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32))(m, l0, v24+int32(6), v121, v5, v24, v123&int32(1), v183, v183)
							mBase = m.M
							v186 = m.ExcPending
							if v186 != 0 {
								return int32(0)
							} else {
								v187 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v187)
								v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
								v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
								m.G0 = v16 + int32(80)
								return v256
							}
						}
					}
				}
			} else {
				v85 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+36)))
				switch v85&int32(_a_F_spgTestLeafTuple_6) - int32(1) {
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
						*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v60
						if v109 == int32(0) {
							v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
							v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
							m.G0 = v16 + int32(80)
							return v256
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
								v127 = int32(_a_F_spgTestLeafTuple_4)
								v128 = *(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0]))
								v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
								*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v130
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
										if v142 <= int32(0) {
										} else {
											v146 = v142 << (uint(int32(3)) % 32)
											if v146 == int32(0) {
											} else {
												base.MemoryCopy(m, v139+int32(40), v122, v146)
											}
										}
										*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = v132
										v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
										*(*uint16)(unsafe.Add(mBase, uint32(v139)+32)) = uint16(v154)
										v156 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
										*(*int32)(unsafe.Add(mBase, uint32(v139)+28)) = v156
										v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
										if v158 == int32(0) {
											v175 = v139
											*(*int64)(unsafe.Add(mBase, uint32(v175)+12)) = int64(0)
											v224 = v175
											v228 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
											*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
											v233 = v120 & v228
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
											v236 = v123 & v228
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
											v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
											F_pairingheap_add(m, v238, v224)
											mBase = m.M
											v240 = m.ExcPending
											if v240 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
												v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
												v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
												m.G0 = v16 + int32(80)
												return v256
											}
										} else {
											v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
											v162 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
											v163 = F_datumCopy(m, v121, v161, v162)
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												v202 = v139
												v204 = v163
												*(*int32)(unsafe.Add(mBase, uint32(v202)+12)) = v204
												v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
												if int32(2) <= v207 {
													v210 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
													v213 = F_palloc(m, int32(base.Ui32(v210)>>(uint(int32(2))%32)))
													mBase = m.M
													v214 = m.ExcPending
													if v214 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = v213
														v216 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
														v218 = int32(base.Ui32(v216) >> (uint(int32(2)) % 32))
														if v218 == int32(0) {
															v224 = v202
														} else {
															base.MemoryCopy(m, v213, v24, v218)
															v224 = v202
														}
														v228 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
														*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
														v233 = v120 & v228
														*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
														v236 = v123 & v228
														*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
														v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
														F_pairingheap_add(m, v238, v224)
														mBase = m.M
														v240 = m.ExcPending
														if v240 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
															v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
															v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
															m.G0 = v16 + int32(80)
															return v256
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = int32(0)
													v224 = v202
													v228 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
													*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
													v233 = v120 & v228
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
													v236 = v123 & v228
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
													v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v238, v224)
													mBase = m.M
													v240 = m.ExcPending
													if v240 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
														v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
														v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
														m.G0 = v16 + int32(80)
														return v256
													}
												}
											}
										}
									}
								} else {
									v166 = F_palloc(m, int32(40))
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v166)+24)) = v132
										*(*uint8)(unsafe.Add(mBase, uint32(v166)+34)) = uint8(v5)
										v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
										*(*uint16)(unsafe.Add(mBase, uint32(v166)+32)) = uint16(v170)
										v172 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
										*(*int32)(unsafe.Add(mBase, uint32(v166)+28)) = v172
										v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
										if v174 != 0 {
											v202 = v166
											v204 = v9
											*(*int32)(unsafe.Add(mBase, uint32(v202)+12)) = v204
											v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
											if int32(2) <= v207 {
												v210 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
												v213 = F_palloc(m, int32(base.Ui32(v210)>>(uint(int32(2))%32)))
												mBase = m.M
												v214 = m.ExcPending
												if v214 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = v213
													v216 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
													v218 = int32(base.Ui32(v216) >> (uint(int32(2)) % 32))
													if v218 == int32(0) {
														v224 = v202
													} else {
														base.MemoryCopy(m, v213, v24, v218)
														v224 = v202
													}
													v228 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
													*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
													v233 = v120 & v228
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
													v236 = v123 & v228
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
													v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v238, v224)
													mBase = m.M
													v240 = m.ExcPending
													if v240 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
														v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
														v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
														m.G0 = v16 + int32(80)
														return v256
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = int32(0)
												v224 = v202
												v228 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
												*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
												v233 = v120 & v228
												*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
												v236 = v123 & v228
												*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
												v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v238, v224)
												mBase = m.M
												v240 = m.ExcPending
												if v240 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
													v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
													v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
													m.G0 = v16 + int32(80)
													return v256
												}
											}
										} else {
											v175 = v166
											*(*int64)(unsafe.Add(mBase, uint32(v175)+12)) = int64(0)
											v224 = v175
											v228 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
											*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
											v233 = v120 & v228
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
											v236 = v123 & v228
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
											v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
											F_pairingheap_add(m, v238, v224)
											mBase = m.M
											v240 = m.ExcPending
											if v240 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
												v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
												v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
												m.G0 = v16 + int32(80)
												return v256
											}
										}
									}
								}
							} else {
								v183 = int32(0)
								m.T0[l7].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32))(m, l0, v24+int32(6), v121, v5, v24, v123&int32(1), v183, v183)
								mBase = m.M
								v186 = m.ExcPending
								if v186 != 0 {
									return int32(0)
								} else {
									v187 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v187)
									v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
									v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
									m.G0 = v16 + int32(80)
									return v256
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
						*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v60
						if v109 == int32(0) {
							v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
							v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
							m.G0 = v16 + int32(80)
							return v256
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
								v127 = int32(_a_F_spgTestLeafTuple_4)
								v128 = *(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0]))
								v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
								*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v130
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
										if v142 <= int32(0) {
										} else {
											v146 = v142 << (uint(int32(3)) % 32)
											if v146 == int32(0) {
											} else {
												base.MemoryCopy(m, v139+int32(40), v122, v146)
											}
										}
										*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = v132
										v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
										*(*uint16)(unsafe.Add(mBase, uint32(v139)+32)) = uint16(v154)
										v156 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
										*(*int32)(unsafe.Add(mBase, uint32(v139)+28)) = v156
										v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
										if v158 == int32(0) {
											v175 = v139
											*(*int64)(unsafe.Add(mBase, uint32(v175)+12)) = int64(0)
											v224 = v175
											v228 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
											*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
											v233 = v120 & v228
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
											v236 = v123 & v228
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
											v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
											F_pairingheap_add(m, v238, v224)
											mBase = m.M
											v240 = m.ExcPending
											if v240 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
												v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
												v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
												m.G0 = v16 + int32(80)
												return v256
											}
										} else {
											v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
											v162 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
											v163 = F_datumCopy(m, v121, v161, v162)
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												v202 = v139
												v204 = v163
												*(*int32)(unsafe.Add(mBase, uint32(v202)+12)) = v204
												v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
												if int32(2) <= v207 {
													v210 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
													v213 = F_palloc(m, int32(base.Ui32(v210)>>(uint(int32(2))%32)))
													mBase = m.M
													v214 = m.ExcPending
													if v214 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = v213
														v216 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
														v218 = int32(base.Ui32(v216) >> (uint(int32(2)) % 32))
														if v218 == int32(0) {
															v224 = v202
														} else {
															base.MemoryCopy(m, v213, v24, v218)
															v224 = v202
														}
														v228 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
														*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
														v233 = v120 & v228
														*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
														v236 = v123 & v228
														*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
														v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
														F_pairingheap_add(m, v238, v224)
														mBase = m.M
														v240 = m.ExcPending
														if v240 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
															v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
															v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
															m.G0 = v16 + int32(80)
															return v256
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = int32(0)
													v224 = v202
													v228 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
													*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
													v233 = v120 & v228
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
													v236 = v123 & v228
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
													v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v238, v224)
													mBase = m.M
													v240 = m.ExcPending
													if v240 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
														v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
														v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
														m.G0 = v16 + int32(80)
														return v256
													}
												}
											}
										}
									}
								} else {
									v166 = F_palloc(m, int32(40))
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v166)+24)) = v132
										*(*uint8)(unsafe.Add(mBase, uint32(v166)+34)) = uint8(v5)
										v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
										*(*uint16)(unsafe.Add(mBase, uint32(v166)+32)) = uint16(v170)
										v172 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
										*(*int32)(unsafe.Add(mBase, uint32(v166)+28)) = v172
										v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
										if v174 != 0 {
											v202 = v166
											v204 = v9
											*(*int32)(unsafe.Add(mBase, uint32(v202)+12)) = v204
											v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
											if int32(2) <= v207 {
												v210 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
												v213 = F_palloc(m, int32(base.Ui32(v210)>>(uint(int32(2))%32)))
												mBase = m.M
												v214 = m.ExcPending
												if v214 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = v213
													v216 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
													v218 = int32(base.Ui32(v216) >> (uint(int32(2)) % 32))
													if v218 == int32(0) {
														v224 = v202
													} else {
														base.MemoryCopy(m, v213, v24, v218)
														v224 = v202
													}
													v228 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
													*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
													v233 = v120 & v228
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
													v236 = v123 & v228
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
													v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v238, v224)
													mBase = m.M
													v240 = m.ExcPending
													if v240 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
														v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
														v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
														m.G0 = v16 + int32(80)
														return v256
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = int32(0)
												v224 = v202
												v228 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
												*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
												v233 = v120 & v228
												*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
												v236 = v123 & v228
												*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
												v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v238, v224)
												mBase = m.M
												v240 = m.ExcPending
												if v240 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
													v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
													v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
													m.G0 = v16 + int32(80)
													return v256
												}
											}
										} else {
											v175 = v166
											*(*int64)(unsafe.Add(mBase, uint32(v175)+12)) = int64(0)
											v224 = v175
											v228 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
											*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
											v233 = v120 & v228
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
											v236 = v123 & v228
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
											v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
											F_pairingheap_add(m, v238, v224)
											mBase = m.M
											v240 = m.ExcPending
											if v240 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
												v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
												v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
												m.G0 = v16 + int32(80)
												return v256
											}
										}
									}
								}
							} else {
								v183 = int32(0)
								m.T0[l7].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32))(m, l0, v24+int32(6), v121, v5, v24, v123&int32(1), v183, v183)
								mBase = m.M
								v186 = m.ExcPending
								if v186 != 0 {
									return int32(0)
								} else {
									v187 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v187)
									v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
									v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
									m.G0 = v16 + int32(80)
									return v256
								}
							}
						}
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v192 = m.ExcPending
					if v192 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = v85
						F_errmsg_internal(m, int32(_a_F_spgTestLeafTuple_7), v16)
						mBase = m.M
						v196 = m.ExcPending
						if v196 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_spgTestLeafTuple_8), int32(70), int32(_a_F_spgTestLeafTuple_9))
							mBase = m.M
							v201 = m.ExcPending
							if v201 != 0 {
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
						*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v60
						if v109 == int32(0) {
							v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
							v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
							m.G0 = v16 + int32(80)
							return v256
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
								v127 = int32(_a_F_spgTestLeafTuple_4)
								v128 = *(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0]))
								v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
								*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v130
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
										if v142 <= int32(0) {
										} else {
											v146 = v142 << (uint(int32(3)) % 32)
											if v146 == int32(0) {
											} else {
												base.MemoryCopy(m, v139+int32(40), v122, v146)
											}
										}
										*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = v132
										v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
										*(*uint16)(unsafe.Add(mBase, uint32(v139)+32)) = uint16(v154)
										v156 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
										*(*int32)(unsafe.Add(mBase, uint32(v139)+28)) = v156
										v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
										if v158 == int32(0) {
											v175 = v139
											*(*int64)(unsafe.Add(mBase, uint32(v175)+12)) = int64(0)
											v224 = v175
											v228 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
											*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
											v233 = v120 & v228
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
											v236 = v123 & v228
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
											v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
											F_pairingheap_add(m, v238, v224)
											mBase = m.M
											v240 = m.ExcPending
											if v240 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
												v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
												v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
												m.G0 = v16 + int32(80)
												return v256
											}
										} else {
											v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
											v162 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
											v163 = F_datumCopy(m, v121, v161, v162)
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												v202 = v139
												v204 = v163
												*(*int32)(unsafe.Add(mBase, uint32(v202)+12)) = v204
												v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
												v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
												if int32(2) <= v207 {
													v210 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
													v213 = F_palloc(m, int32(base.Ui32(v210)>>(uint(int32(2))%32)))
													mBase = m.M
													v214 = m.ExcPending
													if v214 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = v213
														v216 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
														v218 = int32(base.Ui32(v216) >> (uint(int32(2)) % 32))
														if v218 == int32(0) {
															v224 = v202
														} else {
															base.MemoryCopy(m, v213, v24, v218)
															v224 = v202
														}
														v228 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
														*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
														v233 = v120 & v228
														*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
														v236 = v123 & v228
														*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
														v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
														F_pairingheap_add(m, v238, v224)
														mBase = m.M
														v240 = m.ExcPending
														if v240 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
															v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
															v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
															m.G0 = v16 + int32(80)
															return v256
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = int32(0)
													v224 = v202
													v228 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
													*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
													v233 = v120 & v228
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
													v236 = v123 & v228
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
													v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v238, v224)
													mBase = m.M
													v240 = m.ExcPending
													if v240 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
														v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
														v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
														m.G0 = v16 + int32(80)
														return v256
													}
												}
											}
										}
									}
								} else {
									v166 = F_palloc(m, int32(40))
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v166)+24)) = v132
										*(*uint8)(unsafe.Add(mBase, uint32(v166)+34)) = uint8(v5)
										v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)))
										*(*uint16)(unsafe.Add(mBase, uint32(v166)+32)) = uint16(v170)
										v172 = *(*int32)(unsafe.Add(mBase, uint32(v24)+6))
										*(*int32)(unsafe.Add(mBase, uint32(v166)+28)) = v172
										v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
										if v174 != 0 {
											v202 = v166
											v204 = v9
											*(*int32)(unsafe.Add(mBase, uint32(v202)+12)) = v204
											v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
											v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
											if int32(2) <= v207 {
												v210 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
												v213 = F_palloc(m, int32(base.Ui32(v210)>>(uint(int32(2))%32)))
												mBase = m.M
												v214 = m.ExcPending
												if v214 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = v213
													v216 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
													v218 = int32(base.Ui32(v216) >> (uint(int32(2)) % 32))
													if v218 == int32(0) {
														v224 = v202
													} else {
														base.MemoryCopy(m, v213, v24, v218)
														v224 = v202
													}
													v228 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
													*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
													v233 = v120 & v228
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
													v236 = v123 & v228
													*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
													v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
													F_pairingheap_add(m, v238, v224)
													mBase = m.M
													v240 = m.ExcPending
													if v240 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
														v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
														v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
														m.G0 = v16 + int32(80)
														return v256
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v202)+16)) = int32(0)
												v224 = v202
												v228 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
												*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
												v233 = v120 & v228
												*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
												v236 = v123 & v228
												*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
												v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
												F_pairingheap_add(m, v238, v224)
												mBase = m.M
												v240 = m.ExcPending
												if v240 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
													v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
													v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
													m.G0 = v16 + int32(80)
													return v256
												}
											}
										} else {
											v175 = v166
											*(*int64)(unsafe.Add(mBase, uint32(v175)+12)) = int64(0)
											v224 = v175
											v228 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+35)) = uint8(v228)
											*(*int32)(unsafe.Add(mBase, uint32(v224)+20)) = int32(0)
											v233 = v120 & v228
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+37)) = uint8(v233)
											v236 = v123 & v228
											*(*uint8)(unsafe.Add(mBase, uint32(v224)+36)) = uint8(v236)
											v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
											F_pairingheap_add(m, v238, v224)
											mBase = m.M
											v240 = m.ExcPending
											if v240 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_spgTestLeafTuple[0])) = v128
												v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
												v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
												m.G0 = v16 + int32(80)
												return v256
											}
										}
									}
								}
							} else {
								v183 = int32(0)
								m.T0[l7].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32))(m, l0, v24+int32(6), v121, v5, v24, v123&int32(1), v183, v183)
								mBase = m.M
								v186 = m.ExcPending
								if v186 != 0 {
									return int32(0)
								} else {
									v187 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v187)
									v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+4)))
									v256 = v253 & int32(_a_F_spgTestLeafTuple_5)
									m.G0 = v16 + int32(80)
									return v256
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
	var v47 int32
	_ = v47
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
			v47 = v42
		} else {
			v47 = v18
		}
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		if v55&int32(1) != 0 {
			v58 = int32(246)
		} else {
			v58 = int32(247)
		}
		F_pg_qsort(m, v14, v47, int32(8), v58)
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
	var v27 int32
	_ = v27
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
	var v111 int32
	_ = v111
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
	return v111
L2:
	;
	v27 = v2
	goto L5
L3:
	;
	goto L4
L4:
	;
	v101 = int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v102 <= int32(0) {
		v111 = v101
		goto L1
	} else {
		goto L32
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v32 = v29 + v27*int32(48)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+6)))
	switch v34 - int32(1) {
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
	v91 = v27 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v91 < v92 {
		v27 = v91
		goto L5
	} else {
		goto L31
	}
L8:
	;
	v82 = int32(0)
	v85 = F_DirectFunctionCall2Coll(m, int32(250), v82, v13, v33)
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
	v40 = F_DirectFunctionCall2Coll(m, int32(253), v37, v13, v33)
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
	v111 = v37
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
	v111 = v44
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
	v111 = v49
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
	v111 = v54
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
	v111 = v59
	goto L1
L26:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68+v27*int32(48))+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v72
	F_errmsg_internal(m, int32(_a_F_spg_quad_leaf_consistent_0), v10)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_spg_quad_leaf_consistent_1), int32(457), int32(_a_F_spg_quad_leaf_consistent_2))
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
		v111 = v82
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
	v111 = v101
	goto L1
}
func F_spg_range_quad_choose(m *base.Module, l0 int32) int32 {
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
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+12)))
		if v20 == int32(1) {
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(1)
			v93 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v16
			*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v93
			m.G0 = v11 + int32(48)
			return int32(0)
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
			v27 = F_range_get_typcache(m, l0, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+13)))
				if v29 == int32(0) {
					v32 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v32
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v16+int32(base.Ui32(v35)>>(uint(int32(2))%32))-v32))))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = (v41 ^ int32(-1)) & int32(1)
					v93 = v32
					*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v16
					*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v93
					m.G0 = v11 + int32(48)
					return int32(0)
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
					v48 = F_pg_detoast_datum(m, v47)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						v51 = v11 + int32(40)
						v53 = v11 + int32(32)
						F_range_deserialize(m, v27, v48, v51, v53, v11+int32(31))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							v59 = v11 + int32(20)
							v61 = v11 + int32(12)
							F_range_deserialize(m, v27, v16, v59, v61, v11+int32(11))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+11)))
								if v67 != 0 {
									v86 = int32(5)
									v87 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v86 - v87
									*(*int32)(unsafe.Add(mBase, uint32(v13))) = v87
									v93 = v87
									*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v16
									*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v93
									m.G0 = v11 + int32(48)
									return int32(0)
								} else {
									v68 = F_range_cmp_bounds(m, v27, v59, v51)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										v72 = F_range_cmp_bounds(m, v27, v61, v53)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
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
											*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v86 - v87
											*(*int32)(unsafe.Add(mBase, uint32(v13))) = v87
											v93 = v87
											*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v16
											*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v93
											m.G0 = v11 + int32(48)
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
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
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
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
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v253 int32
	_ = v253
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
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
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v32 == int32(18) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v43 = int32(1)
	if v26 != 0 {
		v53 = int32(base.Ui32(v24)>>(uint(v43)%32)) - v43
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v35 = int32(16)
	goto L9
L8:
	;
	v35 = int32(0)
	goto L9
L9:
	;
	if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v42 = int32(4)
	goto L12
L11:
	;
	v42 = v35
	goto L12
L12:
	;
	v53 = v42
	goto L3
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L14:
	;
	v54 = v23
	goto L16
L15:
	;
	v54 = v16 + int32(4)
	goto L16
L16:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+13)))
	if v55 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	if int32(0) < v267 {
		goto L79
	} else {
		goto L80
	}
L18:
	;
	if v71 <= v98 {
		v261 = v98
		v266 = int32(_a_F_spg_text_choose_0)
		goto L17
	} else {
		goto L76
	}
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v59 = F_pg_detoast_datum_packed(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	if v53 <= v234 {
		v261 = v2
		v266 = int32(_a_F_spg_text_choose_0)
		goto L17
	} else {
		goto L75
	}
L22:
	;
	v63 = int32(1)
	v64 = v59 + v63
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	v67 = v65 & v63
	if v67 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v68 = v64
	goto L25
L24:
	;
	v68 = v59 + int32(4)
	goto L25
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v70 = v54 + v69
	v71 = v53 - v69
	if v65 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(1)
	v188 = F_palloc(m, int32(4))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L63
	}
L27:
	;
	v151 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)) = uint8(v151)
	v155 = v127 + int32(4)
	v156 = F_palloc(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L55
	}
L28:
	;
	v148 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)) = uint8(v148)
	v174 = v148
	goto L26
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(3)
	if v127 != 0 {
		goto L27
	} else {
		goto L54
	}
L30:
	;
	if v71 < v98 {
		goto L41
	} else {
		goto L42
	}
L31:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	if v77 == int32(18) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v88 = int32(1)
	if v67 != 0 {
		v98 = int32(base.Ui32(v65)>>(uint(v88)%32)) - v88
		goto L30
	} else {
		goto L40
	}
L34:
	;
	v80 = int32(16)
	goto L36
L35:
	;
	v80 = int32(0)
	goto L36
L36:
	;
	if base.Ui32((v77-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v87 = int32(4)
	goto L39
L38:
	;
	v87 = v80
	goto L39
L39:
	;
	v98 = v87
	goto L30
L40:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v98 = int32(base.Ui32(v92)>>(uint(int32(2))%32)) - int32(4)
	goto L30
L41:
	;
	v100 = v71
	goto L43
L42:
	;
	v100 = v98
	goto L43
L43:
	;
	if int32(0) < v100 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v104 = v70
	v105 = int32(0)
	v107 = v68
	goto L48
L45:
	;
	goto L46
L46:
	;
	if v98 == int32(0) {
		goto L18
	} else {
		goto L53
	}
L47:
	;
	if v127 != v98 {
		goto L29
	} else {
		goto L52
	}
L48:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v116 != v117 {
		v127 = v105
		goto L47
	} else {
		goto L50
	}
L49:
	;
	v127 = v100
	goto L47
L50:
	;
	v119 = int32(1)
	v124 = v105 + v119
	if v124 != v100 {
		v104 = v104 + v119
		v105 = v124
		v107 = v107 + v119
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
	goto L28
L54:
	;
	goto L28
L55:
	;
	if base.Ui32(v127) <= base.Ui32(int32(126)) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v127 != 0 {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v163 = v127<<(uint(int32(1))%32) + int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v163)
	v169 = v151
	goto L56
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v155 << (uint(int32(2)) % 32)
	v169 = int32(4)
	goto L56
L60:
	;
	base.MemoryCopy(m, v156+v169, v68, v127)
	goto L62
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v156
	v174 = v127
	goto L26
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v188
	v191 = v174 + v68
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(0)
	v196 = v98 - v174
	if v196 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v199 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)) = uint8(v199)
	return v199
L65:
	;
	goto L66
L66:
	;
	v203 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)) = uint8(v203)
	v206 = v196 - v203
	v208 = v196 + int32(3)
	v209 = F_palloc(m, v208)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if base.Ui32(v196) <= base.Ui32(int32(127)) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v209
	return int32(0)
L69:
	;
	if v206 == int32(0) {
		goto L68
	} else {
		goto L74
	}
L70:
	;
	v213 = int32(1)
	v216 = v196<<(uint(v213)%32) | v213
	*(*uint8)(unsafe.Add(mBase, uint32(v209))) = uint8(v216)
	if v206 != 0 {
		v223 = v213
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v208 << (uint(int32(2)) % 32)
	v223 = int32(4)
	goto L69
L73:
	;
	goto L68
L74:
	;
	base.MemoryCopy(m, v223+v209, v191+int32(1), v206)
	goto L68
L75:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234+v54))))
	v261 = v2
	v266 = v237
	goto L17
L76:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v98))))
	v261 = v98
	v266 = v253
	goto L17
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v300
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = base.I32_extend16_s(v266)
	return int32(0)
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v287
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(1)
	v338 = int32(0)
	v340 = v261 + base.B2i32(v338 <= v284)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v340
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v344 = v53 - (v340 + v342)
	if v338 < v344 {
		goto L92
	} else {
		goto L93
	}
L79:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v272 = v267
	v275 = int32(0)
	goto L82
L80:
	;
	v300 = v267
	goto L81
L81:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+12)))
	if v312 != int32(1) {
		goto L77
	} else {
		goto L90
	}
L82:
	;
	v284 = base.I32_extend16_s(v266)
	v287 = (v272 + v275) >> (uint(int32(1)) % 32)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v270+v287<<(uint(int32(2))%32))))
	v292 = base.I32_extend16_s(v291)
	if v284 < v292 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v300 = v297
	goto L81
L84:
	;
	if v298 < v297 {
		v272 = v297
		v275 = v298
		goto L82
	} else {
		goto L89
	}
L85:
	;
	v297 = v287
	v298 = v275
	goto L84
L86:
	;
	goto L87
L87:
	;
	if v284 <= v292 {
		goto L78
	} else {
		goto L88
	}
L88:
	;
	v297 = v272
	v298 = v287 + int32(1)
	goto L84
L89:
	;
	goto L83
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(3)
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)) = uint8(v317)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v319
	v324 = F_palloc(m, int32(4))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v324
	*(*int32)(unsafe.Add(mBase, uint32(v324))) = int32(-2)
	v329 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)) = uint8(v329)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v329
	return v329
L92:
	;
	v347 = int32(4)
	v349 = v344 + v347
	v350 = F_palloc(m, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v372 = F_palloc(m, int32(4))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L103
	}
L95:
	;
	if base.Ui32(v344) <= base.Ui32(int32(126)) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v344 != 0 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v354 = int32(1)
	v358 = v344<<(uint(v354)%32) + int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v358)
	v363 = v354
	goto L96
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v350))) = v349 << (uint(int32(2)) % 32)
	v363 = v347
	goto L96
L100:
	;
	base.MemoryCopy(m, v350+v363, v342+v54+v340, v344)
	goto L102
L101:
	;
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v350
	return int32(0)
L103:
	;
	v374 = int32(3)
	*(*uint8)(unsafe.Add(mBase, uint32(v372))) = uint8(v374)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v372
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
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v163 int32
	_ = v163
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
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
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
	var v231 int32
	_ = v231
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
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
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
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v59 = v57 + v19
	v60 = int32(0)
	if v57|base.B2i32(v19 <= v60) == v60 {
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
		v57 = int32(base.Ui32(v28)>>(uint(v31)%32)) - v31
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v45 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v57 = int32(base.Ui32(v37)>>(uint(int32(2))%32)) - int32(4)
	goto L1
L8:
	;
	v48 = int32(16)
	goto L10
L9:
	;
	v48 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v45-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v55 = int32(4)
	goto L13
L12:
	;
	v55 = v48
	goto L13
L13:
	;
	v57 = v55
	goto L1
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v141
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v146 <= int32(0) {
		v347 = int32(1)
		goto L39
	} else {
		goto L40
	}
L15:
	;
	v141 = v58
	v143 = v58 + int32(4)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v68 = v59 + int32(4)
	v69 = F_palloc(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v68 << (uint(int32(2)) % 32)
	v75 = v69 + int32(4)
	if v19 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	base.MemoryCopy(m, v75, v58+int32(4), v19)
	goto L21
L20:
	;
	goto L21
L21:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v79 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v135 == int32(0) {
		v141 = v69
		v143 = v75
		goto L14
	} else {
		goto L38
	}
L23:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if base.B2i32(base.Ui32(int32(18)) < base.Ui32(v82))|base.B2i32(int32(1)<<(uint(v82)%32)&int32(_a_F_spg_text_leaf_consistent_0) == int32(0)) != 0 {
		v141 = v69
		v143 = v75
		goto L14
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v79&int32(1) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L26:
	;
	if v82 == int32(18) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v99 = int32(16)
	goto L29
L28:
	;
	v99 = int32(0)
	goto L29
L29:
	;
	if base.Ui32((v82-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v106 = int32(4)
	goto L32
L31:
	;
	v106 = v99
	goto L32
L32:
	;
	v134 = v24 + int32(1)
	v135 = v106
	goto L22
L33:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v111&int32(-4) == int32(16) {
		v141 = v69
		v143 = v75
		goto L14
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v79&int32(254) == int32(2) {
		v141 = v69
		v143 = v75
		goto L14
	} else {
		goto L37
	}
L36:
	;
	v116 = int32(4)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v134 = v24 + v116
	v135 = int32(base.Ui32(v118)>>(uint(int32(2))%32)) - v116
	goto L22
L37:
	;
	v127 = int32(1)
	v134 = v24 + v127
	v135 = int32(base.Ui32(v79)>>(uint(v127)%32)) - v127
	goto L22
L38:
	;
	base.MemoryCopy(m, v19+v75, v134, v135)
	v141 = v69
	v143 = v75
	goto L14
L39:
	;
	m.G0 = v16 + int32(16)
	return v347
L40:
	;
	v153 = int32(0)
	goto L41
L41:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v166 = v163 + v153*int32(48)
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+6)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v166)+44))
	v169 = F_pg_detoast_datum_packed(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L2
	} else {
		goto L44
	}
L42:
	;
	v347 = v340
	goto L39
L43:
	;
	v202 = v167 & int32(_a_F_spg_text_leaf_consistent_1)
	if v202 == int32(28) {
		goto L56
	} else {
		goto L57
	}
L44:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if v171 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)))
	if v177 == int32(18) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v188 = int32(1)
	if v171&v188 != 0 {
		v200 = int32(base.Ui32(v171)>>(uint(v188)%32)) - v188
		goto L43
	} else {
		goto L54
	}
L48:
	;
	v180 = int32(16)
	goto L50
L49:
	;
	v180 = int32(0)
	goto L50
L50:
	;
	if base.Ui32((v177-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v187 = int32(4)
	goto L53
L52:
	;
	v187 = v180
	goto L53
L53:
	;
	v200 = v187
	goto L43
L54:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v200 = int32(base.Ui32(v194)>>(uint(int32(2))%32)) - int32(4)
	goto L43
L55:
	;
	v340 = int32(1)
	v342 = v153 + v340
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v342 < v343 {
		v153 = v342
		goto L41
	} else {
		goto L109
	}
L56:
	;
	if v200 <= v19 {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if base.Ui32(int32(11)) <= base.Ui32(v202) {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v209 = F_DirectFunctionCall2Coll(m, int32(260), v207, v208, v169)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	if v209 != 0 {
		goto L55
	} else {
		goto L61
	}
L61:
	;
	v347 = int32(0)
	goto L39
L62:
	;
	switch v298&int32(_a_F_spg_text_leaf_consistent_1) - int32(1) {
	case 0:
		goto L100
	case 1:
		goto L95
	case 2:
		goto L99
	case 3:
		goto L98
	case 4:
		goto L97
	default:
		goto L96
	}
L63:
	;
	v216 = int32(1)
	if v171&v216 != 0 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v225 = int32(1)
	if v171&v225 != 0 {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	v220 = v216
	goto L68
L67:
	;
	v220 = int32(4)
	goto L68
L68:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v223 = F_varstr_cmp(m, v143, v59, v169+v220, v200, v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	v297 = v223
	v298 = v167 - int32(10)
	goto L62
L70:
	;
	v229 = v225
	goto L72
L71:
	;
	v229 = int32(4)
	goto L72
L72:
	;
	v230 = v169 + v229
	v231 = base.B2i32(v200 < v59)
	if v200 < v59 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v232 = v200
	goto L75
L74:
	;
	v232 = v59
	goto L75
L75:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v232) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	if v294 != 0 {
		v297 = v294
		v298 = v167
		goto L62
	} else {
		goto L94
	}
L77:
	;
	v294 = int32(0)
	goto L76
L78:
	;
	v268 = v263
	v269 = v264
	v270 = v265
	goto L88
L79:
	;
	if (v143|v230)&int32(3) != 0 {
		v263 = v143
		v264 = v230
		v265 = v232
		goto L78
	} else {
		goto L82
	}
L80:
	;
	v256 = v143
	v257 = v230
	v258 = v232
	goto L81
L81:
	;
	if v258 == int32(0) {
		goto L77
	} else {
		goto L87
	}
L82:
	;
	v240 = v143
	v241 = v230
	v242 = v232
	goto L83
L83:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	if v245 != v246 {
		v263 = v240
		v264 = v241
		v265 = v242
		goto L78
	} else {
		goto L85
	}
L84:
	;
	v256 = v251
	v257 = v249
	v258 = v253
	goto L81
L85:
	;
	v248 = int32(4)
	v249 = v241 + v248
	v251 = v240 + v248
	v253 = v242 - v248
	if base.Ui32(int32(3)) < base.Ui32(v253) {
		v240 = v251
		v241 = v249
		v242 = v253
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v263 = v256
	v264 = v257
	v265 = v258
	goto L78
L88:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	if v273 == v274 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v294 = v273 - v274
	goto L76
L90:
	;
	v276 = int32(1)
	v281 = v270 - v276
	if v281 != 0 {
		v268 = v268 + v276
		v269 = v269 + v276
		v270 = v281
		goto L88
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	goto L89
L93:
	;
	goto L77
L94:
	;
	v297 = v231 - base.B2i32(v59 < v200)
	v298 = v167
	goto L62
L95:
	;
	v334 = int32(0)
	if v334 < v297 {
		v347 = v334
		goto L39
	} else {
		goto L108
	}
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L2
	} else {
		goto L105
	}
L97:
	;
	v313 = int32(0)
	if v313 < v297 {
		goto L55
	} else {
		goto L104
	}
L98:
	;
	if int32(0) <= v297 {
		goto L55
	} else {
		goto L103
	}
L99:
	;
	if v297 == int32(0) {
		goto L55
	} else {
		goto L102
	}
L100:
	;
	v304 = int32(0)
	if v297 < v304 {
		goto L55
	} else {
		goto L101
	}
L101:
	;
	v347 = v304
	goto L39
L102:
	;
	v347 = int32(0)
	goto L39
L103:
	;
	v347 = int32(0)
	goto L39
L104:
	;
	v347 = v313
	goto L39
L105:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v324 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v320+v153*int32(48))+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v324
	F_errmsg_internal(m, int32(_a_F_spg_text_leaf_consistent_2), v16)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L2
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(_a_F_spg_text_leaf_consistent_3), int32(691), int32(_a_F_spg_text_leaf_consistent_4))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L2
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L108:
	;
	goto L55
L109:
	;
	goto L42
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
