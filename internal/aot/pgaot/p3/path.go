package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_GetSearchPathMatcher(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int64
	_ = v47
	F_recomputeNamespacePath(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v9 = int32(4487040)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = l0
	v14 = F_palloc0(m, int32(16))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[428]))
	v18 = F_list_copy(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v43
	v47 = *(*int64)(unsafe.Add(mBase, _consts[432]))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v47
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v10
	return v14
L5:
	;
	if v18 == int32(0) {
		v43 = int32(0)
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v22 = v18
	goto L7
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v29 = *(*int32)(unsafe.Add(mBase, _consts[431]))
	if v27 == v29 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v43 = int32(0)
	goto L4
L9:
	;
	v43 = v22
	goto L4
L10:
	;
	goto L11
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	if v32 == v27 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v38 = F_list_delete_first(m, v22)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L16
	}
L13:
	;
	v34 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+5)) = uint8(v34)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v36 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+4)) = uint8(v36)
	goto L12
L16:
	;
	if v38 != 0 {
		v22 = v38
		goto L7
	} else {
		goto L17
	}
L17:
	;
	goto L8
}
func F_path_add(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 float64
	_ = v71
	var v73 float64
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 float64
	_ = v111
	var v113 int32
	_ = v113
	var v118 float64
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			if v17 == int32(0) {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				if v20 == int32(0) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
					v29 = v27 + v28
					if base.Ui32(int32(268435455)) < base.Ui32(v29) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(438431), int32(0))
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(490692), int32(4368), int32(460678))
									mBase = m.M
									v149 = m.ExcPending
									if v149 != 0 {
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
						v33 = v29 << (uint(int32(4)) % 32)
						if v33 == int32(2147483632) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(438431), int32(0))
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(490692), int32(4368), int32(460678))
										mBase = m.M
										v149 = m.ExcPending
										if v149 != 0 {
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
							v37 = v33 + int32(16)
							v38 = F_palloc(m, v37)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v38))) = v37 << (uint(int32(2)) % 32)
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v43 + v44
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
								v48 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v48
								*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v47
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
								if v48 < v52 {
									v55 = int32(16)
									v59 = v48
									for {
										v68 = v59 << (uint(int32(4)) % 32)
										v69 = v38 + v55 + v68
										v70 = v68 + (v10 + v55)
										v71 = *(*float64)(unsafe.Add(mBase, uint32(v70)))
										*(*float64)(unsafe.Add(mBase, uint32(v69))) = v71
										v73 = *(*float64)(unsafe.Add(mBase, uint32(v70)+8))
										*(*float64)(unsafe.Add(mBase, uint32(v69)+8)) = v73
										v76 = v59 + int32(1)
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
										if v76 < v77 {
											v59 = v76
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								v87 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
								if int32(0) < v87 {
									v90 = int32(16)
									v91 = v38 + v90
									v95 = int32(0)
									for {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
										v105 = int32(4)
										v110 = v15 + v90 + v95<<(uint(v105)%32)
										v111 = *(*float64)(unsafe.Add(mBase, uint32(v110)))
										*(*float64)(unsafe.Add(mBase, uint32(v91+(v103+v95)<<(uint(v105)%32)))) = v111
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
										v118 = *(*float64)(unsafe.Add(mBase, uint32(v110)+8))
										*(*float64)(unsafe.Add(mBase, uint32(v91+(v113+v95)<<(uint(v105)%32))+8)) = v118
										v121 = v95 + int32(1)
										v122 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
										if v121 < v122 {
											v95 = v121
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								return v38
							}
						}
					}
				} else {
					v23 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
					return int32(0)
				}
			} else {
				v23 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
				return int32(0)
			}
		}
	}
}
func F_path_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = int32(44)
	v17 = F___strchrnul(m, v15, v16)
	mBase = m.M
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v19 == v16 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v12 + int32(48)
	return v280
L2:
	;
	v258 = F_errsave_start(m, v14)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L35
	} else {
		goto L69
	}
L3:
	;
	if v23 == int32(0) {
		goto L2
	} else {
		goto L7
	}
L4:
	;
	v23 = v17
	goto L6
L5:
	;
	v23 = v2
	goto L6
L6:
	;
	goto L3
L7:
	;
	v28 = v23
	v29 = v2
	goto L8
L8:
	;
	v35 = int32(1)
	v39 = int32(44)
	v40 = F___strchrnul(m, v28+v35, v39)
	mBase = m.M
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v42 == v39 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v50 = int32(1)
	if v29&v50 != 0 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	if v46 != 0 {
		v28 = v46
		v29 = v29 + v35
		goto L8
	} else {
		goto L14
	}
