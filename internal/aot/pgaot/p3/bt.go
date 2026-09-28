package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__bt_binsrch_insert(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
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
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	v3 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v20 < v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
	if v39 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F__bt_binsrch_insert[0]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24+(v20^int32(-1))<<(uint(int32(2))%32))))
	v38 = v30
	goto L1
L3:
	;
	goto L4
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F__bt_binsrch_insert[1]))
	v38 = v32 + v20<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L25
	} else {
		goto L58
	}
L6:
	;
	v62 = int32(_a_F__bt_binsrch_insert_0)
	v63 = v61 & v62
	if base.Ui32(v60&v62) < base.Ui32(v63) {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v42) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)))
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)))
	v60 = v57
	v61 = v58
	goto L6
L10:
	;
	v50 = int32(base.Ui32(v42+int32(_a_F__bt_binsrch_insert_1)) >> (uint(int32(2)) % 32))
	goto L12
L11:
	;
	v50 = int32(0)
	goto L12
L12:
	;
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v38+v53)+4))
	if v55 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v56 = int32(2)
	goto L15
L14:
	;
	v56 = int32(1)
	goto L15
L15:
	;
	v60 = v50
	v61 = v56
	goto L6
L16:
	;
	v265 = v61
	v266 = v3
	v268 = v3
	v274 = int32(0)
	goto L18
L17:
	;
	v68 = int32(1)
	v71 = v60 + (v39 ^ v68)
	if base.Ui32(v63) < base.Ui32(v71&int32(_a_F__bt_binsrch_insert_0)) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)) = uint16(v268)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)) = uint16(v274)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v266)
	m.G0 = v18 + int32(32)
	return v265 & int32(_a_F__bt_binsrch_insert_0)
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v82 = v61
	v83 = v71
	v85 = v71
	goto L22
L20:
	;
	v250 = v61
	v251 = v68
	v253 = v71
	goto L21
L21:
	;
	v265 = v250
	v266 = v251
	v268 = v253
	v274 = v250
	goto L18
L22:
	;
	v96 = int32(base.Ui32((v83-v82)&int32(_a_F__bt_binsrch_insert_2))>>(uint(int32(1))%32)) + v82
	v98 = v96 & int32(_a_F__bt_binsrch_insert_0)
	v99 = F__bt_compare(m, l0, v75, v38, v98)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v250 = v239
	v251 = int32(1)
	v253 = v231
	goto L21
L24:
	;
	if v99 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L25:
	;
	return int32(0)
L26:
	;
	if v99 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	if v103 == int32(0) {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v106 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v107 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v38+v98<<(uint(int32(2))%32))+20))
	v114 = v38 + v111&int32(_a_F__bt_binsrch_insert_3)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+7)))
	if v115&int32(32) == v107 {
		v200 = v107
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v200
	goto L24
L31:
	;
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+4)))
	if v120&int32(_a_F__bt_binsrch_insert_4) == int32(0) {
		v200 = v107
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v126 = int32(_a_F__bt_binsrch_insert_5)
	if v111&v126 == v126 {
		v200 = int32(-1)
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v130 = int32(0)
	v132 = v120 & int32(4095)
	if v132 == v130 {
		v200 = v130
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v139 = int32(0)
	v140 = v132
	goto L35
L35:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+2)))
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114))))
	v154 = int32(16)
	v160 = base.I32_div_s(v140-v139, int32(2))
	v161 = v160 + v139
	v164 = v152 + (v114 + v153<<(uint(v154)%32)) + v161*int32(6)
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151)+2)))
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151))))
	v172 = v168 | v169<<(uint(v154)%32)
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+2)))
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164))))
	v177 = v173 | v174<<(uint(v154)%32)
	if base.Ui32(v172) < base.Ui32(v177) {
		v188 = int32(-1)
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v200 = v195
	goto L30
L37:
	;
	if v195 < v196 {
		v139 = v195
		v140 = v196
		goto L35
	} else {
		goto L47
	}
L38:
	;
	if int32(0) < v188 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	goto L38
L40:
	;
	if base.Ui32(v177) < base.Ui32(v172) {
		v188 = int32(1)
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151)+4)))
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+4)))
	if base.Ui32(v182) < base.Ui32(v183) {
		v188 = int32(-1)
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v188 = base.B2i32(base.Ui32(v183) < base.Ui32(v182))
	goto L39
L43:
	;
	v195 = v161 + int32(1)
	v196 = v140
	goto L37
L44:
	;
	goto L45
L45:
	;
	if int32(0) <= v188 {
		v200 = v161
		goto L30
	} else {
		goto L46
	}
L46:
	;
	v195 = v139
	v196 = v161
	goto L37
L47:
	;
	goto L36
L48:
	;
	v231 = v96
	goto L50
L49:
	;
	v231 = v85
	goto L50
L50:
	;
	v233 = base.B2i32(int32(0) < v99)
	if int32(0) < v99 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v234 = v83
	goto L53
L52:
	;
	v234 = v96
	goto L53
L53:
	;
	if int32(0) < v99 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v239 = v96 + int32(1)
	goto L56
L55:
	;
	v239 = v82
	goto L56
L56:
	;
	if base.Ui32(v239&int32(_a_F__bt_binsrch_insert_0)) < base.Ui32(v234&int32(_a_F__bt_binsrch_insert_0)) {
		v82 = v239
		v83 = v234
		v85 = v231
		goto L22
	} else {
		goto L57
	}
L57:
	;
	goto L23
L58:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L25
	} else {
		goto L59
	}
L59:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v291)+2)))
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v291))))
	v294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v291)+4)))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v295 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v315 + int32(4)
	v320 = int32(_a_F__bt_binsrch_insert_0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v85 & v320
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v82 & v320
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v292 | v293<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F__bt_binsrch_insert_6), v18)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L25
	} else {
		goto L64
	}
L61:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _c_F__bt_binsrch_insert[2]))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v299+(v295^int32(-1))*int32(56))+16))
	v314 = v305
	goto L60
L62:
	;
	goto L63
L63:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _c_F__bt_binsrch_insert[3]))
	v308 = int32(56)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v307+v295*v308-v308)+16))
	v314 = v313
	goto L60
