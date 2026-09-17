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
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v299+(v295^int32(-1))<<(uint(int32(6))%32))+16))
	v314 = v305
	goto L60
L62:
	;
	goto L63
L63:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _c_F__bt_binsrch_insert[3]))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v307+v295<<(uint(int32(6))%32)+int32(-64))+16))
	v314 = v313
	goto L60
L64:
	;
	F_errfinish(m, int32(_a_F__bt_binsrch_insert_7), int32(575), int32(_a_F__bt_binsrch_insert_8))
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
func F__bt_binsrch_skiparray_skey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+19)))
	if v9 != 0 {
		v65 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l6))) = v65
		return
	} else {
		if l3 != 0 {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+3)))
			if v12&int32(2) != 0 {
				v15 = int32(-1)
			} else {
				v15 = int32(1)
			}
			v65 = v15
			*(*int32)(unsafe.Add(mBase, uint32(l6))) = v65
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
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
						v39 = F_FunctionCall2Coll(m, v32+int32(16), v37, l2, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							if v39 != 0 {
							} else {
								v65 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l6))) = v65
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
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
							v39 = F_FunctionCall2Coll(m, v32+int32(16), v37, l2, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								if v39 != 0 {
								} else {
									v65 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(l6))) = v65
								}
								return
							}
						}
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
						v27 = F_FunctionCall2Coll(m, v20+int32(16), v25, l2, v26)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							if v27 == int32(0) {
								v65 = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(l6))) = v65
								return
							} else {
								v32 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
								if v32 == int32(0) {
									return
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
									v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+44))
									v39 = F_FunctionCall2Coll(m, v32+int32(16), v37, l2, v38)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										if v39 != 0 {
										} else {
											v65 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(l6))) = v65
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
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
					if v53 == int32(0) {
						return
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+44))
						v60 = F_FunctionCall2Coll(m, v53+int32(16), v58, l2, v59)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							if v60 != 0 {
							} else {
								v65 = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(l6))) = v65
							}
							return
						}
					}
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l4)+28))
					if v42 == int32(0) {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
						if v53 == int32(0) {
							return
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+44))
							v60 = F_FunctionCall2Coll(m, v53+int32(16), v58, l2, v59)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								if v60 != 0 {
								} else {
									v65 = int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(l6))) = v65
								}
								return
							}
						}
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)+44))
						v49 = F_FunctionCall2Coll(m, v42+int32(16), v47, l2, v48)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							if v49 != 0 {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(l4)+24))
								if v53 == int32(0) {
									return
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+44))
									v60 = F_FunctionCall2Coll(m, v53+int32(16), v58, l2, v59)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										if v60 != 0 {
										} else {
											v65 = int32(-1)
											*(*int32)(unsafe.Add(mBase, uint32(l6))) = v65
										}
										return
									}
								}
							} else {
								v65 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l6))) = v65
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
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
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
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
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
				v21 = v18 + v12<<(uint(int32(5))%32)
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
				v143 = v12 + int32(1)
				v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+20))
				if v143 < v145 {
					v12 = v143
					continue
				} else {
					break
				}
				break
			}
			v149 = v144
		} else {
			v149 = v7
		}
		v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
		switch v151 {
		case 0, 5:
			F_UnregisterSnapshot(m, v150)
			mBase = m.M
			v153 = m.ExcPending
			if v153 != 0 {
				return
			} else {
				v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v155 = v154
				F_DestroyParallelContext(m, v155)
				mBase = m.M
				v157 = m.ExcPending
				if v157 != 0 {
					return
				} else {
					v160 = *(*int32)(unsafe.Add(mBase, _c_F__bt_end_parallel[20]))
					v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+72))
					*(*int32)(unsafe.Add(mBase, uint32(v160)+72)) = v161 - int32(1)
					return
				}
			}
		default:
			v155 = v149
			F_DestroyParallelContext(m, v155)
			mBase = m.M
			v157 = m.ExcPending
			if v157 != 0 {
				return
			} else {
				v160 = *(*int32)(unsafe.Add(mBase, _c_F__bt_end_parallel[20]))
				v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+72))
				*(*int32)(unsafe.Add(mBase, uint32(v160)+72)) = v161 - int32(1)
				return
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v6 == int32(0) {
		v10 = F_ReadBuffer(m, l0, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			F_LockBuffer(m, v10, int32(1))
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
							F_LockBuffer(m, v10, int32(0))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								F_ReleaseBuffer(m, v10)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									return
								}
							}
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							v35 = F_MemoryContextAlloc(m, v33, int32(48))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v35
								v38 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v35)+40)) = v38
								v40 = *(*int64)(unsafe.Add(mBase, uint32(v17)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v35)+32)) = v40
								v42 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v35)+24)) = v42
								v44 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v44
								v46 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v46
								v48 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
								*(*int64)(unsafe.Add(mBase, uint32(v35))) = v48
								F_LockBuffer(m, v10, int32(0))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return
								} else {
									F_ReleaseBuffer(m, v10)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
										v57 = v55
										v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
										*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v59)))
										v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+40)))
										*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v63)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v57 = v6
		v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v59)))
		v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+40)))
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v63)
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
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70+(v24^int32(-1))<<(uint(int32(6))%32))+16))
	v85 = v76
	goto L12
L14:
	;
	goto L15
L15:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F__bt_moveright[3]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v78+v24<<(uint(int32(6))%32)+int32(-64))+16))
	v85 = v84
	goto L12
L16:
	;
	F__bt_unlockbuf(m, v24)
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
	F_LockBuffer(m, v24, int32(2))
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
	F__bt_relbuf(m, v24)
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
	F_errfinish(m, int32(_a_F__bt_moveright_1), int32(323), int32(_a_F__bt_moveright_2))
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
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
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v399 int32
	_ = v399
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v444 int32
	_ = v444
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v775 int32
	_ = v775
	var v788 int32
	_ = v788
	var v797 int32
	_ = v797
	var v818 int32
	_ = v818
	var v831 int32
	_ = v831
	var v841 int32
	_ = v841
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v880 int32
	_ = v880
	var v893 int32
	_ = v893
	v4 = l3
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+56))
	if v26 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+16)))
	if v26 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readpage[0]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v30+(v26^int32(-1))<<(uint(int32(2))%32))))
	v44 = v36
	goto L1
L3:
	;
	goto L4
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readpage[1]))
	v44 = v38 + v26<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+60)) = v64
	v66 = v45 + v44
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+64)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+84)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+80)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v25)+68)) = v69
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v74 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readpage[2]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49+(v26^int32(-1))<<(uint(int32(6))%32))+16))
	v64 = v55
	goto L5
L7:
	;
	goto L8
L8:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readpage[3]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57+v26<<(uint(int32(6))%32)+int32(-64))+16))
	v64 = v63
	goto L5
L9:
	;
	if l1 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v25)+60))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_PredicateLockPage(m, v24, v82, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L15
	} else {
		goto L17
	}
L12:
	;
	v77 = v69
	goto L14
L13:
	;
	v77 = v67
	goto L14
L14:
	;
	F__bt_parallel_release(m, l0, v77, v64)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	goto L11
L17:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v24)+192))
	v88 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87)+8)))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+12)))
	v91 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+30)) = uint16(v91)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+26)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+13)) = uint8(v91)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+12)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v91
	if base.Ui32(int32(25)) <= base.Ui32(v90) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v110 = int32(base.Ui32(v90+int32(_a_F__bt_readpage_0)) >> (uint(int32(2)) % 32))
	goto L20
L19:
	;
	v110 = v91
	goto L20
L20:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+2)) = uint16(v110)
	v112 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+24)) = uint8(v112)
	if v89 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v116 = int32(2)
	goto L23
L22:
	;
	v116 = v112
	goto L23