L11:
	;
	v46 = v40
	goto L13
L12:
	;
	v46 = int32(0)
	goto L13
L13:
	;
	goto L10
L14:
	;
	goto L9
L15:
	;
	v54 = int32(-1)
	goto L17
L16:
	;
	v54 = (v29 + int32(2)) >> (uint(v50) % 32)
	goto L17
L17:
	;
	if v54 <= int32(0) {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v59 = v15
	goto L20
L19:
	;
	v109 = v54 << (uint(int32(4)) % 32)
	v110 = base.I32_div_s(v109, v54)
	if base.B2i32(v110 == int32(16))&base.B2i32(v109 != int32(2147483632)) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v59
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if base.Ui32(v67-int32(9)) < base.Ui32(int32(5)) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v81 = F_strlen(m, v59)
	mBase = m.M
	v88 = v81 + int32(1)
	goto L27
L22:
	;
	goto L21
L23:
	;
	v59 = v59 + int32(1)
	goto L20
L24:
	;
	v72 = int32(0)
	switch v67 - int32(32) {
	case 0:
		goto L23
	default:
		v106 = v59
		v107 = v72
		goto L19
	case 8:
		goto L22
	}
L25:
	;
	if v100 != v59 {
		v106 = v59
		v107 = v72
		goto L19
	} else {
		goto L31
	}
L26:
	;
	goto L25
L27:
	;
	v90 = int32(0)
	if v88 == v90 {
		v100 = v90
		goto L26
	} else {
		goto L29
	}
L28:
	;
	v100 = v95
	goto L26
L29:
	;
	v94 = v88 - int32(1)
	v95 = v59 + v94
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v96 != int32(40) {
		v88 = v94
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v102 = int32(1)
	v104 = v59 + v102
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v104
	v106 = v104
	v107 = v102
	goto L19
L32:
	;
	v118 = int32(0)
	v119 = F_errsave_start(m, v14)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v138 = v109 + int32(16)
	v139 = F_palloc(m, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L35
	} else {
		goto L41
	}
L35:
	;
	return int32(0)
L36:
	;
	if v119 == int32(0) {
		v280 = v118
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(438431), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	F_errsave_finish(m, v14, int32(490692), int32(1438), int32(277992))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L35
	} else {
		goto L40
	}
L40:
	;
	v280 = v118
	goto L1
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139)+4)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v138 << (uint(int32(2)) % 32)
	v153 = F_path_decode(m, v106, int32(1), v54, v139+int32(16), v12+int32(47), v12+int32(40), int32(320116), v15, v14)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L35
	} else {
		goto L42
	}
L42:
	;
	if v153 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v157 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v157)
	v280 = int32(0)
	goto L1
L44:
	;
	goto L45
L45:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	if v107 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v214&int32(255) != 0 {
		goto L61
	} else {
		goto L62
	}
L47:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	v214 = v163
	goto L46
L48:
	;
	goto L49
L49:
	;
	v165 = v160 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v165
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if v167 == int32(41) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v172 = v165
	goto L53
L51:
	;
	goto L52
L52:
	;
	v190 = int32(0)
	v191 = F_errsave_start(m, v14)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L35
	} else {
		goto L56
	}
L53:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v179-int32(9)))&base.B2i32(v179 != int32(32)) != 0 {
		v214 = v179
		goto L46
	} else {
		goto L55
	}
L55:
	;
	v188 = v172 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v188
	v172 = v188
	goto L53
L56:
	;
	if v191 == int32(0) {
		v280 = v190
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L35
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(320116)
	F_errmsg(m, int32(706670), v12+int32(32))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L35
	} else {
		goto L59
	}
L59:
	;
	F_errsave_finish(m, v14, int32(490692), int32(1455), int32(277992))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L35
	} else {
		goto L60
	}
L60:
	;
	v280 = v190
	goto L1
L61:
	;
	v222 = int32(0)
	v223 = F_errsave_start(m, v14)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L35
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+47)))
	*(*int32)(unsafe.Add(mBase, uint32(v139)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v139)+8)) = v243 ^ int32(1)
	v280 = v139
	goto L1
L64:
	;
	if v223 == int32(0) {
		v280 = v222
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L35
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(320116)
	F_errmsg(m, int32(706670), v12+int32(16))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L35
	} else {
		goto L67
	}
L67:
	;
	F_errsave_finish(m, v14, int32(490692), int32(1463), int32(277992))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L35
	} else {
		goto L68
	}
L68:
	;
	v280 = v222
	goto L1
L69:
	;
	if v258 == int32(0) {
		v280 = v2
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L35
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(320116)
	F_errmsg(m, int32(706670), v12)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L35
	} else {
		goto L72
	}
