package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_r_mark_suffix_with_optional_y_consonant(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
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
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v222 int32
	_ = v222
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v10 <= v11 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return v369
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v365
	v369 = int32(1)
	goto L1
L3:
	;
	v365 = v20 - v8 + v151
	goto L2
L4:
	;
	v159 = v10 - v8
	v160 = v158 + v159
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v160
	if v157 < v160 {
		goto L34
	} else {
		goto L35
	}
L5:
	;
	v156 = v9
	v157 = v11
	v158 = v8
	goto L4
L6:
	;
	goto L7
L7:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v9-int32(1)))))
	if v16 != int32(121) {
		v156 = v9
		v157 = v11
		v158 = v8
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v20 = v10 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L11
L9:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v150 == int32(0) {
		goto L3
	} else {
		goto L33
	}
L10:
	;
	v150 = v143
	goto L9
L11:
	;
	if v20 <= v39 {
		v143 = int32(-1)
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v143 = int32(0)
	goto L10
L13:
	;
	v56 = int32(1)
	v57 = v20 - v56
	v59 = int32(*(*int8)(unsafe.Add(mBase, uint32(v35+v57))))
	v61 = v59 & int32(255)
	if v57 == v39 {
		v116 = v61
		v117 = v56
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if int32(305) < v116 {
		goto L23
	} else {
		goto L24
	}
L15:
	;
	if int32(0) <= v59 {
		v116 = v61
		v117 = v56
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v67 = v61 & int32(63)
	v69 = v20 - int32(2)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v69))))
	v73 = v71 << (uint(int32(6)) % 32)
	if base.B2i32(v69 != v39)&base.B2i32(base.Ui32(v71) < base.Ui32(int32(192))) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v116 = v73&int32(1984) | v67
	v117 = int32(2)
	goto L14
L18:
	;
	goto L19
L19:
	;
	v86 = v73&int32(4032) | v67
	v88 = v20 - int32(3)
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v88))))
	if base.B2i32(v88 != v39)&base.B2i32(base.Ui32(v90) < base.Ui32(int32(224))) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v116 = v90<<(uint(int32(12))%32)&int32(61440) | v86
	v117 = int32(3)
	goto L14
L21:
	;
	goto L22
L22:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+(v35-int32(4))))))
	v116 = v90<<(uint(int32(12))%32)&int32(258048) | v108&int32(7)<<(uint(int32(18))%32) | v86
	v117 = int32(4)
	goto L14
L23:
	;
	v150 = v117
	goto L9
L24:
	;
	goto L25
L25:
	;
	v121 = v116 - int32(97)
	if v121 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v150 = v117
	goto L9
L27:
	;
	goto L28
L28:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v121)>>(uint(int32(3))%32)))+uint32(_consts[1476]))))
	if int32(base.Ui32(v127)>>(uint(v121&int32(7))%32))&int32(1) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v150 = v117
	goto L9
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20 - v117
	goto L32
L32:
	;
	goto L12
L33:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v156 = v154
	v157 = v155
	v158 = v151
	goto L4
L34:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160+v156-int32(1)))))
	if v167 == int32(121) {
		v369 = int32(0)
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v171 = int32(0)
	goto L40
L37:
	;
	goto L36
L38:
	;
	if v222 < int32(0) {
		v369 = v171
		goto L1
	} else {
		goto L58
	}
L40:
	;
	goto L41
L41:
	;
	goto L42
L42:
	;
	v178 = v160
	v180 = int32(1)
	goto L45
L44:
	;
	v222 = v204
	goto L38
L45:
	;
	if v178 <= v157 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L44
L47:
	;
	v222 = int32(-1)
	goto L38
L48:
	;
	goto L49
L49:
	;
	v185 = v178 - int32(1)
	v187 = int32(*(*int8)(unsafe.Add(mBase, uint32(v156+v185))))
	if int32(0) <= v187 {
		v204 = v185
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v208 = int32(1)
	if v208 < v180 {
		v178 = v204
		v180 = v180 - v208
		goto L45
	} else {
		goto L57
	}
L51:
	;
	if v185 <= v157 {
		v204 = v185
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v192 = v185
	goto L53
L53:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156+v192))))
	if base.Ui32(int32(191)) < base.Ui32(v197) {
		v204 = v192
		goto L50
	} else {
		goto L55
	}
L54:
	;
	v204 = v157
	goto L50