L23:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v22))) = uint16(v116)
	if l1 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	m.G0 = v22 + int32(32)
	return v893
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+100)) = v880
	*(*int32)(unsafe.Add(mBase, uint32(v25)+96)) = v864
	*(*int32)(unsafe.Add(mBase, uint32(v25)+92)) = v867
	v893 = base.B2i32(v867 <= v864)
	goto L24
L26:
	;
	if v86 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v86 != 0 {
		goto L102
	} else {
		goto L103
	}
L29:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v120 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	if v4|base.B2i32(base.Ui32(v110&int32(_a_F__bt_readpage_1)) <= base.Ui32(v116)) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L32:
	;
	v146 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+18)) = uint16(v146)
	goto L31
L33:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	v126 = v44 + v123&int32(_a_F__bt_readpage_2)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v126
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+18)))
	if v128 != int32(1) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v132 = F__bt_scanbehind_checkkeys(m, l0, int32(1), v126)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L15
	} else {
		goto L35
	}
L35:
	;
	if v132 != 0 {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v134 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+17)) = uint8(v134)
	v136 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+89)) = uint8(v136)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v139 == v136 {
		v893 = v136
		goto L24
	} else {
		goto L37
	}
L37:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v25)+60))
	F__bt_parallel_primscan_schedule(m, l0, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L15
	} else {
		goto L38
	}
L38:
	;
	v893 = v136
	goto L24
L39:
	;
	F__bt_set_startikey(m, l0, v22)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L15
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(v116) < base.Ui32(l2) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	v159 = l2
	goto L45
L44:
	;
	v159 = v116
	goto L45
L45:
	;
	v161 = v110 & int32(_a_F__bt_readpage_1)
	if base.Ui32(v159) <= base.Ui32(v161) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v164 = v25 + int32(104)
	v170 = int32(0)
	v176 = v159
	goto L49
L47:
	;
	v444 = int32(-1)
	goto L48
L48:
	;
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+24)))
	if v460 != int32(1) {
		goto L88
	} else {
		goto L89
	}
L49:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v44+int32(20)+v176&int32(_a_F__bt_readpage_1)<<(uint(int32(2))%32))))
	v196 = int32(_a_F__bt_readpage_3)
	if base.B2i32(v187 == int32(1))&base.B2i32(v195&v196 == v196) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v444 = v422 - int32(1)
	goto L48
L51:
	;
	goto L50
L52:
	;
	if base.Ui32(v416&int32(_a_F__bt_readpage_1)) <= base.Ui32(v161) {
		v170 = v399
		v176 = v416
		goto L49
	} else {
		goto L87
	}
L53:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+20)) = uint16(v176)
	v208 = v44 + v195&int32(_a_F__bt_readpage_2)
	v209 = F__bt_checkkeys(m, l0, v22, base.B2i32(v86 != int32(0)), v208, v88)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L15
	} else {
		goto L56
	}
L54:
	;
	v378 = v170
	goto L55
L55:
	;
	v399 = v378
	v416 = v176 + int32(1)
	goto L52
L56:
	;
	if v86 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if v209 == int32(0) {
		v356 = v170
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+22)))
	if base.Ui32(int32(2047)) < base.Ui32((v213-int32(1))&int32(_a_F__bt_readpage_1)) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v220 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+22)) = uint16(v220)
	v399 = v170
	v416 = v213
	goto L52
L60:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+24)))
	if v373 != int32(1) {
		v422 = v356
		goto L51
	} else {
		goto L86
	}
L61:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+7)))
	if v225&int32(32) != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v260 = v164 + v170*int32(10)
	v261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v208)+2)))
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v208))))
	v266 = v261 + (v208 + v262<<(uint(int32(16))%32))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v267
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v266)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v260)+6)) = uint16(v176)
	*(*uint16)(unsafe.Add(mBase, uint32(v260)+4)) = uint16(v269)
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	if v272 != 0 {
		goto L73
	} else {
		goto L74
	}
L63:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+5)))
	if v228&int32(32) != 0 {
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v233 = v164 + v170*int32(10)
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v208)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v233)+4)) = uint16(v234)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = v236
	*(*uint16)(unsafe.Add(mBase, uint32(v233)+6)) = uint16(v176)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	if v239 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L65
L67:
	;
	v240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v208)+6)))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v25)+84))
	*(*uint16)(unsafe.Add(mBase, uint32(v233)+8)) = uint16(v241)
	v244 = v240 & int32(_a_F__bt_readpage_4)
	if v244 != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	v356 = v170 + int32(1)
	goto L60
L70:
	;
	base.MemoryCopy(m, v239+v241, v208, v244)
	goto L72
L71:
	;
	goto L72
L72:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v25)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+84)) = v247 + (v244+int32(7))&int32(_a_F__bt_readpage_5)
	goto L69
L73:
	;
	v273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v208)+2)))
	v274 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v208))))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v25)+84))
	*(*uint16)(unsafe.Add(mBase, uint32(v260)+8)) = uint16(v275)
	v277 = v272 + v275
	v284 = (v273 | v274<<(uint(int32(16))%32) + int32(7)) & int32(-8)
	if v284 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v299 = int32(0)
	goto L75
L75:
	;
	v300 = int32(1)
	v302 = v170 + v300
	v303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v208)+4)))
	if v303&int32(4094) == int32(0) {
		v356 = v302
		goto L60
	} else {
		goto L79
	}
L76:
	;
	base.MemoryCopy(m, v277, v208, v284)
	goto L78
L77:
	;
	goto L78
L78:
	;
	v286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277)+6)))
	v289 = v286&int32(_a_F__bt_readpage_6) | v284
	*(*uint16)(unsafe.Add(mBase, uint32(v277)+6)) = uint16(v289)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v25)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+84)) = v291 + v284
	v294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v260)+8)))
	v299 = v294
	goto L75
L79:
	;
	v310 = v302
	v314 = v300
	goto L80
L80:
	;
	v329 = v164 + v310*int32(10)
	v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v208)+2)))
	v331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v208))))
	v338 = v330 + (v208 + v331<<(uint(int32(16))%32)) + v314*int32(6)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
	*(*int32)(unsafe.Add(mBase, uint32(v329))) = v339
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v338)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v329)+6)) = uint16(v176)
	*(*uint16)(unsafe.Add(mBase, uint32(v329)+4)) = uint16(v341)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	if v344 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v356 = v347
	goto L60
L82:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v329)+8)) = uint16(v299)
	goto L84
L83:
	;
	goto L84
L84:
	;
	v346 = int32(1)
	v347 = v310 + v346
	v349 = v314 + v346
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v208)+4)))
	if base.Ui32(v349) < base.Ui32(v350&int32(4095)) {
		v310 = v347
		v314 = v349
		goto L80
	} else {
		goto L85
	}
L85:
	;
	goto L81
L86:
	;
	v378 = v356
	goto L55
L87:
	;
	v422 = v399
	goto L51
L88:
	;
	v510 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+89)) = uint8(v510)
	v864 = v444
	v867 = v510
	v880 = v510
	goto L25
L89:
	;
	v463 = int32(0)
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+18)))
	if v465 != 0 {
		v864 = v444
		v867 = v463
		v880 = v463
		goto L25
	} else {
		goto L90
	}
L90:
	;
	v466 = int32(0)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	if v467 == v466 {
		v864 = v444
		v867 = v463
		v880 = v466
		goto L25
	} else {
		goto L91
	}
L91:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	v473 = v44 + v470&int32(_a_F__bt_readpage_2)
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+13)))
	if v474 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	F__bt_start_array_keys(m, l0, int32(1))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L15
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v480 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v480
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+13)) = uint8(v480)
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v473)+7)))
	if v486&int32(32) == v480 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L94
L96:
	;
	v501 = F__bt_checkkeys(m, l0, v22, base.B2i32(v86 != v480), v473, v500)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L15
	} else {
		goto L100
	}
L97:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v24)+192))
	v498 = int32(*(*int16)(unsafe.Add(mBase, uint32(v497)+8)))
	v500 = v498
	goto L96
