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
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
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
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24+(v20^int32(-1))<<(uint(int32(2))%32))))
	v38 = v30
	goto L1
L3:
	;
	goto L4
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v38 = v32 + v20<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L26
	} else {
		goto L59
	}
L6:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)) = uint16(v254)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)) = uint16(v260)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)) = uint8(v250)
	m.G0 = v18 + int32(32)
	return v248 & int32(65535)
L7:
	;
	v62 = int32(65535)
	v63 = v61 & v62
	if base.Ui32(v59&v62) < base.Ui32(v63) {
		goto L17
	} else {
		goto L18
	}
L8:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v42) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)))
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)))
	v59 = v57
	v61 = v58
	goto L7
L11:
	;
	v50 = int32(base.Ui32(v42+int32(262120)) >> (uint(int32(2)) % 32))
	goto L13
L12:
	;
	v50 = int32(0)
	goto L13
L13:
	;
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v38+v53)+4))
	if v55 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v56 = int32(2)
	goto L16
L15:
	;
	v56 = int32(1)
	goto L16
L16:
	;
	v59 = v50
	v61 = v56
	goto L7
L17:
	;
	v67 = int32(0)
	v248 = v61
	v250 = v3
	v254 = v67
	v260 = v67
	goto L6
L18:
	;
	goto L19
L19:
	;
	v69 = int32(1)
	v72 = v59 + (v39 ^ v69)
	if base.Ui32(v72&int32(65535)) <= base.Ui32(v63) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v248 = v61
	v250 = v69
	v254 = v72
	v260 = v61
	goto L6
L21:
	;
	goto L22
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v80 = v61
	v81 = v72
	v86 = v72
	goto L23
L23:
	;
	v97 = int32(base.Ui32((v81-v80)&int32(65534))>>(uint(int32(1))%32)) + v80
	v99 = v97 & int32(65535)
	v100 = F__bt_compare(m, l0, v76, v38, v99)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v248 = v240
	v250 = int32(1)
	v254 = v232
	v260 = v240
	goto L6
L25:
	;
	if v100 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L26:
	;
	return int32(0)
L27:
	;
	if v100 != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	if v104 == int32(0) {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v107 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v108 = int32(0)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v99<<(uint(int32(2))%32)+v38)+20))
	v115 = v38 + v112&int32(32767)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+7)))
	if v116&int32(32) == v108 {
		v209 = v108
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v209
	goto L25
L32:
	;
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115)+4)))
	if v121&int32(8192) == int32(0) {
		v209 = v108
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v127 = int32(98304)
	if v112&v127 == v127 {
		v209 = int32(-1)
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v131 = int32(0)
	v133 = v121 & int32(4095)
	if v133 == v131 {
		v209 = v131
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v143 = int32(0)
	v149 = v133
	goto L36
L36:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115)+2)))
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v115))))
	v155 = int32(16)
	v161 = base.I32_div_s(v149-v143, int32(2))
	v162 = v161 + v143
	v165 = v115 + (v153 | v154<<(uint(v155)%32)) + v162*int32(6)
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+2)))
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	v173 = v169 | v170<<(uint(v155)%32)
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+2)))
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165))))
	v178 = v174 | v175<<(uint(v155)%32)
	if base.Ui32(v173) < base.Ui32(v178) {
		v189 = int32(-1)
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v209 = v196
	goto L31
L38:
	;
	if v196 < v197 {
		v143 = v196
		v149 = v197
		goto L36
	} else {
		goto L48
	}
L39:
	;
	if int32(0) < v189 {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	goto L39
L41:
	;
	if base.Ui32(v178) < base.Ui32(v173) {
		v189 = int32(1)
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+4)))
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+4)))
	if base.Ui32(v183) < base.Ui32(v184) {
		v189 = int32(-1)
		goto L40
	} else {
		goto L43
	}
L43:
	;
	v189 = base.B2i32(base.Ui32(v184) < base.Ui32(v183))
	goto L40
L44:
	;
	v196 = v162 + int32(1)
	v197 = v149
	goto L38
L45:
	;
	goto L46
L46:
	;
	if int32(0) <= v189 {
		v209 = v162
		goto L31
	} else {
		goto L47
	}
L47:
	;
	v196 = v143
	v197 = v162
	goto L38