L64:
	;
	F_errfinish(m, int32(_a_F__bt_binsrch_insert_7), int32(571), int32(_a_F__bt_binsrch_insert_8))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L25
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__bt_binsrch_skiparray_skey(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+19)))
	if v9 != 0 {
		v71 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l6))) = v71
		return
	} else {
		if l3 != 0 {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+3)))
			if v12&int32(2) != 0 {
				v15 = int32(-1)
			} else {
				v15 = int32(1)
			}
			v71 = v15
			*(*int32)(unsafe.Add(mBase, uint32(l6))) = v71
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(0)
			if l1 == int32(1) {
				if l0 != 0 {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
					if v32 == int32(0) {
						return
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
						v38 = *(*int64)(unsafe.Add(mBase, uint32(v32)+48))
						v39 = F_FunctionCall2Coll(m, v32+int32(16), v37, l2, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							if v39 != int64(0) {
							} else {
								v71 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l6))) = v71
							}
							return
						}
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
					if v20 == int32(0) {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
						if v32 == int32(0) {
							return
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
							v38 = *(*int64)(unsafe.Add(mBase, uint32(v32)+48))
							v39 = F_FunctionCall2Coll(m, v32+int32(16), v37, l2, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								if v39 != int64(0) {
								} else {
									v71 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(l6))) = v71
								}
								return
							}
						}
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
						v26 = *(*int64)(unsafe.Add(mBase, uint32(v20)+48))
						v27 = F_FunctionCall2Coll(m, v20+int32(16), v25, l2, v26)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							if v27 == int64(0) {
								v71 = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(l6))) = v71
								return
							} else {
								v32 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
								if v32 == int32(0) {
									return
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
									v38 = *(*int64)(unsafe.Add(mBase, uint32(v32)+48))
									v39 = F_FunctionCall2Coll(m, v32+int32(16), v37, l2, v38)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										if v39 != int64(0) {
										} else {
											v71 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(l6))) = v71
										}
										return
									}
								}
							}
						}
					}
				}
			} else {
				if l0 != 0 {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
					if v57 == int32(0) {
						return
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
						v63 = *(*int64)(unsafe.Add(mBase, uint32(v57)+48))
						v64 = F_FunctionCall2Coll(m, v57+int32(16), v62, l2, v63)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							if v64 != int64(0) {
							} else {
								v71 = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(l6))) = v71
							}
							return
						}
					}
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
					if v44 == int32(0) {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
						if v57 == int32(0) {
							return
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
							v63 = *(*int64)(unsafe.Add(mBase, uint32(v57)+48))
							v64 = F_FunctionCall2Coll(m, v57+int32(16), v62, l2, v63)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								if v64 != int64(0) {
								} else {
									v71 = int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(l6))) = v71
								}
								return
							}
						}
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
						v50 = *(*int64)(unsafe.Add(mBase, uint32(v44)+48))
						v51 = F_FunctionCall2Coll(m, v44+int32(16), v49, l2, v50)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							if v51 != int64(0) {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
								if v57 == int32(0) {
									return
								} else {
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
									v63 = *(*int64)(unsafe.Add(mBase, uint32(v57)+48))
									v64 = F_FunctionCall2Coll(m, v57+int32(16), v62, l2, v63)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										if v64 != int64(0) {
										} else {
											v71 = int32(-1)
											*(*int32)(unsafe.Add(mBase, uint32(l6))) = v71
										}
										return
									}
								}
							} else {
								v71 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l6))) = v71
								return
							}
						}
					}
				}
			}
		}
	}
}
func F__bt_check_natts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	v13 = l2 + v12
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)))
	if v14&int32(20) != 0 {
		v128 = int32(1)
		return v128
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
		v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+10)))
		v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+8)))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l2+l3<<(uint(int32(2))%32))+20))
		v26 = l2 + v23&int32(_a_F__bt_check_natts_0)
		v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+6)))
		v29 = v27 & int32(_a_F__bt_check_natts_1)
		if v29 == int32(0) {
			v47 = v19
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v50 != 0 {
				v51 = int32(2)
			} else {
				v51 = int32(1)
			}
			if v14&int32(1) != 0 {
				if base.Ui32(v51) <= base.Ui32(l3) {
					if v29 == int32(0) {
						return base.B2i32(v47 == v19)
					} else {
						v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
						return int32(base.Ui32(v59)>>(uint(int32(13))%32)) & base.B2i32(v47 == v19)
					}
				} else {
					if l1 != 0 {
						v82 = int32(0)
						if v29 == v82 {
							v128 = v82
						} else {
							v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+5)))
							if v85&int32(32) != 0 {
								v128 = v82
							} else {
								v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+6)))
								if v90&int32(_a_F__bt_check_natts_1) == int32(0) {
									v118 = v26
								} else {
									v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
									if v95&int32(_a_F__bt_check_natts_1) == int32(0) {
										v100 = int32(0)
										if v95&int32(_a_F__bt_check_natts_2) == v100 {
											v116 = v100
											v118 = v116
										} else {
											v118 = v26 + v90&int32(_a_F__bt_check_natts_3) - int32(6)
										}
									} else {
										v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+2)))
										v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26))))
										v116 = v110 + (v26 + v111<<(uint(int32(16))%32))
										v118 = v116
									}
								}
								if v47 != v18 {
									v121 = v118
								} else {
									v121 = int32(0)
								}
								if v121 != 0 {
									v128 = v82
								} else {
									v128 = base.B2i32(v47 <= v18) & base.B2i32(int32(0) < v47)
								}
							}
						}
						return v128
					} else {
						return base.B2i32(v47 == v18)
					}
				}
			} else {
				if l3 == v51 {
					v69 = base.B2i32(v47 == int32(0))
					if l1|v69 != 0 {
						v128 = v69 | (l1 ^ int32(1))
						return v128
					} else {
						v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
						return base.B2i32(v76 == int32(1))
					}
				} else {
					if l1 != 0 {
						v82 = int32(0)
						if v29 == v82 {
							v128 = v82
						} else {
							v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+5)))
							if v85&int32(32) != 0 {
								v128 = v82
							} else {
								v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+6)))
								if v90&int32(_a_F__bt_check_natts_1) == int32(0) {
									v118 = v26
								} else {
									v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
									if v95&int32(_a_F__bt_check_natts_1) == int32(0) {
										v100 = int32(0)
										if v95&int32(_a_F__bt_check_natts_2) == v100 {
											v116 = v100
											v118 = v116
										} else {
											v118 = v26 + v90&int32(_a_F__bt_check_natts_3) - int32(6)
										}
									} else {
										v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+2)))
										v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26))))
										v116 = v110 + (v26 + v111<<(uint(int32(16))%32))
										v118 = v116
									}
								}
								if v47 != v18 {
									v121 = v118
								} else {
									v121 = int32(0)
								}
								if v121 != 0 {
									v128 = v82
								} else {
									v128 = base.B2i32(v47 <= v18) & base.B2i32(int32(0) < v47)
								}
							}
						}
						return v128
					} else {
						return base.B2i32(v47 == v18)
					}
				}
			}
		} else {
			v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
			if v32&int32(_a_F__bt_check_natts_1) != 0 {
				v35 = int32(0)
				if base.B2i32(l1 == v35)|v32&int32(_a_F__bt_check_natts_2)|base.B2i32(v18 != v19) != 0 {
					v128 = v35
					return v128
				} else {
					v47 = v19
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					if v50 != 0 {
						v51 = int32(2)
					} else {
						v51 = int32(1)
					}
					if v14&int32(1) != 0 {
						if base.Ui32(v51) <= base.Ui32(l3) {
							if v29 == int32(0) {
								return base.B2i32(v47 == v19)
							} else {
								v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
								return int32(base.Ui32(v59)>>(uint(int32(13))%32)) & base.B2i32(v47 == v19)
							}
						} else {
							if l1 != 0 {
								v82 = int32(0)
								if v29 == v82 {
									v128 = v82
								} else {
									v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+5)))
									if v85&int32(32) != 0 {
										v128 = v82
									} else {
										v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+6)))
										if v90&int32(_a_F__bt_check_natts_1) == int32(0) {
											v118 = v26
										} else {
											v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
											if v95&int32(_a_F__bt_check_natts_1) == int32(0) {
												v100 = int32(0)
												if v95&int32(_a_F__bt_check_natts_2) == v100 {
													v116 = v100
													v118 = v116
												} else {
													v118 = v26 + v90&int32(_a_F__bt_check_natts_3) - int32(6)
												}
											} else {
												v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+2)))
												v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26))))
												v116 = v110 + (v26 + v111<<(uint(int32(16))%32))
												v118 = v116
											}
										}
										if v47 != v18 {
											v121 = v118
										} else {
											v121 = int32(0)
										}
										if v121 != 0 {
											v128 = v82
										} else {
											v128 = base.B2i32(v47 <= v18) & base.B2i32(int32(0) < v47)
										}
									}
								}
								return v128
							} else {
								return base.B2i32(v47 == v18)
							}
						}
					} else {
						if l3 == v51 {
							v69 = base.B2i32(v47 == int32(0))
							if l1|v69 != 0 {
								v128 = v69 | (l1 ^ int32(1))
								return v128
							} else {
								v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
								return base.B2i32(v76 == int32(1))
							}
						} else {
							if l1 != 0 {
								v82 = int32(0)
								if v29 == v82 {
									v128 = v82
								} else {
									v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+5)))
									if v85&int32(32) != 0 {
										v128 = v82
									} else {
										v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+6)))
										if v90&int32(_a_F__bt_check_natts_1) == int32(0) {
											v118 = v26
										} else {
											v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
											if v95&int32(_a_F__bt_check_natts_1) == int32(0) {
												v100 = int32(0)
												if v95&int32(_a_F__bt_check_natts_2) == v100 {
													v116 = v100
													v118 = v116
												} else {
													v118 = v26 + v90&int32(_a_F__bt_check_natts_3) - int32(6)
												}
											} else {
												v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+2)))
												v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26))))
												v116 = v110 + (v26 + v111<<(uint(int32(16))%32))
												v118 = v116
											}
										}
										if v47 != v18 {
											v121 = v118
										} else {
											v121 = int32(0)
										}
										if v121 != 0 {
											v128 = v82
										} else {
											v128 = base.B2i32(v47 <= v18) & base.B2i32(int32(0) < v47)
										}
									}
								}
								return v128
							} else {
								return base.B2i32(v47 == v18)
							}
						}
					}
				}
			} else {
				v47 = v32 & int32(4095)
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
				if v50 != 0 {
					v51 = int32(2)
				} else {
					v51 = int32(1)
				}
				if v14&int32(1) != 0 {
					if base.Ui32(v51) <= base.Ui32(l3) {
						if v29 == int32(0) {
							return base.B2i32(v47 == v19)
						} else {
							v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
							return int32(base.Ui32(v59)>>(uint(int32(13))%32)) & base.B2i32(v47 == v19)
						}
					} else {
						if l1 != 0 {
							v82 = int32(0)
							if v29 == v82 {
								v128 = v82
							} else {
								v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+5)))
								if v85&int32(32) != 0 {
									v128 = v82
								} else {
									v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+6)))
									if v90&int32(_a_F__bt_check_natts_1) == int32(0) {
										v118 = v26
									} else {
										v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
										if v95&int32(_a_F__bt_check_natts_1) == int32(0) {
											v100 = int32(0)
											if v95&int32(_a_F__bt_check_natts_2) == v100 {
												v116 = v100
												v118 = v116
											} else {
												v118 = v26 + v90&int32(_a_F__bt_check_natts_3) - int32(6)
											}
										} else {
											v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+2)))
											v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26))))
											v116 = v110 + (v26 + v111<<(uint(int32(16))%32))
											v118 = v116
										}
									}
									if v47 != v18 {
										v121 = v118
									} else {
										v121 = int32(0)
									}
									if v121 != 0 {
										v128 = v82
									} else {
										v128 = base.B2i32(v47 <= v18) & base.B2i32(int32(0) < v47)
									}
								}
							}
							return v128
						} else {
							return base.B2i32(v47 == v18)
						}
					}
				} else {
					if l3 == v51 {
						v69 = base.B2i32(v47 == int32(0))
						if l1|v69 != 0 {
							v128 = v69 | (l1 ^ int32(1))
							return v128
						} else {
							v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
							return base.B2i32(v76 == int32(1))
						}
					} else {
						if l1 != 0 {
							v82 = int32(0)
							if v29 == v82 {
								v128 = v82
							} else {
								v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+5)))
								if v85&int32(32) != 0 {
									v128 = v82
								} else {
									v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+6)))
									if v90&int32(_a_F__bt_check_natts_1) == int32(0) {
										v118 = v26
									} else {
										v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
										if v95&int32(_a_F__bt_check_natts_1) == int32(0) {
											v100 = int32(0)
											if v95&int32(_a_F__bt_check_natts_2) == v100 {
												v116 = v100
												v118 = v116
											} else {
												v118 = v26 + v90&int32(_a_F__bt_check_natts_3) - int32(6)
											}
										} else {
											v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+2)))
											v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26))))
											v116 = v110 + (v26 + v111<<(uint(int32(16))%32))
											v118 = v116
										}
									}
									if v47 != v18 {
										v121 = v118
									} else {
										v121 = int32(0)
									}
									if v121 != 0 {
										v128 = v82
									} else {
										v128 = base.B2i32(v47 <= v18) & base.B2i32(int32(0) < v47)
									}
								}
							}
							return v128
						} else {
							return base.B2i32(v47 == v18)
						}
					}
				}
			}
		}
	}
}
func F__bt_dedup_start_pending(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	v3 = l2
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
	if v6&int32(32) != 0 {
		v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		if v9&int32(_a_F__bt_dedup_start_pending_0) != 0 {
			v24 = v9 & int32(4095)
			v26 = v24 * int32(6)
			if v26 != 0 {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
				v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
				base.MemoryCopy(m, v27, v28+(l1+v29<<(uint(int32(16))%32)), v26)
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v24
			v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
			v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
			v43 = v36 | v37<<(uint(int32(16))%32)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)) = uint16(v14)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = v16
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(1)
			v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
			v43 = v20 & int32(_a_F__bt_dedup_start_pending_1)
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)) = uint16(v14)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = v16
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(1)
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
		v43 = v20 & int32(_a_F__bt_dedup_start_pending_1)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v43
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = (v49&int32(_a_F__bt_dedup_start_pending_1)+int32(7))&int32(_a_F__bt_dedup_start_pending_2) | int32(4)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v59<<(uint(int32(2))%32))+44)) = uint16(v3)
	return
}
func F__bt_end_parallel(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v82 int32
	_ = v82
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v106 int32
	_ = v106
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v130 int32
	_ = v130
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v142 int32
	_ = v142
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_WaitForParallelWorkersToFinish(m, v4)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
		if int32(0) < v8 {
			v12 = int32(0)
			for {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v17 = v14 + v12<<(uint(int32(7))%32)
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v21 = v18 + v12*int32(40)
				v22 = int32(_a_F__bt_end_parallel_0)
				v24 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[0]))
				v25 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[0])) = v24 + v25
				v28 = int32(_a_F__bt_end_parallel_1)
				v30 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[1]))
				v31 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[1])) = v30 + v31
				v34 = int32(_a_F__bt_end_parallel_2)
				v36 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[2]))
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[2])) = v36 + v37
				v40 = int32(_a_F__bt_end_parallel_3)
				v42 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[3]))
				v43 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[3])) = v42 + v43
				v46 = int32(_a_F__bt_end_parallel_4)
				v48 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[4]))
				v49 = *(*int64)(unsafe.Add(mBase, uint32(v17)+32))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[4])) = v48 + v49
				v52 = int32(_a_F__bt_end_parallel_5)
				v54 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[5]))
				v55 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[5])) = v54 + v55
				v58 = int32(_a_F__bt_end_parallel_6)
				v60 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[6]))
				v61 = *(*int64)(unsafe.Add(mBase, uint32(v17)+48))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[6])) = v60 + v61
				v64 = int32(_a_F__bt_end_parallel_7)
				v66 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[7]))
				v67 = *(*int64)(unsafe.Add(mBase, uint32(v17)+56))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[7])) = v66 + v67
				v70 = int32(_a_F__bt_end_parallel_8)
				v72 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[8]))
				v73 = *(*int64)(unsafe.Add(mBase, uint32(v17)+64))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[8])) = v72 + v73
				v76 = int32(_a_F__bt_end_parallel_9)
				v78 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[9]))
				v79 = *(*int64)(unsafe.Add(mBase, uint32(v17)+72))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[9])) = v78 + v79
				v82 = int32(_a_F__bt_end_parallel_10)
				v84 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[10]))
				v85 = *(*int64)(unsafe.Add(mBase, uint32(v17)+80))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[10])) = v84 + v85
				v88 = int32(_a_F__bt_end_parallel_11)
				v90 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[11]))
				v91 = *(*int64)(unsafe.Add(mBase, uint32(v17)+88))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[11])) = v90 + v91
				v94 = int32(_a_F__bt_end_parallel_12)
				v96 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[12]))
				v97 = *(*int64)(unsafe.Add(mBase, uint32(v17)+96))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[12])) = v96 + v97
				v100 = int32(_a_F__bt_end_parallel_13)
				v102 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[13]))
				v103 = *(*int64)(unsafe.Add(mBase, uint32(v17)+104))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[13])) = v102 + v103
				v106 = int32(_a_F__bt_end_parallel_14)
				v108 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[14]))
				v109 = *(*int64)(unsafe.Add(mBase, uint32(v17)+112))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[14])) = v108 + v109
				v112 = int32(_a_F__bt_end_parallel_15)
				v114 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[15]))
				v115 = *(*int64)(unsafe.Add(mBase, uint32(v17)+120))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[15])) = v114 + v115
				v118 = int32(_a_F__bt_end_parallel_16)
				v120 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[16]))
				v121 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[16])) = v120 + v121
				v124 = int32(_a_F__bt_end_parallel_17)
				v126 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[17]))
				v127 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[17])) = v126 + v127
				v130 = int32(_a_F__bt_end_parallel_18)
				v132 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[18]))
				v133 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[18])) = v132 + v133
				v136 = int32(_a_F__bt_end_parallel_19)
				v138 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[19]))
				v139 = *(*int64)(unsafe.Add(mBase, uint32(v21)+24))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[19])) = v138 + v139
				v142 = int32(_a_F__bt_end_parallel_20)
				v144 = *(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[20]))
				v145 = *(*int64)(unsafe.Add(mBase, uint32(v21)+32))
				*(*int64)(unsafe.Add(mBase, _c_F__bt_end_parallel[20])) = v144 + v145
				v149 = v12 + int32(1)
				v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+20))
				if v149 < v151 {
					v12 = v149
					continue
				} else {
					break
				}
				break
			}
			v155 = v150
		} else {
			v155 = v7
		}
		v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
		if v157 != 0 {
			v161 = v155
			F_DestroyParallelContext(m, v161)
			mBase = m.M
			v163 = m.ExcPending
			if v163 != 0 {
				return
			} else {
				v166 = *(*int32)(unsafe.Add(mBase, _c_F__bt_end_parallel[21]))
				v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+72))
				*(*int32)(unsafe.Add(mBase, uint32(v166)+72)) = v167 - int32(1)
				return
			}
		} else {
			F_UnregisterSnapshot(m, v156)
			mBase = m.M
			v159 = m.ExcPending
			if v159 != 0 {
				return
			} else {
				v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v161 = v160
				F_DestroyParallelContext(m, v161)
				mBase = m.M
				v163 = m.ExcPending
				if v163 != 0 {
					return
				} else {
					v166 = *(*int32)(unsafe.Add(mBase, _c_F__bt_end_parallel[21]))
					v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+72))
					*(*int32)(unsafe.Add(mBase, uint32(v166)+72)) = v167 - int32(1)
					return
				}
			}
		}
	}
}
func F__bt_metaversion(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v6 == int32(0) {
		v10 = F_ReadBuffer(m, l0, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			F_LockBufferInternal(m, v10, int32(1))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F__bt_checkpage(m, l0, v10)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					v17 = F__bt_getmeta(m, l0, v10)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
						if v19 == int32(0) {
							v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
							*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v22)))
							v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+40)))
							*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v26)
							F_UnlockReleaseBuffer(m, v10)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return
							} else {
								return
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							v32 = F_MemoryContextAlloc(m, v30, int32(48))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v32
								v35 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v32)+40)) = v35
								v37 = *(*int64)(unsafe.Add(mBase, uint32(v17)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v32)+32)) = v37
								v39 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v32)+24)) = v39
								v41 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = v41
								v43 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v43
								v45 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
								*(*int64)(unsafe.Add(mBase, uint32(v32))) = v45
								F_UnlockReleaseBuffer(m, v10)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
									v51 = v49
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
									*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v53)))
									v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+40)))
									*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v57)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		v51 = v6
		v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v53)))
		v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+40)))
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v57)
		return
	}
}
func F__bt_moveright(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+3)))
	v24 = l3
	goto L1
L1:
	;
	if v24 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
	if v117&int32(20) != 0 {
		goto L35
	} else {
		goto L36
	}
L3:
	;
	goto L2
L4:
	;
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+16)))
	v53 = v52 + v51
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v54 == int32(0) {
		goto L3
	} else {
		goto L8
	}
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F__bt_moveright[0]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+(v24^int32(-1))<<(uint(int32(2))%32))))
	v51 = v43
	goto L4
L6:
	;
	goto L7
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F__bt_moveright[1]))
	v51 = v45 + v24<<(uint(int32(13))%32) + int32(-8192)
	goto L4
L8:
	;
	v57 = int32(0)
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
	if base.B2i32(l4 == v57)|base.B2i32(v59&int32(128) == v57) == v57 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if v24 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	if v59&int32(20) != 0 {
		goto L29
	} else {
		goto L30
	}
L12:
	;
	if l6 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F__bt_moveright[2]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70+(v24^int32(-1))*int32(56))+16))
	v85 = v76
	goto L12
L14:
	;
	goto L15
L15:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F__bt_moveright[3]))
	v79 = int32(56)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v78+v24*v79-v79)+16))
	v85 = v84
	goto L12
L16:
	;
	F_UnlockBuffer(m, v24)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
	if v95&int32(128) != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	return int32(0)
L20:
	;
	F__bt_lockbuf(m, v24, int32(3))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	F__bt_finish_split(m, l0, l1, v24, l5)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L19
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_UnlockReleaseBuffer(m, v24)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L19
	} else {
		goto L27
	}
L25:
	;
	v100 = F__bt_getbuf(m, l0, v85, l6)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	v24 = v100
	goto L1