L98:
	;
	v491 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v473)+4)))
	if v491&int32(_a_F__bt_readpage_7) != 0 {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v500 = v491 & int32(4095)
	goto L96
L100:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+24)))
	if v503 == int32(0) {
		goto L88
	} else {
		goto L101
	}
L101:
	;
	v864 = v444
	v867 = v463
	v880 = int32(0)
	goto L25
L102:
	;
	if base.Ui32(v110&int32(_a_F__bt_readpage_1)) < base.Ui32(v116) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	if v4|base.B2i32(base.Ui32(v110&int32(_a_F__bt_readpage_1)) <= base.Ui32(v116)) == int32(0) {
		goto L113
	} else {
		goto L114
	}
L105:
	;
	v545 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+18)) = uint16(v545)
	goto L104
L106:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v517 == int32(0) {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v44+v116<<(uint(int32(2))%32))+20))
	v526 = v44 + v523&int32(_a_F__bt_readpage_2)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v526
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+18)))
	if v528 != int32(1) {
		goto L105
	} else {
		goto L108
	}
L108:
	;
	v531 = F__bt_scanbehind_checkkeys(m, l0, l1, v526)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L15
	} else {
		goto L109
	}
L109:
	;
	if v531 != 0 {
		goto L105
	} else {
		goto L110
	}
L110:
	;
	v533 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+17)) = uint8(v533)
	v535 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+88)) = uint8(v535)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v538 == v535 {
		v893 = v535
		goto L24
	} else {
		goto L111
	}
L111:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v25)+60))
	F__bt_parallel_primscan_schedule(m, l0, v541)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L15
	} else {
		goto L112
	}
L112:
	;
	v893 = v535
	goto L24
L113:
	;
	F__bt_set_startikey(m, l0, v22)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L15
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v556 = int32(1358)
	v558 = v110 & int32(_a_F__bt_readpage_1)
	if base.Ui32(l2) < base.Ui32(v558) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L115
L117:
	;
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+24)))
	if v855 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L118:
	;
	v560 = l2
	goto L120
L119:
	;
	v560 = v558
	goto L120
L120:
	;
	if base.Ui32(v560) < base.Ui32(v116) {
		v841 = v556
		goto L117
	} else {
		goto L121
	}
L121:
	;
	v563 = v25 + int32(104)
	v572 = v556
	v574 = v560
	goto L122
L122:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	v587 = v574 & int32(_a_F__bt_readpage_1)
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v44+int32(20)+v587<<(uint(int32(2))%32))))
	v592 = int32(_a_F__bt_readpage_3)
	v596 = v585 & base.B2i32(v591&v592 == v592)
	if v596&base.B2i32(base.Ui32(v116) < base.Ui32(v587)) == int32(0) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	v841 = v818
	goto L117
L124:
	;
	if base.Ui32(v116) <= base.Ui32(v831&int32(_a_F__bt_readpage_1)) {
		v572 = v818
		v574 = v831
		goto L122
	} else {
		goto L165
	}
L125:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+20)) = uint16(v574)
	v604 = v44 + v591&int32(_a_F__bt_readpage_2)
	if v86 != 0 {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	v797 = v572
	goto L127
L127:
	;
	v818 = v797
	v831 = v574 - int32(1)
	goto L124
L128:
	;
	v635 = int32(0)
	if v633&base.B2i32(v596 == v635) == v635 {
		v775 = v572
		goto L140
	} else {
		goto L141
	}
L129:
	;
	if v587 != v116 {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	goto L131
L131:
	;
	v631 = F__bt_checkkeys(m, l0, v22, int32(0), v604, v88)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L15
	} else {
		goto L139
	}
L132:
	;
	v618 = F__bt_checkkeys(m, l0, v22, int32(1), v604, v88)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L15
	} else {
		goto L136
	}
L133:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+13)))
	if v606&int32(1) == int32(0) {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v611 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v611
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+13)) = uint8(v611)
	F__bt_start_array_keys(m, l0, l1)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L15
	} else {
		goto L135
	}
L135:
	;
	goto L132
L136:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+18)))
	if v620 != 0 {
		v841 = v572
		goto L117
	} else {
		goto L137
	}
L137:
	;
	v621 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+22)))
	if base.Ui32(int32(2047)) < base.Ui32((v621-int32(1))&int32(_a_F__bt_readpage_1)) {
		v633 = v618
		goto L128
	} else {
		goto L138
	}
L138:
	;
	v628 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+22)) = uint16(v628)
	v818 = v572
	v831 = v621
	goto L124
L139:
	;
	v633 = v631
	goto L128
L140:
	;
	v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+24)))
	if v788 != int32(1) {
		v841 = v775
		goto L117
	} else {
		goto L164
	}
L141:
	;
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604)+7)))
	if v640&int32(32) != 0 {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v674 = v572 - int32(1)
	v677 = v563 + v674*int32(10)
	v678 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604)+2)))
	v679 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604))))
	v683 = v678 + (v604 + v679<<(uint(int32(16))%32))
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v683)))
	*(*int32)(unsafe.Add(mBase, uint32(v677))) = v684
	v686 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v683)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v677)+6)) = uint16(v574)
	*(*uint16)(unsafe.Add(mBase, uint32(v677)+4)) = uint16(v686)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	if v689 != 0 {
		goto L151
	} else {
		goto L152
	}
L143:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604)+5)))
	if v643&int32(32) != 0 {
		goto L142
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v647 = v572 - int32(1)
	v650 = v563 + v647*int32(10)
	v651 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v650)+4)) = uint16(v651)
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v604)))
	*(*int32)(unsafe.Add(mBase, uint32(v650))) = v653
	*(*uint16)(unsafe.Add(mBase, uint32(v650)+6)) = uint16(v574)
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	if v656 == int32(0) {
		v775 = v647
		goto L140
	} else {
		goto L147
	}
L146:
	;
	goto L145
L147:
	;
	v659 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604)+6)))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v25)+84))
	*(*uint16)(unsafe.Add(mBase, uint32(v650)+8)) = uint16(v660)
	v663 = v659 & int32(_a_F__bt_readpage_4)
	if v663 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	base.MemoryCopy(m, v656+v660, v604, v663)
	goto L150
L149:
	;
	goto L150
L150:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v25)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+84)) = v666 + (v663+int32(7))&int32(_a_F__bt_readpage_5)
	v775 = v647
	goto L140
L151:
	;
	v690 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604)+2)))
	v691 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604))))
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v25)+84))
	*(*uint16)(unsafe.Add(mBase, uint32(v677)+8)) = uint16(v692)
	v694 = v689 + v692
	v701 = (v690 | v691<<(uint(int32(16))%32) + int32(7)) & int32(-8)
	if v701 != 0 {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	v716 = int32(0)
	goto L153
L153:
	;
	v718 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604)+4)))
	if v718&int32(4094) == int32(0) {
		v775 = v674
		goto L140
	} else {
		goto L157
	}
L154:
	;
	base.MemoryCopy(m, v694, v604, v701)
	goto L156
L155:
	;
	goto L156
L156:
	;
	v703 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v694)+6)))
	v706 = v703&int32(_a_F__bt_readpage_6) | v701
	*(*uint16)(unsafe.Add(mBase, uint32(v694)+6)) = uint16(v706)
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v25)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+84)) = v708 + v701
	v711 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v677)+8)))
	v716 = v711
	goto L153
L157:
	;
	v725 = int32(1)
	v729 = v674
	goto L158
L158:
	;
	v743 = v729 - int32(1)
	v746 = v563 + v743*int32(10)
	v747 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604)+2)))
	v748 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604))))
	v755 = v747 + (v604 + v748<<(uint(int32(16))%32)) + v725*int32(6)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v755)))
	*(*int32)(unsafe.Add(mBase, uint32(v746))) = v756
	v758 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v755)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v746)+6)) = uint16(v574)
	*(*uint16)(unsafe.Add(mBase, uint32(v746)+4)) = uint16(v758)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	if v761 != 0 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v775 = v743
	goto L140