L48:
	;
	goto L37
L49:
	;
	v232 = v97
	goto L51
L50:
	;
	v232 = v86
	goto L51
L51:
	;
	v234 = base.B2i32(int32(0) < v100)
	if int32(0) < v100 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v235 = v81
	goto L54
L53:
	;
	v235 = v97
	goto L54
L54:
	;
	if int32(0) < v100 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v240 = v97 + int32(1)
	goto L57
L56:
	;
	v240 = v80
	goto L57
L57:
	;
	if base.Ui32(v240&int32(65535)) < base.Ui32(v235&int32(65535)) {
		v80 = v240
		v81 = v235
		v86 = v232
		goto L23
	} else {
		goto L58
	}
L58:
	;
	goto L24
L59:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L26
	} else {
		goto L60
	}
L60:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277)+2)))
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277))))
	v280 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277)+4)))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v281 < int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v300
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v301 + int32(4)
	v306 = int32(65535)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v86 & v306
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v80 & v306
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v280
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v278 | v279<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(665115), v18)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L26
	} else {
		goto L65
	}
L62:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v285+(v281^int32(-1))<<(uint(int32(6))%32))+16))
	v300 = v291
	goto L61
L63:
	;
	goto L64
L64:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v293+v281<<(uint(int32(6))%32)+int32(-64))+16))
	v300 = v299
	goto L61
L65:
	;
	F_errfinish(m, int32(486808), int32(575), int32(80079))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L26
	} else {
		goto L66
	}