L27:
	;
	v104 = F__bt_getbuf(m, l0, v85, l6)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	v24 = v104
	goto L1
L29:
	;
	v113 = v54
	goto L31
L30:
	;
	v109 = F__bt_compare(m, l0, l2, v51, int32(1))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L19
	} else {
		goto L32
	}
L31:
	;
	v114 = F__bt_relandgetbuf(m, l0, v24, v113, l6)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L19
	} else {
		goto L34
	}
L32:
	;
	if v109 < v18^int32(1) {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v113 = v112
	goto L31
L34:
	;
	v24 = v114
	goto L1
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L19
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	m.G0 = v16 + int32(16)
	return v24
L38:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v124 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_moveright_0), v16)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L19
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F__bt_moveright_1), int32(319), int32(_a_F__bt_moveright_2))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L19
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__bt_parallel_release(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+24))
	v7 = v5 + v6
	v9 = v7 + int32(12)
	v11 = F_LWLockAcquire(m, v9, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(3)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
		F_LWLockRelease(m, v9)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			F_ConditionVariableSignal(m, v7+int32(28))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F__bt_readpage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
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
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v415 int32
	_ = v415
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v463 int32
	_ = v463
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v539 int64
	_ = v539
	var v541 int32
	_ = v541
	var v542 int64
	_ = v542
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v599 int32
	_ = v599
	var v622 int32
	_ = v622
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v673 int32
	_ = v673
	var v680 int32
	_ = v680
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v902 int32
	_ = v902
	var v908 int32
	_ = v908
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v954 int32
	_ = v954
	var v966 int32
	_ = v966
	var v978 int32
	_ = v978
	var v1001 int32
	_ = v1001
	var v1013 int32
	_ = v1013
	var v1026 int32
	_ = v1026
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1048 int32
	_ = v1048
	var v1054 int32
	_ = v1054
	var v1066 int32
	_ = v1066
	var v1077 int32
	_ = v1077
	v4 = l3
	v22 = m.G0
	v24 = v22 - int32(48)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+56))
	if v29 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+16)))
	if v29 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readpage[0]))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33+(v29^int32(-1))<<(uint(int32(2))%32))))
	v47 = v39
	goto L1
L3:
	;
	goto L4
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readpage[1]))
	v47 = v41 + v29<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+60)) = v67
	v69 = v48 + v47
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+68)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = l1
	v76 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v26)+192))
	v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80)+8)))
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+46)) = uint16(v76)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+42)) = v76
	*(*int64)(unsafe.Add(mBase, uint32(v24)+32)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+29)) = uint8(v76)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+28)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v76
	if base.Ui32(int32(25)) <= base.Ui32(v82) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readpage[2]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52+(v29^int32(-1))*int32(56))+16))
	v67 = v58
	goto L5
L7:
	;
	goto L8
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readpage[3]))
	v61 = int32(56)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60+v29*v61-v61)+16))
	v67 = v66
	goto L5
L9:
	;
	v102 = int32(base.Ui32(v82+int32(_a_F__bt_readpage_0)) >> (uint(int32(2)) % 32))
	goto L11
L10:
	;
	v102 = v76
	goto L11
L11:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+18)) = uint16(v102)
	v104 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+40)) = uint8(v104)
	if v78 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v108 = int32(2)
	goto L14
L13:
	;
	v108 = v104
	goto L14
L14:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+16)) = uint16(v108)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v110 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if l1 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_PredicateLockPage(m, v26, v118, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L21
	} else {
		goto L23
	}
L18:
	;
	v113 = v72
	goto L20
L19:
	;
	v113 = v70
	goto L20
L20:
	;
	F__bt_parallel_release(m, l0, v113, v67)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	goto L17
L23:
	;
	if l1 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	m.G0 = v24 + int32(48)
	return v1077
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+100)) = v1066
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v1048
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = v1054
	v1077 = base.B2i32(v1054 <= v1048)
	goto L24
L26:
	;
	if v79 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v79 != 0 {
		goto L118
	} else {
		goto L119
	}
L29:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v124 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	if v4|base.B2i32(base.Ui32(v102&int32(_a_F__bt_readpage_1)) <= base.Ui32(v108)) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L32:
	;
	v150 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+18)) = uint16(v150)
	goto L31
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	v130 = v47 + v127&int32(_a_F__bt_readpage_2)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v130
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+18)))
	if v132 != int32(1) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v136 = F__bt_scanbehind_checkkeys(m, l0, int32(1), v130)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L21
	} else {
		goto L35
	}
L35:
	;
	if v136 != 0 {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v138 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+17)) = uint8(v138)
	v140 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+89)) = uint8(v140)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v143 == v140 {
		v1077 = v140
		goto L24
	} else {
		goto L37
	}
L37:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	F__bt_parallel_primscan_schedule(m, l0, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L21
	} else {
		goto L38
	}
L38:
	;
	v1077 = v140
	goto L24
L39:
	;
	F__bt_set_startikey(m, l0, v24+int32(12))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L21
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(v108) < base.Ui32(l2) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	v165 = l2
	goto L45
L44:
	;
	v165 = v108
	goto L45
L45:
	;
	v167 = v102 & int32(_a_F__bt_readpage_1)
	if base.Ui32(v165) <= base.Ui32(v167) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v170 = v28 + int32(104)
	v179 = int32(0)
	v182 = v165
	goto L49
L47:
	;
	v463 = int32(-1)
	goto L48
L48:
	;
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+40)))
	if v481 != int32(1) {
		goto L88
	} else {
		goto L89
	}
L49:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(20)+v182&int32(_a_F__bt_readpage_1)<<(uint(int32(2))%32))))
	v203 = int32(_a_F__bt_readpage_3)
	if v27&int32(1)&base.B2i32(v202&v203 == v203) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v463 = v440 - int32(1)
	goto L48
L51:
	;
	goto L50
L52:
	;
	if base.Ui32(v433&int32(_a_F__bt_readpage_1)) <= base.Ui32(v167) {
		v179 = v415
		v182 = v433
		goto L49
	} else {
		goto L87
	}
L53:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+36)) = uint16(v182)
	v217 = v47 + v202&int32(_a_F__bt_readpage_2)
	v218 = F__bt_checkkeys(m, l0, v24+int32(12), base.B2i32(v79 != int32(0)), v217, v81)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L21
	} else {
		goto L56
	}
L54:
	;
	v392 = v179
	goto L55
L55:
	;
	v415 = v392
	v433 = v182 + int32(1)
	goto L52
L56:
	;
	if v79 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if v218 == int32(0) {
		v368 = v179
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+38)))
	if base.Ui32(int32(2047)) < base.Ui32((v222-int32(1))&int32(_a_F__bt_readpage_1)) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v229 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+38)) = uint16(v229)
	v415 = v179
	v433 = v222
	goto L52
L60:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+40)))
	if v386 != int32(1) {
		v440 = v368
		goto L51
	} else {
		goto L86
	}
L61:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+7)))
	if v234&int32(32) != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v269 = v170 + v179*int32(10)
	v270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+2)))
	v271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217))))
	v275 = v270 + (v217 + v271<<(uint(int32(16))%32))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	*(*int32)(unsafe.Add(mBase, uint32(v269))) = v276
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v275)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v269)+6)) = uint16(v182)
	*(*uint16)(unsafe.Add(mBase, uint32(v269)+4)) = uint16(v278)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v28)+44))
	if v281 != 0 {
		goto L73
	} else {
		goto L74
	}
L63:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+5)))
	if v237&int32(32) != 0 {
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v242 = v170 + v179*int32(10)
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v242)+4)) = uint16(v243)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v245
	*(*uint16)(unsafe.Add(mBase, uint32(v242)+6)) = uint16(v182)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v28)+44))
	if v248 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L65
L67:
	;
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+6)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v28)+84))
	*(*uint16)(unsafe.Add(mBase, uint32(v242)+8)) = uint16(v250)
	v253 = v249 & int32(_a_F__bt_readpage_4)
	if v253 != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	v368 = v179 + int32(1)
	goto L60
L70:
	;
	base.MemoryCopy(m, v248+v250, v217, v253)
	goto L72
L71:
	;
	goto L72
L72:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v28)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v256 + (v253+int32(7))&int32(_a_F__bt_readpage_5)
	goto L69
L73:
	;
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+2)))
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217))))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v28)+84))
	*(*uint16)(unsafe.Add(mBase, uint32(v269)+8)) = uint16(v284)
	v286 = v281 + v284
	v293 = (v282 | v283<<(uint(int32(16))%32) + int32(7)) & int32(-8)
	if v293 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v308 = int32(0)
	goto L75
L75:
	;
	v309 = int32(1)
	v311 = v179 + v309
	v312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)))
	if v312&int32(4094) == int32(0) {
		v368 = v311
		goto L60
	} else {
		goto L79
	}
L76:
	;
	base.MemoryCopy(m, v286, v217, v293)
	goto L78
L77:
	;
	goto L78
L78:
	;
	v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v286)+6)))
	v298 = v295&int32(_a_F__bt_readpage_6) | v293
	*(*uint16)(unsafe.Add(mBase, uint32(v286)+6)) = uint16(v298)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v28)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v300 + v293
	v303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v269)+8)))
	v308 = v303
	goto L75
L79:
	;
	v319 = v309
	v320 = v311
	goto L80
L80:
	;
	v340 = v170 + v320*int32(10)
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+2)))
	v342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217))))
	v349 = v341 + (v217 + v342<<(uint(int32(16))%32)) + v319*int32(6)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	*(*int32)(unsafe.Add(mBase, uint32(v340))) = v350
	v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v349)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v340)+6)) = uint16(v182)
	*(*uint16)(unsafe.Add(mBase, uint32(v340)+4)) = uint16(v352)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v28)+44))
	if v355 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v368 = v358
	goto L60
L82:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v340)+8)) = uint16(v308)
	goto L84
L83:
	;
	goto L84
L84:
	;
	v357 = int32(1)
	v358 = v320 + v357
	v360 = v319 + v357
	v361 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)))
	if base.Ui32(v360) < base.Ui32(v361&int32(4095)) {
		v319 = v360
		v320 = v358
		goto L80
	} else {
		goto L85
	}
L85:
	;
	goto L81
L86:
	;
	v392 = v368
	goto L55
L87:
	;
	v440 = v415
	goto L51
L88:
	;
	v673 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+89)) = uint8(v673)
	v1048 = v463
	v1054 = v673
	v1066 = v673
	goto L25
L89:
	;
	v484 = int32(0)
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+18)))
	if v486 != 0 {
		v1048 = v463
		v1054 = v484
		v1066 = v484
		goto L25
	} else {
		goto L90
	}
L90:
	;
	v487 = int32(0)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if v488 == v487 {
		v1048 = v463
		v1054 = v484
		v1066 = v487
		goto L25
	} else {
		goto L91
	}
L91:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	v494 = v47 + v491&int32(_a_F__bt_readpage_2)
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+29)))
	if v495 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v498)+12))
	if int32(0) < v499 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	v622 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v622
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+29)) = uint8(v622)
	v631 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+7)))
	if v631&int32(32) == v622 {
		goto L113
	} else {
		goto L114
	}
L95:
	;
	v505 = int32(0)
	goto L98
L96:
	;
	goto L97
L97:
	;
	v599 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v498)+18)) = uint16(v599)
	goto L94
L98:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v498)+8))
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v498)+20))
	v528 = v525 + v505<<(uint(int32(5))%32)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v528)))
	v532 = v524 + v529*int32(56)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v528)+4))
	if v533 != int32(-1) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L97
L100:
	;
	v575 = v505 + int32(1)
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v498)+12))
	if v575 < v576 {
		v505 = v575
		goto L98
	} else {
		goto L111
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v528)+12)) = int32(0)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v528)+8))
	v539 = *(*int64)(unsafe.Add(mBase, uint32(v538)))
	*(*int64)(unsafe.Add(mBase, uint32(v532)+48)) = v539
	goto L100
L102:
	;
	goto L103
L103:
	;
	v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+18)))
	if v541 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v532)+48)) = int64(0)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	v553 = v551 & int32(-7864386)
	*(*int32)(unsafe.Add(mBase, uint32(v532))) = v553
	v557 = int32(0)
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v528)+19)))
	if base.B2i32(v551&int32(33554432) == v557)|base.B2i32(v559 != int32(1)) == v557 {
		goto L108
	} else {
		goto L109
	}
L105:
	;
	v542 = *(*int64)(unsafe.Add(mBase, uint32(v532)+48))
	if v542 == int64(0) {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	F_pfree(m, base.I32_wrap_i64(v542))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L21
	} else {
		goto L107
	}
L107:
	;
	goto L104
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v532))) = v553 | int32(65)
	goto L100
L109:
	;
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v532))) = v553 | int32(_a_F__bt_readpage_7)
	goto L100
L111:
	;
	goto L99
L112:
	;
	v646 = F__bt_checkkeys(m, l0, v24+int32(12), base.B2i32(v79 != v622), v494, v645)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L21
	} else {
		goto L116
	}
L113:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v26)+192))
	v643 = int32(*(*int16)(unsafe.Add(mBase, uint32(v642)+8)))
	v645 = v643
	goto L112
L114:
	;
	v636 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v494)+4)))
	if v636&int32(_a_F__bt_readpage_8) != 0 {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v645 = v636 & int32(4095)
	goto L112
L116:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+40)))
	if v648 == int32(0) {
		goto L88
	} else {
		goto L117
	}
L117:
	;
	v1048 = v463
	v1054 = v622
	v1066 = int32(0)
	goto L25
L118:
	;
	if base.Ui32(v102&int32(_a_F__bt_readpage_1)) < base.Ui32(v108) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	if v4|base.B2i32(base.Ui32(v102&int32(_a_F__bt_readpage_1)) <= base.Ui32(v108)) == int32(0) {
		goto L129
	} else {
		goto L130
	}
L121:
	;
	v708 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+18)) = uint16(v708)
	goto L120
L122:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v680 == int32(0) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v47+v108<<(uint(int32(2))%32))+20))
	v689 = v47 + v686&int32(_a_F__bt_readpage_2)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v689
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+18)))
	if v691 != int32(1) {
		goto L121
	} else {
		goto L124
	}
L124:
	;
	v694 = F__bt_scanbehind_checkkeys(m, l0, l1, v689)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L21
	} else {
		goto L125
	}