L72:
	;
	F_errsave_finish(m, v14, int32(490692), int32(1418), int32(277992))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L35
	} else {
		goto L73
	}
L73:
	;
	v280 = v2
	goto L1
}
func F_path_n_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = F_pg_detoast_datum(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			return base.B2i32(v11 == v12)
		}
	}
}
func F_path_sub_pt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v77 int32
	_ = v77
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_copy(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L18
	}
L2:
	;
	return int32(0)
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if int32(0) < v15 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v23 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	return v11
L7:
	;
	v32 = v11 + int32(16) + v23<<(uint(int32(4))%32)
	v33 = *(*float64)(unsafe.Add(mBase, uint32(v32)))
	v34 = *(*float64)(unsafe.Add(mBase, uint32(v18)))
	v35 = base.F64_sub(v33, v34)
	if base.F64_ne(base.F64_abs(v35), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v45 = *(*float64)(unsafe.Add(mBase, uint32(v32)+8))
	v46 = *(*float64)(unsafe.Add(mBase, uint32(v18)+8))
	v47 = base.F64_sub(v45, v46)
	if base.F64_ne(base.F64_abs(v47), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	if base.F64_eq(base.F64_abs(v33), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	if base.F64_ne(base.F64_abs(v34), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v32)+8)) = v47
	*(*float64)(unsafe.Add(mBase, uint32(v32))) = v35
	v60 = v23 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v60 < v61 {
		v23 = v60
		goto L7
	} else {
		goto L17
	}
L14:
	;
	if base.F64_eq(base.F64_abs(v45), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if base.F64_ne(base.F64_abs(v46), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	goto L8
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_setPath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
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
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v149 int32
	_ = v149
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v232 int32
	_ = v232
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v321 int32
	_ = v321
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v373 int32
	_ = v373
	var v384 int32
	_ = v384
	var v402 int32
	_ = v402
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v432 int32
	_ = v432
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v516 int32
	_ = v516
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v570 int32
	_ = v570
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v686 int32
	_ = v686
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v889 int32
	_ = v889
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v941 int32
	_ = v941
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v976 int32
	_ = v976
	var v992 int32
	_ = v992
	var v1004 int32
	_ = v1004
	var v1021 int32
	_ = v1021
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1071 int32
	_ = v1071
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1133 int32
	_ = v1133
	var v1141 int32
	_ = v1141
	var v1144 int32
	_ = v1144
	var v1150 int32
	_ = v1150
	var v1155 int32
	_ = v1155
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1207 int32
	_ = v1207
	var v1211 int32
	_ = v1211
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1227 int32
	_ = v1227
	var v1231 int32
	_ = v1231
	var v1236 int32
	_ = v1236
	var v1240 int32
	_ = v1240
	var v1246 int32
	_ = v1246
	var v1251 int32
	_ = v1251
	v9 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(144)
	m.G0 = v29
	F_check_stack_depth(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v35 = l2 + l5
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v36 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L1
	} else {
		goto L291
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L1
	} else {
		goto L286
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L1
	} else {
		goto L281
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L1
	} else {
		goto L277
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L1
	} else {
		goto L272
	}
L8:
	;
	v42 = F_JsonbIteratorNext(m, l0, v29-int32(-64), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L15
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L1
	} else {
		goto L268
	}
L11:
	;
	m.G0 = v29 + int32(144)
	return v1133
L12:
	;
	if l7&int32(32) != 0 {
		goto L263
	} else {
		goto L264
	}
L13:
	;
	v607 = F_pushJsonbValue(m, l4, int32(6), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L130
	}
L14:
	;
	v47 = l7 & int32(32)
	v49 = l3 - int32(1)
	if v49 < l5 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	switch v42 - int32(2) {
	case 0, 1:
		goto L12
	case 2:
		goto L14
	default:
		goto L3
	case 4:
		goto L13
	}
L16:
	;
	v58 = F_pushJsonbValue(m, l4, int32(4), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L20
	}
L17:
	;
	if v47 == int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+76)))
	if v53 == int32(1) {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	v61 = base.B2i32(l3 <= l5)
	if l3 <= l5 {
		v82 = v60
		goto L21
	} else {
		goto L22
	}
L21:
	;
	if int32(0) <= v82 {
		v114 = v82
		goto L29
	} else {
		goto L30
	}
L22:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v62 != 0 {
		v82 = v60
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1+l5<<(uint(int32(2))%32))))
	v67 = F_text_to_cstring(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0)
	v75 = F_strtol(m, v67, v29+int32(124), int32(10))
	mBase = m.M
	goto L25
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v29)+124))
	if v67 == v76 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	if v78 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v80 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	v82 = v75
	goto L21