L66:
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
	var v67 int32
	_ = v67
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+19)))
	if v9 != 0 {
		v67 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l6))) = v67
		return
	} else {
		if l3 != 0 {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+3)))
			if v12&int32(2) != 0 {
				v15 = int32(-1)
			} else {
				v15 = int32(1)
			}
			v67 = v15
			*(*int32)(unsafe.Add(mBase, uint32(l6))) = v67
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
								v67 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l6))) = v67
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
									v67 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(l6))) = v67
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
								v67 = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(l6))) = v67
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
											v67 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(l6))) = v67
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
								v67 = int32(-1)
								*(*int32)(unsafe.Add(mBase, uint32(l6))) = v67
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
									v67 = int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(l6))) = v67
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
											v67 = int32(-1)
											*(*int32)(unsafe.Add(mBase, uint32(l6))) = v67
										}
										return
									}
								}
							} else {
								v67 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l6))) = v67
								return
							}
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
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	v3 = l2
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
	if v5&int32(32) != 0 {
		v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		if v8&int32(8192) != 0 {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
			v30 = v8 & int32(4095)
			v32 = v30 * int32(6)
			if v32 != 0 {
				v33 = F__emscripten_memcpy_bulkmem(m, v22, l1+(v23|v24<<(uint(int32(16))%32)), v32)
				mBase = m.M
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v30
			v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
			v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
			v42 = v36 | v37<<(uint(int32(16))%32)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v13
			v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)) = uint16(v15)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(1)
			v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
			v42 = v19 & int32(8191)
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v12))) = v13
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)) = uint16(v15)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(1)
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
		v42 = v19 & int32(8191)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v42
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v3)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = (v48&int32(8191)+int32(7))&int32(16376) | int32(4)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v58<<(uint(int32(2))%32))+44)) = uint16(v3)
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
				v22 = int32(4362760)
				v24 = *(*int64)(unsafe.Add(mBase, _consts[52]))
				v25 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
				*(*int64)(unsafe.Add(mBase, _consts[52])) = v24 + v25
				v28 = int32(4362768)
				v30 = *(*int64)(unsafe.Add(mBase, _consts[53]))
				v31 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
				*(*int64)(unsafe.Add(mBase, _consts[53])) = v30 + v31
				v34 = int32(4362776)
				v36 = *(*int64)(unsafe.Add(mBase, _consts[54]))
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
				*(*int64)(unsafe.Add(mBase, _consts[54])) = v36 + v37
				v40 = int32(4362784)
				v42 = *(*int64)(unsafe.Add(mBase, _consts[55]))
				v43 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
				*(*int64)(unsafe.Add(mBase, _consts[55])) = v42 + v43
				v46 = int32(4362792)
				v48 = *(*int64)(unsafe.Add(mBase, _consts[56]))
				v49 = *(*int64)(unsafe.Add(mBase, uint32(v17)+32))
				*(*int64)(unsafe.Add(mBase, _consts[56])) = v48 + v49
				v52 = int32(4362800)
				v54 = *(*int64)(unsafe.Add(mBase, _consts[57]))
				v55 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
				*(*int64)(unsafe.Add(mBase, _consts[57])) = v54 + v55
				v58 = int32(4362808)
				v60 = *(*int64)(unsafe.Add(mBase, _consts[58]))
				v61 = *(*int64)(unsafe.Add(mBase, uint32(v17)+48))
				*(*int64)(unsafe.Add(mBase, _consts[58])) = v60 + v61
				v64 = int32(4362816)
				v66 = *(*int64)(unsafe.Add(mBase, _consts[59]))
				v67 = *(*int64)(unsafe.Add(mBase, uint32(v17)+56))
				*(*int64)(unsafe.Add(mBase, _consts[59])) = v66 + v67
				v70 = int32(4362824)
				v72 = *(*int64)(unsafe.Add(mBase, _consts[60]))
				v73 = *(*int64)(unsafe.Add(mBase, uint32(v17)+64))
				*(*int64)(unsafe.Add(mBase, _consts[60])) = v72 + v73
				v76 = int32(4362832)
				v78 = *(*int64)(unsafe.Add(mBase, _consts[61]))
				v79 = *(*int64)(unsafe.Add(mBase, uint32(v17)+72))
				*(*int64)(unsafe.Add(mBase, _consts[61])) = v78 + v79
				v82 = int32(4362840)
				v84 = *(*int64)(unsafe.Add(mBase, _consts[62]))
				v85 = *(*int64)(unsafe.Add(mBase, uint32(v17)+80))
				*(*int64)(unsafe.Add(mBase, _consts[62])) = v84 + v85
				v88 = int32(4362848)
				v90 = *(*int64)(unsafe.Add(mBase, _consts[63]))
				v91 = *(*int64)(unsafe.Add(mBase, uint32(v17)+88))
				*(*int64)(unsafe.Add(mBase, _consts[63])) = v90 + v91
				v94 = int32(4362856)
				v96 = *(*int64)(unsafe.Add(mBase, _consts[64]))
				v97 = *(*int64)(unsafe.Add(mBase, uint32(v17)+96))
				*(*int64)(unsafe.Add(mBase, _consts[64])) = v96 + v97
				v100 = int32(4362864)
				v102 = *(*int64)(unsafe.Add(mBase, _consts[65]))
				v103 = *(*int64)(unsafe.Add(mBase, uint32(v17)+104))
				*(*int64)(unsafe.Add(mBase, _consts[65])) = v102 + v103
				v106 = int32(4362872)
				v108 = *(*int64)(unsafe.Add(mBase, _consts[66]))
				v109 = *(*int64)(unsafe.Add(mBase, uint32(v17)+112))
				*(*int64)(unsafe.Add(mBase, _consts[66])) = v108 + v109
				v112 = int32(4362880)
				v114 = *(*int64)(unsafe.Add(mBase, _consts[67]))
				v115 = *(*int64)(unsafe.Add(mBase, uint32(v17)+120))
				*(*int64)(unsafe.Add(mBase, _consts[67])) = v114 + v115
				v118 = int32(4362904)
				v120 = *(*int64)(unsafe.Add(mBase, _consts[68]))
				v121 = *(*int64)(unsafe.Add(mBase, uint32(v21)+16))
				*(*int64)(unsafe.Add(mBase, _consts[68])) = v120 + v121
				v124 = int32(4362888)
				v126 = *(*int64)(unsafe.Add(mBase, _consts[69]))
				v127 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
				*(*int64)(unsafe.Add(mBase, _consts[69])) = v126 + v127
				v130 = int32(4362896)
				v132 = *(*int64)(unsafe.Add(mBase, _consts[70]))
				v133 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
				*(*int64)(unsafe.Add(mBase, _consts[70])) = v132 + v133
				v136 = int32(4362912)
				v138 = *(*int64)(unsafe.Add(mBase, _consts[71]))
				v139 = *(*int64)(unsafe.Add(mBase, uint32(v21)+24))
				*(*int64)(unsafe.Add(mBase, _consts[71])) = v138 + v139
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
					v160 = *(*int32)(unsafe.Add(mBase, _consts[72]))
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
				v160 = *(*int32)(unsafe.Add(mBase, _consts[72]))
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
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
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
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
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
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
	if v114&int32(20) != 0 {
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
	v37 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37+(v24^int32(-1))<<(uint(int32(2))%32))))
	v51 = v43
	goto L4