L125:
	;
	if v694 != 0 {
		goto L121
	} else {
		goto L126
	}
L126:
	;
	v696 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+17)) = uint8(v696)
	v698 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+88)) = uint8(v698)
	v701 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v701 == v698 {
		v1077 = v698
		goto L24
	} else {
		goto L127
	}
L127:
	;
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	F__bt_parallel_primscan_schedule(m, l0, v704)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L21
	} else {
		goto L128
	}
L128:
	;
	v1077 = v698
	goto L24
L129:
	;
	F__bt_set_startikey(m, l0, v24+int32(12))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L21
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v721 = int32(1358)
	v723 = v102 & int32(_a_F__bt_readpage_1)
	if base.Ui32(l2) < base.Ui32(v723) {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	goto L131
L133:
	;
	v1039 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+40)))
	if v1039 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L134:
	;
	v725 = l2
	goto L136
L135:
	;
	v725 = v723
	goto L136
L136:
	;
	if base.Ui32(v725) < base.Ui32(v108) {
		v1026 = v721
		goto L133
	} else {
		goto L137
	}
L137:
	;
	v728 = v28 + int32(104)
	v737 = v725
	v740 = v721
	goto L138
L138:
	;
	v753 = v737 & int32(_a_F__bt_readpage_1)
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(20)+v753<<(uint(int32(2))%32))))
	v758 = int32(_a_F__bt_readpage_3)
	v762 = v27 & base.B2i32(v757&v758 == v758)
	if v762&base.B2i32(base.Ui32(v108) < base.Ui32(v753)) == int32(0) {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	v1026 = v1001
	goto L133
L140:
	;
	if base.Ui32(v108) <= base.Ui32(v1013&int32(_a_F__bt_readpage_1)) {
		v737 = v1013
		v740 = v1001
		goto L138
	} else {
		goto L181
	}
L141:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+36)) = uint16(v737)
	v770 = v47 + v757&int32(_a_F__bt_readpage_2)
	if v79 != 0 {
		goto L145
	} else {
		goto L146
	}
L142:
	;
	v978 = v740
	goto L143
L143:
	;
	v1001 = v978
	v1013 = v737 - int32(1)
	goto L140
L144:
	;
	v805 = int32(0)
	if v803&base.B2i32(v762 == v805) == v805 {
		v954 = v740
		goto L156
	} else {
		goto L157
	}
L145:
	;
	if v753 != v108 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L147
L147:
	;
	v801 = F__bt_checkkeys(m, l0, v24+int32(12), int32(0), v770, v81)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L21
	} else {
		goto L155
	}
L148:
	;
	v786 = F__bt_checkkeys(m, l0, v24+int32(12), int32(1), v770, v81)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L21
	} else {
		goto L152
	}
L149:
	;
	v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+29)))
	if v772&int32(1) == int32(0) {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v777 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v777
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+29)) = uint8(v777)
	F__bt_start_array_keys(m, l0, l1)
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L21
	} else {
		goto L151
	}
L151:
	;
	goto L148
L152:
	;
	v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+18)))
	if v788 != 0 {
		v1026 = v740
		goto L133
	} else {
		goto L153
	}
L153:
	;
	v789 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+38)))
	if base.Ui32(int32(2047)) < base.Ui32((v789-int32(1))&int32(_a_F__bt_readpage_1)) {
		v803 = v786
		goto L144
	} else {
		goto L154
	}
L154:
	;
	v796 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+38)) = uint16(v796)
	v1001 = v740
	v1013 = v789
	goto L140
L155:
	;
	v803 = v801
	goto L144
L156:
	;
	v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+40)))
	if v966 != int32(1) {
		v1026 = v954
		goto L133
	} else {
		goto L180
	}
L157:
	;
	v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v770)+7)))
	if v810&int32(32) != 0 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v844 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v770)+2)))
	v845 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v770))))
	v851 = v813 & int32(4095)
	v852 = int32(6)
	v856 = v844 + (v770 + v845<<(uint(int32(16))%32)) + v851*v852 - v852
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v856)))
	v858 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v856)+4)))
	v860 = v740 - int32(1)
	v863 = v728 + v860*int32(10)
	*(*uint16)(unsafe.Add(mBase, uint32(v863)+6)) = uint16(v737)
	*(*uint16)(unsafe.Add(mBase, uint32(v863)+4)) = uint16(v858)
	*(*int32)(unsafe.Add(mBase, uint32(v863))) = v857
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v28)+44))
	if v867 != 0 {
		goto L167
	} else {
		goto L168
	}
L159:
	;
	v813 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v770)+4)))
	if v813&int32(_a_F__bt_readpage_8) != 0 {
		goto L158
	} else {
		goto L162
	}
L160:
	;
	goto L161
L161:
	;
	v818 = v740 - int32(1)
	v821 = v728 + v818*int32(10)
	v822 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v770)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v821)+4)) = uint16(v822)
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v770)))
	*(*int32)(unsafe.Add(mBase, uint32(v821))) = v824
	*(*uint16)(unsafe.Add(mBase, uint32(v821)+6)) = uint16(v737)
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v28)+44))
	if v827 == int32(0) {
		v954 = v818
		goto L156
	} else {
		goto L163
	}
L162:
	;
	goto L161
L163:
	;
	v830 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v770)+6)))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v28)+84))
	*(*uint16)(unsafe.Add(mBase, uint32(v821)+8)) = uint16(v831)
	v834 = v830 & int32(_a_F__bt_readpage_4)
	if v834 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	base.MemoryCopy(m, v827+v831, v770, v834)
	goto L166
L165:
	;
	goto L166
L166:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v28)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v837 + (v834+int32(7))&int32(_a_F__bt_readpage_5)
	v954 = v818
	goto L156
L167:
	;
	v868 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v770)+2)))
	v869 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v770))))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v28)+84))
	*(*uint16)(unsafe.Add(mBase, uint32(v863)+8)) = uint16(v870)
	v872 = v867 + v870
	v879 = (v868 | v869<<(uint(int32(16))%32) + int32(7)) & int32(-8)
	if v879 != 0 {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	v894 = int32(0)
	goto L169
L169:
	;
	if base.Ui32(v851) < base.Ui32(int32(2)) {
		v954 = v860
		goto L156
	} else {
		goto L173
	}
L170:
	;
	base.MemoryCopy(m, v872, v770, v879)
	goto L172
L171:
	;
	goto L172
L172:
	;
	v881 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v872)+6)))
	v884 = v881&int32(_a_F__bt_readpage_6) | v879
	*(*uint16)(unsafe.Add(mBase, uint32(v872)+6)) = uint16(v884)
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v28)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v886 + v879
	v889 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v863)+8)))
	v894 = v889
	goto L169
L173:
	;
	v902 = v851 - int32(2)
	v908 = v860
	goto L174
L174:
	;
	v921 = v908 - int32(1)
	v924 = v728 + v921*int32(10)
	v925 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v770)+2)))
	v926 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v770))))
	v933 = v925 + (v770 + v926<<(uint(int32(16))%32)) + v902*int32(6)
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v933)))
	*(*int32)(unsafe.Add(mBase, uint32(v924))) = v934
	v936 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v933)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v924)+6)) = uint16(v737)
	*(*uint16)(unsafe.Add(mBase, uint32(v924)+4)) = uint16(v936)
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v28)+44))
	if v939 != 0 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v954 = v921
	goto L156
L176:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v924)+8)) = uint16(v894)
	goto L178
L177:
	;
	goto L178
L178:
	;
	if int32(0) < v902 {
		v902 = v902 - int32(1)
		v908 = v921
		goto L174
	} else {
		goto L179
	}
L179:
	;
	goto L175
L180:
	;
	v978 = v954
	goto L143
L181:
	;
	goto L139
L182:
	;
	v1042 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+88)) = uint8(v1042)
	goto L184
L183:
	;
	goto L184
L184:
	;
	v1048 = int32(1357)
	v1054 = v1026
	v1066 = int32(1357)
	goto L25
}
func F_bt_index_parent_check(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v21 int64
	_ = v21
	var v27 int64
	_ = v27
	var v40 int32
	_ = v40
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(1)
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v12 < int32(2) {
	} else {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(base.B2i32(v15 != int64(0)))
		if v12 == int32(2) {
		} else {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(base.B2i32(v21 != int64(0)))
			if base.Ui32(v12) < base.Ui32(int32(4)) {
			} else {
				v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(base.B2i32(v27 != int64(0)))
			}
		}
	}
	F_amcheck_lock_relation_and_check(m, base.I32_wrap_i64(v9), int32(403), int32(_a_F_bt_index_parent_check_0), int32(5), v7+int32(12))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		return int64(0)
	} else {
		m.G0 = v7 + int32(16)
		return int64(0)
	}
}
func F_bt_normalize_tuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v129 int64
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int64
	_ = v134
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v276 int32
	_ = v276
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	v3 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(352)
	m.G0 = v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
	if v20&int32(64) == v3 {
		v276 = l1
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L22
	} else {
		goto L63
	}
L2:
	;
	m.G0 = v18 + int32(352)
	return v276
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v27 <= int32(0) {
		v276 = l1
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v31 = l1 + int32(8)
	v36 = v3
	v37 = v27
	v40 = v3
	goto L5
L5:
	;
	v51 = v18 + int32(32) + v40
	v52 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v51))) = uint8(v52)
	v59 = v26 + v37<<(uint(int32(3))%32) + v40*int32(100)
	v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59)+102)))
	v63 = v18 - int32(-64) + v40
	*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v52)
	v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v52 <= v66 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	if v199&int32(1) == int32(0) {
		v276 = l1
		goto L2
	} else {
		goto L51
	}
L7:
	;
	v139 = v18 + int32(96) + v40<<(uint(int32(3))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v139))) = v134
	v142 = v59 + int32(28)
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+82)))
	if v143 != 0 {
		goto L29
	} else {
		goto L30
	}
L8:
	;
	v129 = F_nocache_index_getattr(m, l1, v60, v26)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L22
	} else {
		goto L27
	}
L9:
	;
	v71 = v26 + int32(20) + v60<<(uint(int32(3))%32)
	v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(v71))))
	if v72 < int32(0) {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v111 = int32(1)
	v112 = v60 - v111
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v112>>(uint(int32(3))%32)))))
	if int32(base.Ui32(v116)>>(uint(v112&int32(7))%32))&v111 != 0 {
		goto L8
	} else {
		goto L26
	}
L12:
	;
	v75 = v72 + v31
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+4)))
	if v76 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(v71)+2)))
	if base.I32_popcnt(v79) != int32(1) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v131 = int32(0)
	v134 = base.I64_extend_i32_u(v75)
	goto L7
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	switch base.I32_ctz(v79) {
	case 0:
		goto L21
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		goto L18
	default:
		goto L16
	}
L18:
	;
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
	v131 = int32(0)
	v134 = v91
	goto L7
L19:
	;
	v89 = int64(*(*int32)(unsafe.Add(mBase, uint32(v75))))
	v131 = int32(0)
	v134 = v89
	goto L7
L20:
	;
	v87 = int64(*(*int16)(unsafe.Add(mBase, uint32(v75))))
	v131 = int32(0)
	v134 = v87
	goto L7
L21:
	;
	v85 = int64(*(*int8)(unsafe.Add(mBase, uint32(v75))))
	v131 = int32(0)
	v134 = v85
	goto L7
L22:
	;
	return int32(0)
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v79
	F_errmsg_internal(m, int32(_a_F_bt_normalize_tuple_0), v18+int32(16))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_bt_normalize_tuple_1), int32(123), int32(_a_F_bt_normalize_tuple_2))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	v122 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v122)
	v131 = v122
	v134 = int64(0)
	goto L7
L27:
	;
	v131 = int32(0)
	v134 = v129
	goto L7
L28:
	;
	v205 = v40 + int32(1)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v205 < v206 {
		v36 = v199
		v37 = v206
		v40 = v205
		goto L5
	} else {
		goto L50
	}
L29:
	;
	v199 = v36
	goto L28
L30:
	;
	goto L31
L31:
	;
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+72)))
	if base.B2i32(v144 != int32(_a_F_bt_normalize_tuple_3))|v131 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v199 = v36
	goto L28
L33:
	;
	goto L34
L34:
	;
	v148 = base.I32_wrap_i64(v134)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v149 == int32(1) {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v153 = v149 & int32(3)
	if v153 != int32(2) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v139))) = base.I64_extend_i32_u(v191)
	v196 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v51))) = uint8(v196)
	v199 = v196
	goto L28
L37:
	;
	if v153 != 0 {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	if v149&int32(1) != 0 {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v165 = F_pg_detoast_datum(m, v148)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L22
	} else {
		goto L43
	}
L41:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if base.Ui32(v158) < base.Ui32(int32(2044)) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+84)))
	switch v162 - int32(109) {
	case 0, 11:
		v199 = int32(1)
		goto L28
	default:
		goto L37
	}
L43:
	;
	v191 = v165
	goto L36
L44:
	;
	v199 = v36
	goto L28
L45:
	;
	goto L46
L46:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v170 = int32(base.Ui32(v168) >> (uint(int32(2)) % 32))
	v172 = v170 - int32(3)
	if base.Ui32(int32(127)) < base.Ui32(v172) {
		v199 = v36
		goto L28
	} else {
		goto L47
	}
L47:
	;
	v175 = F_palloc(m, v172)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L22
	} else {
		goto L48
	}
L48:
	;
	v177 = int32(1)
	v180 = v172<<(uint(v177)%32) | v177
	*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v180)
	v183 = v170 - int32(4)
	if v183 == int32(0) {
		v191 = v175
		goto L36
	} else {
		goto L49
	}
L49:
	;
	base.MemoryCopy(m, v175+int32(1), v148+int32(4), v183)
	v191 = v175
	goto L36
L50:
	;
	goto L6
L51:
	;
	v216 = F_index_form_tuple(m, v26, v18+int32(96), v18-int32(-64))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L22
	} else {
		goto L52
	}
L52:
	;
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v216)+4)) = uint16(v218)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v216))) = v220
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if int32(0) < v222 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v228 = v222
	v232 = int32(0)
	goto L56
L54:
	;
	goto L55
L55:
	;
	v276 = v216
	goto L2
L56:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(32)+v232))))
	if v244 == int32(1) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L55
L58:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(96)+v232<<(uint(int32(3))%32))))
	F_pfree(m, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L22
	} else {
		goto L61
	}