L29:
	;
	v116 = l7 & int32(25)
	v117 = int32(0)
	v119 = base.B2i32(l5 != v49)
	if base.Ui32(v60) < base.Ui32(v114) {
		goto L39
	} else {
		goto L40
	}
L30:
	;
	if base.Ui32(v60) < base.Ui32(int32(0)-v82) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	if base.Ui32(l7) < base.Ui32(int32(64)) {
		v114 = int32(-2147483648)
		goto L29
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v114 = v60 + v82
	goto L29
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = l5 + int32(1)
	F_errmsg(m, int32(483396), v29+int32(32))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(491552), int32(5463), int32(26214))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	v123 = v60
	goto L41
L40:
	;
	v123 = v114
	goto L41
L41:
	;
	if int32(0) < v114 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v126 = v123
	goto L44
L43:
	;
	v126 = v114
	goto L44
L44:
	;
	if v47 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v127 = v114
	goto L47
L46:
	;
	v127 = v126
	goto L47
L47:
	;
	v132 = base.B2i32(v116 == v117) | (v119 | base.B2i32(v60 != v117)&base.B2i32(v127 != int32(-2147483648)))
	if v132 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v60 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v232 = v132 ^ int32(1)
	if v60 != 0 {
		goto L60
	} else {
		goto L61
	}
L51:
	;
	v203 = F_pushJsonbValue(m, l4, int32(3), l6)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L59
	}
L52:
	;
	if v47 == int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	if v127 <= int32(0) {
		goto L51
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+124)) = int32(0)
	v149 = v127
	goto L55
L55:
	;
	v170 = F_pushJsonbValue(m, l4, int32(3), v29+int32(124))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	goto L51
L57:
	;
	v172 = int32(1)
	if base.Ui32(v172) < base.Ui32(v149) {
		v149 = v149 - v172
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	goto L50
L60:
	;
	v254 = int32(0)
	v258 = v232
	goto L63
L61:
	;
	v402 = v232
	goto L62
L62:
	;
	if v116 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L63:
	;
	if l3 <= l5 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v402 = v373
	goto L62
L65:
	;
	v384 = v254 + int32(1)
	if v384 != v60 {
		v254 = v384
		v258 = v373
		goto L63
	} else {
		goto L98
	}
L66:
	;
	v295 = F_JsonbIteratorNext(m, l0, v29+int32(104), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L84
	}
L67:
	;
	if v254 != v127 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	if v119 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v274 = F_JsonbIteratorNext(m, l0, v29+int32(104), int32(1))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v289 = F_setPath(m, l0, l1, l2, l3, l4, l5+int32(1), l6, l7)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L83
	}
L72:
	;
	if l7&int32(9) != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v277 = F_pushJsonbValue(m, l4, int32(3), l6)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	if l7&int32(24) != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L75
L77:
	;
	v281 = F_pushJsonbValue(m, l4, v274, v29+int32(104))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v283 = int32(1)
	if l7&int32(20) == int32(0) {
		v373 = v283
		goto L65
	} else {
		goto L81
	}
L80:
	;
	goto L79
L81:
	;
	v287 = F_pushJsonbValue(m, l4, int32(3), l6)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v373 = v283
	goto L65
L83:
	;
	v373 = int32(1)
	goto L65
L84:
	;
	if base.Ui32(v295) < base.Ui32(int32(4)) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v302 = v29 + int32(104)
	goto L87
L86:
	;
	v302 = int32(0)
	goto L87
L87:
	;
	v303 = F_pushJsonbValue(m, l4, v295, v302)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	if v295&int32(-3) != int32(4) {
		v373 = v258
		goto L65
	} else {
		goto L89
	}
L89:
	;
	v321 = int32(1)
	goto L90
L90:
	;
	v339 = F_JsonbIteratorNext(m, l0, v29+int32(104), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L92
	}
L91:
	;
	v373 = v258
	goto L65