L6:
	;
	goto L7
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v51 = v45 + v24<<(uint(int32(13))%32) + int32(-8192)
	goto L4
L8:
	;
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
	if l4 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if v57&int32(20) != 0 {
		goto L29
	} else {
		goto L30
	}
L10:
	;
	if v57&int32(128) == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	if v24 < int32(0) {
		goto L13
	} else {
		goto L14
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
	v67 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v67+(v24^int32(-1))<<(uint(int32(6))%32))+16))
	v82 = v73
	goto L12
L14:
	;
	goto L15
L15:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75+v24<<(uint(int32(6))%32)+int32(-64))+16))
	v82 = v81
	goto L12
L16:
	;
	F__bt_unlockbuf(m, v24)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+12)))
	if v92&int32(128) != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	return int32(0)
L20:
	;
	F__bt_lockbuf(m, v24, int32(2))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
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
	v96 = m.ExcPending
	if v96 != 0 {
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
	v100 = m.ExcPending
	if v100 != 0 {
		goto L19
	} else {
		goto L27
	}
L25:
	;
	v97 = F__bt_getbuf(m, l0, v82, l6)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	v24 = v97
	goto L1
L27:
	;
	v101 = F__bt_getbuf(m, l0, v82, l6)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	v24 = v101
	goto L1
L29:
	;
	v110 = v54
	goto L31
L30:
	;
	v106 = F__bt_compare(m, l0, l2, v51, int32(1))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L19
	} else {
		goto L32
	}
L31:
	;
	v111 = F__bt_relandgetbuf(m, l0, v24, v110, l6)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L19
	} else {
		goto L34
	}
L32:
	;
	if v106 < v18^int32(1) {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v110 = v109
	goto L31
L34:
	;
	v24 = v111
	goto L1
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
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
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v121 + int32(4)
	F_errmsg_internal(m, int32(667544), v16)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L19
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(486808), int32(323), int32(101965))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
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
	var v5 int32
	_ = v5
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
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
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v381 int32
	_ = v381
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v408 int32
	_ = v408
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v526 int32
	_ = v526
	var v533 int32
	_ = v533
	var v547 int32
	_ = v547
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v768 int32
	_ = v768
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v795 int32
	_ = v795
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v814 int32
	_ = v814
	var v821 int32
	_ = v821
	var v828 int32
	_ = v828
	var v836 int32
	_ = v836
	v4 = l3
	v5 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	if v25 < v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+16)))
	if v25 < int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29+(v25^int32(-1))<<(uint(int32(2))%32))))
	v43 = v35
	goto L1
L3:
	;
	goto L4
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v43 = v37 + v25<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+60)) = v63
	v65 = v44 + v43
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+64)) = v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+84)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = v68
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v73 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v48+(v25^int32(-1))<<(uint(int32(6))%32))+16))
	v63 = v54
	goto L5
L7:
	;
	goto L8
L8:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56+v25<<(uint(int32(6))%32)+int32(-64))+16))
	v63 = v62
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
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_PredicateLockPage(m, v23, v81, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L15
	} else {
		goto L17
	}
L12:
	;
	v76 = v68
	goto L14
L13:
	;
	v76 = v66
	goto L14
L14:
	;
	F__bt_parallel_release(m, l0, v76, v63)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
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
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+192))
	v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v86)+8)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+12)))
	v90 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+30)) = uint16(v90)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+26)) = v90
	*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+13)) = uint8(v90)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+12)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v90
	if base.Ui32(int32(25)) <= base.Ui32(v89) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v109 = int32(base.Ui32(v89+int32(262120)) >> (uint(int32(2)) % 32))
	goto L20
L19:
	;
	v109 = v90
	goto L20