L59:
	;
	v256 = v228
	goto L60
L60:
	;
	v258 = v232 + int32(1)
	if v258 < v256 {
		v228 = v256
		v232 = v258
		goto L56
	} else {
		goto L62
	}
L61:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v256 = v255
	goto L60
L62:
	;
	goto L57
L63:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L22
	} else {
		goto L64
	}
L64:
	;
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)+48))
	v305 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v305
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v304 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v301 | v302<<(uint(int32(16))%32)
	F_errmsg(m, int32(_a_F_bt_normalize_tuple_4), v18)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L22
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_bt_normalize_tuple_5), int32(2892), int32(_a_F_bt_normalize_tuple_6))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L22
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_bt_page_items_1_9(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_bt_page_items_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_bt_page_items_bytea(m *base.Module, l0 int32) int64 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v26 int32
	_ = v26
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v180 int64
	_ = v180
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = F_superuser(m)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			if v18 != 0 {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
				if v21 == int32(0) {
					v24 = F_init_MultiFuncCall(m, l0)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int64(0)
					} else {
						v26 = int32(_a_F_bt_page_items_bytea_0)
						v27 = *(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0]))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
						*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v29
						v32 = F_palloc(m, int32(12))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							v34 = F_get_page_from_raw(m, v14)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v32))) = v34
								v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+14)))
								if v37 == int32(0) {
									*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
									v172 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
									v180 = int64(0)
									m.G0 = v11 + int32(32)
									return v180
								} else {
									v42 = int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v32)+4)) = uint16(v42)
									v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+19)))
									v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+16)))
									if (v44<<(uint(int32(8))%32)-v47)&int32(_a_F_bt_page_items_bytea_1) != int32(16) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v204 = m.ExcPending
										if v204 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v207 = m.ExcPending
											if v207 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_bt_page_items_bytea_2)
												F_errmsg(m, int32(_a_F_bt_page_items_bytea_3), v11+int32(16))
												mBase = m.M
												v214 = m.ExcPending
												if v214 != 0 {
													return int64(0)
												} else {
													v215 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
													v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215)+16)))
													v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+19)))
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(16)
													*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = (v217<<(uint(int32(8))%32) - v216) & int32(_a_F_bt_page_items_bytea_1)
													v227 = F_errdetail(m, int32(_a_F_bt_page_items_bytea_4), v11)
													mBase = m.M
													v228 = m.ExcPending
													if v228 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(774), int32(_a_F_bt_page_items_bytea_6))
														mBase = m.M
														v233 = m.ExcPending
														if v233 != 0 {
															return int64(0)
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
										v53 = v34 + v47
										v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
										if v54&int32(8) != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v237 = m.ExcPending
											if v237 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v240 = m.ExcPending
												if v240 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_bt_page_items_bytea_7), int32(0))
													mBase = m.M
													v244 = m.ExcPending
													if v244 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(781), int32(_a_F_bt_page_items_bytea_6))
														mBase = m.M
														v249 = m.ExcPending
														if v249 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											if v54&int32(1) != 0 {
												v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
												if v59 != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v253 = m.ExcPending
													if v253 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v256 = m.ExcPending
														if v256 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_bt_page_items_bytea_8), int32(0))
															mBase = m.M
															v260 = m.ExcPending
															if v260 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(786), int32(_a_F_bt_page_items_bytea_6))
																mBase = m.M
																v265 = m.ExcPending
																if v265 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													}
												} else {
													if v54&int32(4) == int32(0) {
														v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
														if v79&int32(4) == int32(0) {
															v84 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
															v85 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v84)+12)))
															if base.Ui64(int64(25)) <= base.Ui64(v85) {
																v95 = int64(base.Ui64(v85+int64(262120))>>(uint(int64(2))%64)) & int64(65535)
															} else {
																v95 = int64(0)
															}
															*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v95
															v113 = v79
															v116 = v113 & int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
															v119 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
															v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return int64(0)
															} else {
																if v125 != int32(1) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v269 = m.ExcPending
																	if v269 != 0 {
																		return int64(0)
																	} else {
																		F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																		mBase = m.M
																		v273 = m.ExcPending
																		if v273 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																			mBase = m.M
																			v278 = m.ExcPending
																			if v278 != 0 {
																				return int64(0)
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		}
																	}
																} else {
																	v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																	v130 = F_BlessTupleDesc(m, v129)
																	mBase = m.M
																	v131 = m.ExcPending
																	if v131 != 0 {
																		return int64(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																		*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																		*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																		v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																		v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																		v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																		v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																		if base.Ui64(v144) < base.Ui64(v145) {
																			v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																			v148 = F_bt_page_print_tuples(m, v147)
																			mBase = m.M
																			v149 = m.ExcPending
																			if v149 != 0 {
																				return int64(0)
																			} else {
																				v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																				v151 = int32(1)
																				v152 = v150 + v151
																				*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																				v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																				*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																				v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																				*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																				v180 = v148
																				m.G0 = v11 + int32(32)
																				return v180
																			}
																		} else {
																			F_end_MultiFuncCall(m, l0)
																			mBase = m.M
																			v162 = m.ExcPending
																			if v162 != 0 {
																				return int64(0)
																			} else {
																				v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																				*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																				v172 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																				v180 = int64(0)
																				m.G0 = v11 + int32(32)
																				return v180
																			}
																		}
																	}
																}
															}
														} else {
															v99 = F_errstart(m, int32(18), int32(0))
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return int64(0)
															} else {
																if v99 != 0 {
																	F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_10), int32(0))
																	mBase = m.M
																	v104 = m.ExcPending
																	if v104 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(796), int32(_a_F_bt_page_items_bytea_6))
																		mBase = m.M
																		v109 = m.ExcPending
																		if v109 != 0 {
																			return int64(0)
																		} else {
																			*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = int64(0)
																			v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
																			v113 = v112
																			v116 = v113 & int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
																			v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
																			v119 = int32(0)
																			*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
																			v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return int64(0)
																			} else {
																				if v125 != int32(1) {
																					F_errstart_cold(m, int32(21), int32(0))
																					mBase = m.M
																					v269 = m.ExcPending
																					if v269 != 0 {
																						return int64(0)
																					} else {
																						F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																						mBase = m.M
																						v273 = m.ExcPending
																						if v273 != 0 {
																							return int64(0)
																						} else {
																							F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																							mBase = m.M
																							v278 = m.ExcPending
																							if v278 != 0 {
																								return int64(0)
																							} else {
																								base.Wasm_trap_unreachable()
																								for {
																								}
																							}
																						}
																					}
																				} else {
																					v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																					v130 = F_BlessTupleDesc(m, v129)
																					mBase = m.M
																					v131 = m.ExcPending
																					if v131 != 0 {
																						return int64(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																						*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																						*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																						v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																						v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																						v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																						v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																						if base.Ui64(v144) < base.Ui64(v145) {
																							v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																							v148 = F_bt_page_print_tuples(m, v147)
																							mBase = m.M
																							v149 = m.ExcPending
																							if v149 != 0 {
																								return int64(0)
																							} else {
																								v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																								v151 = int32(1)
																								v152 = v150 + v151
																								*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																								v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																								*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																								v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																								v180 = v148
																								m.G0 = v11 + int32(32)
																								return v180
																							}
																						} else {
																							F_end_MultiFuncCall(m, l0)
																							mBase = m.M
																							v162 = m.ExcPending
																							if v162 != 0 {
																								return int64(0)
																							} else {
																								v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																								v172 = int32(1)
																								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																								v180 = int64(0)
																								m.G0 = v11 + int32(32)
																								return v180
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = int64(0)
																	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
																	v113 = v112
																	v116 = v113 & int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
																	v119 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
																	v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
																	mBase = m.M
																	v126 = m.ExcPending
																	if v126 != 0 {
																		return int64(0)
																	} else {
																		if v125 != int32(1) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v269 = m.ExcPending
																			if v269 != 0 {
																				return int64(0)
																			} else {
																				F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																				mBase = m.M
																				v273 = m.ExcPending
																				if v273 != 0 {
																					return int64(0)
																				} else {
																					F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																					mBase = m.M
																					v278 = m.ExcPending
																					if v278 != 0 {
																						return int64(0)
																					} else {
																						base.Wasm_trap_unreachable()
																						for {
																						}
																					}
																				}
																			}
																		} else {
																			v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																			v130 = F_BlessTupleDesc(m, v129)
																			mBase = m.M
																			v131 = m.ExcPending
																			if v131 != 0 {
																				return int64(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																				*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																				*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																				v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																				v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																				v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																				v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																				if base.Ui64(v144) < base.Ui64(v145) {
																					v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																					v148 = F_bt_page_print_tuples(m, v147)
																					mBase = m.M
																					v149 = m.ExcPending
																					if v149 != 0 {
																						return int64(0)
																					} else {
																						v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																						v151 = int32(1)
																						v152 = v150 + v151
																						*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																						v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																						*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																						v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																						*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																						v180 = v148
																						m.G0 = v11 + int32(32)
																						return v180
																					}
																				} else {
																					F_end_MultiFuncCall(m, l0)
																					mBase = m.M
																					v162 = m.ExcPending
																					if v162 != 0 {
																						return int64(0)
																					} else {
																						v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																						*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																						v172 = int32(1)
																						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																						v180 = int64(0)
																						m.G0 = v11 + int32(32)
																						return v180
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														v66 = F_errstart(m, int32(18), int32(0))
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return int64(0)
														} else {
															if v66 == int32(0) {
																v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
																if v79&int32(4) == int32(0) {
																	v84 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
																	v85 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v84)+12)))
																	if base.Ui64(int64(25)) <= base.Ui64(v85) {
																		v95 = int64(base.Ui64(v85+int64(262120))>>(uint(int64(2))%64)) & int64(65535)
																	} else {
																		v95 = int64(0)
																	}
																	*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v95
																	v113 = v79
																	v116 = v113 & int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
																	v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
																	v119 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
																	v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
																	mBase = m.M
																	v126 = m.ExcPending
																	if v126 != 0 {
																		return int64(0)
																	} else {
																		if v125 != int32(1) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v269 = m.ExcPending
																			if v269 != 0 {
																				return int64(0)
																			} else {
																				F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																				mBase = m.M
																				v273 = m.ExcPending
																				if v273 != 0 {
																					return int64(0)
																				} else {
																					F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																					mBase = m.M
																					v278 = m.ExcPending
																					if v278 != 0 {
																						return int64(0)
																					} else {
																						base.Wasm_trap_unreachable()
																						for {
																						}
																					}
																				}
																			}
																		} else {
																			v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																			v130 = F_BlessTupleDesc(m, v129)
																			mBase = m.M
																			v131 = m.ExcPending
																			if v131 != 0 {
																				return int64(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																				*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																				*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																				v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																				v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																				v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																				v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																				if base.Ui64(v144) < base.Ui64(v145) {
																					v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																					v148 = F_bt_page_print_tuples(m, v147)
																					mBase = m.M
																					v149 = m.ExcPending
																					if v149 != 0 {
																						return int64(0)
																					} else {
																						v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																						v151 = int32(1)
																						v152 = v150 + v151
																						*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																						v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																						*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																						v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																						*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																						v180 = v148
																						m.G0 = v11 + int32(32)
																						return v180
																					}
																				} else {
																					F_end_MultiFuncCall(m, l0)
																					mBase = m.M
																					v162 = m.ExcPending
																					if v162 != 0 {
																						return int64(0)
																					} else {
																						v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																						*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																						v172 = int32(1)
																						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																						v180 = int64(0)
																						m.G0 = v11 + int32(32)
																						return v180
																					}
																				}
																			}
																		}
																	}
																} else {
																	v99 = F_errstart(m, int32(18), int32(0))
																	mBase = m.M
																	v100 = m.ExcPending
																	if v100 != 0 {
																		return int64(0)
																	} else {
																		if v99 != 0 {
																			F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_10), int32(0))
																			mBase = m.M
																			v104 = m.ExcPending
																			if v104 != 0 {
																				return int64(0)
																			} else {
																				F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(796), int32(_a_F_bt_page_items_bytea_6))
																				mBase = m.M
																				v109 = m.ExcPending
																				if v109 != 0 {
																					return int64(0)
																				} else {
																					*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = int64(0)
																					v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
																					v113 = v112
																					v116 = v113 & int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
																					v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
																					v119 = int32(0)
																					*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
																					v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
																					mBase = m.M
																					v126 = m.ExcPending
																					if v126 != 0 {
																						return int64(0)
																					} else {
																						if v125 != int32(1) {
																							F_errstart_cold(m, int32(21), int32(0))
																							mBase = m.M
																							v269 = m.ExcPending
																							if v269 != 0 {
																								return int64(0)
																							} else {
																								F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																								mBase = m.M
																								v273 = m.ExcPending
																								if v273 != 0 {
																									return int64(0)
																								} else {
																									F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																									mBase = m.M
																									v278 = m.ExcPending
																									if v278 != 0 {
																										return int64(0)
																									} else {
																										base.Wasm_trap_unreachable()
																										for {
																										}
																									}
																								}
																							}
																						} else {
																							v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																							v130 = F_BlessTupleDesc(m, v129)
																							mBase = m.M
																							v131 = m.ExcPending
																							if v131 != 0 {
																								return int64(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																								*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																								*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																								v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																								v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																								v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																								v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																								if base.Ui64(v144) < base.Ui64(v145) {
																									v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																									v148 = F_bt_page_print_tuples(m, v147)
																									mBase = m.M
																									v149 = m.ExcPending
																									if v149 != 0 {
																										return int64(0)
																									} else {
																										v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																										v151 = int32(1)
																										v152 = v150 + v151
																										*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																										v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																										*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																										v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																										*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																										v180 = v148
																										m.G0 = v11 + int32(32)
																										return v180
																									}
																								} else {
																									F_end_MultiFuncCall(m, l0)
																									mBase = m.M
																									v162 = m.ExcPending
																									if v162 != 0 {
																										return int64(0)
																									} else {
																										v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																										*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																										v172 = int32(1)
																										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																										v180 = int64(0)
																										m.G0 = v11 + int32(32)
																										return v180
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		} else {
																			*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = int64(0)
																			v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
																			v113 = v112
																			v116 = v113 & int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
																			v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
																			v119 = int32(0)
																			*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
																			v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return int64(0)
																			} else {
																				if v125 != int32(1) {
																					F_errstart_cold(m, int32(21), int32(0))
																					mBase = m.M
																					v269 = m.ExcPending
																					if v269 != 0 {
																						return int64(0)
																					} else {
																						F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																						mBase = m.M
																						v273 = m.ExcPending
																						if v273 != 0 {
																							return int64(0)
																						} else {
																							F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																							mBase = m.M
																							v278 = m.ExcPending
																							if v278 != 0 {
																								return int64(0)
																							} else {
																								base.Wasm_trap_unreachable()
																								for {
																								}
																							}
																						}
																					}
																				} else {
																					v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																					v130 = F_BlessTupleDesc(m, v129)
																					mBase = m.M
																					v131 = m.ExcPending
																					if v131 != 0 {
																						return int64(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																						*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																						*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																						v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																						v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																						v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																						v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																						if base.Ui64(v144) < base.Ui64(v145) {
																							v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																							v148 = F_bt_page_print_tuples(m, v147)
																							mBase = m.M
																							v149 = m.ExcPending
																							if v149 != 0 {
																								return int64(0)
																							} else {
																								v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																								v151 = int32(1)
																								v152 = v150 + v151
																								*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																								v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																								*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																								v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																								v180 = v148
																								m.G0 = v11 + int32(32)
																								return v180
																							}
																						} else {
																							F_end_MultiFuncCall(m, l0)
																							mBase = m.M
																							v162 = m.ExcPending
																							if v162 != 0 {
																								return int64(0)
																							} else {
																								v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																								v172 = int32(1)
																								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																								v180 = int64(0)
																								m.G0 = v11 + int32(32)
																								return v180
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_11), int32(0))
																mBase = m.M
																v73 = m.ExcPending
																if v73 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(789), int32(_a_F_bt_page_items_bytea_6))
																	mBase = m.M
																	v78 = m.ExcPending
																	if v78 != 0 {
																		return int64(0)
																	} else {
																		v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
																		if v79&int32(4) == int32(0) {
																			v84 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
																			v85 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v84)+12)))
																			if base.Ui64(int64(25)) <= base.Ui64(v85) {
																				v95 = int64(base.Ui64(v85+int64(262120))>>(uint(int64(2))%64)) & int64(65535)
																			} else {
																				v95 = int64(0)
																			}
																			*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v95
																			v113 = v79
																			v116 = v113 & int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
																			v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
																			v119 = int32(0)
																			*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
																			v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return int64(0)
																			} else {
																				if v125 != int32(1) {
																					F_errstart_cold(m, int32(21), int32(0))
																					mBase = m.M
																					v269 = m.ExcPending
																					if v269 != 0 {
																						return int64(0)
																					} else {
																						F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																						mBase = m.M
																						v273 = m.ExcPending
																						if v273 != 0 {
																							return int64(0)
																						} else {
																							F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																							mBase = m.M
																							v278 = m.ExcPending
																							if v278 != 0 {
																								return int64(0)
																							} else {
																								base.Wasm_trap_unreachable()
																								for {
																								}
																							}
																						}
																					}
																				} else {
																					v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																					v130 = F_BlessTupleDesc(m, v129)
																					mBase = m.M
																					v131 = m.ExcPending
																					if v131 != 0 {
																						return int64(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																						*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																						*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																						v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																						v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																						v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																						v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																						if base.Ui64(v144) < base.Ui64(v145) {
																							v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																							v148 = F_bt_page_print_tuples(m, v147)
																							mBase = m.M
																							v149 = m.ExcPending
																							if v149 != 0 {
																								return int64(0)
																							} else {
																								v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																								v151 = int32(1)
																								v152 = v150 + v151
																								*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																								v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																								*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																								v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																								v180 = v148
																								m.G0 = v11 + int32(32)
																								return v180
																							}
																						} else {
																							F_end_MultiFuncCall(m, l0)
																							mBase = m.M
																							v162 = m.ExcPending
																							if v162 != 0 {
																								return int64(0)
																							} else {
																								v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																								v172 = int32(1)
																								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																								v180 = int64(0)
																								m.G0 = v11 + int32(32)
																								return v180
																							}
																						}
																					}
																				}
																			}
																		} else {
																			v99 = F_errstart(m, int32(18), int32(0))
																			mBase = m.M
																			v100 = m.ExcPending
																			if v100 != 0 {
																				return int64(0)
																			} else {
																				if v99 != 0 {
																					F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_10), int32(0))
																					mBase = m.M
																					v104 = m.ExcPending
																					if v104 != 0 {
																						return int64(0)
																					} else {
																						F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(796), int32(_a_F_bt_page_items_bytea_6))
																						mBase = m.M
																						v109 = m.ExcPending
																						if v109 != 0 {
																							return int64(0)
																						} else {
																							*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = int64(0)
																							v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
																							v113 = v112
																							v116 = v113 & int32(1)
																							*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
																							v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
																							v119 = int32(0)
																							*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
																							v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
																							mBase = m.M
																							v126 = m.ExcPending
																							if v126 != 0 {
																								return int64(0)
																							} else {
																								if v125 != int32(1) {
																									F_errstart_cold(m, int32(21), int32(0))
																									mBase = m.M
																									v269 = m.ExcPending
																									if v269 != 0 {
																										return int64(0)
																									} else {
																										F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																										mBase = m.M
																										v273 = m.ExcPending
																										if v273 != 0 {
																											return int64(0)
																										} else {
																											F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																											mBase = m.M
																											v278 = m.ExcPending
																											if v278 != 0 {
																												return int64(0)
																											} else {
																												base.Wasm_trap_unreachable()
																												for {
																												}
																											}
																										}
																									}
																								} else {
																									v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																									v130 = F_BlessTupleDesc(m, v129)
																									mBase = m.M
																									v131 = m.ExcPending
																									if v131 != 0 {
																										return int64(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																										*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																										*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																										v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																										v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																										v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																										v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																										if base.Ui64(v144) < base.Ui64(v145) {
																											v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																											v148 = F_bt_page_print_tuples(m, v147)
																											mBase = m.M
																											v149 = m.ExcPending
																											if v149 != 0 {
																												return int64(0)
																											} else {
																												v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																												v151 = int32(1)
																												v152 = v150 + v151
																												*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																												v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																												*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																												v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																												*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																												v180 = v148
																												m.G0 = v11 + int32(32)
																												return v180
																											}
																										} else {
																											F_end_MultiFuncCall(m, l0)
																											mBase = m.M
																											v162 = m.ExcPending
																											if v162 != 0 {
																												return int64(0)
																											} else {
																												v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																												*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																												v172 = int32(1)
																												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																												v180 = int64(0)
																												m.G0 = v11 + int32(32)
																												return v180
																											}
																										}
																									}
																								}
																							}
																						}
																					}
																				} else {
																					*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = int64(0)
																					v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
																					v113 = v112
																					v116 = v113 & int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
																					v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
																					v119 = int32(0)
																					*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
																					v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
																					mBase = m.M
																					v126 = m.ExcPending
																					if v126 != 0 {
																						return int64(0)
																					} else {
																						if v125 != int32(1) {
																							F_errstart_cold(m, int32(21), int32(0))
																							mBase = m.M
																							v269 = m.ExcPending
																							if v269 != 0 {
																								return int64(0)
																							} else {
																								F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																								mBase = m.M
																								v273 = m.ExcPending
																								if v273 != 0 {
																									return int64(0)
																								} else {
																									F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																									mBase = m.M
																									v278 = m.ExcPending
																									if v278 != 0 {
																										return int64(0)
																									} else {
																										base.Wasm_trap_unreachable()
																										for {
																										}
																									}
																								}
																							}
																						} else {
																							v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																							v130 = F_BlessTupleDesc(m, v129)
																							mBase = m.M
																							v131 = m.ExcPending
																							if v131 != 0 {
																								return int64(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																								*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																								*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																								v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																								v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																								v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																								v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																								if base.Ui64(v144) < base.Ui64(v145) {
																									v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																									v148 = F_bt_page_print_tuples(m, v147)
																									mBase = m.M
																									v149 = m.ExcPending
																									if v149 != 0 {
																										return int64(0)
																									} else {
																										v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																										v151 = int32(1)
																										v152 = v150 + v151
																										*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																										v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																										*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																										v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																										*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																										v180 = v148
																										m.G0 = v11 + int32(32)
																										return v180
																									}
																								} else {
																									F_end_MultiFuncCall(m, l0)
																									mBase = m.M
																									v162 = m.ExcPending
																									if v162 != 0 {
																										return int64(0)
																									} else {
																										v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																										*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																										v172 = int32(1)
																										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																										v180 = int64(0)
																										m.G0 = v11 + int32(32)
																										return v180
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
														}
													}
												}
											} else {
												if v54&int32(4) == int32(0) {
													v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
													if v79&int32(4) == int32(0) {
														v84 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
														v85 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v84)+12)))
														if base.Ui64(int64(25)) <= base.Ui64(v85) {
															v95 = int64(base.Ui64(v85+int64(262120))>>(uint(int64(2))%64)) & int64(65535)
														} else {
															v95 = int64(0)
														}
														*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v95
														v113 = v79
														v116 = v113 & int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
														v119 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
														v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return int64(0)
														} else {
															if v125 != int32(1) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v269 = m.ExcPending
																if v269 != 0 {
																	return int64(0)
																} else {
																	F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																	mBase = m.M
																	v273 = m.ExcPending
																	if v273 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																		mBase = m.M
																		v278 = m.ExcPending
																		if v278 != 0 {
																			return int64(0)
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	}
																}
															} else {
																v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																v130 = F_BlessTupleDesc(m, v129)
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return int64(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																	*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																	v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																	v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																	if base.Ui64(v144) < base.Ui64(v145) {
																		v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																		v148 = F_bt_page_print_tuples(m, v147)
																		mBase = m.M
																		v149 = m.ExcPending
																		if v149 != 0 {
																			return int64(0)
																		} else {
																			v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																			v151 = int32(1)
																			v152 = v150 + v151
																			*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																			v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																			*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																			v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																			*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																			v180 = v148
																			m.G0 = v11 + int32(32)
																			return v180
																		}
																	} else {
																		F_end_MultiFuncCall(m, l0)
																		mBase = m.M
																		v162 = m.ExcPending
																		if v162 != 0 {
																			return int64(0)
																		} else {
																			v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																			*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																			v172 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																			v180 = int64(0)
																			m.G0 = v11 + int32(32)
																			return v180
																		}
																	}
																}
															}
														}
													} else {
														v99 = F_errstart(m, int32(18), int32(0))
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return int64(0)
														} else {
															if v99 != 0 {
																F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_10), int32(0))
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(796), int32(_a_F_bt_page_items_bytea_6))
																	mBase = m.M
																	v109 = m.ExcPending
																	if v109 != 0 {
																		return int64(0)
																	} else {
																		*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = int64(0)
																		v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
																		v113 = v112
																		v116 = v113 & int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
																		v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
																		v119 = int32(0)
																		*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
																		v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
																		mBase = m.M
																		v126 = m.ExcPending
																		if v126 != 0 {
																			return int64(0)
																		} else {
																			if v125 != int32(1) {
																				F_errstart_cold(m, int32(21), int32(0))
																				mBase = m.M
																				v269 = m.ExcPending
																				if v269 != 0 {
																					return int64(0)
																				} else {
																					F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																					mBase = m.M
																					v273 = m.ExcPending
																					if v273 != 0 {
																						return int64(0)
																					} else {
																						F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																						mBase = m.M
																						v278 = m.ExcPending
																						if v278 != 0 {
																							return int64(0)
																						} else {
																							base.Wasm_trap_unreachable()
																							for {
																							}
																						}
																					}
																				}
																			} else {
																				v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																				v130 = F_BlessTupleDesc(m, v129)
																				mBase = m.M
																				v131 = m.ExcPending
																				if v131 != 0 {
																					return int64(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																					*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																					*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																					v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																					v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																					v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																					v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																					if base.Ui64(v144) < base.Ui64(v145) {
																						v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																						v148 = F_bt_page_print_tuples(m, v147)
																						mBase = m.M
																						v149 = m.ExcPending
																						if v149 != 0 {
																							return int64(0)
																						} else {
																							v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																							v151 = int32(1)
																							v152 = v150 + v151
																							*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																							v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																							*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																							v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																							*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																							v180 = v148
																							m.G0 = v11 + int32(32)
																							return v180
																						}
																					} else {
																						F_end_MultiFuncCall(m, l0)
																						mBase = m.M
																						v162 = m.ExcPending
																						if v162 != 0 {
																							return int64(0)
																						} else {
																							v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																							*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																							v172 = int32(1)
																							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																							v180 = int64(0)
																							m.G0 = v11 + int32(32)
																							return v180
																						}
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = int64(0)
																v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
																v113 = v112
																v116 = v113 & int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
																v119 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
																v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
																mBase = m.M
																v126 = m.ExcPending
																if v126 != 0 {
																	return int64(0)
																} else {
																	if v125 != int32(1) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v269 = m.ExcPending
																		if v269 != 0 {
																			return int64(0)
																		} else {
																			F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																			mBase = m.M
																			v273 = m.ExcPending
																			if v273 != 0 {
																				return int64(0)
																			} else {
																				F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																				mBase = m.M
																				v278 = m.ExcPending
																				if v278 != 0 {
																					return int64(0)
																				} else {
																					base.Wasm_trap_unreachable()
																					for {
																					}
																				}
																			}
																		}
																	} else {
																		v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																		v130 = F_BlessTupleDesc(m, v129)
																		mBase = m.M
																		v131 = m.ExcPending
																		if v131 != 0 {
																			return int64(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																			*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																			*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																			v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																			v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																			v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																			v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																			if base.Ui64(v144) < base.Ui64(v145) {
																				v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																				v148 = F_bt_page_print_tuples(m, v147)
																				mBase = m.M
																				v149 = m.ExcPending
																				if v149 != 0 {
																					return int64(0)
																				} else {
																					v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																					v151 = int32(1)
																					v152 = v150 + v151
																					*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																					v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																					*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																					v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																					*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																					v180 = v148
																					m.G0 = v11 + int32(32)
																					return v180
																				}
																			} else {
																				F_end_MultiFuncCall(m, l0)
																				mBase = m.M
																				v162 = m.ExcPending
																				if v162 != 0 {
																					return int64(0)
																				} else {
																					v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																					*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																					v172 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																					v180 = int64(0)
																					m.G0 = v11 + int32(32)
																					return v180
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v66 = F_errstart(m, int32(18), int32(0))
													mBase = m.M
													v67 = m.ExcPending
													if v67 != 0 {
														return int64(0)
													} else {
														if v66 == int32(0) {
															v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
															if v79&int32(4) == int32(0) {
																v84 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
																v85 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v84)+12)))
																if base.Ui64(int64(25)) <= base.Ui64(v85) {
																	v95 = int64(base.Ui64(v85+int64(262120))>>(uint(int64(2))%64)) & int64(65535)
																} else {
																	v95 = int64(0)
																}
																*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v95
																v113 = v79
																v116 = v113 & int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
																v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
																v119 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
																v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
																mBase = m.M
																v126 = m.ExcPending
																if v126 != 0 {
																	return int64(0)
																} else {
																	if v125 != int32(1) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v269 = m.ExcPending
																		if v269 != 0 {
																			return int64(0)
																		} else {
																			F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																			mBase = m.M
																			v273 = m.ExcPending
																			if v273 != 0 {
																				return int64(0)
																			} else {
																				F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																				mBase = m.M
																				v278 = m.ExcPending
																				if v278 != 0 {
																					return int64(0)
																				} else {
																					base.Wasm_trap_unreachable()
																					for {
																					}
																				}
																			}
																		}
																	} else {
																		v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																		v130 = F_BlessTupleDesc(m, v129)
																		mBase = m.M
																		v131 = m.ExcPending
																		if v131 != 0 {
																			return int64(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																			*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																			*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																			v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																			v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																			v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																			v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																			if base.Ui64(v144) < base.Ui64(v145) {
																				v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																				v148 = F_bt_page_print_tuples(m, v147)
																				mBase = m.M
																				v149 = m.ExcPending
																				if v149 != 0 {
																					return int64(0)
																				} else {
																					v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																					v151 = int32(1)
																					v152 = v150 + v151
																					*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																					v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																					*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																					v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																					*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																					v180 = v148
																					m.G0 = v11 + int32(32)
																					return v180
																				}
																			} else {
																				F_end_MultiFuncCall(m, l0)
																				mBase = m.M
																				v162 = m.ExcPending
																				if v162 != 0 {
																					return int64(0)
																				} else {
																					v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																					*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																					v172 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																					v180 = int64(0)
																					m.G0 = v11 + int32(32)
																					return v180
																				}
																			}
																		}
																	}
																}
															} else {
																v99 = F_errstart(m, int32(18), int32(0))
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return int64(0)
																} else {
																	if v99 != 0 {
																		F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_10), int32(0))
																		mBase = m.M
																		v104 = m.ExcPending
																		if v104 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(796), int32(_a_F_bt_page_items_bytea_6))
																			mBase = m.M
																			v109 = m.ExcPending
																			if v109 != 0 {
																				return int64(0)
																			} else {
																				*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = int64(0)
																				v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
																				v113 = v112
																				v116 = v113 & int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
																				v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
																				v119 = int32(0)
																				*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
																				v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
																				mBase = m.M
																				v126 = m.ExcPending
																				if v126 != 0 {
																					return int64(0)
																				} else {
																					if v125 != int32(1) {
																						F_errstart_cold(m, int32(21), int32(0))
																						mBase = m.M
																						v269 = m.ExcPending
																						if v269 != 0 {
																							return int64(0)
																						} else {
																							F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																							mBase = m.M
																							v273 = m.ExcPending
																							if v273 != 0 {
																								return int64(0)
																							} else {
																								F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																								mBase = m.M
																								v278 = m.ExcPending
																								if v278 != 0 {
																									return int64(0)
																								} else {
																									base.Wasm_trap_unreachable()
																									for {
																									}
																								}
																							}
																						}
																					} else {
																						v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																						v130 = F_BlessTupleDesc(m, v129)
																						mBase = m.M
																						v131 = m.ExcPending
																						if v131 != 0 {
																							return int64(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																							*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																							*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																							v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																							v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																							v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																							v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																							if base.Ui64(v144) < base.Ui64(v145) {
																								v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																								v148 = F_bt_page_print_tuples(m, v147)
																								mBase = m.M
																								v149 = m.ExcPending
																								if v149 != 0 {
																									return int64(0)
																								} else {
																									v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																									v151 = int32(1)
																									v152 = v150 + v151
																									*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																									v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																									*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																									v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																									*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																									v180 = v148
																									m.G0 = v11 + int32(32)
																									return v180
																								}
																							} else {
																								F_end_MultiFuncCall(m, l0)
																								mBase = m.M
																								v162 = m.ExcPending
																								if v162 != 0 {
																									return int64(0)
																								} else {
																									v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																									*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																									v172 = int32(1)
																									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																									v180 = int64(0)
																									m.G0 = v11 + int32(32)
																									return v180
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	} else {
																		*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = int64(0)
																		v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
																		v113 = v112
																		v116 = v113 & int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
																		v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
																		v119 = int32(0)
																		*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
																		v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
																		mBase = m.M
																		v126 = m.ExcPending
																		if v126 != 0 {
																			return int64(0)
																		} else {
																			if v125 != int32(1) {
																				F_errstart_cold(m, int32(21), int32(0))
																				mBase = m.M
																				v269 = m.ExcPending
																				if v269 != 0 {
																					return int64(0)
																				} else {
																					F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																					mBase = m.M
																					v273 = m.ExcPending
																					if v273 != 0 {
																						return int64(0)
																					} else {
																						F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																						mBase = m.M
																						v278 = m.ExcPending
																						if v278 != 0 {
																							return int64(0)
																						} else {
																							base.Wasm_trap_unreachable()
																							for {
																							}
																						}
																					}
																				}
																			} else {
																				v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																				v130 = F_BlessTupleDesc(m, v129)
																				mBase = m.M
																				v131 = m.ExcPending
																				if v131 != 0 {
																					return int64(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																					*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																					*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																					v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																					v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																					v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																					v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																					if base.Ui64(v144) < base.Ui64(v145) {
																						v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																						v148 = F_bt_page_print_tuples(m, v147)
																						mBase = m.M
																						v149 = m.ExcPending
																						if v149 != 0 {
																							return int64(0)
																						} else {
																							v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																							v151 = int32(1)
																							v152 = v150 + v151
																							*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																							v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																							*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																							v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																							*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																							v180 = v148
																							m.G0 = v11 + int32(32)
																							return v180
																						}
																					} else {
																						F_end_MultiFuncCall(m, l0)
																						mBase = m.M
																						v162 = m.ExcPending
																						if v162 != 0 {
																							return int64(0)
																						} else {
																							v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																							*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																							v172 = int32(1)
																							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																							v180 = int64(0)
																							m.G0 = v11 + int32(32)
																							return v180
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														} else {
															F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_11), int32(0))
															mBase = m.M
															v73 = m.ExcPending
															if v73 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(789), int32(_a_F_bt_page_items_bytea_6))
																mBase = m.M
																v78 = m.ExcPending
																if v78 != 0 {
																	return int64(0)
																} else {
																	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
																	if v79&int32(4) == int32(0) {
																		v84 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
																		v85 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v84)+12)))
																		if base.Ui64(int64(25)) <= base.Ui64(v85) {
																			v95 = int64(base.Ui64(v85+int64(262120))>>(uint(int64(2))%64)) & int64(65535)
																		} else {
																			v95 = int64(0)
																		}
																		*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v95
																		v113 = v79
																		v116 = v113 & int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
																		v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
																		v119 = int32(0)
																		*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
																		v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
																		mBase = m.M
																		v126 = m.ExcPending
																		if v126 != 0 {
																			return int64(0)
																		} else {
																			if v125 != int32(1) {
																				F_errstart_cold(m, int32(21), int32(0))
																				mBase = m.M
																				v269 = m.ExcPending
																				if v269 != 0 {
																					return int64(0)
																				} else {
																					F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																					mBase = m.M
																					v273 = m.ExcPending
																					if v273 != 0 {
																						return int64(0)
																					} else {
																						F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																						mBase = m.M
																						v278 = m.ExcPending
																						if v278 != 0 {
																							return int64(0)
																						} else {
																							base.Wasm_trap_unreachable()
																							for {
																							}
																						}
																					}
																				}
																			} else {
																				v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																				v130 = F_BlessTupleDesc(m, v129)
																				mBase = m.M
																				v131 = m.ExcPending
																				if v131 != 0 {
																					return int64(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																					*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																					*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																					v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																					v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																					v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																					v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																					if base.Ui64(v144) < base.Ui64(v145) {
																						v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																						v148 = F_bt_page_print_tuples(m, v147)
																						mBase = m.M
																						v149 = m.ExcPending
																						if v149 != 0 {
																							return int64(0)
																						} else {
																							v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																							v151 = int32(1)
																							v152 = v150 + v151
																							*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																							v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																							*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																							v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																							*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																							v180 = v148
																							m.G0 = v11 + int32(32)
																							return v180
																						}
																					} else {
																						F_end_MultiFuncCall(m, l0)
																						mBase = m.M
																						v162 = m.ExcPending
																						if v162 != 0 {
																							return int64(0)
																						} else {
																							v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																							*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																							v172 = int32(1)
																							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																							v180 = int64(0)
																							m.G0 = v11 + int32(32)
																							return v180
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v99 = F_errstart(m, int32(18), int32(0))
																		mBase = m.M
																		v100 = m.ExcPending
																		if v100 != 0 {
																			return int64(0)
																		} else {
																			if v99 != 0 {
																				F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_10), int32(0))
																				mBase = m.M
																				v104 = m.ExcPending
																				if v104 != 0 {
																					return int64(0)
																				} else {
																					F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(796), int32(_a_F_bt_page_items_bytea_6))
																					mBase = m.M
																					v109 = m.ExcPending
																					if v109 != 0 {
																						return int64(0)
																					} else {
																						*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = int64(0)
																						v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
																						v113 = v112
																						v116 = v113 & int32(1)
																						*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
																						v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
																						v119 = int32(0)
																						*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
																						v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
																						mBase = m.M
																						v126 = m.ExcPending
																						if v126 != 0 {
																							return int64(0)
																						} else {
																							if v125 != int32(1) {
																								F_errstart_cold(m, int32(21), int32(0))
																								mBase = m.M
																								v269 = m.ExcPending
																								if v269 != 0 {
																									return int64(0)
																								} else {
																									F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																									mBase = m.M
																									v273 = m.ExcPending
																									if v273 != 0 {
																										return int64(0)
																									} else {
																										F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																										mBase = m.M
																										v278 = m.ExcPending
																										if v278 != 0 {
																											return int64(0)
																										} else {
																											base.Wasm_trap_unreachable()
																											for {
																											}
																										}
																									}
																								}
																							} else {
																								v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																								v130 = F_BlessTupleDesc(m, v129)
																								mBase = m.M
																								v131 = m.ExcPending
																								if v131 != 0 {
																									return int64(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																									*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																									*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																									v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																									v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																									v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																									v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																									if base.Ui64(v144) < base.Ui64(v145) {
																										v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																										v148 = F_bt_page_print_tuples(m, v147)
																										mBase = m.M
																										v149 = m.ExcPending
																										if v149 != 0 {
																											return int64(0)
																										} else {
																											v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																											v151 = int32(1)
																											v152 = v150 + v151
																											*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																											v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																											*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																											v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																											*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																											v180 = v148
																											m.G0 = v11 + int32(32)
																											return v180
																										}
																									} else {
																										F_end_MultiFuncCall(m, l0)
																										mBase = m.M
																										v162 = m.ExcPending
																										if v162 != 0 {
																											return int64(0)
																										} else {
																											v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																											*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																											v172 = int32(1)
																											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																											v180 = int64(0)
																											m.G0 = v11 + int32(32)
																											return v180
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			} else {
																				*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = int64(0)
																				v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
																				v113 = v112
																				v116 = v113 & int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(v32)+6)) = uint8(v116)
																				v118 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
																				v119 = int32(0)
																				*(*uint8)(unsafe.Add(mBase, uint32(v32)+7)) = uint8(base.B2i32(v118 == v119))
																				v125 = F_get_call_result_type(m, l0, v119, v11+int32(28))
																				mBase = m.M
																				v126 = m.ExcPending
																				if v126 != 0 {
																					return int64(0)
																				} else {
																					if v125 != int32(1) {
																						F_errstart_cold(m, int32(21), int32(0))
																						mBase = m.M
																						v269 = m.ExcPending
																						if v269 != 0 {
																							return int64(0)
																						} else {
																							F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																							mBase = m.M
																							v273 = m.ExcPending
																							if v273 != 0 {
																								return int64(0)
																							} else {
																								F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(804), int32(_a_F_bt_page_items_bytea_6))
																								mBase = m.M
																								v278 = m.ExcPending
																								if v278 != 0 {
																									return int64(0)
																								} else {
																									base.Wasm_trap_unreachable()
																									for {
																									}
																								}
																							}
																						}
																					} else {
																						v129 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
																						v130 = F_BlessTupleDesc(m, v129)
																						mBase = m.M
																						v131 = m.ExcPending
																						if v131 != 0 {
																							return int64(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v130
																							*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
																							*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
																							v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																							v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
																							v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																							v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
																							if base.Ui64(v144) < base.Ui64(v145) {
																								v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
																								v148 = F_bt_page_print_tuples(m, v147)
																								mBase = m.M
																								v149 = m.ExcPending
																								if v149 != 0 {
																									return int64(0)
																								} else {
																									v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																									v151 = int32(1)
																									v152 = v150 + v151
																									*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																									v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																									*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																									v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																									*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																									v180 = v148
																									m.G0 = v11 + int32(32)
																									return v180
																								}
																							} else {
																								F_end_MultiFuncCall(m, l0)
																								mBase = m.M
																								v162 = m.ExcPending
																								if v162 != 0 {
																									return int64(0)
																								} else {
																									v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																									*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																									v172 = int32(1)
																									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																									v180 = int64(0)
																									m.G0 = v11 + int32(32)
																									return v180
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
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
					v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
					v145 = *(*int64)(unsafe.Add(mBase, uint32(v143)+8))
					if base.Ui64(v144) < base.Ui64(v145) {
						v147 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
						v148 = F_bt_page_print_tuples(m, v147)
						mBase = m.M
						v149 = m.ExcPending
						if v149 != 0 {
							return int64(0)
						} else {
							v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
							v151 = int32(1)
							v152 = v150 + v151
							*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
							v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
							*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
							v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
							v180 = v148
							m.G0 = v11 + int32(32)
							return v180
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v162 = m.ExcPending
						if v162 != 0 {
							return int64(0)
						} else {
							v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
							v172 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
							v180 = int64(0)
							m.G0 = v11 + int32(32)
							return v180
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v188 = m.ExcPending
				if v188 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v191 = m.ExcPending
					if v191 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_bt_page_items_bytea_12), int32(0))
						mBase = m.M
						v195 = m.ExcPending
						if v195 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(744), int32(_a_F_bt_page_items_bytea_6))
							mBase = m.M
							v200 = m.ExcPending
							if v200 != 0 {
								return int64(0)
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
}
func F_bt_page_items_internal(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v23 int32
	_ = v23
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int64
	_ = v93
	var v103 int64
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int64
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int64
	_ = v178
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v20 = F_superuser(m)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			if v20 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
				if v23 == int32(0) {
					v26 = F_init_MultiFuncCall(m, l0)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						v28 = F_textToQualifiedNameList(m, v15)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int64(0)
						} else {
							v30 = F_makeRangeVarFromNameList(m, v28)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int64(0)
							} else {
								v33 = F_relation_openrv(m, v30, int32(1))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int64(0)
								} else {
									if l1 != 0 {
										v37 = v19
									} else {
										v37 = v19 & int64(4294967295)
									}
									F_bt_index_block_validate(m, v33, v37)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int64(0)
									} else {
										v41 = F_ReadBuffer(m, v33, base.I32_wrap_i64(v37))
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return int64(0)
										} else {
											F_LockBufferInternal(m, v41, int32(1))
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return int64(0)
											} else {
												v46 = int32(_a_F_bt_page_items_internal_0)
												v47 = *(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_internal[0]))
												v49 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
												*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_internal[0])) = v49
												v52 = F_palloc(m, int32(12))
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return int64(0)
												} else {
													v55 = F_palloc(m, int32(_a_F_bt_page_items_internal_1))
													mBase = m.M
													v56 = m.ExcPending
													if v56 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v52))) = v55
														if v41 < int32(0) {
															v61 = *(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_internal[1]))
															v67 = *(*int32)(unsafe.Add(mBase, uint32(v61+(v41^int32(-1))<<(uint(int32(2))%32))))
															v75 = v67
														} else {
															v69 = *(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_internal[2]))
															v75 = v69 + v41<<(uint(int32(13))%32) + int32(-8192)
														}
														base.MemoryCopy(m, v55, v75, int32(_a_F_bt_page_items_internal_1))
														F_UnlockReleaseBuffer(m, v41)
														mBase = m.M
														v79 = m.ExcPending
														if v79 != 0 {
															return int64(0)
														} else {
															F_relation_close(m, v33, int32(1))
															mBase = m.M
															v82 = m.ExcPending
															if v82 != 0 {
																return int64(0)
															} else {
																v83 = int32(1)
																*(*uint16)(unsafe.Add(mBase, uint32(v52)+4)) = uint16(v83)
																v85 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
																v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+16)))
																v87 = v85 + v86
																v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+12)))
																if v88&int32(4) == int32(0) {
																	v93 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v85)+12)))
																	if base.Ui64(int64(25)) <= base.Ui64(v93) {
																		v103 = int64(base.Ui64(v93+int64(262120))>>(uint(int64(2))%64)) & int64(65535)
																	} else {
																		v103 = int64(0)
																	}
																	*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = v103
																	v122 = v88
																	v124 = v122 & int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v52)+6)) = uint8(v124)
																	v126 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
																	v127 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(v52)+7)) = uint8(base.B2i32(v126 == v127))
																	v133 = F_get_call_result_type(m, l0, v127, v12+int32(12))
																	mBase = m.M
																	v134 = m.ExcPending
																	if v134 != 0 {
																		return int64(0)
																	} else {
																		if v133 != int32(1) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v203 = m.ExcPending
																			if v203 != 0 {
																				return int64(0)
																			} else {
																				F_errmsg_internal(m, int32(_a_F_bt_page_items_internal_2), int32(0))
																				mBase = m.M
																				v207 = m.ExcPending
																				if v207 != 0 {
																					return int64(0)
																				} else {
																					F_errfinish(m, int32(_a_F_bt_page_items_internal_3), int32(688), int32(_a_F_bt_page_items_internal_4))
																					mBase = m.M
																					v212 = m.ExcPending
																					if v212 != 0 {
																						return int64(0)
																					} else {
																						base.Wasm_trap_unreachable()
																						for {
																						}
																					}
																				}
																			}
																		} else {
																			v137 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
																			v138 = F_BlessTupleDesc(m, v137)
																			mBase = m.M
																			v139 = m.ExcPending
																			if v139 != 0 {
																				return int64(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v138
																				*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v52
																				*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_internal[0])) = v47
																				v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																				v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+16))
																				v153 = *(*int64)(unsafe.Add(mBase, uint32(v152)))
																				v154 = *(*int64)(unsafe.Add(mBase, uint32(v152)+8))
																				if base.Ui64(v153) < base.Ui64(v154) {
																					v156 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
																					v157 = F_bt_page_print_tuples(m, v156)
																					mBase = m.M
																					v158 = m.ExcPending
																					if v158 != 0 {
																						return int64(0)
																					} else {
																						v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)))
																						v160 = int32(1)
																						v161 = v159 + v160
																						*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)) = uint16(v161)
																						v163 = *(*int64)(unsafe.Add(mBase, uint32(v152)))
																						*(*int64)(unsafe.Add(mBase, uint32(v152))) = v163 + int64(1)
																						v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																						*(*int32)(unsafe.Add(mBase, uint32(v167)+20)) = v160
																						v178 = v157
																						m.G0 = v12 + int32(16)
																						return v178
																					}
																				} else {
																					F_end_MultiFuncCall(m, l0)
																					mBase = m.M
																					v171 = m.ExcPending
																					if v171 != 0 {
																						return int64(0)
																					} else {
																						v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																						*(*int32)(unsafe.Add(mBase, uint32(v172)+20)) = int32(2)
																						v175 = int32(1)
																						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v175)
																						v178 = int64(0)
																						m.G0 = v12 + int32(16)
																						return v178
																					}
																				}
																			}
																		}
																	}
																} else {
																	v107 = F_errstart(m, int32(18), int32(0))
																	mBase = m.M
																	v108 = m.ExcPending
																	if v108 != 0 {
																		return int64(0)
																	} else {
																		if v107 != 0 {
																			*(*int64)(unsafe.Add(mBase, uint32(v12))) = v37
																			F_errmsg_internal(m, int32(_a_F_bt_page_items_internal_5), v12)
																			mBase = m.M
																			v112 = m.ExcPending
																			if v112 != 0 {
																				return int64(0)
																			} else {
																				F_errfinish(m, int32(_a_F_bt_page_items_internal_3), int32(680), int32(_a_F_bt_page_items_internal_4))
																				mBase = m.M
																				v117 = m.ExcPending
																				if v117 != 0 {
																					return int64(0)
																				} else {
																					*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = int64(0)
																					v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+12)))
																					v122 = v120
																					v124 = v122 & int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(v52)+6)) = uint8(v124)
																					v126 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
																					v127 = int32(0)
																					*(*uint8)(unsafe.Add(mBase, uint32(v52)+7)) = uint8(base.B2i32(v126 == v127))
																					v133 = F_get_call_result_type(m, l0, v127, v12+int32(12))
																					mBase = m.M
																					v134 = m.ExcPending
																					if v134 != 0 {
																						return int64(0)
																					} else {
																						if v133 != int32(1) {
																							F_errstart_cold(m, int32(21), int32(0))
																							mBase = m.M
																							v203 = m.ExcPending
																							if v203 != 0 {
																								return int64(0)
																							} else {
																								F_errmsg_internal(m, int32(_a_F_bt_page_items_internal_2), int32(0))
																								mBase = m.M
																								v207 = m.ExcPending
																								if v207 != 0 {
																									return int64(0)
																								} else {
																									F_errfinish(m, int32(_a_F_bt_page_items_internal_3), int32(688), int32(_a_F_bt_page_items_internal_4))
																									mBase = m.M
																									v212 = m.ExcPending
																									if v212 != 0 {
																										return int64(0)
																									} else {
																										base.Wasm_trap_unreachable()
																										for {
																										}
																									}
																								}
																							}
																						} else {
																							v137 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
																							v138 = F_BlessTupleDesc(m, v137)
																							mBase = m.M
																							v139 = m.ExcPending
																							if v139 != 0 {
																								return int64(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v138
																								*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v52
																								*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_internal[0])) = v47
																								v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																								v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+16))
																								v153 = *(*int64)(unsafe.Add(mBase, uint32(v152)))
																								v154 = *(*int64)(unsafe.Add(mBase, uint32(v152)+8))
																								if base.Ui64(v153) < base.Ui64(v154) {
																									v156 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
																									v157 = F_bt_page_print_tuples(m, v156)
																									mBase = m.M
																									v158 = m.ExcPending
																									if v158 != 0 {
																										return int64(0)
																									} else {
																										v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)))
																										v160 = int32(1)
																										v161 = v159 + v160
																										*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)) = uint16(v161)
																										v163 = *(*int64)(unsafe.Add(mBase, uint32(v152)))
																										*(*int64)(unsafe.Add(mBase, uint32(v152))) = v163 + int64(1)
																										v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																										*(*int32)(unsafe.Add(mBase, uint32(v167)+20)) = v160
																										v178 = v157
																										m.G0 = v12 + int32(16)
																										return v178
																									}
																								} else {
																									F_end_MultiFuncCall(m, l0)
																									mBase = m.M
																									v171 = m.ExcPending
																									if v171 != 0 {
																										return int64(0)
																									} else {
																										v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																										*(*int32)(unsafe.Add(mBase, uint32(v172)+20)) = int32(2)
																										v175 = int32(1)
																										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v175)
																										v178 = int64(0)
																										m.G0 = v12 + int32(16)
																										return v178
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		} else {
																			*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = int64(0)
																			v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+12)))
																			v122 = v120
																			v124 = v122 & int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v52)+6)) = uint8(v124)
																			v126 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
																			v127 = int32(0)
																			*(*uint8)(unsafe.Add(mBase, uint32(v52)+7)) = uint8(base.B2i32(v126 == v127))
																			v133 = F_get_call_result_type(m, l0, v127, v12+int32(12))
																			mBase = m.M
																			v134 = m.ExcPending
																			if v134 != 0 {
																				return int64(0)
																			} else {
																				if v133 != int32(1) {
																					F_errstart_cold(m, int32(21), int32(0))
																					mBase = m.M
																					v203 = m.ExcPending
																					if v203 != 0 {
																						return int64(0)
																					} else {
																						F_errmsg_internal(m, int32(_a_F_bt_page_items_internal_2), int32(0))
																						mBase = m.M
																						v207 = m.ExcPending
																						if v207 != 0 {
																							return int64(0)
																						} else {
																							F_errfinish(m, int32(_a_F_bt_page_items_internal_3), int32(688), int32(_a_F_bt_page_items_internal_4))
																							mBase = m.M
																							v212 = m.ExcPending
																							if v212 != 0 {
																								return int64(0)
																							} else {
																								base.Wasm_trap_unreachable()
																								for {
																								}
																							}
																						}
																					}
																				} else {
																					v137 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
																					v138 = F_BlessTupleDesc(m, v137)
																					mBase = m.M
																					v139 = m.ExcPending
																					if v139 != 0 {
																						return int64(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v138
																						*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v52
																						*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_internal[0])) = v47
																						v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																						v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+16))
																						v153 = *(*int64)(unsafe.Add(mBase, uint32(v152)))
																						v154 = *(*int64)(unsafe.Add(mBase, uint32(v152)+8))
																						if base.Ui64(v153) < base.Ui64(v154) {
																							v156 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
																							v157 = F_bt_page_print_tuples(m, v156)
																							mBase = m.M
																							v158 = m.ExcPending
																							if v158 != 0 {
																								return int64(0)
																							} else {
																								v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)))
																								v160 = int32(1)
																								v161 = v159 + v160
																								*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)) = uint16(v161)
																								v163 = *(*int64)(unsafe.Add(mBase, uint32(v152)))
																								*(*int64)(unsafe.Add(mBase, uint32(v152))) = v163 + int64(1)
																								v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								*(*int32)(unsafe.Add(mBase, uint32(v167)+20)) = v160
																								v178 = v157
																								m.G0 = v12 + int32(16)
																								return v178
																							}
																						} else {
																							F_end_MultiFuncCall(m, l0)
																							mBase = m.M
																							v171 = m.ExcPending
																							if v171 != 0 {
																								return int64(0)
																							} else {
																								v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								*(*int32)(unsafe.Add(mBase, uint32(v172)+20)) = int32(2)
																								v175 = int32(1)
																								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v175)
																								v178 = int64(0)
																								m.G0 = v12 + int32(16)
																								return v178
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
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+16))
					v153 = *(*int64)(unsafe.Add(mBase, uint32(v152)))
					v154 = *(*int64)(unsafe.Add(mBase, uint32(v152)+8))
					if base.Ui64(v153) < base.Ui64(v154) {
						v156 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
						v157 = F_bt_page_print_tuples(m, v156)
						mBase = m.M
						v158 = m.ExcPending
						if v158 != 0 {
							return int64(0)
						} else {
							v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)))
							v160 = int32(1)
							v161 = v159 + v160
							*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)) = uint16(v161)
							v163 = *(*int64)(unsafe.Add(mBase, uint32(v152)))
							*(*int64)(unsafe.Add(mBase, uint32(v152))) = v163 + int64(1)
							v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v167)+20)) = v160
							v178 = v157
							m.G0 = v12 + int32(16)
							return v178
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int64(0)
						} else {
							v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v172)+20)) = int32(2)
							v175 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v175)
							v178 = int64(0)
							m.G0 = v12 + int32(16)
							return v178
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v187 = m.ExcPending
				if v187 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v190 = m.ExcPending
					if v190 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_bt_page_items_internal_6), int32(0))
						mBase = m.M
						v194 = m.ExcPending
						if v194 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_bt_page_items_internal_3), int32(636), int32(_a_F_bt_page_items_internal_4))
							mBase = m.M
							v199 = m.ExcPending
							if v199 != 0 {
								return int64(0)
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
}