L92:
	;
	if base.Ui32(v339) < base.Ui32(int32(4)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v346 = v29 + int32(104)
	goto L95
L94:
	;
	v346 = int32(0)
	goto L95
L95:
	;
	v347 = F_pushJsonbValue(m, l4, v339, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v350 = v339 & int32(-3)
	v356 = v321 + base.B2i32(v350 == int32(4)) - base.B2i32(v350 == int32(5))
	if v356 != 0 {
		v321 = v356
		goto L90
	} else {
		goto L97
	}
L97:
	;
	goto L91
L98:
	;
	goto L64
L99:
	;
	v600 = F_JsonbIteratorNext(m, l0, v29-int32(-64), int32(0))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L128
	}
L100:
	;
	if v127 <= int32(0) {
		goto L120
	} else {
		goto L121
	}
L101:
	;
	if v49 <= l5 {
		goto L99
	} else {
		goto L118
	}
L102:
	;
	if v47 == int32(0) {
		goto L99
	} else {
		goto L115
	}
L103:
	;
	if v402&int32(1) != 0 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	if l5 != v49 {
		goto L101
	} else {
		goto L105
	}
L105:
	;
	if v47 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v486 = F_pushJsonbValue(m, l4, int32(3), l6)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L114
	}
L107:
	;
	if base.Ui32(v127) <= base.Ui32(v60) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v419 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+124)) = v419
	v421 = v127 - v60
	if v421 <= v419 {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v432 = v421
	goto L110
L110:
	;
	v453 = F_pushJsonbValue(m, l4, int32(3), v29+int32(124))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L112
	}
L111:
	;
	goto L106
L112:
	;
	v455 = int32(1)
	if base.Ui32(v455) < base.Ui32(v432) {
		v432 = v432 - v455
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	goto L99
L115:
	;
	if v49 <= l5 {
		goto L99
	} else {
		goto L116
	}
L116:
	;
	v491 = int32(0)
	if (base.B2i32(v116 != v491)|v402)&int32(1) == v491 {
		goto L100
	} else {
		goto L117
	}
L117:
	;
	goto L99
L118:
	;
	if v47 == int32(0) {
		goto L99
	} else {
		goto L119
	}
L119:
	;
	goto L100
L120:
	;
	F_push_path(m, l4, l5, l1, l2, l3, l6)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L127
	}
L121:
	;
	v503 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+124)) = v503
	v505 = v127 - v60
	if v505 <= v503 {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v516 = v505
	goto L123
L123:
	;
	v537 = F_pushJsonbValue(m, l4, int32(3), v29+int32(124))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L125
	}
L124:
	;
	goto L120
L125:
	;
	v539 = int32(1)
	if base.Ui32(v539) < base.Ui32(v516) {
		v516 = v516 - v539
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	goto L99
L128:
	;
	v603 = F_pushJsonbValue(m, l4, v600, int32(0))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v1133 = v603
	goto L11
L130:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	v610 = int32(1)
	if l3 <= l5 {
		v620 = v9
		v621 = v610
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v623 = l7 & int32(25)
	v625 = l3 - int32(1)
	v626 = base.B2i32(l5 != v625)
	if l5 != v625 {
		goto L136
	} else {
		goto L137
	}
L132:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v612 != 0 {
		v620 = v9
		v621 = v610
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(l1+l5<<(uint(int32(2))%32))))
	v618 = F_pg_detoast_datum_packed(m, v617)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v620 = v618
	v621 = int32(0)
	goto L131
L135:
	;
	if v625 <= l5 {
		goto L241
	} else {
		goto L242
	}
L136:
	;
	if v609 == int32(0) {
		v1021 = v621
		goto L135
	} else {
		goto L156
	}
L137:
	;
	if v623 == int32(0) {
		goto L136
	} else {
		goto L138
	}
L138:
	;
	if v609 != 0 {
		goto L136
	} else {
		goto L139
	}
L139:
	;
	v629 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+124)) = v629
	v632 = v620 + v629
	v635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620))))
	if v635&v629 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v638 = v632
	goto L142
L141:
	;
	v638 = v620 + int32(4)
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+132)) = v638
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620))))
	if v640 == int32(1) {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+128)) = v670
	v675 = F_pushJsonbValue(m, l4, int32(1), v29+int32(124))
	mBase = m.M
	v676 = m.ExcPending
	if v676 != 0 {
		goto L1
	} else {
		goto L154
	}
L144:
	;
	v643 = int32(4)
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632))))
	if v645&int32(254) == int32(2) {
		goto L147
	} else {
		goto L148
	}
L145:
	;
	goto L146
L146:
	;
	v658 = int32(1)
	if v640&v658 != 0 {
		v670 = int32(base.Ui32(v640)>>(uint(v658)%32)) - v658
		goto L143
	} else {
		goto L153
	}
L147:
	;
	v654 = v643
	goto L149
L148:
	;
	v654 = base.B2i32(v645 == int32(18)) << (uint(v643) % 32)
	goto L149
L149:
	;
	if v645 == int32(1) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v657 = v643
	goto L152
L151:
	;
	v657 = v654
	goto L152
L152:
	;
	v670 = v657
	goto L143
L153:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	v670 = int32(base.Ui32(v664)>>(uint(int32(2))%32)) - int32(4)
	goto L143