L20:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+2)) = uint16(v109)
	v111 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+24)) = uint8(v111)
	if v88 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v115 = int32(2)
	goto L23
L22:
	;
	v115 = v111
	goto L23
L23:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v21))) = uint16(v115)
	if l1 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	m.G0 = v21 + int32(32)
	return v836
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+100)) = v828
	*(*int32)(unsafe.Add(mBase, uint32(v24)+96)) = v814
	v836 = base.B2i32(v821 <= v814)
	goto L24
L26:
	;
	if v85 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	if v85 != 0 {
		goto L107
	} else {
		goto L108
	}
L29:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v119 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	if v4 != 0 {
		goto L39
	} else {
		goto L40
	}
L32:
	;
	v145 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+18)) = uint16(v145)
	goto L31
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
	v125 = v43 + v122&int32(32767)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v125
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+18)))
	if v127 != int32(1) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v131 = F__bt_scanbehind_checkkeys(m, l0, int32(1), v125)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L15
	} else {
		goto L35
	}
L35:
	;
	if v131 != 0 {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+17)) = uint8(v133)
	v135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+89)) = uint8(v135)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v138 == v135 {
		v836 = v135
		goto L24
	} else {
		goto L37
	}
L37:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
	F__bt_parallel_primscan_schedule(m, l0, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L15
	} else {
		goto L38
	}
L38:
	;
	v836 = v135
	goto L24
L39:
	;
	if base.Ui32(v115) < base.Ui32(l2) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	if base.Ui32(v109&int32(65535)) <= base.Ui32(v115) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	F__bt_set_startikey(m, l0, v21)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L15
	} else {
		goto L42
	}
L42:
	;
	goto L39
L43:
	;
	v154 = l2
	goto L45
L44:
	;
	v154 = v115
	goto L45
L45:
	;
	v156 = v24 + int32(104)
	v163 = v154
	v168 = v5
	goto L46
L46:
	;
	v177 = int32(65535)
	v178 = v109 & v177
	if base.Ui32(v178) < base.Ui32(v163&v177) {
		v408 = v168
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+24)))
	if v417 != int32(1) {
		goto L92
	} else {
		goto L93
	}
L48:
	;
	goto L47
L49:
	;
	v186 = v163
	goto L50
L50:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v186&int32(65535)<<(uint(int32(2))%32)+(v43+int32(24))-int32(4))))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	if v208 != int32(1) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v408 = v168
	goto L48
L52:
	;
	if base.Ui32(v392&int32(65535)) <= base.Ui32(v178) {
		v186 = v392
		goto L50
	} else {
		goto L89
	}
L53:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+20)) = uint16(v186)
	v222 = v43 + v207&int32(32767)
	v223 = F__bt_checkkeys(m, l0, v21, base.B2i32(v85 != int32(0)), v222, v87)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L15
	} else {
		goto L56
	}
L54:
	;
	v211 = int32(98304)
	if v207&v211 != v211 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v392 = v186 + int32(1)
	goto L52
L56:
	;
	if v85 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if v223 == int32(0) {
		v381 = v168
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+22)))
	if base.Ui32(int32(2047)) < base.Ui32((v227-int32(1))&int32(65535)) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v234 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+22)) = uint16(v234)
	v392 = v227
	goto L52
L60:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+24)))
	if v390 != 0 {
		v163 = v186 + v390
		v168 = v381
		goto L46
	} else {
		goto L88
	}
L61:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+7)))
	if v239&int32(32) != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v277 = v156 + v168*int32(10)
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+2)))
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222))))
	v283 = v222 + (v278 | v279<<(uint(int32(16))%32))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = v284
	v286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v283)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v277)+6)) = uint16(v186)
	*(*uint16)(unsafe.Add(mBase, uint32(v277)+4)) = uint16(v286)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v289 != 0 {
		goto L74
	} else {
		goto L75
	}
L63:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+5)))
	if v242&int32(32) != 0 {
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v247 = v156 + v168*int32(10)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v248
	v250 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v247)+6)) = uint16(v186)
	*(*uint16)(unsafe.Add(mBase, uint32(v247)+4)) = uint16(v250)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v253 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L65