L160:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v746)+8)) = uint16(v716)
	goto L162
L161:
	;
	goto L162
L162:
	;
	v764 = v725 + int32(1)
	v765 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v604)+4)))
	if base.Ui32(v764) < base.Ui32(v765&int32(4095)) {
		v725 = v764
		v729 = v743
		goto L158
	} else {
		goto L163
	}
L163:
	;
	goto L159
L164:
	;
	v797 = v775
	goto L127
L165:
	;
	goto L123
L166:
	;
	v858 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+88)) = uint8(v858)
	goto L168
L167:
	;
	goto L168
L168:
	;
	v864 = int32(1357)
	v867 = v841
	v880 = int32(1357)
	goto L25
}
func F_bt_index_parent_check(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v39 int32
	_ = v39
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(1)
	v12 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v12 < int32(2) {
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+13)) = uint8(base.B2i32(v15 != int32(0)))
		if v12 == int32(2) {
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(base.B2i32(v21 != int32(0)))
			if base.Ui32(v12) < base.Ui32(int32(4)) {
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(base.B2i32(v27 != int32(0)))
			}
		}
	}
	F_amcheck_lock_relation_and_check(m, v9, int32(403), int32(_a_F_bt_index_parent_check_0), int32(5), v7+int32(12))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return int32(0)
	} else {
		m.G0 = v7 + int32(16)
		return int32(0)
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v267 int32
	_ = v267
	var v288 int32
	_ = v288
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
	var v296 int32
	_ = v296
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	v3 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(224)
	m.G0 = v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
	if v20&int32(64) == v3 {
		v267 = l1
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L20
	} else {
		goto L60
	}
L2:
	;
	m.G0 = v18 + int32(224)
	return v267
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+52))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v27 <= int32(0) {
		v267 = l1
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
	v59 = v26 + v37<<(uint(int32(4))%32) + v40*int32(100)
	v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59)+94)))
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
	if v190&int32(1) == int32(0) {
		v267 = l1
		goto L2
	} else {
		goto L48
	}
L7:
	;
	v134 = v18 + int32(96) + v40<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v128
	v137 = v59 + int32(20)
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+82)))
	if v138 != 0 {
		goto L27
	} else {
		goto L28
	}
L8:
	;
	v125 = F_nocache_index_getattr(m, l1, v60, v26)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L20
	} else {
		goto L25
	}
L9:
	;
	v71 = v26 + int32(4) + v60<<(uint(int32(4))%32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
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
	v107 = int32(1)
	v108 = v60 - v107
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v108>>(uint(int32(3))%32)))))
	if int32(base.Ui32(v112)>>(uint(v108&int32(7))%32))&v107 != 0 {
		goto L8
	} else {
		goto L24
	}
L12:
	;
	v75 = v72 + v31
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+6)))
	if v76 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
	switch v79 - int32(1) {
	case 0:
		goto L19
	case 1:
		goto L18
	default:
		goto L16
	case 3:
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v127 = int32(0)
	v128 = v75
	goto L7
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v127 = int32(0)
	v128 = v86
	goto L7
L18:
	;
	v84 = int32(*(*int16)(unsafe.Add(mBase, uint32(v75))))
	v127 = int32(0)
	v128 = v84
	goto L7
L19:
	;
	v82 = int32(*(*int8)(unsafe.Add(mBase, uint32(v75))))
	v127 = int32(0)
	v128 = v82
	goto L7
L20:
	;
	return int32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = base.I32_extend16_s(v79)
	F_errmsg_internal(m, int32(_a_F_bt_normalize_tuple_0), v18+int32(16))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_bt_normalize_tuple_1), int32(70), int32(_a_F_bt_normalize_tuple_2))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v118 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v118)
	v127 = v118
	v128 = int32(0)
	goto L7
L25:
	;
	v127 = int32(0)
	v128 = v125
	goto L7
L26:
	;
	v196 = v40 + int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v196 < v197 {
		v36 = v190
		v37 = v197
		v40 = v196
		goto L5
	} else {
		goto L47
	}
L27:
	;
	v190 = v36
	goto L26
L28:
	;
	goto L29
L29:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137)+72)))
	if base.B2i32(v139 != int32(_a_F_bt_normalize_tuple_3))|v127 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v190 = v36
	goto L26
L31:
	;
	goto L32
L32:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v143 == int32(1) {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v147 = v143 & int32(3)
	if v147 != int32(2) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v182
	v187 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v51))) = uint8(v187)
	v190 = v187
	goto L26
L35:
	;
	if v147 != 0 {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	if base.Ui32(v150) < base.Ui32(int32(2044)) {
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v157 = F_pg_detoast_datum(m, v128)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L20
	} else {
		goto L40
	}
L39:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+84)))
	switch v154 - int32(109) {
	case 0, 11:
		v190 = int32(1)
		goto L26
	default:
		goto L35
	}
L40:
	;
	v182 = v157
	goto L34
L41:
	;
	v190 = v36
	goto L26
L42:
	;
	goto L43
L43:
	;
	v161 = int32(base.Ui32(v150) >> (uint(int32(2)) % 32))
	v163 = v161 - int32(3)
	if base.Ui32(int32(127)) < base.Ui32(v163) {
		v190 = v36
		goto L26
	} else {
		goto L44
	}
L44:
	;
	v166 = F_palloc(m, v163)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L20
	} else {
		goto L45
	}
L45:
	;
	v168 = int32(1)
	v171 = v163<<(uint(v168)%32) | v168
	*(*uint8)(unsafe.Add(mBase, uint32(v166))) = uint8(v171)
	v174 = v161 - int32(4)
	if v174 == int32(0) {
		v182 = v166
		goto L34
	} else {
		goto L46
	}
L46:
	;
	base.MemoryCopy(m, v166+int32(1), v128+int32(4), v174)
	v182 = v166
	goto L34
L47:
	;
	goto L6
L48:
	;
	v207 = F_index_form_tuple(m, v26, v18+int32(96), v18-int32(-64))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L20
	} else {
		goto L49
	}
L49:
	;
	v209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v207)+4)) = uint16(v209)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v207))) = v211
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if int32(0) < v213 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v220 = v213
	v223 = int32(0)
	goto L53
L51:
	;
	goto L52
L52:
	;
	v267 = v207
	goto L2
L53:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+int32(32)+v223))))
	if v235 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L52
L55:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(96)+v223<<(uint(int32(2))%32))))
	F_pfree(m, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L20
	} else {
		goto L58
	}
L56:
	;
	v247 = v220
	goto L57
L57:
	;
	v249 = v223 + int32(1)
	if v249 < v247 {
		v220 = v247
		v223 = v249
		goto L53
	} else {
		goto L59
	}
L58:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v247 = v246
	goto L57
L59:
	;
	goto L54
L60:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L20
	} else {
		goto L61
	}