L55:
	;
	v201 = v192 - int32(1)
	if v157 < v201 {
		v192 = v201
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	goto L46
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v222
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L61
L59:
	;
	if v354 != 0 {
		v369 = v171
		goto L1
	} else {
		goto L83
	}
L60:
	;
	v354 = v347
	goto L59
L61:
	;
	if v222 <= v243 {
		v347 = int32(-1)
		goto L60
	} else {
		goto L63
	}
L62:
	;
	v347 = int32(0)
	goto L60
L63:
	;
	v260 = int32(1)
	v261 = v222 - v260
	v263 = int32(*(*int8)(unsafe.Add(mBase, uint32(v239+v261))))
	v265 = v263 & int32(255)
	if v261 == v243 {
		v320 = v265
		v321 = v260
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if int32(305) < v320 {
		goto L73
	} else {
		goto L74
	}
L65:
	;
	if int32(0) <= v263 {
		v320 = v265
		v321 = v260
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v271 = v265 & int32(63)
	v273 = v222 - int32(2)
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239+v273))))
	v277 = v275 << (uint(int32(6)) % 32)
	if base.B2i32(v273 != v243)&base.B2i32(base.Ui32(v275) < base.Ui32(int32(192))) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v320 = v277&int32(1984) | v271
	v321 = int32(2)
	goto L64
L68:
	;
	goto L69
L69:
	;
	v290 = v277&int32(4032) | v271
	v292 = v222 - int32(3)
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239+v292))))
	if base.B2i32(v292 != v243)&base.B2i32(base.Ui32(v294) < base.Ui32(int32(224))) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v320 = v294<<(uint(int32(12))%32)&int32(61440) | v290
	v321 = int32(3)
	goto L64
L71:
	;
	goto L72
L72:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222+(v239-int32(4))))))
	v320 = v294<<(uint(int32(12))%32)&int32(258048) | v312&int32(7)<<(uint(int32(18))%32) | v290
	v321 = int32(4)
	goto L64
L73:
	;
	v354 = v321
	goto L59
L74:
	;
	goto L75
L75:
	;
	v325 = v320 - int32(97)
	if v325 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v354 = v321
	goto L59
L77:
	;
	goto L78
L78:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v325)>>(uint(int32(3))%32)))+uint32(_consts[1476]))))
	if int32(base.Ui32(v331)>>(uint(v325&int32(7))%32))&int32(1) == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v354 = v321
	goto L59
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v222 - v321
	goto L82
L82:
	;
	goto L62