L67:
	;
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+6)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v24)+84))
	*(*uint16)(unsafe.Add(mBase, uint32(v247)+8)) = uint16(v255)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v24)+84))
	v261 = v254 & int32(8191)
	if v261 != 0 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	goto L69
L69:
	;
	v381 = v168 + int32(1)
	goto L60
L70:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v24)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+84)) = v264 + (v261+int32(7))&int32(16376)
	goto L69
L71:
	;
	v262 = F__emscripten_memcpy_bulkmem(m, v257+v258, v222, v261)
	mBase = m.M
	goto L73
L72:
	;
	goto L73
L73:
	;
	goto L70
L74:
	;
	v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+2)))
	v291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222))))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v24)+84))
	*(*uint16)(unsafe.Add(mBase, uint32(v277)+8)) = uint16(v292)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v24)+84))
	v296 = v294 + v295
	v303 = (v290 | v291<<(uint(int32(16))%32) + int32(7)) & int32(-8)
	if v303 != 0 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v318 = int32(0)
	goto L76
L76:
	;
	v319 = int32(1)
	v321 = v168 + v319
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+4)))
	if v322&int32(4094) == int32(0) {
		v381 = v321
		goto L60
	} else {
		goto L81
	}
L77:
	;
	v306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v305)+6)))
	v309 = v306&int32(57344) | v303
	*(*uint16)(unsafe.Add(mBase, uint32(v305)+6)) = uint16(v309)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v24)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+84)) = v311 + v303
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277)+8)))
	v318 = v314
	goto L76
L78:
	;
	v304 = F__emscripten_memcpy_bulkmem(m, v296, v222, v303)
	mBase = m.M
	v305 = v304
	goto L80
L79:
	;
	v305 = v296
	goto L80
L80:
	;
	goto L77
L81:
	;
	v329 = v319
	v336 = v321
	goto L82
L82:
	;
	v347 = v156 + v336*int32(10)
	v348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+2)))
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222))))
	v356 = v222 + (v348 | v349<<(uint(int32(16))%32)) + v329*int32(6)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	*(*int32)(unsafe.Add(mBase, uint32(v347))) = v357
	v359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v356)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v347)+6)) = uint16(v186)
	*(*uint16)(unsafe.Add(mBase, uint32(v347)+4)) = uint16(v359)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v362 != 0 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v381 = v365
	goto L60
L84:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v347)+8)) = uint16(v318)
	goto L86
L85:
	;
	goto L86
L86:
	;
	v364 = int32(1)
	v365 = v336 + v364
	v367 = v329 + v364
	v368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+4)))
	if base.Ui32(v367) < base.Ui32(v368&int32(4095)) {
		v329 = v367
		v336 = v365
		goto L82
	} else {
		goto L87
	}
L87:
	;
	goto L83
L88:
	;
	v408 = v381
	goto L48
L89:
	;
	goto L51
L90:
	;
	v468 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+92)) = v468
	v814 = v408 - int32(1)
	v821 = v468
	v828 = v468
	goto L25
L91:
	;
	v464 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+89)) = uint8(v464)
	goto L90
L92:
	;
	if v417 != 0 {
		goto L90
	} else {
		goto L106
	}
L93:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+18)))
	if v420 != 0 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v421 == int32(0) {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
	v427 = v43 + v424&int32(32767)
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+13)))
	if v428 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	F__bt_start_array_keys(m, l0, int32(1))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L15
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v434 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v434
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+13)) = uint8(v434)
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427)+7)))
	if v440&int32(32) == v434 {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L98
L100:
	;
	v455 = F__bt_checkkeys(m, l0, v21, base.B2i32(v85 != v434), v427, v454)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L15
	} else {
		goto L104
	}
L101:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v23)+192))
	v452 = int32(*(*int16)(unsafe.Add(mBase, uint32(v451)+8)))
	v454 = v452
	goto L100
L102:
	;
	v445 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v427)+4)))
	if v445&int32(8192) != 0 {
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v454 = v445 & int32(4095)
	goto L100
L104:
	;
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+24)))
	if v457&int32(1) == int32(0) {
		goto L91
	} else {
		goto L105
	}
L105:
	;
	goto L90
L106:
	;
	goto L91