L154:
	;
	v678 = F_pushJsonbValue(m, l4, int32(2), l6)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v1021 = v621
	goto L135
L156:
	;
	v686 = int32(1)
	v693 = v620 + v686
	v695 = int32(0)
	v712 = v9
	v715 = v621
	goto L157
L157:
	;
	v729 = F_JsonbIteratorNext(m, l0, v29+int32(124), int32(1))
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L159
	}
L158:
	;
	v1021 = v992
	goto L135
L159:
	;
	if v715 != 0 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	v1004 = v712 + int32(1)
	if v1004 != v609 {
		v712 = v1004
		v715 = v992
		goto L157
	} else {
		goto L240
	}
L161:
	;
	v910 = F_pushJsonbValue(m, l4, v729, v29+int32(124))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L1
	} else {
		goto L225
	}
L162:
	;
	v731 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620))))
	if v731 == int32(1) {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	if base.B2i32(v712 != v609-v686)|base.B2i32(base.B2i32(l5 == v625)&base.B2i32(v623 != v695) == v695) != 0 {
		goto L161
	} else {
		goto L208
	}
L164:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v29)+128))
	if v761 != v762 {
		goto L163
	} else {
		goto L175
	}
L165:
	;
	v734 = int32(4)
	v736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v693))))
	if v736&int32(254) == int32(2) {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	goto L167
L167:
	;
	v749 = int32(1)
	if v731&v749 != 0 {
		v761 = int32(base.Ui32(v731)>>(uint(v749)%32)) - v749
		goto L164
	} else {
		goto L174
	}
L168:
	;
	v745 = v734
	goto L170
L169:
	;
	v745 = base.B2i32(v736 == int32(18)) << (uint(v734) % 32)
	goto L170
L170:
	;
	if v736 == int32(1) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v748 = v734
	goto L173
L172:
	;
	v748 = v745
	goto L173
L173:
	;
	v761 = v748
	goto L164
L174:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	v761 = int32(base.Ui32(v755)>>(uint(int32(2))%32)) - int32(4)
	goto L164
L175:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v29)+132))
	v765 = int32(1)
	if v731&v765 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v769 = v765
	goto L178
L177:
	;
	v769 = int32(4)
	goto L178
L178:
	;
	v770 = v620 + v769
	if base.Ui32(int32(4)) <= base.Ui32(v762) {
		goto L182
	} else {
		goto L183
	}
L179:
	;
	if v832 != 0 {
		goto L163
	} else {
		goto L197
	}
L180:
	;
	v832 = int32(0)
	goto L179
L181:
	;
	v806 = v801
	v807 = v802
	v808 = v803
	goto L191
L182:
	;
	if (v764|v770)&int32(3) != 0 {
		v801 = v764
		v802 = v770
		v803 = v762
		goto L181
	} else {
		goto L185
	}
L183:
	;
	v794 = v764
	v795 = v770
	v796 = v762
	goto L184
L184:
	;
	if v796 == int32(0) {
		goto L180
	} else {
		goto L190
	}
L185:
	;
	v778 = v764
	v779 = v770
	v780 = v762
	goto L186
L186:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v778)))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v779)))
	if v783 != v784 {
		v801 = v778
		v802 = v779
		v803 = v780
		goto L181
	} else {
		goto L188
	}
L187:
	;
	v794 = v789
	v795 = v787
	v796 = v791
	goto L184
L188:
	;
	v786 = int32(4)
	v787 = v779 + v786
	v789 = v778 + v786
	v791 = v780 - v786
	if base.Ui32(int32(3)) < base.Ui32(v791) {
		v778 = v789
		v779 = v787
		v780 = v791
		goto L186
	} else {
		goto L189
	}
L189:
	;
	goto L187
L190:
	;
	v801 = v794
	v802 = v795
	v803 = v796
	goto L181
L191:
	;
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v806))))
	v812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v807))))
	if v811 == v812 {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	v832 = v811 - v812
	goto L179
L193:
	;
	v814 = int32(1)
	v819 = v808 - v814
	if v819 != 0 {
		v806 = v806 + v814
		v807 = v807 + v814
		v808 = v819
		goto L191
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	goto L192
L196:
	;
	goto L180
L197:
	;
	if v626 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	if l7&int32(24) != 0 {
		goto L5
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	v851 = F_pushJsonbValue(m, l4, v729, v29+int32(124))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L1
	} else {
		goto L206
	}
L201:
	;
	v835 = int32(1)
	v839 = F_JsonbIteratorNext(m, l0, v29+int32(104), v835)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	if l7&int32(2) != 0 {
		v992 = v835
		goto L160
	} else {
		goto L203
	}
L203:
	;
	v844 = F_pushJsonbValue(m, l4, int32(1), v29+int32(124))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v847 = F_pushJsonbValue(m, l4, int32(2), l6)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	v992 = v835
	goto L160
L206:
	;
	v853 = F_setPath(m, l0, l1, l2, l3, l4, l5+v686, l6, l7)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	v992 = int32(1)
	goto L160
L208:
	;
	v858 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+84)) = v858
	v860 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620))))
	if v860&v858 != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v863 = v693
	goto L211