L61:
	;
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+48))
	v296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v296
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v295 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v292 | v293<<(uint(int32(16))%32)
	F_errmsg(m, int32(_a_F_bt_normalize_tuple_4), v18)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L20
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_bt_normalize_tuple_5), int32(2892), int32(_a_F_bt_normalize_tuple_6))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L20
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_bt_page_items_1_9(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_bt_page_items_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_bt_page_items_bytea(m *base.Module, l0 int32) int32 {
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
	var v148 int32
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
	var v175 int32
	_ = v175
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = F_superuser(m)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v18 != 0 {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
				if v21 == int32(0) {
					v24 = F_init_MultiFuncCall(m, l0)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = int32(_a_F_bt_page_items_bytea_0)
						v27 = *(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0]))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
						*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v29
						v32 = F_palloc(m, int32(12))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = F_get_page_from_raw(m, v14)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v32))) = v34
								v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+14)))
								if v37 == int32(0) {
									*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_bytea[0])) = v27
									v172 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
									v175 = int32(0)
									m.G0 = v11 + int32(32)
									return v175
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
											return int32(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v207 = m.ExcPending
											if v207 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_bt_page_items_bytea_2)
												F_errmsg(m, int32(_a_F_bt_page_items_bytea_3), v11+int32(16))
												mBase = m.M
												v214 = m.ExcPending
												if v214 != 0 {
													return int32(0)
												} else {
													v215 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
													v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215)+16)))
													v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+19)))
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(16)
													*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = (v217<<(uint(int32(8))%32) - v216) & int32(_a_F_bt_page_items_bytea_1)
													F_errdetail(m, int32(_a_F_bt_page_items_bytea_4), v11)
													mBase = m.M
													v228 = m.ExcPending
													if v228 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(773), int32(_a_F_bt_page_items_bytea_6))
														mBase = m.M
														v233 = m.ExcPending
														if v233 != 0 {
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
										v53 = v34 + v47
										v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
										if v54&int32(8) != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v237 = m.ExcPending
											if v237 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v240 = m.ExcPending
												if v240 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_bt_page_items_bytea_7), int32(0))
													mBase = m.M
													v244 = m.ExcPending
													if v244 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(780), int32(_a_F_bt_page_items_bytea_6))
														mBase = m.M
														v249 = m.ExcPending
														if v249 != 0 {
															return int32(0)
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
														return int32(0)
													} else {
														F_errcode(m, int32(50856066))
														mBase = m.M
														v256 = m.ExcPending
														if v256 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_bt_page_items_bytea_8), int32(0))
															mBase = m.M
															v260 = m.ExcPending
															if v260 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(785), int32(_a_F_bt_page_items_bytea_6))
																mBase = m.M
																v265 = m.ExcPending
																if v265 != 0 {
																	return int32(0)
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
																return int32(0)
															} else {
																if v125 != int32(1) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v269 = m.ExcPending
																	if v269 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																		mBase = m.M
																		v273 = m.ExcPending
																		if v273 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																			mBase = m.M
																			v278 = m.ExcPending
																			if v278 != 0 {
																				return int32(0)
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
																		return int32(0)
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
																				return int32(0)
																			} else {
																				v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																				v151 = int32(1)
																				v152 = v150 + v151
																				*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																				v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																				*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																				v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																				*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																				v175 = v148
																				m.G0 = v11 + int32(32)
																				return v175
																			}
																		} else {
																			F_end_MultiFuncCall(m, l0)
																			mBase = m.M
																			v162 = m.ExcPending
																			if v162 != 0 {
																				return int32(0)
																			} else {
																				v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																				*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																				v172 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																				v175 = int32(0)
																				m.G0 = v11 + int32(32)
																				return v175
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
																return int32(0)
															} else {
																if v99 != 0 {
																	F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_10), int32(0))
																	mBase = m.M
																	v104 = m.ExcPending
																	if v104 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(795), int32(_a_F_bt_page_items_bytea_6))
																		mBase = m.M
																		v109 = m.ExcPending
																		if v109 != 0 {
																			return int32(0)
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
																				return int32(0)
																			} else {
																				if v125 != int32(1) {
																					F_errstart_cold(m, int32(21), int32(0))
																					mBase = m.M
																					v269 = m.ExcPending
																					if v269 != 0 {
																						return int32(0)
																					} else {
																						F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																						mBase = m.M
																						v273 = m.ExcPending
																						if v273 != 0 {
																							return int32(0)
																						} else {
																							F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																							mBase = m.M
																							v278 = m.ExcPending
																							if v278 != 0 {
																								return int32(0)
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
																						return int32(0)
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
																								return int32(0)
																							} else {
																								v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																								v151 = int32(1)
																								v152 = v150 + v151
																								*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																								v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																								*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																								v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																								v175 = v148
																								m.G0 = v11 + int32(32)
																								return v175
																							}
																						} else {
																							F_end_MultiFuncCall(m, l0)
																							mBase = m.M
																							v162 = m.ExcPending
																							if v162 != 0 {
																								return int32(0)
																							} else {
																								v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																								v172 = int32(1)
																								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																								v175 = int32(0)
																								m.G0 = v11 + int32(32)
																								return v175
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
																		return int32(0)
																	} else {
																		if v125 != int32(1) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v269 = m.ExcPending
																			if v269 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																				mBase = m.M
																				v273 = m.ExcPending
																				if v273 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																					mBase = m.M
																					v278 = m.ExcPending
																					if v278 != 0 {
																						return int32(0)
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
																				return int32(0)
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
																						return int32(0)
																					} else {
																						v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																						v151 = int32(1)
																						v152 = v150 + v151
																						*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																						v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																						*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																						v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																						*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																						v175 = v148
																						m.G0 = v11 + int32(32)
																						return v175
																					}
																				} else {
																					F_end_MultiFuncCall(m, l0)
																					mBase = m.M
																					v162 = m.ExcPending
																					if v162 != 0 {
																						return int32(0)
																					} else {
																						v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																						*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																						v172 = int32(1)
																						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																						v175 = int32(0)
																						m.G0 = v11 + int32(32)
																						return v175
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
															return int32(0)
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
																		return int32(0)
																	} else {
																		if v125 != int32(1) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v269 = m.ExcPending
																			if v269 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																				mBase = m.M
																				v273 = m.ExcPending
																				if v273 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																					mBase = m.M
																					v278 = m.ExcPending
																					if v278 != 0 {
																						return int32(0)
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
																				return int32(0)
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
																						return int32(0)
																					} else {
																						v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																						v151 = int32(1)
																						v152 = v150 + v151
																						*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																						v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																						*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																						v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																						*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																						v175 = v148
																						m.G0 = v11 + int32(32)
																						return v175
																					}
																				} else {
																					F_end_MultiFuncCall(m, l0)
																					mBase = m.M
																					v162 = m.ExcPending
																					if v162 != 0 {
																						return int32(0)
																					} else {
																						v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																						*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																						v172 = int32(1)
																						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																						v175 = int32(0)
																						m.G0 = v11 + int32(32)
																						return v175
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
																		return int32(0)
																	} else {
																		if v99 != 0 {
																			F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_10), int32(0))
																			mBase = m.M
																			v104 = m.ExcPending
																			if v104 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(795), int32(_a_F_bt_page_items_bytea_6))
																				mBase = m.M
																				v109 = m.ExcPending
																				if v109 != 0 {
																					return int32(0)
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
																						return int32(0)
																					} else {
																						if v125 != int32(1) {
																							F_errstart_cold(m, int32(21), int32(0))
																							mBase = m.M
																							v269 = m.ExcPending
																							if v269 != 0 {
																								return int32(0)
																							} else {
																								F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																								mBase = m.M
																								v273 = m.ExcPending
																								if v273 != 0 {
																									return int32(0)
																								} else {
																									F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																									mBase = m.M
																									v278 = m.ExcPending
																									if v278 != 0 {
																										return int32(0)
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
																								return int32(0)
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
																										return int32(0)
																									} else {
																										v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																										v151 = int32(1)
																										v152 = v150 + v151
																										*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																										v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																										*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																										v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																										*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																										v175 = v148
																										m.G0 = v11 + int32(32)
																										return v175
																									}
																								} else {
																									F_end_MultiFuncCall(m, l0)
																									mBase = m.M
																									v162 = m.ExcPending
																									if v162 != 0 {
																										return int32(0)
																									} else {
																										v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																										*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																										v172 = int32(1)
																										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																										v175 = int32(0)
																										m.G0 = v11 + int32(32)
																										return v175
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
																				return int32(0)
																			} else {
																				if v125 != int32(1) {
																					F_errstart_cold(m, int32(21), int32(0))
																					mBase = m.M
																					v269 = m.ExcPending
																					if v269 != 0 {
																						return int32(0)
																					} else {
																						F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																						mBase = m.M
																						v273 = m.ExcPending
																						if v273 != 0 {
																							return int32(0)
																						} else {
																							F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																							mBase = m.M
																							v278 = m.ExcPending
																							if v278 != 0 {
																								return int32(0)
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
																						return int32(0)
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
																								return int32(0)
																							} else {
																								v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																								v151 = int32(1)
																								v152 = v150 + v151
																								*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																								v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																								*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																								v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																								v175 = v148
																								m.G0 = v11 + int32(32)
																								return v175
																							}
																						} else {
																							F_end_MultiFuncCall(m, l0)
																							mBase = m.M
																							v162 = m.ExcPending
																							if v162 != 0 {
																								return int32(0)
																							} else {
																								v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																								v172 = int32(1)
																								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																								v175 = int32(0)
																								m.G0 = v11 + int32(32)
																								return v175
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
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(788), int32(_a_F_bt_page_items_bytea_6))
																	mBase = m.M
																	v78 = m.ExcPending
																	if v78 != 0 {
																		return int32(0)
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
																				return int32(0)
																			} else {
																				if v125 != int32(1) {
																					F_errstart_cold(m, int32(21), int32(0))
																					mBase = m.M
																					v269 = m.ExcPending
																					if v269 != 0 {
																						return int32(0)
																					} else {
																						F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																						mBase = m.M
																						v273 = m.ExcPending
																						if v273 != 0 {
																							return int32(0)
																						} else {
																							F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																							mBase = m.M
																							v278 = m.ExcPending
																							if v278 != 0 {
																								return int32(0)
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
																						return int32(0)
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
																								return int32(0)
																							} else {
																								v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																								v151 = int32(1)
																								v152 = v150 + v151
																								*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																								v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																								*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																								v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																								v175 = v148
																								m.G0 = v11 + int32(32)
																								return v175
																							}
																						} else {
																							F_end_MultiFuncCall(m, l0)
																							mBase = m.M
																							v162 = m.ExcPending
																							if v162 != 0 {
																								return int32(0)
																							} else {
																								v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																								v172 = int32(1)
																								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																								v175 = int32(0)
																								m.G0 = v11 + int32(32)
																								return v175
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
																				return int32(0)
																			} else {
																				if v99 != 0 {
																					F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_10), int32(0))
																					mBase = m.M
																					v104 = m.ExcPending
																					if v104 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(795), int32(_a_F_bt_page_items_bytea_6))
																						mBase = m.M
																						v109 = m.ExcPending
																						if v109 != 0 {
																							return int32(0)
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
																								return int32(0)
																							} else {
																								if v125 != int32(1) {
																									F_errstart_cold(m, int32(21), int32(0))
																									mBase = m.M
																									v269 = m.ExcPending
																									if v269 != 0 {
																										return int32(0)
																									} else {
																										F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																										mBase = m.M
																										v273 = m.ExcPending
																										if v273 != 0 {
																											return int32(0)
																										} else {
																											F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																											mBase = m.M
																											v278 = m.ExcPending
																											if v278 != 0 {
																												return int32(0)
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
																										return int32(0)
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
																												return int32(0)
																											} else {
																												v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																												v151 = int32(1)
																												v152 = v150 + v151
																												*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																												v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																												*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																												v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																												*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																												v175 = v148
																												m.G0 = v11 + int32(32)
																												return v175
																											}
																										} else {
																											F_end_MultiFuncCall(m, l0)
																											mBase = m.M
																											v162 = m.ExcPending
																											if v162 != 0 {
																												return int32(0)
																											} else {
																												v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																												*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																												v172 = int32(1)
																												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																												v175 = int32(0)
																												m.G0 = v11 + int32(32)
																												return v175
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
																						return int32(0)
																					} else {
																						if v125 != int32(1) {
																							F_errstart_cold(m, int32(21), int32(0))
																							mBase = m.M
																							v269 = m.ExcPending
																							if v269 != 0 {
																								return int32(0)
																							} else {
																								F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																								mBase = m.M
																								v273 = m.ExcPending
																								if v273 != 0 {
																									return int32(0)
																								} else {
																									F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																									mBase = m.M
																									v278 = m.ExcPending
																									if v278 != 0 {
																										return int32(0)
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
																								return int32(0)
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
																										return int32(0)
																									} else {
																										v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																										v151 = int32(1)
																										v152 = v150 + v151
																										*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																										v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																										*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																										v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																										*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																										v175 = v148
																										m.G0 = v11 + int32(32)
																										return v175
																									}
																								} else {
																									F_end_MultiFuncCall(m, l0)
																									mBase = m.M
																									v162 = m.ExcPending
																									if v162 != 0 {
																										return int32(0)
																									} else {
																										v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																										*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																										v172 = int32(1)
																										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																										v175 = int32(0)
																										m.G0 = v11 + int32(32)
																										return v175
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
															return int32(0)
														} else {
															if v125 != int32(1) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v269 = m.ExcPending
																if v269 != 0 {
																	return int32(0)
																} else {
																	F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																	mBase = m.M
																	v273 = m.ExcPending
																	if v273 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																		mBase = m.M
																		v278 = m.ExcPending
																		if v278 != 0 {
																			return int32(0)
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
																	return int32(0)
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
																			return int32(0)
																		} else {
																			v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																			v151 = int32(1)
																			v152 = v150 + v151
																			*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																			v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																			*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																			v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																			*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																			v175 = v148
																			m.G0 = v11 + int32(32)
																			return v175
																		}
																	} else {
																		F_end_MultiFuncCall(m, l0)
																		mBase = m.M
																		v162 = m.ExcPending
																		if v162 != 0 {
																			return int32(0)
																		} else {
																			v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																			*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																			v172 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																			v175 = int32(0)
																			m.G0 = v11 + int32(32)
																			return v175
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
															return int32(0)
														} else {
															if v99 != 0 {
																F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_10), int32(0))
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(795), int32(_a_F_bt_page_items_bytea_6))
																	mBase = m.M
																	v109 = m.ExcPending
																	if v109 != 0 {
																		return int32(0)
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
																			return int32(0)
																		} else {
																			if v125 != int32(1) {
																				F_errstart_cold(m, int32(21), int32(0))
																				mBase = m.M
																				v269 = m.ExcPending
																				if v269 != 0 {
																					return int32(0)
																				} else {
																					F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																					mBase = m.M
																					v273 = m.ExcPending
																					if v273 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																						mBase = m.M
																						v278 = m.ExcPending
																						if v278 != 0 {
																							return int32(0)
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
																					return int32(0)
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
																							return int32(0)
																						} else {
																							v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																							v151 = int32(1)
																							v152 = v150 + v151
																							*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																							v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																							*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																							v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																							*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																							v175 = v148
																							m.G0 = v11 + int32(32)
																							return v175
																						}
																					} else {
																						F_end_MultiFuncCall(m, l0)
																						mBase = m.M
																						v162 = m.ExcPending
																						if v162 != 0 {
																							return int32(0)
																						} else {
																							v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																							*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																							v172 = int32(1)
																							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																							v175 = int32(0)
																							m.G0 = v11 + int32(32)
																							return v175
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
																	return int32(0)
																} else {
																	if v125 != int32(1) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v269 = m.ExcPending
																		if v269 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																			mBase = m.M
																			v273 = m.ExcPending
																			if v273 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																				mBase = m.M
																				v278 = m.ExcPending
																				if v278 != 0 {
																					return int32(0)
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
																			return int32(0)
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
																					return int32(0)
																				} else {
																					v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																					v151 = int32(1)
																					v152 = v150 + v151
																					*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																					v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																					*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																					v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																					*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																					v175 = v148
																					m.G0 = v11 + int32(32)
																					return v175
																				}
																			} else {
																				F_end_MultiFuncCall(m, l0)
																				mBase = m.M
																				v162 = m.ExcPending
																				if v162 != 0 {
																					return int32(0)
																				} else {
																					v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																					*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																					v172 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																					v175 = int32(0)
																					m.G0 = v11 + int32(32)
																					return v175
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
														return int32(0)
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
																	return int32(0)
																} else {
																	if v125 != int32(1) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v269 = m.ExcPending
																		if v269 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																			mBase = m.M
																			v273 = m.ExcPending
																			if v273 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																				mBase = m.M
																				v278 = m.ExcPending
																				if v278 != 0 {
																					return int32(0)
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
																			return int32(0)
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
																					return int32(0)
																				} else {
																					v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																					v151 = int32(1)
																					v152 = v150 + v151
																					*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																					v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																					*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																					v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																					*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																					v175 = v148
																					m.G0 = v11 + int32(32)
																					return v175
																				}
																			} else {
																				F_end_MultiFuncCall(m, l0)
																				mBase = m.M
																				v162 = m.ExcPending
																				if v162 != 0 {
																					return int32(0)
																				} else {
																					v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																					*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																					v172 = int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																					v175 = int32(0)
																					m.G0 = v11 + int32(32)
																					return v175
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
																	return int32(0)
																} else {
																	if v99 != 0 {
																		F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_10), int32(0))
																		mBase = m.M
																		v104 = m.ExcPending
																		if v104 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(795), int32(_a_F_bt_page_items_bytea_6))
																			mBase = m.M
																			v109 = m.ExcPending
																			if v109 != 0 {
																				return int32(0)
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
																					return int32(0)
																				} else {
																					if v125 != int32(1) {
																						F_errstart_cold(m, int32(21), int32(0))
																						mBase = m.M
																						v269 = m.ExcPending
																						if v269 != 0 {
																							return int32(0)
																						} else {
																							F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																							mBase = m.M
																							v273 = m.ExcPending
																							if v273 != 0 {
																								return int32(0)
																							} else {
																								F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																								mBase = m.M
																								v278 = m.ExcPending
																								if v278 != 0 {
																									return int32(0)
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
																							return int32(0)
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
																									return int32(0)
																								} else {
																									v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																									v151 = int32(1)
																									v152 = v150 + v151
																									*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																									v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																									*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																									v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																									*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																									v175 = v148
																									m.G0 = v11 + int32(32)
																									return v175
																								}
																							} else {
																								F_end_MultiFuncCall(m, l0)
																								mBase = m.M
																								v162 = m.ExcPending
																								if v162 != 0 {
																									return int32(0)
																								} else {
																									v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																									*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																									v172 = int32(1)
																									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																									v175 = int32(0)
																									m.G0 = v11 + int32(32)
																									return v175
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
																			return int32(0)
																		} else {
																			if v125 != int32(1) {
																				F_errstart_cold(m, int32(21), int32(0))
																				mBase = m.M
																				v269 = m.ExcPending
																				if v269 != 0 {
																					return int32(0)
																				} else {
																					F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																					mBase = m.M
																					v273 = m.ExcPending
																					if v273 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																						mBase = m.M
																						v278 = m.ExcPending
																						if v278 != 0 {
																							return int32(0)
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
																					return int32(0)
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
																							return int32(0)
																						} else {
																							v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																							v151 = int32(1)
																							v152 = v150 + v151
																							*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																							v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																							*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																							v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																							*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																							v175 = v148
																							m.G0 = v11 + int32(32)
																							return v175
																						}
																					} else {
																						F_end_MultiFuncCall(m, l0)
																						mBase = m.M
																						v162 = m.ExcPending
																						if v162 != 0 {
																							return int32(0)
																						} else {
																							v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																							*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																							v172 = int32(1)
																							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																							v175 = int32(0)
																							m.G0 = v11 + int32(32)
																							return v175
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
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(788), int32(_a_F_bt_page_items_bytea_6))
																mBase = m.M
																v78 = m.ExcPending
																if v78 != 0 {
																	return int32(0)
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
																			return int32(0)
																		} else {
																			if v125 != int32(1) {
																				F_errstart_cold(m, int32(21), int32(0))
																				mBase = m.M
																				v269 = m.ExcPending
																				if v269 != 0 {
																					return int32(0)
																				} else {
																					F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																					mBase = m.M
																					v273 = m.ExcPending
																					if v273 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																						mBase = m.M
																						v278 = m.ExcPending
																						if v278 != 0 {
																							return int32(0)
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
																					return int32(0)
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
																							return int32(0)
																						} else {
																							v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																							v151 = int32(1)
																							v152 = v150 + v151
																							*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																							v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																							*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																							v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																							*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																							v175 = v148
																							m.G0 = v11 + int32(32)
																							return v175
																						}
																					} else {
																						F_end_MultiFuncCall(m, l0)
																						mBase = m.M
																						v162 = m.ExcPending
																						if v162 != 0 {
																							return int32(0)
																						} else {
																							v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																							*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																							v172 = int32(1)
																							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																							v175 = int32(0)
																							m.G0 = v11 + int32(32)
																							return v175
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
																			return int32(0)
																		} else {
																			if v99 != 0 {
																				F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_10), int32(0))
																				mBase = m.M
																				v104 = m.ExcPending
																				if v104 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(795), int32(_a_F_bt_page_items_bytea_6))
																					mBase = m.M
																					v109 = m.ExcPending
																					if v109 != 0 {
																						return int32(0)
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
																							return int32(0)
																						} else {
																							if v125 != int32(1) {
																								F_errstart_cold(m, int32(21), int32(0))
																								mBase = m.M
																								v269 = m.ExcPending
																								if v269 != 0 {
																									return int32(0)
																								} else {
																									F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																									mBase = m.M
																									v273 = m.ExcPending
																									if v273 != 0 {
																										return int32(0)
																									} else {
																										F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																										mBase = m.M
																										v278 = m.ExcPending
																										if v278 != 0 {
																											return int32(0)
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
																									return int32(0)
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
																											return int32(0)
																										} else {
																											v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																											v151 = int32(1)
																											v152 = v150 + v151
																											*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																											v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																											*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																											v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																											*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																											v175 = v148
																											m.G0 = v11 + int32(32)
																											return v175
																										}
																									} else {
																										F_end_MultiFuncCall(m, l0)
																										mBase = m.M
																										v162 = m.ExcPending
																										if v162 != 0 {
																											return int32(0)
																										} else {
																											v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																											*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																											v172 = int32(1)
																											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																											v175 = int32(0)
																											m.G0 = v11 + int32(32)
																											return v175
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
																					return int32(0)
																				} else {
																					if v125 != int32(1) {
																						F_errstart_cold(m, int32(21), int32(0))
																						mBase = m.M
																						v269 = m.ExcPending
																						if v269 != 0 {
																							return int32(0)
																						} else {
																							F_errmsg_internal(m, int32(_a_F_bt_page_items_bytea_9), int32(0))
																							mBase = m.M
																							v273 = m.ExcPending
																							if v273 != 0 {
																								return int32(0)
																							} else {
																								F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(803), int32(_a_F_bt_page_items_bytea_6))
																								mBase = m.M
																								v278 = m.ExcPending
																								if v278 != 0 {
																									return int32(0)
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
																							return int32(0)
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
																									return int32(0)
																								} else {
																									v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
																									v151 = int32(1)
																									v152 = v150 + v151
																									*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
																									v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
																									*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
																									v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																									*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
																									v175 = v148
																									m.G0 = v11 + int32(32)
																									return v175
																								}
																							} else {
																								F_end_MultiFuncCall(m, l0)
																								mBase = m.M
																								v162 = m.ExcPending
																								if v162 != 0 {
																									return int32(0)
																								} else {
																									v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																									*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
																									v172 = int32(1)
																									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
																									v175 = int32(0)
																									m.G0 = v11 + int32(32)
																									return v175
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
							return int32(0)
						} else {
							v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)))
							v151 = int32(1)
							v152 = v150 + v151
							*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v152)
							v154 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
							*(*int64)(unsafe.Add(mBase, uint32(v143))) = v154 + int64(1)
							v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v151
							v175 = v148
							m.G0 = v11 + int32(32)
							return v175
						}
					} else {
						F_end_MultiFuncCall(m, l0)
						mBase = m.M
						v162 = m.ExcPending
						if v162 != 0 {
							return int32(0)
						} else {
							v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(2)
							v172 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v172)
							v175 = int32(0)
							m.G0 = v11 + int32(32)
							return v175
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v188 = m.ExcPending
				if v188 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v191 = m.ExcPending
					if v191 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_bt_page_items_bytea_12), int32(0))
						mBase = m.M
						v195 = m.ExcPending
						if v195 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_bt_page_items_bytea_5), int32(743), int32(_a_F_bt_page_items_bytea_6))
							mBase = m.M
							v200 = m.ExcPending
							if v200 != 0 {
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
}
func F_bt_page_items_internal(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
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
	var v121 int32
	_ = v121
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
	var v157 int32
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
	var v178 int32
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if l1 != 0 {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
			v22 = v20
		} else {
			v22 = base.I64_extend_i32_u(v19)
		}
		v23 = F_superuser(m)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			if v23 != 0 {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
				if v26 == int32(0) {
					v29 = F_init_MultiFuncCall(m, l0)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = F_textToQualifiedNameList(m, v15)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = F_makeRangeVarFromNameList(m, v31)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v36 = F_relation_openrv(m, v33, int32(1))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									F_bt_index_block_validate(m, v36, v22)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										v41 = F_ReadBuffer(m, v36, base.I32_wrap_i64(v22))
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return int32(0)
										} else {
											F_LockBuffer(m, v41, int32(1))
											mBase = m.M
											v45 = m.ExcPending
											if v45 != 0 {
												return int32(0)
											} else {
												v46 = int32(_a_F_bt_page_items_internal_0)
												v47 = *(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_internal[0]))
												v49 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
												*(*int32)(unsafe.Add(mBase, _c_F_bt_page_items_internal[0])) = v49
												v52 = F_palloc(m, int32(12))
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return int32(0)
												} else {
													v55 = F_palloc(m, int32(_a_F_bt_page_items_internal_1))
													mBase = m.M
													v56 = m.ExcPending
													if v56 != 0 {
														return int32(0)
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
															return int32(0)
														} else {
															F_relation_close(m, v36, int32(1))
															mBase = m.M
															v82 = m.ExcPending
															if v82 != 0 {
																return int32(0)
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
																	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v103
																	v121 = v88
																	v124 = v121 & int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v52)+6)) = uint8(v124)
																	v126 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
																	v127 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(v52)+7)) = uint8(base.B2i32(v126 == v127))
																	v133 = F_get_call_result_type(m, l0, v127, v12+int32(12))
																	mBase = m.M
																	v134 = m.ExcPending
																	if v134 != 0 {
																		return int32(0)
																	} else {
																		if v133 != int32(1) {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v203 = m.ExcPending
																			if v203 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg_internal(m, int32(_a_F_bt_page_items_internal_2), int32(0))
																				mBase = m.M
																				v207 = m.ExcPending
																				if v207 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(_a_F_bt_page_items_internal_3), int32(687), int32(_a_F_bt_page_items_internal_4))
																					mBase = m.M
																					v212 = m.ExcPending
																					if v212 != 0 {
																						return int32(0)
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
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v138
																				*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v52
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
																						return int32(0)
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
																						return int32(0)
																					} else {
																						v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																						*(*int32)(unsafe.Add(mBase, uint32(v172)+20)) = int32(2)
																						v175 = int32(1)
																						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v175)
																						v178 = int32(0)
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
																		return int32(0)
																	} else {
																		if v107 != 0 {
																			*(*int64)(unsafe.Add(mBase, uint32(v12))) = v22
																			F_errmsg_internal(m, int32(_a_F_bt_page_items_internal_5), v12)
																			mBase = m.M
																			v112 = m.ExcPending
																			if v112 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_bt_page_items_internal_3), int32(679), int32(_a_F_bt_page_items_internal_4))
																				mBase = m.M
																				v117 = m.ExcPending
																				if v117 != 0 {
																					return int32(0)
																				} else {
																					*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = int64(0)
																					v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+12)))
																					v121 = v120
																					v124 = v121 & int32(1)
																					*(*uint8)(unsafe.Add(mBase, uint32(v52)+6)) = uint8(v124)
																					v126 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
																					v127 = int32(0)
																					*(*uint8)(unsafe.Add(mBase, uint32(v52)+7)) = uint8(base.B2i32(v126 == v127))
																					v133 = F_get_call_result_type(m, l0, v127, v12+int32(12))
																					mBase = m.M
																					v134 = m.ExcPending
																					if v134 != 0 {
																						return int32(0)
																					} else {
																						if v133 != int32(1) {
																							F_errstart_cold(m, int32(21), int32(0))
																							mBase = m.M
																							v203 = m.ExcPending
																							if v203 != 0 {
																								return int32(0)
																							} else {
																								F_errmsg_internal(m, int32(_a_F_bt_page_items_internal_2), int32(0))
																								mBase = m.M
																								v207 = m.ExcPending
																								if v207 != 0 {
																									return int32(0)
																								} else {
																									F_errfinish(m, int32(_a_F_bt_page_items_internal_3), int32(687), int32(_a_F_bt_page_items_internal_4))
																									mBase = m.M
																									v212 = m.ExcPending
																									if v212 != 0 {
																										return int32(0)
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
																								return int32(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v138
																								*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v52
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
																										return int32(0)
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
																										return int32(0)
																									} else {
																										v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																										*(*int32)(unsafe.Add(mBase, uint32(v172)+20)) = int32(2)
																										v175 = int32(1)
																										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v175)
																										v178 = int32(0)
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
																			*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = int64(0)
																			v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+12)))
																			v121 = v120
																			v124 = v121 & int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(v52)+6)) = uint8(v124)
																			v126 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
																			v127 = int32(0)
																			*(*uint8)(unsafe.Add(mBase, uint32(v52)+7)) = uint8(base.B2i32(v126 == v127))
																			v133 = F_get_call_result_type(m, l0, v127, v12+int32(12))
																			mBase = m.M
																			v134 = m.ExcPending
																			if v134 != 0 {
																				return int32(0)
																			} else {
																				if v133 != int32(1) {
																					F_errstart_cold(m, int32(21), int32(0))
																					mBase = m.M
																					v203 = m.ExcPending
																					if v203 != 0 {
																						return int32(0)
																					} else {
																						F_errmsg_internal(m, int32(_a_F_bt_page_items_internal_2), int32(0))
																						mBase = m.M
																						v207 = m.ExcPending
																						if v207 != 0 {
																							return int32(0)
																						} else {
																							F_errfinish(m, int32(_a_F_bt_page_items_internal_3), int32(687), int32(_a_F_bt_page_items_internal_4))
																							mBase = m.M
																							v212 = m.ExcPending
																							if v212 != 0 {
																								return int32(0)
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
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v138
																						*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v52
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
																								return int32(0)
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
																								return int32(0)
																							} else {
																								v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																								*(*int32)(unsafe.Add(mBase, uint32(v172)+20)) = int32(2)
																								v175 = int32(1)
																								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v175)
																								v178 = int32(0)
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
							return int32(0)
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
							return int32(0)
						} else {
							v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v172)+20)) = int32(2)
							v175 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v175)
							v178 = int32(0)
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
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v190 = m.ExcPending
					if v190 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_bt_page_items_internal_6), int32(0))
						mBase = m.M
						v194 = m.ExcPending
						if v194 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_bt_page_items_internal_3), int32(635), int32(_a_F_bt_page_items_internal_4))
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
				}
			}
		}
	}
}