L107:
	;
	if base.Ui32(v109&int32(65535)) < base.Ui32(v115) {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	goto L109
L109:
	;
	if v4 != 0 {
		goto L118
	} else {
		goto L119
	}
L110:
	;
	v505 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+18)) = uint16(v505)
	goto L109
L111:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v477 == int32(0) {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v115<<(uint(int32(2))%32)+v43)+20))
	v486 = v43 + v483&int32(32767)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v486
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+18)))
	if v488 != int32(1) {
		goto L110
	} else {
		goto L113
	}
L113:
	;
	v491 = F__bt_scanbehind_checkkeys(m, l0, l1, v486)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L15
	} else {
		goto L114
	}
L114:
	;
	if v491 != 0 {
		goto L110
	} else {
		goto L115
	}
L115:
	;
	v493 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+17)) = uint8(v493)
	v495 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+88)) = uint8(v495)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v498 == v495 {
		v836 = v495
		goto L24
	} else {
		goto L116
	}
L116:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
	F__bt_parallel_primscan_schedule(m, l0, v501)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L15
	} else {
		goto L117
	}
L117:
	;
	v836 = v495
	goto L24
L118:
	;
	v514 = v109 & int32(65535)
	if base.Ui32(l2) < base.Ui32(v514) {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	if base.Ui32(v109&int32(65535)) <= base.Ui32(v115) {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	F__bt_set_startikey(m, l0, v21)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L15
	} else {
		goto L121
	}
L121:
	;
	goto L118
L122:
	;
	v516 = l2
	goto L124
L123:
	;
	v516 = v514
	goto L124
L124:
	;
	v518 = v24 + int32(104)
	v526 = v516
	v533 = int32(1358)
	goto L125
L125:
	;
	if base.Ui32(v526&int32(65535)) < base.Ui32(v115) {
		v795 = v533
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+24)))
	if v802 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L127:
	;
	goto L126
L128:
	;
	v547 = v526
	goto L129
L129:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
	v563 = v547 & int32(65535)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v563<<(uint(int32(2))%32)+(v43+int32(24))-int32(4))))
	v570 = int32(98304)
	v574 = v561 & base.B2i32(v569&v570 == v570)
	if v574 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L130:
	;
	v795 = v533
	goto L127
L131:
	;
	if base.Ui32(v115) <= base.Ui32(v779&int32(65535)) {
		v547 = v779
		goto L129
	} else {
		goto L177
	}
L132:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+20)) = uint16(v547)
	v583 = v43 + v569&int32(32767)
	if v85 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L133:
	;
	if base.Ui32(v563) <= base.Ui32(v115) {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v779 = v547 - int32(1)
	goto L131
L135:
	;
	if v617 == int32(0) {
		v768 = v533
		goto L149
	} else {
		goto L150
	}
L136:
	;
	v606 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+18)))
	if v606 != 0 {
		v795 = v533
		goto L127
	} else {
		goto L147
	}
L137:
	;
	v601 = F__bt_checkkeys(m, l0, v21, base.B2i32(v85 != int32(0)), v583, v87)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L15
	} else {
		goto L145
	}
L138:
	;
	if v563 != v115 {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v587 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+13)))
	if v587 == int32(1) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v590 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v590
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+13)) = uint8(v590)
	F__bt_start_array_keys(m, l0, l1)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L15
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	v597 = F__bt_checkkeys(m, l0, v21, int32(1), v583, v87)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L15
	} else {
		goto L144
	}
L143:
	;
	goto L142
L144:
	;
	v605 = v597
	goto L136
L145:
	;
	if v85 == int32(0) {
		v617 = v601
		goto L135
	} else {
		goto L146
	}
L146:
	;
	v605 = v601
	goto L136
L147:
	;
	v607 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+22)))
	if base.Ui32(int32(2047)) < base.Ui32((v607-int32(1))&int32(65535)) {
		v617 = v605
		goto L135
	} else {
		goto L148
	}
L148:
	;
	v614 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+22)) = uint16(v614)
	v779 = v607
	goto L131
L149:
	;
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+24)))
	if v775 != 0 {
		v526 = v547 - v775
		v533 = v768
		goto L125
	} else {
		goto L176
	}