L210:
	;
	v863 = v620 + int32(4)
	goto L211
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+92)) = v863
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620))))
	if v865 == int32(1) {
		goto L213
	} else {
		goto L214
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+88)) = v895
	v900 = F_pushJsonbValue(m, l4, int32(1), v29+int32(84))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L1
	} else {
		goto L223
	}
L213:
	;
	v868 = int32(4)
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v693))))
	if v870&int32(254) == int32(2) {
		goto L216
	} else {
		goto L217
	}
L214:
	;
	goto L215
L215:
	;
	v883 = int32(1)
	if v865&v883 != 0 {
		v895 = int32(base.Ui32(v865)>>(uint(v883)%32)) - v883
		goto L212
	} else {
		goto L222
	}
L216:
	;
	v879 = v868
	goto L218
L217:
	;
	v879 = base.B2i32(v870 == int32(18)) << (uint(v868) % 32)
	goto L218
L218:
	;
	if v870 == int32(1) {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v882 = v868
	goto L221
L220:
	;
	v882 = v879
	goto L221
L221:
	;
	v895 = v882
	goto L212
L222:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	v895 = int32(base.Ui32(v889)>>(uint(int32(2))%32)) - int32(4)
	goto L212
L223:
	;
	v903 = F_pushJsonbValue(m, l4, int32(2), l6)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	goto L161
L225:
	;
	v915 = F_JsonbIteratorNext(m, l0, v29+int32(104), int32(0))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	if base.Ui32(v915) < base.Ui32(int32(4)) {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v922 = v29 + int32(104)
	goto L229
L228:
	;
	v922 = int32(0)
	goto L229
L229:
	;
	v923 = F_pushJsonbValue(m, l4, v915, v922)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	if v915&int32(-3) != int32(4) {
		v992 = v715
		goto L160
	} else {
		goto L231
	}
L231:
	;
	v941 = int32(1)
	goto L232
L232:
	;
	v959 = F_JsonbIteratorNext(m, l0, v29+int32(104), int32(0))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L1
	} else {
		goto L234
	}
L233:
	;
	v992 = v715
	goto L160
L234:
	;
	if base.Ui32(v959) < base.Ui32(int32(4)) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v966 = v29 + int32(104)
	goto L237
L236:
	;
	v966 = int32(0)
	goto L237
L237:
	;
	v967 = F_pushJsonbValue(m, l4, v959, v966)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	v970 = v959 & int32(-3)
	v976 = v941 + base.B2i32(v970 == int32(4)) - base.B2i32(v970 == int32(5))
	if v976 != 0 {
		v941 = v976
		goto L232
	} else {
		goto L239
	}
L239:
	;
	goto L233
L240:
	;
	goto L158
L241:
	;
	v1091 = F_JsonbIteratorNext(m, l0, v29-int32(-64), int32(1))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L1
	} else {
		goto L261
	}
L242:
	;
	if l7&int32(32) == int32(0) {
		goto L241
	} else {
		goto L243
	}
L243:
	;
	if v1021 != 0 {
		goto L241
	} else {
		goto L244
	}
L244:
	;
	v1037 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+84)) = v1037
	v1040 = v620 + v1037
	v1043 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620))))
	if v1043&v1037 != 0 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v1046 = v1040
	goto L247
L246:
	;
	v1046 = v620 + int32(4)
	goto L247
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+92)) = v1046
	if v1043 == int32(1) {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+88)) = v1077
	v1082 = F_pushJsonbValue(m, l4, int32(1), v29+int32(84))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L1
	} else {
		goto L259
	}
L249:
	;
	v1050 = int32(4)
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1040))))
	if v1052&int32(254) == int32(2) {
		goto L252
	} else {
		goto L253
	}
L250:
	;
	goto L251
L251:
	;
	v1065 = int32(1)
	if v1043&v1065 != 0 {
		v1077 = int32(base.Ui32(v1043)>>(uint(v1065)%32)) - v1065
		goto L248
	} else {
		goto L258
	}
L252:
	;
	v1061 = v1050
	goto L254