L83:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v365 = v355 + v159
	goto L2
}
func F_r_mark_yUm(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	v2 = int32(0)
	v4 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v4 == v2 {
		v100 = v2
		return v100
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v9 = v7 - int32(1)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v9 <= v10 {
			v100 = v2
			return v100
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v9))))
			if v14 != int32(109) {
				v100 = v2
				return v100
			} else {
				v19 = F_find_among_b(m, l0, int32(4308528), int32(4))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					if v19 == int32(0) {
						v100 = v2
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v33 <= v34 {
							v55 = v32
							v56 = v34
							v57 = v31
							v58 = v33 - v31
							v59 = v57 + v58
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v59
							if v56 < v59 {
								v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v55-int32(1)))))
								if v66 == int32(121) {
									v95 = int32(0)
								} else {
									v70 = int32(0)
									v72 = F_skip_b_utf8(m, v55, v59, v56, int32(1))
									mBase = m.M
									if v72 < v70 {
										v95 = v70
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
										v80 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
										mBase = m.M
										if v80 != 0 {
											v95 = v70
										} else {
											v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v91 = v81 + v58
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
											v95 = int32(1)
										}
									}
								}
							} else {
								v70 = int32(0)
								v72 = F_skip_b_utf8(m, v55, v59, v56, int32(1))
								mBase = m.M
								if v72 < v70 {
									v95 = v70
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
									v80 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
									mBase = m.M
									if v80 != 0 {
										v95 = v70
									} else {
										v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v91 = v81 + v58
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
										v95 = int32(1)
									}
								}
							}
						} else {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33+v32-int32(1)))))
							if v39 != int32(121) {
								v55 = v32
								v56 = v34
								v57 = v31
								v58 = v33 - v31
								v59 = v57 + v58
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v59
								if v56 < v59 {
									v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v55-int32(1)))))
									if v66 == int32(121) {
										v95 = int32(0)
									} else {
										v70 = int32(0)
										v72 = F_skip_b_utf8(m, v55, v59, v56, int32(1))
										mBase = m.M
										if v72 < v70 {
											v95 = v70
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
											v80 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
											mBase = m.M
											if v80 != 0 {
												v95 = v70
											} else {
												v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v91 = v81 + v58
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
												v95 = int32(1)
											}
										}
									}
								} else {
									v70 = int32(0)
									v72 = F_skip_b_utf8(m, v55, v59, v56, int32(1))
									mBase = m.M
									if v72 < v70 {
										v95 = v70
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
										v80 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
										mBase = m.M
										if v80 != 0 {
											v95 = v70
										} else {
											v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v91 = v81 + v58
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
											v95 = int32(1)
										}
									}
								}
							} else {
								v43 = v33 - int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43
								v48 = int32(0)
								v49 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), v48)
								mBase = m.M
								v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v49 == v48 {
									v91 = v43 - v31 + v50
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
									v95 = int32(1)
								} else {
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v55 = v53
									v56 = v54
									v57 = v50
									v58 = v33 - v31
									v59 = v57 + v58
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v59
									if v56 < v59 {
										v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v55-int32(1)))))
										if v66 == int32(121) {
											v95 = int32(0)
										} else {
											v70 = int32(0)
											v72 = F_skip_b_utf8(m, v55, v59, v56, int32(1))
											mBase = m.M
											if v72 < v70 {
												v95 = v70
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
												v80 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
												mBase = m.M
												if v80 != 0 {
													v95 = v70
												} else {
													v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v91 = v81 + v58
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
													v95 = int32(1)
												}
											}
										}
									} else {
										v70 = int32(0)
										v72 = F_skip_b_utf8(m, v55, v59, v56, int32(1))
										mBase = m.M
										if v72 < v70 {
											v95 = v70
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v72
											v80 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
											mBase = m.M
											if v80 != 0 {
												v95 = v70
											} else {
												v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v91 = v81 + v58
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v91
												v95 = int32(1)
											}
										}
									}
								}
							}
						}
						v100 = v95
					}
					return v100
				}
			}
		}
	}
}
func F_r_mark_ymUs_(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v2 = int32(0)
	v4 = F_r_check_vowel_harmony(m, l0)
	mBase = m.M
	if v4 == v2 {
		v102 = v2
		return v102
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v8-int32(3) <= v7 {
			v102 = v2
			return v102
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v8-int32(1)))))
			if v16 != int32(159) {
				v102 = v2
				return v102
			} else {
				v21 = F_find_among_b(m, l0, int32(4307472), int32(4))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					if v21 == int32(0) {
						v102 = v2
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v35 <= v36 {
							v57 = v34
							v58 = v36
							v59 = v33
							v60 = v35 - v33
							v61 = v59 + v60
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61
							if v58 < v61 {
								v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v57-int32(1)))))
								if v68 == int32(121) {
									v97 = int32(0)
								} else {
									v72 = int32(0)
									v74 = F_skip_b_utf8(m, v57, v61, v58, int32(1))
									mBase = m.M
									if v74 < v72 {
										v97 = v72
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v74
										v82 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
										mBase = m.M
										if v82 != 0 {
											v97 = v72
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v93 = v83 + v60
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
											v97 = int32(1)
										}
									}
								}
							} else {
								v72 = int32(0)
								v74 = F_skip_b_utf8(m, v57, v61, v58, int32(1))
								mBase = m.M
								if v74 < v72 {
									v97 = v72
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v74
									v82 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
									mBase = m.M
									if v82 != 0 {
										v97 = v72
									} else {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v93 = v83 + v60
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
										v97 = int32(1)
									}
								}
							}
						} else {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v34-int32(1)))))
							if v41 != int32(121) {
								v57 = v34
								v58 = v36
								v59 = v33
								v60 = v35 - v33
								v61 = v59 + v60
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61
								if v58 < v61 {
									v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v57-int32(1)))))
									if v68 == int32(121) {
										v97 = int32(0)
									} else {
										v72 = int32(0)
										v74 = F_skip_b_utf8(m, v57, v61, v58, int32(1))
										mBase = m.M
										if v74 < v72 {
											v97 = v72
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v74
											v82 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
											mBase = m.M
											if v82 != 0 {
												v97 = v72
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v93 = v83 + v60
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
												v97 = int32(1)
											}
										}
									}
								} else {
									v72 = int32(0)
									v74 = F_skip_b_utf8(m, v57, v61, v58, int32(1))
									mBase = m.M
									if v74 < v72 {
										v97 = v72
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v74
										v82 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
										mBase = m.M
										if v82 != 0 {
											v97 = v72
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v93 = v83 + v60
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
											v97 = int32(1)
										}
									}
								}
							} else {
								v45 = v35 - int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v45
								v50 = int32(0)
								v51 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), v50)
								mBase = m.M
								v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v51 == v50 {
									v93 = v45 - v33 + v52
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
									v97 = int32(1)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v57 = v55
									v58 = v56
									v59 = v52
									v60 = v35 - v33
									v61 = v59 + v60
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v61
									if v58 < v61 {
										v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v57-int32(1)))))
										if v68 == int32(121) {
											v97 = int32(0)
										} else {
											v72 = int32(0)
											v74 = F_skip_b_utf8(m, v57, v61, v58, int32(1))
											mBase = m.M
											if v74 < v72 {
												v97 = v72
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v74
												v82 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
												mBase = m.M
												if v82 != 0 {
													v97 = v72
												} else {
													v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v93 = v83 + v60
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
													v97 = int32(1)
												}
											}
										}
									} else {
										v72 = int32(0)
										v74 = F_skip_b_utf8(m, v57, v61, v58, int32(1))
										mBase = m.M
										if v74 < v72 {
											v97 = v72
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v74
											v82 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
											mBase = m.M
											if v82 != 0 {
												v97 = v72
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v93 = v83 + v60
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v93
												v97 = int32(1)
											}
										}
									}
								}
							}
						}
						v102 = v97
					}
					return v102
				}
			}
		}
	}
}
func F_r_mark_ysA(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = v4 - int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 <= v7 {
		v106 = v2
		return v106
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v6))))
		if v11&int32(224) != int32(96) {
			v106 = v2
			return v106
		} else {
			if int32(1)<<(uint(v11)%32)&int32(26658) == int32(0) {
				v106 = v2
				return v106
			} else {
				v24 = F_find_among_b(m, l0, int32(4308192), int32(8))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					if v24 == int32(0) {
						v106 = v2
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v38 <= v39 {
							v60 = v37
							v61 = v39
							v62 = v36
							v63 = v38 - v36
							v64 = v62 + v63
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v64
							if v61 < v64 {
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v60-int32(1)))))
								if v71 == int32(121) {
									v100 = int32(0)
								} else {
									v75 = int32(0)
									v77 = F_skip_b_utf8(m, v60, v64, v61, int32(1))
									mBase = m.M
									if v77 < v75 {
										v100 = v75
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v77
										v85 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
										mBase = m.M
										if v85 != 0 {
											v100 = v75
										} else {
											v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v96 = v86 + v63
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v96
											v100 = int32(1)
										}
									}
								}
							} else {
								v75 = int32(0)
								v77 = F_skip_b_utf8(m, v60, v64, v61, int32(1))
								mBase = m.M
								if v77 < v75 {
									v100 = v75
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v77
									v85 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
									mBase = m.M
									if v85 != 0 {
										v100 = v75
									} else {
										v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v96 = v86 + v63
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v96
										v100 = int32(1)
									}
								}
							}
						} else {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v37-int32(1)))))
							if v44 != int32(121) {
								v60 = v37
								v61 = v39
								v62 = v36
								v63 = v38 - v36
								v64 = v62 + v63
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v64
								if v61 < v64 {
									v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v60-int32(1)))))
									if v71 == int32(121) {
										v100 = int32(0)
									} else {
										v75 = int32(0)
										v77 = F_skip_b_utf8(m, v60, v64, v61, int32(1))
										mBase = m.M
										if v77 < v75 {
											v100 = v75
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v77
											v85 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
											mBase = m.M
											if v85 != 0 {
												v100 = v75
											} else {
												v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v96 = v86 + v63
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v96
												v100 = int32(1)
											}
										}
									}
								} else {
									v75 = int32(0)
									v77 = F_skip_b_utf8(m, v60, v64, v61, int32(1))
									mBase = m.M
									if v77 < v75 {
										v100 = v75
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v77
										v85 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
										mBase = m.M
										if v85 != 0 {
											v100 = v75
										} else {
											v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v96 = v86 + v63
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v96
											v100 = int32(1)
										}
									}
								}
							} else {
								v48 = v38 - int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v48
								v53 = int32(0)
								v54 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), v53)
								mBase = m.M
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v54 == v53 {
									v96 = v48 - v36 + v55
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v96
									v100 = int32(1)
								} else {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v60 = v58
									v61 = v59
									v62 = v55
									v63 = v38 - v36
									v64 = v62 + v63
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v64
									if v61 < v64 {
										v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64+v60-int32(1)))))
										if v71 == int32(121) {
											v100 = int32(0)
										} else {
											v75 = int32(0)
											v77 = F_skip_b_utf8(m, v60, v64, v61, int32(1))
											mBase = m.M
											if v77 < v75 {
												v100 = v75
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v77
												v85 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
												mBase = m.M
												if v85 != 0 {
													v100 = v75
												} else {
													v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v96 = v86 + v63
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v96
													v100 = int32(1)
												}
											}
										}
									} else {
										v75 = int32(0)
										v77 = F_skip_b_utf8(m, v60, v64, v61, int32(1))
										mBase = m.M
										if v77 < v75 {
											v100 = v75
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v77
											v85 = F_in_grouping_b_U(m, l0, int32(2164368), int32(97), int32(305), int32(0))
											mBase = m.M
											if v85 != 0 {
												v100 = v75
											} else {
												v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v96 = v86 + v63
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v96
												v100 = int32(1)
											}
										}
									}
								}
							}
						}
						v106 = v100
					}
					return v106
				}
			}
		}
	}
}