L150:
	;
	if v574^int32(1) == int32(0) {
		v768 = v533
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+7)))
	if v624&int32(32) != 0 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v661 = v533 - int32(1)
	v664 = v518 + v661*int32(10)
	v665 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583)+2)))
	v666 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583))))
	v670 = v583 + (v665 | v666<<(uint(int32(16))%32))
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v670)))
	*(*int32)(unsafe.Add(mBase, uint32(v664))) = v671
	v673 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v670)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v664)+6)) = uint16(v547)
	*(*uint16)(unsafe.Add(mBase, uint32(v664)+4)) = uint16(v673)
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v676 != 0 {
		goto L162
	} else {
		goto L163
	}
L153:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583)+5)))
	if v627&int32(32) != 0 {
		goto L152
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v631 = v533 - int32(1)
	v634 = v518 + v631*int32(10)
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v583)))
	*(*int32)(unsafe.Add(mBase, uint32(v634))) = v635
	v637 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v634)+6)) = uint16(v547)
	*(*uint16)(unsafe.Add(mBase, uint32(v634)+4)) = uint16(v637)
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v640 == int32(0) {
		v768 = v631
		goto L149
	} else {
		goto L157
	}
L156:
	;
	goto L155
L157:
	;
	v643 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583)+6)))
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v24)+84))
	*(*uint16)(unsafe.Add(mBase, uint32(v634)+8)) = uint16(v644)
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v24)+84))
	v650 = v643 & int32(8191)
	if v650 != 0 {
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v24)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+84)) = v653 + (v650+int32(7))&int32(16376)
	v768 = v631
	goto L149
L159:
	;
	v651 = F__emscripten_memcpy_bulkmem(m, v646+v647, v583, v650)
	mBase = m.M
	goto L161
L160:
	;
	goto L161
L161:
	;
	goto L158
L162:
	;
	v677 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583)+2)))
	v678 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583))))
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v24)+84))
	*(*uint16)(unsafe.Add(mBase, uint32(v664)+8)) = uint16(v679)
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v24)+84))
	v683 = v681 + v682
	v690 = (v677 | v678<<(uint(int32(16))%32) + int32(7)) & int32(-8)
	if v690 != 0 {
		goto L166
	} else {
		goto L167
	}
L163:
	;
	v705 = int32(0)
	goto L164
L164:
	;
	v707 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583)+4)))
	if v707&int32(4094) == int32(0) {
		v768 = v661
		goto L149
	} else {
		goto L169
	}
L165:
	;
	v693 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v692)+6)))
	v696 = v693&int32(57344) | v690
	*(*uint16)(unsafe.Add(mBase, uint32(v692)+6)) = uint16(v696)
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v24)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+84)) = v698 + v690
	v701 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v664)+8)))
	v705 = v701
	goto L164
L166:
	;
	v691 = F__emscripten_memcpy_bulkmem(m, v683, v583, v690)
	mBase = m.M
	v692 = v691
	goto L168
L167:
	;
	v692 = v683
	goto L168
L168:
	;
	goto L165
L169:
	;
	v720 = int32(1)
	v723 = v661
	goto L170
L170:
	;
	v731 = v723 - int32(1)
	v734 = v518 + v731*int32(10)
	v735 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583)+2)))
	v736 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583))))
	v743 = v583 + (v735 | v736<<(uint(int32(16))%32)) + v720*int32(6)
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v743)))
	*(*int32)(unsafe.Add(mBase, uint32(v734))) = v744
	v746 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v743)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v734)+6)) = uint16(v547)
	*(*uint16)(unsafe.Add(mBase, uint32(v734)+4)) = uint16(v746)
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v749 != 0 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v768 = v731
	goto L149
L172:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v734)+8)) = uint16(v705)
	goto L174
L173:
	;
	goto L174
L174:
	;
	v752 = v720 + int32(1)
	v753 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583)+4)))
	if base.Ui32(v752) < base.Ui32(v753&int32(4095)) {
		v720 = v752
		v723 = v731
		goto L170
	} else {
		goto L175
	}
L175:
	;
	goto L171
L176:
	;
	v795 = v768
	goto L127
L177:
	;
	goto L130
L178:
	;
	v805 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+88)) = uint8(v805)
	goto L180
L179:
	;
	goto L180
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+92)) = v795
	v808 = int32(1357)
	v814 = v808
	v821 = v795
	v828 = v808
	goto L25
}