L253:
	;
	v1061 = base.B2i32(v1052 == int32(18)) << (uint(v1050) % 32)
	goto L254
L254:
	;
	if v1052 == int32(1) {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v1064 = v1050
	goto L257
L256:
	;
	v1064 = v1061
	goto L257
L257:
	;
	v1077 = v1064
	goto L248
L258:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v620)))
	v1077 = int32(base.Ui32(v1071)>>(uint(int32(2))%32)) - int32(4)
	goto L248
L259:
	;
	F_push_path(m, l4, l5, l1, l2, l3, l6)
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	goto L241
L261:
	;
	v1094 = F_pushJsonbValue(m, l4, v1091, int32(0))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	v1133 = v1094
	goto L11
L263:
	;
	v1102 = base.B2i32(l5 <= l3-int32(1))
	goto L265
L264:
	;
	v1102 = int32(0)
	goto L265
L265:
	;
	if v1102 != 0 {
		goto L4
	} else {
		goto L266
	}
L266:
	;
	v1105 = F_pushJsonbValue(m, l4, v42, v29-int32(-64))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	v1133 = v1105
	goto L11
L268:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = l5 + int32(1)
	F_errmsg(m, int32(300395), v29)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	F_errfinish(m, int32(491552), int32(5218), int32(320134))
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L272:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	F_errmsg(m, int32(22258), int32(0))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	F_errdetail(m, int32(610252), int32(0))
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L1
	} else {
		goto L275
	}
L275:
	;
	F_errfinish(m, int32(491552), int32(5238), int32(320134))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L277:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = l5 + int32(1)
	F_errmsg(m, int32(707384), v29+int32(48))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	F_errfinish(m, int32(491552), int32(5446), int32(26214))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L281:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1203 = m.ExcPending
	if v1203 != 0 {
		goto L1
	} else {
		goto L282
	}
L282:
	;
	F_errmsg(m, int32(22258), int32(0))
	mBase = m.M
	v1207 = m.ExcPending
	if v1207 != 0 {
		goto L1
	} else {
		goto L283
	}
L283:
	;
	F_errhint(m, int32(610125), int32(0))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L1
	} else {
		goto L284
	}
L284:
	;
	F_errfinish(m, int32(491552), int32(5342), int32(110907))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L286:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	F_errmsg(m, int32(22258), int32(0))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	F_errdetail(m, int32(610252), int32(0))
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	F_errfinish(m, int32(491552), int32(5269), int32(320134))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v42
	F_errmsg_internal(m, int32(478309), v29+int32(16))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L1
	} else {
		goto L292
	}
L292:
	;
	F_errfinish(m, int32(491552), int32(5274), int32(320134))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_substitute_path_macro(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 != int32(36) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L36
	}
L2:
	;
	m.G0 = v9 + int32(32)
	return v89
L3:
	;
	v14 = F_pstrdup(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v19 = l0
	goto L9
L6:
	;
	return int32(0)
L7:
	;
	v89 = v14
	goto L2
L8:
	;
	if v29 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v21 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L8
L11:
	;
	goto L10
L12:
	;
	v29 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	if v21 == int32(47) {
		v29 = v19
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v19 = v19 + int32(1)
	goto L9
L16:
	;
	v32 = F_strlen(m, l0)
	mBase = m.M
	v34 = v32 + l0
	goto L18
L17:
	;
	v34 = v29
	goto L18
L18:
	;
	v35 = F_strlen(m, l1)
	mBase = m.M
	if v35 != v34-l0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v35 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v81 != 0 {
		goto L1
	} else {
		goto L34
	}
L21:
	;
	v81 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v43 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v44 = l0
	v45 = l1
	v46 = v35
	v47 = v43
	goto L28
L25:
	;
	v69 = l1
	v73 = int32(0)
	goto L26
L26:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	v81 = v73 - v74
	goto L20
L27:
	;
	v69 = v64
	v73 = v66
	goto L26
L28:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v47 != v49 {
		v64 = v45
		v66 = v47
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v64 = v58
	v66 = int32(0)
	goto L27
L30:
	;
	if v49 == int32(0) {
		v64 = v45
		v66 = v47
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v54 = v46 - int32(1)
	if v54 == int32(0) {
		v64 = v45
		v66 = v47
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v57 = int32(1)
	v58 = v45 + v57
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if v59 != 0 {
		v44 = v44 + v57
		v45 = v58
		v46 = v54
		v47 = v59
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
	v85 = F_psprintf(m, int32(174963), v9)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v89 = v85
	goto L2
L36:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
	F_errmsg(m, int32(201807), v9+int32(16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(492034), int32(555), int32(239085))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
